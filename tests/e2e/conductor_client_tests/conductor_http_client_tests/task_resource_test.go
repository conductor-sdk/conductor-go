package conductor_http_client_tests

import (
	"context"
	"testing"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
)

var taskClient = conductor_http_client.TaskResourceApiService{
	APIClient: apiClient,
}

func TestUpdateTaskRefByName(t *testing.T) {
	workflowId := startWorkflow(t, WORKFLOW_NAME)
	_ = updateTaskByRefName(
		t,
		TASK_OUTPUT,
		workflowId,
		TASK_REFERENCE_NAME,
		"COMPLETED",
	)
	workflow := getWorkflowExecutionStatus(t, workflowId)
	if workflow.Status != "COMPLETED" {
		t.Error("Workflow status is not completed: ", workflow.Status)
	}
}

func updateTaskByRefName(t *testing.T, taskOutput map[string]interface{}, workflowId string, taskReferenceName string, status string) string {
	returnValue, response, err := taskClient.UpdateTaskByRefName(
		context.Background(),
		taskOutput,
		workflowId,
		taskReferenceName,
		status,
	)
	if err != nil {
		t.Error("returnValue: ", returnValue, ", response: ", response, ", error: ", err)
	}
	return returnValue
}
