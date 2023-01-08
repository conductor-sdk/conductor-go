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
	"fmt"

	"github.com/conductor-sdk/conductor-go/sdk/model"
)

type TaskType string

const (
	SIMPLE            TaskType = "SIMPLE"
	DYNAMIC           TaskType = "DYNAMIC"
	FORK_JOIN         TaskType = "FORK_JOIN"
	FORK_JOIN_DYNAMIC TaskType = "FORK_JOIN_DYNAMIC"
	SWITCH            TaskType = "SWITCH"
	JOIN              TaskType = "JOIN"
	DO_WHILE          TaskType = "DO_WHILE"
	SUB_WORKFLOW      TaskType = "SUB_WORKFLOW"
	START_WORKFLOW    TaskType = "START_WORKFLOW"
	EVENT             TaskType = "EVENT"
	WAIT              TaskType = "WAIT"
	HUMAN             TaskType = "HUMAN"
	HTTP              TaskType = "HTTP"
	INLINE            TaskType = "INLINE"
	TERMINATE         TaskType = "TERMINATE"
	KAFKA_PUBLISH     TaskType = "KAFKA_PUBLISH"
	JSON_JQ_TRANSFORM TaskType = "JSON_JQ_TRANSFORM"
	SET_VARIABLE      TaskType = "SET_VARIABLE"
)

type Task struct {
	WorkflowTask model.WorkflowTask
}

func (task *Task) ToTaskDef() *model.TaskDef {
	return &model.TaskDef{
		Name:        task.WorkflowTask.Name,
		Description: task.WorkflowTask.Description,
	}
}

func (task *Task) ReferenceName() string {
	return task.WorkflowTask.TaskReferenceName
}
func (task *Task) OutputRef(path string) string {
	if path == "" {
		return fmt.Sprintf("${%s.output}", task.WorkflowTask.TaskReferenceName)
	}
	return fmt.Sprintf("${%s.output.%s}", task.WorkflowTask.TaskReferenceName, path)
}

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *Task) Input(key string, value interface{}) *Task {
	task.WorkflowTask.InputParameters[key] = value
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *Task) InputMap(inputMap map[string]interface{}) *Task {
	for k, v := range inputMap {
		task.WorkflowTask.InputParameters[k] = v
	}
	return task
}

// Description of the task
func (task *Task) Description(description string) *Task {
	task.WorkflowTask.Description = description
	return task
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *Task) Optional(optional bool) *Task {
	task.WorkflowTask.Optional = optional
	return task
}
