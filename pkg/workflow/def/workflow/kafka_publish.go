package workflow

type KafkaPublishTask struct {
	Task
}

type KafkaPublishTaskInput struct {
	BootStrapServers string                 `json:"bootStrapServers"`
	Key              string                 `json:"key"`
	KeySerializer    string                 `json:"keySerializer,omitempty"`
	Value            string                 `json:"value"`
	RequestTimeoutMs string                 `json:"requestTimeoutMs,omitempty"`
	MaxBlockMs       string                 `json:"maxBlockMs,omitempty"`
	Headers          map[string]interface{} `json:"headers,omitempty"`
	Topic            string                 `json:"topic"`
}

func NewKafkaPublishTask(taskRefName string, kafkaPublishTaskInput *KafkaPublishTaskInput) *KafkaPublishTask {
	return &KafkaPublishTask{
		Task: Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			taskType:          KAFKA_PUBLISH,
			inputParameters: map[string]interface{}{
				"kafka_request": kafkaPublishTaskInput,
			},
		},
	}
}

// Input to the task
func (task *KafkaPublishTask) Input(key string, value interface{}) *KafkaPublishTask {
	task.Task.Input(key, value)
	return task
}
func (task *KafkaPublishTask) Optional(optional bool) *KafkaPublishTask {
	task.Task.Optional(optional)
	return task
}
func (task *KafkaPublishTask) Description(description string) *KafkaPublishTask {
	task.Task.Description(description)
	return task
}
