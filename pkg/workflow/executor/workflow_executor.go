package executor

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/model"
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

func (e *WorkflowExecutor) ExecuteWorkflow(name string, version int32, input interface{}) (string, WorkflowExecutionChannel, error) {
	workflowId, err := e.startWorkflow(
		name,
		version,
		input,
	)
	if err != nil {
		return "", nil, err
	}
	executionChannel, err := e.workflowMonitor.GenerateWorkflowExecutionChannel(workflowId)
	if err != nil {
		return "", nil, err
	}
	return workflowId, executionChannel, nil
}

func (e *WorkflowExecutor) RegisterWorkflow(workflow *http_model.WorkflowDef) (*http.Response, error) {
	return e.metadataClient.Update(
		context.Background(),
		[]http_model.WorkflowDef{
			*workflow,
		},
	)
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

func (e *WorkflowExecutor) startWorkflow(name string, version int32, input interface{}) (string, error) {
	inputAsMap, err := model.ConvertToMap(input)
	if err != nil {
		return "", err
	}
	startWorkflowRequest := http_model.StartWorkflowRequest{
		Name:    name,
		Version: version,
		Input:   inputAsMap,
	}
	workflowId, response, err := e.workflowClient.StartWorkflow1(
		context.Background(),
		startWorkflowRequest,
	)
	if err != nil {
		log.Debug(
			"Failed to start workflow",
			", reason: ", err.Error(),
			", name: ", name,
			", version: ", version,
			", input: ", input,
			", workflowId: ", workflowId,
			", response: ", response,
		)
		return "", err
	}
	log.Debug(
		"Started workflow",
		", workflowId: ", workflowId,
		", name: ", name,
		", version: ", version,
		", input: ", input,
	)
	return workflowId, err
}
