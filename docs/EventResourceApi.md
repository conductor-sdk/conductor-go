# {{classname}}

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEventHandler**](EventResourceApi.md#AddEventHandler) | **Post** /api/event | Add a new event handler.
[**GetEventHandlers**](EventResourceApi.md#GetEventHandlers) | **Get** /api/event | Get all the event handlers
[**GetEventHandlersForEvent**](EventResourceApi.md#GetEventHandlersForEvent) | **Get** /api/event/{event} | Get event handlers for a given event
[**RemoveEventHandlerStatus**](EventResourceApi.md#RemoveEventHandlerStatus) | **Delete** /api/event/{name} | Remove an event handler
[**UpdateEventHandler**](EventResourceApi.md#UpdateEventHandler) | **Put** /api/event | Update an existing event handler.

# **AddEventHandler**
> AddEventHandler(ctx, body)
Add a new event handler.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventHandler**](EventHandler.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventHandlers**
> []EventHandler GetEventHandlers(ctx, )
Get all the event handlers

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]EventHandler**](EventHandler.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventHandlersForEvent**
> []EventHandler GetEventHandlersForEvent(ctx, event, optional)
Get event handlers for a given event

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **event** | **string**|  | 
 **optional** | ***EventResourceApiGetEventHandlersForEventOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventResourceApiGetEventHandlersForEventOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activeOnly** | **optional.Bool**|  | [default to true]

### Return type

[**[]EventHandler**](EventHandler.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveEventHandlerStatus**
> RemoveEventHandlerStatus(ctx, name)
Remove an event handler

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventHandler**
> UpdateEventHandler(ctx, body)
Update an existing event handler.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventHandler**](EventHandler.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

