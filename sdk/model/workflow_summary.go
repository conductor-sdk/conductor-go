//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package model

type WorkflowSummary struct {
	WorkflowType                     string `json:"workflowType,omitempty"`
	Version                          int32  `json:"version,omitempty"`
	WorkflowId                       string `json:"workflowId,omitempty"`
	CorrelationId                    string `json:"correlationId,omitempty"`
	StartTime                        string `json:"startTime,omitempty"`
	UpdateTime                       string `json:"updateTime,omitempty"`
	EndTime                          string `json:"endTime,omitempty"`
	Status                           string `json:"status,omitempty"`
	Input                            string `json:"input,omitempty"`
	Output                           string `json:"output,omitempty"`
	ReasonForIncompletion            string `json:"reasonForIncompletion,omitempty"`
	ExecutionTime                    int64  `json:"executionTime,omitempty"`
	Event                            string `json:"event,omitempty"`
	FailedReferenceTaskNames         string `json:"failedReferenceTaskNames,omitempty"`
	ExternalInputPayloadStoragePath  string `json:"externalInputPayloadStoragePath,omitempty"`
	ExternalOutputPayloadStoragePath string `json:"externalOutputPayloadStoragePath,omitempty"`
	Priority                         int32  `json:"priority,omitempty"`
	OutputSize                       int64  `json:"outputSize,omitempty"`
	InputSize                        int64  `json:"inputSize,omitempty"`
}
