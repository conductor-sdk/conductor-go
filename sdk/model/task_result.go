// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.
package model

type TaskResult struct {
	CallbackAfterSeconds             int64                  `json:"callbackAfterSeconds,omitempty"`
	ExternalOutputPayloadStoragePath string                 `json:"externalOutputPayloadStoragePath,omitempty"`
	Logs                             []TaskExecLog          `json:"logs,omitempty"`
	OutputData                       map[string]interface{} `json:"outputData,omitempty"`
	ReasonForIncompletion            string                 `json:"reasonForIncompletion,omitempty"`
	Status                           string                 `json:"status,omitempty"`
	SubWorkflowId                    string                 `json:"subWorkflowId,omitempty"`
	TaskId                           string                 `json:"taskId"`
	WorkerId                         string                 `json:"workerId,omitempty"`
	WorkflowInstanceId               string                 `json:"workflowInstanceId"`
}
