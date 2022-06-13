package def

import (
	"github.com/conductor-sdk/conductor-go/model"
)

const (
	sqsEventPrefix       = "sqs"
	conductorEventPrefix = "conductor"
)

// EventTask Task to publish Events to external queuing systems like SQS, NATS, AMQP etc.
type EventTask struct {
	Task
	sink string
}

func NewSqsEventTask(taskRefName string, queueName string) *EventTask {
	return newEventTask(
		taskRefName,
		sqsEventPrefix,
		queueName,
	)
}

func NewConductorEventTask(taskRefName string, eventName string) *EventTask {
	return newEventTask(
		taskRefName,
		conductorEventPrefix,
		eventName,
	)
}

func newEventTask(taskRefName string, eventPrefix string, eventSuffix string) *EventTask {
	return &EventTask{
		Task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			taskType:          EVENT,
		},
		sink: eventPrefix + ":" + eventSuffix,
	}
}

func (task *EventTask) toWorkflowTask() []model.WorkflowTask {
	workflowTasks := task.Task.toWorkflowTask()
	workflowTasks[0].Sink = task.sink
	return workflowTasks
}

// Input to the task
func (task *EventTask) Input(key string, value interface{}) *EventTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *EventTask) InputMap(inputMap map[string]interface{}) *EventTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *EventTask) Optional(optional bool) *EventTask {
	task.Task.Optional(optional)
	return task
}

// Description of the task
func (task *EventTask) Description(description string) *EventTask {
	task.Task.Description(description)
	return task
}
