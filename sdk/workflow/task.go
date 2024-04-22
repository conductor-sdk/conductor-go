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

type RetryLogic string

const (
	FixedRetry             RetryLogic = "FIXED"
	LinearBackoffRetry     RetryLogic = "LINEAR_BACKOFF"
	ExponenialBackOffRetry RetryLogic = "EXPONENTIAL_BACKOFF"
)
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

type TaskInterface interface {
	toWorkflowTask() []model.WorkflowTask
	OutputRef(path string) string
	ToTaskDef() *model.TaskDef
}

type Task struct {
	name              string
	taskReferenceName string
	description       string
	taskType          TaskType
	optional          bool
	inputParameters   map[string]interface{}
}

func (task *Task) toWorkflowTask() []model.WorkflowTask {
	return []model.WorkflowTask{
		{
			Name:              task.name,
			TaskReferenceName: task.taskReferenceName,
			Description:       task.description,
			InputParameters:   task.inputParameters,
			Optional:          task.optional,
			Type_:             string(task.taskType),
		},
	}
}

func (task *Task) ToTaskDef() *model.TaskDef {
	return &model.TaskDef{
		Name:        task.name,
		Description: task.description,
	}
}

func (task *Task) ReferenceName() string {
	return task.taskReferenceName
}
func (task *Task) OutputRef(path string) string {
	if path == "" {
		return fmt.Sprintf("${%s.output}", task.taskReferenceName)
	}
	return fmt.Sprintf("${%s.output.%s}", task.taskReferenceName, path)
}

//Note: All the below method should be implemented by the
//Implementing interface given its a fluent interface
//If not, the return type is a Task which makes it impossible to use fluent interface
//For the tasks like Switch which has other methods too - quirks with Golang!

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *Task) Input(key string, value interface{}) *Task {
	task.inputParameters[key] = value
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *Task) InputMap(inputMap map[string]interface{}) *Task {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}

// Description of the task
func (task *Task) Description(description string) *Task {
	task.description = description
	return task
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *Task) Optional(optional bool) *Task {
	task.optional = optional
	return task
}
