package model

type Action struct {
	Action           string         `json:"action,omitempty"`
	StartWorkflow    *StartWorkflow `json:"start_workflow,omitempty"`
	CompleteTask     *TaskDetails   `json:"complete_task,omitempty"`
	FailTask         *TaskDetails   `json:"fail_task,omitempty"`
	ExpandInlineJSON bool           `json:"expandInlineJSON,omitempty"`
}
