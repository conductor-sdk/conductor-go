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
	"github.com/conductor-sdk/conductor-go/sdk/model/human"
	"net/http"
	"net/url"
	"strings"
)

type HumanTaskApiService struct {
	*APIClient
}

/*
   HumanTaskApiService Claim a task to an external user
   * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    * @param taskId
    * @param userId
    * @param optional nil or *HumanTaskApiAssignAndClaimOpts - Optional Parameters:
        * @param "OverrideAssignment" (optional.Bool) -
    * @param "WithTemplate" (optional.Bool) -
   @return HumanTaskEntry
*/

type HumanTaskApiAssignAndClaimOpts struct {
	OverrideAssignment optional.Bool
	WithTemplate       optional.Bool
}

func (a *HumanTaskApiService) AssignAndClaim(ctx context.Context, taskId string, userId string, optionals *HumanTaskApiAssignAndClaimOpts) (human.HumanTaskEntry, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Post")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue human.HumanTaskEntry
	)

	// create path and map variables
	path := "/human/tasks/{taskId}/externalUser/{userId}"
	path = strings.Replace(path, "{"+"taskId"+"}", fmt.Sprintf("%v", taskId), -1)
	path = strings.Replace(path, "{"+"userId"+"}", fmt.Sprintf("%v", userId), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	if optionals != nil && optionals.OverrideAssignment.IsSet() {
		queryParams.Add("overrideAssignment", parameterToString(optionals.OverrideAssignment.Value(), ""))
	}
	if optionals != nil && optionals.WithTemplate.IsSet() {
		queryParams.Add("withTemplate", parameterToString(optionals.WithTemplate.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypes := []string{}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
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

	if httpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.decode(&returnValue, responseBody, httpResponse.Header.Get("Content-Type"))
		if err == nil {
			return returnValue, httpResponse, err
		}
	}

	if httpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  responseBody,
			error: httpResponse.Status,
		}
		if httpResponse.StatusCode == 200 {
			var v human.HumanTaskEntry
			err = a.decode(&v, responseBody, httpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, httpResponse, newErr
			}
			newErr.model = v
			return returnValue, httpResponse, newErr
		}
		return returnValue, httpResponse, newErr
	}

	return returnValue, httpResponse, nil
}

/*
HumanTaskApiService API for backpopulating index data
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param var100
    @return map[string]interface{}
*/
func (a *HumanTaskApiService) BackPopulateFullTextIndex(ctx context.Context, var100 int32) (map[string]interface{}, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue map[string]interface{}
	)

	// create path and map variables
	path := "/human/tasks/backPopulateFullTextIndex"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	queryParams.Add("100", parameterToString(var100, ""))
	// to determine the Content-Type header
	contentTypes := []string{}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
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

	if httpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.decode(&returnValue, responseBody, httpResponse.Header.Get("Content-Type"))
		if err == nil {
			return returnValue, httpResponse, err
		}
	}

	if httpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  responseBody,
			error: httpResponse.Status,
		}
		if httpResponse.StatusCode == 200 {
			var v map[string]interface{}
			err = a.decode(&v, responseBody, httpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, httpResponse, newErr
			}
			newErr.model = v
			return returnValue, httpResponse, newErr
		}
		return returnValue, httpResponse, newErr
	}

	return returnValue, httpResponse, nil
}

/*
   HumanTaskApiService Claim a task by authenticated Conductor user
   * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    * @param taskId
    * @param optional nil or *HumanTaskApiClaimTaskOpts - Optional Parameters:
        * @param "OverrideAssignment" (optional.Bool) -
    * @param "WithTemplate" (optional.Bool) -
   @return human.HumanTaskEntry
*/

type HumanTaskApiClaimTaskOpts struct {
	OverrideAssignment optional.Bool
	WithTemplate       optional.Bool
}

