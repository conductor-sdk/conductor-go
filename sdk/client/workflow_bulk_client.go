package client

import (
	"context"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"net/http"
)

type WorkflowBulkClient interface {
	PauseWorkflow1(ctx context.Context, workflowIds []string) (model.BulkResponse, *http.Response, error)
	Restart(ctx context.Context, workflowIds []string, opts *WorkflowBulkResourceApiRestart1Opts) (model.BulkResponse, *http.Response, error)
	ResumeWorkflow(ctx context.Context, workflowIds []string) (model.BulkResponse, *http.Response, error)
	Retry1(ctx context.Context, workflowIds []string) (model.BulkResponse, *http.Response, error)
	Terminate(ctx context.Context, workflowIds []string, opts *WorkflowBulkResourceApiTerminateOpts) (model.BulkResponse, *http.Response, error)
}

func NewWorkflowBulkClient(client *APIClient) WorkflowBulkClient {
	return &WorkflowBulkResourceApiService{client}
}
