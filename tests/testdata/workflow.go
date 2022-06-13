//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package examples

import (
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/definition"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
)

func IsWorkflowCompleted(workflow *model.Workflow) bool {
	return workflow.Status == model.CompletedWorkflow
}

func NewHttpTaskConductorWorkflow(workflowExecutor *executor.WorkflowExecutor) *definition.ConductorWorkflow {
	return newConductorWorkflow(
		workflowExecutor,
		"go_workflow_with_http_task",
		HttpTask,
	)
}

func NewSimpleTaskConductorWorkflow(workflowExecutor *executor.WorkflowExecutor) *definition.ConductorWorkflow {
	return newConductorWorkflow(
		workflowExecutor,
		"go_workflow_with_simple_task",
		SimpleTask,
	)
}

func newConductorWorkflow(workflowExecutor *executor.WorkflowExecutor, workflowName string, task definition.TaskInterface) *definition.ConductorWorkflow {
	return definition.NewConductorWorkflow(workflowExecutor).
		Name(workflowName).
		Version(1).
		Add(task)
}
