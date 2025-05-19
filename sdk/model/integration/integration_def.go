// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package integration

// IntegrationDef represents an integration definition
type IntegrationDef struct {
	Type           string                     `json:"type,omitempty"`
	Name           string                     `json:"name,omitempty"`
	Description    string                     `json:"description,omitempty"`
	Category       Category                   `json:"category,omitempty"`
	CategoryLabel  string                     `json:"categoryLabel,omitempty"`
	IconName       string                     `json:"iconName,omitempty"`
	Configuration  []IntegrationDefFormField  `json:"configuration,omitempty"`
	Tags           []string                   `json:"tags,omitempty"`
	Enabled        bool                       `json:"enabled,omitempty"`
}

const (
	TypeCustomAiModel = "TYPE_CUSTOM_AI_MODEL"
)