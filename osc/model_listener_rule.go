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

// ListenerRule Information about the listener rule.
type ListenerRule struct {
	// The type of action for the rule (always `forward`).
	Action string `json:"Action,omitempty"`
	// A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?].
	HostNamePattern string `json:"HostNamePattern,omitempty"`
	// The ID of the listener.
	ListenerId int32 `json:"ListenerId,omitempty"`
	// The ID of the listener rule.
	ListenerRuleId int32 `json:"ListenerRuleId,omitempty"`
	// A human-readable name for the listener rule.
	ListenerRuleName string `json:"ListenerRuleName,omitempty"`
	// A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~\"'@:+?].
	PathPattern string `json:"PathPattern,omitempty"`
	// The priority level of the listener rule, between `1` and `19999` both included. Each rule must have a unique priority level. Otherwise, an error is returned.
	Priority int32 `json:"Priority,omitempty"`
	// The IDs of the backend VMs.
	VmIds []string `json:"VmIds,omitempty"`
}
