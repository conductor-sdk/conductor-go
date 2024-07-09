// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package human

type HumanTaskSearch struct {
	Assignees       []HumanTaskUser `json:"assignees,omitempty"`
	Claimants       []HumanTaskUser `json:"claimants,omitempty"`
	DefinitionNames []string        `json:"definitionNames,omitempty"`
	DisplayNames    []string        `json:"displayNames,omitempty"`
	FullTextQuery   string          `json:"fullTextQuery,omitempty"`
	SearchType      string          `json:"searchType,omitempty"`
	Size            int32           `json:"size,omitempty"`
	Start           int32           `json:"start,omitempty"`
	States          []string        `json:"states,omitempty"`
	TaskInputQuery  string          `json:"taskInputQuery,omitempty"`
	TaskOutputQuery string          `json:"taskOutputQuery,omitempty"`
	TaskRefNames    []string        `json:"taskRefNames,omitempty"`
	UpdateEndTime   int64           `json:"updateEndTime,omitempty"`
	UpdateStartTime int64           `json:"updateStartTime,omitempty"`
	WorkflowNames   []string        `json:"workflowNames,omitempty"`
}
