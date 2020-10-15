/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.4
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc
// FiltersClientGateway One or more filters.
type FiltersClientGateway struct {
	// The Border Gateway Protocol (BGP) Autonomous System Numbers (ASNs) of the connections.
	BgpAsns []int32 `json:"BgpAsns,omitempty"`
	// The IDs of the client gateways.
	ClientGatewayIds []string `json:"ClientGatewayIds,omitempty"`
	// The types of communication tunnels used by the client gateways (only `ipsec.1` is supported).
	ConnectionTypes []string `json:"ConnectionTypes,omitempty"`
	// The public IPv4 addresses of the client gateways.
	PublicIps []string `json:"PublicIps,omitempty"`
	// The states of the client gateways (`pending` \\| `available` \\| `deleting` \\| `deleted`).
	States []string `json:"States,omitempty"`
	// The keys of the tags associated with the client gateways.
	TagKeys []string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the client gateways.
	TagValues []string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the client gateways, in the following format: \"Filters\":{\"Tags\":[\"TAGKEY=TAGVALUE\"]}.
	Tags []string `json:"Tags,omitempty"`
}
