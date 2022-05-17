package http_client_e2e

import (
	"context"
	"testing"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/model/enum/task_result_status"
	"github.com/conductor-sdk/conductor-go/tests/e2e/e2e_properties"
	"github.com/conductor-sdk/conductor-go/tests/e2e/http_client_e2e/http_client_e2e_properties"
)

var taskClient = conductor_http_client.TaskResourceApiService{
	APIClient: e2e_properties.API_CLIENT,
}

func TestUpdateTaskRefByName(t *testing.T) {
	workflowId := StartWorkflow(t, http_client_e2e_properties.WORKFLOW_NAME)
	_, _, err := taskClient.UpdateTaskByRefName(
		context.Background(),
		http_client_e2e_properties.TASK_OUTPUT,
		workflowId,
		http_client_e2e_properties.TASK_REFERENCE_NAME,
		string(task_result_status.COMPLETED),
	)
	if err != nil {
		t.Error(err)
	}
	workflow := GetWorkflowExecutionStatus(t, workflowId)
	if workflow.Status != string(task_result_status.COMPLETED) {
		t.Error("Workflow status is not completed: ", workflow.Status)
	}
}
