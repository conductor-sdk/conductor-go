package model

type WorkflowStatus string

const (
	RUNNING    WorkflowStatus = "RUNNING"
	COMPLETED  WorkflowStatus = "COMPLETED"
	FAILED     WorkflowStatus = "FAILED"
	TIMED_OUT  WorkflowStatus = "TIMED_OUT"
	TERMINATED WorkflowStatus = "TERMINATED"
	PAUSED     WorkflowStatus = "PAUSED"
)

var (
	WorkflowTerminalStates = []WorkflowStatus{
		COMPLETED,
		FAILED,
		TIMED_OUT,
		TERMINATED,
	}
)
