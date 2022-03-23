# {{classname}}

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**All**](TaskResourceApi.md#All) | **Get** /api/tasks/queue/all | Get the details about each queue
[**AllVerbose**](TaskResourceApi.md#AllVerbose) | **Get** /api/tasks/queue/all/verbose | Get the details about each queue
[**BatchPoll**](TaskResourceApi.md#BatchPoll) | **Get** /api/tasks/poll/batch/{tasktype} | Batch poll for a task of a certain type
[**GetAllPollData**](TaskResourceApi.md#GetAllPollData) | **Get** /api/tasks/queue/polldata/all | Get the last poll data for all task types
[**GetExternalStorageLocation1**](TaskResourceApi.md#GetExternalStorageLocation1) | **Get** /api/tasks/externalstoragelocation | Get the external uri where the task payload is to be stored
[**GetPollData**](TaskResourceApi.md#GetPollData) | **Get** /api/tasks/queue/polldata | Get the last poll data for a given task type
[**GetTask**](TaskResourceApi.md#GetTask) | **Get** /api/tasks/{taskId} | Get task by Id
[**GetTaskLogs**](TaskResourceApi.md#GetTaskLogs) | **Get** /api/tasks/{taskId}/log | Get Task Execution Logs
[**Log**](TaskResourceApi.md#Log) | **Post** /api/tasks/{taskId}/log | Log Task Execution Details
[**Poll**](TaskResourceApi.md#Poll) | **Get** /api/tasks/poll/{tasktype} | Poll for a task of a certain type
[**RequeuePendingTask**](TaskResourceApi.md#RequeuePendingTask) | **Post** /api/tasks/queue/requeue/{taskType} | Requeue pending tasks
[**Search1**](TaskResourceApi.md#Search1) | **Get** /api/tasks/search | Search for tasks based in payload and other parameters
[**SearchV21**](TaskResourceApi.md#SearchV21) | **Get** /api/tasks/search-v2 | Search for tasks based in payload and other parameters
[**Size**](TaskResourceApi.md#Size) | **Get** /api/tasks/queue/sizes | Get Task type queue sizes
[**UpdateTask**](TaskResourceApi.md#UpdateTask) | **Post** /api/tasks | Update a task

# **All**
> map[string]int64 All(ctx, )
Get the details about each queue

### Required Parameters
This endpoint does not need any parameter.

### Return type

**map[string]int64**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AllVerbose**
> map[string]map[string]map[string]int64 AllVerbose(ctx, )
Get the details about each queue

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**map[string]map[string]map[string]int64**](map.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BatchPoll**
> []Task BatchPoll(ctx, tasktype, optional)
Batch poll for a task of a certain type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tasktype** | **string**|  | 
 **optional** | ***TaskResourceApiBatchPollOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TaskResourceApiBatchPollOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workerid** | **optional.String**|  | 
 **domain** | **optional.String**|  | 
 **count** | **optional.Int32**|  | [default to 1]
 **timeout** | **optional.Int32**|  | [default to 100]

### Return type

[**[]Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllPollData**
> []PollData GetAllPollData(ctx, )
Get the last poll data for all task types

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]PollData**](PollData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExternalStorageLocation1**
> ExternalStorageLocation GetExternalStorageLocation1(ctx, path, operation, payloadType)
Get the external uri where the task payload is to be stored

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**|  | 
  **operation** | **string**|  | 
  **payloadType** | **string**|  | 

### Return type

[**ExternalStorageLocation**](ExternalStorageLocation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPollData**
> []PollData GetPollData(ctx, taskType)
Get the last poll data for a given task type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskType** | **string**|  | 

### Return type

[**[]PollData**](PollData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTask**
> Task GetTask(ctx, taskId)
Get task by Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 

### Return type

[**Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaskLogs**
> []TaskExecLog GetTaskLogs(ctx, taskId)
Get Task Execution Logs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 

### Return type

[**[]TaskExecLog**](TaskExecLog.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Log**
> Log(ctx, body, taskId)
Log Task Execution Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)|  | 
  **taskId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Poll**
> Task Poll(ctx, tasktype, optional)
Poll for a task of a certain type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tasktype** | **string**|  | 
 **optional** | ***TaskResourceApiPollOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TaskResourceApiPollOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workerid** | **optional.String**|  | 
 **domain** | **optional.String**|  | 

### Return type

[**Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequeuePendingTask**
> string RequeuePendingTask(ctx, taskType)
Requeue pending tasks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskType** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Search1**
> SearchResultTaskSummary Search1(ctx, optional)
Search for tasks based in payload and other parameters

use sort options as sort=<field>:ASC|DESC e.g. sort=name&sort=workflowId:DESC. If order is not specified, defaults to ASC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TaskResourceApiSearch1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TaskResourceApiSearch1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**|  | [default to 100]
 **sort** | **optional.String**|  | 
 **freeText** | **optional.String**|  | [default to *]
 **query** | **optional.String**|  | 

### Return type

[**SearchResultTaskSummary**](SearchResultTaskSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchV21**
> SearchResultTask SearchV21(ctx, optional)
Search for tasks based in payload and other parameters

use sort options as sort=<field>:ASC|DESC e.g. sort=name&sort=workflowId:DESC. If order is not specified, defaults to ASC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TaskResourceApiSearchV21Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TaskResourceApiSearchV21Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**|  | [default to 100]
 **sort** | **optional.String**|  | 
 **freeText** | **optional.String**|  | [default to *]
 **query** | **optional.String**|  | 

### Return type

[**SearchResultTask**](SearchResultTask.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Size**
> map[string]int32 Size(ctx, optional)
Get Task type queue sizes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TaskResourceApiSizeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TaskResourceApiSizeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskType** | [**optional.Interface of []string**](string.md)|  | 

### Return type

**map[string]int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTask**
> string UpdateTask(ctx, body)
Update a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TaskResult**](TaskResult.md)|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

