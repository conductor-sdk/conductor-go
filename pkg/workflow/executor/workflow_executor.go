package executor

import (
	"context"
	"net/http"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/sirupsen/logrus"
)

type WorkflowExecutor struct {
	metadataClient  *conductor_http_client.MetadataResourceApiService
	taskClient      *conductor_http_client.TaskResourceApiService
	workflowClient  *conductor_http_client.WorkflowResourceApiService
	workflowMonitor *WorkflowMonitor
}

func NewWorkflowExecutor(apiClient *conductor_http_client.APIClient) *WorkflowExecutor {
	workflowClient := &conductor_http_client.WorkflowResourceApiService{apiClient}
	workflowExecutor := WorkflowExecutor{
		metadataClient:  &conductor_http_client.MetadataResourceApiService{apiClient},
		taskClient:      &conductor_http_client.TaskResourceApiService{apiClient},
		workflowClient:  workflowClient,
		workflowMonitor: NewWorkflowMonitor(workflowClient),
	}
	go workflowExecutor.workflowMonitor.MonitorRunningWorkflows()
	return &workflowExecutor
}

func (e *WorkflowExecutor) ExecuteWorkflow(name string, version int32, input interface{}) (WorkflowExecutionChannel, error) {
	workflowId, err := e.startWorkflow(
		name,
		version,
		input,
	)
	if err != nil {
		return nil, err
	}
	return e.workflowMonitor.GenerateWorkflowExecutionChannel(workflowId)
}

func (e *WorkflowExecutor) ExecuteWorkflowWithTimeout(name string, version int32, input interface{}, timeout time.Duration) (WorkflowExecutionChannel, error) {
	workflowId, err := e.startWorkflow(
		name,
		version,
		input,
	)
	if err != nil {
		return nil, err
	}
	return e.workflowMonitor.GenerateWorkflowExecutionChannelWithTimeout(workflowId, timeout)
}

func (e *WorkflowExecutor) RegisterWorkflow(workflow *http_model.WorkflowDef) (*http.Response, error) {
	return e.metadataClient.Update(
		context.Background(),
		[]http_model.WorkflowDef{
			*workflow,
		},
	)
}

func (e *WorkflowExecutor) startWorkflow(name string, version int32, input interface{}) (string, error) {
	startWorkflowRequest := http_model.StartWorkflowRequest{
		Name:    name,
		Version: version,
		Input:   getInputAsMap(input),
	}
	workflowId, _, err := e.workflowClient.StartWorkflow1(
		context.Background(),
		startWorkflowRequest,
	)
	if err != nil {
		return "", err
	}
	logrus.Debug(
		"Started workflow",
		", name: ", name,
		", version: ", version,
		", input: ", getInputAsMap(input),
		", workflowId: ", workflowId,
	)
	return workflowId, err
}
