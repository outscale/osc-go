/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.1
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc
// ConsumptionEntry Information about the resources consumed during the specified time period.
type ConsumptionEntry struct {
	// The category of the resource (for example, `network`).
	Category string `json:"Category,omitempty"`
	// The beginning of the time period.
	FromDate string `json:"FromDate,omitempty"`
	// The API call that triggered the resource consumption (for example, `RunInstances` or `CreateVolume`).
	Operation string `json:"Operation,omitempty"`
	// The service of the API call (`TinaOS-FCU`, `TinaOS-LBU`, `TinaOS-OSU` or `TinaOS-DirectLink`).
	Service string `json:"Service,omitempty"`
	// A description of the consumed resource.
	Title string `json:"Title,omitempty"`
	// The end of the time period.
	ToDate string `json:"ToDate,omitempty"`
	// The type of resource, depending on the API call.
	Type string `json:"Type,omitempty"`
	// The consumed amount for the resource. The unit depends on the resource type. For more information, see the `Title` element.
	Value string `json:"Value,omitempty"`
}
