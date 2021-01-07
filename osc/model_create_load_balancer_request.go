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

// CreateLoadBalancerRequest struct for CreateLoadBalancerRequest
type CreateLoadBalancerRequest struct {
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// One or more listeners to create.
	Listeners []ListenerForCreation `json:"Listeners"`
	// The unique name of the load balancer (32 alphanumeric or hyphen characters maximum, but cannot start or end with a hyphen).
	LoadBalancerName string `json:"LoadBalancerName"`
	// The type of load balancer: `internet-facing` or `internal`. Use this parameter only for load balancers in a Net.
	LoadBalancerType string `json:"LoadBalancerType,omitempty"`
	// (Net only) One or more IDs of security groups you want to assign to the load balancer. If not specified, the default security group of the Net is assigned to the load balancer.
	SecurityGroups []string `json:"SecurityGroups,omitempty"`
	// One or more IDs of Subnets in your Net that you want to attach to the load balancer.
	Subnets []string `json:"Subnets,omitempty"`
	// One or more names of Subregions (currently, only one Subregion is supported). This parameter is not required if you create a load balancer in a Net. To create an internal load balancer, use the `LoadBalancerType` parameter.
	SubregionNames []string `json:"SubregionNames,omitempty"`
	// One or more tags assigned to the load balancer.
	Tags []ResourceTag `json:"Tags,omitempty"`
}
