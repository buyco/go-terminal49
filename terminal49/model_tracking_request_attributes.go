/*
 * Terminal49 API Reference
 *
 * The Terminal 49 API offers a convenient way to programmatically track your shipments from origin to destination.
 *
 * API version: 0.2.0
 * Contact: support@terminal49.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package terminal49
import (
	"time"
)
// TrackingRequestAttributes struct for TrackingRequestAttributes
type TrackingRequestAttributes struct {
	RequestNumber string `json:"request_number"`
	RefNumbers []string `json:"ref_numbers,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Status string `json:"status"`
	// If the tracking request has failed, or is currently failing, the last reason we were unable to complete the request
	FailedReason *string `json:"failed_reason,omitempty"`
	RequestType string `json:"request_type"`
	Scac string `json:"scac"`
	CreatedAt time.Time `json:"created_at"`
	IsRetrying bool `json:"is_retrying,omitempty"`
	// How many times T49 has attempted to get the shipment from the shipping line
	RetryCount int32 `json:"retry_count,omitempty"`
}
