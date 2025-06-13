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
	CorrelationId                   string                         `json:"correlationId,omitempty"`
	CreatedBy                       string                         `json:"createdBy,omitempty"`
	ExternalInputPayloadStoragePath string                         `json:"externalInputPayloadStoragePath,omitempty"`
	IdempotencyKey                  string                         `json:"idempotencyKey,omitempty"`
	IdempotencyStrategy             string                         `json:"idempotencyStrategy,omitempty"`
	Input                           map[string]interface{}         `json:"input,omitempty"`
	Name                            string                         `json:"name"`
	Priority                        int32                          `json:"priority,omitempty"`
	SubWorkflowTestRequest          map[string]WorkflowTestRequest `json:"subWorkflowTestRequest,omitempty"`
	TaskRefToMockOutput             map[string][]TaskMock          `json:"taskRefToMockOutput,omitempty"`
	TaskToDomain                    map[string]string              `json:"taskToDomain,omitempty"`
	Version                         int32                          `json:"version,omitempty"`
	WorkflowDef                     *WorkflowDef                   `json:"workflowDef,omitempty"`
}
