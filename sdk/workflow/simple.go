//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package workflow

import "github.com/conductor-sdk/conductor-go/sdk/model"

type SimpleTask struct {
	Task
	workflowTask model.WorkflowTask
}

func NewSimpleTask(taskType string, taskRefName string) *SimpleTask {
	return &SimpleTask{
		Task: Task{
			name:              taskType,
			taskReferenceName: taskRefName,
			taskType:          SIMPLE,
			inputParameters:   map[string]interface{}{},
		},
		workflowTask: model.WorkflowTask{
			Name:              taskType,
			TaskReferenceName: taskRefName,
			Type_:             string(SIMPLE),
			TaskDefinition:    &model.TaskDef{Name: taskType},
		},
	}
}

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *SimpleTask) Input(key string, value interface{}) *SimpleTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *SimpleTask) InputMap(inputMap map[string]interface{}) *SimpleTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *SimpleTask) Optional(optional bool) *SimpleTask {
	task.Task.Optional(optional)
	return task
}

// Description of the task
func (task *SimpleTask) Description(description string) *SimpleTask {
	task.Task.Description(description)
	return task
}

func (task *SimpleTask) toWorkflowTask() []model.WorkflowTask {
	task.workflowTask.InputParameters = task.inputParameters
	task.workflowTask.Optional = task.optional
	task.workflowTask.Description = task.description
	task.workflowTask.CacheConfig = task.cacheConfig
	return []model.WorkflowTask{task.workflowTask}
}

// RetryPolicy for the task
func (task *SimpleTask) RetryPolicy(retryCount int32, policy RetryLogic, retryDelay int32, backoffScaleFactor int32) *SimpleTask {

	taskDefinition := task.ensureTaskDef()
	taskDefinition.RetryCount = retryCount
	taskDefinition.RetryLogic = string(policy)
	taskDefinition.RetryDelaySeconds = retryDelay
	taskDefinition.BackoffScaleFactor = backoffScaleFactor
	return task
}

// RateLimitFrequency based on the frequency window for the task
func (task *SimpleTask) RateLimitFrequency(rateLimitFrequencyInSeconds int32, rateLimitPerFrequency int32) *SimpleTask {
	taskDefinition := task.ensureTaskDef()
	taskDefinition.RateLimitPerFrequency = rateLimitPerFrequency
	taskDefinition.RateLimitFrequencyInSeconds = rateLimitFrequencyInSeconds
	return task
}

// ConcurrentExecutionLimit limits the max no. of concurrent execution of the tasks in the cluster
func (task *SimpleTask) ConcurrentExecutionLimit(limit int32) *SimpleTask {
	taskDefinition := task.ensureTaskDef()
	taskDefinition.ConcurrentExecLimit = limit
	return task
}

// ExecutionTimeout time in seconds by when the task MUST complete
// See #TimeoutPolicy
func (task *SimpleTask) ExecutionTimeout(timoutInSecond int64) *SimpleTask {
	taskDefinition := task.ensureTaskDef()
	taskDefinition.TimeoutSeconds = timoutInSecond
	return task
}

// PollTimeout time in seconds by when the task MUST be polled after getting scheduled
// See #TimeoutPolicy
func (task *SimpleTask) PollTimeout(timoutInSecond int32) *SimpleTask {
	taskDefinition := task.ensureTaskDef()
	taskDefinition.PollTimeoutSeconds = timoutInSecond
	return task
}

// ResponseTimeout time in seconds by which long-running task MUST send back the updates.
// See #TimeoutPolicy
func (task *SimpleTask) ResponseTimeout(timoutInSecond int64) *SimpleTask {
	taskDefinition := task.ensureTaskDef()
	taskDefinition.ResponseTimeoutSeconds = timoutInSecond
	return task
}

// TimeoutPolicy how to handle any of the timeout cases.
func (task *SimpleTask) TimeoutPolicy(timeoutPolicy TaskTimeoutPolicy) *SimpleTask {
	taskDefinition := task.ensureTaskDef()
	taskDefinition.TimeoutPolicy = string(timeoutPolicy)
	return task
}

// CacheConfig When set, the task's execution output is cached with the key and ttl as specified
// CacheKey can be parameterized.  e.g. if
func (task *Task) CacheConfig(cacheKey string, ttlInSeconds int) *Task {
	task.cacheConfig = &model.CacheConfig{
		Key:          cacheKey,
		TtlInSeconds: ttlInSeconds,
	}
	return task
}

func (task *SimpleTask) ensureTaskDef() *model.TaskDef {
	if task.workflowTask.TaskDefinition == nil {
		task.workflowTask.TaskDefinition = &model.TaskDef{Name: task.name}
	}
	return task.workflowTask.TaskDefinition
}
