# \ContainersApi

All URIs are relative to *https://api.terminal49.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContainers**](ContainersApi.md#GetContainers) | **Get** /containers | List containers
[**GetContainersId**](ContainersApi.md#GetContainersId) | **Get** /containers/{id} | Get a container
[**GetContainersIdRawEvents**](ContainersApi.md#GetContainersIdRawEvents) | **Get** /containers/{id}/raw_events | Get a container&#39;s raw events
[**GetContainersIdTransportEvents**](ContainersApi.md#GetContainersIdTransportEvents) | **Get** /containers/{id}/transport_events | Get a container&#39;s transport events



## GetContainers

> InlineResponse2008 GetContainers(ctx, optional)

List containers

Returns a list of container. The containers are returned sorted by creation date, with the most recent container appearing first.  This API will return all containers associated with the account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetContainersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetContainersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**|  | [default to 1]
 **pageSize** | **optional.Int32**|  | [default to 30]

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainersId

> InlineResponse2009 GetContainersId(ctx, id)

Get a container

Retrieves the details of a container.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainersIdRawEvents

> InlineResponse20010 GetContainersIdRawEvents(ctx, id)

Get a container's raw events

Returns the carrier raw_events for a container

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainersIdTransportEvents

> InlineResponse20011 GetContainersIdTransportEvents(ctx, id)

Get a container's transport events

The canonical transport events for the container.  These are a verified subset of the raw events may also be sent as push notifications. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

