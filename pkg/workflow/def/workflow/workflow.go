package workflow

import (
	"encoding/json"
	"net/http"

	"github.com/conductor-sdk/conductor-go/pkg/model"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
	log "github.com/sirupsen/logrus"
)

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

//FailureWorkflow name of the workflow to execute when this workflow fails.
//Failure workflows can be used for handling compensation logic
func (workflow *ConductorWorkflow) FailureWorkflow(failureWorkflow string) *ConductorWorkflow {
	workflow.failureWorkflow = failureWorkflow
	return workflow
}

//Restartable if the workflow can be restarted after it has reached terminal state.
//Set this to false if restarting workflow can have side effects
func (workflow *ConductorWorkflow) Restartable(restartable bool) *ConductorWorkflow {
	workflow.restartable = restartable
	return workflow
}

//OutputParameters Workflow outputs. Workflow output follows similar structure as task inputs
//See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for more details
func (workflow *ConductorWorkflow) OutputParameters(outputParameters map[string]interface{}) *ConductorWorkflow {
	workflow.outputParameters = outputParameters
	return workflow
}

//InputTemplate template input to the workflow.  Can have combination of variables (e.g. ${workflow.input.abc}) and
//static values
func (workflow *ConductorWorkflow) InputTemplate(inputTemplate map[string]interface{}) *ConductorWorkflow {
	workflow.inputTemplate = inputTemplate
	return workflow
}

//Variables Workflow variables are set using SET_VARIABLE task.  Excellent way to maintain business state
//e.g. Variables can maintain business/user specific states which can be queried and inspected to find out the state of the workflow
func (workflow *ConductorWorkflow) Variables(variables map[string]interface{}) *ConductorWorkflow {
	workflow.variables = variables
	return workflow
}

//InputParameters List of the input parameters to the workflow.  Used ONLY for the documentation purpose.
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

//Register the workflow definition with the server. If overwrite is set, the definition on the server will be overwritten.
//When not set, the call fails if there is any change in the workflow definition between the server and what is being registered.
func (workflow *ConductorWorkflow) Register(overwrite bool) (*http.Response, error) {
	return workflow.executor.RegisterWorkflow(overwrite, workflow.ToWorkflowDef())
}

// TODO update description
//ExecuteWorkflowWithInput Execute the workflow with specific input.  The input struct MUST be serializable to JSON
//Returns the workflow Id that can be used to monitor and get the status of the workflow execution
//Optionally, for short-lived workflows the channel can be used to monitor the status of the workflow
func (workflow *ConductorWorkflow) StartWorkflowWithInput(input interface{}) ([]*executor.RunningWorkflow, error) {
	return workflow.executor.StartWorkflow(
		&model.StartWorkflowRequest{
			Name:        workflow.GetName(),
			Version:     workflow.GetVersion(),
			Input:       getInputAsMap(input),
			WorkflowDef: workflow.ToWorkflowDef(),
		},
	)
}

//ExecuteWorkflow Execute the workflow with start request, that allows you to pass more details like correlationId, domain mapping etc.
//Returns the workflow Id that can be used to monitor and get the status of the workflow execution
//Optionally, for short-lived workflows the channel can be used to monitor the status of the workflow
func (workflow *ConductorWorkflow) StartWorkflow(startWorkflowRequests ...*model.StartWorkflowRequest) ([]*executor.RunningWorkflow, error) {
	for i := range startWorkflowRequests {
		startWorkflowRequests[i] = workflow.decorateStartWorkflowRequest(startWorkflowRequests[i])
	}
	return workflow.executor.StartWorkflow(startWorkflowRequests...)
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

//ToWorkflowDef converts the workflow to the JSON serializable format
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

func (workflow *ConductorWorkflow) decorateStartWorkflowRequest(startWorkflowRequest *model.StartWorkflowRequest) *model.StartWorkflowRequest {
	return &model.StartWorkflowRequest{
		Name:                            workflow.GetName(),
		Version:                         workflow.version,
		CorrelationId:                   startWorkflowRequest.CorrelationId,
		Input:                           getInputAsMap(startWorkflowRequest.Input),
		WorkflowDef:                     startWorkflowRequest.WorkflowDef,
		TaskToDomain:                    startWorkflowRequest.TaskToDomain,
		ExternalInputPayloadStoragePath: startWorkflowRequest.ExternalInputPayloadStoragePath,
		Priority:                        startWorkflowRequest.Priority,
	}
}
