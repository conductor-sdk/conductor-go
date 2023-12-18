package model

type WorkflowRun struct {
	CorrelationId string                 `json:"correlationId,omitempty"`
	CreateTime    int64                  `json:"createTime,omitempty"`
	CreatedBy     string                 `json:"createdBy,omitempty"`
	Input         map[string]interface{} `json:"input,omitempty"`
	Output        map[string]interface{} `json:"output,omitempty"`
	Priority      int32                  `json:"priority,omitempty"`
	RequestId     string                 `json:"requestId,omitempty"`
	Status        string                 `json:"status,omitempty"`
	Tasks         []Task                 `json:"tasks,omitempty"`
	UpdateTime    int64                  `json:"updateTime,omitempty"`
	Variables     map[string]interface{} `json:"variables,omitempty"`
	WorkflowId    string                 `json:"workflowId,omitempty"`
}
