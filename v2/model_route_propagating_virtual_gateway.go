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

// RoutePropagatingVirtualGateway Information about the route propagating virtual gateway.
type RoutePropagatingVirtualGateway struct {
	// The ID of the virtual gateway.
	VirtualGatewayId *string `json:"VirtualGatewayId,omitempty"`
}

// NewRoutePropagatingVirtualGateway instantiates a new RoutePropagatingVirtualGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutePropagatingVirtualGateway() *RoutePropagatingVirtualGateway {
	this := RoutePropagatingVirtualGateway{}
	return &this
}

// NewRoutePropagatingVirtualGatewayWithDefaults instantiates a new RoutePropagatingVirtualGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutePropagatingVirtualGatewayWithDefaults() *RoutePropagatingVirtualGateway {
	this := RoutePropagatingVirtualGateway{}
	return &this
}

// GetVirtualGatewayId returns the VirtualGatewayId field value if set, zero value otherwise.
func (o *RoutePropagatingVirtualGateway) GetVirtualGatewayId() string {
	if o == nil || o.VirtualGatewayId == nil {
		var ret string
		return ret
	}
	return *o.VirtualGatewayId
}

// GetVirtualGatewayIdOk returns a tuple with the VirtualGatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutePropagatingVirtualGateway) GetVirtualGatewayIdOk() (*string, bool) {
	if o == nil || o.VirtualGatewayId == nil {
		return nil, false
	}
	return o.VirtualGatewayId, true
}

// HasVirtualGatewayId returns a boolean if a field has been set.
func (o *RoutePropagatingVirtualGateway) HasVirtualGatewayId() bool {
	if o != nil && o.VirtualGatewayId != nil {
		return true
	}

	return false
}

// SetVirtualGatewayId gets a reference to the given string and assigns it to the VirtualGatewayId field.
func (o *RoutePropagatingVirtualGateway) SetVirtualGatewayId(v string) {
	o.VirtualGatewayId = &v
}

func (o RoutePropagatingVirtualGateway) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VirtualGatewayId != nil {
		toSerialize["VirtualGatewayId"] = o.VirtualGatewayId
	}
	return json.Marshal(toSerialize)
}

type NullableRoutePropagatingVirtualGateway struct {
	value *RoutePropagatingVirtualGateway
	isSet bool
}

func (v NullableRoutePropagatingVirtualGateway) Get() *RoutePropagatingVirtualGateway {
	return v.value
}

func (v *NullableRoutePropagatingVirtualGateway) Set(val *RoutePropagatingVirtualGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutePropagatingVirtualGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutePropagatingVirtualGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutePropagatingVirtualGateway(val *RoutePropagatingVirtualGateway) *NullableRoutePropagatingVirtualGateway {
	return &NullableRoutePropagatingVirtualGateway{value: val, isSet: true}
}

func (v NullableRoutePropagatingVirtualGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutePropagatingVirtualGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
