# \WebhookNotificationsApi

All URIs are relative to *https://api.terminal49.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWebhookNotificationId**](WebhookNotificationsApi.md#GetWebhookNotificationId) | **Get** /webhook_notification/{id} | Get a single webhook notification
[**GetWebhookNotifications**](WebhookNotificationsApi.md#GetWebhookNotifications) | **Get** /webhook_notifications | List webhook notifications



## GetWebhookNotificationId

> InlineResponse2006 GetWebhookNotificationId(ctx, id, optional)

Get a single webhook notification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***GetWebhookNotificationIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWebhookNotificationIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **optional.String**| Comma delimited list of relations to include. | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookNotifications

> InlineResponse2007 GetWebhookNotifications(ctx, optional)

List webhook notifications

Return the list of  webhook notifications. This can be useful for reconciling your data if your endpoint has been down. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetWebhookNotificationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetWebhookNotificationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **include** | **optional.String**| Comma delimited list of relations to include. | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

