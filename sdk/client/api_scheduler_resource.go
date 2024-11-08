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
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

type SchedulerResourceApiService struct {
	*APIClient
}

/*
SchedulerResourceApiService Deletes an existing workflow schedule by name
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
    @return interface{}
*/
func (a *SchedulerResourceApiService) DeleteSchedule(ctx context.Context, name string) (interface{}, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Delete")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue interface{}
	)

	path := "/scheduler/schedules/{name}"
	path = strings.Replace(path, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

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
SchedulerResourceApiService Delete a tag for schedule
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param name
*/
func (a *SchedulerResourceApiService) DeleteTagForSchedule(ctx context.Context, body []model.Tag, name string) (*http.Response, error) {
	var (
		httpMethod = strings.ToUpper("Delete")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)

	path := "/scheduler/schedules/{name}/tags"
	path = strings.Replace(path, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

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
   SchedulerResourceApiService Get all existing workflow schedules and optionally filter by workflow name
   * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    * @param optional nil or *SchedulerResourceApiGetAllSchedulesOpts - Optional Parameters:
        * @param "WorkflowName" (optional.String) -
   @return []WorkflowScheduleModel
*/

type SchedulerResourceApiGetAllSchedulesOpts struct {
	WorkflowName optional.String
}

func (a *SchedulerResourceApiService) GetAllSchedules(ctx context.Context, optionals *SchedulerResourceApiGetAllSchedulesOpts) ([]model.WorkflowScheduleModel, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue []model.WorkflowScheduleModel
	)

	path := "/scheduler/schedules"

	headerParams := make(map[string]string)
	headerParams["Accept"] = "application/json"

	queryParams := url.Values{}
	formParams := url.Values{}

	if optionals != nil && optionals.WorkflowName.IsSet() {
		queryParams.Add("workflowName", parameterToString(optionals.WorkflowName.Value(), ""))
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
   SchedulerResourceApiService Get list of the next x (default 3, max 5) execution times for a scheduler
   * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    * @param cronExpression
    * @param optional nil or *SchedulerResourceApiGetNextFewSchedulesOpts - Optional Parameters:
        * @param "ScheduleStartTime" (optional.Int64) -
    * @param "ScheduleEndTime" (optional.Int64) -
    * @param "Limit" (optional.Int32) -
   @return []int64
*/

type SchedulerResourceApiGetNextFewSchedulesOpts struct {
	ScheduleStartTime optional.Int64
	ScheduleEndTime   optional.Int64
	Limit             optional.Int32
}

func (a *SchedulerResourceApiService) GetNextFewSchedules(ctx context.Context, cronExpression string, optionals *SchedulerResourceApiGetNextFewSchedulesOpts) ([]int64, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue []int64
	)

	path := "/scheduler/nextFewSchedules"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	queryParams.Add("cronExpression", parameterToString(cronExpression, ""))
	if optionals != nil && optionals.ScheduleStartTime.IsSet() {
		queryParams.Add("scheduleStartTime", parameterToString(optionals.ScheduleStartTime.Value(), ""))
	}
	if optionals != nil && optionals.ScheduleEndTime.IsSet() {
		queryParams.Add("scheduleEndTime", parameterToString(optionals.ScheduleEndTime.Value(), ""))
	}
	if optionals != nil && optionals.Limit.IsSet() {
		queryParams.Add("limit", parameterToString(optionals.Limit.Value(), ""))
	}
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
SchedulerResourceApiService Get an existing workflow schedule by name
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
    @return WorkflowSchedule
*/
func (a *SchedulerResourceApiService) GetSchedule(ctx context.Context, name string) (model.WorkflowSchedule, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue model.WorkflowSchedule
	)

	path := "/scheduler/schedules/{name}"
	path = strings.Replace(path, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

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
SchedulerResourceApiService Get tags by schedule
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
    @return []Tag
*/
func (a *SchedulerResourceApiService) GetTagsForSchedule(ctx context.Context, name string) ([]model.Tag, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue []model.Tag
	)

	path := "/scheduler/schedules/{name}/tags"
	path = strings.Replace(path, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

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
SchedulerResourceApiService Pause all scheduling in a single conductor server instance (for debugging only)
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return map[string]interface{}
*/
func (a *SchedulerResourceApiService) PauseAllSchedules(ctx context.Context) (map[string]interface{}, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue map[string]interface{}
	)

	path := "/scheduler/admin/pause"

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
SchedulerResourceApiService Pauses an existing schedule by name
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
    @return interface{}
*/
func (a *SchedulerResourceApiService) PauseSchedule(ctx context.Context, name string) (interface{}, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue interface{}
	)

	path := "/scheduler/schedules/{name}/pause"
	path = strings.Replace(path, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

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
SchedulerResourceApiService Put a tag to schedule
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param name
*/
func (a *SchedulerResourceApiService) PutTagForSchedule(ctx context.Context, body []model.Tag, name string) (*http.Response, error) {
	var (
		httpMethod = strings.ToUpper("Put")
		postBody   interface{}
		fileName   string
		fileBytes  []byte
	)

	path := "/scheduler/schedules/{name}/tags"
	path = strings.Replace(path, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

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
SchedulerResourceApiService Requeue all execution records
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return map[string]interface{}
*/
func (a *SchedulerResourceApiService) RequeueAllExecutionRecords(ctx context.Context) (map[string]interface{}, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue map[string]interface{}
	)

	path := "/scheduler/admin/requeue"

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
SchedulerResourceApiService Resume all scheduling
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @return map[string]interface{}
*/
func (a *SchedulerResourceApiService) ResumeAllSchedules(ctx context.Context) (map[string]interface{}, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue map[string]interface{}
	)

	path := "/scheduler/admin/resume"

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
SchedulerResourceApiService Resume a paused schedule by name
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param name
    @return interface{}
*/
func (a *SchedulerResourceApiService) ResumeSchedule(ctx context.Context, name string) (interface{}, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue interface{}
	)

	path := "/scheduler/schedules/{name}/resume"
	path = strings.Replace(path, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

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
SchedulerResourceApiService Create or update a schedule for a specified workflow with a corresponding start workflow request
* @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
    @return interface{}
*/
func (a *SchedulerResourceApiService) SaveSchedule(ctx context.Context, body model.SaveScheduleRequest) (interface{}, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Post")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue interface{}
	)

	path := "/scheduler/schedules"

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

/*
   SchedulerResourceApiService Search for workflows based on payload and other parameters
       use sort options as sort&#x3D;&lt;field&gt;:ASC|DESC e.g. sort&#x3D;name&amp;sort&#x3D;workflowId:DESC. If order is not specified, defaults to ASC.
   * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    * @param optional nil or *SchedulerSearchOpts - Optional Parameters:
        * @param "Start" (optional.Int32) -
    * @param "Size" (optional.Int32) -
    * @param "Sort" (optional.String) -
    * @param "FreeText" (optional.String) -
    * @param "Query" (optional.String) -
   @return SearchResultWorkflowSchedule
*/

type SchedulerSearchOpts struct {
	Start    optional.Int32
	Size     optional.Int32
	Sort     optional.String
	FreeText optional.String
	Query    optional.String
}

func (a *SchedulerResourceApiService) SearchV2(ctx context.Context, optionals *SchedulerSearchOpts) (model.SearchResultWorkflowSchedule, *http.Response, error) {
	var (
		httpMethod  = strings.ToUpper("Get")
		postBody    interface{}
		fileName    string
		fileBytes   []byte
		returnValue model.SearchResultWorkflowSchedule
	)

	path := "/scheduler/search/executions"

	headerParams := make(map[string]string)
	headerParams["Accept"] = "application/json"

	queryParams := url.Values{}
	formParams := url.Values{}

	if optionals != nil && optionals.Start.IsSet() {
		queryParams.Add("start", parameterToString(optionals.Start.Value(), ""))
	}
	if optionals != nil && optionals.Size.IsSet() {
		queryParams.Add("size", parameterToString(optionals.Size.Value(), ""))
	}
	if optionals != nil && optionals.Sort.IsSet() {
		queryParams.Add("sort", parameterToString(optionals.Sort.Value(), ""))
	}
	if optionals != nil && optionals.FreeText.IsSet() {
		queryParams.Add("freeText", parameterToString(optionals.FreeText.Value(), ""))
	}
	if optionals != nil && optionals.Query.IsSet() {
		queryParams.Add("query", parameterToString(optionals.Query.Value(), ""))
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
