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
// CreateNetRequest struct for CreateNetRequest
type CreateNetRequest struct {
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// The IP range for the Net, in CIDR notation (for example, 10.0.0.0/16).
	IpRange string `json:"IpRange"`
	// The tenancy options for the VMs (`default` if a VM created in a Net can be launched with any tenancy, `dedicated` if it can be launched with dedicated tenancy VMs running on single-tenant hardware).
	Tenancy string `json:"Tenancy,omitempty"`
}
