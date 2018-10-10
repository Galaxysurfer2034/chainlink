package services

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/smartcontractkit/chainlink/adapters"
	"github.com/smartcontractkit/chainlink/logger"
	"github.com/smartcontractkit/chainlink/store"
	"github.com/smartcontractkit/chainlink/store/assets"
	"github.com/smartcontractkit/chainlink/store/models"
	"github.com/smartcontractkit/chainlink/utils"
)

// ExecuteJob saves and immediately begins executing a run for a specified job
// if it is ready.
func ExecuteJob(
	job models.JobSpec,
	initiator models.Initiator,
	input models.RunResult,
	creationHeight *big.Int,
	store *store.Store) (models.JobRun, error) {

	logger.Debugw(fmt.Sprintf("New run triggered by %s", initiator.Type), []interface{}{
		"job", job.ID,
		"input_status", input.Status,
		"creation_height", creationHeight,
	}...)

	run, err := NewRun(job, initiator, input, creationHeight, store)
	if err != nil {
		return models.JobRun{}, err
	}

	return run, saveAndTrigger(&run, store)
}

// NewRun returns a run from an input job, in an initial state ready for
// processing by the job runner system
func NewRun(
	job models.JobSpec,
	initiator models.Initiator,
	input models.RunResult,
	currentHeight *big.Int,
	store *store.Store) (models.JobRun, error) {

	now := store.Clock.Now()
	if !job.Started(now) {
		return models.JobRun{}, RecurringScheduleJobError{
			msg: fmt.Sprintf("Job runner: Job %v unstarted: %v before job's start time %v", job.ID, now, job.EndAt),
		}
	}

	if job.Ended(now) {
		return models.JobRun{}, RecurringScheduleJobError{
			msg: fmt.Sprintf("Job runner: Job %v ended: %v past job's end time %v", job.ID, now, job.EndAt),
		}
	}

	run := job.NewRun(initiator)

	run.Overrides = input
	run.ApplyResult(input)

	if currentHeight != nil {
		creationHeightHex := hexutil.Big(*currentHeight)
		run.CreationHeight = &creationHeightHex
		run.ObservedHeight = &creationHeightHex
	}

	cost := assets.NewLink(0)
	for i, taskRun := range run.TaskRuns {
		adapter, err := adapters.For(taskRun.Task, store)

		if err != nil {
			run.ApplyResult(run.Result.WithError(err))
			return run, nil
		}

		mp := adapter.MinContractPayment()
		cost.Add(cost, &mp)

		if currentHeight != nil {
			run.TaskRuns[i].MinimumConfirmations = utils.MaxUint64(
				store.Config.MinIncomingConfirmations,
				taskRun.Task.Confirmations,
				adapter.MinConfs())
		}
	}

	// input.Amount is always present for runs triggered by ethlogs
	if input.Amount != nil {
		if cost.Cmp(input.Amount) > 0 {
			logger.Debugw("Rejecting run with insufficient paymen", []interface{}{
				"run", run.ID,
				"job", run.JobID,
				"input_amount", input.Amount,
				"required_amount", cost,
			}...)

			err := fmt.Errorf(
				"Rejecting job %s with payment %s below minimum threshold (%s)",
				job.ID,
				input.Amount,
				store.Config.MinimumContractPayment.Text(10))
			run.ApplyResult(input.WithError(err))
			return run, nil
		}
	}

	initialTask := run.TaskRuns[0]
	if meetsMinimumConfirmations(&run, &initialTask, run.CreationHeight) {
		run.Status = models.RunStatusInProgress
	} else {
		logger.Debugw("Insufficient confirmations to begin job", run.ForLogger()...)
		run.Status = models.RunStatusPendingConfirmations
	}

	return run, nil
}

// ResumeConfirmingTask resumes a confirming run if the minimum confirmations have been met
func ResumeConfirmingTask(
	run *models.JobRun,
	store *store.Store,
	currentBlockHeight *big.Int,
) error {

	logger.Debugw("New head resuming run", run.ForLogger()...)

	if !run.Status.PendingConfirmations() {
		return fmt.Errorf("Attempt to resume non confirming task")
	}

	currentTaskRun := run.NextTaskRun()
	if currentTaskRun == nil {
		return fmt.Errorf("Attempting to resume confirming run with no remaining tasks %s", run.ID)
	}

	observedHeight := hexutil.Big(*currentBlockHeight)
	run.ObservedHeight = &observedHeight

	if meetsMinimumConfirmations(run, currentTaskRun, run.ObservedHeight) {
		logger.Debugw("Minimum confirmations met, resuming job", []interface{}{
			"run", run.ID,
			"job", run.JobID,
			"observed_height", currentBlockHeight,
		}...)
		run.Status = models.RunStatusInProgress
	} else {
		logger.Debugw("Insufficient confirmations to wake job", []interface{}{
			"run", run.ID,
			"job", run.JobID,
			"observed_height", currentBlockHeight,
		}...)
		run.Status = models.RunStatusPendingConfirmations
	}

	return saveAndTrigger(run, store)
}

