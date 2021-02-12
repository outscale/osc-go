/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.7
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// NatService Information about the NAT service.
type NatService struct {
	// The ID of the NAT service.
	NatServiceId string `json:"NatServiceId,omitempty"`
	// The ID of the Net in which the NAT service is.
	NetId string `json:"NetId,omitempty"`
	// Information about the External IP address or addresses (EIPs) associated with the NAT service.
	PublicIps []PublicIpLight `json:"PublicIps,omitempty"`
	// The state of the NAT service (`pending` \\| `available` \\| `deleting` \\| `deleted`).
	State string `json:"State,omitempty"`
	// The ID of the Subnet in which the NAT service is.
	SubnetId string `json:"SubnetId,omitempty"`
	// One or more tags associated with the NAT service.
	Tags []ResourceTag `json:"Tags,omitempty"`
}
