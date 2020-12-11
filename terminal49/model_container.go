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
// Container Represents the equipment during a specific journey.
type Container struct {
	Id string `json:"id"`
	Attributes ContainerAttributes `json:"attributes,omitempty"`
}
