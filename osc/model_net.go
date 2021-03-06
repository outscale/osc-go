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

// Net Information about the Net.
type Net struct {
	// The ID of the DHCP options set (or `default` if you want to associate the default one).
	DhcpOptionsSetId string `json:"DhcpOptionsSetId,omitempty"`
	// The IP range for the Net, in CIDR notation (for example, 10.0.0.0/16).
	IpRange string `json:"IpRange,omitempty"`
	// The ID of the Net.
	NetId string `json:"NetId,omitempty"`
	// The state of the Net (`pending` \\| `available`).
	State string `json:"State,omitempty"`
	// One or more tags associated with the Net.
	Tags []ResourceTag `json:"Tags,omitempty"`
	// The VM tenancy in a Net.
	Tenancy string `json:"Tenancy,omitempty"`
}
