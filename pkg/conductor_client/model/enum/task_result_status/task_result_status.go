package task_result_status

type TaskResultStatus string

const (
	IN_PROGRESS TaskResultStatus = "IN_PROGRESS"
	CANCELED                     = "CANCELED"
	FAILED                       = "FAILED"
	COMPLETED                    = "COMPLETED"
	SCHEDULED                    = "SCHEDULED"
	TIMED_OUT                    = "TIMED_OUT"
	SKIPPED                      = "SKIPPED"
)