// ResumePendingTask takes the body provided from an external adapter,
// saves it for the next task to process, then tells the job runner to execute
// it
func ResumePendingTask(
	run *models.JobRun,
	store *store.Store,
	input models.RunResult,
) error {

	logger.Debugw("External adapter resuming job", []interface{}{
		"run", run.ID,
		"job", run.JobID,
		"status", run.Status,
		"input_data", input.Data,
		"input_result", input.Status,
	}...)

	if !run.Status.PendingBridge() {
		return fmt.Errorf("Attempting to resume non pending run %s", run.ID)
	}

	currentTaskRun := run.NextTaskRun()
	if currentTaskRun == nil {
		return fmt.Errorf("Attempting to resume pending run with no remaining tasks %s", run.ID)
	}

	var err error
	if run.Overrides, err = run.Overrides.Merge(input); err != nil {
		currentTaskRun.ApplyResult(input.WithError(err))
		run.ApplyResult(input.WithError(err))
		return store.Save(run)
	}

	currentTaskRun.ApplyResult(input)
	if currentTaskRun.Status.Finished() && run.TasksRemain() {
		run.Status = models.RunStatusInProgress
	} else {
		run.ApplyResult(input)
	}

	return saveAndTrigger(run, store)
}

// QueueSleepingTask creates a go routine which will wake up the job runner
// once the sleep's time has elapsed
func QueueSleepingTask(
	run *models.JobRun,
	store *store.Store,
) error {
	if !run.Status.PendingSleep() {
		return fmt.Errorf("Attempting to resume non sleeping run %s", run.ID)
	}
	currentTaskRun := run.NextTaskRun()
	if currentTaskRun == nil {
		return fmt.Errorf("Attempting to resume sleeping run with no remaining tasks %s", run.ID)
	}

	if !currentTaskRun.Status.PendingSleep() {
		return fmt.Errorf("Attempting to resume sleeping run with non sleeping task %s", run.ID)
	}

	adapter, err := adapters.For(currentTaskRun.Task, store)

	if err != nil {
		currentTaskRun.ApplyResult(run.Result.WithError(err))
		run.ApplyResult(run.Result.WithError(err))
		return store.Save(run)
	}

	if sleepAdapter, ok := adapter.BaseAdapter.(*adapters.Sleep); ok {
		return performTaskSleep(run, currentTaskRun, sleepAdapter, store)
	}

	return fmt.Errorf("Attempting to resume non sleeping task for run %s (%s)", run.ID, currentTaskRun.Task.Type)
}

func performTaskSleep(
	run *models.JobRun,
	task *models.TaskRun,
	adapter *adapters.Sleep,
	store *store.Store) error {

	duration := adapter.Duration()
	if duration <= 0 {
		logger.Debugw("Sleep duration has already elapsed, completing task", run.ForLogger()...)
		task.Status = models.RunStatusCompleted
		run.Status = models.RunStatusInProgress
		return saveAndTrigger(run, store)
	}

	go func() {
		logger.Debugw("Task sleeping...", run.ForLogger()...)

		<-store.Clock.After(duration)

		task.Status = models.RunStatusCompleted
		run.Status = models.RunStatusInProgress

		logger.Debugw("Waking job up after sleep", run.ForLogger()...)

		if err := saveAndTrigger(run, store); err != nil {
			logger.Errorw("Error resuming sleeping job:", "error", err)
		}
	}()

	return nil
}

func meetsMinimumConfirmations(
	run *models.JobRun,
	taskRun *models.TaskRun,
	currentHeight *hexutil.Big) bool {
	if run.CreationHeight == nil {
		return true
	}

	diff := new(big.Int).Sub(currentHeight.ToInt(), run.CreationHeight.ToInt())
	min := new(big.Int).SetUint64(taskRun.MinimumConfirmations)
	min = min.Sub(min, big.NewInt(1))
	return diff.Cmp(min) >= 0
}

func saveAndTrigger(run *models.JobRun, store *store.Store) error {
	if err := store.Save(run); err != nil {
		return err
	}

	if run.Status == models.RunStatusInProgress {
		logger.Debugw(fmt.Sprintf("Executing new run initiated by %s", run.Initiator.Type), run.ForLogger()...)
		return store.RunChannel.Send(run.ID)
	}

	logger.Debugw(fmt.Sprintf("Queueing run initiated by %s", run.Initiator.Type), run.ForLogger()...)
	return nil
}

// RecurringScheduleJobError contains the field for the error message.
type RecurringScheduleJobError struct {
	msg string
}

// Error returns the error message for the run.
func (err RecurringScheduleJobError) Error() string {
	return err.msg
}
