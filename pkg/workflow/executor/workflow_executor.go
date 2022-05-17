package executor

import (
	"context"
	"encoding/json"
	"net/http"
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

func (e *WorkflowExecutor) RegisterWorkflow(workflow *http_model.WorkflowDef) (*http.Response, error) {
	return e.metadataClient.Update(
		context.Background(),
		[]http_model.WorkflowDef{
			*workflow,
		},
	)
}

func (e *WorkflowExecutor) monitorRunningWorkflows() {
	for {
		e.notifyFinishedWorkflows()
		time.Sleep(RUNNING_WORKFLOWS_REFRESH_INTERVAL)
	}
}

func (e *WorkflowExecutor) notifyFinishedWorkflows() {
	for _, workflowId := range e.getRunningWorkflowIdList() {
		if workflow := e.getWorkflowIfFinished(workflowId); workflow != nil {
			e.notifyFinishedWorkflow(workflowId, workflow)
		}
	}
}

func (e *WorkflowExecutor) getWorkflowIfFinished(workflowId string) *http_model.Workflow {
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
		return nil
	}
	if !isWorkflowFinished(workflow) {
		return nil
	}
	return &workflow
}

func (e *WorkflowExecutor) notifyFinishedWorkflow(workflowId string, workflow *http_model.Workflow) {
	logrus.Debug("Workflow finished: ", workflowId)
	workflowExecutionChannel, ok := e.getWorkflowExecutionChannel(workflowId)
	if !ok {
		logrus.Error("Failed to get workflow execution channel for workflowId: ", workflowId)
	} else {
		workflowExecutionChannel <- *workflow
		close(workflowExecutionChannel)
		logrus.Debug("Remove workflow execution channel: ", workflowId)
		delete(e.runningWorkflowById, workflowId)
	}
}

func (e *WorkflowExecutor) addWorkflowExecutionChannel(workflowId string, workflowExecutionChannel WorkflowExecutionChannel) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	logrus.Debug("Add WorkflowExecutionChannel, workflowId: ", workflowId)
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

func (e *WorkflowExecutor) getWorkflowExecutionChannel(workflowId string) (WorkflowExecutionChannel, bool) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	workflowExecutionChannel, ok := e.runningWorkflowById[workflowId]
	return workflowExecutionChannel, ok
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
