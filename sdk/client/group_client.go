// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package client

import (
	"context"
	"net/http"

	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
)

type GroupClient interface {
	AddUserToGroup(ctx context.Context, groupId string, userId string) (interface{}, *http.Response, error)
	AddUsersToGroup(ctx context.Context, body []string, groupId string) (*http.Response, error)
	DeleteGroup(ctx context.Context, id string) (*http.Response, error)
	GetGrantedPermissions1(ctx context.Context, groupId string) (rbac.GrantedAccessResponse, *http.Response, error)
	GetGroup(ctx context.Context, id string) (interface{}, *http.Response, error)
	GetUsersInGroup(ctx context.Context, id string) (interface{}, *http.Response, error)
	ListGroups(ctx context.Context) ([]rbac.Group, *http.Response, error)
	RemoveUserFromGroup(ctx context.Context, groupId string, userId string) (interface{}, *http.Response, error)
	RemoveUsersFromGroup(ctx context.Context, body []string, groupId string) (*http.Response, error)
	UpsertGroup(ctx context.Context, body rbac.UpsertGroupRequest, id string) (interface{}, *http.Response, error)
}

func NewGroupClient(client *APIClient) GroupClient {
	return &GroupResourceApiService{client}
}
