// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.
package model

type WorkflowSummary struct {
	CorrelationId                    string `json:"correlationId,omitempty"`
	CreatedBy                        string `json:"createdBy,omitempty"`
	EndTime                          string `json:"endTime,omitempty"`
	Event                            string `json:"event,omitempty"`
	ExecutionTime                    int64  `json:"executionTime,omitempty"`
	ExternalInputPayloadStoragePath  string `json:"externalInputPayloadStoragePath,omitempty"`
	ExternalOutputPayloadStoragePath string `json:"externalOutputPayloadStoragePath,omitempty"`
	FailedReferenceTaskNames         string `json:"failedReferenceTaskNames,omitempty"`
	Input                            string `json:"input,omitempty"`
	InputSize                        int64  `json:"inputSize,omitempty"`
	Output                           string `json:"output,omitempty"`
	OutputSize                       int64  `json:"outputSize,omitempty"`
	Priority                         int32  `json:"priority,omitempty"`
	ReasonForIncompletion            string `json:"reasonForIncompletion,omitempty"`
	StartTime                        string `json:"startTime,omitempty"`
	Status                           string `json:"status,omitempty"`
	UpdateTime                       string `json:"updateTime,omitempty"`
	Version                          int32  `json:"version,omitempty"`
	WorkflowId                       string `json:"workflowId,omitempty"`
	WorkflowType                     string `json:"workflowType,omitempty"`
}
