// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package model

type WorkflowScheduleModel struct {
	CreateTime                  int64                 `json:"createTime,omitempty"`
	CreatedBy                   string                `json:"createdBy,omitempty"`
	CronExpression              string                `json:"cronExpression,omitempty"`
	Description                 string                `json:"description,omitempty"`
	Name                        string                `json:"name,omitempty"`
	Paused                      bool                  `json:"paused,omitempty"`
	PausedReason                string                `json:"pausedReason,omitempty"`
	RunCatchupScheduleInstances bool                  `json:"runCatchupScheduleInstances,omitempty"`
	ScheduleEndTime             int64                 `json:"scheduleEndTime,omitempty"`
	ScheduleStartTime           int64                 `json:"scheduleStartTime,omitempty"`
	StartWorkflowRequest        *StartWorkflowRequest `json:"startWorkflowRequest,omitempty"`
	Tags                        []Tag                 `json:"tags,omitempty"`
	UpdatedBy                   string                `json:"updatedBy,omitempty"`
	UpdatedTime                 int64                 `json:"updatedTime,omitempty"`
	ZoneId                      string                `json:"zoneId,omitempty"`
}
