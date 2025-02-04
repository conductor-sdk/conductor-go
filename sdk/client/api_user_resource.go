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
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
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
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue map[string]interface{}
	)

	path := "/users/{userId}/checkPermissions"
	path = strings.Replace(path, "{"+"userId"+"}", fmt.Sprintf("%v", userId), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	queryParams.Add("type", parameterToString(type_, ""))
	queryParams.Add("id", parameterToString(id, ""))
	headerParams["Accept"] = "application/json"
	r, err := a.prepareRequest(ctx, path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return returnValue, nil, err
	}

	httpResponse, err := a.callAPI(r)
	if err != nil || httpResponse == nil {
		return returnValue, httpResponse, err
	}

	responseBody, err := getDecompressedBody(httpResponse)
	httpResponse.Body.Close()
	if err != nil {
		return returnValue, httpResponse, err
	}

	if isSuccessfulStatus(httpResponse.StatusCode) {
		err = a.decode(&returnValue, responseBody, httpResponse.Header.Get("Content-Type"))
	} else {
		newErr := NewGenericSwaggerError(responseBody, httpResponse.Status, nil, httpResponse.StatusCode)
		return returnValue, httpResponse, newErr
	}

	return returnValue, httpResponse, err
}

/*
UserResourceApiService Delete a user
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id
    @return Response
*/
func (a *UserResourceApiService) DeleteUser(ctx context.Context, id string) (*http.Response, error) {
	var (
		httpMethod = strings.ToUpper("Delete")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)

	path := "/users/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	headerParams := make(map[string]string)
	headerParams["Accept"] = "application/json"

	queryParams := url.Values{}
	formParams := url.Values{}
	r, err := a.prepareRequest(ctx, path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return nil, err
	}

	httpResponse, err := a.callAPI(r)
	if err != nil || httpResponse == nil {
		return httpResponse, err
	}

	responseBody, err := getDecompressedBody(httpResponse)
	httpResponse.Body.Close()
	if err != nil {
		return httpResponse, err
	}

	if !isSuccessfulStatus(httpResponse.StatusCode) {
		return httpResponse, NewGenericSwaggerError(responseBody, httpResponse.Status, nil, httpResponse.StatusCode)
	}

	return httpResponse, nil
}

/*
UserResourceApiService Get the permissions this user has over workflows and tasks
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param userId
    @return interface{}
*/
func (a *UserResourceApiService) GetGrantedPermissions(ctx context.Context, userId string) (rbac.GrantedAccessResponse, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue rbac.GrantedAccessResponse
	)

	path := "/users/{userId}/permissions"
	path = strings.Replace(path, "{"+"userId"+"}", fmt.Sprintf("%v", userId), -1)

	headerParams := make(map[string]string)
	headerParams["Accept"] = "application/json"

	queryParams := url.Values{}
	formParams := url.Values{}
	r, err := a.prepareRequest(ctx, path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return returnValue, nil, err
	}

	httpResponse, err := a.callAPI(r)
	if err != nil || httpResponse == nil {
		return returnValue, httpResponse, err
	}

	responseBody, err := getDecompressedBody(httpResponse)
	httpResponse.Body.Close()
	if err != nil {
		return returnValue, httpResponse, err
	}

	if isSuccessfulStatus(httpResponse.StatusCode) {
		err = a.decode(&returnValue, responseBody, httpResponse.Header.Get("Content-Type"))
	} else {
		newErr := NewGenericSwaggerError(responseBody, httpResponse.Status, nil, httpResponse.StatusCode)
		return returnValue, httpResponse, newErr
	}

	return returnValue, httpResponse, err
}

/*
UserResourceApiService Get a user by id
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id
    @return interface{}
*/
func (a *UserResourceApiService) GetUser(ctx context.Context, id string) (*rbac.ConductorUser, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue rbac.ConductorUser
	)

	path := "/users/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	headerParams := make(map[string]string)
	headerParams["Accept"] = "application/json"

	queryParams := url.Values{}
	formParams := url.Values{}
	r, err := a.prepareRequest(ctx, path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return nil, nil, err
	}

	httpResponse, err := a.callAPI(r)
	if err != nil || httpResponse == nil {
		return nil, httpResponse, err
	}

	responseBody, err := getDecompressedBody(httpResponse)
	httpResponse.Body.Close()
	if err != nil {
		return nil, httpResponse, err
	}

	if isSuccessfulStatus(httpResponse.StatusCode) {
		err = a.decode(&returnValue, responseBody, httpResponse.Header.Get("Content-Type"))
	} else {
		newErr := NewGenericSwaggerError(responseBody, httpResponse.Status, nil, httpResponse.StatusCode)
		return nil, httpResponse, newErr
	}

	return &returnValue, httpResponse, err
}

/*
   UserResourceApiService Get all users
   * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    * @param optional nil or *UserResourceApiListUsersOpts - Optional Parameters:
        * @param "Apps" (optional.Bool) -
   @return []ConductorUser
*/

func (a *UserResourceApiService) ListUsers(ctx context.Context, optionals *UserResourceApiListUsersOpts) ([]rbac.ConductorUser, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue []rbac.ConductorUser
	)

	path := "/users"

	headerParams := make(map[string]string)
	headerParams["Accept"] = "application/json"

	queryParams := url.Values{}
	formParams := url.Values{}

	if optionals != nil && optionals.Apps.IsSet() {
		queryParams.Add("apps", parameterToString(optionals.Apps.Value(), ""))
	}

	r, err := a.prepareRequest(ctx, path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return returnValue, nil, err
	}

	httpResponse, err := a.callAPI(r)
	if err != nil || httpResponse == nil {
		return returnValue, httpResponse, err
	}

	responseBody, err := getDecompressedBody(httpResponse)
	httpResponse.Body.Close()
	if err != nil {
		return returnValue, httpResponse, err
	}

	if isSuccessfulStatus(httpResponse.StatusCode) {
		err = a.decode(&returnValue, responseBody, httpResponse.Header.Get("Content-Type"))
	} else {
		newErr := NewGenericSwaggerError(responseBody, httpResponse.Status, nil, httpResponse.StatusCode)
		return returnValue, httpResponse, newErr
	}

	return returnValue, httpResponse, err
}

/*
UserResourceApiService Create or update a user
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param id
    @return interface{}
*/
func (a *UserResourceApiService) UpsertUser(ctx context.Context, body rbac.UpsertUserRequest, id string) (*rbac.ConductorUser, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Put")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue rbac.ConductorUser
	)

	path := "/users/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	headerParams := make(map[string]string)
	headerParams["Accept"] = "application/json"
	headerParams["Content-Type"] = "application/json"

	queryParams := url.Values{}
	formParams := url.Values{}

	postBody = &body
	r, err := a.prepareRequest(ctx, path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return nil, nil, err
	}

	httpResponse, err := a.callAPI(r)
	if err != nil || httpResponse == nil {
		return nil, httpResponse, err
	}

	responseBody, err := getDecompressedBody(httpResponse)
	httpResponse.Body.Close()
	if err != nil {
		return nil, httpResponse, err
	}

	if isSuccessfulStatus(httpResponse.StatusCode) {
		err = a.decode(&returnValue, responseBody, httpResponse.Header.Get("Content-Type"))
	} else {
		newErr := NewGenericSwaggerError(responseBody, httpResponse.Status, nil, httpResponse.StatusCode)
		return nil, httpResponse, newErr
	}

	return &returnValue, httpResponse, err
}
