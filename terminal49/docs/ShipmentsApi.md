# \ShipmentsApi

All URIs are relative to *https://api.terminal49.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetShipmentId**](ShipmentsApi.md#GetShipmentId) | **Get** /shipments/{id} | Get a shipment
[**GetShipments**](ShipmentsApi.md#GetShipments) | **Get** /shipments | List shipments



## GetShipmentId

> InlineResponse2001 GetShipmentId(ctx, id, optional)

Get a shipment

Retrieves the details of an existing shipment. You need only supply the unique shipment `id` that was returned upon `tracking_request` creation.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| Shipment Id | 
 **optional** | ***GetShipmentIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetShipmentIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **optional.String**| Comma delimited list of relations to include | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShipments

> InlineResponse200 GetShipments(ctx, optional)

List shipments

Returns a list of your shipments. The shipments are returned sorted by creation date, with the most recent shipments appearing first.  This api will return all shipments associated with the account. Shipments created via the `tracking_request` API aswell as the ones added via the dashboard will be retuned via this endpoint. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetShipmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetShipmentsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **optional.Int32**|  | [default to 1]
 **pageSize** | **optional.Int32**|  | [default to 30]
 **q** | **optional.String**|  Search shipments by master bill of lading, reference number, or container number. | 
 **include** | **optional.String**| Comma delimited list of relations to include | 
 **number** | **optional.String**| Search shipments by the original request tracking &#x60;request_number&#x60; | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[authorization](../README.md#authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

