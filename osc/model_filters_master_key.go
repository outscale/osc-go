/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.6
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc
// FiltersMasterKey One or more filters.
type FiltersMasterKey struct {
	// The descriptions of the master keys.
	Descriptions []string `json:"Descriptions,omitempty"`
	// The IDs of the master keys.
	MasterKeyIds []string `json:"MasterKeyIds,omitempty"`
	// The states of the master keys (`Enabled` \\| `Disabled` \\| `PendingDeletion`).
	States []string `json:"States,omitempty"`
}
