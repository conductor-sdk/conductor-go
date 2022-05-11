package worker_tests

import (
	"testing"
)

func TestWorkerOrkestratorExecution(t *testing.T) {
	// workflowName := "workflow_with_go_task_example_from_code"
	// workflows := make([]string, 0)
	// for workflowCounter := 0; workflowCounter < 5; workflowCounter += 1 {
	// 	workflowId, err := conductor_client_tests.StartWorkflow(workflowName)
	// 	if err != nil {
	// 		t.Error("Failed to create workflow of type: ", workflowName)
	// 	}
	// 	workflows = append(workflows, workflowId)
	// }
	// apiClient := conductor_client_tests.GetApiClientWithAuthentication()
	// workerOrkestrator := worker.NewWorkerOrkestratorWithApiClient(apiClient)
	// workerOrkestrator.StartWorker(
	// 	"go_task_example_from_code",
	// 	task_execute_function.Example1,
	// 	3,
	// 	1000,
	// )
	// time.Sleep(
	// 	time.Duration(10000) * time.Millisecond,
	// )
	// logrus.Warning("workflows: ", workflows)
}
