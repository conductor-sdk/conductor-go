package client

import (
	"context"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"net/http"
)

type WorkflowBulkClient interface {
	PauseWorkflow1(ctx context.Context, body []string) (model.BulkResponse, *http.Response, error)
	Restart1(ctx context.Context, body []string, localVarOptionals *WorkflowBulkResourceApiRestart1Opts) (model.BulkResponse, *http.Response, error)
	ResumeWorkflow1(ctx context.Context, body []string) (model.BulkResponse, *http.Response, error)
	Retry1(ctx context.Context, body []string) (model.BulkResponse, *http.Response, error)
	Terminate(ctx context.Context, body []string, localVarOptionals *WorkflowBulkResourceApiTerminateOpts) (model.BulkResponse, *http.Response, error)
}
