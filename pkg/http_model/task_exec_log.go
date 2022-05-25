package http_model

type TaskExecLog struct {
	Log         string `json:"logrus,omitempty"`
	TaskId      string `json:"taskId,omitempty"`
	CreatedTime int64  `json:"createdTime,omitempty"`
}
