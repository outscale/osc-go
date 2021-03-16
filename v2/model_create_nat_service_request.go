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

// CreateNatServiceRequest struct for CreateNatServiceRequest
type CreateNatServiceRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The allocation ID of the EIP to associate with the NAT service.<br /> If the EIP is already associated with another resource, you must first disassociate it.
	PublicIpId string `json:"PublicIpId"`
	// The ID of the Subnet in which you want to create the NAT service.
	SubnetId string `json:"SubnetId"`
}

// NewCreateNatServiceRequest instantiates a new CreateNatServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNatServiceRequest(publicIpId string, subnetId string) *CreateNatServiceRequest {
	this := CreateNatServiceRequest{}
	this.PublicIpId = publicIpId
	this.SubnetId = subnetId
	return &this
}

// NewCreateNatServiceRequestWithDefaults instantiates a new CreateNatServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNatServiceRequestWithDefaults() *CreateNatServiceRequest {
	this := CreateNatServiceRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateNatServiceRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNatServiceRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateNatServiceRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateNatServiceRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetPublicIpId returns the PublicIpId field value
func (o *CreateNatServiceRequest) GetPublicIpId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicIpId
}

// GetPublicIpIdOk returns a tuple with the PublicIpId field value
// and a boolean to check if the value has been set.
func (o *CreateNatServiceRequest) GetPublicIpIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicIpId, true
}

// SetPublicIpId sets field value
func (o *CreateNatServiceRequest) SetPublicIpId(v string) {
	o.PublicIpId = v
}

// GetSubnetId returns the SubnetId field value
func (o *CreateNatServiceRequest) GetSubnetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubnetId
}

// GetSubnetIdOk returns a tuple with the SubnetId field value
// and a boolean to check if the value has been set.
func (o *CreateNatServiceRequest) GetSubnetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubnetId, true
}

// SetSubnetId sets field value
func (o *CreateNatServiceRequest) SetSubnetId(v string) {
	o.SubnetId = v
}

func (o CreateNatServiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["PublicIpId"] = o.PublicIpId
	}
	if true {
		toSerialize["SubnetId"] = o.SubnetId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNatServiceRequest struct {
	value *CreateNatServiceRequest
	isSet bool
}

func (v NullableCreateNatServiceRequest) Get() *CreateNatServiceRequest {
	return v.value
}

func (v *NullableCreateNatServiceRequest) Set(val *CreateNatServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNatServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNatServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNatServiceRequest(val *CreateNatServiceRequest) *NullableCreateNatServiceRequest {
	return &NullableCreateNatServiceRequest{value: val, isSet: true}
}

func (v NullableCreateNatServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNatServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
