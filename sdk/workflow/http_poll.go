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

type HttpPollTask struct {
	SimpleTask
}

// NewHttpPollTask Create a new HTTP Poll Task
func NewHttpPollTask(taskRefName string, input *HttpPollInput) *HttpPollTask {
	if len(input.Method) == 0 {
		input.Method = GET
	}
	httpPoll := &HttpPollTask{
		SimpleTask{
			Task: Task{
				name:              string(HTTP_POLL),
				taskReferenceName: taskRefName,
				taskType:          HTTP_POLL,
				inputParameters:   map[string]interface{}{},
			},
			workflowTask: model.WorkflowTask{
				Name:              string(HTTP_POLL),
				TaskReferenceName: taskRefName,
				Type_:             string(HTTP_POLL),
				InputParameters:   map[string]interface{}{},
			},
		},
	}

	// Convert input to a map for the http_request parameter
	httpRequestMap := map[string]interface{}{
		"uri":    input.Uri,
		"method": string(input.Method),
	}

	// Add optional parameters only if they're set
	if len(input.Headers) > 0 {
		httpRequestMap["headers"] = input.Headers
	}
	if input.Accept != "" {
		httpRequestMap["accept"] = input.Accept
	}
	if input.ContentType != "" {
		httpRequestMap["contentType"] = input.ContentType
	}
	if input.ConnectionTimeOut > 0 {
		httpRequestMap["connectionTimeOut"] = input.ConnectionTimeOut
	}
	if input.ReadTimeout > 0 {
		httpRequestMap["readTimeout"] = input.ReadTimeout
	}
	if input.Body != nil {
		httpRequestMap["body"] = input.Body
	}
	if input.PollingInterval > 0 {
		httpRequestMap["pollingInterval"] = input.PollingInterval
	}
	if input.PollingTimeout > 0 {
		httpRequestMap["pollingTimeout"] = input.PollingTimeout
	}
	if input.PollingStrategy != "" {
		httpRequestMap["pollingStrategy"] = input.PollingStrategy
	}
	if input.SuccessCriteria != "" {
		httpRequestMap["successCondition"] = input.SuccessCriteria
	}
	if input.TerminationCriteria != "" {
		httpRequestMap["terminationCondition"] = input.TerminationCriteria
	}
	if input.Encode {
		httpRequestMap["encode"] = input.Encode
	}

	// Set the input parameters
	inputParams := map[string]interface{}{
		"http_request": httpRequestMap,
	}

	// Set both inputParameters and workflowTask.InputParameters
	httpPoll.inputParameters = inputParams
	httpPoll.workflowTask.InputParameters = inputParams

	return httpPoll
}

// HttpPollInput Input to the HTTP Poll task
type HttpPollInput struct {
	Method              HttpMethod          `json:"method"`
	Uri                 string              `json:"uri"`
	Headers             map[string][]string `json:"headers,omitempty"`
	Accept              string              `json:"accept,omitempty"`
	ContentType         string              `json:"contentType,omitempty"`
	ConnectionTimeOut   int16               `json:"connectionTimeOut,omitempty"`
	ReadTimeout         int16               `json:"readTimeout,omitempty"`
	Body                interface{}         `json:"body,omitempty"`
	PollingInterval     int                 `json:"pollingInterval,omitempty"`      // In seconds
	PollingTimeout      int                 `json:"pollingTimeout,omitempty"`       // In seconds
	PollingStrategy     string              `json:"pollingStrategy,omitempty"`      // "FIXED" or other strategy
	SuccessCriteria     string              `json:"successCondition,omitempty"`     // JavaScript expression
	TerminationCriteria string              `json:"terminationCondition,omitempty"` // JavaScript expression
	Encode              bool                `json:"encode,omitempty"`               // Whether to URL encode the request
}

// Input to the task. See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *HttpPollTask) Input(key string, value interface{}) *HttpPollTask {
	task.Task.Input(key, value)
	return task
}

// InputMap to the task. See https://conductor.netflix.com/how-tos/Tasks/task-inputs.html for details
func (task *HttpPollTask) InputMap(inputMap map[string]interface{}) *HttpPollTask {
	for k, v := range inputMap {
		task.inputParameters[k] = v
	}
	return task
}

// Optional if set to true, the task will not fail the workflow if the task fails
func (task *HttpPollTask) Optional(optional bool) *HttpPollTask {
	task.Task.Optional(optional)
	return task
}

// Description of the task
func (task *HttpPollTask) Description(description string) *HttpPollTask {
	task.Task.Description(description)
	return task
}