func (a *HumanTaskApiService) ClaimTask(ctx context.Context, taskId string, optionals *HumanTaskApiClaimTaskOpts) (human.HumanTaskEntry, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Post")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue human.HumanTaskEntry
	)

	// create path and map variables
	path := "/human/tasks/{taskId}/claim"
	path = strings.Replace(path, "{"+"taskId"+"}", fmt.Sprintf("%v", taskId), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	if optionals != nil && optionals.OverrideAssignment.IsSet() {
		queryParams.Add("overrideAssignment", parameterToString(optionals.OverrideAssignment.Value(), ""))
	}
	if optionals != nil && optionals.WithTemplate.IsSet() {
		queryParams.Add("withTemplate", parameterToString(optionals.WithTemplate.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypes := []string{}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
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

	if httpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.decode(&returnValue, responseBody, httpResponse.Header.Get("Content-Type"))
		if err == nil {
			return returnValue, httpResponse, err
		}
	}

	if httpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  responseBody,
			error: httpResponse.Status,
		}
		if httpResponse.StatusCode == 200 {
			var v human.HumanTaskEntry
			err = a.decode(&v, responseBody, httpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, httpResponse, newErr
			}
			newErr.model = v
			return returnValue, httpResponse, newErr
		}
		return returnValue, httpResponse, newErr
	}

	return returnValue, httpResponse, nil
}

/*
HumanTaskApiService If the workflow is disconnected from tasks, this API can be used to clean up (in bulk)
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
*/
func (a *HumanTaskApiService) DeleteTaskFromHumanTaskRecords(ctx context.Context, body []string) (*http.Response, error) {
	var (
		httpMethod = strings.ToUpper("Delete")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)

	// create path and map variables
	path := "/human/tasks/delete"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
	}
	// body params
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

	if httpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  responseBody,
			error: httpResponse.Status,
		}
		return httpResponse, newErr
	}

	return httpResponse, nil
}

/*
HumanTaskApiService If the workflow is disconnected from tasks, this API can be used to clean up
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param taskId
*/
func (a *HumanTaskApiService) DeleteTaskFromHumanTaskRecords1(ctx context.Context, taskId string) (*http.Response, error) {
	var (
		httpMethod = strings.ToUpper("Delete")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)

	// create path and map variables
	path := "/human/tasks/delete/{taskId}"
	path = strings.Replace(path, "{"+"taskId"+"}", fmt.Sprintf("%v", taskId), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
	}
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

	if httpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  responseBody,
			error: httpResponse.Status,
		}
		return httpResponse, newErr
	}

	return httpResponse, nil
}

/*
HumanTaskApiService Delete all versions of user form template by name
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
*/
func (a *HumanTaskApiService) DeleteTemplateByName(ctx context.Context, name string) (*http.Response, error) {
	var (
		httpMethod = strings.ToUpper("Delete")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)

	// create path and map variables
	path := "/human/template/{name}"
	path = strings.Replace(path, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
	}
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

	if httpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  responseBody,
			error: httpResponse.Status,
		}
		return httpResponse, newErr
	}

	return httpResponse, nil
}

/*
HumanTaskApiService Delete a version of form template by name
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
  - @param version
*/
func (a *HumanTaskApiService) DeleteTemplatesByNameAndVersion(ctx context.Context, name string, version int32) (*http.Response, error) {
	var (
		httpMethod = strings.ToUpper("Delete")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)

	// create path and map variables
	path := "/human/template/{name}/{version}"
	path = strings.Replace(path, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	path = strings.Replace(path, "{"+"version"+"}", fmt.Sprintf("%v", version), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
	}
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

	if httpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  responseBody,
			error: httpResponse.Status,
		}
		return httpResponse, newErr
	}

	return httpResponse, nil
}

/*
   HumanTaskApiService List all user form templates or get templates by name, or a template by name and version
   * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    * @param optional nil or *HumanTaskApiGetAllTemplatesOpts - Optional Parameters:
        * @param "Name" (optional.String) -
    * @param "Version" (optional.Int32) -
   @return []human.human.HumanTaskSearch
*/

