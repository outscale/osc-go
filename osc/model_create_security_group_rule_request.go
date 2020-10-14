/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.4
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateSecurityGroupRuleRequest struct for CreateSecurityGroupRuleRequest
type CreateSecurityGroupRuleRequest struct {
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The direction of the flow: `Inbound` or `Outbound`. You can specify `Outbound` for Nets only.
	Flow string `json:"Flow"`
	// The beginning of the port range for the TCP and UDP protocols, or an ICMP type number.
	FromPortRange *int32 `json:"FromPortRange,omitempty"`
	// The IP protocol name (`tcp`, `udp`, `icmp`) or protocol number. By default, `-1`, which means all protocols.
	IpProtocol *string `json:"IpProtocol,omitempty"`
	// The IP range for the security group rule, in CIDR notation (for example, 10.0.0.0/16).
	IpRange *string `json:"IpRange,omitempty"`
	// Information about the security group rule to create.
	Rules *[]SecurityGroupRule `json:"Rules,omitempty"`
	// The account ID of the owner of the security group for which you want to create a rule.
	SecurityGroupAccountIdToLink *string `json:"SecurityGroupAccountIdToLink,omitempty"`
	// The ID of the security group for which you want to create a rule.
	SecurityGroupId string `json:"SecurityGroupId"`
	// The ID of the source security group. If you are in the Public Cloud, you can also specify the name of the source security group.
	SecurityGroupNameToLink *string `json:"SecurityGroupNameToLink,omitempty"`
	// The end of the port range for the TCP and UDP protocols, or an ICMP type number.
	ToPortRange *int32 `json:"ToPortRange,omitempty"`
}

// NewCreateSecurityGroupRuleRequest instantiates a new CreateSecurityGroupRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSecurityGroupRuleRequest(flow string, securityGroupId string, ) *CreateSecurityGroupRuleRequest {
	this := CreateSecurityGroupRuleRequest{}
	this.Flow = flow
	this.SecurityGroupId = securityGroupId
	return &this
}

// NewCreateSecurityGroupRuleRequestWithDefaults instantiates a new CreateSecurityGroupRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSecurityGroupRuleRequestWithDefaults() *CreateSecurityGroupRuleRequest {
	this := CreateSecurityGroupRuleRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateSecurityGroupRuleRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRuleRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateSecurityGroupRuleRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateSecurityGroupRuleRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFlow returns the Flow field value
func (o *CreateSecurityGroupRuleRequest) GetFlow() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Flow
}

// GetFlowOk returns a tuple with the Flow field value
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRuleRequest) GetFlowOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Flow, true
}

// SetFlow sets field value
func (o *CreateSecurityGroupRuleRequest) SetFlow(v string) {
	o.Flow = v
}

// GetFromPortRange returns the FromPortRange field value if set, zero value otherwise.
func (o *CreateSecurityGroupRuleRequest) GetFromPortRange() int32 {
	if o == nil || o.FromPortRange == nil {
		var ret int32
		return ret
	}
	return *o.FromPortRange
}

// GetFromPortRangeOk returns a tuple with the FromPortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRuleRequest) GetFromPortRangeOk() (*int32, bool) {
	if o == nil || o.FromPortRange == nil {
		return nil, false
	}
	return o.FromPortRange, true
}

// HasFromPortRange returns a boolean if a field has been set.
func (o *CreateSecurityGroupRuleRequest) HasFromPortRange() bool {
	if o != nil && o.FromPortRange != nil {
		return true
	}

	return false
}

// SetFromPortRange gets a reference to the given int32 and assigns it to the FromPortRange field.
func (o *CreateSecurityGroupRuleRequest) SetFromPortRange(v int32) {
	o.FromPortRange = &v
}

// GetIpProtocol returns the IpProtocol field value if set, zero value otherwise.
func (o *CreateSecurityGroupRuleRequest) GetIpProtocol() string {
	if o == nil || o.IpProtocol == nil {
		var ret string
		return ret
	}
	return *o.IpProtocol
}

