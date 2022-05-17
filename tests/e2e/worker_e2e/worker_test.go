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

func init() {
	taskRunner.StartWorker(
		http_client_e2e_properties.TASK_NAME,
		examples.SimpleWorker,
		http_client_e2e_properties.WORKER_THREAD_COUNT,
		http_client_e2e_properties.WORKER_POLLING_INTERVAL,
	)
	taskRunner.StartWorker(
		http_client_e2e_properties.TREASURE_CHEST_TASK_NAME,
		examples.OpenTreasureChest,
		http_client_e2e_properties.WORKER_THREAD_COUNT,
		http_client_e2e_properties.WORKER_POLLING_INTERVAL,
	)
}

func TestTaskRunnerExecution(t *testing.T) {
	workflowIdList := http_client_e2e.StartWorkflows(
		t,
		http_client_e2e_properties.WORKFLOW_EXECUTION_AMOUNT,
		http_client_e2e_properties.WORKFLOW_NAME,
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
	workflow := http_client_e2e.GetWorkflowExecutionStatus(
		t,
		workflowId,
	)
	if workflow.Status != "COMPLETED" {
		t.Error("Incomplete workflow: ", workflowId)
	}
}
