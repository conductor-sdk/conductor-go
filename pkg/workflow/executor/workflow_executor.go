package executor

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/model"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	log "github.com/sirupsen/logrus"
)

type WorkflowExecutor struct {
	metadataClient  *conductor_http_client.MetadataResourceApiService
	taskClient      *conductor_http_client.TaskResourceApiService
	workflowClient  *conductor_http_client.WorkflowResourceApiService
	workflowMonitor *WorkflowMonitor
}

//Create a new workflow executor
func NewWorkflowExecutor(apiClient *conductor_http_client.APIClient) *WorkflowExecutor {
	workflowClient := &conductor_http_client.WorkflowResourceApiService{
		APIClient: apiClient,
	}
	workflowExecutor := WorkflowExecutor{
		metadataClient: &conductor_http_client.MetadataResourceApiService{
			APIClient: apiClient,
		},
		taskClient: &conductor_http_client.TaskResourceApiService{
			APIClient: apiClient,
		},
		workflowClient:  workflowClient,
		workflowMonitor: NewWorkflowMonitor(workflowClient),
	}
	return &workflowExecutor
}

//RegisterWorkflow Registers the workflow on the server.  Overwrites if the flag is set.  If the 'overwrite' flag is not set
//and the workflow definition differs from the one on the server, the call will fail with response code 409
func (e *WorkflowExecutor) RegisterWorkflow(overwrite bool, workflow *model.WorkflowDef) (*http.Response, error) {
	return e.metadataClient.RegisterWorkflowDef(
		context.Background(),
		overwrite,
		*workflow,
	)
}

//StartWorkflow Starts a new workflow execution and returns workflow id and channel on which workflow execution can be monitored.
func (e *WorkflowExecutor) StartWorkflow(request *model.StartWorkflowRequest) (string, WorkflowExecutionChannel, error) {
	workflowId, err := e.executeWorkflow(nil, request)
	if err != nil {
		return "", nil, err
	}
	executionChannel, err := e.workflowMonitor.GenerateWorkflowExecutionChannel(workflowId)
	if err != nil {
		return "", nil, err
	}
	return workflowId, executionChannel, nil
}

func (e *WorkflowExecutor) ExecuteWorkflow(workflow *model.WorkflowDef, request *model.StartWorkflowRequest) (string, WorkflowExecutionChannel, error) {
	workflowId, err := e.executeWorkflow(workflow, request)
	if err != nil {
		return "", nil, err
	}
	executionChannel, err := e.workflowMonitor.GenerateWorkflowExecutionChannel(workflowId)
	if err != nil {
		return "", nil, err
	}
	return workflowId, executionChannel, nil
}

func WaitForWorkflowCompletionUntilTimeout(executionChannel WorkflowExecutionChannel, timeout time.Duration) (*model.Workflow, error) {
	select {
	case workflow, ok := <-executionChannel:
		if !ok {
			return nil, fmt.Errorf("channel closed")
		}
		return workflow, nil
	case <-time.After(timeout):
		return nil, fmt.Errorf("timeout")
	}
}

// ExecuteWorkflow Executes a workflow
// Returns workflow Id for the newly started workflow
func (e *WorkflowExecutor) executeWorkflow(workflow *model.WorkflowDef, request *model.StartWorkflowRequest) (string, error) {
	startWorkflowRequest := model.StartWorkflowRequest{
		Name:                            request.Name,
		Version:                         request.Version,
		CorrelationId:                   request.CorrelationId,
		Input:                           request.Input,
		TaskToDomain:                    request.TaskToDomain,
		ExternalInputPayloadStoragePath: request.ExternalInputPayloadStoragePath,
		Priority:                        request.Priority,
	}
	if workflow != nil {
		startWorkflowRequest.WorkflowDef = workflow
	}
	workflowId, response, err := e.workflowClient.StartWorkflow1(
		context.Background(),
		startWorkflowRequest,
	)
	if err != nil {
		log.Debug(
			"Failed to start workflow",
			", reason: ", err.Error(),
			", name: ", request.Name,
			", version: ", request.Version,
			", input: ", request.Input,
			", workflowId: ", workflowId,
			", response: ", response,
		)
		return "", err
	}
	log.Debug(
		"Started workflow",
		", workflowId: ", workflowId,
		", name: ", request.Name,
		", version: ", request.Version,
		", input: ", request.Input,
	)
	return workflowId, err
}
