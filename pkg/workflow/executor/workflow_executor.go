package executor

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/concurrency"
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

// NewWorkflowExecutor Create a new workflow executor
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

//StartWorkflow Start workflows
//Returns the id of the newly created workflow
func (e *WorkflowExecutor) StartWorkflow(startWorkflowRequest *model.StartWorkflowRequest) (workflowId string, error error) {
	id, _, err := e.workflowClient.StartWorkflowWithRequest(
		context.Background(),
		*startWorkflowRequest,
	)
	if err != nil {
		return "", err
	}
	return id, err
}

//MonitorExecution monitors the workflow execution
//Returns the channel with the execution result of the workflow
//Note: Channels will continue to grow if the workflows do not complete and/or are not taken out
func (e *WorkflowExecutor) MonitorExecution(workflowId string) (workflowMonitor WorkflowExecutionChannel) {
	return e.getWorkflowMonitorChannel(workflowId)
}

//StartWorkflows Start workflows in bulk
func (e *WorkflowExecutor) StartWorkflows(startWorkflowRequests []model.StartWorkflowRequest, monitorExecution bool) []*RunningWorkflow {
	amount := len(startWorkflowRequests)
	executionChannel := make(chan *RunningWorkflow, amount)
	for i := 0; i < amount; i += 1 {
		req := &startWorkflowRequests[i]
		go func() {
			workflowId, err := e.executeWorkflow(nil, req)
			executionChannel <- NewRunningWorkflow(workflowId, nil, err)
		}()
	}
	runningWorkflows := make([]*RunningWorkflow, amount)
	for i := 0; i < amount; i += 1 {
		runningWorkflows[i] = <-executionChannel
		if monitorExecution {
			runningWorkflows[i].WorkflowExecutionChannel = e.getWorkflowMonitorChannel(runningWorkflows[i].WorkflowId)
		}
	}

	return runningWorkflows
}

func (e *WorkflowExecutor) startWorkflows(startWorkflowRequests ...*model.StartWorkflowRequest) []*RunningWorkflow {
	amount := len(startWorkflowRequests)
	runningWorkflowsChannel := make([]chan *RunningWorkflow, amount)
	for i := 0; i < amount; i += 1 {
		runningWorkflowsChannel[i] = make(chan *RunningWorkflow)
		go e.startWorkflowDaemon(startWorkflowRequests[i], runningWorkflowsChannel[i])
	}
	runningWorkflows := make([]*RunningWorkflow, amount)
	for i := 0; i < amount; i += 1 {
		runningWorkflows[i] = <-runningWorkflowsChannel[i]
	}
	return runningWorkflows
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
	workflowId, response, err := e.workflowClient.StartWorkflowWithRequest(
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

func (e *WorkflowExecutor) startWorkflowDaemon(request *model.StartWorkflowRequest, runningWorkflowChannel chan *RunningWorkflow) {
	defer concurrency.HandlePanicError("start_workflow")
	workflowId, err := e.executeWorkflow(nil, request)
	if err != nil {
		runningWorkflowChannel <- NewRunningWorkflow("", nil, err)
	}
	executionChannel, err := e.workflowMonitor.GenerateWorkflowExecutionChannel(workflowId)
	if err != nil {
		runningWorkflowChannel <- NewRunningWorkflow("", nil, err)
	}
	runningWorkflowChannel <- NewRunningWorkflow(workflowId, executionChannel, nil)
}

func (e *WorkflowExecutor) getWorkflowMonitorChannel(workflowId string) WorkflowExecutionChannel {
	defer concurrency.HandlePanicError("monitor_workflow")
	channel, _ := e.workflowMonitor.GenerateWorkflowExecutionChannel(workflowId)
	return channel
}

func (e *WorkflowExecutor) startWorkflowsAndChannel(request *model.StartWorkflowRequest) {
	defer concurrency.HandlePanicError("start_workflow")
	runningWorkflowChannel := make(chan *RunningWorkflow)
	workflowId, err := e.executeWorkflow(nil, request)
	if err != nil {
		runningWorkflowChannel <- NewRunningWorkflow(workflowId, nil, err)
	}
	executionChannel, err := e.workflowMonitor.GenerateWorkflowExecutionChannel(workflowId)
	if err != nil {
		runningWorkflowChannel <- NewRunningWorkflow("", nil, err)
	}
	runningWorkflowChannel <- NewRunningWorkflow(workflowId, executionChannel, nil)
}
