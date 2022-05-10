package orkestrator_tests

import (
	"testing"

	"github.com/conductor-sdk/conductor-go/pkg/orkestrator"
	"github.com/conductor-sdk/conductor-go/tests/e2e/conductor_client_tests"
)

func TestWorkerOrkestratorExecution(t *testing.T) {
	workflowName := "workflow_with_go_task_example_from_code"
	workflows := make([]string, 0)
	for workflowCounter := 0; workflowCounter < 5; workflowCounter += 1 {
		workflowId, err := conductor_client_tests.StartWorkflow(workflowName)
		if err != nil {
			t.Error("Failed to create workflow of type: ", workflowName)
		}
		workflows = append(workflows, workflowId)
	}
	apiClient := conductor_client_tests.GetApiClientWithAuthentication()
	workerOrkestrator := orkestrator.NewWorkerOrkestratorWithApiClient(apiClient)
	workerOrkestrator.StartWorker(
		"go_task_example_from_code",
	)
}
