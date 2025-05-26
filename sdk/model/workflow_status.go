//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package model

// WorkflowStatus represents the status information of a workflow
type WorkflowStatus struct {
	WorkflowId    string                 `json:"workflowId,omitempty"`
	CorrelationId string                 `json:"correlationId,omitempty"`
	Output        map[string]interface{} `json:"output,omitempty"`
	Variables     map[string]interface{} `json:"variables,omitempty"`
	Status        string                 `json:"status,omitempty"`
}

const (
	RunningWorkflow    string = "RUNNING"
	CompletedWorkflow  string = "COMPLETED"
	FailedWorkflow     string = "FAILED"
	TimedOutWorkflow   string = "TIMED_OUT"
	TerminatedWorkflow string = "TERMINATED"
	PausedWorkflow     string = "PAUSED"
)

var (
	WorkflowTerminalStates = []string{
		CompletedWorkflow,
		FailedWorkflow,
		TimedOutWorkflow,
		TerminatedWorkflow,
	}
)