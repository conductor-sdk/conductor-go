package model

type RerunWorkflowRequest struct {
	ReRunFromWorkflowId string                 `json:"reRunFromWorkflowId,omitempty"`
	WorkflowInput       map[string]interface{} `json:"workflowInput,omitempty"`
	ReRunFromTaskId     string                 `json:"reRunFromTaskId,omitempty"`
	TaskInput           map[string]interface{} `json:"taskInput,omitempty"`
	CorrelationId       string                 `json:"correlationId,omitempty"`
}
