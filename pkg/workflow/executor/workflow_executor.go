package executor

import (
	"context"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
)

type RunningWorkflow chan http_model.Workflow

type WorkflowExecutor struct {
	runningWorkflowById map[string]RunningWorkflow

	taskClient     conductor_http_client.TaskResourceApiService
	workflowClient conductor_http_client.WorkflowResourceApiService
	metadataClient conductor_http_client.MetadataResourceApiService
}

func NewWorkflowExecutor(apiClient *conductor_http_client.APIClient) *WorkflowExecutor {
	return &WorkflowExecutor{
		runningWorkflowById: make(map[string]RunningWorkflow),
		taskClient: conductor_http_client.TaskResourceApiService{
			APIClient: apiClient,
		},
		workflowClient: conductor_http_client.WorkflowResourceApiService{
			APIClient: apiClient,
		},
		metadataClient: conductor_http_client.MetadataResourceApiService{
			APIClient: apiClient,
		},
	}
}

func (e *WorkflowExecutor) ExecuteWorkflow(name string, version int32, input map[string]interface{}) (RunningWorkflow, error) {
	startWorkflowRequest := http_model.StartWorkflowRequest{
		Name:    name,
		Version: version,
		Input:   input,
	}
	workflowId, _, err := e.workflowClient.StartWorkflow1(
		context.Background(),
		startWorkflowRequest,
	)
	if err != nil {
		return nil, err
	}
	e.runningWorkflowById[workflowId] = make(RunningWorkflow)
	return e.runningWorkflowById[workflowId], nil
}
