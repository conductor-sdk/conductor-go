package def

import (
	"encoding/json"
	"github.com/conductor-sdk/conductor-go/sdk/model"
	executor3 "github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type TimeoutPolicy string

const (
	TimeOutWorkflow TimeoutPolicy = "TIME_OUT_WF"
	AlertOnly       TimeoutPolicy = "ALERT_ONLY"
)

type ConductorWorkflow struct {
	executor         *executor3.WorkflowExecutor
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

func NewConductorWorkflow(executor *executor3.WorkflowExecutor) *ConductorWorkflow {
	return &ConductorWorkflow{
		executor:      executor,
		timeoutPolicy: AlertOnly,
		restartable:   true,
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
func (workflow *ConductorWorkflow) OutputParameters(outputParameters interface{}) *ConductorWorkflow {
	workflow.outputParameters = getInputAsMap(outputParameters)
	return workflow
}

//InputTemplate template input to the workflow.  Can have combination of variables (e.g. ${workflow.input.abc}) and
//static values
func (workflow *ConductorWorkflow) InputTemplate(inputTemplate interface{}) *ConductorWorkflow {
	workflow.inputTemplate = getInputAsMap(inputTemplate)
	return workflow
}

//Variables Workflow variables are set using SET_VARIABLE task.  Excellent way to maintain business state
//e.g. Variables can maintain business/user specific states which can be queried and inspected to find out the state of the workflow
func (workflow *ConductorWorkflow) Variables(variables interface{}) *ConductorWorkflow {
	workflow.variables = getInputAsMap(variables)
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

func (workflow *ConductorWorkflow) GetName() (name string) {
	return workflow.name
}

func (workflow *ConductorWorkflow) GetVersion() (version int32) {
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

// StartWorkflowWithInput ExecuteWorkflowWithInput Execute the workflow with specific input.  The input struct MUST be serializable to JSON
//Returns the workflow Id that can be used to monitor and get the status of the workflow execution
func (workflow *ConductorWorkflow) StartWorkflowWithInput(input interface{}) (workflowId string, err error) {
	version := workflow.GetVersion()
	return workflow.executor.StartWorkflow(
		&model.StartWorkflowRequest{
			Name:        workflow.GetName(),
			Version:     &version,
			Input:       getInputAsMap(input),
			WorkflowDef: workflow.ToWorkflowDef(),
		},
	)
}

//StartWorkflow starts the workflow execution with startWorkflowRequest that allows you to specify more details like task domains, correlationId etc.
//Returns the ID of the newly created workflow
func (workflow *ConductorWorkflow) StartWorkflow(startWorkflowRequest *model.StartWorkflowRequest) (workflowId string, err error) {
	startWorkflowRequest.WorkflowDef = workflow.ToWorkflowDef()
	return workflow.executor.StartWorkflow(startWorkflowRequest)
}

//StartWorkflowsAndMonitorExecution Starts the workflow execution and returns a channel that can be used to monitor the workflow execution
//This method is useful for short duration workflows that are expected to complete in few seconds.  For long-running workflows use GetStatus APIs to periodically check the status
func (workflow *ConductorWorkflow) StartWorkflowsAndMonitorExecution(startWorkflowRequest *model.StartWorkflowRequest) (executionChannel executor3.WorkflowExecutionChannel, err error) {
	workflowId, err := workflow.StartWorkflow(startWorkflowRequest)
	if err != nil {
		return nil, err
	}
	return workflow.executor.MonitorExecution(workflowId)
}

func getInputAsMap(input interface{}) map[string]interface{} {
	if input == nil {
		return nil
	}
	casted, ok := input.(map[string]interface{})
	if ok {
		return casted
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
