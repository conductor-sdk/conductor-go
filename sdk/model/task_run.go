package model

import (
	"time"
)

// TaskRun represents a task run in the Conductor system
type TaskRun struct {
	TaskType              string                 `json:"taskType,omitempty"`
	TaskId                string                 `json:"taskId,omitempty"`
	ReferenceTaskName     string                 `json:"referenceTaskName,omitempty"`
	RetryCount            int32                  `json:"retryCount,omitempty"`
	TaskDefName           string                 `json:"taskDefName,omitempty"`
	RetriedTaskId         string                 `json:"retriedTaskId,omitempty"`
	WorkflowType          string                 `json:"workflowType,omitempty"`
	ReasonForIncompletion string                 `json:"reasonForIncompletion,omitempty"`
	Priority              int                    `json:"priority,omitempty"`
	Variables             map[string]interface{} `json:"variables,omitempty"`
	Tasks                 []Task                 `json:"tasks,omitempty"`
	CreatedBy             string                 `json:"createdBy,omitempty"`
	CreateTime            int64                  `json:"createTime,omitempty"`
	UpdateTime            int64                  `json:"updateTime,omitempty"`
	Status                TaskResultStatus       `json:"status,omitempty"`
	InputData             map[string]interface{} `json:"inputData,omitempty"`
	OutputData            map[string]interface{} `json:"outputData,omitempty"`
}

// GetCreateTimeFormatted formats CreateTime as a readable string
func (t *TaskRun) GetCreateTimeFormatted() string {
	return time.Unix(0, t.CreateTime*int64(time.Millisecond)).String()
}

// GetUpdateTimeFormatted formats UpdateTime as a readable string
func (t *TaskRun) GetUpdateTimeFormatted() string {
	return time.Unix(0, t.UpdateTime*int64(time.Millisecond)).String()
}
