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

// MasterKey Information about a master key.
type MasterKey struct {
	// The date and time of creation of the master key.
	CreationDate *string `json:"CreationDate,omitempty"`
	// The date and time of scheduled deletion of the master key.
	DeletionDate *string `json:"DeletionDate,omitempty"`
	//  The description of the master key.
	Description *string `json:"Description,omitempty"`
	// The ID of the master key.
	MasterKeyId *string `json:"MasterKeyId,omitempty"`
	// The state of the master key (`Enabled` \\| `Disabled` \\| `PendingDeletion`).
	State *string `json:"State,omitempty"`
}

// NewMasterKey instantiates a new MasterKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMasterKey() *MasterKey {
	this := MasterKey{}
	return &this
}

// NewMasterKeyWithDefaults instantiates a new MasterKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMasterKeyWithDefaults() *MasterKey {
	this := MasterKey{}
	return &this
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *MasterKey) GetCreationDate() string {
	if o == nil || o.CreationDate == nil {
		var ret string
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterKey) GetCreationDateOk() (*string, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *MasterKey) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given string and assigns it to the CreationDate field.
func (o *MasterKey) SetCreationDate(v string) {
	o.CreationDate = &v
}

// GetDeletionDate returns the DeletionDate field value if set, zero value otherwise.
func (o *MasterKey) GetDeletionDate() string {
	if o == nil || o.DeletionDate == nil {
		var ret string
		return ret
	}
	return *o.DeletionDate
}

// GetDeletionDateOk returns a tuple with the DeletionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterKey) GetDeletionDateOk() (*string, bool) {
	if o == nil || o.DeletionDate == nil {
		return nil, false
	}
	return o.DeletionDate, true
}

// HasDeletionDate returns a boolean if a field has been set.
func (o *MasterKey) HasDeletionDate() bool {
	if o != nil && o.DeletionDate != nil {
		return true
	}

	return false
}

// SetDeletionDate gets a reference to the given string and assigns it to the DeletionDate field.
func (o *MasterKey) SetDeletionDate(v string) {
	o.DeletionDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MasterKey) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterKey) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MasterKey) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MasterKey) SetDescription(v string) {
	o.Description = &v
}

// GetMasterKeyId returns the MasterKeyId field value if set, zero value otherwise.
func (o *MasterKey) GetMasterKeyId() string {
	if o == nil || o.MasterKeyId == nil {
		var ret string
		return ret
	}
	return *o.MasterKeyId
}

// GetMasterKeyIdOk returns a tuple with the MasterKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterKey) GetMasterKeyIdOk() (*string, bool) {
	if o == nil || o.MasterKeyId == nil {
		return nil, false
	}
	return o.MasterKeyId, true
}

// HasMasterKeyId returns a boolean if a field has been set.
func (o *MasterKey) HasMasterKeyId() bool {
	if o != nil && o.MasterKeyId != nil {
		return true
	}

	return false
}

// SetMasterKeyId gets a reference to the given string and assigns it to the MasterKeyId field.
func (o *MasterKey) SetMasterKeyId(v string) {
	o.MasterKeyId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *MasterKey) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MasterKey) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MasterKey) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *MasterKey) SetState(v string) {
	o.State = &v
}

func (o MasterKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreationDate != nil {
		toSerialize["CreationDate"] = o.CreationDate
	}
	if o.DeletionDate != nil {
		toSerialize["DeletionDate"] = o.DeletionDate
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.MasterKeyId != nil {
		toSerialize["MasterKeyId"] = o.MasterKeyId
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableMasterKey struct {
	value *MasterKey
	isSet bool
}

func (v NullableMasterKey) Get() *MasterKey {
	return v.value
}

func (v *NullableMasterKey) Set(val *MasterKey) {
	v.value = val
	v.isSet = true
}

func (v NullableMasterKey) IsSet() bool {
	return v.isSet
}

func (v *NullableMasterKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMasterKey(val *MasterKey) *NullableMasterKey {
	return &NullableMasterKey{value: val, isSet: true}
}

func (v NullableMasterKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMasterKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

