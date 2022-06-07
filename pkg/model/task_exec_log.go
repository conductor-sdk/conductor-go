package model

type TaskExecLog struct {
	Log         string `json:"log,omitempty"`
	TaskId      string `json:"taskId,omitempty"`
	CreatedTime int64  `json:"createdTime,omitempty"`
}
