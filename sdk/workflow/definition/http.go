//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package definition

type HttpTask struct {
	Task
}

type HttpMethod string

const (
	GET     HttpMethod = "GET"
	PUT     HttpMethod = "PUT"
	POST    HttpMethod = "POST"
	DELETE  HttpMethod = "DELETE"
	HEAD    HttpMethod = "HEAD"
	OPTIONS HttpMethod = "OPTIONS"
)

// NewHttpTask Create a new HTTP Task
func NewHttpTask(taskRefName string, input *HttpInput) *HttpTask {
	if len(input.Method) == 0 {
		input.Method = GET
	}
	return &HttpTask{
		Task{
			name:              taskRefName,
			taskReferenceName: taskRefName,
			description:       "",
			taskType:          HTTP,
			optional:          false,
			inputParameters: map[string]interface{}{
				"http_request": input,
			},
		},
	}
}

// HttpInput Input to the HTTP task
type HttpInput struct {
	Method            HttpMethod          `json:"method"`
	Uri               string              `json:"uri"`
	Headers           map[string][]string `json:"headers,omitempty"`
	Accept            string              `json:"accept,omitempty"`
	ContentType       string              `json:"contentType,omitempty"`
	ConnectionTimeOut int16               `json:"ConnectionTimeOut,omitempty"`
	ReadTimeout       int16               `json:"readTimeout,omitempty"`
	Body              interface{}         `json:"body,omitempty"`
}

// Input to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *HttpTask) Input(key string, value interface{}) *HttpTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task.  See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *HttpTask) InputMap(inputMap map[string]interface{}) *HttpTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *HttpTask) Optional(optional bool) *HttpTask {
	task.Task.Optional(optional)
	return task
}

// Description of the task
func (task *HttpTask) Description(description string) *HttpTask {
	task.Task.Description(description)
	return task
}
