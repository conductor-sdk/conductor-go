package e2e

import (
	"sync"
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
	workflowIdList := startWorkflows(
		t,
		WORKFLOW_EXECUTION_AMOUNT,
		WORKFLOW_NAME,
	)
	var waitGroup sync.WaitGroup
	for _, workflowId := range workflowIdList {
		waitGroup.Add(1)
		go validateWorkflow(t, &waitGroup, workflowId)
	}
	waitGroup.Wait()
}

func validateWorkflow(t *testing.T, waitGroup *sync.WaitGroup, workflowId string) {
	defer waitGroup.Done()
	time.Sleep(3 * time.Second)
	workflow := getWorkflowExecutionStatus(
		t,
		workflowId,
	)
	if workflow.Status != "COMPLETED" {
		t.Error("Incomplete workflow: ", workflowId)
	}
}
