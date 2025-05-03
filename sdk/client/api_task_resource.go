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
	resp, err := a.Get(ctx, "/tasks/queue/all", nil, &result)
	return result, resp, err
}

/*
TaskResourceApiService Get the details about each queue
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return map[string]map[string]map[string]int64
*/
func (a *TaskResourceApiService) AllVerbose(ctx context.Context) (map[string]map[string]map[string]int64, *http.Response, error) {
	var result map[string]map[string]map[string]int64
	resp, err := a.Get(ctx, "/tasks/queue/all/verbose", nil, &result)
	return result, resp, err
}

/*
TaskResourceApiService Batch poll for a task of a certain type
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param tasktype
 * @param optional nil or *TaskResourceApiBatchPollOpts - Optional Parameters:
     * @param "Workerid" (optional.String) -
     * @param "Domain" (optional.String) -
     * @param "Count" (optional.Int32) -
     * @param "Timeout" (optional.Int32) -
@return []Task
*/

type TaskResourceApiBatchPollOpts struct {
	Workerid optional.String
	Domain   optional.String
	Count    optional.Int32
	Timeout  optional.Int32
}

// BatchPoll polls for multiple tasks of the specified type
func (a *TaskResourceApiService) BatchPoll(ctx context.Context, tasktype string, localVarOptionals *TaskResourceApiBatchPollOpts) ([]model.Task, *http.Response, error) {
	returnValue := new([]model.Task)
	_, response, err := a.batchPoll(ctx, tasktype, localVarOptionals, returnValue)
	return *returnValue, response, err
}

// BatchPollTask polls for multiple tasks of the specified type and returns them as PolledTask
func (a *TaskResourceApiService) BatchPollTask(ctx context.Context, tasktype string, localVarOptionals *TaskResourceApiBatchPollOpts) ([]model.PolledTask, *http.Response, error) {
	returnValue := new([]model.PolledTask)
	_, response, err := a.batchPoll(ctx, tasktype, localVarOptionals, returnValue)
	return *returnValue, response, err
}

// batchPoll is a helper method for batch polling
func (a *TaskResourceApiService) batchPoll(ctx context.Context, tasktype string, opts *TaskResourceApiBatchPollOpts, returnValue interface{}) (interface{}, *http.Response, error) {
	// Build path
	path := fmt.Sprintf("/tasks/poll/batch/%s", tasktype)

	// Build query parameters
	queryParams := url.Values{}
	if opts != nil {
		if opts.Workerid.IsSet() {
			queryParams.Add("workerid", parameterToString(opts.Workerid.Value(), ""))
		}
		if opts.Domain.IsSet() {
			queryParams.Add("domain", parameterToString(opts.Domain.Value(), ""))
		}
		if opts.Count.IsSet() {
			queryParams.Add("count", parameterToString(opts.Count.Value(), ""))
		}
		if opts.Timeout.IsSet() {
			queryParams.Add("timeout", parameterToString(opts.Timeout.Value(), ""))
		}
	}

	// Make request
	resp, err := a.Get(ctx, path, queryParams, returnValue)
	return returnValue, resp, err
}

/*
TaskResourceApiService Get the last poll data for all task types
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return []PollData
*/
func (a *TaskResourceApiService) GetAllPollData(ctx context.Context) ([]model.PollData, *http.Response, error) {
	var result []model.PollData
	resp, err := a.Get(ctx, "/tasks/queue/polldata/all", nil, &result)
	return result, resp, err
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

	// Build query parameters
	queryParams := url.Values{}
	queryParams.Add("path", parameterToString(path, ""))
	queryParams.Add("operation", parameterToString(operation, ""))
	queryParams.Add("payloadType", parameterToString(payloadType, ""))

	resp, err := a.Get(ctx, "/tasks/externalstoragelocation", queryParams, &result)
	return result, resp, err
}

/*
TaskResourceApiService Get the last poll data for a given task type
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param taskType

@return []PollData
*/
func (a *TaskResourceApiService) GetPollData(ctx context.Context, taskType string) ([]model.PollData, *http.Response, error) {
	var result []model.PollData

	// Build query parameters
	queryParams := url.Values{}
	queryParams.Add("taskType", parameterToString(taskType, ""))

	resp, err := a.Get(ctx, "/tasks/queue/polldata", queryParams, &result)
	return result, resp, err
}

/*
TaskResourceApiService Get task by Id
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param taskId

@return Task
*/
func (a *TaskResourceApiService) GetTask(ctx context.Context, taskId string) (model.Task, *http.Response, error) {
	var result model.Task
	resp, err := a.Get(ctx, fmt.Sprintf("/tasks/%s", taskId), nil, &result)
	return result, resp, err
}

/*
TaskResourceApiService Get Task Execution Logs
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param taskId

@return []TaskExecLog
*/
func (a *TaskResourceApiService) GetTaskLogs(ctx context.Context, taskId string) ([]model.TaskExecLog, *http.Response, error) {
	var result []model.TaskExecLog
	resp, err := a.Get(ctx, fmt.Sprintf("/tasks/%s/log", taskId), nil, &result)
	return result, resp, err
}

