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

type WorkflowValidator func(*http_model.Workflow) bool

func TestWorkflowExecutor(t *testing.T) {
	workflowExecutionChannelList := getWorkflowExecutionChannelList(
		t,
		WORKFLOW_NAME,
		1,
		nil,
	)
	waitForCompletionOfWorkflows(
		t,
		workflowExecutionChannelList,
	)
}

func TestWorkflowExecutorWithCustomInput(t *testing.T) {
	workflowExecutionChannelList := getWorkflowExecutionChannelList(
		t,
		TREASURE_CHEST_WORKFLOW_NAME,
		1,
		TREASURE_WORKFLOW_INPUT,
	)
	waitForCompletionOfWorkflows(
		t,
		workflowExecutionChannelList,
	)
}

func getWorkflowExecutionChannelList(t *testing.T, workflowName string, version int32, input interface{}) []executor.WorkflowExecutionChannel {
	workflowExecutionChannelList := make([]executor.WorkflowExecutionChannel, WORKFLOW_EXECUTION_AMOUNT)
	for i := 0; i < WORKFLOW_EXECUTION_AMOUNT; i += 1 {
		workflowExecutionChannel, err := workflowExecutor.ExecuteWorkflow(
			workflowName,
			version,
			input,
		)
		if err != nil {
			t.Error(err)
		}
		workflowExecutionChannelList[i] = workflowExecutionChannel
	}
	return workflowExecutionChannelList
}

func waitForCompletionOfWorkflows(t *testing.T, workflowExecutionChannelList []executor.WorkflowExecutionChannel) {
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
