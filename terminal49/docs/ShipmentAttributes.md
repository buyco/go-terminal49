# ShipmentAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillOfLadingNumber** | **string** |  | 
**RefNumbers** | **[]string** |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**Tags** | **[]string** | &#x60;unreleased&#x60; | [optional] 
**PortOfLadingLocode** | **string** | UN/LOCODE | [optional] 
**PortOfLadingName** | **string** |  | [optional] 
**PortOfDischargeLocode** | **string** | UN/LOCODE | [optional] 
**PortOfDischargeName** | **string** |  | [optional] 
**DestinationLocode** | **string** | UN/LOCODE | [optional] 
**DestinationName** | **string** |  | [optional] 
**ShippingLineScac** | **string** |  | [optional] 
**ShippingLineName** | **string** |  | [optional] 
**PodVesselName** | **string** |  | [optional] 
**PodVesselImo** | **string** |  | [optional] 
**PodVoyageNumber** | **string** |  | [optional] 
**PolEtdAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**PolAtdAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**PodEtaAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**PodAtaAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**DestinationEtaAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**DestinationAtaAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**PolTimezone** | Pointer to **string** | IANA tz | [optional] 
**PodTimezone** | Pointer to **string** | IANA tz | [optional] 
**DestinationTimezone** | Pointer to **string** | IANA tz | [optional] 
**LineTrackingLastAttemptedAt** | Pointer to [**time.Time**](time.Time.md) | When Terminal49 last tried to update the shipment status from the shipping line | [optional] 
**LineTrackingLastSucceededAt** | Pointer to [**time.Time**](time.Time.md) | When Terminal49 last successfully updated the shipment status from the shipping line | [optional] 
**LineTrackingStoppedAt** | Pointer to [**time.Time**](time.Time.md) | When Terminal49 stopped checking at the shipping line | [optional] 
**LineTrackingStoppedReason** | Pointer to **string** | The reason Terminal49 stopped checking | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


