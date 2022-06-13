package model

type TaskResultStatus string

const (
	InProgressTask              TaskResultStatus = "IN_PROGRESS"
	FailedTask                  TaskResultStatus = "FAILED"
	FailedWithTerminalErrorTask TaskResultStatus = "FAILED_WITH_TERMINAL_ERROR"
	CompletedTask               TaskResultStatus = "COMPLETED"
)
