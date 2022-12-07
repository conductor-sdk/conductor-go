//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package workflow

type InlineTask struct {
	Task
}

const (
	JavascriptEvaluator = "javascript"
	graalJS             = "graaljs"
)

// NewInlineTask An inline task with that executes the given javascript
// Legacy -- please use NewInlineGraalJSTask which provides better performance
func NewInlineTask(name string, script string) *InlineTask {
	return &InlineTask{
		Task{
			name:              name,
			taskReferenceName: name,
			taskType:          INLINE,
			inputParameters: map[string]interface{}{
				"evaluatorType": JavascriptEvaluator,
				"expression":    script,
			},
		},
	}
}

// NewInlineGraalJSTask An inline task with that executes the given javascript.  Uses GraalVM for faster execution
func NewInlineGraalJSTask(name string, script string) *InlineTask {
	return &InlineTask{
		Task{
			name:              name,
			taskReferenceName: name,
			taskType:          INLINE,
			inputParameters: map[string]interface{}{
				"evaluatorType": graalJS,
				"expression":    script,
			},
		},
	}
}

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *InlineTask) Input(key string, value interface{}) *InlineTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *InlineTask) InputMap(inputMap map[string]interface{}) *InlineTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *InlineTask) Optional(optional bool) *InlineTask {
	task.Task.Optional(optional)
	return task
}

// Description of the task
func (task *InlineTask) Description(description string) *InlineTask {
	task.Task.Description(description)
	return task
}
