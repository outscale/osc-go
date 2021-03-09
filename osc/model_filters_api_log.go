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

// FiltersApiLog One or more filters.
type FiltersApiLog struct {
	// One or more API keys used for the query.
	QueryAccessKeys []string `json:"QueryAccessKeys,omitempty"`
	// The name of one or more API services used for the query.
	QueryApiNames []string `json:"QueryApiNames,omitempty"`
	// The name of one or more calls.
	QueryCallNames []string `json:"QueryCallNames,omitempty"`
	// The logs of the queries made after the date you specify, in ISO 8601 format (for example, `2017-06-14`).
	QueryDateAfter string `json:"QueryDateAfter,omitempty"`
	// The logs of the queries made before the date you specify, in ISO 8601 format (for example, `2017-06-14`).
	QueryDateBefore string `json:"QueryDateBefore,omitempty"`
	// One or more IP addresses used for the query.
	QueryIpAddresses []string `json:"QueryIpAddresses,omitempty"`
	// One or more user agents used for the HTTP request.
	QueryUserAgents []string `json:"QueryUserAgents,omitempty"`
	// One or more request IDs.
	RequestIds []string `json:"RequestIds,omitempty"`
	// One or more HTTP codes provided by the responses.
	ResponseStatusCodes []int32 `json:"ResponseStatusCodes,omitempty"`
}
