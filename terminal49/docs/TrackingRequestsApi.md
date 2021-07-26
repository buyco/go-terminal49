# \TrackingRequestsApi

All URIs are relative to *https://api.terminal49.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTrackRequestById**](TrackingRequestsApi.md#GetTrackRequestById) | **Get** /tracking_requests/{id} | Get a single tracking request
[**GetTrackingRequests**](TrackingRequestsApi.md#GetTrackingRequests) | **Get** /tracking_requests | List tracking requests
[**PostTrack**](TrackingRequestsApi.md#PostTrack) | **Post** /tracking_requests | Create a tracking request



## GetTrackRequestById

> InlineResponse2003 GetTrackRequestById(ctx, id, optional)

Get a single tracking request

Get the details and status of an existing tracking request. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| Tracking Request ID | 
 **optional** | ***GetTrackRequestByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTrackRequestByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **optional.String**| Comma delimited list of relations to include. &#39;tracked_object&#39; is included by default. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrackingRequests

> InlineResponse2002 GetTrackingRequests(ctx, optional)

List tracking requests

Returns a list of your tracking requests. The trackig requests are returned sorted by creation date, with the most recent tracking request appearing first.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTrackingRequestsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTrackingRequestsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| A search term to be applied against request_number and reference_numbers. | 
 **filterStatus** | **optional.String**| filter by &#x60;status&#x60; | 
 **filterScac** | **optional.String**| filter by shipping line &#x60;scac&#x60; | 
 **filterCreatedAtStart** | **optional.Time**| filter by tracking_requests &#x60;created_at&#x60; after a certain ISO8601 timestamp | 
 **filterCreatedAtEnd** | **optional.Time**| filter by tracking_requests &#x60;created_at&#x60; before a certain ISO8601 timestamp | 
 **include** | **optional.String**| Comma delimited list of relations to include. &#39;tracked_object&#39; is included by default. | 
 **pageNumber** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 
 **requestNumber** | **optional.String**| filter by tracking_requests &#x60;request_number&#x60; | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTrack

> InlineResponse201 PostTrack(ctx, optional)

Create a tracking request

To track an ocean shipment, you create a new tracking request.  Two attributes are required to track a shipment. A `bill of lading/booking number` and a shipping line `SCAC`.   Once a tracking request is created we will attempt to fetch the shipment details and it's related containers from the shipping line. If the attempt is successful we will create in new shipment object including any related container objects. We will send a `tracking_Request.succeeded` webhook notification to your webhooks.    If the attempt to fetch fails then we will send a `tracking_Request.failed` webhook notification to your `webhooks`.    A `tracking_Request.succeeded` or `tracking_Request.failed` webhook notificaiton will only be sent  if you have  atleast one active webhook. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostTrackOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostTrackOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject** | [**optional.Interface of InlineObject**](InlineObject.md)|  | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

