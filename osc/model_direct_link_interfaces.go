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
// DirectLinkInterfaces Information about the DirectLink interfaces.
type DirectLinkInterfaces struct {
	// The account ID of the owner of the DirectLink interface.
	AccountId string `json:"AccountId,omitempty"`
	// The BGP (Border Gateway Protocol) ASN (Autonomous System Number) on the customer's side of the DirectLink interface.
	BgpAsn int32 `json:"BgpAsn,omitempty"`
	// The BGP authentication key.
	BgpKey string `json:"BgpKey,omitempty"`
	// The IP address on the customer's side of the DirectLink interface.
	ClientPrivateIp string `json:"ClientPrivateIp,omitempty"`
	// The ID of the DirectLink.
	DirectLinkId string `json:"DirectLinkId,omitempty"`
	// The ID of the DirectLink interface.
	DirectLinkInterfaceId string `json:"DirectLinkInterfaceId,omitempty"`
	// The name of the DirectLink interface.
	DirectLinkInterfaceName string `json:"DirectLinkInterfaceName,omitempty"`
	// The type of the DirectLink interface (always `private`).
	InterfaceType string `json:"InterfaceType,omitempty"`
	// The datacenter where the DirectLink interface is located.
	Location string `json:"Location,omitempty"`
	// The IP address on 3DS OUTSCALE's side of the DirectLink interface.
	OutscalePrivateIp string `json:"OutscalePrivateIp,omitempty"`
	// The state of the DirectLink interface (`pending` \\| `available` \\| `deleting` \\| `deleted` \\| `confirming` \\| `rejected` \\| `expired`).
	State string `json:"State,omitempty"`
	// The ID of the target virtual gateway.
	VirtualGatewayId string `json:"VirtualGatewayId,omitempty"`
	// The VLAN number associated with the DirectLink interface.
	Vlan int32 `json:"Vlan,omitempty"`
}
