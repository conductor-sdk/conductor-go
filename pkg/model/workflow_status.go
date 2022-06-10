package model

type WorkflowState struct {
	WorkflowId    string                 `json:"workflowId,omitempty"`
	Output        map[string]interface{} `json:"output,omitempty"`
	CorrelationId string                 `json:"correlationId,omitempty"`
	Variables     map[string]interface{} `json:"variables,omitempty"`
	Status        string                 `json:"status,omitempty"`
}
