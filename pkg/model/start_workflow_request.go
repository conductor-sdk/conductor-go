package model

type StartWorkflowRequest struct {
	Name                            string            `json:"name"`
	Version                         int32             `json:"version,omitempty"`
	CorrelationId                   string            `json:"correlationId,omitempty"`
	Input                           interface{}       `json:"input,omitempty"`
	TaskToDomain                    map[string]string `json:"taskToDomain,omitempty"`
	WorkflowDef                     *WorkflowDef      `json:"workflowDef,omitempty"`
	ExternalInputPayloadStoragePath string            `json:"externalInputPayloadStoragePath,omitempty"`
	Priority                        int32             `json:"priority,omitempty"`
}

func NewStartWorkflowRequest(name string, version int32, correlationId string, input interface{}) *StartWorkflowRequest {
	return &StartWorkflowRequest{
		Name:          name,
		Version:       version,
		CorrelationId: correlationId,
		Input:         input,
	}
}
