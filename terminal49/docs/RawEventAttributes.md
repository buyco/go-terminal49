# RawEventAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** | Normalized string representing the event | [optional] 
**OriginalEvent** | **string** | The event name as returned by the carrier | [optional] 
**ActualOn** | **string** | The date of the event at the event location when no time information is available.  | [optional] 
**EstimatedOn** | **string** | The estimated date of the event at the event location when no time information is available.  | [optional] 
**ActualAt** | [**time.Time**](time.Time.md) | The datetime the event transpired in UTC | [optional] 
**EstimatedAt** | [**time.Time**](time.Time.md) | The estimated datetime the event will occur in UTC | [optional] 
**Timezone** | **string** | IANA tz where the event occured | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | When the raw_event was created in UTC | [optional] 
**LocationName** | **string** | The city or facility name of the event location | [optional] 
**LocationLocode** | **string** | UNLOCODE of the event location | [optional] 
**VesselName** | **string** | The name of the vessel where applicable | [optional] 
**VesselImo** | **string** | The IMO of the vessel where applicable | [optional] 
**Index** | **int32** | The order of the event. Thishelpful when only dates (i.e. actual_on) are available. | [optional] 
**VoyageNumber** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


