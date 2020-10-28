# ContainerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** |  | [optional] 
**EquipmentType** | **string** |  | [optional] 
**EquipmentLength** | **int32** |  | [optional] 
**EquipmentHeight** | **string** |  | [optional] 
**WeightInLbs** | **string** |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**SealNumber** | **string** |  | [optional] 
**PickupLfd** | [**time.Time**](time.Time.md) | The last free day for pickup before deummurage accrues. | [optional] 
**PickupApopintmentAt** | [**time.Time**](time.Time.md) | When available the pickup appointment time at the terminal is returned. | [optional] 
**AvailabilityKnown** | **bool** | Whether Terminal 49 is receiving availability status from the terminal. | [optional] 
**AvailableForPickup** | **bool** | If availability_known is true, then whether container is available to be picked up at terminal. | [optional] 
**CurrentTransportationMode** | **string** | &#x60;unreleased&#x60; | [optional] 
**PodDischargedAt** | [**time.Time**](time.Time.md) | &#x60;unreleased&#x60; Discharge time at the port of discharge | [optional] 
**PodFullOutAt** | [**time.Time**](time.Time.md) | Full Out time at port of discharge. Null for inland moves. | [optional] 
**DestinationUnloadedAt** | [**time.Time**](time.Time.md) | &#x60;unreleased&#x60; Arrival or unloading time at final destination for inland moves. | [optional] 
**FinalDestinationFullOutAt** | [**time.Time**](time.Time.md) | &#x60;unreleased&#x60; Pickup time at final destination for inland moves. | [optional] 
**EmptyTerminatedAt** | [**time.Time**](time.Time.md) | &#x60;unreleased&#x60; Time empty container was returned. | [optional] 
**PodTimezone** | **string** | &#x60;unreleased&#x60; IANA tz | [optional] 
**DestinationTimezone** | **string** | &#x60;unreleaed&#x60; IANA tz | [optional] 
**EmptyReturnedTimezone** | **string** | &#x60;unreleased&#x60; IANA tz | [optional] 
**HoldsAtPodTerminal** | [**[]TerminalHold**](terminal_hold.md) |  | [optional] 
**FeesAtPodTerminal** | [**[]TerminalFee**](terminal_fee.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


