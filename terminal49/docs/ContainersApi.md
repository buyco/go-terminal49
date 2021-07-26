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
 **include** | **optional.String**| Comma delimited list of relations to include | 

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

> InlineResponse2009 GetContainersId(ctx, id, optional)

Get a container

Retrieves the details of a container.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***GetContainersIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetContainersIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **optional.String**| Comma delimited list of relations to include | 

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

Get a list of past and futuer (estimated) milestones for a container as reported by the carrier. Some of the data is normalized even though the API is called raw_events.   Normalized attributes: `event`, `estimated_at`, `actual_at` timestamps. Not all of the `event` values have been normalized. You can expect the the events related to container movements to be normalized but there are cases (paperwork \"milestones\" like customs release) where events are not normalized. 

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

> InlineResponse20011 GetContainersIdTransportEvents(ctx, id, optional)

Get a container's transport events

Get a list of past transport events (canonical) for a container. All data has been normalized across all carriers. These are a verified subset of the raw events may also be sent as Webhook Notifications to a webhook endpoint.  This does not provide any estiamted future events. See `container/:id/raw_events` endpoint for that.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***GetContainersIdTransportEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetContainersIdTransportEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **optional.String**| Comma delimited list of relations to include | 

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

