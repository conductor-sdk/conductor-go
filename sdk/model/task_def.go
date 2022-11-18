// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.
package model

type TaskDef struct {
	BackoffScaleFactor          int32                  `json:"backoffScaleFactor,omitempty"`
	ConcurrentExecLimit         int32                  `json:"concurrentExecLimit,omitempty"`
	CreateTime                  int64                  `json:"createTime,omitempty"`
	CreatedBy                   string                 `json:"createdBy,omitempty"`
	Description                 string                 `json:"description,omitempty"`
	ExecutionNameSpace          string                 `json:"executionNameSpace,omitempty"`
	InputKeys                   []string               `json:"inputKeys,omitempty"`
	InputTemplate               map[string]interface{} `json:"inputTemplate,omitempty"`
	IsolationGroupId            string                 `json:"isolationGroupId,omitempty"`
	Name                        string                 `json:"name"`
	OutputKeys                  []string               `json:"outputKeys,omitempty"`
	OwnerApp                    string                 `json:"ownerApp,omitempty"`
	OwnerEmail                  string                 `json:"ownerEmail,omitempty"`
	PollTimeoutSeconds          int32                  `json:"pollTimeoutSeconds,omitempty"`
	RateLimitFrequencyInSeconds int32                  `json:"rateLimitFrequencyInSeconds,omitempty"`
	RateLimitPerFrequency       int32                  `json:"rateLimitPerFrequency,omitempty"`
	ResponseTimeoutSeconds      int64                  `json:"responseTimeoutSeconds,omitempty"`
	RetryCount                  int32                  `json:"retryCount,omitempty"`
	RetryDelaySeconds           int32                  `json:"retryDelaySeconds,omitempty"`
	RetryLogic                  string                 `json:"retryLogic,omitempty"`
	TimeoutPolicy               string                 `json:"timeoutPolicy,omitempty"`
	TimeoutSeconds              int64                  `json:"timeoutSeconds"`
	UpdateTime                  int64                  `json:"updateTime,omitempty"`
	UpdatedBy                   string                 `json:"updatedBy,omitempty"`
}
