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

// LoadBalancer Information about the load balancer.
type LoadBalancer struct {
	AccessLog AccessLog `json:"AccessLog,omitempty"`
	// The stickiness policies defined for the load balancer.
	ApplicationStickyCookiePolicies []ApplicationStickyCookiePolicy `json:"ApplicationStickyCookiePolicies,omitempty"`
	// One or more IDs of back-end VMs for the load balancer.
	BackendVmIds []string `json:"BackendVmIds,omitempty"`
	// The DNS name of the load balancer.
	DnsName     string      `json:"DnsName,omitempty"`
	HealthCheck HealthCheck `json:"HealthCheck,omitempty"`
	// The listeners for the load balancer.
	Listeners []Listener `json:"Listeners,omitempty"`
	// The name of the load balancer.
	LoadBalancerName string `json:"LoadBalancerName,omitempty"`
	// The policies defined for the load balancer.
	LoadBalancerStickyCookiePolicies []LoadBalancerStickyCookiePolicy `json:"LoadBalancerStickyCookiePolicies,omitempty"`
	// The type of load balancer. Valid only for load balancers in a Net.<br /> If `LoadBalancerType` is `internet-facing`, the load balancer has a public DNS name that resolves to a public IP address.<br /> If `LoadBalancerType` is `internal`, the load balancer has a public DNS name that resolves to a private IP address.
	LoadBalancerType string `json:"LoadBalancerType,omitempty"`
	// The ID of the Net for the load balancer.
	NetId string `json:"NetId,omitempty"`
	// One or more IDs of security groups for the load balancers. Valid only for load balancers in a Net.
	SecurityGroups      []string            `json:"SecurityGroups,omitempty"`
	SourceSecurityGroup SourceSecurityGroup `json:"SourceSecurityGroup,omitempty"`
	// The IDs of the Subnets for the load balancer.
	Subnets []string `json:"Subnets,omitempty"`
	// One or more names of Subregions for the load balancer.
	SubregionNames []string `json:"SubregionNames,omitempty"`
	// One or more tags associated with the load balancer.
	Tags []ResourceTag `json:"Tags,omitempty"`
}
