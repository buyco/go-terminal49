/*
 * Terminal49 API Reference
 *
 * The Terminal 49 API offers a convenient way to programmatically track your shipments from origin to destination.
 *
 * API version: 0.1.0
 * Contact: support@terminal49.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package terminal49
import (
	"time"
)
// ContainerAttributes struct for ContainerAttributes
type ContainerAttributes struct {
	Number string `json:"number,omitempty"`
	EquipmentType string `json:"equipment_type,omitempty"`
	EquipmentLength int32 `json:"equipment_length,omitempty"`
	EquipmentHeight string `json:"equipment_height,omitempty"`
	WeightInLbs string `json:"weight_in_lbs,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	SealNumber *string `json:"seal_number,omitempty"`
	// The last free day for pickup before deummurage accrues.
	PickupLfd *time.Time `json:"pickup_lfd,omitempty"`
	// When available the pickup appointment time at the terminal is returned.
	PickupApopintmentAt time.Time `json:"pickup_apopintment_at,omitempty"`
	// Whether Terminal 49 is receiving availability status from the terminal.
	AvailabilityKnown bool `json:"availability_known,omitempty"`
	// If availability_known is true, then whether container is available to be picked up at terminal.
	AvailableForPickup *bool `json:"available_for_pickup,omitempty"`
	// `unreleased`
	CurrentTransportationMode string `json:"current_transportation_mode,omitempty"`
	// `unreleased` Discharge time at the port of discharge
	PodDischargedAt *time.Time `json:"pod_discharged_at,omitempty"`
	// Full Out time at port of discharge. Null for inland moves.
	PodFullOutAt *time.Time `json:"pod_full_out_at,omitempty"`
	// `unreleased` Arrival or unloading time at final destination for inland moves.
	DestinationUnloadedAt *time.Time `json:"destination_unloaded_at,omitempty"`
	// `unreleased` Pickup time at final destination for inland moves.
	FinalDestinationFullOutAt *time.Time `json:"final_destination_full_out_at,omitempty"`
	// `unreleased` Time empty container was returned.
	EmptyTerminatedAt *time.Time `json:"empty_terminated_at,omitempty"`
	// `unreleased` IANA tz
	PodTimezone string `json:"pod_timezone,omitempty"`
	// `unreleaed` IANA tz
	DestinationTimezone string `json:"destination_timezone,omitempty"`
	// `unreleased` IANA tz
	EmptyReturnedTimezone *string `json:"empty_returned_timezone,omitempty"`
	HoldsAtPodTerminal []TerminalHold `json:"holds_at_pod_terminal,omitempty"`
	FeesAtPodTerminal []TerminalFee `json:"fees_at_pod_terminal,omitempty"`
}