package model

type SaveScheduleRequest struct {
	CreatedBy                   string                `json:"createdBy,omitempty"`
	CronExpression              string                `json:"cronExpression"`
	Name                        string                `json:"name"`
	Paused                      bool                  `json:"paused,omitempty"`
	RunCatchupScheduleInstances bool                  `json:"runCatchupScheduleInstances,omitempty"`
	ScheduleEndTime             int64                 `json:"scheduleEndTime,omitempty"`
	ScheduleStartTime           int64                 `json:"scheduleStartTime,omitempty"`
	StartWorkflowRequest        *StartWorkflowRequest `json:"startWorkflowRequest"`
	UpdatedBy                   string                `json:"updatedBy,omitempty"`
}
