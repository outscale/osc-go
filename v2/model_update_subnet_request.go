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

// UpdateSubnetRequest struct for UpdateSubnetRequest
type UpdateSubnetRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// If true, a public IP address is assigned to the network interface cards (NICs) created in the specified Subnet.
	MapPublicIpOnLaunch bool `json:"MapPublicIpOnLaunch"`
	// The ID of the Subnet.
	SubnetId string `json:"SubnetId"`
}

// NewUpdateSubnetRequest instantiates a new UpdateSubnetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSubnetRequest(mapPublicIpOnLaunch bool, subnetId string) *UpdateSubnetRequest {
	this := UpdateSubnetRequest{}
	this.MapPublicIpOnLaunch = mapPublicIpOnLaunch
	this.SubnetId = subnetId
	return &this
}

// NewUpdateSubnetRequestWithDefaults instantiates a new UpdateSubnetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSubnetRequestWithDefaults() *UpdateSubnetRequest {
	this := UpdateSubnetRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateSubnetRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSubnetRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateSubnetRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateSubnetRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetMapPublicIpOnLaunch returns the MapPublicIpOnLaunch field value
func (o *UpdateSubnetRequest) GetMapPublicIpOnLaunch() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MapPublicIpOnLaunch
}

// GetMapPublicIpOnLaunchOk returns a tuple with the MapPublicIpOnLaunch field value
// and a boolean to check if the value has been set.
func (o *UpdateSubnetRequest) GetMapPublicIpOnLaunchOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MapPublicIpOnLaunch, true
}

// SetMapPublicIpOnLaunch sets field value
func (o *UpdateSubnetRequest) SetMapPublicIpOnLaunch(v bool) {
	o.MapPublicIpOnLaunch = v
}

// GetSubnetId returns the SubnetId field value
func (o *UpdateSubnetRequest) GetSubnetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubnetId
}

// GetSubnetIdOk returns a tuple with the SubnetId field value
// and a boolean to check if the value has been set.
func (o *UpdateSubnetRequest) GetSubnetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubnetId, true
}

// SetSubnetId sets field value
func (o *UpdateSubnetRequest) SetSubnetId(v string) {
	o.SubnetId = v
}

func (o UpdateSubnetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["MapPublicIpOnLaunch"] = o.MapPublicIpOnLaunch
	}
	if true {
		toSerialize["SubnetId"] = o.SubnetId
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSubnetRequest struct {
	value *UpdateSubnetRequest
	isSet bool
}

func (v NullableUpdateSubnetRequest) Get() *UpdateSubnetRequest {
	return v.value
}

func (v *NullableUpdateSubnetRequest) Set(val *UpdateSubnetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSubnetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSubnetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSubnetRequest(val *UpdateSubnetRequest) *NullableUpdateSubnetRequest {
	return &NullableUpdateSubnetRequest{value: val, isSet: true}
}

func (v NullableUpdateSubnetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSubnetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
