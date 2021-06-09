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

// UpdateVmRequest struct for UpdateVmRequest
type UpdateVmRequest struct {
	// One or more block device mappings of the VM.
	BlockDeviceMappings []BlockDeviceMappingVmUpdate `json:"BlockDeviceMappings,omitempty"`
	// If true, the VM is optimized for BSU I/O.
	BsuOptimized bool `json:"BsuOptimized,omitempty"`
	// If true, you cannot terminate the VM using Cockpit, the CLI or the API. If false, you can.
	DeletionProtection bool `json:"DeletionProtection,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// (Net only) If true, the source/destination check is enabled. If false, it is disabled. This value must be false for a NAT VM to perform network address translation (NAT) in a Net.
	IsSourceDestChecked bool `json:"IsSourceDestChecked,omitempty"`
	// The name of the keypair.<br /> To complete the replacement, manually replace the old public key with the new public key in the ~/.ssh/authorized_keys file located in the VM. Restart the VM to apply the change.
	KeypairName string `json:"KeypairName,omitempty"`
	// The performance of the VM (`medium` \\| `high` \\|  `highest`).
	Performance string `json:"Performance,omitempty"`
	// One or more IDs of security groups for the VM.
	SecurityGroupIds []string `json:"SecurityGroupIds,omitempty"`
	// The Base64-encoded MIME user data.
	UserData string `json:"UserData,omitempty"`
	// The ID of the VM.
	VmId string `json:"VmId"`
	// The VM behavior when you stop it. By default or if set to `stop`, the VM stops. If set to `restart`, the VM stops then automatically restarts. If set to `terminate`, the VM stops and is terminated.
	VmInitiatedShutdownBehavior string `json:"VmInitiatedShutdownBehavior,omitempty"`
	// The type of VM. For more information, see [Instance Types](https://wiki.outscale.net/display/EN/Instance+Types).
	VmType string `json:"VmType,omitempty"`
}
