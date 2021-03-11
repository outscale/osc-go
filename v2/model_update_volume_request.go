/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.7
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// UpdateVolumeRequest struct for UpdateVolumeRequest
type UpdateVolumeRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The new size of the volume, in gibibytes (GiB). This value must be equal to or greater than the current size of the volume.
	Size *int32 `json:"Size,omitempty"`
	// The ID of the volume you want to update.
	VolumeId string `json:"VolumeId"`
}

// NewUpdateVolumeRequest instantiates a new UpdateVolumeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVolumeRequest(volumeId string) *UpdateVolumeRequest {
	this := UpdateVolumeRequest{}
	this.VolumeId = volumeId
	return &this
}

// NewUpdateVolumeRequestWithDefaults instantiates a new UpdateVolumeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVolumeRequestWithDefaults() *UpdateVolumeRequest {
	this := UpdateVolumeRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateVolumeRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVolumeRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateVolumeRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateVolumeRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *UpdateVolumeRequest) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVolumeRequest) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *UpdateVolumeRequest) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *UpdateVolumeRequest) SetSize(v int32) {
	o.Size = &v
}

// GetVolumeId returns the VolumeId field value
func (o *UpdateVolumeRequest) GetVolumeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VolumeId
}

// GetVolumeIdOk returns a tuple with the VolumeId field value
// and a boolean to check if the value has been set.
func (o *UpdateVolumeRequest) GetVolumeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeId, true
}

// SetVolumeId sets field value
func (o *UpdateVolumeRequest) SetVolumeId(v string) {
	o.VolumeId = v
}

func (o UpdateVolumeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if true {
		toSerialize["VolumeId"] = o.VolumeId
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateVolumeRequest struct {
	value *UpdateVolumeRequest
	isSet bool
}

func (v NullableUpdateVolumeRequest) Get() *UpdateVolumeRequest {
	return v.value
}

func (v *NullableUpdateVolumeRequest) Set(val *UpdateVolumeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVolumeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVolumeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVolumeRequest(val *UpdateVolumeRequest) *NullableUpdateVolumeRequest {
	return &NullableUpdateVolumeRequest{value: val, isSet: true}
}

func (v NullableUpdateVolumeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVolumeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
