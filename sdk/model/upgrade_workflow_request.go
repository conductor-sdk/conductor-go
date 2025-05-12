package model

type UpgradeWorkflowRequest struct {
	Name          string                 `json:"name"`
	TaskOutput    map[string]interface{} `json:"taskOutput,omitempty"`
	Version       int32                  `json:"version,omitempty"`
	WorkflowInput map[string]interface{} `json:"workflowInput,omitempty"`
}
