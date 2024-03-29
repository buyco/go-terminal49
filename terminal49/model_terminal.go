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
// Terminal struct for Terminal
type Terminal struct {
	Id string `json:"id,omitempty"`
	Relationships TerminalRelationships `json:"relationships"`
	Attributes TerminalAttributes `json:"attributes"`
	Type string `json:"type,omitempty"`
}
