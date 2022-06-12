package model

type StartWorkflow struct {
	Name          string                 `json:"name,omitempty"`
	Version       int32                  `json:"version,omitempty"`
	CorrelationId string                 `json:"correlationId,omitempty"`
	Input         map[string]interface{} `json:"input,omitempty"`
	TaskToDomain  map[string]string      `json:"taskToDomain,omitempty"`
}
