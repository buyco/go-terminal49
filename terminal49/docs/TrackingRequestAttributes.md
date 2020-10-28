# TrackingRequestAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestNumber** | **string** |  | 
**RefNumbers** | **[]string** |  | [optional] 
**Tags** | **[]string** |  | [optional] 
**Status** | **string** |  | 
**FailedReason** | **string** | If the tracking request has failed, or is currently failing, the last reason we were unable to complete the request | [optional] 
**RequestType** | **string** |  | 
**Scac** | **string** |  | 
**CreatedAt** | [**time.Time**](time.Time.md) |  | 
**IsRetrying** | **bool** |  | [optional] 
**RetryCount** | **int32** | How many times T49 has attempted to get the shipment from the shipping line | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


