// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.
package model

type IntegrationDefFormField struct {
	DefaultValue string   `json:"defaultValue,omitempty"`
	Description  string   `json:"description,omitempty"`
	FieldName    string   `json:"fieldName,omitempty"`
	FieldType    string   `json:"fieldType,omitempty"`
	Label        string   `json:"label,omitempty"`
	Optional     bool     `json:"optional,omitempty"`
	Value        string   `json:"value,omitempty"`
	ValueOptions []Option `json:"valueOptions,omitempty"`
}

type Option struct {
	Label string `json:"label,omitempty"`
	Value string `json:"value,omitempty"`
}
