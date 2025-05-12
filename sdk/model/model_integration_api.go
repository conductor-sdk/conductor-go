// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.
package model

    
    type IntegrationApi struct {
            Api string `json:"api,omitempty"`
            Configuration map[string]interface{} `json:"configuration,omitempty"`
            CreateTime int64 `json:"createTime,omitempty"`
            CreatedBy string `json:"createdBy,omitempty"`
            Description string `json:"description,omitempty"`
            Enabled bool `json:"enabled,omitempty"`
            IntegrationName string `json:"integrationName,omitempty"`
            OwnerApp string `json:"ownerApp,omitempty"`
            Tags []Tag `json:"tags,omitempty"`
            UpdateTime int64 `json:"updateTime,omitempty"`
            UpdatedBy string `json:"updatedBy,omitempty"`
    }