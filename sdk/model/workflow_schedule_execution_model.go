package model

type WorkflowScheduleExecution struct {
	ExecutionId          string                `json:"executionId,omitempty"`
	ExecutionTime        int64                 `json:"executionTime,omitempty"`
	Reason               string                `json:"reason,omitempty"`
	ScheduleName         string                `json:"scheduleName,omitempty"`
	ScheduledTime        int64                 `json:"scheduledTime,omitempty"`
	StackTrace           string                `json:"stackTrace,omitempty"`
	StartWorkflowRequest *StartWorkflowRequest `json:"startWorkflowRequest,omitempty"`
	State                string                `json:"state,omitempty"`
	WorkflowId           string                `json:"workflowId,omitempty"`
	WorkflowName         string                `json:"workflowName,omitempty"`
}
