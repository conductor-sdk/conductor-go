package model

type TaskDetails struct {
	WorkflowId  string                 `json:"workflowId,omitempty"`
	TaskRefName string                 `json:"taskRefName,omitempty"`
	Output      map[string]interface{} `json:"output,omitempty"`
	TaskId      string                 `json:"taskId,omitempty"`
}
