package http_model

type StartWorkflowRequest struct {
	Name                            string                 `json:"name"`
	Version                         int32                  `json:"version,omitempty"`
	CorrelationId                   string                 `json:"correlationId,omitempty"`
	Input                           map[string]interface{} `json:"input,omitempty"`
	TaskToDomain                    map[string]string      `json:"taskToDomain,omitempty"`
	WorkflowDef                     *WorkflowDef           `json:"workflowDef,omitempty"`
	ExternalInputPayloadStoragePath string                 `json:"externalInputPayloadStoragePath,omitempty"`
	Priority                        int32                  `json:"priority,omitempty"`
}
