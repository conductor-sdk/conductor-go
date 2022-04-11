package task_result_status

type TaskResultStatus string

const (
	IN_PROGRESS TaskResultStatus = "IN_PROGRESS"
	CANCELED    TaskResultStatus = "CANCELED"
	FAILED      TaskResultStatus = "FAILED"
	COMPLETED   TaskResultStatus = "COMPLETED"
	SCHEDULED   TaskResultStatus = "SCHEDULED"
	TIMED_OUT   TaskResultStatus = "TIMED_OUT"
	SKIPPED     TaskResultStatus = "SKIPPED"
)
