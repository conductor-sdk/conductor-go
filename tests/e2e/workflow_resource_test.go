package e2e

import (
	"context"
	"testing"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
)

var workflowClient = conductor_http_client.WorkflowResourceApiService{
	APIClient: API_CLIENT,
}

func startWorkflows(t *testing.T, workflowQty int, workflowName string) []string {
	workflowIdList := make([]string, workflowQty)
	for i := 0; i < workflowQty; i += 1 {
		workflowIdList[i] = startWorkflow(t, workflowName)
	}
	return workflowIdList
}

func startWorkflow(t *testing.T, workflowName string) string {
	workflowId, response, err := workflowClient.StartWorkflow(
		context.Background(),
		make(map[string]interface{}),
		workflowName,
		nil,
	)
	if err != nil {
		t.Error("workflowId: ", workflowId, ", response: ", response, ", error: ", err)
	}
	return workflowId
}

func getWorkflowExecutionStatus(t *testing.T, workflowId string) http_model.Workflow {
	workflow, response, err := workflowClient.GetExecutionStatus(
		context.Background(),
		workflowId,
		nil,
	)
	if err != nil {
		t.Error("workflow: ", workflow, ", response: ", response, ", error: ", err)
	}
	return workflow
}
