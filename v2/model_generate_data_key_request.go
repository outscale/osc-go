/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.6
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// GenerateDataKeyRequest struct for GenerateDataKeyRequest
type GenerateDataKeyRequest struct {
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// Information about the encryption context.
	EncryptionContext *map[string]string `json:"EncryptionContext,omitempty"`
	// If `true`, the data key is generated in plaintext and ciphertext. If `false`, the data key is generated only in ciphertext.
	GeneratePlaintext *bool `json:"GeneratePlaintext,omitempty"`
	// The ID of the master key used to generate the data key.
	MasterKeyId string `json:"MasterKeyId"`
	// The size of the data key you want to generate, in bytes (between 1 and 1024).
	Size int32 `json:"Size"`
}

// NewGenerateDataKeyRequest instantiates a new GenerateDataKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateDataKeyRequest(masterKeyId string, size int32, ) *GenerateDataKeyRequest {
	this := GenerateDataKeyRequest{}
	this.MasterKeyId = masterKeyId
	this.Size = size
	return &this
}

// NewGenerateDataKeyRequestWithDefaults instantiates a new GenerateDataKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateDataKeyRequestWithDefaults() *GenerateDataKeyRequest {
	this := GenerateDataKeyRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *GenerateDataKeyRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateDataKeyRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *GenerateDataKeyRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *GenerateDataKeyRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetEncryptionContext returns the EncryptionContext field value if set, zero value otherwise.
func (o *GenerateDataKeyRequest) GetEncryptionContext() map[string]string {
	if o == nil || o.EncryptionContext == nil {
		var ret map[string]string
		return ret
	}
	return *o.EncryptionContext
}

// GetEncryptionContextOk returns a tuple with the EncryptionContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateDataKeyRequest) GetEncryptionContextOk() (*map[string]string, bool) {
	if o == nil || o.EncryptionContext == nil {
		return nil, false
	}
	return o.EncryptionContext, true
}

// HasEncryptionContext returns a boolean if a field has been set.
func (o *GenerateDataKeyRequest) HasEncryptionContext() bool {
	if o != nil && o.EncryptionContext != nil {
		return true
	}

	return false
}

// SetEncryptionContext gets a reference to the given map[string]string and assigns it to the EncryptionContext field.
func (o *GenerateDataKeyRequest) SetEncryptionContext(v map[string]string) {
	o.EncryptionContext = &v
}

// GetGeneratePlaintext returns the GeneratePlaintext field value if set, zero value otherwise.
func (o *GenerateDataKeyRequest) GetGeneratePlaintext() bool {
	if o == nil || o.GeneratePlaintext == nil {
		var ret bool
		return ret
	}
	return *o.GeneratePlaintext
}

// GetGeneratePlaintextOk returns a tuple with the GeneratePlaintext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateDataKeyRequest) GetGeneratePlaintextOk() (*bool, bool) {
	if o == nil || o.GeneratePlaintext == nil {
		return nil, false
	}
	return o.GeneratePlaintext, true
}

// HasGeneratePlaintext returns a boolean if a field has been set.
func (o *GenerateDataKeyRequest) HasGeneratePlaintext() bool {
	if o != nil && o.GeneratePlaintext != nil {
		return true
	}

	return false
}

// SetGeneratePlaintext gets a reference to the given bool and assigns it to the GeneratePlaintext field.
func (o *GenerateDataKeyRequest) SetGeneratePlaintext(v bool) {
	o.GeneratePlaintext = &v
}

// GetMasterKeyId returns the MasterKeyId field value
func (o *GenerateDataKeyRequest) GetMasterKeyId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.MasterKeyId
}

// GetMasterKeyIdOk returns a tuple with the MasterKeyId field value
// and a boolean to check if the value has been set.
func (o *GenerateDataKeyRequest) GetMasterKeyIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MasterKeyId, true
}

// SetMasterKeyId sets field value
func (o *GenerateDataKeyRequest) SetMasterKeyId(v string) {
	o.MasterKeyId = v
}

// GetSize returns the Size field value
func (o *GenerateDataKeyRequest) GetSize() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *GenerateDataKeyRequest) GetSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *GenerateDataKeyRequest) SetSize(v int32) {
	o.Size = v
}

func (o GenerateDataKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.EncryptionContext != nil {
		toSerialize["EncryptionContext"] = o.EncryptionContext
	}
	if o.GeneratePlaintext != nil {
		toSerialize["GeneratePlaintext"] = o.GeneratePlaintext
	}
	if true {
		toSerialize["MasterKeyId"] = o.MasterKeyId
	}
	if true {
		toSerialize["Size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableGenerateDataKeyRequest struct {
	value *GenerateDataKeyRequest
	isSet bool
}

func (v NullableGenerateDataKeyRequest) Get() *GenerateDataKeyRequest {
	return v.value
}

func (v *NullableGenerateDataKeyRequest) Set(val *GenerateDataKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateDataKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateDataKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateDataKeyRequest(val *GenerateDataKeyRequest) *NullableGenerateDataKeyRequest {
	return &NullableGenerateDataKeyRequest{value: val, isSet: true}
}

func (v NullableGenerateDataKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateDataKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

