/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.8
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// UpdateRouteRequest struct for UpdateRouteRequest
type UpdateRouteRequest struct {
	// The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).
	DestinationIpRange string `json:"DestinationIpRange"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// The ID of an Internet service or virtual gateway attached to your Net.
	GatewayId string `json:"GatewayId,omitempty"`
	// The ID of a NAT service.
	NatServiceId string `json:"NatServiceId,omitempty"`
	// The ID of a Net peering connection.
	NetPeeringId string `json:"NetPeeringId,omitempty"`
	// The ID of a network interface card (NIC).
	NicId string `json:"NicId,omitempty"`
	// The ID of the route table.
	RouteTableId string `json:"RouteTableId"`
	// The ID of a NAT VM in your Net.
	VmId string `json:"VmId,omitempty"`
}
