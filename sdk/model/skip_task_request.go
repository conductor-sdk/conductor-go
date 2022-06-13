package model

type SkipTaskRequest struct {
	TaskInput  map[string]interface{} `json:"taskInput,omitempty"`
	TaskOutput map[string]interface{} `json:"taskOutput,omitempty"`
}