// GetIpProtocolOk returns a tuple with the IpProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRuleRequest) GetIpProtocolOk() (*string, bool) {
	if o == nil || o.IpProtocol == nil {
		return nil, false
	}
	return o.IpProtocol, true
}

// HasIpProtocol returns a boolean if a field has been set.
func (o *CreateSecurityGroupRuleRequest) HasIpProtocol() bool {
	if o != nil && o.IpProtocol != nil {
		return true
	}

	return false
}

// SetIpProtocol gets a reference to the given string and assigns it to the IpProtocol field.
func (o *CreateSecurityGroupRuleRequest) SetIpProtocol(v string) {
	o.IpProtocol = &v
}

// GetIpRange returns the IpRange field value if set, zero value otherwise.
func (o *CreateSecurityGroupRuleRequest) GetIpRange() string {
	if o == nil || o.IpRange == nil {
		var ret string
		return ret
	}
	return *o.IpRange
}

// GetIpRangeOk returns a tuple with the IpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRuleRequest) GetIpRangeOk() (*string, bool) {
	if o == nil || o.IpRange == nil {
		return nil, false
	}
	return o.IpRange, true
}

// HasIpRange returns a boolean if a field has been set.
func (o *CreateSecurityGroupRuleRequest) HasIpRange() bool {
	if o != nil && o.IpRange != nil {
		return true
	}

	return false
}

// SetIpRange gets a reference to the given string and assigns it to the IpRange field.
func (o *CreateSecurityGroupRuleRequest) SetIpRange(v string) {
	o.IpRange = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *CreateSecurityGroupRuleRequest) GetRules() []SecurityGroupRule {
	if o == nil || o.Rules == nil {
		var ret []SecurityGroupRule
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRuleRequest) GetRulesOk() (*[]SecurityGroupRule, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *CreateSecurityGroupRuleRequest) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given []SecurityGroupRule and assigns it to the Rules field.
func (o *CreateSecurityGroupRuleRequest) SetRules(v []SecurityGroupRule) {
	o.Rules = &v
}

// GetSecurityGroupAccountIdToLink returns the SecurityGroupAccountIdToLink field value if set, zero value otherwise.
func (o *CreateSecurityGroupRuleRequest) GetSecurityGroupAccountIdToLink() string {
	if o == nil || o.SecurityGroupAccountIdToLink == nil {
		var ret string
		return ret
	}
	return *o.SecurityGroupAccountIdToLink
}

// GetSecurityGroupAccountIdToLinkOk returns a tuple with the SecurityGroupAccountIdToLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRuleRequest) GetSecurityGroupAccountIdToLinkOk() (*string, bool) {
	if o == nil || o.SecurityGroupAccountIdToLink == nil {
		return nil, false
	}
	return o.SecurityGroupAccountIdToLink, true
}

// HasSecurityGroupAccountIdToLink returns a boolean if a field has been set.
func (o *CreateSecurityGroupRuleRequest) HasSecurityGroupAccountIdToLink() bool {
	if o != nil && o.SecurityGroupAccountIdToLink != nil {
		return true
	}

	return false
}

// SetSecurityGroupAccountIdToLink gets a reference to the given string and assigns it to the SecurityGroupAccountIdToLink field.
func (o *CreateSecurityGroupRuleRequest) SetSecurityGroupAccountIdToLink(v string) {
	o.SecurityGroupAccountIdToLink = &v
}

// GetSecurityGroupId returns the SecurityGroupId field value
func (o *CreateSecurityGroupRuleRequest) GetSecurityGroupId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SecurityGroupId
}

// GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field value
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRuleRequest) GetSecurityGroupIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SecurityGroupId, true
}

// SetSecurityGroupId sets field value
func (o *CreateSecurityGroupRuleRequest) SetSecurityGroupId(v string) {
	o.SecurityGroupId = v
}

