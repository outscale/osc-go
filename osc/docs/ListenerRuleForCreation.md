# ListenerRuleForCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The type of action for the rule (always &#x60;forward&#x60;). | [optional] 
**HostNamePattern** | **string** | A host-name pattern for the rule, with a maximum length of 128 characters. This host-name pattern supports maximum three wildcards, and must not contain any special characters except [-.?].  | [optional] 
**ListenerRuleName** | **string** | A human-readable name for the listener rule. | 
**PathPattern** | **string** | A path pattern for the rule, with a maximum length of 128 characters. This path pattern supports maximum three wildcards, and must not contain any special characters except [_-.$/~&amp;quot;&#39;@:+?]. | [optional] 
**Priority** | **int32** | The priority level of the listener rule, between &#x60;1&#x60; and &#x60;19999&#x60; both included. Each rule must have a unique priority level. Otherwise, an error is returned. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


