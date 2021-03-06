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

// FiltersRouteTable One or more filters.
type FiltersRouteTable struct {
	// The IDs of the route tables involved in the associations.
	LinkRouteTableIds []string `json:"LinkRouteTableIds,omitempty"`
	// The IDs of the associations between the route tables and the Subnets.
	LinkRouteTableLinkRouteTableIds []string `json:"LinkRouteTableLinkRouteTableIds,omitempty"`
	// If true, the route tables are the main ones for their Nets.
	LinkRouteTableMain bool `json:"LinkRouteTableMain,omitempty"`
	// The IDs of the Subnets involved in the associations.
	LinkSubnetIds []string `json:"LinkSubnetIds,omitempty"`
	// The IDs of the Nets for the route tables.
	NetIds []string `json:"NetIds,omitempty"`
	// The methods used to create a route.
	RouteCreationMethods []string `json:"RouteCreationMethods,omitempty"`
	// The IP ranges specified in routes in the tables.
	RouteDestinationIpRanges []string `json:"RouteDestinationIpRanges,omitempty"`
	// The service IDs specified in routes in the tables.
	RouteDestinationServiceIds []string `json:"RouteDestinationServiceIds,omitempty"`
	// The IDs of the gateways specified in routes in the tables.
	RouteGatewayIds []string `json:"RouteGatewayIds,omitempty"`
	// The IDs of the NAT services specified in routes in the tables.
	RouteNatServiceIds []string `json:"RouteNatServiceIds,omitempty"`
	// The IDs of the Net peering connections specified in routes in the tables.
	RouteNetPeeringIds []string `json:"RouteNetPeeringIds,omitempty"`
	// The states of routes in the route tables (`active` \\| `blackhole`). The `blackhole` state indicates that the target of the route is not available.
	RouteStates []string `json:"RouteStates,omitempty"`
	// The IDs of the route tables.
	RouteTableIds []string `json:"RouteTableIds,omitempty"`
	// The IDs of the VMs specified in routes in the tables.
	RouteVmIds []string `json:"RouteVmIds,omitempty"`
	// The keys of the tags associated with the route tables.
	TagKeys []string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the route tables.
	TagValues []string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the route tables, in the following format: \"Filters\":{\"Tags\":[\"TAGKEY=TAGVALUE\"]}.
	Tags []string `json:"Tags,omitempty"`
}
