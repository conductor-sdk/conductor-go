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
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

//NewForkTask creates a new fork task that executes the given tasks in parallel
/**
 * execute task specified in the forkedTasks parameter in parallel.
 *
 * <p>forkedTask is a two-dimensional list that executes the outermost list in parallel and list
 * within that is executed sequentially.
 *
 * <p>e.g. [[task1, task2],[task3, task4],[task5]] are executed as:
 *
 * <pre>
 *                    ---------------
 *                    |     fork    |
 *                    ---------------
 *                    |       |     |
 *                    |       |     |
 *                  task1  task3  task5
 *                  task2  task4    |
 *                    |      |      |
 *                 ---------------------
 *                 |       join        |
 *                 ---------------------
 * </pre>
 *
 *
 */
func NewForkTask(taskRefName string, forkedTask ...[]model.WorkflowTask) *Task {
	return &Task{
		model.WorkflowTask{
			Name:              taskRefName,
			TaskReferenceName: taskRefName,
			Description:       "",
			WorkflowTaskType:  string(FORK_JOIN),
			Optional:          false,
			InputParameters:   map[string]interface{}{},
			ForkTasks:         forkedTask,
		},
	}
}
