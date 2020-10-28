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
// InlineResponse200 struct for InlineResponse200
type InlineResponse200 struct {
	Data []Shipment `json:"data,omitempty"`
	Included []AnyOfcontainerportterminal `json:"included,omitempty"`
	Links Links `json:"links,omitempty"`
	Meta Meta `json:"meta,omitempty"`
}
