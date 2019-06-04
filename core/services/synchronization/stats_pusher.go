package synchronization

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jpillora/backoff"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/store/models"
	"github.com/smartcontractkit/chainlink/core/store/orm"
	"github.com/smartcontractkit/chainlink/core/utils"
)

// StatsPusher polls for events and pushes them via a WebSocketClient
type StatsPusher struct {
	ORM            *orm.ORM
	WSClient       WebSocketClient
	Period         time.Duration
	cancel         context.CancelFunc
	clock          utils.Afterer
	backoffSleeper backoff.Backoff
	waker          chan struct{}
}

const (
	createCallbackName = "sync:run_after_create"
	updateCallbackName = "sync:run_after_update"
)

// NewStatsPusher returns a new event queuer
func NewStatsPusher(orm *orm.ORM, url *url.URL, accessKey, secret string, afters ...utils.Afterer) *StatsPusher {
	var clock utils.Afterer
	if len(afters) == 0 {
		clock = utils.Clock{}
	} else {
		clock = afters[0]
	}

	sp := &StatsPusher{
		ORM:      orm,
		WSClient: noopWebSocketClient{},
		Period:   30 * time.Minute,
		clock:    clock,
		backoffSleeper: backoff.Backoff{
			Min: 1 * time.Second,
			Max: 5 * time.Minute,
		},
		waker: make(chan struct{}, 1),
	}

	if url != nil {
		sp.WSClient = NewWebSocketClient(url, accessKey, secret)
		orm.DB.Callback().Create().Register(createCallbackName, createSyncEventWithStatsPusher(sp))
		orm.DB.Callback().Update().Register(updateCallbackName, createSyncEventWithStatsPusher(sp))
	}
	return sp
}

// Start starts the stats pusher
func (sp *StatsPusher) Start() error {
	err := sp.WSClient.Start()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	sp.cancel = cancel
	go sp.eventLoop(ctx)
	return nil
}

// Close shuts down the stats pusher
func (sp *StatsPusher) Close() error {
	if sp.cancel != nil {
		sp.cancel()
	}
	sp.ORM.DB.Callback().Create().Remove(createCallbackName)
	sp.ORM.DB.Callback().Update().Remove(updateCallbackName)
	return sp.WSClient.Close()
}

// PushNow wakes up the stats pusher, asking it to push all queued events immediately.
func (sp *StatsPusher) PushNow() {
	select {
	case sp.waker <- struct{}{}:
	default:
	}
}

type response struct {
	Status int `json:"status"`
}

func (sp *StatsPusher) syncEvent(event *models.SyncEvent) error {
	sp.WSClient.Send([]byte(event.Body))

	message, err := sp.WSClient.Receive()
	if err != nil {
		return err
	}

	var response response
	err = json.Unmarshal(message, &response)
	if err != nil {
		return err
	}

	if response.Status != 201 {
		return errors.New("event not created")
	}

	err = sp.ORM.DB.Delete(event).Error
	if err != nil {
		return err
	}

	return nil
}

func (sp *StatsPusher) eventLoop(parentCtx context.Context) {
	for {
		err := sp.pusherLoop(parentCtx)
		if err == nil {
			return
		}

		duration := sp.backoffSleeper.Duration()
		logger.Errorw("Error during event synchronization", "error", err, "sleep_duration", duration)

		select {
		case <-parentCtx.Done():
			return
		case <-sp.clock.After(duration):
			continue
		}
	}
}

func (sp *StatsPusher) pushEvents() error {
	err := sp.ORM.AllSyncEvents(func(event *models.SyncEvent) error {
		logger.Debugw("StatsPusher got event", "event", event.ID)
		return sp.syncEvent(event)
	})

	if err != nil {
		return err
	}

	sp.backoffSleeper.Reset()
	return nil
}

func (sp *StatsPusher) pusherLoop(parentCtx context.Context) error {
	logger.Debugw("Entered main pusher loop")

	for {
		select {
		case <-parentCtx.Done():
			logger.Debugw("StatsPusher got done signal, shutting down")
			return nil
		case <-sp.waker:
			err := sp.pushEvents()
			if err != nil {
				return err
			}
		case <-sp.clock.After(sp.Period):
			err := sp.pushEvents()
			if err != nil {
				return err
			}
		}
	}
}

func createSyncEventWithStatsPusher(sp *StatsPusher) func(*gorm.Scope) {
	return func(scope *gorm.Scope) {
		if scope.HasError() {
			return
		}

		if scope.TableName() != "job_runs" {
			return
		}

		run, ok := scope.Value.(*models.JobRun)
		if !ok {
			return
		}

		presenter := SyncJobRunPresenter{run}
		bodyBytes, err := json.Marshal(presenter)
		if err != nil {
			scope.Err(err)
			return
		}

		event := models.SyncEvent{
			Body: string(bodyBytes),
		}
		err = scope.DB().Save(&event).Error
		if err != nil {
			scope.Err(err)
			return
		}

		sp.PushNow()
	}
}
