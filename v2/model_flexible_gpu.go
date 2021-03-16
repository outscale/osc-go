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

// FlexibleGpu Information about the flexible GPU (fGPU).
type FlexibleGpu struct {
	// If true, the fGPU is deleted when the VM is terminated.
	DeleteOnVmDeletion *bool `json:"DeleteOnVmDeletion,omitempty"`
	// The ID of the fGPU.
	FlexibleGpuId *string `json:"FlexibleGpuId,omitempty"`
	// The compatible processor generation.
	Generation *string `json:"Generation,omitempty"`
	// The model of fGPU. For more information, see [About Flexible GPUs](https://wiki.outscale.net/display/EN/About+Flexible+GPUs).
	ModelName *string `json:"ModelName,omitempty"`
	// The state of the fGPU (`allocated` \\| `attaching` \\| `attached` \\| `detaching`).
	State *string `json:"State,omitempty"`
	// The Subregion where the fGPU is located.
	SubregionName *string `json:"SubregionName,omitempty"`
	// The ID of the VM the fGPU is attached to, if any.
	VmId *string `json:"VmId,omitempty"`
}

// NewFlexibleGpu instantiates a new FlexibleGpu object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlexibleGpu() *FlexibleGpu {
	this := FlexibleGpu{}
	return &this
}

// NewFlexibleGpuWithDefaults instantiates a new FlexibleGpu object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlexibleGpuWithDefaults() *FlexibleGpu {
	this := FlexibleGpu{}
	return &this
}

// GetDeleteOnVmDeletion returns the DeleteOnVmDeletion field value if set, zero value otherwise.
func (o *FlexibleGpu) GetDeleteOnVmDeletion() bool {
	if o == nil || o.DeleteOnVmDeletion == nil {
		var ret bool
		return ret
	}
	return *o.DeleteOnVmDeletion
}

// GetDeleteOnVmDeletionOk returns a tuple with the DeleteOnVmDeletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleGpu) GetDeleteOnVmDeletionOk() (*bool, bool) {
	if o == nil || o.DeleteOnVmDeletion == nil {
		return nil, false
	}
	return o.DeleteOnVmDeletion, true
}

// HasDeleteOnVmDeletion returns a boolean if a field has been set.
func (o *FlexibleGpu) HasDeleteOnVmDeletion() bool {
	if o != nil && o.DeleteOnVmDeletion != nil {
		return true
	}

	return false
}

// SetDeleteOnVmDeletion gets a reference to the given bool and assigns it to the DeleteOnVmDeletion field.
func (o *FlexibleGpu) SetDeleteOnVmDeletion(v bool) {
	o.DeleteOnVmDeletion = &v
}

// GetFlexibleGpuId returns the FlexibleGpuId field value if set, zero value otherwise.
func (o *FlexibleGpu) GetFlexibleGpuId() string {
	if o == nil || o.FlexibleGpuId == nil {
		var ret string
		return ret
	}
	return *o.FlexibleGpuId
}

// GetFlexibleGpuIdOk returns a tuple with the FlexibleGpuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleGpu) GetFlexibleGpuIdOk() (*string, bool) {
	if o == nil || o.FlexibleGpuId == nil {
		return nil, false
	}
	return o.FlexibleGpuId, true
}

// HasFlexibleGpuId returns a boolean if a field has been set.
func (o *FlexibleGpu) HasFlexibleGpuId() bool {
	if o != nil && o.FlexibleGpuId != nil {
		return true
	}

	return false
}

// SetFlexibleGpuId gets a reference to the given string and assigns it to the FlexibleGpuId field.
func (o *FlexibleGpu) SetFlexibleGpuId(v string) {
	o.FlexibleGpuId = &v
}

// GetGeneration returns the Generation field value if set, zero value otherwise.
func (o *FlexibleGpu) GetGeneration() string {
	if o == nil || o.Generation == nil {
		var ret string
		return ret
	}
	return *o.Generation
}

// GetGenerationOk returns a tuple with the Generation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleGpu) GetGenerationOk() (*string, bool) {
	if o == nil || o.Generation == nil {
		return nil, false
	}
	return o.Generation, true
}

