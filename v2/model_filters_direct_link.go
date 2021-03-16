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

// FiltersDirectLink One or more filters.
type FiltersDirectLink struct {
	// The IDs of the DirectLinks.
	DirectLinkIds *[]string `json:"DirectLinkIds,omitempty"`
}

// NewFiltersDirectLink instantiates a new FiltersDirectLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersDirectLink() *FiltersDirectLink {
	this := FiltersDirectLink{}
	return &this
}

// NewFiltersDirectLinkWithDefaults instantiates a new FiltersDirectLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersDirectLinkWithDefaults() *FiltersDirectLink {
	this := FiltersDirectLink{}
	return &this
}

// GetDirectLinkIds returns the DirectLinkIds field value if set, zero value otherwise.
func (o *FiltersDirectLink) GetDirectLinkIds() []string {
	if o == nil || o.DirectLinkIds == nil {
		var ret []string
		return ret
	}
	return *o.DirectLinkIds
}

// GetDirectLinkIdsOk returns a tuple with the DirectLinkIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersDirectLink) GetDirectLinkIdsOk() (*[]string, bool) {
	if o == nil || o.DirectLinkIds == nil {
		return nil, false
	}
	return o.DirectLinkIds, true
}

// HasDirectLinkIds returns a boolean if a field has been set.
func (o *FiltersDirectLink) HasDirectLinkIds() bool {
	if o != nil && o.DirectLinkIds != nil {
		return true
	}

	return false
}

// SetDirectLinkIds gets a reference to the given []string and assigns it to the DirectLinkIds field.
func (o *FiltersDirectLink) SetDirectLinkIds(v []string) {
	o.DirectLinkIds = &v
}

func (o FiltersDirectLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DirectLinkIds != nil {
		toSerialize["DirectLinkIds"] = o.DirectLinkIds
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersDirectLink struct {
	value *FiltersDirectLink
	isSet bool
}

func (v NullableFiltersDirectLink) Get() *FiltersDirectLink {
	return v.value
}

func (v *NullableFiltersDirectLink) Set(val *FiltersDirectLink) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersDirectLink) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersDirectLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersDirectLink(val *FiltersDirectLink) *NullableFiltersDirectLink {
	return &NullableFiltersDirectLink{value: val, isSet: true}
}

func (v NullableFiltersDirectLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersDirectLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
