package workflow_e2e

import (
	"testing"

	"github.com/conductor-sdk/conductor-go/examples"
	"github.com/conductor-sdk/conductor-go/pkg/worker"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/def/workflow"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
	"github.com/conductor-sdk/conductor-go/tests/e2e/e2e_properties"
	"github.com/conductor-sdk/conductor-go/tests/e2e/http_client_e2e/http_client_e2e_properties"
	"github.com/conductor-sdk/conductor-go/tests/e2e/workflow_e2e/workflow_e2e_properties"
	"github.com/conductor-sdk/conductor-go/tests/e2e/workflow_e2e/workflow_task_e2e"
)

var taskRunner = worker.NewTaskRunnerWithApiClient(e2e_properties.API_CLIENT)

var (
	workflowExecutorList = []*workflow.ConductorWorkflow{
		workflow_task_e2e.HTTP_WORKFLOW,
		// workflow_task_e2e.SIMPLE_WORKFLOW,
	}
)

func init() {
	taskRunner.StartWorker(
		workflow_task_e2e.SIMPLE_WORKFLOW.GetName(),
		examples.SimpleWorker,
		http_client_e2e_properties.WORKER_THREAD_COUNT,
		http_client_e2e_properties.WORKER_POLLING_INTERVAL,
	)
}

func TestValidateWorkflowDefinitions(t *testing.T) {
	for _, conductorWorkflow := range workflowExecutorList {
		response, err := conductorWorkflow.Register()
		if err != nil {
			t.Error("Response: ", response, ", error: ", err)
		}
	}
}

func TestWorkflowDefExecution(t *testing.T) {
	workflowExecutionChannelList := make([]*executor.WorkflowExecutionChannel, len(workflowExecutorList))
	for i, conductorWorkflow := range workflowExecutorList {
		workflowExecutionChannel, err := conductorWorkflow.Start(nil)
		if err != nil {
			t.Error(err)
		}
		workflowExecutionChannelList[i] = &workflowExecutionChannel
	}
	workflow_e2e_properties.WaitForCompletionOfWorkflows(
		t,
		workflowExecutionChannelList,
		workflow_e2e_properties.IsWorkflowCompleted,
	)
}
