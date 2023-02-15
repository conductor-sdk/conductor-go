//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package workflow

// import (
// 	"github.com/conductor-sdk/conductor-go/sdk/model"
// )

// const (
// 	forkedTasks       = "forkedTasks"
// 	forkedTasksInputs = "forkedTasksInputs"
// )

// func NewDynamicForkTask(taskRefName string, forkPrepareTask Task) *Task {
// 	// preForkTask: &forkPrepareTask,
// 	return &Task{
// 		model.WorkflowTask{
// 			Name:              taskRefName,
// 			TaskReferenceName: taskRefName,
// 			Description:       "",
// 			WorkflowTaskType:  string(FORK_JOIN_DYNAMIC),
// 			Optional:          false,
// 			InputParameters:   make(map[string]interface{}),
// 		},
// 	}
// }

// func NewDynamicForkTaskWithoutPrepareTask(taskRefName string) *Task {
// 	return &Task{
// 		model.WorkflowTask{
// 			Name:              taskRefName,
// 			TaskReferenceName: taskRefName,
// 			Description:       "",
// 			WorkflowTaskType:  string(FORK_JOIN_DYNAMIC),
// 			Optional:          false,
// 			InputParameters:   make(map[string]interface{}),
// 		},
// 	}
// }

// func NewDynamicForkWithJoinTask(taskRefName string, forkPrepareTask TaskInterface, join JoinTask) *Task {
// 	// preForkTask: &forkPrepareTask,
// 	// join:        join,
// 	return &model.WorkflowTask{
// 		Name:              taskRefName,
// 		TaskReferenceName: taskRefName,
// 		Description:       "",
// 		WorkflowTaskType:  string(FORK_JOIN_DYNAMIC),
// 		Optional:          false,
// 		InputParameters:   make(map[string]interface{}),
// 	},
// }

// func (task *DynamicForkTask) toWorkflowTask() []model.WorkflowTask {
// 	forkWorkflowTask := task.Task.toWorkflowTask()[0]
// 	forkWorkflowTask.DynamicForkTasksParam = forkedTasks
// 	forkWorkflowTask.DynamicForkTasksInputParamName = forkedTasksInputs
// 	if task.preForkTask != nil {
// 		forkWorkflowTask.InputParameters[forkedTasks] = (*task.preForkTask).OutputRef(forkedTasks)
// 		forkWorkflowTask.InputParameters[forkedTasksInputs] = (*task.preForkTask).OutputRef(forkedTasksInputs)
// 		tasks := (*task.preForkTask).toWorkflowTask()
// 		tasks = append(tasks, forkWorkflowTask, task.getJoinTask())
// 		return tasks
// 	}
// 	return []model.WorkflowTask{
// 		forkWorkflowTask,
// 		task.getJoinTask(),
// 	}
// }
