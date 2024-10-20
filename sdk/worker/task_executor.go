package worker

import (
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"time"
)

type TaskWorker[T any, R any] struct {
	taskName   string
	function   func(ctx TaskContext, task *model.PolledTask) (interface{}, error)
	Options    *TaskWorkerOptions
	taskRunner *TaskRunner
}
type TaskWorkerOptions struct {
	Domain       string
	BatchSize    int
	PollInterval time.Duration
}
type executePolledTaskFunction func(ctx TaskContext, t *model.PolledTask) (interface{}, error)

func defaultOptions() *TaskWorkerOptions {
	return &TaskWorkerOptions{
		BatchSize:    1,
		PollInterval: time.Millisecond * 100,
	}
}

func NewWorker[T any, R any](taskName string, f func(T) (R, error)) *TaskWorker[T, R] {
	executorFn := func(ctx TaskContext, task *model.PolledTask) (interface{}, error) {

		var value T
		_, err := task.ToValue(&value)
		if err != nil {
			return nil, err
		}
		result, err := f(value)
		if err != nil {
			return nil, err
		}
		// Return the result as an interface{}
		return result, nil
	}

	return &TaskWorker[T, R]{
		taskName: taskName,
		function: executorFn,
		Options:  defaultOptions(),
	}
}

func NewWorkerWithCtx[T any, R any](taskName string, f func(ctx TaskContext, t T) (R, error)) *TaskWorker[T, R] {
	executorFn := func(ctx TaskContext, task *model.PolledTask) (interface{}, error) {

		var value T
		_, err := task.ToValue(&value)
		if err != nil {
			return nil, err
		}
		result, err := f(ctx, value)
		if err != nil {
			return nil, err
		}
		// Return the result as an interface{}
		return result, nil
	}

	return &TaskWorker[T, R]{
		taskName: taskName,
		function: executorFn,
		Options:  defaultOptions(),
	}
}

func (executor *TaskWorker[T, R]) UpdateBatchSize(batchSize int) error {
	executor.Options.BatchSize = batchSize
	if executor.taskRunner != nil {
		return executor.taskRunner.SetBatchSize(executor.taskName, batchSize)
	}
	return nil
}
func (executor *TaskWorker[T, R]) UpdatePollInterval(pollInterval time.Duration) error {
	executor.Options.PollInterval = pollInterval
	if executor.taskRunner != nil {
		return executor.taskRunner.SetPollIntervalForTask(executor.taskName, pollInterval)
	}
	return nil
}

func (executor *TaskWorker[T, R]) Start(taskRunner *TaskRunner) error {
	executor.taskRunner = taskRunner
	return taskRunner.StartWorkerWithExecFn(executor.taskName, executor.function, executor.Options)
}
