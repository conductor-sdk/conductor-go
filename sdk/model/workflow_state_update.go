package model

type WorkflowStateUpdate struct {
	TaskReferenceName string                 `json:"taskReferenceName,omitempty"`
	TaskResult        *TaskResult            `json:"taskResult,omitempty"`
	Variables         map[string]interface{} `json:"variables,omitempty"`
}
