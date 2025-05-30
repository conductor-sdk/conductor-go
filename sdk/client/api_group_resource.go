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
	"fmt"
	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
	"net/http"
)

type GroupResourceApiService struct {
	*APIClient
}

/*
GroupResourceApiService Add user to group
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param groupId
  - @param userId
    @return interface{}
*/
func (a *GroupResourceApiService) AddUserToGroup(ctx context.Context, groupId string, userId string) (interface{}, *http.Response, error) {
	var result interface{}
	path := fmt.Sprintf("/groups/%s/users/%s", groupId, userId)
	resp, err := a.Post(ctx, path, nil, &result)

	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

/*
GroupResourceApiService Add users to group
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param groupId
*/
func (a *GroupResourceApiService) AddUsersToGroup(ctx context.Context, body []string, groupId string) (*http.Response, error) {
	path := fmt.Sprintf("/groups/%s/users", groupId)
	resp, err := a.Post(ctx, path, body, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

/*
GroupResourceApiService Delete a group
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id
    @return Response
*/
func (a *GroupResourceApiService) DeleteGroup(ctx context.Context, id string) (*http.Response, error) {
	path := fmt.Sprintf("/groups/%s", id)
	resp, err := a.Delete(ctx, path, nil, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
GroupResourceApiService Get the permissions this group has over workflows and tasks
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param groupId
    @return rbac.GrantedAccessResponse
*/
func (a *GroupResourceApiService) GetGrantedPermissions1(ctx context.Context, groupId string) (rbac.GrantedAccessResponse, *http.Response, error) {
	var result rbac.GrantedAccessResponse
	path := fmt.Sprintf("/groups/%s/permissions", groupId)
	resp, err := a.Get(ctx, path, nil, &result)

	if err != nil {
		return rbac.GrantedAccessResponse{}, resp, err
	}

	return result, resp, nil
}

/*
GroupResourceApiService Get a group by id
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id
    @return interface{}
*/
func (a *GroupResourceApiService) GetGroup(ctx context.Context, id string) (interface{}, *http.Response, error) {
	var result interface{}
	path := fmt.Sprintf("/groups/%s", id)
	resp, err := a.Get(ctx, path, nil, &result)

	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

/*
GroupResourceApiService Get all users in group
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id
    @return interface{}
*/
func (a *GroupResourceApiService) GetUsersInGroup(ctx context.Context, id string) (interface{}, *http.Response, error) {
	var result interface{}
	path := fmt.Sprintf("/groups/%s/users", id)
	resp, err := a.Get(ctx, path, nil, &result)

	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

/*
GroupResourceApiService Get all groups
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return []rbac.Group
*/
func (a *GroupResourceApiService) ListGroups(ctx context.Context) ([]rbac.Group, *http.Response, error) {
	var result []rbac.Group
	resp, err := a.Get(ctx, "/groups", nil, &result)

	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

/*
GroupResourceApiService Remove user from group
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param groupId
  - @param userId
    @return interface{}
*/
func (a *GroupResourceApiService) RemoveUserFromGroup(ctx context.Context, groupId string, userId string) (interface{}, *http.Response, error) {
	var result interface{}
	path := fmt.Sprintf("/groups/%s/users/%s", groupId, userId)
	resp, err := a.Delete(ctx, path, nil, &result)

	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

//TODO test this method
/*
GroupResourceApiService Remove users from group
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param groupId
*/
func (a *GroupResourceApiService) RemoveUsersFromGroup(ctx context.Context, body []string, groupId string) (*http.Response, error) {
	path := fmt.Sprintf("/groups/%s/users", groupId)

	resp, err := a.DeleteWithBody(ctx, path, body, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

/*
GroupResourceApiService Create or update a group
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param id
    @return interface{}
*/
func (a *GroupResourceApiService) UpsertGroup(ctx context.Context, body rbac.UpsertGroupRequest, id string) (interface{}, *http.Response, error) {
	var result interface{}
	path := fmt.Sprintf("/groups/%s", id)

	resp, err := a.Put(ctx, path, body, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
