// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package rbac

import "github.com/conductor-sdk/conductor-go/sdk/model"

type ConductorApplication struct {
	CreateTime int64       `json:"createTime,omitempty"`
	CreatedBy  string      `json:"createdBy,omitempty"`
	Id         string      `json:"id,omitempty"`
	Name       string      `json:"name,omitempty"`
	Tags       []model.Tag `json:"tags,omitempty"`
	UpdateTime int64       `json:"updateTime,omitempty"`
	UpdatedBy  string      `json:"updatedBy,omitempty"`
}
