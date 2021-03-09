/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.7
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// Route Information about the route.
type Route struct {
	// The method used to create the route.
	CreationMethod string `json:"CreationMethod,omitempty"`
	// The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).
	DestinationIpRange string `json:"DestinationIpRange,omitempty"`
	// The ID of the OUTSCALE service.
	DestinationServiceId string `json:"DestinationServiceId,omitempty"`
	// The ID of the Internet service or virtual gateway attached to the Net.
	GatewayId string `json:"GatewayId,omitempty"`
	// The ID of a NAT service attached to the Net.
	NatServiceId string `json:"NatServiceId,omitempty"`
	// The ID of the Net access point.
	NetAccessPointId string `json:"NetAccessPointId,omitempty"`
	// The ID of the Net peering connection.
	NetPeeringId string `json:"NetPeeringId,omitempty"`
	// The ID of the NIC.
	NicId string `json:"NicId,omitempty"`
	// The state of a route in the route table (`active` \\| `blackhole`). The `blackhole` state indicates that the target of the route is not available.
	State string `json:"State,omitempty"`
	// The account ID of the owner of the VM.
	VmAccountId string `json:"VmAccountId,omitempty"`
	// The ID of a VM specified in a route in the table.
	VmId string `json:"VmId,omitempty"`
}
