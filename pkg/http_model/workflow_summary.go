package http_model

type WorkflowSummary struct {
	WorkflowType                     string `json:"workflowType,omitempty"`
	Version                          int32  `json:"version,omitempty"`
	WorkflowId                       string `json:"workflowId,omitempty"`
	CorrelationId                    string `json:"correlationId,omitempty"`
	StartTime                        string `json:"startTime,omitempty"`
	UpdateTime                       string `json:"updateTime,omitempty"`
	EndTime                          string `json:"endTime,omitempty"`
	Status                           string `json:"status,omitempty"`
	Input                            string `json:"input,omitempty"`
	Output                           string `json:"output,omitempty"`
	ReasonForIncompletion            string `json:"reasonForIncompletion,omitempty"`
	ExecutionTime                    int64  `json:"executionTime,omitempty"`
	Event                            string `json:"event,omitempty"`
	FailedReferenceTaskNames         string `json:"failedReferenceTaskNames,omitempty"`
	ExternalInputPayloadStoragePath  string `json:"externalInputPayloadStoragePath,omitempty"`
	ExternalOutputPayloadStoragePath string `json:"externalOutputPayloadStoragePath,omitempty"`
	Priority                         int32  `json:"priority,omitempty"`
	OutputSize                       int64  `json:"outputSize,omitempty"`
	InputSize                        int64  `json:"inputSize,omitempty"`
}
