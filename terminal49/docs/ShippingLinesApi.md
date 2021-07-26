# \ShippingLinesApi

All URIs are relative to *https://api.terminal49.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetShippingLines**](ShippingLinesApi.md#GetShippingLines) | **Get** /shipping_lines | Shipping Lines
[**GetShippingLinesId**](ShippingLinesApi.md#GetShippingLinesId) | **Get** /shipping_lines/{id} | Get a single shipping line



## GetShippingLines

> InlineResponse20012 GetShippingLines(ctx, optional)

Shipping Lines

Return a list of shipping lines supported by Terminal49.  N.B. There is no pagination for this endpoint.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetShippingLinesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetShippingLinesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **optional.Map[string]interface{}**|  | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShippingLinesId

> InlineResponse20013 GetShippingLinesId(ctx, id)

Get a single shipping line

Return the details of a single shipping line.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

