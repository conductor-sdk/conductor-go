package http_client_e2e

import (
	"context"
	"testing"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/tests/e2e/e2e_properties"
	"github.com/conductor-sdk/conductor-go/tests/e2e/http_client_e2e/http_client_e2e_properties"
)

var metadataClient = conductor_http_client.MetadataResourceApiService{
	APIClient: e2e_properties.API_CLIENT,
}

func TestRegisterTaskDefinition(t *testing.T) {
	_, err := metadataClient.RegisterTaskDef(
		context.Background(),
		http_client_e2e_properties.TASK_DEFINITIONS,
	)
	if err != nil {
		t.Error(err)
	}
}

func TestRegisterWorkflowDefinition(t *testing.T) {
	response, err := metadataClient.RegisterWorkflowDef(
		context.Background(),
		http_client_e2e_properties.WORKFLOW_DEFINITION,
	)
	if err != nil && response.StatusCode != 409 {
		t.Error(err)
	}
}
