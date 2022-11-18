// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.
package model

type WorkflowRun struct {
	CorrelationId string                 `json:"correlationId,omitempty"`
	CreateTime    int64                  `json:"createTime,omitempty"`
	CreatedBy     string                 `json:"createdBy,omitempty"`
	Input         map[string]interface{} `json:"input,omitempty"`
	Output        map[string]interface{} `json:"output,omitempty"`
	Priority      int32                  `json:"priority,omitempty"`
	RequestId     string                 `json:"requestId,omitempty"`
	Status        string                 `json:"status,omitempty"`
	Tasks         []Task                 `json:"tasks,omitempty"`
	UpdateTime    int64                  `json:"updateTime,omitempty"`
	Variables     map[string]interface{} `json:"variables,omitempty"`
	WorkflowId    string                 `json:"workflowId,omitempty"`
}
