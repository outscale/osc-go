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

// FlexibleGpuCatalog Information about the flexible GPU (fGPU) that is available in the public catalog.
type FlexibleGpuCatalog struct {
	// The generations of VMs that the fGPU is compatible with.
	Generations *[]string `json:"Generations,omitempty"`
	// The maximum number of VM vCores that the fGPU is compatible with.
	MaxCpu *int32 `json:"MaxCpu,omitempty"`
	// The maximum amount of VM memory that the fGPU is compatible with.
	MaxRam *int32 `json:"MaxRam,omitempty"`
	// The model of fGPU.
	ModelName *string `json:"ModelName,omitempty"`
	// The amount of video RAM (VRAM) of the fGPU.
	VRam *int32 `json:"VRam,omitempty"`
}

// NewFlexibleGpuCatalog instantiates a new FlexibleGpuCatalog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlexibleGpuCatalog() *FlexibleGpuCatalog {
	this := FlexibleGpuCatalog{}
	return &this
}

// NewFlexibleGpuCatalogWithDefaults instantiates a new FlexibleGpuCatalog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlexibleGpuCatalogWithDefaults() *FlexibleGpuCatalog {
	this := FlexibleGpuCatalog{}
	return &this
}

// GetGenerations returns the Generations field value if set, zero value otherwise.
func (o *FlexibleGpuCatalog) GetGenerations() []string {
	if o == nil || o.Generations == nil {
		var ret []string
		return ret
	}
	return *o.Generations
}

// GetGenerationsOk returns a tuple with the Generations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleGpuCatalog) GetGenerationsOk() (*[]string, bool) {
	if o == nil || o.Generations == nil {
		return nil, false
	}
	return o.Generations, true
}

// HasGenerations returns a boolean if a field has been set.
func (o *FlexibleGpuCatalog) HasGenerations() bool {
	if o != nil && o.Generations != nil {
		return true
	}

	return false
}

// SetGenerations gets a reference to the given []string and assigns it to the Generations field.
func (o *FlexibleGpuCatalog) SetGenerations(v []string) {
	o.Generations = &v
}

// GetMaxCpu returns the MaxCpu field value if set, zero value otherwise.
func (o *FlexibleGpuCatalog) GetMaxCpu() int32 {
	if o == nil || o.MaxCpu == nil {
		var ret int32
		return ret
	}
	return *o.MaxCpu
}

// GetMaxCpuOk returns a tuple with the MaxCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleGpuCatalog) GetMaxCpuOk() (*int32, bool) {
	if o == nil || o.MaxCpu == nil {
		return nil, false
	}
	return o.MaxCpu, true
}

// HasMaxCpu returns a boolean if a field has been set.
func (o *FlexibleGpuCatalog) HasMaxCpu() bool {
	if o != nil && o.MaxCpu != nil {
		return true
	}

	return false
}

// SetMaxCpu gets a reference to the given int32 and assigns it to the MaxCpu field.
func (o *FlexibleGpuCatalog) SetMaxCpu(v int32) {
	o.MaxCpu = &v
}

// GetMaxRam returns the MaxRam field value if set, zero value otherwise.
func (o *FlexibleGpuCatalog) GetMaxRam() int32 {
	if o == nil || o.MaxRam == nil {
		var ret int32
		return ret
	}
	return *o.MaxRam
}

// GetMaxRamOk returns a tuple with the MaxRam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleGpuCatalog) GetMaxRamOk() (*int32, bool) {
	if o == nil || o.MaxRam == nil {
		return nil, false
	}
	return o.MaxRam, true
}

// HasMaxRam returns a boolean if a field has been set.
func (o *FlexibleGpuCatalog) HasMaxRam() bool {
	if o != nil && o.MaxRam != nil {
		return true
	}

	return false
}

// SetMaxRam gets a reference to the given int32 and assigns it to the MaxRam field.
func (o *FlexibleGpuCatalog) SetMaxRam(v int32) {
	o.MaxRam = &v
}

// GetModelName returns the ModelName field value if set, zero value otherwise.
func (o *FlexibleGpuCatalog) GetModelName() string {
	if o == nil || o.ModelName == nil {
		var ret string
		return ret
	}
	return *o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleGpuCatalog) GetModelNameOk() (*string, bool) {
	if o == nil || o.ModelName == nil {
		return nil, false
	}
	return o.ModelName, true
}

// HasModelName returns a boolean if a field has been set.
func (o *FlexibleGpuCatalog) HasModelName() bool {
	if o != nil && o.ModelName != nil {
		return true
	}

	return false
}

// SetModelName gets a reference to the given string and assigns it to the ModelName field.
func (o *FlexibleGpuCatalog) SetModelName(v string) {
	o.ModelName = &v
}

// GetVRam returns the VRam field value if set, zero value otherwise.
func (o *FlexibleGpuCatalog) GetVRam() int32 {
	if o == nil || o.VRam == nil {
		var ret int32
		return ret
	}
	return *o.VRam
}

// GetVRamOk returns a tuple with the VRam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleGpuCatalog) GetVRamOk() (*int32, bool) {
	if o == nil || o.VRam == nil {
		return nil, false
	}
	return o.VRam, true
}

// HasVRam returns a boolean if a field has been set.
func (o *FlexibleGpuCatalog) HasVRam() bool {
	if o != nil && o.VRam != nil {
		return true
	}

	return false
}

// SetVRam gets a reference to the given int32 and assigns it to the VRam field.
func (o *FlexibleGpuCatalog) SetVRam(v int32) {
	o.VRam = &v
}

func (o FlexibleGpuCatalog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Generations != nil {
		toSerialize["Generations"] = o.Generations
	}
	if o.MaxCpu != nil {
		toSerialize["MaxCpu"] = o.MaxCpu
	}
	if o.MaxRam != nil {
		toSerialize["MaxRam"] = o.MaxRam
	}
	if o.ModelName != nil {
		toSerialize["ModelName"] = o.ModelName
	}
	if o.VRam != nil {
		toSerialize["VRam"] = o.VRam
	}
	return json.Marshal(toSerialize)
}

type NullableFlexibleGpuCatalog struct {
	value *FlexibleGpuCatalog
	isSet bool
}

func (v NullableFlexibleGpuCatalog) Get() *FlexibleGpuCatalog {
	return v.value
}

func (v *NullableFlexibleGpuCatalog) Set(val *FlexibleGpuCatalog) {
	v.value = val
	v.isSet = true
}

func (v NullableFlexibleGpuCatalog) IsSet() bool {
	return v.isSet
}

func (v *NullableFlexibleGpuCatalog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlexibleGpuCatalog(val *FlexibleGpuCatalog) *NullableFlexibleGpuCatalog {
	return &NullableFlexibleGpuCatalog{value: val, isSet: true}
}

func (v NullableFlexibleGpuCatalog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlexibleGpuCatalog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
