package executor

import (
	"context"
	"encoding/json"
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

func (e *WorkflowExecutor) ExecuteWorkflow(name string, version int32, input interface{}) (WorkflowExecutionChannel, error) {
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
	workflowExecutionChannel := make(WorkflowExecutionChannel)
	e.addWorkflowExecutionChannel(workflowId, workflowExecutionChannel)
	return workflowExecutionChannel, nil
}

func (e *WorkflowExecutor) RegisterWorkflow(workflow *http_model.WorkflowDef) error {
	response, error := e.metadataClient.RegisterWorkflowDef(context.Background(), *workflow)
	if response.StatusCode != 200 {
		return error
	}
	return nil
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
			logrus.Debug("Workflow finished: ", workflowId)
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
	for _, workflowId := range finishedWorkflowIdList {
		delete(e.runningWorkflowById, workflowId)
	}
}

func (e *WorkflowExecutor) notifyWorkflowExecutionStatus(workflowId string, workflow http_model.Workflow) {
	logrus.Debug("notifyWorkflowExecutionStatus, workflow: ", workflow)
	if workflowExecutionChannel, ok := e.runningWorkflowById[workflowId]; ok {
		workflowExecutionChannel <- workflow
		close(workflowExecutionChannel)
	}
}

func (e *WorkflowExecutor) addWorkflowExecutionChannel(workflowId string, workflowExecutionChannel WorkflowExecutionChannel) {
	logrus.Debug("addWorkflowExecutionChannel, workflowId: ", workflowId)
	e.runningWorkflowById[workflowId] = workflowExecutionChannel
}

func (e *WorkflowExecutor) getRunningWorkflowIdList() []string {
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

func getInputAsMap(input interface{}) map[string]interface{} {
	if input == nil {
		return nil
	}
	data, _ := json.Marshal(input)
	var parsedInput map[string]interface{}
	json.Unmarshal(data, &parsedInput)
	return parsedInput
}
