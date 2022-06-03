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

func NewKafkaPublishTask(taskName string, kafkaPublishTaskInput *KafkaPublishTaskInput) *KafkaPublishTask {
	return &KafkaPublishTask{
		Task: Task{
			name:              taskName,
			taskReferenceName: taskName,
			taskType:          KAFKA_PUBLISH,
			inputParameters: map[string]interface{}{
				"kafka_request": kafkaPublishTaskInput,
			},
		},
	}
}
