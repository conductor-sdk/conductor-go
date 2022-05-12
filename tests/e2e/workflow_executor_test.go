package e2e

import (
	"fmt"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/examples/task_execute_function"
	"github.com/conductor-sdk/conductor-go/pkg/worker"
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
	taskRunner := worker.NewWorkerOrkestratorWithApiClient(
		apiClient,
	)
	taskRunner.StartWorker(
		TASK_NAME,
		task_execute_function.Example1,
		WORKER_THREAD_COUNT,
		WORKER_POLLING_INTERVAL,
	)
	select {
	case workflow := <-workflowExecutionChannel:
		fmt.Println(workflow.WorkflowId)
	case <-time.After(5 * time.Second):
		t.Error()
	}
}
