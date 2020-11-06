# \WebhooksApi

All URIs are relative to *https://api.terminal49.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWebhooksId**](WebhooksApi.md#DeleteWebhooksId) | **Delete** /webhooks/{id} | Delete a webhook
[**GetWebhooks**](WebhooksApi.md#GetWebhooks) | **Get** /webhooks | List webhooks
[**GetWebhooksId**](WebhooksApi.md#GetWebhooksId) | **Get** /webhooks/{id} | Get single webhook
[**PatchWebhooksId**](WebhooksApi.md#PatchWebhooksId) | **Patch** /webhooks/{id} | Edit a webhook
[**PostWebhooks**](WebhooksApi.md#PostWebhooks) | **Post** /webhooks | Create a webhook



## DeleteWebhooksId

> DeleteWebhooksId(ctx, id)

Delete a webhook

Delete a webhook

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooks

> InlineResponse2005 GetWebhooks(ctx, )

List webhooks

Get a list of all the webhooks

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooksId

> InlineResponse2004 GetWebhooksId(ctx, id)

Get single webhook

Get the details of a single webhook

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchWebhooksId

> InlineResponse2004 PatchWebhooksId(ctx, id, optional)

Edit a webhook

Update a single webhook

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***PatchWebhooksIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PatchWebhooksIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject1** | [**optional.Interface of InlineObject1**](InlineObject1.md)|  | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWebhooks

> InlineResponse2004 PostWebhooks(ctx, optional)

Create a webhook

You can configure a webhook via the API to be notified about events that happen in your Terminal49 account. These events can be realted to tracking_requests, shipments and containers.   This is the recommended way tracking shipments and containers via the API. You should use this instead of polling our the API periodically.  

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostWebhooksOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostWebhooksOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject2** | [**optional.Interface of InlineObject2**](InlineObject2.md)|  | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

