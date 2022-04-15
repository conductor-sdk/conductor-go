# {{classname}}

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](MetadataResourceApi.md#Create) | **Post** /api/metadata/workflow | Create a new workflow definition
[**Get**](MetadataResourceApi.md#Get) | **Get** /api/metadata/workflow/{name} | Retrieves workflow definition along with blueprint
[**GetAll**](MetadataResourceApi.md#GetAll) | **Get** /api/metadata/workflow | Retrieves all workflow definition along with blueprint
[**GetTaskDef**](MetadataResourceApi.md#GetTaskDef) | **Get** /api/metadata/taskdefs/{tasktype} | Gets the task definition
[**GetTaskDefs**](MetadataResourceApi.md#GetTaskDefs) | **Get** /api/metadata/taskdefs | Gets all task definition
[**RegisterTaskDef**](MetadataResourceApi.md#RegisterTaskDef) | **Put** /api/metadata/taskdefs | Update an existing task
[**RegisterTaskDef1**](MetadataResourceApi.md#RegisterTaskDef1) | **Post** /api/metadata/taskdefs | Create new task definition(s)
[**UnregisterTaskDef**](MetadataResourceApi.md#UnregisterTaskDef) | **Delete** /api/metadata/taskdefs/{tasktype} | Remove a task definition
[**UnregisterWorkflowDef**](MetadataResourceApi.md#UnregisterWorkflowDef) | **Delete** /api/metadata/workflow/{name}/{version} | Removes workflow definition. It does not remove workflows associated with the definition.
[**Update**](MetadataResourceApi.md#Update) | **Put** /api/metadata/workflow | Create or update workflow definition

# **Create**
> Create(ctx, body)
Create a new workflow definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WorkflowDef**](WorkflowDef.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get**
> WorkflowDef Get(ctx, name, optional)
Retrieves workflow definition along with blueprint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
 **optional** | ***MetadataResourceApiGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MetadataResourceApiGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.Int32**|  | 

### Return type

[**WorkflowDef**](WorkflowDef.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAll**
> []WorkflowDef GetAll(ctx, )
Retrieves all workflow definition along with blueprint

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]WorkflowDef**](WorkflowDef.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaskDef**
> TaskDef GetTaskDef(ctx, tasktype)
Gets the task definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tasktype** | **string**|  | 

### Return type

[**TaskDef**](TaskDef.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTaskDefs**
> []TaskDef GetTaskDefs(ctx, )
Gets all task definition

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]TaskDef**](TaskDef.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterTaskDef**
> RegisterTaskDef(ctx, body)
Update an existing task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TaskDef**](TaskDef.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterTaskDef1**
> RegisterTaskDef1(ctx, body)
Create new task definition(s)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]TaskDef**](TaskDef.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnregisterTaskDef**
> UnregisterTaskDef(ctx, tasktype)
Remove a task definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tasktype** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnregisterWorkflowDef**
> UnregisterWorkflowDef(ctx, name, version)
Removes workflow definition. It does not remove workflows associated with the definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **version** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Update**
> Update(ctx, body)
Create or update workflow definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]WorkflowDef**](WorkflowDef.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

