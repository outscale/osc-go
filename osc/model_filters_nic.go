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

// FiltersNic One or more filters.
type FiltersNic struct {
	// The descriptions of the NICs.
	Descriptions []string `json:"Descriptions,omitempty"`
	// Whether the source/destination checking is enabled (true) or disabled (false).
	IsSourceDestCheck bool `json:"IsSourceDestCheck,omitempty"`
	// Whether the NICs are deleted when the VMs they are attached to are terminated.
	LinkNicDeleteOnVmDeletion bool `json:"LinkNicDeleteOnVmDeletion,omitempty"`
	// The device numbers the NICs are attached to.
	LinkNicDeviceNumbers []int32 `json:"LinkNicDeviceNumbers,omitempty"`
	// The attachment IDs of the NICs.
	LinkNicLinkNicIds []string `json:"LinkNicLinkNicIds,omitempty"`
	// The states of the attachments.
	LinkNicStates []string `json:"LinkNicStates,omitempty"`
	// The account IDs of the owners of the VMs the NICs are attached to.
	LinkNicVmAccountIds []string `json:"LinkNicVmAccountIds,omitempty"`
	// The IDs of the VMs the NICs are attached to.
	LinkNicVmIds []string `json:"LinkNicVmIds,omitempty"`
	// The account IDs of the owners of the EIPs associated with the NICs.
	LinkPublicIpAccountIds []string `json:"LinkPublicIpAccountIds,omitempty"`
	// The association IDs returned when the EIPs were associated with the FNIs.
	LinkPublicIpLinkPublicIpIds []string `json:"LinkPublicIpLinkPublicIpIds,omitempty"`
	// The allocation IDs returned when the EIPs were allocated to their accounts.
	LinkPublicIpPublicIpIds []string `json:"LinkPublicIpPublicIpIds,omitempty"`
	// The EIPs associated with the NICs.
	LinkPublicIpPublicIps []string `json:"LinkPublicIpPublicIps,omitempty"`
	// The Media Access Control (MAC) addresses of the NICs.
	MacAddresses []string `json:"MacAddresses,omitempty"`
	// The IDs of the Nets where the NICs are located.
	NetIds []string `json:"NetIds,omitempty"`
	// The IDs of the NICs.
	NicIds []string `json:"NicIds,omitempty"`
	// The private DNS names associated with the primary private IP addresses.
	PrivateDnsNames []string `json:"PrivateDnsNames,omitempty"`
	// The account IDs of the owner of the EIPs associated with the private IP addresses.
	PrivateIpsLinkPublicIpAccountIds []string `json:"PrivateIpsLinkPublicIpAccountIds,omitempty"`
	// The EIPs associated with the private IP addresses.
	PrivateIpsLinkPublicIpPublicIps []string `json:"PrivateIpsLinkPublicIpPublicIps,omitempty"`
	// The primary private IP addresses of the NICs.
	PrivateIpsPrimaryIp bool `json:"PrivateIpsPrimaryIp,omitempty"`
	// The private IP addresses of the NICs.
	PrivateIpsPrivateIps []string `json:"PrivateIpsPrivateIps,omitempty"`
	// The IDs of the security groups associated with the NICs.
	SecurityGroupIds []string `json:"SecurityGroupIds,omitempty"`
	// The names of the security groups associated with the NICs.
	SecurityGroupNames []string `json:"SecurityGroupNames,omitempty"`
	// The states of the NICs.
	States []string `json:"States,omitempty"`
	// The IDs of the Subnets for the NICs.
	SubnetIds []string `json:"SubnetIds,omitempty"`
	// The Subregions where the NICs are located.
	SubregionNames []string `json:"SubregionNames,omitempty"`
	// The keys of the tags associated with the NICs.
	TagKeys []string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the NICs.
	TagValues []string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the NICs, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}.
	Tags []string `json:"Tags,omitempty"`
}
