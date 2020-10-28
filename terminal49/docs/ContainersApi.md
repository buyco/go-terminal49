# \ContainersApi

All URIs are relative to *https://api.terminal49.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContainers**](ContainersApi.md#GetContainers) | **Get** /containers | List containers
[**GetContainersId**](ContainersApi.md#GetContainersId) | **Get** /containers/{id} | Get a container



## GetContainers

> InlineResponse2006 GetContainers(ctx, optional)

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

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[Authorization_Header](../README.md#Authorization_Header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainersId

> InlineResponse2007 GetContainersId(ctx, id)

Get a container

Retrieves the details of a container.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[Authorization_Header](../README.md#Authorization_Header)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