// GetSecurityGroupNameToLink returns the SecurityGroupNameToLink field value if set, zero value otherwise.
func (o *CreateSecurityGroupRuleRequest) GetSecurityGroupNameToLink() string {
	if o == nil || o.SecurityGroupNameToLink == nil {
		var ret string
		return ret
	}
	return *o.SecurityGroupNameToLink
}

// GetSecurityGroupNameToLinkOk returns a tuple with the SecurityGroupNameToLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRuleRequest) GetSecurityGroupNameToLinkOk() (*string, bool) {
	if o == nil || o.SecurityGroupNameToLink == nil {
		return nil, false
	}
	return o.SecurityGroupNameToLink, true
}

// HasSecurityGroupNameToLink returns a boolean if a field has been set.
func (o *CreateSecurityGroupRuleRequest) HasSecurityGroupNameToLink() bool {
	if o != nil && o.SecurityGroupNameToLink != nil {
		return true
	}

	return false
}

// SetSecurityGroupNameToLink gets a reference to the given string and assigns it to the SecurityGroupNameToLink field.
func (o *CreateSecurityGroupRuleRequest) SetSecurityGroupNameToLink(v string) {
	o.SecurityGroupNameToLink = &v
}

// GetToPortRange returns the ToPortRange field value if set, zero value otherwise.
func (o *CreateSecurityGroupRuleRequest) GetToPortRange() int32 {
	if o == nil || o.ToPortRange == nil {
		var ret int32
		return ret
	}
	return *o.ToPortRange
}

// GetToPortRangeOk returns a tuple with the ToPortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSecurityGroupRuleRequest) GetToPortRangeOk() (*int32, bool) {
	if o == nil || o.ToPortRange == nil {
		return nil, false
	}
	return o.ToPortRange, true
}

// HasToPortRange returns a boolean if a field has been set.
func (o *CreateSecurityGroupRuleRequest) HasToPortRange() bool {
	if o != nil && o.ToPortRange != nil {
		return true
	}

	return false
}

// SetToPortRange gets a reference to the given int32 and assigns it to the ToPortRange field.
func (o *CreateSecurityGroupRuleRequest) SetToPortRange(v int32) {
	o.ToPortRange = &v
}

func (o CreateSecurityGroupRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["Flow"] = o.Flow
	}
	if o.FromPortRange != nil {
		toSerialize["FromPortRange"] = o.FromPortRange
	}
	if o.IpProtocol != nil {
		toSerialize["IpProtocol"] = o.IpProtocol
	}
	if o.IpRange != nil {
		toSerialize["IpRange"] = o.IpRange
	}
	if o.Rules != nil {
		toSerialize["Rules"] = o.Rules
	}
	if o.SecurityGroupAccountIdToLink != nil {
		toSerialize["SecurityGroupAccountIdToLink"] = o.SecurityGroupAccountIdToLink
	}
	if true {
		toSerialize["SecurityGroupId"] = o.SecurityGroupId
	}
	if o.SecurityGroupNameToLink != nil {
		toSerialize["SecurityGroupNameToLink"] = o.SecurityGroupNameToLink
	}
	if o.ToPortRange != nil {
		toSerialize["ToPortRange"] = o.ToPortRange
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSecurityGroupRuleRequest struct {
	value *CreateSecurityGroupRuleRequest
	isSet bool
}

func (v NullableCreateSecurityGroupRuleRequest) Get() *CreateSecurityGroupRuleRequest {
	return v.value
}

func (v *NullableCreateSecurityGroupRuleRequest) Set(val *CreateSecurityGroupRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSecurityGroupRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSecurityGroupRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSecurityGroupRuleRequest(val *CreateSecurityGroupRuleRequest) *NullableCreateSecurityGroupRuleRequest {
	return &NullableCreateSecurityGroupRuleRequest{value: val, isSet: true}
}

func (v NullableCreateSecurityGroupRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSecurityGroupRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


