package workflow_e2e

import (
	"testing"

	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
	"github.com/conductor-sdk/conductor-go/tests/e2e/http_client_e2e/http_client_e2e_properties"
	"github.com/conductor-sdk/conductor-go/tests/e2e/workflow_e2e/workflow_e2e_properties"
)

func TestWorkflowExecutor(t *testing.T) {
	workflowExecutionChannelList := getWorkflowExecutionChannelList(
		t,
		http_client_e2e_properties.WORKFLOW_NAME,
		1,
		nil,
	)
	workflow_e2e_properties.WaitForCompletionOfWorkflows(
		t,
		workflowExecutionChannelList,
		workflow_e2e_properties.IsWorkflowCompleted,
	)
}

func TestWorkflowExecutorWithCustomInput(t *testing.T) {
	workflowExecutionChannelList := getWorkflowExecutionChannelList(
		t,
		http_client_e2e_properties.TREASURE_CHEST_WORKFLOW_NAME,
		1,
		http_client_e2e_properties.TREASURE_WORKFLOW_INPUT,
	)
	workflow_e2e_properties.WaitForCompletionOfWorkflows(
		t,
		workflowExecutionChannelList,
		workflow_e2e_properties.IsWorkflowCompleted,
	)
}

func getWorkflowExecutionChannelList(t *testing.T, workflowName string, version int32, input interface{}) []*executor.WorkflowExecutionChannel {
	workflowExecutionChannelList := make([]*executor.WorkflowExecutionChannel, http_client_e2e_properties.WORKFLOW_EXECUTION_AMOUNT)
	for i := 0; i < http_client_e2e_properties.WORKFLOW_EXECUTION_AMOUNT; i += 1 {
		workflowExecutionChannel, err := workflow_e2e_properties.WorkflowExecutor.ExecuteWorkflow(
			workflowName,
			version,
			input,
		)
		if err != nil {
			t.Error(err)
		}
		workflowExecutionChannelList[i] = &workflowExecutionChannel
	}
	return workflowExecutionChannelList
}
