package executor

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	log "github.com/sirupsen/logrus"
)

type WorkflowExecutor struct {
	metadataClient  *conductor_http_client.MetadataResourceApiService
	taskClient      *conductor_http_client.TaskResourceApiService
	workflowClient  *conductor_http_client.WorkflowResourceApiService
	workflowMonitor *WorkflowMonitor
}

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

func (e *WorkflowExecutor) RegisterWorkflow(workflow *http_model.WorkflowDef) (*http.Response, error) {
	return e.metadataClient.Update(
		context.Background(),
		[]http_model.WorkflowDef{
			*workflow,
		},
	)
}

func (e *WorkflowExecutor) StartWorkflow(request *http_model.StartWorkflowRequest) (string, WorkflowExecutionChannel, error) {
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

func (e *WorkflowExecutor) ExecuteWorkflow(workflow *http_model.WorkflowDef, request *http_model.StartWorkflowRequest) (string, WorkflowExecutionChannel, error) {
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

func WaitForWorkflowCompletionUntilTimeout(executionChannel WorkflowExecutionChannel, timeout time.Duration) (*http_model.Workflow, error) {
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
func (e *WorkflowExecutor) executeWorkflow(workflow *http_model.WorkflowDef, request *http_model.StartWorkflowRequest) (string, error) {

	//inputAsMap, err := getInputAsMap(request.Input)
	//if err != nil {
	//	return "", err
	//}
	startWorkflowRequest := http_model.StartWorkflowRequest{
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

func getInputAsMap(input interface{}) (map[string]interface{}, error) {
	if input == nil {
		return nil, nil
	}
	data, err := json.Marshal(input)
	if err != nil {
		log.Debug(
			"Failed to parse input",
			", reason: ", err.Error(),
		)
		return nil, err
	}
	var parsedInput map[string]interface{}
	json.Unmarshal(data, &parsedInput)
	return parsedInput, nil
}
