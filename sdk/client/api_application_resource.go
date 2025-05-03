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

	"github.com/conductor-sdk/conductor-go/sdk/model"
	"github.com/conductor-sdk/conductor-go/sdk/model/rbac"
)

type ApplicationResourceApiService struct {
	*APIClient
}

/*
ApplicationResourceApiService
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param applicationId
  - @param role
    @return interface{}
*/
func (a *ApplicationResourceApiService) AddRoleToApplicationUser(ctx context.Context, applicationId string, role string) (interface{}, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Post")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue interface{}
	)

	path := "/applications/{applicationId}/roles/{role}"
	path = strings.Replace(path, "{"+"applicationId"+"}", fmt.Sprintf("%v", applicationId), -1)
	path = strings.Replace(path, "{"+"role"+"}", fmt.Sprintf("%v", role), -1)

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
ApplicationResourceApiService Create an access key for an application
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id
    @return interface{}
*/
func (a *ApplicationResourceApiService) CreateAccessKey(ctx context.Context, id string) (*rbac.ConductorApplication, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Post")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue rbac.ConductorApplication
	)

	path := "/applications/{id}/accessKeys"
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
ApplicationResourceApiService Create an application
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
    @return interface{}
*/
func (a *ApplicationResourceApiService) CreateApplication(ctx context.Context, body rbac.CreateOrUpdateApplicationRequest) (*rbac.ConductorApplication, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Post")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue rbac.ConductorApplication
	)

	path := "/applications"

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

/*
ApplicationResourceApiService Delete an access key
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param applicationId
  - @param keyId
    @return interface{}
*/
func (a *ApplicationResourceApiService) DeleteAccessKey(ctx context.Context, applicationId string, keyId string) (*http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Delete")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue interface{}
	)

	path := "/applications/{applicationId}/accessKeys/{keyId}"
	path = strings.Replace(path, "{"+"applicationId"+"}", fmt.Sprintf("%v", applicationId), -1)
	path = strings.Replace(path, "{"+"keyId"+"}", fmt.Sprintf("%v", keyId), -1)

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

	if isSuccessfulStatus(httpResponse.StatusCode) {
		err = a.decode(&returnValue, responseBody, httpResponse.Header.Get("Content-Type"))
	} else {
		return httpResponse, NewGenericSwaggerError(responseBody, httpResponse.Status, nil, httpResponse.StatusCode)
	}

	return httpResponse, err
}

/*
ApplicationResourceApiService Delete an application
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id
    @return interface{}
*/
func (a *ApplicationResourceApiService) DeleteApplication(ctx context.Context, id string) (interface{}, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Delete")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue interface{}
	)

	path := "/applications/{id}"
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
ApplicationResourceApiService Delete a tag for application
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param id
*/
func (a *ApplicationResourceApiService) DeleteTagForApplication(ctx context.Context, body []model.Tag, id string) (*http.Response, error) {
	var (
		httpMethod = strings.ToUpper("Delete")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)

	path := "/applications/{id}/tags"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

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
ApplicationResourceApiService Get application&#x27;s access keys
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id
    @return interface{}
*/
func (a *ApplicationResourceApiService) GetAccessKeys(ctx context.Context, id string) ([]rbac.AccessKeyResponse, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue []rbac.AccessKeyResponse
	)

	path := "/applications/{id}/accessKeys"
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
ApplicationResourceApiService Get application id by access key id
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param accessKeyId
    @return interface{}
*/
func (a *ApplicationResourceApiService) GetAppByAccessKeyId(ctx context.Context, accessKeyId string) (interface{}, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue interface{}
	)

	path := "/applications/key/{accessKeyId}"
	path = strings.Replace(path, "{"+"accessKeyId"+"}", fmt.Sprintf("%v", accessKeyId), -1)

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
ApplicationResourceApiService Get an application by id
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id
    @return interface{}
*/
func (a *ApplicationResourceApiService) GetApplication(ctx context.Context, id string) (*rbac.ConductorApplication, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue rbac.ConductorApplication
	)

	path := "/applications/{id}"
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
ApplicationResourceApiService Get tags by application
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id
    @return []model.Tag
*/
func (a *ApplicationResourceApiService) GetTagsForApplication(ctx context.Context, id string) ([]model.Tag, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue []model.Tag
	)

	path := "/applications/{id}/tags"
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
ApplicationResourceApiService Get all applications
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return []rbac.ConductorApplication
*/
func (a *ApplicationResourceApiService) ListApplications(ctx context.Context) ([]rbac.ConductorApplication, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue []rbac.ConductorApplication
	)

	path := "/applications"

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
ApplicationResourceApiService Put a tag to application
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param id
*/
func (a *ApplicationResourceApiService) PutTagForApplication(ctx context.Context, body []model.Tag, id string) (*http.Response, error) {
	var (
		httpMethod = strings.ToUpper("Put")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)

	path := "/applications/{id}/tags"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

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
ApplicationResourceApiService
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param applicationId
  - @param role
    @return interface{}
*/
func (a *ApplicationResourceApiService) RemoveRoleFromApplicationUser(ctx context.Context, applicationId string, role string) (interface{}, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Delete")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue interface{}
	)

	path := "/applications/{applicationId}/roles/{role}"
	path = strings.Replace(path, "{"+"applicationId"+"}", fmt.Sprintf("%v", applicationId), -1)
	path = strings.Replace(path, "{"+"role"+"}", fmt.Sprintf("%v", role), -1)

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
ApplicationResourceApiService Toggle the status of an access key
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param applicationId
  - @param keyId
    @return interface{}
*/
func (a *ApplicationResourceApiService) ToggleAccessKeyStatus(ctx context.Context, applicationId string, keyId string) (interface{}, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Post")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue interface{}
	)

	path := "/applications/{applicationId}/accessKeys/{keyId}/status"
	path = strings.Replace(path, "{"+"applicationId"+"}", fmt.Sprintf("%v", applicationId), -1)
	path = strings.Replace(path, "{"+"keyId"+"}", fmt.Sprintf("%v", keyId), -1)

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

	if !isSuccessfulStatus(httpResponse.StatusCode) {
		newErr := NewGenericSwaggerError(responseBody, httpResponse.Status, nil, httpResponse.StatusCode)
		return returnValue, httpResponse, newErr
	}

	return returnValue, httpResponse, nil
}

/*
ApplicationResourceApiService Update an application
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param id
    @return interface{}
*/
func (a *ApplicationResourceApiService) UpdateApplication(ctx context.Context, body rbac.CreateOrUpdateApplicationRequest, id string) (*rbac.ConductorApplication, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Put")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue rbac.ConductorApplication
	)

	path := "/applications/{id}"
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
