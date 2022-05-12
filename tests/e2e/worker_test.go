package e2e

import (
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/examples/task_execute_function"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/worker"
)

var taskRunner = worker.NewTaskRunnerWithApiClient(API_CLIENT)

func init() {
	taskRunner.StartWorker(
		TASK_NAME,
		task_execute_function.Example1,
		WORKER_THREAD_COUNT,
		WORKER_POLLING_INTERVAL,
	)
}

func TestTaskRunnerExecution(t *testing.T) {
	registerTaskDefinition(
		t,
		[]http_model.TaskDef{
			TASK_DEFINITION,
		},
	)
	registerWorkflowDefinition(
		t,
		WORKFLOW_DEFINITION,
	)
	workflowQty := 10
	workflowIdList := startWorkflows(
		t,
		workflowQty,
		WORKFLOW_NAME,
	)
	time.Sleep(5 * time.Second)
	for i := range workflowIdList {
		workflow := getWorkflowExecutionStatus(
			t,
			workflowIdList[i],
		)
		if workflow.Status != "COMPLETED" {
			t.Error("Incomplete workflow: ", workflowIdList[i])
		}
	}
}
