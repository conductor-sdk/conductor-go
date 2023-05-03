//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package model

type WebhookConfig struct {
	CreatedBy                       string                    `json:"createdBy,omitempty"`
	HeaderKey                       string                    `json:"headerKey,omitempty"`
	Headers                         map[string]string         `json:"headers,omitempty"`
	Id                              string                    `json:"id,omitempty"`
	Name                            string                    `json:"name,omitempty"`
	ReceiverWorkflowNamesToVersions map[string]int32          `json:"receiverWorkflowNamesToVersions,omitempty"`
	SecretKey                       string                    `json:"secretKey,omitempty"`
	SecretValue                     string                    `json:"secretValue,omitempty"`
	SourcePlatform                  string                    `json:"sourcePlatform,omitempty"`
	UrlVerified                     bool                      `json:"urlVerified,omitempty"`
	Verifier                        string                    `json:"verifier,omitempty"`
	WebhookExecutionHistory         []WebhookExecutionHistory `json:"webhookExecutionHistory,omitempty"`
	WorkflowsToStart                map[string]int32          `json:"workflowsToStart,omitempty"`
}
