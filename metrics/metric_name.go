package metrics

type MetricName uint8

const (
	EXTERNAL_PAYLOAD_USED MetricName = iota
	TASK_ACK_ERROR
	TASK_ACK_FAILED
	TASK_EXECUTE_ERROR
	TASK_EXECUTE_TIME
	TASK_EXECUTION_QUEUE_FULL
	TASK_PAUSED
	TASK_POLL
	TASK_POLL_ERROR
	TASK_POLL_TIME
	TASK_RESULT_SIZE
	TASK_UPDATE_ERROR
	THREAD_UNCAUGHT_EXCEPTION
	WORKFLOW_INPUT_SIZE
	WORKFLOW_START_ERROR
)
