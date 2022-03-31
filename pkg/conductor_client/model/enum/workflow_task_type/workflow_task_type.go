package workflow_task_type

type WorkflowTaskType uint8

const (
	SIMPLE WorkflowTaskType = iota
	DYNAMIC
	FORK_JOIN
	FORK_JOIN_DYNAMIC
	DECISION
	JOIN
	SUB_WORKFLOW
	EVENT
	WAIT
	USER_DEFINED
)
