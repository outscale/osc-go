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

// ReadRouteTablesRequest struct for ReadRouteTablesRequest
type ReadRouteTablesRequest struct {
	// If true, checks whether you have the required permissions to perform the action.
	DryRun  *bool              `json:"DryRun,omitempty"`
	Filters *FiltersRouteTable `json:"Filters,omitempty"`
}

// NewReadRouteTablesRequest instantiates a new ReadRouteTablesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadRouteTablesRequest() *ReadRouteTablesRequest {
	this := ReadRouteTablesRequest{}
	return &this
}

// NewReadRouteTablesRequestWithDefaults instantiates a new ReadRouteTablesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadRouteTablesRequestWithDefaults() *ReadRouteTablesRequest {
	this := ReadRouteTablesRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ReadRouteTablesRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadRouteTablesRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ReadRouteTablesRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ReadRouteTablesRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ReadRouteTablesRequest) GetFilters() FiltersRouteTable {
	if o == nil || o.Filters == nil {
		var ret FiltersRouteTable
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadRouteTablesRequest) GetFiltersOk() (*FiltersRouteTable, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ReadRouteTablesRequest) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given FiltersRouteTable and assigns it to the Filters field.
func (o *ReadRouteTablesRequest) SetFilters(v FiltersRouteTable) {
	o.Filters = &v
}

func (o ReadRouteTablesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.Filters != nil {
		toSerialize["Filters"] = o.Filters
	}
	return json.Marshal(toSerialize)
}

type NullableReadRouteTablesRequest struct {
	value *ReadRouteTablesRequest
	isSet bool
}

func (v NullableReadRouteTablesRequest) Get() *ReadRouteTablesRequest {
	return v.value
}

func (v *NullableReadRouteTablesRequest) Set(val *ReadRouteTablesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReadRouteTablesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReadRouteTablesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadRouteTablesRequest(val *ReadRouteTablesRequest) *NullableReadRouteTablesRequest {
	return &NullableReadRouteTablesRequest{value: val, isSet: true}
}

func (v NullableReadRouteTablesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadRouteTablesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