type HumanTaskApiGetAllTemplatesOpts struct {
	Name    optional.String
	Version optional.Int32
}

func (a *HumanTaskApiService) GetAllTemplates(ctx context.Context, optionals *HumanTaskApiGetAllTemplatesOpts) ([]human.HumanTaskSearch, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue []human.HumanTaskSearch
	)

	// create path and map variables
	path := "/human/template"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	if optionals != nil && optionals.Name.IsSet() {
		queryParams.Add("name", parameterToString(optionals.Name.Value(), ""))
	}
	if optionals != nil && optionals.Version.IsSet() {
		queryParams.Add("version", parameterToString(optionals.Version.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypes := []string{}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
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

	if httpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.decode(&returnValue, responseBody, httpResponse.Header.Get("Content-Type"))
		if err == nil {
			return returnValue, httpResponse, err
		}
	}

	if httpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  responseBody,
			error: httpResponse.Status,
		}
		if httpResponse.StatusCode == 200 {
			var v []human.HumanTaskSearch
			err = a.decode(&v, responseBody, httpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, httpResponse, newErr
			}
			newErr.model = v
			return returnValue, httpResponse, newErr
		}
		return returnValue, httpResponse, newErr
	}

	return returnValue, httpResponse, nil
}

/*
   HumanTaskApiService Get a task
   * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    * @param taskId
    * @param optional nil or *HumanTaskApiGetTask1Opts - Optional Parameters:
        * @param "WithTemplate" (optional.Bool) -
   @return human.HumanTaskEntry
*/

type HumanTaskApiGetTask1Opts struct {
	WithTemplate optional.Bool
}

func (a *HumanTaskApiService) GetTask1(ctx context.Context, taskId string, optionals *HumanTaskApiGetTask1Opts) (human.HumanTaskEntry, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue human.HumanTaskEntry
	)

	// create path and map variables
	path := "/human/tasks/{taskId}"
	path = strings.Replace(path, "{"+"taskId"+"}", fmt.Sprintf("%v", taskId), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	if optionals != nil && optionals.WithTemplate.IsSet() {
		queryParams.Add("withTemplate", parameterToString(optionals.WithTemplate.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypes := []string{}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
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

	if httpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.decode(&returnValue, responseBody, httpResponse.Header.Get("Content-Type"))
		if err == nil {
			return returnValue, httpResponse, err
		}
	}

	if httpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  responseBody,
			error: httpResponse.Status,
		}
		if httpResponse.StatusCode == 200 {
			var v human.HumanTaskEntry
			err = a.decode(&v, responseBody, httpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, httpResponse, newErr
			}
			newErr.model = v
			return returnValue, httpResponse, newErr
		}
		return returnValue, httpResponse, newErr
	}

	return returnValue, httpResponse, nil
}

/*
HumanTaskApiService Get list of task display names applicable for the user
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param searchType
    @return []string
*/
func (a *HumanTaskApiService) GetTaskDisplayNames(ctx context.Context, searchType string) ([]string, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue []string
	)

	// create path and map variables
	path := "/human/tasks/getTaskDisplayNames"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	queryParams.Add("searchType", parameterToString(searchType, ""))
	// to determine the Content-Type header
	contentTypes := []string{}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
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

	if httpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.decode(&returnValue, responseBody, httpResponse.Header.Get("Content-Type"))
		if err == nil {
			return returnValue, httpResponse, err
		}
	}

	if httpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  responseBody,
			error: httpResponse.Status,
		}
		if httpResponse.StatusCode == 200 {
			var v []string
			err = a.decode(&v, responseBody, httpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, httpResponse, newErr
			}
			newErr.model = v
			return returnValue, httpResponse, newErr
		}
		return returnValue, httpResponse, newErr
	}

	return returnValue, httpResponse, nil
}

