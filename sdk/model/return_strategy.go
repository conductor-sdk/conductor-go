package model

type ReturnStrategy string

const (
	ReturnTargetWorkflow    ReturnStrategy = "TARGET_WORKFLOW" // Default
	ReturnBlockingWorkflow  ReturnStrategy = "BLOCKING_WORKFLOW"
	ReturnBlockingTask      ReturnStrategy = "BLOCKING_TASK"
	ReturnBlockingTaskInput ReturnStrategy = "BLOCKING_TASK_INPUT"
)
