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

// CreateLoadBalancerRequest struct for CreateLoadBalancerRequest
type CreateLoadBalancerRequest struct {
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// One or more listeners to create.
	Listeners []ListenerForCreation `json:"Listeners"`
	// The unique name of the load balancer (32 alphanumeric or hyphen characters maximum, but cannot start or end with a hyphen).
	LoadBalancerName string `json:"LoadBalancerName"`
	// The type of load balancer: `internet-facing` or `internal`. Use this parameter only for load balancers in a Net.
	LoadBalancerType *string `json:"LoadBalancerType,omitempty"`
	// (Net only) One or more IDs of security groups you want to assign to the load balancer. If not specified, the default security group of the Net is assigned to the load balancer.
	SecurityGroups *[]string `json:"SecurityGroups,omitempty"`
	// One or more IDs of Subnets in your Net that you want to attach to the load balancer.
	Subnets *[]string `json:"Subnets,omitempty"`
	// One or more names of Subregions (currently, only one Subregion is supported). This parameter is not required if you create a load balancer in a Net. To create an internal load balancer, use the `LoadBalancerType` parameter.
	SubregionNames *[]string `json:"SubregionNames,omitempty"`
	// One or more tags assigned to the load balancer.
	Tags *[]ResourceTag `json:"Tags,omitempty"`
}

// NewCreateLoadBalancerRequest instantiates a new CreateLoadBalancerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLoadBalancerRequest(listeners []ListenerForCreation, loadBalancerName string, ) *CreateLoadBalancerRequest {
	this := CreateLoadBalancerRequest{}
	this.Listeners = listeners
	this.LoadBalancerName = loadBalancerName
	return &this
}

// NewCreateLoadBalancerRequestWithDefaults instantiates a new CreateLoadBalancerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLoadBalancerRequestWithDefaults() *CreateLoadBalancerRequest {
	this := CreateLoadBalancerRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateLoadBalancerRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateLoadBalancerRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateLoadBalancerRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetListeners returns the Listeners field value
func (o *CreateLoadBalancerRequest) GetListeners() []ListenerForCreation {
	if o == nil  {
		var ret []ListenerForCreation
		return ret
	}

	return o.Listeners
}

// GetListenersOk returns a tuple with the Listeners field value
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerRequest) GetListenersOk() (*[]ListenerForCreation, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Listeners, true
}

// SetListeners sets field value
func (o *CreateLoadBalancerRequest) SetListeners(v []ListenerForCreation) {
	o.Listeners = v
}

// GetLoadBalancerName returns the LoadBalancerName field value
func (o *CreateLoadBalancerRequest) GetLoadBalancerName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.LoadBalancerName
}

// GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field value
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerRequest) GetLoadBalancerNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LoadBalancerName, true
}

// SetLoadBalancerName sets field value
func (o *CreateLoadBalancerRequest) SetLoadBalancerName(v string) {
	o.LoadBalancerName = v
}

// GetLoadBalancerType returns the LoadBalancerType field value if set, zero value otherwise.
func (o *CreateLoadBalancerRequest) GetLoadBalancerType() string {
	if o == nil || o.LoadBalancerType == nil {
		var ret string
		return ret
	}
	return *o.LoadBalancerType
}

// GetLoadBalancerTypeOk returns a tuple with the LoadBalancerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerRequest) GetLoadBalancerTypeOk() (*string, bool) {
	if o == nil || o.LoadBalancerType == nil {
		return nil, false
	}
	return o.LoadBalancerType, true
}

// HasLoadBalancerType returns a boolean if a field has been set.
func (o *CreateLoadBalancerRequest) HasLoadBalancerType() bool {
	if o != nil && o.LoadBalancerType != nil {
		return true
	}

	return false
}

