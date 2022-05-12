package e2e

import (
	"fmt"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
)

var workflowExecutor = executor.NewWorkflowExecutor(API_CLIENT)

func TestWorkflowExecutor(t *testing.T) {
	workflowExecutionChannelList := make([]executor.WorkflowExecutionChannel, WORKFLOW_EXECUTION_AMOUNT)
	for i := 0; i < WORKFLOW_EXECUTION_AMOUNT; i += 1 {
		workflowExecutionChannel, err := workflowExecutor.ExecuteWorkflow(
			WORKFLOW_NAME,
			1,
			nil,
		)
		if err != nil {
			t.Error(err)
		}
		workflowExecutionChannelList[i] = workflowExecutionChannel
	}
	for _, workflowExecutionChannel := range workflowExecutionChannelList {
		select {
		case workflow := <-workflowExecutionChannel:
			fmt.Println(workflow)
		case <-time.After(5 * time.Second):
			t.Error()
		}
	}
}
