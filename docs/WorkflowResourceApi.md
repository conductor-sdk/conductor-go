# {{classname}}

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Decide**](WorkflowResourceApi.md#Decide) | **Put** /api/workflow/decide/{workflowId} | Starts the decision task for a workflow
[**Delete**](WorkflowResourceApi.md#Delete) | **Delete** /api/workflow/{workflowId}/remove | Removes the workflow from the system
[**GetExecutionStatus**](WorkflowResourceApi.md#GetExecutionStatus) | **Get** /api/workflow/{workflowId} | Gets the workflow by workflow id
[**GetExternalStorageLocation**](WorkflowResourceApi.md#GetExternalStorageLocation) | **Get** /api/workflow/externalstoragelocation | Get the uri and path of the external storage where the workflow payload is to be stored
[**GetRunningWorkflow**](WorkflowResourceApi.md#GetRunningWorkflow) | **Get** /api/workflow/running/{name} | Retrieve all the running workflows
[**GetWorkflows**](WorkflowResourceApi.md#GetWorkflows) | **Post** /api/workflow/{name}/correlated | Lists workflows for the given correlation id list
[**GetWorkflows1**](WorkflowResourceApi.md#GetWorkflows1) | **Get** /api/workflow/{name}/correlated/{correlationId} | Lists workflows for the given correlation id
[**PauseWorkflow**](WorkflowResourceApi.md#PauseWorkflow) | **Put** /api/workflow/{workflowId}/pause | Pauses the workflow
[**Rerun**](WorkflowResourceApi.md#Rerun) | **Post** /api/workflow/{workflowId}/rerun | Reruns the workflow from a specific task
[**ResetWorkflow**](WorkflowResourceApi.md#ResetWorkflow) | **Post** /api/workflow/{workflowId}/resetcallbacks | Resets callback times of all non-terminal SIMPLE tasks to 0
[**Restart**](WorkflowResourceApi.md#Restart) | **Post** /api/workflow/{workflowId}/restart | Restarts a completed workflow
[**ResumeWorkflow**](WorkflowResourceApi.md#ResumeWorkflow) | **Put** /api/workflow/{workflowId}/resume | Resumes the workflow
[**Retry**](WorkflowResourceApi.md#Retry) | **Post** /api/workflow/{workflowId}/retry | Retries the last failed task
[**Search**](WorkflowResourceApi.md#Search) | **Get** /api/workflow/search | Search for workflows based on payload and other parameters
[**SearchV2**](WorkflowResourceApi.md#SearchV2) | **Get** /api/workflow/search-v2 | Search for workflows based on payload and other parameters
[**SearchWorkflowsByTasks**](WorkflowResourceApi.md#SearchWorkflowsByTasks) | **Get** /api/workflow/search-by-tasks | Search for workflows based on task parameters
[**SearchWorkflowsByTasksV2**](WorkflowResourceApi.md#SearchWorkflowsByTasksV2) | **Get** /api/workflow/search-by-tasks-v2 | Search for workflows based on task parameters
[**SkipTaskFromWorkflow**](WorkflowResourceApi.md#SkipTaskFromWorkflow) | **Put** /api/workflow/{workflowId}/skiptask/{taskReferenceName} | Skips a given task from a current running workflow
[**StartWorkflow**](WorkflowResourceApi.md#StartWorkflow) | **Post** /api/workflow/{name} | Start a new workflow. Returns the ID of the workflow instance that can be later used for tracking
[**StartWorkflow1**](WorkflowResourceApi.md#StartWorkflow1) | **Post** /api/workflow | Start a new workflow with StartWorkflowRequest, which allows task to be executed in a domain
[**Terminate1**](WorkflowResourceApi.md#Terminate1) | **Delete** /api/workflow/{workflowId} | Terminate workflow execution

# **Decide**
> Decide(ctx, workflowId)
Starts the decision task for a workflow

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Delete**
> Delete(ctx, workflowId, optional)
Removes the workflow from the system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**|  | 
 **optional** | ***WorkflowResourceApiDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowResourceApiDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archiveWorkflow** | **optional.Bool**|  | [default to true]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExecutionStatus**
> Workflow GetExecutionStatus(ctx, workflowId, optional)
Gets the workflow by workflow id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**|  | 
 **optional** | ***WorkflowResourceApiGetExecutionStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowResourceApiGetExecutionStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeTasks** | **optional.Bool**|  | [default to true]

### Return type

[**Workflow**](Workflow.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExternalStorageLocation**
> ExternalStorageLocation GetExternalStorageLocation(ctx, path, operation, payloadType)
Get the uri and path of the external storage where the workflow payload is to be stored

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

# **GetRunningWorkflow**
> []string GetRunningWorkflow(ctx, name, optional)
Retrieve all the running workflows

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
 **optional** | ***WorkflowResourceApiGetRunningWorkflowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowResourceApiGetRunningWorkflowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.Int32**|  | [default to 1]
 **startTime** | **optional.Int64**|  | 
 **endTime** | **optional.Int64**|  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWorkflows**
> map[string][]Workflow GetWorkflows(ctx, body, name, optional)
Lists workflows for the given correlation id list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]string**](string.md)|  | 
  **name** | **string**|  | 
 **optional** | ***WorkflowResourceApiGetWorkflowsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowResourceApiGetWorkflowsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeClosed** | **optional.**|  | [default to false]
 **includeTasks** | **optional.**|  | [default to false]

### Return type

[**map[string][]Workflow**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWorkflows1**
> []Workflow GetWorkflows1(ctx, name, correlationId, optional)
Lists workflows for the given correlation id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **correlationId** | **string**|  | 
 **optional** | ***WorkflowResourceApiGetWorkflows1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowResourceApiGetWorkflows1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeClosed** | **optional.Bool**|  | [default to false]
 **includeTasks** | **optional.Bool**|  | [default to false]

### Return type

[**[]Workflow**](Workflow.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PauseWorkflow**
> PauseWorkflow(ctx, workflowId)
Pauses the workflow

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Rerun**
> string Rerun(ctx, body, workflowId)
Reruns the workflow from a specific task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RerunWorkflowRequest**](RerunWorkflowRequest.md)|  | 
  **workflowId** | **string**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetWorkflow**
> ResetWorkflow(ctx, workflowId)
Resets callback times of all non-terminal SIMPLE tasks to 0

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Restart**
> Restart(ctx, workflowId, optional)
Restarts a completed workflow

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**|  | 
 **optional** | ***WorkflowResourceApiRestartOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowResourceApiRestartOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **useLatestDefinitions** | **optional.Bool**|  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResumeWorkflow**
> ResumeWorkflow(ctx, workflowId)
Resumes the workflow

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Retry**
> Retry(ctx, workflowId, optional)
Retries the last failed task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**|  | 
 **optional** | ***WorkflowResourceApiRetryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowResourceApiRetryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resumeSubworkflowTasks** | **optional.Bool**|  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Search**
> SearchResultWorkflowSummary Search(ctx, optional)
Search for workflows based on payload and other parameters

use sort options as sort=<field>:ASC|DESC e.g. sort=name&sort=workflowId:DESC. If order is not specified, defaults to ASC.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WorkflowResourceApiSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowResourceApiSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**|  | [default to 100]
 **sort** | **optional.String**|  | 
 **freeText** | **optional.String**|  | [default to *]
 **query** | **optional.String**|  | 

### Return type

[**SearchResultWorkflowSummary**](SearchResultWorkflowSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchV2**
> SearchResultWorkflow SearchV2(ctx, optional)
Search for workflows based on payload and other parameters

use sort options as sort=<field>:ASC|DESC e.g. sort=name&sort=workflowId:DESC. If order is not specified, defaults to ASC.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WorkflowResourceApiSearchV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowResourceApiSearchV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**|  | [default to 100]
 **sort** | **optional.String**|  | 
 **freeText** | **optional.String**|  | [default to *]
 **query** | **optional.String**|  | 

### Return type

[**SearchResultWorkflow**](SearchResultWorkflow.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchWorkflowsByTasks**
> SearchResultWorkflowSummary SearchWorkflowsByTasks(ctx, optional)
Search for workflows based on task parameters

use sort options as sort=<field>:ASC|DESC e.g. sort=name&sort=workflowId:DESC. If order is not specified, defaults to ASC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WorkflowResourceApiSearchWorkflowsByTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowResourceApiSearchWorkflowsByTasksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**|  | [default to 100]
 **sort** | **optional.String**|  | 
 **freeText** | **optional.String**|  | [default to *]
 **query** | **optional.String**|  | 

### Return type

[**SearchResultWorkflowSummary**](SearchResultWorkflowSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchWorkflowsByTasksV2**
> SearchResultWorkflow SearchWorkflowsByTasksV2(ctx, optional)
Search for workflows based on task parameters

use sort options as sort=<field>:ASC|DESC e.g. sort=name&sort=workflowId:DESC. If order is not specified, defaults to ASC

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WorkflowResourceApiSearchWorkflowsByTasksV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowResourceApiSearchWorkflowsByTasksV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**|  | [default to 100]
 **sort** | **optional.String**|  | 
 **freeText** | **optional.String**|  | [default to *]
 **query** | **optional.String**|  | 

### Return type

[**SearchResultWorkflow**](SearchResultWorkflow.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SkipTaskFromWorkflow**
> SkipTaskFromWorkflow(ctx, workflowId, taskReferenceName, skipTaskRequest)
Skips a given task from a current running workflow

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**|  | 
  **taskReferenceName** | **string**|  | 
  **skipTaskRequest** | [**SkipTaskRequest**](.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartWorkflow**
> string StartWorkflow(ctx, body, name, optional)
Start a new workflow. Returns the ID of the workflow instance that can be later used for tracking

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**map[string]interface{}**](map.md)|  | 
  **name** | **string**|  | 
 **optional** | ***WorkflowResourceApiStartWorkflowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowResourceApiStartWorkflowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **optional.**|  | 
 **correlationId** | **optional.**|  | 
 **priority** | **optional.**|  | [default to 0]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartWorkflow1**
> string StartWorkflow1(ctx, body)
Start a new workflow with StartWorkflowRequest, which allows task to be executed in a domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**StartWorkflowRequest**](StartWorkflowRequest.md)|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Terminate1**
> Terminate1(ctx, workflowId, optional)
Terminate workflow execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **workflowId** | **string**|  | 
 **optional** | ***WorkflowResourceApiTerminate1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowResourceApiTerminate1Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