// SetLoadBalancerType gets a reference to the given string and assigns it to the LoadBalancerType field.
func (o *CreateLoadBalancerRequest) SetLoadBalancerType(v string) {
	o.LoadBalancerType = &v
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise.
func (o *CreateLoadBalancerRequest) GetSecurityGroups() []string {
	if o == nil || o.SecurityGroups == nil {
		var ret []string
		return ret
	}
	return *o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerRequest) GetSecurityGroupsOk() (*[]string, bool) {
	if o == nil || o.SecurityGroups == nil {
		return nil, false
	}
	return o.SecurityGroups, true
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *CreateLoadBalancerRequest) HasSecurityGroups() bool {
	if o != nil && o.SecurityGroups != nil {
		return true
	}

	return false
}

// SetSecurityGroups gets a reference to the given []string and assigns it to the SecurityGroups field.
func (o *CreateLoadBalancerRequest) SetSecurityGroups(v []string) {
	o.SecurityGroups = &v
}

// GetSubnets returns the Subnets field value if set, zero value otherwise.
func (o *CreateLoadBalancerRequest) GetSubnets() []string {
	if o == nil || o.Subnets == nil {
		var ret []string
		return ret
	}
	return *o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerRequest) GetSubnetsOk() (*[]string, bool) {
	if o == nil || o.Subnets == nil {
		return nil, false
	}
	return o.Subnets, true
}

// HasSubnets returns a boolean if a field has been set.
func (o *CreateLoadBalancerRequest) HasSubnets() bool {
	if o != nil && o.Subnets != nil {
		return true
	}

	return false
}

// SetSubnets gets a reference to the given []string and assigns it to the Subnets field.
func (o *CreateLoadBalancerRequest) SetSubnets(v []string) {
	o.Subnets = &v
}

// GetSubregionNames returns the SubregionNames field value if set, zero value otherwise.
func (o *CreateLoadBalancerRequest) GetSubregionNames() []string {
	if o == nil || o.SubregionNames == nil {
		var ret []string
		return ret
	}
	return *o.SubregionNames
}

// GetSubregionNamesOk returns a tuple with the SubregionNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerRequest) GetSubregionNamesOk() (*[]string, bool) {
	if o == nil || o.SubregionNames == nil {
		return nil, false
	}
	return o.SubregionNames, true
}

// HasSubregionNames returns a boolean if a field has been set.
func (o *CreateLoadBalancerRequest) HasSubregionNames() bool {
	if o != nil && o.SubregionNames != nil {
		return true
	}

	return false
}

// SetSubregionNames gets a reference to the given []string and assigns it to the SubregionNames field.
func (o *CreateLoadBalancerRequest) SetSubregionNames(v []string) {
	o.SubregionNames = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateLoadBalancerRequest) GetTags() []ResourceTag {
	if o == nil || o.Tags == nil {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerRequest) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateLoadBalancerRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *CreateLoadBalancerRequest) SetTags(v []ResourceTag) {
	o.Tags = &v
}

func (o CreateLoadBalancerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["Listeners"] = o.Listeners
	}
	if true {
		toSerialize["LoadBalancerName"] = o.LoadBalancerName
	}
	if o.LoadBalancerType != nil {
		toSerialize["LoadBalancerType"] = o.LoadBalancerType
	}
	if o.SecurityGroups != nil {
		toSerialize["SecurityGroups"] = o.SecurityGroups
	}
	if o.Subnets != nil {
		toSerialize["Subnets"] = o.Subnets
	}
	if o.SubregionNames != nil {
		toSerialize["SubregionNames"] = o.SubregionNames
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableCreateLoadBalancerRequest struct {
	value *CreateLoadBalancerRequest
	isSet bool
}

func (v NullableCreateLoadBalancerRequest) Get() *CreateLoadBalancerRequest {
	return v.value
}

func (v *NullableCreateLoadBalancerRequest) Set(val *CreateLoadBalancerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLoadBalancerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLoadBalancerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLoadBalancerRequest(val *CreateLoadBalancerRequest) *NullableCreateLoadBalancerRequest {
	return &NullableCreateLoadBalancerRequest{value: val, isSet: true}
}

func (v NullableCreateLoadBalancerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLoadBalancerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


