// workflowrun.go
package model

import (
	"time"
)

// WorkflowRun represents a workflow run in the Conductor system
type WorkflowRun struct {
	SignalResponse
	Priority   int                    `json:"priority,omitempty"`
	Variables  map[string]interface{} `json:"variables,omitempty"`
	Tasks      []Task                 `json:"tasks,omitempty"`
	CreatedBy  string                 `json:"createdBy,omitempty"`
	CreateTime int64                  `json:"createTime,omitempty"`
	Status     WorkflowStatus         `json:"status,omitempty"`
	UpdateTime int64                  `json:"updateTime,omitempty"`

	WorkflowId    string                 `json:"workflowId,omitempty"`
	CorrelationId string                 `json:"correlationId,omitempty"`
	Input         map[string]interface{} `json:"input,omitempty"`
	Output        map[string]interface{} `json:"output,omitempty"`
}

// GetCreateTimeFormatted formats CreateTime as a readable string
func (w *WorkflowRun) GetCreateTimeFormatted() string {
	return time.Unix(0, w.CreateTime*int64(time.Millisecond)).String()
}

// GetUpdateTimeFormatted formats UpdateTime as a readable string
func (w *WorkflowRun) GetUpdateTimeFormatted() string {
	return time.Unix(0, w.UpdateTime*int64(time.Millisecond)).String()
}
