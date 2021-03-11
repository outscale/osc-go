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

// UpdateSnapshotRequest struct for UpdateSnapshotRequest
type UpdateSnapshotRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun                    *bool                         `json:"DryRun,omitempty"`
	PermissionsToCreateVolume PermissionsOnResourceCreation `json:"PermissionsToCreateVolume"`
	// The ID of the snapshot.
	SnapshotId string `json:"SnapshotId"`
}

// NewUpdateSnapshotRequest instantiates a new UpdateSnapshotRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSnapshotRequest(permissionsToCreateVolume PermissionsOnResourceCreation, snapshotId string) *UpdateSnapshotRequest {
	this := UpdateSnapshotRequest{}
	this.PermissionsToCreateVolume = permissionsToCreateVolume
	this.SnapshotId = snapshotId
	return &this
}

// NewUpdateSnapshotRequestWithDefaults instantiates a new UpdateSnapshotRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSnapshotRequestWithDefaults() *UpdateSnapshotRequest {
	this := UpdateSnapshotRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *UpdateSnapshotRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSnapshotRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *UpdateSnapshotRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *UpdateSnapshotRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetPermissionsToCreateVolume returns the PermissionsToCreateVolume field value
func (o *UpdateSnapshotRequest) GetPermissionsToCreateVolume() PermissionsOnResourceCreation {
	if o == nil {
		var ret PermissionsOnResourceCreation
		return ret
	}

	return o.PermissionsToCreateVolume
}

// GetPermissionsToCreateVolumeOk returns a tuple with the PermissionsToCreateVolume field value
// and a boolean to check if the value has been set.
func (o *UpdateSnapshotRequest) GetPermissionsToCreateVolumeOk() (*PermissionsOnResourceCreation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PermissionsToCreateVolume, true
}

// SetPermissionsToCreateVolume sets field value
func (o *UpdateSnapshotRequest) SetPermissionsToCreateVolume(v PermissionsOnResourceCreation) {
	o.PermissionsToCreateVolume = v
}

// GetSnapshotId returns the SnapshotId field value
func (o *UpdateSnapshotRequest) GetSnapshotId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value
// and a boolean to check if the value has been set.
func (o *UpdateSnapshotRequest) GetSnapshotIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SnapshotId, true
}

// SetSnapshotId sets field value
func (o *UpdateSnapshotRequest) SetSnapshotId(v string) {
	o.SnapshotId = v
}

func (o UpdateSnapshotRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["PermissionsToCreateVolume"] = o.PermissionsToCreateVolume
	}
	if true {
		toSerialize["SnapshotId"] = o.SnapshotId
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSnapshotRequest struct {
	value *UpdateSnapshotRequest
	isSet bool
}

func (v NullableUpdateSnapshotRequest) Get() *UpdateSnapshotRequest {
	return v.value
}

func (v *NullableUpdateSnapshotRequest) Set(val *UpdateSnapshotRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSnapshotRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSnapshotRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSnapshotRequest(val *UpdateSnapshotRequest) *NullableUpdateSnapshotRequest {
	return &NullableUpdateSnapshotRequest{value: val, isSet: true}
}

func (v NullableUpdateSnapshotRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSnapshotRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
