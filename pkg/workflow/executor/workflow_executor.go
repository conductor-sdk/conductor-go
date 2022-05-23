package executor

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/concurrency"
	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/sirupsen/logrus"
)

type WorkflowExecutionChannel chan *http_model.Workflow

type executionChannel struct {
	ClientWorkflowExecutionChannel  WorkflowExecutionChannel
	TimeoutWorkflowExecutionChannel chan bool
}

type WorkflowExecutor struct {
	mutex                        sync.Mutex
	executionChannelByWorkflowId map[string]executionChannel

	taskClient     conductor_http_client.TaskResourceApiService
	workflowClient conductor_http_client.WorkflowResourceApiService
	metadataClient conductor_http_client.MetadataResourceApiService
}

const (
	monitorRunningWorkflowsRefreshInterval = 100 * time.Millisecond
)

func NewWorkflowExecutor(apiClient *conductor_http_client.APIClient) *WorkflowExecutor {
	workflowExecutor := WorkflowExecutor{
		executionChannelByWorkflowId: make(map[string]executionChannel),
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
	go workflowExecutor.monitorRunningWorkflows()
	return &workflowExecutor
}

func (e *WorkflowExecutor) ExecuteWorkflow(name string, version int32, input interface{}) (WorkflowExecutionChannel, error) {
	return e.ExecuteWorkflowWithTimeout(
		name,
		version,
		input,
		nil,
	)
}

func (e *WorkflowExecutor) ExecuteWorkflowWithTimeout(name string, version int32, input interface{}, timeout *time.Duration) (WorkflowExecutionChannel, error) {
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
		return nil, err
	}
	logrus.Debug("Workflow started: ", workflowId)
	return e.addWorkflowExecution(workflowId, timeout)
}

func (e *WorkflowExecutor) RegisterWorkflow(workflow *http_model.WorkflowDef) (*http.Response, error) {
	return e.metadataClient.Update(
		context.Background(),
		[]http_model.WorkflowDef{
			*workflow,
		},
	)
}

func (e *WorkflowExecutor) monitorRunningWorkflows() {
	defer concurrency.OnError("monitorRunningWorkflows")
	for {
		e.notifyFinishedWorkflows()
		time.Sleep(monitorRunningWorkflowsRefreshInterval)
	}
}

func (e *WorkflowExecutor) notifyFinishedWorkflows() {
	for _, workflowId := range e.getRunningWorkflowIdList() {
		workflow, response, err := e.workflowClient.GetExecutionStatus(
			context.Background(),
			workflowId,
			nil,
		)
		if err != nil {
			logrus.Warning(
				"Failed to get workflow execution status",
				"workflowId: ", workflowId,
				"response: ", response,
				"error: ", err.Error(),
			)
		} else if isWorkflowInTerminalState(&workflow) {
			e.notifyFinishedWorkflow(workflowId, &workflow)
		}
	}
}

func (e *WorkflowExecutor) notifyFinishedWorkflow(workflowId string, workflow *http_model.Workflow) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	logrus.Debug("Workflow finished: ", workflowId)
	executionChannel, ok := e.executionChannelByWorkflowId[workflowId]
	if !ok {
		logrus.Error("Failed to get execution channel for workflowId: ", workflowId)
	} else {
		executionChannel.ClientWorkflowExecutionChannel <- workflow
		executionChannel.TimeoutWorkflowExecutionChannel <- true
		logrus.Debug("Remove workflow execution channel: ", workflowId)
		delete(e.executionChannelByWorkflowId, workflowId)
	}
}

func (e *WorkflowExecutor) addWorkflowExecution(workflowId string, timeout *time.Duration) (WorkflowExecutionChannel, error) {
	logrus.Debug("addWorkflowExecution, workflowId: ", workflowId)
	workflowExecutionChannel := make(WorkflowExecutionChannel, 1)
	timeoutExecutionChannel := make(chan bool, 1)
	if timeout != nil {
		go e.monitorWorkflowExecution(
			workflowId,
			timeoutExecutionChannel,
			*timeout,
		)
	}
	e.mutex.Lock()
	e.executionChannelByWorkflowId[workflowId] = executionChannel{
		ClientWorkflowExecutionChannel:  workflowExecutionChannel,
		TimeoutWorkflowExecutionChannel: timeoutExecutionChannel,
	}
	e.mutex.Unlock()
	return workflowExecutionChannel, nil
}

func (e *WorkflowExecutor) getRunningWorkflowIdList() []string {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	i := 0
	runningWorkflowIdList := make([]string, len(e.executionChannelByWorkflowId))
	for workflowId := range e.executionChannelByWorkflowId {
		runningWorkflowIdList[i] = workflowId
		i += 1
	}
	return runningWorkflowIdList
}

func (e *WorkflowExecutor) monitorWorkflowExecution(workflowId string, timeoutWorkflowExecutionChannel chan bool, timeout time.Duration) {
	defer concurrency.OnError(
		fmt.Sprint(
			"monitorWorkflowExecution",
			", workflowId: ", workflowId,
			", timeout: ", timeout,
		),
	)
	select {
	case <-e.executionChannelByWorkflowId[workflowId].TimeoutWorkflowExecutionChannel:
		return
	case <-time.After(timeout):
		logrus.Warning("Timeout waiting for completion of workflow, with id: ", workflowId)
		e.notifyFinishedWorkflow(workflowId, nil)
	}
}
