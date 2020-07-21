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
// Service Information about the service.
type Service struct {
	// The list of network prefixes used by the service, in CIDR notation.
	IpRanges []string `json:"IpRanges,omitempty"`
	// The ID of the service.
	ServiceId string `json:"ServiceId,omitempty"`
	// The name of the prefix list, which identifies the 3DS OUTSCALE service it is associated with.
	ServiceName string `json:"ServiceName,omitempty"`
}