/*
TaskResourceApiService Log Task Execution Details
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body
  - @param taskId
*/
func (a *TaskResourceApiService) Log(ctx context.Context, body string, taskId string) (*http.Response, error) {
	return a.Post(ctx, fmt.Sprintf("/tasks/%s/log", taskId), body, nil)
}

/*
TaskResourceApiService Poll for a task of a certain type
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param tasktype
 * @param optional nil or *TaskResourceApiPollOpts - Optional Parameters:
     * @param "Workerid" (optional.String) -
     * @param "Domain" (optional.String) -
@return Task
*/

type TaskResourceApiPollOpts struct {
	Workerid optional.String
	Domain   optional.String
}

func (a *TaskResourceApiService) Poll(ctx context.Context, tasktype string, opts *TaskResourceApiPollOpts) (model.Task, *http.Response, error) {
	var result model.Task

	// Build path
	path := fmt.Sprintf("/tasks/poll/%s", tasktype)

	// Build query parameters
	queryParams := url.Values{}
	if opts != nil {
		if opts.Workerid.IsSet() {
			queryParams.Add("workerid", opts.Workerid.Value())
		}
		if opts.Domain.IsSet() {
			queryParams.Add("domain", opts.Domain.Value())
		}
	}

	resp, err := a.Get(ctx, path, queryParams, &result)
	return result, resp, err
}

/*
TaskResourceApiService Requeue pending tasks
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param taskType

@return string
*/
func (a *TaskResourceApiService) RequeuePendingTask(ctx context.Context, taskType string) (string, *http.Response, error) {
	var result string
	resp, err := a.Post(ctx, fmt.Sprintf("/tasks/queue/requeue/%s", taskType), nil, &result)
	return result, resp, err
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

func (a *TaskResourceApiService) Search(ctx context.Context, opts *TaskResourceApiSearch1Opts) (model.SearchResultTaskSummary, *http.Response, error) {
	var result model.SearchResultTaskSummary

	// Build query parameters
	queryParams := url.Values{}
	if opts != nil {
		if opts.Start.IsSet() {
			queryParams.Add("start", parameterToString(opts.Start.Value(), ""))
		}
		if opts.Size.IsSet() {
			queryParams.Add("size", parameterToString(opts.Size.Value(), ""))
		}
		if opts.Sort.IsSet() {
			queryParams.Add("sort", parameterToString(opts.Sort.Value(), ""))
		}
		if opts.FreeText.IsSet() {
			queryParams.Add("freeText", parameterToString(opts.FreeText.Value(), ""))
		}
		if opts.Query.IsSet() {
			queryParams.Add("query", parameterToString(opts.Query.Value(), ""))
		}
	}

	resp, err := a.Get(ctx, "/tasks/search", queryParams, &result)
	return result, resp, err
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

func (a *TaskResourceApiService) SearchV2(ctx context.Context, opts *TaskResourceApiSearchV21Opts) (model.SearchResultTask, *http.Response, error) {
	var result model.SearchResultTask

	// Build query parameters
	queryParams := url.Values{}
	if opts != nil {
		if opts.Start.IsSet() {
			queryParams.Add("start", parameterToString(opts.Start.Value(), ""))
		}
		if opts.Size.IsSet() {
			queryParams.Add("size", parameterToString(opts.Size.Value(), ""))
		}
		if opts.Sort.IsSet() {
			queryParams.Add("sort", parameterToString(opts.Sort.Value(), ""))
		}
		if opts.FreeText.IsSet() {
			queryParams.Add("freeText", parameterToString(opts.FreeText.Value(), ""))
		}
		if opts.Query.IsSet() {
			queryParams.Add("query", parameterToString(opts.Query.Value(), ""))
		}
	}

	resp, err := a.Get(ctx, "/tasks/search-v2", queryParams, &result)
	return result, resp, err
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

func (a *TaskResourceApiService) Size(ctx context.Context, opts *TaskResourceApiSizeOpts) (map[string]int32, *http.Response, error) {
	var result map[string]int32

	// Build query parameters
	queryParams := url.Values{}
	if opts != nil && opts.TaskType.IsSet() {
		queryParams.Add("taskType", parameterToString(opts.TaskType.Value(), "multi"))
	}

	resp, err := a.Get(ctx, "/tasks/queue/sizes", queryParams, &result)
	return result, resp, err
}

/*
TaskResourceApiService Update a task
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param body

@return string
*/
func (a *TaskResourceApiService) UpdateTask(ctx context.Context, taskResult *model.TaskResult) (string, *http.Response, error) {
	var result string
	resp, err := a.Post(ctx, "/tasks", taskResult, &result)
	return result, resp, err
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

	// Build path
	path := fmt.Sprintf("/tasks/%s/%s/%s", workflowId, taskRefName, status)

	// Build query parameters
	queryParams := url.Values{}
	if workerId.IsSet() {
		queryParams.Add("workerid", workerId.Value())
	}

	resp, err := a.PostWithParams(ctx, path, queryParams, body, &result)
	return result, resp, err
}

func getHostname() string {
	once.Do(updateHostname)
	return hostname
}

func updateHostname() {
	hostname, _ = os.Hostname()
}
