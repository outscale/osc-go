/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.4
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// ReadApiLogsRequest struct for ReadApiLogsRequest
type ReadApiLogsRequest struct {
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	Filters *FiltersApiLog `json:"Filters,omitempty"`
	// The token to request the next page of results.
	NextPageToken *string `json:"NextPageToken,omitempty"`
	// The maximum number of items returned in a single page. By default, 100.
	ResultsPerPage *int32 `json:"ResultsPerPage,omitempty"`
	With *With `json:"With,omitempty"`
}

// NewReadApiLogsRequest instantiates a new ReadApiLogsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadApiLogsRequest() *ReadApiLogsRequest {
	this := ReadApiLogsRequest{}
	return &this
}

// NewReadApiLogsRequestWithDefaults instantiates a new ReadApiLogsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadApiLogsRequestWithDefaults() *ReadApiLogsRequest {
	this := ReadApiLogsRequest{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *ReadApiLogsRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadApiLogsRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *ReadApiLogsRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *ReadApiLogsRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ReadApiLogsRequest) GetFilters() FiltersApiLog {
	if o == nil || o.Filters == nil {
		var ret FiltersApiLog
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadApiLogsRequest) GetFiltersOk() (*FiltersApiLog, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ReadApiLogsRequest) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given FiltersApiLog and assigns it to the Filters field.
func (o *ReadApiLogsRequest) SetFilters(v FiltersApiLog) {
	o.Filters = &v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ReadApiLogsRequest) GetNextPageToken() string {
	if o == nil || o.NextPageToken == nil {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadApiLogsRequest) GetNextPageTokenOk() (*string, bool) {
	if o == nil || o.NextPageToken == nil {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ReadApiLogsRequest) HasNextPageToken() bool {
	if o != nil && o.NextPageToken != nil {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ReadApiLogsRequest) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

// GetResultsPerPage returns the ResultsPerPage field value if set, zero value otherwise.
func (o *ReadApiLogsRequest) GetResultsPerPage() int32 {
	if o == nil || o.ResultsPerPage == nil {
		var ret int32
		return ret
	}
	return *o.ResultsPerPage
}

// GetResultsPerPageOk returns a tuple with the ResultsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadApiLogsRequest) GetResultsPerPageOk() (*int32, bool) {
	if o == nil || o.ResultsPerPage == nil {
		return nil, false
	}
	return o.ResultsPerPage, true
}

// HasResultsPerPage returns a boolean if a field has been set.
func (o *ReadApiLogsRequest) HasResultsPerPage() bool {
	if o != nil && o.ResultsPerPage != nil {
		return true
	}

	return false
}

// SetResultsPerPage gets a reference to the given int32 and assigns it to the ResultsPerPage field.
func (o *ReadApiLogsRequest) SetResultsPerPage(v int32) {
	o.ResultsPerPage = &v
}

// GetWith returns the With field value if set, zero value otherwise.
func (o *ReadApiLogsRequest) GetWith() With {
	if o == nil || o.With == nil {
		var ret With
		return ret
	}
	return *o.With
}

// GetWithOk returns a tuple with the With field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadApiLogsRequest) GetWithOk() (*With, bool) {
	if o == nil || o.With == nil {
		return nil, false
	}
	return o.With, true
}

// HasWith returns a boolean if a field has been set.
func (o *ReadApiLogsRequest) HasWith() bool {
	if o != nil && o.With != nil {
		return true
	}

	return false
}

// SetWith gets a reference to the given With and assigns it to the With field.
func (o *ReadApiLogsRequest) SetWith(v With) {
	o.With = &v
}

func (o ReadApiLogsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.Filters != nil {
		toSerialize["Filters"] = o.Filters
	}
	if o.NextPageToken != nil {
		toSerialize["NextPageToken"] = o.NextPageToken
	}
	if o.ResultsPerPage != nil {
		toSerialize["ResultsPerPage"] = o.ResultsPerPage
	}
	if o.With != nil {
		toSerialize["With"] = o.With
	}
	return json.Marshal(toSerialize)
}

type NullableReadApiLogsRequest struct {
	value *ReadApiLogsRequest
	isSet bool
}

func (v NullableReadApiLogsRequest) Get() *ReadApiLogsRequest {
	return v.value
}

func (v *NullableReadApiLogsRequest) Set(val *ReadApiLogsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReadApiLogsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReadApiLogsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadApiLogsRequest(val *ReadApiLogsRequest) *NullableReadApiLogsRequest {
	return &NullableReadApiLogsRequest{value: val, isSet: true}
}

func (v NullableReadApiLogsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadApiLogsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


