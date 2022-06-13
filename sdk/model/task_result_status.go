package model

type TaskResultStatus string

const (
	IN_PROGRESS                TaskResultStatus = "IN_PROGRESS"
	CANCELED                   TaskResultStatus = "CANCELED"
	FAILED                     TaskResultStatus = "FAILED"
	FAILED_WITH_TERMINAL_ERROR TaskResultStatus = "FAILED_WITH_TERMINAL_ERROR"
	COMPLETED                  TaskResultStatus = "COMPLETED"
	SCHEDULED                  TaskResultStatus = "SCHEDULED"
	TIMED_OUT                  TaskResultStatus = "TIMED_OUT"
	SKIPPED                    TaskResultStatus = "SKIPPED"
)
