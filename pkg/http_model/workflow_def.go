package http_model

type WorkflowDef struct {
	OwnerApp                      string                 `json:"ownerApp,omitempty"`
	CreateTime                    int64                  `json:"createTime,omitempty"`
	UpdateTime                    int64                  `json:"updateTime,omitempty"`
	CreatedBy                     string                 `json:"createdBy,omitempty"`
	UpdatedBy                     string                 `json:"updatedBy,omitempty"`
	Name                          string                 `json:"name"`
	Description                   string                 `json:"description,omitempty"`
	Version                       int32                  `json:"version,omitempty"`
	Tasks                         []WorkflowTask         `json:"tasks"`
	InputParameters               []string               `json:"inputParameters,omitempty"`
	OutputParameters              map[string]interface{} `json:"outputParameters,omitempty"`
	FailureWorkflow               string                 `json:"failureWorkflow,omitempty"`
	SchemaVersion                 int32                  `json:"schemaVersion,omitempty"`
	Restartable                   bool                   `json:"restartable,omitempty"`
	WorkflowStatusListenerEnabled bool                   `json:"workflowStatusListenerEnabled,omitempty"`
	OwnerEmail                    string                 `json:"ownerEmail,omitempty"`
	TimeoutPolicy                 string                 `json:"timeoutPolicy,omitempty"`
	TimeoutSeconds                int64                  `json:"timeoutSeconds"`
	Variables                     map[string]interface{} `json:"variables,omitempty"`
	InputTemplate                 map[string]interface{} `json:"inputTemplate,omitempty"`
}
