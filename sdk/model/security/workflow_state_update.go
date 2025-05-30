package security

// WorkflowStateUpdate represents a workflow state update
type WorkflowStateUpdate struct {
	TaskReferenceName string                 `json:"taskReferenceName,omitempty"`
	Variables         map[string]interface{} `json:"variables,omitempty"`
	TaskResult        interface{}            `json:"taskResult,omitempty"`
}