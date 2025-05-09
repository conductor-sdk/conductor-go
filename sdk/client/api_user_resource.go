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
	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
	"net/http"
	"net/url"
)

type UserResourceApiService struct {
	*APIClient
}

type UserResourceApiListUsersOpts struct {
	Apps optional.Bool
}

/*
UserResourceApiService Get the permissions this user has over workflows and tasks
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param userId
  - @param type_
  - @param id
    @return interface{}
*/
func (a *UserResourceApiService) CheckPermissions(ctx context.Context, userId string, type_ string, id string) (map[string]interface{}, *http.Response, error) {
	var result map[string]interface{}

	path := fmt.Sprintf("/users/%s/checkPermissions", userId)

	queryParams := url.Values{}
	queryParams.Add("type", parameterToString(type_, ""))
	queryParams.Add("id", parameterToString(id, ""))

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, err
}

/*
UserResourceApiService Delete a user
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id
    @return Response
*/
func (a *UserResourceApiService) DeleteUser(ctx context.Context, id string) (*http.Response, error) {
	path := fmt.Sprintf("/users/%s", id)

	resp, err := a.Delete(ctx, path, nil, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
UserResourceApiService Get the permissions this user has over workflows and tasks
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param userId
    @return interface{}
*/
func (a *UserResourceApiService) GetGrantedPermissions(ctx context.Context, userId string) (rbac.GrantedAccessResponse, *http.Response, error) {
	var result rbac.GrantedAccessResponse

	path := fmt.Sprintf("/users/%s/permissions", userId)

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return rbac.GrantedAccessResponse{}, resp, err
	}
	return result, resp, err
}

/*
UserResourceApiService Get a user by id
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id
    @return interface{}
*/
func (a *UserResourceApiService) GetUser(ctx context.Context, id string) (*rbac.ConductorUser, *http.Response, error) {
	var result rbac.ConductorUser

	path := fmt.Sprintf("/users/%s", id)

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, err
}

/*
   UserResourceApiService Get all users
   * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    * @param optional nil or *UserResourceApiListUsersOpts - Optional Parameters:
        * @param "Apps" (optional.Bool) -
   @return []ConductorUser
*/

func (a *UserResourceApiService) ListUsers(ctx context.Context, optionals *UserResourceApiListUsersOpts) ([]rbac.ConductorUser, *http.Response, error) {
	var result []rbac.ConductorUser

	path := "/users"

	queryParams := url.Values{}
	if optionals != nil && optionals.Apps.IsSet() {
		queryParams.Add("apps", parameterToString(optionals.Apps.Value(), ""))
	}

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, err
}

/*
UserResourceApiService Create or update a user
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param id
    @return interface{}
*/
func (a *UserResourceApiService) UpsertUser(ctx context.Context, body rbac.UpsertUserRequest, id string) (*rbac.ConductorUser, *http.Response, error) {
	var result rbac.ConductorUser

	path := fmt.Sprintf("/users/%s", id)
	resp, err := a.Put(ctx, path, body, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, err
}
