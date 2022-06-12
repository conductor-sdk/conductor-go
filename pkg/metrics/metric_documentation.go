package metrics

type MetricDocumentation string

const (
	EXTERNAL_PAYLOAD_USED     MetricDocumentation = "Incremented each time external payload storage is used"
	TASK_ACK_ERROR            MetricDocumentation = "Task ack has encountered an exception"
	TASK_ACK_FAILED           MetricDocumentation = "Task ack failed"
	TASK_EXECUTE_ERROR        MetricDocumentation = "Execution error"
	TASK_EXECUTE_TIME         MetricDocumentation = "Time to execute a task"
	TASK_EXECUTION_QUEUE_FULL MetricDocumentation = "Counter to record execution queue has saturated"
	TASK_PAUSED               MetricDocumentation = "Counter for number of times the task has been polled, when the worker has been paused"
	TASK_POLL                 MetricDocumentation = "Incremented each time polling is done"
	TASK_POLL_ERROR           MetricDocumentation = "Client error when polling for a task queue"
	TASK_POLL_TIME            MetricDocumentation = "Time to poll for a batch of tasks"
	TASK_RESULT_SIZE          MetricDocumentation = "Records output payload size of a task"
	TASK_UPDATE_ERROR         MetricDocumentation = "Task status cannot be updated back to server"
	TASK_UPDATE_TIME          MetricDocumentation = "Time to update for a task"
	THREAD_UNCAUGHT_EXCEPTION MetricDocumentation = "thread_uncaught_exceptions"
	WORKFLOW_START_ERROR      MetricDocumentation = "Counter for workflow start errors"
	WORKFLOW_INPUT_SIZE       MetricDocumentation = "Records input payload size of a workflow"
)
