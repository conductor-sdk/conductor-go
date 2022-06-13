package model

type WorkflowStatus string

const (
	RunningWorkflow    WorkflowStatus = "RUNNING"
	CompletedWorkflow  WorkflowStatus = "COMPLETED"
	FailedWorkflow     WorkflowStatus = "FAILED"
	TimedOutWorkflow   WorkflowStatus = "TIMED_OUT"
	TerminatedWorkflow WorkflowStatus = "TERMINATED"
	PausedWorkflow     WorkflowStatus = "PAUSED"
)

var (
	WorkflowTerminalStates = []WorkflowStatus{
		CompletedWorkflow,
		FailedWorkflow,
		TimedOutWorkflow,
		TerminatedWorkflow,
	}
)
