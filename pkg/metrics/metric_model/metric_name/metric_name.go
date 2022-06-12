package metric_name

type MetricName string

const (
	EXTERNAL_PAYLOAD_USED     MetricName = "external_payload_used"
	TASK_ACK_ERROR            MetricName = "task_ack_error"
	TASK_ACK_FAILED           MetricName = "task_ack_failed"
	TASK_EXECUTE_ERROR        MetricName = "task_execute_error"
	TASK_EXECUTE_TIME         MetricName = "task_execute_time"
	TASK_EXECUTION_QUEUE_FULL MetricName = "task_execution_queue_full"
	TASK_PAUSED               MetricName = "task_paused"
	TASK_POLL                 MetricName = "task_poll"
	TASK_POLL_ERROR           MetricName = "task_poll_error"
	TASK_POLL_TIME            MetricName = "task_poll_time"
	TASK_RESULT_SIZE          MetricName = "task_result_size"
	TASK_UPDATE_ERROR         MetricName = "task_update_error"
	TASK_UPDATE_TIME          MetricName = "task_update_time"
	THREAD_UNCAUGHT_EXCEPTION MetricName = "thread_uncaught_exceptions"
	WORKFLOW_INPUT_SIZE       MetricName = "workflow_input_size"
	WORKFLOW_START_ERROR      MetricName = "workflow_start_error"
)
