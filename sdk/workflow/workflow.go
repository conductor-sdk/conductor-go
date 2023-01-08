//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package workflow

import (
	"encoding/json"

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
	log "github.com/sirupsen/logrus"
)

type TimeoutPolicy string

const (
	TimeOutWorkflow TimeoutPolicy = "TIME_OUT_WF"
	AlertOnly       TimeoutPolicy = "ALERT_ONLY"
)

type ConductorWorkflow struct {
	executor *executor.WorkflowExecutor
	// name             string
	// version          int32
	// description      string
	// ownerEmail       string
	// tasks            []TaskInterface
	// timeoutPolicy    TimeoutPolicy
	// timeoutSeconds   int64
	// failureWorkflow  string
	// inputParameters  []string
	// outputParameters map[string]interface{}
	// inputTemplate    map[string]interface{}
	// variables        map[string]interface{}
	// restartable      bool
	WorkflowDef model.WorkflowDef
}

func NewConductorWorkflow(executor *executor.WorkflowExecutor) *ConductorWorkflow {
	return &ConductorWorkflow{
		executor: executor,
		WorkflowDef: model.WorkflowDef{
			TimeoutPolicy: string(AlertOnly),
			Restartable:   true,
		},
	}
}

func (workflow *ConductorWorkflow) Name(name string) *ConductorWorkflow {
	workflow.WorkflowDef.Name = name
	return workflow
}

func (workflow *ConductorWorkflow) Version(version int32) *ConductorWorkflow {
	workflow.WorkflowDef.Version = version
	return workflow
}

func (workflow *ConductorWorkflow) Description(description string) *ConductorWorkflow {
	workflow.WorkflowDef.Description = description
	return workflow
}

func (workflow *ConductorWorkflow) TimeoutPolicy(timeoutPolicy TimeoutPolicy, timeoutSeconds int64) *ConductorWorkflow {
	workflow.WorkflowDef.TimeoutPolicy = string(timeoutPolicy)
	workflow.WorkflowDef.TimeoutSeconds = timeoutSeconds
	return workflow
}

func (workflow *ConductorWorkflow) TimeoutSeconds(timeoutSeconds int64) *ConductorWorkflow {
	workflow.WorkflowDef.TimeoutSeconds = timeoutSeconds
	return workflow
}

// FailureWorkflow name of the workflow to execute when this workflow fails.
// Failure workflows can be used for handling compensation logic
func (workflow *ConductorWorkflow) FailureWorkflow(failureWorkflow string) *ConductorWorkflow {
	workflow.WorkflowDef.FailureWorkflow = failureWorkflow
	return workflow
}

// Restartable if the workflow can be restarted after it has reached terminal state.
// Set this to false if restarting workflow can have side effects
func (workflow *ConductorWorkflow) Restartable(restartable bool) *ConductorWorkflow {
	workflow.WorkflowDef.Restartable = restartable
	return workflow
}

// OutputParameters Workflow outputs. Workflow output follows similar structure as task inputs
// See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for more details
func (workflow *ConductorWorkflow) OutputParameters(outputParameters interface{}) *ConductorWorkflow {
	workflow.WorkflowDef.OutputParameters = getInputAsMap(outputParameters)
	return workflow
}

// InputTemplate template input to the workflow.  Can have combination of variables (e.g. ${workflow.input.abc}) and
// static values
func (workflow *ConductorWorkflow) InputTemplate(inputTemplate interface{}) *ConductorWorkflow {
	workflow.WorkflowDef.InputTemplate = getInputAsMap(inputTemplate)
	return workflow
}

// Variables Workflow variables are set using SET_VARIABLE task.  Excellent way to maintain business state
// e.g. Variables can maintain business/user specific states which can be queried and inspected to find out the state of the workflow
func (workflow *ConductorWorkflow) Variables(variables interface{}) *ConductorWorkflow {
	workflow.WorkflowDef.Variables = getInputAsMap(variables)
	return workflow
}

// InputParameters List of the input parameters to the workflow.  Used ONLY for the documentation purpose.
func (workflow *ConductorWorkflow) InputParameters(inputParameters ...string) *ConductorWorkflow {
	workflow.WorkflowDef.InputParameters = inputParameters
	return workflow
}

func (workflow *ConductorWorkflow) OwnerEmail(ownerEmail string) *ConductorWorkflow {
	workflow.WorkflowDef.OwnerEmail = ownerEmail
	return workflow
}

func (workflow *ConductorWorkflow) GetName() (name string) {
	return workflow.WorkflowDef.Name
}

func (workflow *ConductorWorkflow) GetVersion() (version int32) {
	return workflow.WorkflowDef.Version
}

func (workflow *ConductorWorkflow) Add(task *Task) *ConductorWorkflow {
	workflow.WorkflowDef.Tasks = append(workflow.WorkflowDef.Tasks, task.WorkflowTask)
	return workflow
}

// Register the workflow definition with the server. If overwrite is set, the definition on the server will be overwritten.
// When not set, the call fails if there is any change in the workflow definition between the server and what is being registered.
func (workflow *ConductorWorkflow) Register(overwrite bool) error {
	return workflow.executor.RegisterWorkflow(overwrite, &workflow.WorkflowDef)
}

// StartWorkflowWithInput ExecuteWorkflowWithInput Execute the workflow with specific input.  The input struct MUST be serializable to JSON
// Returns the workflow Id that can be used to monitor and get the status of the workflow execution
func (workflow *ConductorWorkflow) StartWorkflowWithInput(input interface{}) (workflowId string, err error) {
	version := workflow.GetVersion()
	return workflow.executor.StartWorkflow(
		&model.StartWorkflowRequest{
			Name:        workflow.GetName(),
			Version:     &version,
			Input:       getInputAsMap(input),
			WorkflowDef: &workflow.WorkflowDef,
		},
	)
}

// StartWorkflow starts the workflow execution with startWorkflowRequest that allows you to specify more details like task domains, correlationId etc.
// Returns the ID of the newly created workflow
func (workflow *ConductorWorkflow) StartWorkflow(startWorkflowRequest *model.StartWorkflowRequest) (workflowId string, err error) {
	startWorkflowRequest.WorkflowDef = &workflow.WorkflowDef
	return workflow.executor.StartWorkflow(startWorkflowRequest)
}

// ExecuteWorkflowWithInput Execute the workflow with specific input and wait for the workflow to complete or until the task specified as waitUntil is completed.
// waitUntilTask Reference name of the task which MUST be completed before returning the output.  if specified as empty string, then the call waits until the
// workflow completes or reaches the timeout (as specified on the server)
// The input struct MUST be serializable to JSON
// Returns the workflow output
func (workflow *ConductorWorkflow) ExecuteWorkflowWithInput(input interface{}, waitUntilTask string) (worfklowRun *model.WorkflowRun, err error) {
	version := workflow.GetVersion()
	return workflow.executor.ExecuteWorkflow(
		&model.StartWorkflowRequest{
			Name:        workflow.GetName(),
			Version:     &version,
			Input:       getInputAsMap(input),
			WorkflowDef: &workflow.WorkflowDef,
		},
		waitUntilTask,
	)
}

// StartWorkflowsAndMonitorExecution Starts the workflow execution and returns a channel that can be used to monitor the workflow execution
// This method is useful for short duration workflows that are expected to complete in few seconds.  For long-running workflows use GetStatus APIs to periodically check the status
func (workflow *ConductorWorkflow) StartWorkflowsAndMonitorExecution(startWorkflowRequest *model.StartWorkflowRequest) (executionChannel executor.WorkflowExecutionChannel, err error) {
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
