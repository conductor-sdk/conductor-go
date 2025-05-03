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

	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
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
	var (
		httpMethod  = strings.ToUpper("Post")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue interface{}
	)

	path := "/groups/{groupId}/users/{userId}"
	path = strings.Replace(path, "{"+"groupId"+"}", fmt.Sprintf("%v", groupId), -1)
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
GroupResourceApiService Add users to group
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param groupId
*/
func (a *GroupResourceApiService) AddUsersToGroup(ctx context.Context, body []string, groupId string) (*http.Response, error) {
	var (
		httpMethod = strings.ToUpper("Post")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)

	path := "/groups/{groupId}/users"
	path = strings.Replace(path, "{"+"groupId"+"}", fmt.Sprintf("%v", groupId), -1)

	headerParams := make(map[string]string)
	headerParams["Content-Type"] = "application/json"

	queryParams := url.Values{}
	formParams := url.Values{}

	postBody = &body
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
GroupResourceApiService Delete a group
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id
    @return Response
*/
func (a *GroupResourceApiService) DeleteGroup(ctx context.Context, id string) (*http.Response, error) {
	var (
		httpMethod = strings.ToUpper("Delete")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)

	path := "/groups/{id}"
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
GroupResourceApiService Get the permissions this group has over workflows and tasks
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param groupId
    @return rbac.GrantedAccessResponse
*/
func (a *GroupResourceApiService) GetGrantedPermissions1(ctx context.Context, groupId string) (rbac.GrantedAccessResponse, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue rbac.GrantedAccessResponse
	)

	path := "/groups/{groupId}/permissions"
	path = strings.Replace(path, "{"+"groupId"+"}", fmt.Sprintf("%v", groupId), -1)

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
GroupResourceApiService Get a group by id
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id
    @return interface{}
*/
func (a *GroupResourceApiService) GetGroup(ctx context.Context, id string) (interface{}, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue interface{}
	)

	path := "/groups/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

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
GroupResourceApiService Get all users in group
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id
    @return interface{}
*/
func (a *GroupResourceApiService) GetUsersInGroup(ctx context.Context, id string) (interface{}, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue interface{}
	)

	path := "/groups/{id}/users"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

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
GroupResourceApiService Get all groups
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return []rbac.Group
*/
func (a *GroupResourceApiService) ListGroups(ctx context.Context) ([]rbac.Group, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue []rbac.Group
	)

	path := "/groups"

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
GroupResourceApiService Remove user from group
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param groupId
  - @param userId
    @return interface{}
*/
func (a *GroupResourceApiService) RemoveUserFromGroup(ctx context.Context, groupId string, userId string) (interface{}, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Delete")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue interface{}
	)

	path := "/groups/{groupId}/users/{userId}"
	path = strings.Replace(path, "{"+"groupId"+"}", fmt.Sprintf("%v", groupId), -1)
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

//TODO test this method
/*
GroupResourceApiService Remove users from group
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param groupId
*/
func (a *GroupResourceApiService) RemoveUsersFromGroup(ctx context.Context, body []string, groupId string) (*http.Response, error) {
	var (
		httpMethod = strings.ToUpper("Delete")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)

	path := "/groups/{groupId}/users"
	path = strings.Replace(path, "{"+"groupId"+"}", fmt.Sprintf("%v", groupId), -1)

	headerParams := make(map[string]string)
	headerParams["Content-Type"] = "application/json"

	queryParams := url.Values{}
	formParams := url.Values{}

	postBody = &body
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
GroupResourceApiService Create or update a group
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param id
    @return interface{}
*/
func (a *GroupResourceApiService) UpsertGroup(ctx context.Context, body rbac.UpsertGroupRequest, id string) (interface{}, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Put")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue interface{}
	)

	path := "/groups/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	headerParams := make(map[string]string)
	headerParams["Accept"] = "application/json"
	headerParams["Content-Type"] = "application/json"

	queryParams := url.Values{}
	formParams := url.Values{}

	postBody = &body
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
