package workflow

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
		Task{
			name:              taskName,
			taskReferenceName: taskName,
			taskType:          EVENT,
		},
		eventPrefix + ":" + eventSuffix,
	}
}