/*
HumanTaskApiService Get user form template by name and version
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
  - @param version
    @return human.human.human.HumanTaskSearch
*/
func (a *HumanTaskApiService) GetTemplateByNameAndVersion(ctx context.Context, name string, version int32) (human.HumanTaskSearch, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue human.HumanTaskSearch
	)

	// create path and map variables
	path := "/human/template/{name}/{version}"
	path = strings.Replace(path, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	path = strings.Replace(path, "{"+"version"+"}", fmt.Sprintf("%v", version), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
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

	if httpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.decode(&returnValue, responseBody, httpResponse.Header.Get("Content-Type"))
		if err == nil {
			return returnValue, httpResponse, err
		}
	}

	if httpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  responseBody,
			error: httpResponse.Status,
		}
		if httpResponse.StatusCode == 200 {
			var v human.HumanTaskSearch
			err = a.decode(&v, responseBody, httpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, httpResponse, newErr
			}
			newErr.model = v
			return returnValue, httpResponse, newErr
		}
		return returnValue, httpResponse, newErr
	}

	return returnValue, httpResponse, nil
}

/*
HumanTaskApiService Get user form by human task id
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param humanTaskId
    @return human.HumanTaskSearch
*/
func (a *HumanTaskApiService) GetTemplateByTaskId(ctx context.Context, humanTaskId string) (human.HumanTaskSearch, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue human.HumanTaskSearch
	)

	// create path and map variables
	path := "/human/template/{humanTaskId}"
	path = strings.Replace(path, "{"+"humanTaskId"+"}", fmt.Sprintf("%v", humanTaskId), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
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

	if httpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.decode(&returnValue, responseBody, httpResponse.Header.Get("Content-Type"))
		if err == nil {
			return returnValue, httpResponse, err
		}
	}

	if httpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  responseBody,
			error: httpResponse.Status,
		}
		if httpResponse.StatusCode == 200 {
			var v human.HumanTaskSearch
			err = a.decode(&v, responseBody, httpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, httpResponse, newErr
			}
			newErr.model = v
			return returnValue, httpResponse, newErr
		}
		return returnValue, httpResponse, newErr
	}

	return returnValue, httpResponse, nil
}

/*
HumanTaskApiService Release a task without completing it
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param taskId
*/
func (a *HumanTaskApiService) ReassignTask(ctx context.Context, body []human.HumanTaskAssignment, taskId string) (*http.Response, error) {
	var (
		httpMethod = strings.ToUpper("Post")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)

	// create path and map variables
	path := "/human/tasks/{taskId}/reassign"
	path = strings.Replace(path, "{"+"taskId"+"}", fmt.Sprintf("%v", taskId), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
	}
	// body params
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

	if httpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  responseBody,
			error: httpResponse.Status,
		}
		return httpResponse, newErr
	}

	return httpResponse, nil
}

/*
HumanTaskApiService Release a task without completing it
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param taskId
*/
func (a *HumanTaskApiService) ReleaseTask(ctx context.Context, taskId string) (*http.Response, error) {
	var (
		httpMethod = strings.ToUpper("Post")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)

	// create path and map variables
	path := "/human/tasks/{taskId}/release"
	path = strings.Replace(path, "{"+"taskId"+"}", fmt.Sprintf("%v", taskId), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
	}
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

	if httpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  responseBody,
			error: httpResponse.Status,
		}
		return httpResponse, newErr
	}

	return httpResponse, nil
}

/*
   HumanTaskApiService Save user form template
   * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    * @param body
    * @param optional nil or *HumanTaskApiSaveTemplateOpts - Optional Parameters:
        * @param "NewVersion" (optional.Bool) -
   @return human.HumanTaskSearch
*/

type HumanTaskApiSaveTemplateOpts struct {
	NewVersion optional.Bool
}

