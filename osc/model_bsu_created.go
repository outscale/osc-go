/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.3
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc
// BsuCreated Information about the created BSU volume.
type BsuCreated struct {
	// Set to `true` by default, which means that the volume is deleted when the VM is terminated. If set to `false`, the volume is not deleted when the VM is terminated.
	DeleteOnVmDeletion bool `json:"DeleteOnVmDeletion,omitempty"`
	// The time and date of attachment of the volume to the VM.
	LinkDate string `json:"LinkDate,omitempty"`
	// The state of the volume.
	State string `json:"State,omitempty"`
	// The ID of the volume.
	VolumeId string `json:"VolumeId,omitempty"`
}