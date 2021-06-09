/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.10
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// CreateApiAccessRuleRequest struct for CreateApiAccessRuleRequest
type CreateApiAccessRuleRequest struct {
	//  One or more IDs of Client Certificate Authorities (CAs).
	CaIds []string `json:"CaIds,omitempty"`
	// One or more Client Certificate Common Names (CNs). If this parameter is specified, you must also specify the `CaIds` parameter.
	Cns []string `json:"Cns,omitempty"`
	// A description for the API access rule.
	Description string `json:"Description,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// One or more IP ranges, in CIDR notation (for example, 192.0.2.0/16).
	IpRanges []string `json:"IpRanges,omitempty"`
}
