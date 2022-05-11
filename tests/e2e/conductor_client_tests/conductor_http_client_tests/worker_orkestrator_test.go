package conductor_http_client_tests

import (
	"testing"
	"time"

	"github.com/conductor-sdk/conductor-go/examples/task_execute_function"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/worker"
)

func TestWorkerOrkestratorExecution(t *testing.T) {
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
	workflowIdList := make([]string, 0)
	for i := 0; i < workflowQty; i += 1 {
		workflowId := startWorkflow(t, WORKFLOW_NAME)
		workflowIdList = append(workflowIdList, workflowId)
	}
	taskRunner := worker.NewWorkerOrkestratorWithApiClient(
		apiClient,
	)
	workerThreadCount := 5
	workerPollingInterval := 1000
	taskRunner.StartWorker(
		TASK_NAME,
		task_execute_function.Example1,
		workerThreadCount,
		workerPollingInterval,
	)
	total := workflowQty * workerPollingInterval / workerThreadCount
	time.Sleep(
		time.Duration(total<<1) * time.Millisecond,
	)
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
