/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 0.17
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc
// SecurityGroupRule Information about the security group rule.
type SecurityGroupRule struct {
	// The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.
	FromPortRange int32 `json:"FromPortRange,omitempty"`
	// The IP protocol name (`tcp`, `udp`, `icmp`) or protocol number. By default, `-1`, which means all protocols.
	IpProtocol string `json:"IpProtocol,omitempty"`
	// One or more IP ranges for the security group rules, in CIDR notation (for example, 10.0.0.0/16).
	IpRanges []string `json:"IpRanges,omitempty"`
	// Information about one or more members of a security group.
	SecurityGroupsMembers []SecurityGroupsMember `json:"SecurityGroupsMembers,omitempty"`
	// One or more service IDs to allow traffic from a Net to access the corresponding 3DS OUTSCALE services. For more information, see [ReadNetAccessPointServices](#readnetaccesspointservices).
	ServiceIds []string `json:"ServiceIds,omitempty"`
	// The end of the port range for the TCP and UDP protocols, or an ICMP type number.
	ToPortRange int32 `json:"ToPortRange,omitempty"`
}
