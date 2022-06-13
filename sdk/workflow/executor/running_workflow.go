package executor

type RunningWorkflow struct {
	WorkflowId               string
	WorkflowExecutionChannel WorkflowExecutionChannel
	Err                      error
}

func NewRunningWorkflow(workflowId string, workflowExecutionChannel WorkflowExecutionChannel, err error) *RunningWorkflow {
	return &RunningWorkflow{
		WorkflowId:               workflowId,
		WorkflowExecutionChannel: workflowExecutionChannel,
		Err:                      err,
	}
}
