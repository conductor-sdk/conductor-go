package http_model

type TaskSummary struct {
	WorkflowId                       string `json:"workflowId,omitempty"`
	WorkflowType                     string `json:"workflowType,omitempty"`
	CorrelationId                    string `json:"correlationId,omitempty"`
	ScheduledTime                    string `json:"scheduledTime,omitempty"`
	StartTime                        string `json:"startTime,omitempty"`
	UpdateTime                       string `json:"updateTime,omitempty"`
	EndTime                          string `json:"endTime,omitempty"`
	Status                           string `json:"status,omitempty"`
	ReasonForIncompletion            string `json:"reasonForIncompletion,omitempty"`
	ExecutionTime                    int64  `json:"executionTime,omitempty"`
	QueueWaitTime                    int64  `json:"queueWaitTime,omitempty"`
	TaskDefName                      string `json:"taskDefName,omitempty"`
	TaskType                         string `json:"taskType,omitempty"`
	Input                            string `json:"input,omitempty"`
	Output                           string `json:"output,omitempty"`
	TaskId                           string `json:"taskId,omitempty"`
	ExternalInputPayloadStoragePath  string `json:"externalInputPayloadStoragePath,omitempty"`
	ExternalOutputPayloadStoragePath string `json:"externalOutputPayloadStoragePath,omitempty"`
	WorkflowPriority                 int32  `json:"workflowPriority,omitempty"`
}
