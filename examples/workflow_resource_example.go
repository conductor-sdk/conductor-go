package examples

import (
	"context"
	"github.com/conductor-sdk/conductor-go/pkg/client"
	"net/http"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/pkg/settings"
)

func StartWorkflow(workflowName string, version optional.Int32, correlationId optional.String) (string, *http.Response, error) {
	workflowClient := client.WorkflowResourceApiService{
		APIClient: client.NewAPIClient(
			settings.NewAuthenticationSettings(
				"key",
				"id",
			),
			settings.NewHttpDefaultSettings(),
		),
	}
	return workflowClient.StartWorkflow(
		context.Background(),
		make(map[string]interface{}),
		workflowName,
		&client.WorkflowResourceApiStartWorkflowOpts{
			Version:       version,
			CorrelationId: correlationId,
		},
	)
}
