package workflow

import (
	"encoding/json"
	"net/http"

	"github.com/conductor-sdk/conductor-go/pkg/concurrency"
	"github.com/conductor-sdk/conductor-go/pkg/model"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
	log "github.com/sirupsen/logrus"
)

type executeWorkflowResponse struct {
	ResponseValue            string
	WorkflowExecutionChannel executor.WorkflowExecutionChannel
	Err                      error
}

type TimeoutPolicy string

const (
	TimeOutWorkflow TimeoutPolicy = "TIME_OUT_WF"
	AlertOnly       TimeoutPolicy = "ALERT_ONLY"
)

type ConductorWorkflow struct {
	executor         *executor.WorkflowExecutor
	name             string
	version          int32
	description      string
	ownerEmail       string
	tasks            []TaskInterface
	timeoutPolicy    TimeoutPolicy
	timeoutSeconds   int64
	failureWorkflow  string
	inputParameters  []string
	outputParameters map[string]interface{}
	inputTemplate    map[string]interface{}
	variables        map[string]interface{}
	restartable      bool
}

func NewConductorWorkflow(executor *executor.WorkflowExecutor) *ConductorWorkflow {
	return &ConductorWorkflow{
		executor:      executor,
		timeoutPolicy: AlertOnly,
	}
}

func (workflow *ConductorWorkflow) Name(name string) *ConductorWorkflow {
	workflow.name = name
	return workflow
}

func (workflow *ConductorWorkflow) Version(version int32) *ConductorWorkflow {
	workflow.version = version
	return workflow
}

func (workflow *ConductorWorkflow) Description(description string) *ConductorWorkflow {
	workflow.description = description
	return workflow
}

func (workflow *ConductorWorkflow) TimeoutPolicy(timeoutPolicy TimeoutPolicy, timeoutSeconds int64) *ConductorWorkflow {
	workflow.timeoutPolicy = timeoutPolicy
	workflow.timeoutSeconds = timeoutSeconds
	return workflow
}

func (workflow *ConductorWorkflow) TimeoutSeconds(timeoutSeconds int64) *ConductorWorkflow {
	workflow.timeoutSeconds = timeoutSeconds
	return workflow
}

func (workflow *ConductorWorkflow) FailureWorkflow(failureWorkflow string) *ConductorWorkflow {
	workflow.failureWorkflow = failureWorkflow
	return workflow
}

func (workflow *ConductorWorkflow) Restartable(restartable bool) *ConductorWorkflow {
	workflow.restartable = restartable
	return workflow
}

func (workflow *ConductorWorkflow) OutputParameters(outputParameters map[string]interface{}) *ConductorWorkflow {
	workflow.outputParameters = outputParameters
	return workflow
}

func (workflow *ConductorWorkflow) InputTemplate(inputTemplate map[string]interface{}) *ConductorWorkflow {
	workflow.inputTemplate = inputTemplate
	return workflow
}

func (workflow *ConductorWorkflow) Variables(variables map[string]interface{}) *ConductorWorkflow {
	workflow.variables = variables
	return workflow
}

func (workflow *ConductorWorkflow) InputParameters(inputParameters ...string) *ConductorWorkflow {
	workflow.inputParameters = inputParameters
	return workflow
}

func (workflow *ConductorWorkflow) OwnerEmail(ownerEmail string) *ConductorWorkflow {
	workflow.ownerEmail = ownerEmail
	return workflow
}

func (workflow *ConductorWorkflow) GetName() string {
	return workflow.name
}

func (workflow *ConductorWorkflow) GetVersion() int32 {
	return workflow.version
}

func (workflow *ConductorWorkflow) Add(task TaskInterface) *ConductorWorkflow {
	workflow.tasks = append(workflow.tasks, task)
	return workflow
}

func (workflow *ConductorWorkflow) Register(override bool) (*http.Response, error) {
	return workflow.executor.RegisterWorkflow(override, workflow.ToWorkflowDef())
}

