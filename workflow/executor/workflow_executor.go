package executor

import (
	"context"
	"fmt"
	client2 "github.com/conductor-sdk/conductor-go/client"
	"github.com/conductor-sdk/conductor-go/concurrency"
	model2 "github.com/conductor-sdk/conductor-go/model"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type WorkflowExecutor struct {
	metadataClient  *client2.MetadataResourceApiService
	taskClient      *client2.TaskResourceApiService
	workflowClient  *client2.WorkflowResourceApiService
	workflowMonitor *WorkflowMonitor
}

// NewWorkflowExecutor Create a new workflow executor
func NewWorkflowExecutor(apiClient *client2.APIClient) *WorkflowExecutor {
	workflowClient := &client2.WorkflowResourceApiService{
		APIClient: apiClient,
	}
	workflowExecutor := WorkflowExecutor{
		metadataClient: &client2.MetadataResourceApiService{
			APIClient: apiClient,
		},
		taskClient: &client2.TaskResourceApiService{
			APIClient: apiClient,
		},
		workflowClient:  workflowClient,
		workflowMonitor: NewWorkflowMonitor(workflowClient),
	}
	return &workflowExecutor
}

//RegisterWorkflow Registers the workflow on the server.  Overwrites if the flag is set.  If the 'overwrite' flag is not set
//and the workflow definition differs from the one on the server, the call will fail with response code 409
func (e *WorkflowExecutor) RegisterWorkflow(overwrite bool, workflow *model2.WorkflowDef) (*http.Response, error) {
	return e.metadataClient.RegisterWorkflowDef(
		context.Background(),
		overwrite,
		*workflow,
	)
}

//MonitorExecution monitors the workflow execution
//Returns the channel with the execution result of the workflow
//Note: Channels will continue to grow if the workflows do not complete and/or are not taken out
func (e *WorkflowExecutor) MonitorExecution(workflowId string) (workflowMonitor WorkflowExecutionChannel, err error) {
	return e.workflowMonitor.GenerateWorkflowExecutionChannel(workflowId)
}

//StartWorkflow Start workflows
//Returns the id of the newly created workflow
func (e *WorkflowExecutor) StartWorkflow(startWorkflowRequest *model2.StartWorkflowRequest) (workflowId string, err error) {
	id, _, err := e.workflowClient.StartWorkflowWithRequest(
		context.Background(),
		*startWorkflowRequest,
	)
	if err != nil {
		return "", err
	}
	return id, err
}

//StartWorkflows Start workflows in bulk
//Returns RunningWorkflow struct that contains the workflowId, Err (if failed to start) and an execution channel
//which can be used to monitor the completion of the workflow execution.  The channel is available if monitorExecution is set
func (e *WorkflowExecutor) StartWorkflows(monitorExecution bool, startWorkflowRequests ...*model2.StartWorkflowRequest) []*RunningWorkflow {
	amount := len(startWorkflowRequests)
	startingWorkflowChannel := make([]chan *RunningWorkflow, amount)
	for i := 0; i < amount; i += 1 {
		startingWorkflowChannel[i] = make(chan *RunningWorkflow)
		go e.startWorkflowDaemon(monitorExecution, startWorkflowRequests[i], startingWorkflowChannel[i])
	}
	startedWorkflows := make([]*RunningWorkflow, amount)
	for i := 0; i < amount; i += 1 {
		startedWorkflows[i] = <-startingWorkflowChannel[i]
	}
	return startedWorkflows
}

//WaitForWorkflowCompletionUntilTimeout Helper method to wait on the channel until the timeout for the workflow execution to complete
func WaitForWorkflowCompletionUntilTimeout(executionChannel WorkflowExecutionChannel, timeout time.Duration) (workflow *model2.Workflow, err error) {
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
func (e *WorkflowExecutor) executeWorkflow(workflow *model2.WorkflowDef, request *model2.StartWorkflowRequest) (workflowId string, err error) {
	startWorkflowRequest := model2.StartWorkflowRequest{
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

func (e *WorkflowExecutor) startWorkflowDaemon(monitorExecution bool, request *model2.StartWorkflowRequest, runningWorkflowChannel chan *RunningWorkflow) {
	defer concurrency.HandlePanicError("start_workflow")
	workflowId, err := e.executeWorkflow(nil, request)
	if err != nil {
		runningWorkflowChannel <- NewRunningWorkflow("", nil, err)
		return
	}
	if !monitorExecution {
		runningWorkflowChannel <- NewRunningWorkflow(workflowId, nil, nil)
		return
	}
	executionChannel, err := e.workflowMonitor.GenerateWorkflowExecutionChannel(workflowId)
	if err != nil {
		runningWorkflowChannel <- NewRunningWorkflow(workflowId, nil, err)
		return
	}
	runningWorkflowChannel <- NewRunningWorkflow(workflowId, executionChannel, nil)
}
