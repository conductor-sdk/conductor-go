package executor

import (
	"context"
	"sync"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/conductor_client/conductor_http_client"
	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/sirupsen/logrus"
)

var (
	RUNNING_WORKFLOWS_REFRESH_INTERVAL = 100 * time.Millisecond
)

type WorkflowExecutionChannel chan http_model.Workflow

type WorkflowExecutor struct {
	mutex               sync.Mutex
	runningWorkflowById map[string]WorkflowExecutionChannel

	taskClient     conductor_http_client.TaskResourceApiService
	workflowClient conductor_http_client.WorkflowResourceApiService
	metadataClient conductor_http_client.MetadataResourceApiService
}

func NewWorkflowExecutor(apiClient *conductor_http_client.APIClient) *WorkflowExecutor {
	workflowExecutor := WorkflowExecutor{
		runningWorkflowById: make(map[string]WorkflowExecutionChannel),
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

func (e *WorkflowExecutor) ExecuteWorkflow(name string, version int32, input map[string]interface{}) (WorkflowExecutionChannel, error) {
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
	workflowExecutionChannel := make(WorkflowExecutionChannel)
	e.addWorkflowExecutionChannel(workflowId, workflowExecutionChannel)
	return workflowExecutionChannel, nil
}

func (e *WorkflowExecutor) monitorRunningWorkflows() {
	for {
		finishedWorkflowIdList := e.getFinishedWorkflowIdList()
		e.removeFinishedWorkflows(finishedWorkflowIdList)
		time.Sleep(RUNNING_WORKFLOWS_REFRESH_INTERVAL)
	}
}

func (e *WorkflowExecutor) getFinishedWorkflowIdList() []string {
	finishedWorkflowIdList := make([]string, 0)
	for _, workflowId := range e.getRunningWorkflowIdList() {
		if e.isWorkflowFinished(workflowId) {
			finishedWorkflowIdList = append(finishedWorkflowIdList, workflowId)
		}
	}
	return finishedWorkflowIdList
}

func (e *WorkflowExecutor) isWorkflowFinished(workflowId string) bool {
	workflow, response, err := e.workflowClient.GetExecutionStatus(
		context.Background(),
		workflowId,
		nil,
	)
	if err != nil {
		logrus.Debug(
			"Failed to get workflow execution status",
			"workflowId: ", workflowId,
			"response: ", response,
			"error: ", err.Error(),
		)
		return false
	}
	if !isWorkflowFinished(workflow) {
		return false
	}
	e.notifyWorkflowExecutionStatus(workflowId, workflow)
	return true
}

func (e *WorkflowExecutor) removeFinishedWorkflows(finishedWorkflowIdList []string) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	for _, workflowId := range finishedWorkflowIdList {
		delete(e.runningWorkflowById, workflowId)
	}
}

func (e *WorkflowExecutor) notifyWorkflowExecutionStatus(workflowId string, workflow http_model.Workflow) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	if workflowExecutionChannel, ok := e.runningWorkflowById[workflowId]; ok {
		workflowExecutionChannel <- workflow
	}
}

func (e *WorkflowExecutor) addWorkflowExecutionChannel(workflowId string, workflowExecutionChannel WorkflowExecutionChannel) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.runningWorkflowById[workflowId] = workflowExecutionChannel
}

func (e *WorkflowExecutor) getRunningWorkflowIdList() []string {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	i := 0
	runningWorkflowIdList := make([]string, len(e.runningWorkflowById))
	for workflowId := range e.runningWorkflowById {
		runningWorkflowIdList[i] = workflowId
		i += 1
	}
	return runningWorkflowIdList
}

func isWorkflowFinished(workflow http_model.Workflow) bool {
	return workflow.Status != "PAUSED" && workflow.Status != "RUNNING"
}