func (workflow *ConductorWorkflow) ExecuteWorkflowBulk(startWorkflowRequests ...model.StartWorkflowRequest) ([]executor.WorkflowExecutionChannel, error) {
	amount := len(startWorkflowRequests)
	executeWorkflowResponseChannels := make([]chan executeWorkflowResponse, amount)
	for i := 0; i < amount; i += 1 {
		executeWorkflowResponseChannels[i] = make(chan executeWorkflowResponse)
		go workflow.executeWorkflowDaemon(executeWorkflowResponseChannels[i], &startWorkflowRequests[i])
	}
	workflowExecutionChannelList := make([]executor.WorkflowExecutionChannel, amount)
	for i := 0; i < amount; i += 1 {
		executeWorkflowResponse := <-executeWorkflowResponseChannels[i]
		if executeWorkflowResponse.Err != nil {
			return nil, executeWorkflowResponse.Err
		}
		workflowExecutionChannelList[i] = executeWorkflowResponse.WorkflowExecutionChannel
	}
	return workflowExecutionChannelList, nil
}

func (workflow *ConductorWorkflow) ExecuteWorkflowWithInput(input interface{}) (string, executor.WorkflowExecutionChannel, error) {
	version := workflow.GetVersion()
	modelRequest := model.StartWorkflowRequest{
		Name:    workflow.GetName(),
		Version: &version,
		Input:   getInputAsMap(input),
	}
	return workflow.executor.ExecuteWorkflow(
		workflow.ToWorkflowDef(),
		&modelRequest,
	)
}

func (workflow *ConductorWorkflow) ExecuteWorkflow(startWorkflowRequest *model.StartWorkflowRequest) (string, executor.WorkflowExecutionChannel, error) {
	version := workflow.GetVersion()
	modelRequest := model.StartWorkflowRequest{
		Name:                            workflow.GetName(),
		Version:                         &version,
		CorrelationId:                   startWorkflowRequest.CorrelationId,
		Input:                           getInputAsMap(startWorkflowRequest.Input),
		TaskToDomain:                    startWorkflowRequest.TaskToDomain,
		ExternalInputPayloadStoragePath: startWorkflowRequest.ExternalInputPayloadStoragePath,
		Priority:                        startWorkflowRequest.Priority,
	}
	return workflow.executor.ExecuteWorkflow(
		workflow.ToWorkflowDef(),
		&modelRequest,
	)
}

func (workflow *ConductorWorkflow) executeWorkflowDaemon(executeWorkflowChannel chan executeWorkflowResponse, startWorkflowRequest *model.StartWorkflowRequest) {
	defer concurrency.HandlePanicError("execute_workflow")
	responseValue, workflowExecutionChannel, err := workflow.ExecuteWorkflow(startWorkflowRequest)
	executeWorkflowChannel <- executeWorkflowResponse{
		ResponseValue:            responseValue,
		WorkflowExecutionChannel: workflowExecutionChannel,
		Err:                      err,
	}
}

func getInputAsMap(input interface{}) map[string]interface{} {
	if input == nil {
		return nil
	}
	data, err := json.Marshal(input)
	if err != nil {
		log.Debug(
			"Failed to parse input",
			", reason: ", err.Error(),
		)
		return nil
	}
	var parsedInput map[string]interface{}
	json.Unmarshal(data, &parsedInput)
	return parsedInput
}

func (workflow *ConductorWorkflow) ToWorkflowDef() *model.WorkflowDef {
	return &model.WorkflowDef{
		Name:             workflow.name,
		Description:      workflow.description,
		Version:          workflow.version,
		Tasks:            getWorkflowTasksFromConductorWorkflow(workflow),
		InputParameters:  workflow.inputParameters,
		OutputParameters: workflow.outputParameters,
		FailureWorkflow:  workflow.failureWorkflow,
		SchemaVersion:    2,
		OwnerEmail:       workflow.ownerEmail,
		TimeoutPolicy:    string(workflow.timeoutPolicy),
		TimeoutSeconds:   workflow.timeoutSeconds,
		Variables:        workflow.variables,
		InputTemplate:    workflow.inputTemplate,
	}
}

func getWorkflowTasksFromConductorWorkflow(workflow *ConductorWorkflow) []model.WorkflowTask {
	workflowTasks := make([]model.WorkflowTask, 0)
	for _, task := range workflow.tasks {
		workflowTasks = append(
			workflowTasks,
			task.toWorkflowTask()...,
		)
	}
	return workflowTasks
}