// HasGeneration returns a boolean if a field has been set.
func (o *FlexibleGpu) HasGeneration() bool {
	if o != nil && o.Generation != nil {
		return true
	}

	return false
}

// SetGeneration gets a reference to the given string and assigns it to the Generation field.
func (o *FlexibleGpu) SetGeneration(v string) {
	o.Generation = &v
}

// GetModelName returns the ModelName field value if set, zero value otherwise.
func (o *FlexibleGpu) GetModelName() string {
	if o == nil || o.ModelName == nil {
		var ret string
		return ret
	}
	return *o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleGpu) GetModelNameOk() (*string, bool) {
	if o == nil || o.ModelName == nil {
		return nil, false
	}
	return o.ModelName, true
}

// HasModelName returns a boolean if a field has been set.
func (o *FlexibleGpu) HasModelName() bool {
	if o != nil && o.ModelName != nil {
		return true
	}

	return false
}

// SetModelName gets a reference to the given string and assigns it to the ModelName field.
func (o *FlexibleGpu) SetModelName(v string) {
	o.ModelName = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *FlexibleGpu) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleGpu) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *FlexibleGpu) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *FlexibleGpu) SetState(v string) {
	o.State = &v
}

// GetSubregionName returns the SubregionName field value if set, zero value otherwise.
func (o *FlexibleGpu) GetSubregionName() string {
	if o == nil || o.SubregionName == nil {
		var ret string
		return ret
	}
	return *o.SubregionName
}

// GetSubregionNameOk returns a tuple with the SubregionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleGpu) GetSubregionNameOk() (*string, bool) {
	if o == nil || o.SubregionName == nil {
		return nil, false
	}
	return o.SubregionName, true
}

// HasSubregionName returns a boolean if a field has been set.
func (o *FlexibleGpu) HasSubregionName() bool {
	if o != nil && o.SubregionName != nil {
		return true
	}

	return false
}

// SetSubregionName gets a reference to the given string and assigns it to the SubregionName field.
func (o *FlexibleGpu) SetSubregionName(v string) {
	o.SubregionName = &v
}

// GetVmId returns the VmId field value if set, zero value otherwise.
func (o *FlexibleGpu) GetVmId() string {
	if o == nil || o.VmId == nil {
		var ret string
		return ret
	}
	return *o.VmId
}

// GetVmIdOk returns a tuple with the VmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexibleGpu) GetVmIdOk() (*string, bool) {
	if o == nil || o.VmId == nil {
		return nil, false
	}
	return o.VmId, true
}

// HasVmId returns a boolean if a field has been set.
func (o *FlexibleGpu) HasVmId() bool {
	if o != nil && o.VmId != nil {
		return true
	}

	return false
}

// SetVmId gets a reference to the given string and assigns it to the VmId field.
func (o *FlexibleGpu) SetVmId(v string) {
	o.VmId = &v
}

func (o FlexibleGpu) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteOnVmDeletion != nil {
		toSerialize["DeleteOnVmDeletion"] = o.DeleteOnVmDeletion
	}
	if o.FlexibleGpuId != nil {
		toSerialize["FlexibleGpuId"] = o.FlexibleGpuId
	}
	if o.Generation != nil {
		toSerialize["Generation"] = o.Generation
	}
	if o.ModelName != nil {
		toSerialize["ModelName"] = o.ModelName
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.SubregionName != nil {
		toSerialize["SubregionName"] = o.SubregionName
	}
	if o.VmId != nil {
		toSerialize["VmId"] = o.VmId
	}
	return json.Marshal(toSerialize)
}

type NullableFlexibleGpu struct {
	value *FlexibleGpu
	isSet bool
}

func (v NullableFlexibleGpu) Get() *FlexibleGpu {
	return v.value
}

func (v *NullableFlexibleGpu) Set(val *FlexibleGpu) {
	v.value = val
	v.isSet = true
}

func (v NullableFlexibleGpu) IsSet() bool {
	return v.isSet
}

func (v *NullableFlexibleGpu) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlexibleGpu(val *FlexibleGpu) *NullableFlexibleGpu {
	return &NullableFlexibleGpu{value: val, isSet: true}
}

func (v NullableFlexibleGpu) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlexibleGpu) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
