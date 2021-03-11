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

// CreateRouteRequest struct for CreateRouteRequest
type CreateRouteRequest struct {
	// The IP range used for the destination match, in CIDR notation (for example, 10.0.0.0/24).
	DestinationIpRange string `json:"DestinationIpRange"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of an Internet service or virtual gateway attached to your Net.
	GatewayId *string `json:"GatewayId,omitempty"`
	// The ID of a NAT service.
	NatServiceId *string `json:"NatServiceId,omitempty"`
	// The ID of a Net peering connection.
	NetPeeringId *string `json:"NetPeeringId,omitempty"`
	// The ID of a NIC.
	NicId *string `json:"NicId,omitempty"`
	// The ID of the route table for which you want to create a route.
	RouteTableId string `json:"RouteTableId"`
	// The ID of a NAT VM in your Net (attached to exactly one NIC).
	VmId *string `json:"VmId,omitempty"`
}

// NewCreateRouteRequest instantiates a new CreateRouteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRouteRequest(destinationIpRange string, routeTableId string) *CreateRouteRequest {
	this := CreateRouteRequest{}
	this.DestinationIpRange = destinationIpRange
	this.RouteTableId = routeTableId
	return &this
}

// NewCreateRouteRequestWithDefaults instantiates a new CreateRouteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRouteRequestWithDefaults() *CreateRouteRequest {
	this := CreateRouteRequest{}
	return &this
}

// GetDestinationIpRange returns the DestinationIpRange field value
func (o *CreateRouteRequest) GetDestinationIpRange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationIpRange
}

// GetDestinationIpRangeOk returns a tuple with the DestinationIpRange field value
// and a boolean to check if the value has been set.
func (o *CreateRouteRequest) GetDestinationIpRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationIpRange, true
}

// SetDestinationIpRange sets field value
func (o *CreateRouteRequest) SetDestinationIpRange(v string) {
	o.DestinationIpRange = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateRouteRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateRouteRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateRouteRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *CreateRouteRequest) GetGatewayId() string {
	if o == nil || o.GatewayId == nil {
		var ret string
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteRequest) GetGatewayIdOk() (*string, bool) {
	if o == nil || o.GatewayId == nil {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *CreateRouteRequest) HasGatewayId() bool {
	if o != nil && o.GatewayId != nil {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given string and assigns it to the GatewayId field.
func (o *CreateRouteRequest) SetGatewayId(v string) {
	o.GatewayId = &v
}

// GetNatServiceId returns the NatServiceId field value if set, zero value otherwise.
func (o *CreateRouteRequest) GetNatServiceId() string {
	if o == nil || o.NatServiceId == nil {
		var ret string
		return ret
	}
	return *o.NatServiceId
}

// GetNatServiceIdOk returns a tuple with the NatServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteRequest) GetNatServiceIdOk() (*string, bool) {
	if o == nil || o.NatServiceId == nil {
		return nil, false
	}
	return o.NatServiceId, true
}

// HasNatServiceId returns a boolean if a field has been set.
func (o *CreateRouteRequest) HasNatServiceId() bool {
	if o != nil && o.NatServiceId != nil {
		return true
	}

	return false
}

// SetNatServiceId gets a reference to the given string and assigns it to the NatServiceId field.
func (o *CreateRouteRequest) SetNatServiceId(v string) {
	o.NatServiceId = &v
}

// GetNetPeeringId returns the NetPeeringId field value if set, zero value otherwise.
func (o *CreateRouteRequest) GetNetPeeringId() string {
	if o == nil || o.NetPeeringId == nil {
		var ret string
		return ret
	}
	return *o.NetPeeringId
}

// GetNetPeeringIdOk returns a tuple with the NetPeeringId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteRequest) GetNetPeeringIdOk() (*string, bool) {
	if o == nil || o.NetPeeringId == nil {
		return nil, false
	}
	return o.NetPeeringId, true
}

// HasNetPeeringId returns a boolean if a field has been set.
func (o *CreateRouteRequest) HasNetPeeringId() bool {
	if o != nil && o.NetPeeringId != nil {
		return true
	}

	return false
}

// SetNetPeeringId gets a reference to the given string and assigns it to the NetPeeringId field.
func (o *CreateRouteRequest) SetNetPeeringId(v string) {
	o.NetPeeringId = &v
}

// GetNicId returns the NicId field value if set, zero value otherwise.
func (o *CreateRouteRequest) GetNicId() string {
	if o == nil || o.NicId == nil {
		var ret string
		return ret
	}
	return *o.NicId
}

// GetNicIdOk returns a tuple with the NicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteRequest) GetNicIdOk() (*string, bool) {
	if o == nil || o.NicId == nil {
		return nil, false
	}
	return o.NicId, true
}

// HasNicId returns a boolean if a field has been set.
func (o *CreateRouteRequest) HasNicId() bool {
	if o != nil && o.NicId != nil {
		return true
	}

	return false
}

// SetNicId gets a reference to the given string and assigns it to the NicId field.
func (o *CreateRouteRequest) SetNicId(v string) {
	o.NicId = &v
}

// GetRouteTableId returns the RouteTableId field value
func (o *CreateRouteRequest) GetRouteTableId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RouteTableId
}

// GetRouteTableIdOk returns a tuple with the RouteTableId field value
// and a boolean to check if the value has been set.
func (o *CreateRouteRequest) GetRouteTableIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RouteTableId, true
}

// SetRouteTableId sets field value
func (o *CreateRouteRequest) SetRouteTableId(v string) {
	o.RouteTableId = v
}

// GetVmId returns the VmId field value if set, zero value otherwise.
func (o *CreateRouteRequest) GetVmId() string {
	if o == nil || o.VmId == nil {
		var ret string
		return ret
	}
	return *o.VmId
}

// GetVmIdOk returns a tuple with the VmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRouteRequest) GetVmIdOk() (*string, bool) {
	if o == nil || o.VmId == nil {
		return nil, false
	}
	return o.VmId, true
}

// HasVmId returns a boolean if a field has been set.
func (o *CreateRouteRequest) HasVmId() bool {
	if o != nil && o.VmId != nil {
		return true
	}

	return false
}

// SetVmId gets a reference to the given string and assigns it to the VmId field.
func (o *CreateRouteRequest) SetVmId(v string) {
	o.VmId = &v
}

func (o CreateRouteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["DestinationIpRange"] = o.DestinationIpRange
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if o.GatewayId != nil {
		toSerialize["GatewayId"] = o.GatewayId
	}
	if o.NatServiceId != nil {
		toSerialize["NatServiceId"] = o.NatServiceId
	}
	if o.NetPeeringId != nil {
		toSerialize["NetPeeringId"] = o.NetPeeringId
	}
	if o.NicId != nil {
		toSerialize["NicId"] = o.NicId
	}
	if true {
		toSerialize["RouteTableId"] = o.RouteTableId
	}
	if o.VmId != nil {
		toSerialize["VmId"] = o.VmId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateRouteRequest struct {
	value *CreateRouteRequest
	isSet bool
}

func (v NullableCreateRouteRequest) Get() *CreateRouteRequest {
	return v.value
}

func (v *NullableCreateRouteRequest) Set(val *CreateRouteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRouteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRouteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRouteRequest(val *CreateRouteRequest) *NullableCreateRouteRequest {
	return &NullableCreateRouteRequest{value: val, isSet: true}
}

func (v NullableCreateRouteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRouteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
