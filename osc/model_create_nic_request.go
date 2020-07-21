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
// CreateNicRequest struct for CreateNicRequest
type CreateNicRequest struct {
	// A description for the NIC.
	Description string `json:"Description,omitempty"`
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// The primary private IP address for the NIC.<br /><br />  This IP address must be within the IP address range of the Subnet that you specify with the `SubnetId` attribute.<br /> If you do not specify this attribute, a random private IP address is selected within the IP address range of the Subnet.
	PrivateIps []PrivateIpLight `json:"PrivateIps,omitempty"`
	// One or more IDs of security groups for the NIC.
	SecurityGroupIds []string `json:"SecurityGroupIds,omitempty"`
	// The ID of the Subnet in which you want to create the NIC.
	SubnetId string `json:"SubnetId"`
}