func (a *HumanTaskApiService) SaveTemplate(ctx context.Context, body human.HumanTaskSearch, optionals *HumanTaskApiSaveTemplateOpts) (human.HumanTaskSearch, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Post")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue human.HumanTaskSearch
	)

	// create path and map variables
	path := "/human/template"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	if optionals != nil && optionals.NewVersion.IsSet() {
		queryParams.Add("newVersion", parameterToString(optionals.NewVersion.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
	}
	// body params
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

	if httpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.decode(&returnValue, responseBody, httpResponse.Header.Get("Content-Type"))
		if err == nil {
			return returnValue, httpResponse, err
		}
	}

	if httpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  responseBody,
			error: httpResponse.Status,
		}
		if httpResponse.StatusCode == 200 {
			var v human.HumanTaskSearch
			err = a.decode(&v, responseBody, httpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, httpResponse, newErr
			}
			newErr.model = v
			return returnValue, httpResponse, newErr
		}
		return returnValue, httpResponse, newErr
	}

	return returnValue, httpResponse, nil
}

/*
   HumanTaskApiService Save user form template
   * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    * @param body
    * @param optional nil or *HumanTaskApiSaveTemplatesOpts - Optional Parameters:
        * @param "NewVersion" (optional.Bool) -
   @return []human.HumanTaskSearch
*/

type HumanTaskApiSaveTemplatesOpts struct {
	NewVersion optional.Bool
}

func (a *HumanTaskApiService) SaveTemplates(ctx context.Context, body []human.HumanTaskSearch, optionals *HumanTaskApiSaveTemplatesOpts) ([]human.HumanTaskSearch, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Post")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue []human.HumanTaskSearch
	)

	// create path and map variables
	path := "/human/template/bulk"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	if optionals != nil && optionals.NewVersion.IsSet() {
		queryParams.Add("newVersion", parameterToString(optionals.NewVersion.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
	}
	// body params
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

	if httpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.decode(&returnValue, responseBody, httpResponse.Header.Get("Content-Type"))
		if err == nil {
			return returnValue, httpResponse, err
		}
	}

	if httpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  responseBody,
			error: httpResponse.Status,
		}
		if httpResponse.StatusCode == 200 {
			var v []human.HumanTaskSearch
			err = a.decode(&v, responseBody, httpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, httpResponse, newErr
			}
			newErr.model = v
			return returnValue, httpResponse, newErr
		}
		return returnValue, httpResponse, newErr
	}

	return returnValue, httpResponse, nil
}

/*
HumanTaskApiService Search human tasks
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
    @return human.HumanTaskSearchResult
*/
func (a *HumanTaskApiService) Search(ctx context.Context, body human.HumanTaskSearch) (human.HumanTaskSearchResult, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Post")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue human.HumanTaskSearchResult
	)

	// create path and map variables
	path := "/human/tasks/search"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{"application/json"}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
	}
	// body params
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

	if httpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.decode(&returnValue, responseBody, httpResponse.Header.Get("Content-Type"))
		if err == nil {
			return returnValue, httpResponse, err
		}
	}

	if httpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  responseBody,
			error: httpResponse.Status,
		}
		if httpResponse.StatusCode == 200 {
			var v human.HumanTaskSearchResult
			err = a.decode(&v, responseBody, httpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return returnValue, httpResponse, newErr
			}
			newErr.model = v
			return returnValue, httpResponse, newErr
		}
		return returnValue, httpResponse, newErr
	}

	return returnValue, httpResponse, nil
}

/*
   HumanTaskApiService If a task is assigned to a user, this API can be used to skip that assignment and move to the next assignee
   * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    * @param taskId
    * @param optional nil or *HumanTaskApiSkipTaskOpts - Optional Parameters:
        * @param "Reason" (optional.String) -

*/

type HumanTaskApiSkipTaskOpts struct {
	Reason optional.String
}

