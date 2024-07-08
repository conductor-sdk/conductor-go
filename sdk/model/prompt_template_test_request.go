// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package model

type PromptTemplateTestRequest struct {
	LlmProvider     string                 `json:"llmProvider,omitempty"`
	Model           string                 `json:"model,omitempty"`
	Prompt          string                 `json:"prompt,omitempty"`
	PromptVariables map[string]interface{} `json:"promptVariables,omitempty"`
	StopWords       []string               `json:"stopWords,omitempty"`
	Temperature     float64                `json:"temperature,omitempty"`
	TopP            float64                `json:"topP,omitempty"`
}
