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

//DynamicForkInput struct that represents the output of the dynamic fork preparatory task
//
//DynamicFork requires inputs that specifies the list of tasks to be executed in parallel along with the inputs
//to each of these tasks.
//This usually means having another task before the dynamic task whose output contains
//the tasks to be forked.
//
//This struct represents the output of such a task
type DynamicForkInput struct {
	Tasks     []model.WorkflowTask
	TaskInput map[string]interface{}
}
