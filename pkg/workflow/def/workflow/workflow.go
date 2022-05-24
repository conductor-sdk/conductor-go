package workflow

import (
	"net/http"
	"time"

	"github.com/conductor-sdk/conductor-go/pkg/http_model"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/executor"
)

type TimeoutPolicy string

const (
	TimeOutWorkflow = "TIME_OUT_WF"
	AlertOnly       = "ALERT_ONLY"
)

type ConductorWorkflow struct {
	executor                      *executor.WorkflowExecutor
	name                          string
	version                       int32
	description                   string
	ownerEmail                    string
	tasks                         []TaskInterface
	timeoutPolicy                 TimeoutPolicy
	timeoutSeconds                int64
	failureWorkflow               string
	inputParameters               []string
	outputParameters              map[string]interface{}
	inputTemplate                 map[string]interface{}
	variables                     map[string]interface{}
	restartable                   bool
	workflowStatusListenerEnabled bool
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

func (workflow *ConductorWorkflow) TimeoutPolicy(timeoutPolicy TimeoutPolicy) *ConductorWorkflow {
	workflow.timeoutPolicy = timeoutPolicy
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

func (workflow *ConductorWorkflow) Register() (*http.Response, error) {
	return workflow.executor.RegisterWorkflow(
		workflow.toWorkflowDef(),
	)
}

func (workflow *ConductorWorkflow) Start(input interface{}) (executor.WorkflowExecutionChannel, error) {
	return workflow.executor.ExecuteWorkflow(
		workflow.name,
		workflow.version,
		input,
	)
}

func (workflow *ConductorWorkflow) StartWithTimeout(input interface{}, timeout time.Duration) (executor.WorkflowExecutionChannel, error) {
	return workflow.executor.ExecuteWorkflowWithTimeout(
		workflow.name,
		workflow.version,
		input,
		timeout,
	)
}

func (workflow *ConductorWorkflow) StartMany(amount int) ([]executor.WorkflowExecutionChannel, error) {
	workflowExecutionChannelList := make([]executor.WorkflowExecutionChannel, amount)
	for i := 0; i < amount; i += 1 {
		workflowExecutionChannel, err := workflow.Start(nil)
		if err != nil {
			return nil, err
		}
		workflowExecutionChannelList[i] = workflowExecutionChannel
	}
	return workflowExecutionChannelList, nil
}

func (workflow *ConductorWorkflow) StartManyWithTimeout(amount int, timeout time.Duration) ([]executor.WorkflowExecutionChannel, error) {
	workflowExecutionChannelList := make([]executor.WorkflowExecutionChannel, amount)
	for i := 0; i < amount; i += 1 {
		workflowExecutionChannel, err := workflow.StartWithTimeout(
			nil,
			timeout,
		)
		if err != nil {
			return nil, err
		}
		workflowExecutionChannelList[i] = workflowExecutionChannel
	}
	return workflowExecutionChannelList, nil
}

func (workflow *ConductorWorkflow) toWorkflowDef() *http_model.WorkflowDef {
	return &http_model.WorkflowDef{
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

func getWorkflowTasksFromConductorWorkflow(workflow *ConductorWorkflow) []http_model.WorkflowTask {
	workflowTasks := make([]http_model.WorkflowTask, 0)
	for _, task := range workflow.tasks {
		workflowTasks = append(
			workflowTasks,
			task.toWorkflowTask()...,
		)
	}
	return workflowTasks
}
