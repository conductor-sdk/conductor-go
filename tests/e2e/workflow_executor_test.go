package e2e

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
)

var workflowExecutor = executor.NewWorkflowExecutor(API_CLIENT)

type Treasure struct {
	ImportantValue string `json:"importantValue"`
}

type WorkflowValidator func(*http_model.Workflow) bool

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
	var waitGroup sync.WaitGroup
	for _, workflowExecutionChannel := range workflowExecutionChannelList {
		waitGroup.Add(1)
		go getWorkflowAndValidate(
			t,
			&waitGroup,
			workflowExecutionChannel,
			isWorkflowCompleted,
		)
	}
	waitGroup.Wait()
}

func isWorkflowCompleted(workflow *http_model.Workflow) bool {
	return workflow.Status == "COMPLETED"
}

func getWorkflowAndValidate(t *testing.T, waitGroup *sync.WaitGroup, workflowExecutionChannel executor.WorkflowExecutionChannel, isWorkflowValid WorkflowValidator) {
	defer waitGroup.Done()
	workflow, err := getWorkflow(workflowExecutionChannel)
	if err != nil {
		t.Error(err)
	}
	if !isWorkflowValid(workflow) {
		t.Error("Workflow is not valid: ", workflow)
	}
}

func getWorkflow(workflowExecutionChannel executor.WorkflowExecutionChannel) (*http_model.Workflow, error) {
	select {
	case workflow := <-workflowExecutionChannel:
		return &workflow, nil
	case <-time.After(5 * time.Second):
		return nil, fmt.Errorf("Timeout waiting for workflow")
	}
}
