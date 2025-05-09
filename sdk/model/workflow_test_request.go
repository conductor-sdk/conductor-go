package model

type WorkflowTestRequest struct {
	CorrelationId                   string                         `json:"correlationId,omitempty"`
	CreatedBy                       string                         `json:"createdBy,omitempty"`
	ExternalInputPayloadStoragePath string                         `json:"externalInputPayloadStoragePath,omitempty"`
	IdempotencyKey                  string                         `json:"idempotencyKey,omitempty"`
	IdempotencyStrategy             string                         `json:"idempotencyStrategy,omitempty"`
	Input                           map[string]interface{}         `json:"input,omitempty"`
	Name                            string                         `json:"name"`
	Priority                        int32                          `json:"priority,omitempty"`
	SubWorkflowTestRequest          map[string]WorkflowTestRequest `json:"subWorkflowTestRequest,omitempty"`
	TaskRefToMockOutput             map[string][]TaskMock          `json:"taskRefToMockOutput,omitempty"`
	TaskToDomain                    map[string]string              `json:"taskToDomain,omitempty"`
	Version                         int32                          `json:"version,omitempty"`
	WorkflowDef                     *WorkflowDef                   `json:"workflowDef,omitempty"`
}

type TaskMock struct {
	ExecutionTime int64                  `json:"executionTime,omitempty"`
	Output        map[string]interface{} `json:"output,omitempty"`
	QueueWaitTime int64                  `json:"queueWaitTime,omitempty"`
	Status        string                 `json:"status,omitempty"`
}
