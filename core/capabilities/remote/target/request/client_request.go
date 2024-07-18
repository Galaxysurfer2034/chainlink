package request

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"

	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	commoncap "github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/pb"
	"github.com/smartcontractkit/chainlink/v2/core/capabilities/remote"
	"github.com/smartcontractkit/chainlink/v2/core/capabilities/remote/types"
	"github.com/smartcontractkit/chainlink/v2/core/capabilities/transmission"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	p2ptypes "github.com/smartcontractkit/chainlink/v2/core/services/p2p/types"
)

type ClientRequest struct {
	cancelFn         context.CancelFunc
	responseCh       chan commoncap.CapabilityResponse
	createdAt        time.Time
	responseIDCount  map[[32]byte]int
	errorCount       map[string]int
	responseReceived map[p2ptypes.PeerID]bool
	lggr             logger.Logger

	requiredIdenticalResponses int

	requestTimeout time.Duration

	respSent bool
	mux      sync.Mutex
	wg       *sync.WaitGroup
}

func NewClientRequest(ctx context.Context, lggr logger.Logger, req commoncap.CapabilityRequest, messageID string,
	remoteCapabilityInfo commoncap.CapabilityInfo, localDonInfo capabilities.DON, dispatcher types.Dispatcher,
	requestTimeout time.Duration) (*ClientRequest, error) {
	remoteCapabilityDonInfo := remoteCapabilityInfo.DON
	if remoteCapabilityDonInfo == nil {
		return nil, errors.New("remote capability info missing DON")
	}

	lggr = lggr.Named("ClientRequest-" + uuid.New().String())

	lggr.Debugw("new client request for capability request", "request", req)
	rawRequest, err := proto.MarshalOptions{Deterministic: true}.Marshal(pb.CapabilityRequestToProto(req))

	hash := sha256.Sum256(rawRequest)
	requestID := messageID + hex.EncodeToString(hash[:])
	lggr.Debugw("new raw request for capability request", "request ID", requestID)

	if err != nil {
		return nil, fmt.Errorf("failed to marshal capability request: %w", err)
	}

	peerIDToTransmissionDelay, err := transmission.GetPeerIDToTransmissionDelay(remoteCapabilityDonInfo.Members, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get peer ID to transmission delay: %w", err)
	}

	lggr.Debugw("sending request to peers", "execID", req.Metadata.WorkflowExecutionID, "schedule", peerIDToTransmissionDelay)

	responseReceived := make(map[p2ptypes.PeerID]bool)

	ctxWithCancel, cancelFn := context.WithCancel(ctx)
	wg := &sync.WaitGroup{}
	for peerID, delay := range peerIDToTransmissionDelay {
		responseReceived[peerID] = false
		wg.Add(1)
		go func(ctx context.Context, peerID ragep2ptypes.PeerID, delay time.Duration) {
			defer wg.Done()
			message := &types.MessageBody{
				CapabilityId:    remoteCapabilityInfo.ID,
				CapabilityDonId: remoteCapabilityDonInfo.ID,
				CallerDonId:     localDonInfo.ID,
				Method:          types.MethodExecute,
				Payload:         rawRequest,
				MessageId:       []byte(messageID),
			}

			select {
			case <-ctxWithCancel.Done():
				lggr.Debugw("context done, not sending request to peer", "execID", req.Metadata.WorkflowExecutionID, "peerID", peerID)
				return
			case <-time.After(delay):
				lggr.Debugw("sending request to peer", "execID", req.Metadata.WorkflowExecutionID, "peerID", peerID)
				err := dispatcher.Send(peerID, message)
				if err != nil {
					lggr.Errorw("failed to send message", "peerID", peerID, "err", err)
				}
			}
		}(ctxWithCancel, peerID, delay)
	}

	return &ClientRequest{
		cancelFn:                   cancelFn,
		createdAt:                  time.Now(),
		requestTimeout:             requestTimeout,
		requiredIdenticalResponses: int(remoteCapabilityDonInfo.F + 1),
		responseIDCount:            make(map[[32]byte]int),
		errorCount:                 make(map[string]int),
		responseReceived:           responseReceived,
		responseCh:                 make(chan commoncap.CapabilityResponse, 1),
		wg:                         wg,
		lggr:                       lggr,
	}, nil
}

func (c *ClientRequest) ResponseChan() <-chan commoncap.CapabilityResponse {
	return c.responseCh
}

func (c *ClientRequest) Expired() bool {
	return time.Since(c.createdAt) > c.requestTimeout
}

func (c *ClientRequest) Cancel(err error) {
	c.cancelFn()
	c.wg.Wait()
	c.mux.Lock()
	defer c.mux.Unlock()
	if !c.respSent {
		c.sendResponse(commoncap.CapabilityResponse{Err: err})
	}
}

// TODO OnMessage assumes that only one response is received from each peer, if streaming responses need to be supported this will need to be updated
func (c *ClientRequest) OnMessage(_ context.Context, msg *types.MessageBody) error {
	c.mux.Lock()
	defer c.mux.Unlock()

	if c.respSent {
		return nil
	}

	if msg.Sender == nil {
		return fmt.Errorf("sender missing from message")
	}

	c.lggr.Debugw("OnMessage called for client request", "messageID", msg.MessageId)

	sender := remote.ToPeerID(msg.Sender)

	received, expected := c.responseReceived[sender]
	if !expected {
		return fmt.Errorf("response from peer %s not expected", sender)
	}

	if received {
		return fmt.Errorf("response from peer %s already received", sender)
	}

	c.responseReceived[sender] = true

	if msg.Error == types.Error_OK {
		responseID := sha256.Sum256(msg.Payload)
		c.responseIDCount[responseID]++

		if len(c.responseIDCount) > 1 {
			c.lggr.Warn("received multiple different responses for the same request, number of different responses received: %d", len(c.responseIDCount))
		}

		if c.responseIDCount[responseID] == c.requiredIdenticalResponses {
			capabilityResponse, err := pb.UnmarshalCapabilityResponse(msg.Payload)
			if err != nil {
				c.sendResponse(commoncap.CapabilityResponse{Err: fmt.Errorf("failed to unmarshal capability response: %w", err)})
			} else {
				c.sendResponse(commoncap.CapabilityResponse{Value: capabilityResponse.Value})
			}
		}
	} else {
		c.lggr.Warnw("received error response", "error", msg.ErrorMsg)
		c.errorCount[msg.ErrorMsg]++
		if c.errorCount[msg.ErrorMsg] == c.requiredIdenticalResponses {
			c.sendResponse(commoncap.CapabilityResponse{Err: errors.New(msg.ErrorMsg)})
		}
	}
	return nil
}

func (c *ClientRequest) sendResponse(response commoncap.CapabilityResponse) {
	c.responseCh <- response
	close(c.responseCh)
	c.respSent = true
}
