package model

type SignalResponse struct {
	CorrelationId        string                 `json:"correlationId,omitempty"`
	Input                map[string]interface{} `json:"input,omitempty"`
	Output               map[string]interface{} `json:"output,omitempty"`
	RequestId            string                 `json:"requestId,omitempty"`
	ResponseType         string                 `json:"responseType,omitempty"`
	TargetWorkflowId     string                 `json:"targetWorkflowId,omitempty"`
	TargetWorkflowStatus string                 `json:"targetWorkflowStatus,omitempty"`
	WorkflowId           string                 `json:"workflowId,omitempty"`
}