func (a *HumanTaskApiService) SkipTask(ctx context.Context, taskId string, optionals *HumanTaskApiSkipTaskOpts) (*http.Response, error) {
	var (
		httpMethod = strings.ToUpper("Post")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)

	// create path and map variables
	path := "/human/tasks/{taskId}/skip"
	path = strings.Replace(path, "{"+"taskId"+"}", fmt.Sprintf("%v", taskId), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	if optionals != nil && optionals.Reason.IsSet() {
		queryParams.Add("reason", parameterToString(optionals.Reason.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypes := []string{}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
	}
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

	if httpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  responseBody,
			error: httpResponse.Status,
		}
		return httpResponse, newErr
	}

	return httpResponse, nil
}

/*
   HumanTaskApiService Update task output, optionally complete
   * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    * @param body
    * @param taskId
    * @param optional nil or *HumanTaskApiUpdateTaskOutputOpts - Optional Parameters:
        * @param "Complete" (optional.Bool) -

*/

type HumanTaskApiUpdateTaskOutputOpts struct {
	Complete optional.Bool
}

func (a *HumanTaskApiService) UpdateTaskOutput(ctx context.Context, body map[string]interface{}, taskId string, optionals *HumanTaskApiUpdateTaskOutputOpts) (*http.Response, error) {
	var (
		httpMethod = strings.ToUpper("Post")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)

	// create path and map variables
	path := "/human/tasks/{taskId}/update"
	path = strings.Replace(path, "{"+"taskId"+"}", fmt.Sprintf("%v", taskId), -1)

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	if optionals != nil && optionals.Complete.IsSet() {
		queryParams.Add("complete", parameterToString(optionals.Complete.Value(), ""))
	}
	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
	}
	// body params
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

	if httpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  responseBody,
			error: httpResponse.Status,
		}
		return httpResponse, newErr
	}

	return httpResponse, nil
}

/*
   HumanTaskApiService Update task output, optionally complete
   * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    * @param body
    * @param workflowId
    * @param taskRefName
    * @param optional nil or *HumanTaskApiUpdateTaskOutputByRefOpts - Optional Parameters:
        * @param "Complete" (optional.Bool) -
    * @param "Iteration" (optional.Interface of []int32) -  Populate this value if your task is in a loop and you want to update a specific iteration. If its not in a loop OR if you want to just update the latest iteration, leave this as empty

*/

type HumanTaskApiUpdateTaskOutputByRefOpts struct {
	Complete  optional.Bool
	Iteration optional.Interface
}

func (a *HumanTaskApiService) UpdateTaskOutputByRef(ctx context.Context, body map[string]interface{}, workflowId string, taskRefName string, optionals *HumanTaskApiUpdateTaskOutputByRefOpts) (*http.Response, error) {
	var (
		httpMethod = strings.ToUpper("Post")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)

	// create path and map variables
	path := "/human/tasks/update/taskRef"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	queryParams.Add("workflowId", parameterToString(workflowId, ""))
	queryParams.Add("taskRefName", parameterToString(taskRefName, ""))
	if optionals != nil && optionals.Complete.IsSet() {
		queryParams.Add("complete", parameterToString(optionals.Complete.Value(), ""))
	}
	if optionals != nil && optionals.Iteration.IsSet() {
		queryParams.Add("iteration", parameterToString(optionals.Iteration.Value(), "multi"))
	}
	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// set Content-Type header
	contentType := selectHeaderContentType(contentTypes)
	if contentType != "" {
		headerParams["Content-Type"] = contentType
	}

	// to determine the Accept header
	headerAccepts := []string{}

	// set Accept header
	headerAccept := selectHeaderAccept(headerAccepts)
	if headerAccept != "" {
		headerParams["Accept"] = headerAccept
	}
	// body params
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

	if httpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  responseBody,
			error: httpResponse.Status,
		}
		return httpResponse, newErr
	}

	return httpResponse, nil
}
