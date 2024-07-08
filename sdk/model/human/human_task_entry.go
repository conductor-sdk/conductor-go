// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package human

type HumanTaskEntry struct {
	Assignee       *HumanTaskUser         `json:"assignee,omitempty"`
	Claimant       *HumanTaskUser         `json:"claimant,omitempty"`
	CreatedBy      string                 `json:"createdBy,omitempty"`
	CreatedOn      int64                  `json:"createdOn,omitempty"`
	DefinitionName string                 `json:"definitionName,omitempty"`
	DisplayName    string                 `json:"displayName,omitempty"`
	HumanTaskDef   *HumanTaskDefinition   `json:"humanTaskDef,omitempty"`
	Input          map[string]interface{} `json:"input,omitempty"`
	Output         map[string]interface{} `json:"output,omitempty"`
	State          string                 `json:"state,omitempty"`
	TaskId         string                 `json:"taskId,omitempty"`
	TaskRefName    string                 `json:"taskRefName,omitempty"`
	UpdatedBy      string                 `json:"updatedBy,omitempty"`
	UpdatedOn      int64                  `json:"updatedOn,omitempty"`
	WorkflowId     string                 `json:"workflowId,omitempty"`
	WorkflowName   string                 `json:"workflowName,omitempty"`
}
