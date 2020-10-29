# ContainerUpdatedEventAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changeset** | [**map[string]interface{}**](.md) | A hash of all the changed attributes with the values being an array of the before and after. E.g. &#x60;{\&quot;pickup_lfd\&quot;: [null, \&quot;2020-05-20\&quot;]}&#x60;  The current attributes that can be alerted on are: - &#x60;available_for_pickup&#x60; - &#x60;pickup_lfd&#x60; - &#x60;fees_at_pod_terminal&#x60; - &#x60;holds_at_pod_terminal&#x60; | 
**Timestamp** | [**time.Time**](time.Time.md) |  | 
**Timezone** | **string** | IANA TZ  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


