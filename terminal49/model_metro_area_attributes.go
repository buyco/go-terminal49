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
// MetroAreaAttributes struct for MetroAreaAttributes
type MetroAreaAttributes struct {
	Name string `json:"name,omitempty"`
	// 2 digit country code
	CountryCode string `json:"country_code,omitempty"`
	// IANA tz
	Timezone string `json:"timezone,omitempty"`
	// UN/LOCODE
	Code string `json:"code,omitempty"`
}
