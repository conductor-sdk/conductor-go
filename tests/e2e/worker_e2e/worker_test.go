package worker_e2e

import (
	"sync"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/examples"
	"github.com/conductor-sdk/conductor-go/pkg/worker"
	"github.com/conductor-sdk/conductor-go/tests/e2e/e2e_properties"
	"github.com/conductor-sdk/conductor-go/tests/e2e/http_client_e2e"
	"github.com/conductor-sdk/conductor-go/tests/e2e/http_client_e2e/http_client_e2e_properties"
)

var taskRunner = worker.NewTaskRunnerWithApiClient(e2e_properties.API_CLIENT)

func TestTaskRunnerExecution(t *testing.T) {
	workflowIdList, err := http_client_e2e.StartWorkflows(
		http_client_e2e_properties.WORKFLOW_EXECUTION_AMOUNT,
		http_client_e2e_properties.WORKFLOW_NAME,
	)
	if err != nil {
		t.Error(err)
	}

	taskRunner.StartWorker(
		http_client_e2e_properties.TASK_NAME,
		examples.SimpleWorker,
		http_client_e2e_properties.WORKER_THREAD_COUNT,
		http_client_e2e_properties.WORKER_POLLING_INTERVAL,
	)

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(workflowIdList))
	for _, workflowId := range workflowIdList {
		go validateWorkflow(t, &waitGroup, workflowId)
	}
	waitGroup.Wait()

	taskRunner.RemoveWorker(
		http_client_e2e_properties.TASK_NAME,
		http_client_e2e_properties.WORKER_THREAD_COUNT,
	)
}

func validateWorkflow(t *testing.T, waitGroup *sync.WaitGroup, workflowId string) {
	defer waitGroup.Done()
	time.Sleep(3 * time.Second)
	workflow, _, err := http_client_e2e.GetWorkflowExecutionStatus(
		workflowId,
	)
	if err != nil {
		t.Error(err)
	}
	if workflow.Status != "COMPLETED" {
		t.Error("Incomplete workflow: ", workflowId)
	}
}
