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
// Links struct for Links
type Links struct {
	Last string `json:"last,omitempty"`
	Next string `json:"next,omitempty"`
	Prev string `json:"prev,omitempty"`
	First string `json:"first,omitempty"`
	Self string `json:"self,omitempty"`
}
