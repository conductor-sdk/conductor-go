package e2e

import (
	"fmt"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
)

var workflowExecutor = executor.NewWorkflowExecutor(
	getApiClientWithAuthentication(),
)

func TestWorkflowExecutor(t *testing.T) {
	workflowExecutionChannel, err := workflowExecutor.ExecuteWorkflow(
		WORKFLOW_NAME,
		1,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	select {
	case workflow := <-workflowExecutionChannel:
		fmt.Println(workflow)
	case <-time.After(5 * time.Second):
		t.Error()
	}
}
