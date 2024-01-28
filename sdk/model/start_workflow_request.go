//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package model

type StartWorkflowRequest struct {
	Name                            string              `json:"name"`
	Version                         int32               `json:"version,omitempty"`
	CorrelationId                   string              `json:"correlationId,omitempty"`
	Input                           interface{}         `json:"input,omitempty"`
	TaskToDomain                    map[string]string   `json:"taskToDomain,omitempty"`
	WorkflowDef                     *WorkflowDef        `json:"workflowDef,omitempty"`
	ExternalInputPayloadStoragePath string              `json:"externalInputPayloadStoragePath,omitempty"`
	Priority                        int32               `json:"priority,omitempty"`
	IdempotencyKey                  string              `json:"idempotencyKey,omitempty"`
	IdempotencyStrategy             IdempotencyStrategy `json:"idempotencyStrategy,omitempty"`
}

func NewStartWorkflowRequest(name string, version int32, correlationId string, input interface{}) *StartWorkflowRequest {
	return &StartWorkflowRequest{
		Name:          name,
		Version:       version,
		CorrelationId: correlationId,
		Input:         input,
	}
}

func NewIdempotentStartWorkflowRequest(name string, version int32, correlationId string,
	idempotencyKey string, idempotencyStrategy IdempotencyStrategy, input interface{}) *StartWorkflowRequest {
	return &StartWorkflowRequest{
		Name:                name,
		Version:             version,
		CorrelationId:       correlationId,
		Input:               input,
		IdempotencyKey:      idempotencyKey,
		IdempotencyStrategy: idempotencyStrategy,
	}
}

func NewStartWorkflowRequestLegacy(name string, version int32, correlationId string, input interface{}) *StartWorkflowRequest {
	return &StartWorkflowRequest{
		Name:          name,
		Version:       version,
		CorrelationId: correlationId,
		Input:         input,
	}
}
