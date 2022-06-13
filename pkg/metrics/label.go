package metrics

type MetricLabel string

const (
	ENTITY_NAME      MetricLabel = "entityName"
	EXCEPTION        MetricLabel = "exception"
	OPERATION        MetricLabel = "operation"
	PAYLOAD_TYPE     MetricLabel = "payload_type"
	TASK_TYPE        MetricLabel = "taskType"
	WORKFLOW_TYPE    MetricLabel = "workflowType"
	WORKFLOW_VERSION MetricLabel = "version"
)
