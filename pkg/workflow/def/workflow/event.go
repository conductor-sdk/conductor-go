package workflow

import "github.com/conductor-sdk/conductor-go/pkg/http_model"

const (
	sqsEventPrefix       = "sqs"
	conductorEventPrefix = "conductor"
)

// Task to publish Events to external queuing systems like SQS, NATS, AMQP etc.
type EventTask struct {
	Task
	sink string
}

func NewSqsEventTask(taskName string, queueName string) *EventTask {
	return newEventTask(
		taskName,
		sqsEventPrefix,
		queueName,
	)
}

func NewConductorEventTask(taskName string, eventName string) *EventTask {
	return newEventTask(
		taskName,
		conductorEventPrefix,
		eventName,
	)
}

func newEventTask(taskName string, eventPrefix string, eventSuffix string) *EventTask {
	return &EventTask{
		Task: Task{
			name:              taskName,
			taskReferenceName: taskName,
			taskType:          EVENT,
		},
		sink: eventPrefix + ":" + eventSuffix,
	}
}

func (task *EventTask) toWorkflowTask() []http_model.WorkflowTask {
	workflowTasks := task.Task.toWorkflowTask()
	workflowTasks[0].Sink = task.sink
	return workflowTasks
}
