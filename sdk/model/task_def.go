//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package model

type TaskDef struct {
	OwnerApp                    string                 `json:"ownerApp,omitempty"`
	CreateTime                  int64                  `json:"createTime,omitempty"`
	UpdateTime                  int64                  `json:"updateTime,omitempty"`
	CreatedBy                   string                 `json:"createdBy,omitempty"`
	UpdatedBy                   string                 `json:"updatedBy,omitempty"`
	Name                        string                 `json:"name"`
	Description                 string                 `json:"description,omitempty"`
	RetryCount                  int32                  `json:"retryCount"`
	TimeoutSeconds              int64                  `json:"timeoutSeconds"`
	InputKeys                   []string               `json:"inputKeys,omitempty"`
	OutputKeys                  []string               `json:"outputKeys,omitempty"`
	TimeoutPolicy               string                 `json:"timeoutPolicy,omitempty"`
	RetryLogic                  string                 `json:"retryLogic,omitempty"`
	RetryDelaySeconds           int32                  `json:"retryDelaySeconds,omitempty"`
	ResponseTimeoutSeconds      int64                  `json:"responseTimeoutSeconds,omitempty"`
	ConcurrentExecLimit         int32                  `json:"concurrentExecLimit,omitempty"`
	InputTemplate               map[string]interface{} `json:"inputTemplate,omitempty"`
	RateLimitPerFrequency       int32                  `json:"rateLimitPerFrequency,omitempty"`
	RateLimitFrequencyInSeconds int32                  `json:"rateLimitFrequencyInSeconds,omitempty"`
	IsolationGroupId            string                 `json:"isolationGroupId,omitempty"`
	ExecutionNameSpace          string                 `json:"executionNameSpace,omitempty"`
	OwnerEmail                  string                 `json:"ownerEmail,omitempty"`
	PollTimeoutSeconds          int32                  `json:"pollTimeoutSeconds,omitempty"`
	BackoffScaleFactor          int32                  `json:"backoffScaleFactor,omitempty"`
	Tags                        []TagObject            `json:"tags,omitempty"`
	OverwriteTags               bool                   `json:"overwriteTags"`
}
