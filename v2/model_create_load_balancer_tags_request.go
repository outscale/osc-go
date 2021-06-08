/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.8
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// CreateLoadBalancerTagsRequest struct for CreateLoadBalancerTagsRequest
type CreateLoadBalancerTagsRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// One or more load balancer names.
	LoadBalancerNames []string `json:"LoadBalancerNames"`
	// One or more tags to add to the specified load balancers.
	Tags []ResourceTag `json:"Tags"`
}

// NewCreateLoadBalancerTagsRequest instantiates a new CreateLoadBalancerTagsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLoadBalancerTagsRequest(loadBalancerNames []string, tags []ResourceTag) *CreateLoadBalancerTagsRequest {
	this := CreateLoadBalancerTagsRequest{}
	this.LoadBalancerNames = loadBalancerNames
	this.Tags = tags
	return &this
}

// NewCreateLoadBalancerTagsRequestWithDefaults instantiates a new CreateLoadBalancerTagsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLoadBalancerTagsRequestWithDefaults() *CreateLoadBalancerTagsRequest {
	this := CreateLoadBalancerTagsRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateLoadBalancerTagsRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerTagsRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateLoadBalancerTagsRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateLoadBalancerTagsRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetLoadBalancerNames returns the LoadBalancerNames field value
func (o *CreateLoadBalancerTagsRequest) GetLoadBalancerNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LoadBalancerNames
}

// GetLoadBalancerNamesOk returns a tuple with the LoadBalancerNames field value
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerTagsRequest) GetLoadBalancerNamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadBalancerNames, true
}

// SetLoadBalancerNames sets field value
func (o *CreateLoadBalancerTagsRequest) SetLoadBalancerNames(v []string) {
	o.LoadBalancerNames = v
}

// GetTags returns the Tags field value
func (o *CreateLoadBalancerTagsRequest) GetTags() []ResourceTag {
	if o == nil {
		var ret []ResourceTag
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *CreateLoadBalancerTagsRequest) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value
func (o *CreateLoadBalancerTagsRequest) SetTags(v []ResourceTag) {
	o.Tags = v
}

func (o CreateLoadBalancerTagsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["LoadBalancerNames"] = o.LoadBalancerNames
	}
	if true {
		toSerialize["Tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableCreateLoadBalancerTagsRequest struct {
	value *CreateLoadBalancerTagsRequest
	isSet bool
}

func (v NullableCreateLoadBalancerTagsRequest) Get() *CreateLoadBalancerTagsRequest {
	return v.value
}

func (v *NullableCreateLoadBalancerTagsRequest) Set(val *CreateLoadBalancerTagsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLoadBalancerTagsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLoadBalancerTagsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLoadBalancerTagsRequest(val *CreateLoadBalancerTagsRequest) *NullableCreateLoadBalancerTagsRequest {
	return &NullableCreateLoadBalancerTagsRequest{value: val, isSet: true}
}

func (v NullableCreateLoadBalancerTagsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLoadBalancerTagsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}