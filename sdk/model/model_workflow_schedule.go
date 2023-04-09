package model

type WorkflowSchedule struct {
	CreateTime                  int64                 `json:"createTime,omitempty"`
	CreatedBy                   string                `json:"createdBy,omitempty"`
	CronExpression              string                `json:"cronExpression,omitempty"`
	Name                        string                `json:"name,omitempty"`
	Paused                      bool                  `json:"paused,omitempty"`
	PausedReason                string                `json:"pausedReason,omitempty"`
	RunCatchupScheduleInstances bool                  `json:"runCatchupScheduleInstances,omitempty"`
	ScheduleEndTime             int64                 `json:"scheduleEndTime,omitempty"`
	ScheduleStartTime           int64                 `json:"scheduleStartTime,omitempty"`
	StartWorkflowRequest        *StartWorkflowRequest `json:"startWorkflowRequest,omitempty"`
	Tags                        []Tag                 `json:"tags,omitempty"`
	UpdatedBy                   string                `json:"updatedBy,omitempty"`
	UpdatedTime                 int64                 `json:"updatedTime,omitempty"`
}
