//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package model

// Defines an event handler
type EventHandler struct {
	Name          string   `json:"name"`
	Event         string   `json:"event"`
	Condition     string   `json:"condition,omitempty"`
	Actions       []Action `json:"actions"`
	Active        bool     `json:"active,omitempty"`
	EvaluatorType string   `json:"evaluatorType,omitempty"`
}

type TerminateWorkflow struct {
	WorkflowId        string `json:"workflowId,omitempty"`
	TerminationReason string `json:"terminationReason,omitempty"`
}

type UpdateWorkflowVariables struct {
	WorkflowId  string                 `json:"workflowId,omitempty"`
	Variables   map[string]interface{} `json:"variables,omitempty"`
	AppendArray *bool                  `json:"appendArray,omitempty"`
}
