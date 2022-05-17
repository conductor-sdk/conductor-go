package http_client_e2e

import (
	"context"
	"testing"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/tests/e2e/e2e_properties"
)

var workflowClient = conductor_http_client.WorkflowResourceApiService{
	APIClient: e2e_properties.API_CLIENT,
}

func StartWorkflows(t *testing.T, workflowQty int, workflowName string) []string {
	workflowIdList := make([]string, workflowQty)
	for i := 0; i < workflowQty; i += 1 {
		workflowIdList[i] = StartWorkflow(t, workflowName)
	}
	return workflowIdList
}

func StartWorkflow(t *testing.T, workflowName string) string {
	workflowId, _, err := workflowClient.StartWorkflow(
		context.Background(),
		make(map[string]interface{}),
		workflowName,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	return workflowId
}

func GetWorkflowExecutionStatus(t *testing.T, workflowId string) http_model.Workflow {
	workflow, _, err := workflowClient.GetExecutionStatus(
		context.Background(),
		workflowId,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	return workflow
}
