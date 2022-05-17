package e2e

import (
	"context"
	"testing"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/model/enum/task_result_status"
	"github.com/conductor-sdk/conductor-go/tests"
)

var taskClient = conductor_http_client.TaskResourceApiService{
	APIClient: API_CLIENT,
}

func TestUpdateTaskRefByName(t *testing.T) {
	workflowId := startWorkflow(t, tests.WORKFLOW_NAME)
	_ = updateTaskByRefName(
		t,
		tests.TASK_OUTPUT,
		workflowId,
		tests.TASK_REFERENCE_NAME,
		string(task_result_status.COMPLETED),
	)
	workflow := getWorkflowExecutionStatus(t, workflowId)
	if workflow.Status != string(task_result_status.COMPLETED) {
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
