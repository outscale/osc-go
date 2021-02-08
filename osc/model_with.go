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

// With The information to be displayed in the API logs to retrieve.
type With struct {
	// If set to true, the account ID is displayed in the logs.
	AccountId bool `json:"AccountId,omitempty"`
	// If set to true, the duration of the call is displayed each log.
	CallDuration bool `json:"CallDuration,omitempty"`
	// If set to true, the API key used for the query is displayed each log.
	QueryAccessKey bool `json:"QueryAccessKey,omitempty"`
	// If set to true, the name of the API service used by the call is displayed in each log (`oapi` \\| `fcu` \\| `lbu` \\| `directlink` \\| `eim` \\| `icu`).
	QueryApiName bool `json:"QueryApiName,omitempty"`
	// If set to true, the version of the API service used by the call is displayed in each log.
	QueryApiVersion bool `json:"QueryApiVersion,omitempty"`
	// If set to true, the name of the call is displayed in each log.
	QueryCallName bool `json:"QueryCallName,omitempty"`
	// If set to true, the date of the call is displayed in each log.
	QueryDate bool `json:"QueryDate,omitempty"`
	// If set to true, the query header RAW is displayed in each log.
	QueryHeaderRaw bool `json:"QueryHeaderRaw,omitempty"`
	// If set to true, the query header size is displayed in each log.
	QueryHeaderSize bool `json:"QueryHeaderSize,omitempty"`
	// If set to true, the IP address used to make to query is displayed in each log.
	QueryIpAddress bool `json:"QueryIpAddress,omitempty"`
	// If set to true, the query payload raw is displayed in each log.
	QueryPayloadRaw bool `json:"QueryPayloadRaw,omitempty"`
	// If set to true, the query payload size is displayed in each log.
	QueryPayloadSize bool `json:"QueryPayloadSize,omitempty"`
	// If set to true, the user agent used to make the HTTP request is displayed in each log.
	QueryUserAgent bool `json:"QueryUserAgent,omitempty"`
	// By default ot if set to true, the ID of the call is displayed in each log.
	RequestId bool `json:"RequestId,omitempty"`
	// If set to true, the size of the response (in bytes) is displayed in each log.
	ResponseSize bool `json:"ResponseSize,omitempty"`
	// If set to true, the HTTP code provided by the response is displayed in each log.
	ResponseStatusCode bool `json:"ResponseStatusCode,omitempty"`
}
