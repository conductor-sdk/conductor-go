package worker

import "context"

type TaskContext interface {
	context.Context
	GetRetryCount() int32
	GetCorrelationId() string
	GetPollCount() int32
	GetRetriedTaskId() string
	GetRetried() bool
	GetWorkflowInstanceId() string
	GetWorkflowType() string
	GetTaskId() string
	GetWorkflowPriority() int32
}

type TaskContextImpl struct {
	context.Context    // Embed the standard context.Context
	RetryCount         int32
	CorrelationId      string
	PollCount          int32
	RetriedTaskId      string
	Retried            bool
	WorkflowInstanceId string
	WorkflowType       string
	TaskId             string
	WorkflowPriority   int32
}

// Implement the TaskContext interface by providing the getter methods

func (t *TaskContextImpl) GetRetryCount() int32 {
	return t.RetryCount
}

func (t *TaskContextImpl) GetCorrelationId() string {
	return t.CorrelationId
}

func (t *TaskContextImpl) GetPollCount() int32 {
	return t.PollCount
}

func (t *TaskContextImpl) GetRetriedTaskId() string {
	return t.RetriedTaskId
}

func (t *TaskContextImpl) GetRetried() bool {
	return t.Retried
}

func (t *TaskContextImpl) GetWorkflowInstanceId() string {
	return t.WorkflowInstanceId
}

func (t *TaskContextImpl) GetWorkflowType() string {
	return t.WorkflowType
}

func (t *TaskContextImpl) GetTaskId() string {
	return t.TaskId
}

func (t *TaskContextImpl) GetWorkflowPriority() int32 {
	return t.WorkflowPriority
}
