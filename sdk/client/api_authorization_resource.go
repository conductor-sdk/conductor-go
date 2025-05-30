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

type AuthorizationResourceApiService struct {
	*APIClient
}

/*
AuthorizationResourceApiService Get the access that have been granted over the given object
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param type_
  - @param id
    @return interface{}
*/
func (a *AuthorizationResourceApiService) GetPermissions(ctx context.Context, type_ string, id string) (interface{}, *http.Response, error) {
	var result interface{}
	path := fmt.Sprintf("/auth/authorization/%s/%s", type_, id)
	resp, err := a.Get(ctx, path, nil, &result)

	// Return nil result if there's an error to match original behavior
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

/*
AuthorizationResourceApiService Grant access to a user over the target
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
    @return Response
*/
func (a *AuthorizationResourceApiService) GrantPermissions(ctx context.Context, body rbac.AuthorizationRequest) (*http.Response, error) {
	path := "/auth/authorization"
	resp, err := a.Post(ctx, path, body, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
AuthorizationResourceApiService Remove user&#x27;s access over the target
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
    @return Response
*/
func (a *AuthorizationResourceApiService) RemovePermissions(ctx context.Context, body rbac.AuthorizationRequest) (*http.Response, error) {
	path := "/auth/authorization"
	resp, err := a.DeleteWithBody(ctx, path, body, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
