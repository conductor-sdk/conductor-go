package e2e

import (
	"context"
	"testing"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/tests"
)

var metadataClient = conductor_http_client.MetadataResourceApiService{
	APIClient: API_CLIENT,
}

func TestRegisterTaskDefinition(t *testing.T) {
	registerTaskDefinition(t, tests.TASK_DEFINITIONS)
}

func TestRegisterWorkflowDefinition(t *testing.T) {
	for _, workflowDefinition := range tests.WORKFLOW_DEFINITIONS {
		registerWorkflowDefinition(t, workflowDefinition)
	}
}

func registerTaskDefinition(t *testing.T, taskDefinitionList []http_model.TaskDef) {
	response, err := metadataClient.RegisterTaskDef(
		context.Background(),
		taskDefinitionList,
	)
	if err != nil {
		t.Error("response: ", response, "err: ", err)
	}
}

func registerWorkflowDefinition(t *testing.T, workflowDefinition http_model.WorkflowDef) {
	response, err := metadataClient.RegisterWorkflowDef(
		context.Background(),
		workflowDefinition,
	)
	if err != nil && response.StatusCode != 409 {
		t.Error("response: ", response, ", error: ", err)
	}
}
