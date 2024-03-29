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
// ModelError struct for ModelError
type ModelError struct {
	Detail string `json:"detail,omitempty"`
	Title string `json:"title"`
	Source ErrorSource `json:"source,omitempty"`
	Code string `json:"code,omitempty"`
	Status string `json:"status,omitempty"`
}
