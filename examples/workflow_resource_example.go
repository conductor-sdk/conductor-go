package examples

import (
	"context"
	client2 "github.com/conductor-sdk/conductor-go/client"
	settings2 "github.com/conductor-sdk/conductor-go/settings"
	"net/http"

	"github.com/antihax/optional"
)

func StartWorkflow(workflowName string, version optional.Int32, correlationId optional.String) (string, *http.Response, error) {
	workflowClient := client2.WorkflowResourceApiService{
		APIClient: client2.NewAPIClient(
			settings2.NewAuthenticationSettings(
				"key",
				"id",
			),
			settings2.NewHttpDefaultSettings(),
		),
	}
	return workflowClient.StartWorkflow(
		context.Background(),
		make(map[string]interface{}),
		workflowName,
		&client2.WorkflowResourceApiStartWorkflowOpts{
			Version:       version,
			CorrelationId: correlationId,
		},
	)
}
