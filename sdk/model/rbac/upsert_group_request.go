// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package rbac

type UpsertGroupRequest struct {
	// a default Map<TargetType, Set<Access> to share permissions, allowed target types: WORKFLOW_DEF, TASK_DEF, WORKFLOW_SCHEDULE
	DefaultAccess map[string][]string `json:"defaultAccess,omitempty"`
	// A general description of the group
	Description string   `json:"description"`
	Roles       []string `json:"roles,omitempty"`
}
