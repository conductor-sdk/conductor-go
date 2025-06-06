//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package model

type Workflow struct {
	OwnerApp                         string                 `json:"ownerApp,omitempty"`
	CreateTime                       int64                  `json:"createTime,omitempty"`
	UpdateTime                       int64                  `json:"updateTime,omitempty"`
	CreatedBy                        string                 `json:"createdBy,omitempty"`
	UpdatedBy                        string                 `json:"updatedBy,omitempty"`
	Status                           WorkflowStatus         `json:"status,omitempty"`
	EndTime                          int64                  `json:"endTime,omitempty"`
	WorkflowId                       string                 `json:"workflowId,omitempty"`
	ParentWorkflowId                 string                 `json:"parentWorkflowId,omitempty"`
	ParentWorkflowTaskId             string                 `json:"parentWorkflowTaskId,omitempty"`
	Tasks                            []Task                 `json:"tasks,omitempty"`
	Input                            map[string]interface{} `json:"input,omitempty"`
	Output                           map[string]interface{} `json:"output,omitempty"`
	CorrelationId                    string                 `json:"correlationId,omitempty"`
	ReRunFromWorkflowId              string                 `json:"reRunFromWorkflowId,omitempty"`
	ReasonForIncompletion            string                 `json:"reasonForIncompletion,omitempty"`
	Event                            string                 `json:"event,omitempty"`
	TaskToDomain                     map[string]string      `json:"taskToDomain,omitempty"`
	FailedReferenceTaskNames         []string               `json:"failedReferenceTaskNames,omitempty"`
	WorkflowDefinition               *WorkflowDef           `json:"workflowDefinition,omitempty"`
	ExternalInputPayloadStoragePath  string                 `json:"externalInputPayloadStoragePath,omitempty"`
	ExternalOutputPayloadStoragePath string                 `json:"externalOutputPayloadStoragePath,omitempty"`
	Priority                         int32                  `json:"priority,omitempty"`
	Variables                        map[string]interface{} `json:"variables,omitempty"`
	LastRetriedTime                  int64                  `json:"lastRetriedTime,omitempty"`
	// StartTime is deprecated
	StartTime int64 `json:"startTime,omitempty"`
	// WorkflowName is deprecated
	WorkflowName string `json:"workflowName,omitempty"`
	// WorkflowVersion is deprecated
	WorkflowVersion int32      `json:"workflowVersion,omitempty"`
	FailedTaskNames []string   `json:"failedTaskNames,omitempty"`
	History         []Workflow `json:"history,omitempty"`
	IdempotencyKey  string     `json:"idempotencyKey,omitempty"`
	RateLimitKey    string     `json:"rateLimitKey,omitempty"`
	RateLimited     bool       `json:"rateLimited,omitempty"`
}
