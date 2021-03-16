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

// ReadAccountsRequest struct for ReadAccountsRequest
type ReadAccountsRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
}

// NewReadAccountsRequest instantiates a new ReadAccountsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadAccountsRequest() *ReadAccountsRequest {
	this := ReadAccountsRequest{}
	return &this
}

// NewReadAccountsRequestWithDefaults instantiates a new ReadAccountsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadAccountsRequestWithDefaults() *ReadAccountsRequest {
	this := ReadAccountsRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ReadAccountsRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadAccountsRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ReadAccountsRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ReadAccountsRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

func (o ReadAccountsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	return json.Marshal(toSerialize)
}

type NullableReadAccountsRequest struct {
	value *ReadAccountsRequest
	isSet bool
}

func (v NullableReadAccountsRequest) Get() *ReadAccountsRequest {
	return v.value
}

func (v *NullableReadAccountsRequest) Set(val *ReadAccountsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReadAccountsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReadAccountsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadAccountsRequest(val *ReadAccountsRequest) *NullableReadAccountsRequest {
	return &NullableReadAccountsRequest{value: val, isSet: true}
}

func (v NullableReadAccountsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadAccountsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
