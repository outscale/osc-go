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

// CreateVolumeRequest struct for CreateVolumeRequest
type CreateVolumeRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The number of I/O operations per second (IOPS). This parameter must be specified only if you create an `io1` volume. The maximum number of IOPS allowed for `io1` volumes is `13000`.
	Iops *int32 `json:"Iops,omitempty"`
	// The size of the volume, in gibibytes (GiB). The maximum allowed size for a volume is 14901 GiB. This parameter is required if the volume is not created from a snapshot (`SnapshotId` unspecified).
	Size *int32 `json:"Size,omitempty"`
	// The ID of the snapshot from which you want to create the volume.
	SnapshotId *string `json:"SnapshotId,omitempty"`
	// The Subregion in which you want to create the volume.
	SubregionName string `json:"SubregionName"`
	// The type of volume you want to create (`io1` \\| `gp2` \\| `standard`). If not specified, a `standard` volume is created.<br /> For more information about volume types, see [Volume Types and IOPS](https://wiki.outscale.net/display/EN/About+Volumes#AboutVolumes-VolumeTypesVolumeTypesandIOPS).
	VolumeType *string `json:"VolumeType,omitempty"`
}

// NewCreateVolumeRequest instantiates a new CreateVolumeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVolumeRequest(subregionName string) *CreateVolumeRequest {
	this := CreateVolumeRequest{}
	this.SubregionName = subregionName
	return &this
}

// NewCreateVolumeRequestWithDefaults instantiates a new CreateVolumeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVolumeRequestWithDefaults() *CreateVolumeRequest {
	this := CreateVolumeRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateVolumeRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVolumeRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateVolumeRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateVolumeRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetIops returns the Iops field value if set, zero value otherwise.
func (o *CreateVolumeRequest) GetIops() int32 {
	if o == nil || o.Iops == nil {
		var ret int32
		return ret
	}
	return *o.Iops
}

// GetIopsOk returns a tuple with the Iops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVolumeRequest) GetIopsOk() (*int32, bool) {
	if o == nil || o.Iops == nil {
		return nil, false
	}
	return o.Iops, true
}

// HasIops returns a boolean if a field has been set.
func (o *CreateVolumeRequest) HasIops() bool {
	if o != nil && o.Iops != nil {
		return true
	}

	return false
}

// SetIops gets a reference to the given int32 and assigns it to the Iops field.
func (o *CreateVolumeRequest) SetIops(v int32) {
	o.Iops = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *CreateVolumeRequest) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVolumeRequest) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *CreateVolumeRequest) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *CreateVolumeRequest) SetSize(v int32) {
	o.Size = &v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *CreateVolumeRequest) GetSnapshotId() string {
	if o == nil || o.SnapshotId == nil {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVolumeRequest) GetSnapshotIdOk() (*string, bool) {
	if o == nil || o.SnapshotId == nil {
		return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *CreateVolumeRequest) HasSnapshotId() bool {
	if o != nil && o.SnapshotId != nil {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *CreateVolumeRequest) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetSubregionName returns the SubregionName field value
func (o *CreateVolumeRequest) GetSubregionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubregionName
}

// GetSubregionNameOk returns a tuple with the SubregionName field value
// and a boolean to check if the value has been set.
func (o *CreateVolumeRequest) GetSubregionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubregionName, true
}

// SetSubregionName sets field value
func (o *CreateVolumeRequest) SetSubregionName(v string) {
	o.SubregionName = v
}

// GetVolumeType returns the VolumeType field value if set, zero value otherwise.
func (o *CreateVolumeRequest) GetVolumeType() string {
	if o == nil || o.VolumeType == nil {
		var ret string
		return ret
	}
	return *o.VolumeType
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVolumeRequest) GetVolumeTypeOk() (*string, bool) {
	if o == nil || o.VolumeType == nil {
		return nil, false
	}
	return o.VolumeType, true
}

// HasVolumeType returns a boolean if a field has been set.
func (o *CreateVolumeRequest) HasVolumeType() bool {
	if o != nil && o.VolumeType != nil {
		return true
	}

	return false
}

// SetVolumeType gets a reference to the given string and assigns it to the VolumeType field.
func (o *CreateVolumeRequest) SetVolumeType(v string) {
	o.VolumeType = &v
}

func (o CreateVolumeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.Iops != nil {
		toSerialize["Iops"] = o.Iops
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.SnapshotId != nil {
		toSerialize["SnapshotId"] = o.SnapshotId
	}
	if true {
		toSerialize["SubregionName"] = o.SubregionName
	}
	if o.VolumeType != nil {
		toSerialize["VolumeType"] = o.VolumeType
	}
	return json.Marshal(toSerialize)
}

type NullableCreateVolumeRequest struct {
	value *CreateVolumeRequest
	isSet bool
}

func (v NullableCreateVolumeRequest) Get() *CreateVolumeRequest {
	return v.value
}

func (v *NullableCreateVolumeRequest) Set(val *CreateVolumeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVolumeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVolumeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVolumeRequest(val *CreateVolumeRequest) *NullableCreateVolumeRequest {
	return &NullableCreateVolumeRequest{value: val, isSet: true}
}

func (v NullableCreateVolumeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVolumeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
