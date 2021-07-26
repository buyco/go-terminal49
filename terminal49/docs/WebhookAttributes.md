# WebhookAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | https end point | 
**Active** | **bool** | whether the webhook will be delivered when events are triggered | [default to true]
**Events** | **[]string** |  | 
**Secret** | **string** | A random token that will sign all delivered webhooks | 
**Headers** | Pointer to [**[]WebhookAttributesHeaders**](webhook_attributes_headers.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


