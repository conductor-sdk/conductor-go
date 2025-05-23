//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
//  the License. You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
//  an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
//  specific language governing permissions and limitations under the License.

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"

	"github.com/antihax/optional"
	"github.com/conductor-sdk/conductor-go/sdk/model"
)

var hostname string
var once sync.Once

// Linger please
var (
	_ context.Context
)

type TaskResourceApiService struct {
	*APIClient
}

/*
TaskResourceApiService Get the details about each queue
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return map[string]int64
*/
func (a *TaskResourceApiService) All(ctx context.Context) (map[string]int64, *http.Response, error) {
	var result map[string]int64

	path := "/tasks/queue/all"

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
TaskResourceApiService Get the details about each queue
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return map[string]map[string]map[string]int64
*/
func (a *TaskResourceApiService) AllVerbose(ctx context.Context) (map[string]map[string]map[string]int64, *http.Response, error) {
	var result map[string]map[string]map[string]int64

	path := "/tasks/queue/all/verbose"

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
TaskResourceApiService Batch poll for a task of a certain type
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param tasktype
 * @param optional nil or *TaskResourceApiBatchPollOpts - Optional Parameters:
     * @param "WorkerId" (optional.String) -
     * @param "Domain" (optional.String) -
     * @param "Count" (optional.Int32) -
     * @param "Timeout" (optional.Int32) -
@return []Task
*/

type TaskResourceApiBatchPollOpts struct {
	WorkerId optional.String
	Domain   optional.String
	Count    optional.Int32
	Timeout  optional.Int32
}

func (a *TaskResourceApiService) BatchPollTask(ctx context.Context, tasktype string, localVarOptionals *TaskResourceApiBatchPollOpts) ([]model.PolledTask, *http.Response, error) {
	var result []model.PolledTask

	path := fmt.Sprintf("/tasks/poll/batch/%s", tasktype)

	queryParams := url.Values{}
	if localVarOptionals != nil && localVarOptionals.WorkerId.IsSet() {
		queryParams.Add("workerid", parameterToString(localVarOptionals.WorkerId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Domain.IsSet() {
		queryParams.Add("domain", parameterToString(localVarOptionals.Domain.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Count.IsSet() {
		queryParams.Add("count", parameterToString(localVarOptionals.Count.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Timeout.IsSet() {
		queryParams.Add("timeout", parameterToString(localVarOptionals.Timeout.Value(), ""))
	}

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
TaskResourceApiService Get the last poll data for all task types
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return []PollData
*/
func (a *TaskResourceApiService) GetAllPollData(ctx context.Context) ([]model.PollData, *http.Response, error) {
	var result []model.PollData

	path := "/tasks/queue/polldata/all"

	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
TaskResourceApiService Get the external uri where the task payload is to be stored
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param path
  - @param operation
  - @param payloadType

@return ExternalStorageLocation
*/
func (a *TaskResourceApiService) GetExternalStorageLocation1(ctx context.Context, path string, operation string, payloadType string) (model.ExternalStorageLocation, *http.Response, error) {
	var result model.ExternalStorageLocation

	localVarPath := "/tasks/externalstoragelocation"

	queryParams := url.Values{}

	queryParams.Add("path", parameterToString(path, ""))
	queryParams.Add("operation", parameterToString(operation, ""))
	queryParams.Add("payloadType", parameterToString(payloadType, ""))

	resp, err := a.Get(ctx, localVarPath, queryParams, &result)
	if err != nil {
		return model.ExternalStorageLocation{}, resp, err
	}
	return result, resp, nil
}

/*
TaskResourceApiService Get the last poll data for a given task type
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param taskType

@return []PollData
*/
func (a *TaskResourceApiService) GetPollData(ctx context.Context, taskType string) ([]model.PollData, *http.Response, error) {
	var result []model.PollData

	path := "/tasks/queue/polldata"

	queryParams := url.Values{}
	queryParams.Add("taskType", parameterToString(taskType, ""))

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
TaskResourceApiService Get task by Id
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param taskId

@return Task
*/
func (a *TaskResourceApiService) GetTask(ctx context.Context, taskId string) (model.Task, *http.Response, error) {
	var result model.Task

	path := fmt.Sprintf("/tasks/%s", taskId)
	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return model.Task{}, resp, err
	}

	return result, resp, nil
}

/*
TaskResourceApiService Get Task Execution Logs
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param taskId

@return []TaskExecLog
*/
func (a *TaskResourceApiService) GetTaskLogs(ctx context.Context, taskId string) ([]model.TaskExecLog, *http.Response, error) {
	var result []model.TaskExecLog

	path := fmt.Sprintf("/tasks/%s/log", taskId)
	resp, err := a.Get(ctx, path, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
TaskResourceApiService Log Task Execution Details
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param taskId
*/
func (a *TaskResourceApiService) Log(ctx context.Context, body string, taskId string) (*http.Response, error) {

	path := fmt.Sprintf("/tasks/%s/log", taskId)
	resp, err := a.Post(ctx, path, body, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

/*
TaskResourceApiService Poll for a task of a certain type
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param tasktype
 * @param optional nil or *TaskResourceApiPollOpts - Optional Parameters:
     * @param "WorkerId" (optional.String) -
     * @param "Domain" (optional.String) -
@return Task
*/

type TaskResourceApiPollOpts struct {
	Workerid optional.String
	Domain   optional.String
}

func (a *TaskResourceApiService) Poll(ctx context.Context, tasktype string, localVarOptionals *TaskResourceApiPollOpts) (model.Task, *http.Response, error) {
	var result model.Task

	path := fmt.Sprintf("/tasks/poll/%s", tasktype)

	queryParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Workerid.IsSet() {
		queryParams.Add("workerid", parameterToString(localVarOptionals.Workerid.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Domain.IsSet() {
		queryParams.Add("domain", parameterToString(localVarOptionals.Domain.Value(), ""))
	}

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return model.Task{}, resp, err
	}
	return result, resp, nil
}

/*
TaskResourceApiService Requeue pending tasks
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param taskType

@return string
*/
func (a *TaskResourceApiService) RequeuePendingTask(ctx context.Context, taskType string) (string, *http.Response, error) {
	var result string

	path := fmt.Sprintf("/tasks/queue/requeue/%s", taskType)
	resp, err := a.Post(ctx, path, nil, &result)
	if err != nil {
		return "", resp, err
	}
	return result, resp, nil
}

/*
TaskResourceApiService Search for tasks based in payload and other parameters
use sort options as sort&#x3D;&lt;field&gt;:ASC|DESC e.g. sort&#x3D;name&amp;sort&#x3D;workflowId:DESC. If order is not specified, defaults to ASC
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *TaskResourceApiSearch1Opts - Optional Parameters:
     * @param "Start" (optional.Int32) -
     * @param "Size" (optional.Int32) -
     * @param "Sort" (optional.String) -
     * @param "FreeText" (optional.String) -
     * @param "Query" (optional.String) -
@return SearchResultTaskSummary
*/

type TaskResourceApiSearch1Opts struct {
	Start    optional.Int32
	Size     optional.Int32
	Sort     optional.String
	FreeText optional.String
	Query    optional.String
}

func (a *TaskResourceApiService) Search(ctx context.Context, localVarOptionals *TaskResourceApiSearch1Opts) (model.SearchResultTaskSummary, *http.Response, error) {
	var result model.SearchResultTaskSummary

	path := "/tasks/search"

	queryParams := url.Values{}
	if localVarOptionals != nil && localVarOptionals.Start.IsSet() {
		queryParams.Add("start", parameterToString(localVarOptionals.Start.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Size.IsSet() {
		queryParams.Add("size", parameterToString(localVarOptionals.Size.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		queryParams.Add("sort", parameterToString(localVarOptionals.Sort.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FreeText.IsSet() {
		queryParams.Add("freeText", parameterToString(localVarOptionals.FreeText.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Query.IsSet() {
		queryParams.Add("query", parameterToString(localVarOptionals.Query.Value(), ""))
	}

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return model.SearchResultTaskSummary{}, resp, err
	}
	return result, resp, nil
}

/*
TaskResourceApiService Search for tasks based in payload and other parameters
use sort options as sort&#x3D;&lt;field&gt;:ASC|DESC e.g. sort&#x3D;name&amp;sort&#x3D;workflowId:DESC. If order is not specified, defaults to ASC
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *TaskResourceApiSearchV21Opts - Optional Parameters:
     * @param "Start" (optional.Int32) -
     * @param "Size" (optional.Int32) -
     * @param "Sort" (optional.String) -
     * @param "FreeText" (optional.String) -
     * @param "Query" (optional.String) -
@return SearchResultTask
*/

type TaskResourceApiSearchV21Opts struct {
	Start    optional.Int32
	Size     optional.Int32
	Sort     optional.String
	FreeText optional.String
	Query    optional.String
}

func (a *TaskResourceApiService) SearchV2(ctx context.Context, localVarOptionals *TaskResourceApiSearchV21Opts) (model.SearchResultTask, *http.Response, error) {
	var result model.SearchResultTask

	path := "/tasks/search-v2"

	queryParams := url.Values{}
	if localVarOptionals != nil && localVarOptionals.Start.IsSet() {
		queryParams.Add("start", parameterToString(localVarOptionals.Start.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Size.IsSet() {
		queryParams.Add("size", parameterToString(localVarOptionals.Size.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Sort.IsSet() {
		queryParams.Add("sort", parameterToString(localVarOptionals.Sort.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FreeText.IsSet() {
		queryParams.Add("freeText", parameterToString(localVarOptionals.FreeText.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Query.IsSet() {
		queryParams.Add("query", parameterToString(localVarOptionals.Query.Value(), ""))
	}

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return model.SearchResultTask{}, resp, err
	}
	return result, resp, nil
}

/*
TaskResourceApiService Get Task type queue sizes
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *TaskResourceApiSizeOpts - Optional Parameters:
     * @param "TaskType" (optional.Interface of []string) -
@return map[string]int32
*/

type TaskResourceApiSizeOpts struct {
	TaskType optional.Interface
}

func (a *TaskResourceApiService) Size(ctx context.Context, localVarOptionals *TaskResourceApiSizeOpts) (map[string]int32, *http.Response, error) {
	var result map[string]int32

	path := "/tasks/queue/sizes"

	queryParams := url.Values{}
	if localVarOptionals != nil && localVarOptionals.TaskType.IsSet() {
		queryParams.Add("taskType", parameterToString(localVarOptionals.TaskType.Value(), "multi"))
	}

	resp, err := a.Get(ctx, path, queryParams, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

/*
TaskResourceApiService Update a task
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body

@return string
*/
func (a *TaskResourceApiService) UpdateTask(ctx context.Context, taskResult *model.TaskResult) (string, *http.Response, error) {
	var result string

	path := "/tasks"

	resp, err := a.Post(ctx, path, taskResult, &result)
	if err != nil {
		return "", resp, err
	}
	return result, resp, nil
}

/*
TaskResourceApiService Update a task By Ref Name
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param workflowId
  - @param taskRefName
  - @param status

@return string
*/
func (a *TaskResourceApiService) UpdateTaskByRefName(ctx context.Context, body map[string]interface{}, workflowId string, taskRefName string, status string) (string, *http.Response, error) {
	return a.updateTaskByRefName(ctx, body, workflowId, taskRefName, status, optional.EmptyString())
}

/*
TaskResourceApiService Update a task By Ref Name
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param workflowId
  - @param taskRefName
  - @param status
  - @param workerId

@return string
*/
func (a *TaskResourceApiService) UpdateTaskByRefNameWithWorkerId(ctx context.Context, body map[string]interface{}, workflowId string, taskRefName string, status string, workerId optional.String) (string, *http.Response, error) {
	if workerId.IsSet() {
		return a.updateTaskByRefName(ctx, body, workflowId, taskRefName, status, workerId)
	}
	return a.updateTaskByRefName(ctx, body, workflowId, taskRefName, status, optional.NewString(getHostname()))
}

func (a *TaskResourceApiService) updateTaskByRefName(ctx context.Context, body map[string]interface{}, workflowId string, taskRefName string, status string, workerId optional.String) (string, *http.Response, error) {
	var result string

	localVarPath := fmt.Sprintf("/tasks/%s/%s/%s", workflowId, taskRefName, status)

	queryParams := url.Values{}
	if workerId.IsSet() {
		queryParams.Add("workerid", workerId.Value())
	}

	resp, err := a.PostWithParams(ctx, localVarPath, queryParams, body, &result)
	if err != nil {
		return "", resp, err
	}
	return result, resp, nil
}

func getHostname() string {
	once.Do(updateHostname)
	return hostname
}

func updateHostname() {
	hostname, _ = os.Hostname()
}
