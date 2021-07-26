# ContainerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** |  | [optional] 
**EquipmentType** | Pointer to **string** |  | [optional] 
**EquipmentLength** | Pointer to **int32** |  | [optional] 
**EquipmentHeight** | Pointer to **string** |  | [optional] 
**WeightInLbs** | Pointer to **float32** |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**SealNumber** | Pointer to **string** |  | [optional] 
**PickupLfd** | Pointer to [**time.Time**](time.Time.md) | The last free day for pickup before deummurage accrues. | [optional] 
**PickupAppointmentAt** | Pointer to [**time.Time**](time.Time.md) | When available the pickup appointment time at the terminal is returned. | [optional] 
**AvailabilityKnown** | **bool** | Whether Terminal 49 is receiving availability status from the terminal. | [optional] 
**AvailableForPickup** | Pointer to **bool** | If availability_known is true, then whether container is available to be picked up at terminal. | [optional] 
**PodArrivedAt** | Pointer to [**time.Time**](time.Time.md) | Time the vessel arrived at the POD | [optional] 
**PodDischargedAt** | Pointer to [**time.Time**](time.Time.md) | Discharge time at the port of discharge | [optional] 
**PodFullOutAt** | Pointer to [**time.Time**](time.Time.md) | Full Out time at port of discharge. Null for inland moves. | [optional] 
**PodFullOutChassisNumber** | Pointer to **string** | The chassis number used when container was picked up at POD (if available) | [optional] 
**FinalDestinationFullOutAt** | Pointer to [**time.Time**](time.Time.md) | Pickup time at final destination for inland moves. | [optional] 
**EmptyTerminatedAt** | Pointer to [**time.Time**](time.Time.md) | Time empty container was returned. | [optional] 
**HoldsAtPodTerminal** | [**[]TerminalHold**](terminal_hold.md) |  | [optional] 
**FeesAtPodTerminal** | [**[]TerminalFee**](terminal_fee.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


