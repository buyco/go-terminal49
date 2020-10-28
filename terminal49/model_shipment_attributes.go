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
// ShipmentAttributes struct for ShipmentAttributes
type ShipmentAttributes struct {
	BillOfLadingNumber string `json:"bill_of_lading_number"`
	RefNumbers []string `json:"ref_numbers,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	// `unreleased`
	Tags []string `json:"tags,omitempty"`
	// UN/LOCODE
	PortOfLadingLocode string `json:"port_of_lading_locode,omitempty"`
	PortOfLadingName string `json:"port_of_lading_name,omitempty"`
	// UN/LOCODE
	PortOfDischargeLocode string `json:"port_of_discharge_locode,omitempty"`
	PortOfDischargeName string `json:"port_of_discharge_name,omitempty"`
	// UN/LOCODE
	DestinationLocode string `json:"destination_locode,omitempty"`
	DestinationName string `json:"destination_name,omitempty"`
	ShippingLineScac string `json:"shipping_line_scac,omitempty"`
	PodVesselName string `json:"pod_vessel_name,omitempty"`
	PodVesselImo string `json:"pod_vessel_imo,omitempty"`
	PodVoyageNumber string `json:"pod_voyage_number,omitempty"`
	PolEtdAt time.Time `json:"pol_etd_at,omitempty"`
	PolAtdAt time.Time `json:"pol_atd_at,omitempty"`
	PodEtaAt time.Time `json:"pod_eta_at,omitempty"`
	PodAtaAt time.Time `json:"pod_ata_at,omitempty"`
	DestinationEtaAt time.Time `json:"destination_eta_at,omitempty"`
	DestinationAtaAt time.Time `json:"destination_ata_at,omitempty"`
	// IANA tz
	PolTimezone string `json:"pol_timezone,omitempty"`
	// IANA tz
	PodTimezone string `json:"pod_timezone,omitempty"`
	// IANA tz
	DestinationTimezone string `json:"destination_timezone,omitempty"`
}
