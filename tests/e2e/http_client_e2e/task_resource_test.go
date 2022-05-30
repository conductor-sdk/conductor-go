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
	workflowId, response, err := StartWorkflow(http_client_e2e_properties.WORKFLOW_NAME)
	if err != nil {
		t.Fatal(
			"workflowId: ", workflowId,
			", response:, ", *response,
			", error: ", err.Error(),
		)
	}
	value, response, err := taskClient.UpdateTaskByRefName(
		context.Background(),
		http_client_e2e_properties.TASK_OUTPUT,
		workflowId,
		http_client_e2e_properties.TASK_NAME,
		string(task_result_status.COMPLETED),
	)
	if err != nil {
		t.Fatal(
			"value: ", value,
			", response:, ", *response,
			", error: ", err.Error(),
		)
	}
	workflow, response, err := GetWorkflowExecutionStatus(workflowId)
	if err != nil {
		t.Fatal(
			"workflow: ", workflow,
			", response:, ", *response,
			", error: ", err.Error(),
		)
	}
	if workflow.Status != string(task_result_status.COMPLETED) {
		t.Fatal(
			"Workflow status is not completed: ", workflow.Status,
		)
	}
}
