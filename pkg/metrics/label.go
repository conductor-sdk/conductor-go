package metrics

type MetricLabel string

const (
	ENTITY_NAME      MetricLabel = "entityName"
	EXCEPTION                    = "exception"
	OPERATION                    = "operation"
	PAYLOAD_TYPE                 = "payload_type"
	TASK_TYPE                    = "taskType"
	WORKFLOW_TYPE                = "workflowType"
	WORKFLOW_VERSION             = "version"
)
