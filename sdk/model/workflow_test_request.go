// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package model

type WorkflowTestRequest struct {
	StartWorkflowRequest
	TaskRefToMockOutput    map[string][]TaskMock          `json:"taskRefToMockOutput,omitempty"`
	SubWorkflowTestRequest map[string]WorkflowTestRequest `json:"subWorkflowTestRequest,omitempty"`
}

type TaskMock struct {
	Status        string                 `json:"status,omitempty"`
	Output        map[string]interface{} `json:"output,omitempty"`
	ExecutionTime int64                  `json:"executionTime,omitempty"`
	QueueWaitTime int64                  `json:"queueWaitTime,omitempty"`
}