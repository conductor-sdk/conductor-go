package workflow_e2e_properties

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
	"github.com/conductor-sdk/conductor-go/tests/e2e/e2e_properties"
)

type WorkflowValidator func(*http_model.Workflow) bool

var (
	WorkflowExecutor = executor.NewWorkflowExecutor(e2e_properties.API_CLIENT)
)

func IsWorkflowCompleted(workflow *http_model.Workflow) bool {
	return workflow.Status == "COMPLETED"
}

func WaitForCompletionOfWorkflows(t *testing.T, workflowExecutionChannelList []*executor.WorkflowExecutionChannel, isWorkflowValid WorkflowValidator) {
	var waitGroup sync.WaitGroup
	for _, workflowExecutionChannel := range workflowExecutionChannelList {
		waitGroup.Add(1)
		go getWorkflowAndValidate(
			t,
			&waitGroup,
			workflowExecutionChannel,
			IsWorkflowCompleted,
		)
	}
	waitGroup.Wait()
}

func getWorkflowAndValidate(t *testing.T, waitGroup *sync.WaitGroup, workflowExecutionChannel *executor.WorkflowExecutionChannel, isWorkflowValid WorkflowValidator) {
	defer waitGroup.Done()
	workflow, err := getWorkflow(workflowExecutionChannel)
	if err != nil {
		t.Error(err)
	}
	if !isWorkflowValid(workflow) {
		t.Error("Workflow is not valid: ", workflow)
	}
}

func getWorkflow(workflowExecutionChannel *executor.WorkflowExecutionChannel) (*http_model.Workflow, error) {
	select {
	case workflow := <-*workflowExecutionChannel:
		return &workflow, nil
	case <-time.After(5 * time.Second):
		return nil, fmt.Errorf("timeout waiting for workflow")
	}
}
