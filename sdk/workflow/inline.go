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

const (
	JavascriptEvaluator = "javascript"
	graalJS             = "graaljs"
)

// NewInlineTask An inline task with that executes the given javascript
// Deprecated: please use NewInlineGraalJSTask which provides better performance
func NewInlineTask(name string, script string) *Task {
	return &Task{
		model.WorkflowTask{
			Name:              name,
			TaskReferenceName: name,
			WorkflowTaskType:  string(INLINE),
			InputParameters: map[string]interface{}{
				"evaluatorType": JavascriptEvaluator,
				"expression":    script,
			},
		},
	}
}

// NewInlineGraalJSTask An inline task with that executes the given javascript.  Uses GraalVM for faster execution
func NewInlineGraalJSTask(name string, script string) *Task {
	return &Task{
		model.WorkflowTask{
			Name:              name,
			TaskReferenceName: name,
			WorkflowTaskType:  string(INLINE),
			InputParameters: map[string]interface{}{
				"evaluatorType": graalJS,
				"expression":    script,
			},
		},
	}
}
