package e2e

import (
	"sync"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/worker"
	"github.com/conductor-sdk/conductor-go/tests"
)

var taskRunner = worker.NewTaskRunnerWithApiClient(API_CLIENT)

func init() {
	for taskName, taskExecuteFunction := range tests.TASK_DEFINITION_TO_WORKER {
		taskRunner.StartWorker(
			taskName,
			taskExecuteFunction,
			tests.WORKER_THREAD_COUNT,
			tests.WORKER_POLLING_INTERVAL,
		)
	}
}

func TestTaskRunnerExecution(t *testing.T) {
	registerTaskDefinition(
		t,
		[]http_model.TaskDef{
			tests.TASK_DEFINITION,
		},
	)
	registerWorkflowDefinition(
		t,
		tests.WORKFLOW_DEFINITION,
	)
	workflowIdList := startWorkflows(
		t,
		tests.WORKFLOW_EXECUTION_AMOUNT,
		tests.WORKFLOW_NAME,
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
