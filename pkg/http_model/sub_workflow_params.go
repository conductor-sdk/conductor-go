package http_model

type SubWorkflowParams struct {
	Name               string            `json:"name"`
	Version            int32             `json:"version,omitempty"`
	TaskToDomain       map[string]string `json:"taskToDomain,omitempty"`
	WorkflowDefinition *WorkflowDef      `json:"workflowDefinition,omitempty"`
}
