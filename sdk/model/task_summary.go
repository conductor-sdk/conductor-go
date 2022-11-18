// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.
package model

type TaskSummary struct {
	CorrelationId                    string `json:"correlationId,omitempty"`
	EndTime                          string `json:"endTime,omitempty"`
	ExecutionTime                    int64  `json:"executionTime,omitempty"`
	ExternalInputPayloadStoragePath  string `json:"externalInputPayloadStoragePath,omitempty"`
	ExternalOutputPayloadStoragePath string `json:"externalOutputPayloadStoragePath,omitempty"`
	Input                            string `json:"input,omitempty"`
	Output                           string `json:"output,omitempty"`
	QueueWaitTime                    int64  `json:"queueWaitTime,omitempty"`
	ReasonForIncompletion            string `json:"reasonForIncompletion,omitempty"`
	ScheduledTime                    string `json:"scheduledTime,omitempty"`
	StartTime                        string `json:"startTime,omitempty"`
	Status                           string `json:"status,omitempty"`
	TaskDefName                      string `json:"taskDefName,omitempty"`
	TaskId                           string `json:"taskId,omitempty"`
	TaskType                         string `json:"taskType,omitempty"`
	UpdateTime                       string `json:"updateTime,omitempty"`
	WorkflowId                       string `json:"workflowId,omitempty"`
	WorkflowPriority                 int32  `json:"workflowPriority,omitempty"`
	WorkflowType                     string `json:"workflowType,omitempty"`
}
