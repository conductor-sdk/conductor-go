package worker_e2e

import (
	"os"
	"sync"
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/examples"
	"github.com/conductor-sdk/conductor-go/pkg/model/enum/workflow_status"
	"github.com/conductor-sdk/conductor-go/pkg/worker"
	"github.com/conductor-sdk/conductor-go/tests/e2e/e2e_properties"
	"github.com/conductor-sdk/conductor-go/tests/e2e/http_client_e2e"
	"github.com/conductor-sdk/conductor-go/tests/e2e/http_client_e2e/http_client_e2e_properties"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func TestTaskRunnerExecution(t *testing.T) {
	taskRunner := worker.NewTaskRunnerWithApiClient(e2e_properties.API_CLIENT)
	workflowIdList, err := http_client_e2e.StartWorkflows(
		http_client_e2e_properties.WORKFLOW_EXECUTION_AMOUNT,
		http_client_e2e_properties.WORKFLOW_NAME,
	)
	if err != nil {
		t.Fatal(err)
	}
	err = taskRunner.StartWorker(
		http_client_e2e_properties.TASK_NAME,
		examples.SimpleWorker,
		http_client_e2e_properties.WORKER_THREAD_COUNT,
		http_client_e2e_properties.WORKER_POLLING_INTERVAL,
	)
	if err != nil {
		t.Fatal(err)
	}
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(workflowIdList))
	for _, workflowId := range workflowIdList {
		go testValidateWorkflow(t, &waitGroup, workflowId)
	}
	waitGroup.Wait()
	if err != nil {
		t.Fatal(err)
	}
}

func testValidateWorkflow(t *testing.T, waitGroup *sync.WaitGroup, workflowId string) {
	defer waitGroup.Done()
	time.Sleep(10 * time.Second)
	workflow, _, err := http_client_e2e.GetWorkflowExecutionStatus(
		workflowId,
	)
	if err != nil {
		t.Error(err)
	}
	if workflow.Status != string(workflow_status.COMPLETED) {
		t.Error("Workflow finished with invalid terminal state, workflowId: ", workflowId)
	}
}
