//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package testdata

import (
	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/workflow"
	"github.com/conductor-sdk/conductor-go/sdk/workflow/executor"
)

var HttpTask = workflow.NewHttpTask(
	"go_task_of_http_type", // task name
	&workflow.HttpInput{ // http input
		Uri: "https://catfact.ninja/fact",
	},
)

var SimpleTask = workflow.NewSimpleTask(
	"go_task_of_simple_type",
	"go_task_of_simple_type",
)

func IsWorkflowCompleted(workflow *model.Workflow) bool {
	return workflow.Status == model.CompletedWorkflow
}

func NewHttpTaskConductorWorkflow(workflowExecutor *executor.WorkflowExecutor) *workflow.ConductorWorkflow {
	return NewConductorWorkflow(
		workflowExecutor,
		"go_workflow_with_http_task",
		HttpTask,
	)
}

func NewSimpleTaskConductorWorkflow(workflowExecutor *executor.WorkflowExecutor) *workflow.ConductorWorkflow {
	
	return NewConductorWorkflow(
		workflowExecutor,
		"go_workflow_with_simple_task",
		SimpleTask,
	)
}

func NewConductorWorkflow(workflowExecutor *executor.WorkflowExecutor, workflowName string, task workflow.TaskInterface) *workflow.ConductorWorkflow {
	return workflow.NewConductorWorkflow(workflowExecutor).
		Name(workflowName).
		Version(1).
		Add(task)
}
