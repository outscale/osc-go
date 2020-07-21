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
// LinkPrivateIpsRequest struct for LinkPrivateIpsRequest
type LinkPrivateIpsRequest struct {
	// If `true`, allows an IP address that is already assigned to another NIC in the same Subnet to be assigned to the NIC you specified.
	AllowRelink bool `json:"AllowRelink,omitempty"`
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// The ID of the NIC.
	NicId string `json:"NicId"`
	// The secondary private IP address or addresses you want to assign to the NIC within the IP address range of the Subnet.
	PrivateIps []string `json:"PrivateIps,omitempty"`
	// The number of secondary private IP addresses to assign to the NIC.
	SecondaryPrivateIpCount int32 `json:"SecondaryPrivateIpCount,omitempty"`
}
