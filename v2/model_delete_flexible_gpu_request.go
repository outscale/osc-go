/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.10
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// DeleteFlexibleGpuRequest struct for DeleteFlexibleGpuRequest
type DeleteFlexibleGpuRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the fGPU you want to delete.
	FlexibleGpuId string `json:"FlexibleGpuId"`
}

// NewDeleteFlexibleGpuRequest instantiates a new DeleteFlexibleGpuRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteFlexibleGpuRequest(flexibleGpuId string) *DeleteFlexibleGpuRequest {
	this := DeleteFlexibleGpuRequest{}
	this.FlexibleGpuId = flexibleGpuId
	return &this
}

// NewDeleteFlexibleGpuRequestWithDefaults instantiates a new DeleteFlexibleGpuRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteFlexibleGpuRequestWithDefaults() *DeleteFlexibleGpuRequest {
	this := DeleteFlexibleGpuRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *DeleteFlexibleGpuRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteFlexibleGpuRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *DeleteFlexibleGpuRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *DeleteFlexibleGpuRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFlexibleGpuId returns the FlexibleGpuId field value
func (o *DeleteFlexibleGpuRequest) GetFlexibleGpuId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FlexibleGpuId
}

// GetFlexibleGpuIdOk returns a tuple with the FlexibleGpuId field value
// and a boolean to check if the value has been set.
func (o *DeleteFlexibleGpuRequest) GetFlexibleGpuIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlexibleGpuId, true
}

// SetFlexibleGpuId sets field value
func (o *DeleteFlexibleGpuRequest) SetFlexibleGpuId(v string) {
	o.FlexibleGpuId = v
}

func (o DeleteFlexibleGpuRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["FlexibleGpuId"] = o.FlexibleGpuId
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteFlexibleGpuRequest struct {
	value *DeleteFlexibleGpuRequest
	isSet bool
}

func (v NullableDeleteFlexibleGpuRequest) Get() *DeleteFlexibleGpuRequest {
	return v.value
}

func (v *NullableDeleteFlexibleGpuRequest) Set(val *DeleteFlexibleGpuRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteFlexibleGpuRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteFlexibleGpuRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteFlexibleGpuRequest(val *DeleteFlexibleGpuRequest) *NullableDeleteFlexibleGpuRequest {
	return &NullableDeleteFlexibleGpuRequest{value: val, isSet: true}
}

func (v NullableDeleteFlexibleGpuRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteFlexibleGpuRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
