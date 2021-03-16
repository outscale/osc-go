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

// ListenerForCreation Information about the listener to create.
type ListenerForCreation struct {
	// The port on which the back-end VM is listening (between `1` and `65535`, both included).
	BackendPort int32 `json:"BackendPort"`
	// The protocol for routing traffic to back-end VMs (`HTTP` \\| `HTTPS` \\| `TCP` \\| `SSL` \\| `UDP`).
	BackendProtocol *string `json:"BackendProtocol,omitempty"`
	// The port on which the load balancer is listening (between `1` and `65535`, both included).
	LoadBalancerPort int32 `json:"LoadBalancerPort"`
	// The routing protocol (`HTTP` \\| `HTTPS` \\| `TCP` \\| `SSL` \\| `UDP`).
	LoadBalancerProtocol string `json:"LoadBalancerProtocol"`
	// The Outscale Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > Outscale Resource Names (ORNs)](https://wiki.outscale.net/display/EN/Resource+Identifiers#ResourceIdentifiers-ORNFormat).
	ServerCertificateId *string `json:"ServerCertificateId,omitempty"`
}

// NewListenerForCreation instantiates a new ListenerForCreation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListenerForCreation(backendPort int32, loadBalancerPort int32, loadBalancerProtocol string) *ListenerForCreation {
	this := ListenerForCreation{}
	this.BackendPort = backendPort
	this.LoadBalancerPort = loadBalancerPort
	this.LoadBalancerProtocol = loadBalancerProtocol
	return &this
}

// NewListenerForCreationWithDefaults instantiates a new ListenerForCreation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListenerForCreationWithDefaults() *ListenerForCreation {
	this := ListenerForCreation{}
	return &this
}

// GetBackendPort returns the BackendPort field value
func (o *ListenerForCreation) GetBackendPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BackendPort
}

// GetBackendPortOk returns a tuple with the BackendPort field value
// and a boolean to check if the value has been set.
func (o *ListenerForCreation) GetBackendPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackendPort, true
}

// SetBackendPort sets field value
func (o *ListenerForCreation) SetBackendPort(v int32) {
	o.BackendPort = v
}

// GetBackendProtocol returns the BackendProtocol field value if set, zero value otherwise.
func (o *ListenerForCreation) GetBackendProtocol() string {
	if o == nil || o.BackendProtocol == nil {
		var ret string
		return ret
	}
	return *o.BackendProtocol
}

// GetBackendProtocolOk returns a tuple with the BackendProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListenerForCreation) GetBackendProtocolOk() (*string, bool) {
	if o == nil || o.BackendProtocol == nil {
		return nil, false
	}
	return o.BackendProtocol, true
}

// HasBackendProtocol returns a boolean if a field has been set.
func (o *ListenerForCreation) HasBackendProtocol() bool {
	if o != nil && o.BackendProtocol != nil {
		return true
	}

	return false
}

// SetBackendProtocol gets a reference to the given string and assigns it to the BackendProtocol field.
func (o *ListenerForCreation) SetBackendProtocol(v string) {
	o.BackendProtocol = &v
}

// GetLoadBalancerPort returns the LoadBalancerPort field value
func (o *ListenerForCreation) GetLoadBalancerPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LoadBalancerPort
}

// GetLoadBalancerPortOk returns a tuple with the LoadBalancerPort field value
// and a boolean to check if the value has been set.
func (o *ListenerForCreation) GetLoadBalancerPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadBalancerPort, true
}

// SetLoadBalancerPort sets field value
func (o *ListenerForCreation) SetLoadBalancerPort(v int32) {
	o.LoadBalancerPort = v
}

// GetLoadBalancerProtocol returns the LoadBalancerProtocol field value
func (o *ListenerForCreation) GetLoadBalancerProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoadBalancerProtocol
}

// GetLoadBalancerProtocolOk returns a tuple with the LoadBalancerProtocol field value
// and a boolean to check if the value has been set.
func (o *ListenerForCreation) GetLoadBalancerProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadBalancerProtocol, true
}

// SetLoadBalancerProtocol sets field value
func (o *ListenerForCreation) SetLoadBalancerProtocol(v string) {
	o.LoadBalancerProtocol = v
}

// GetServerCertificateId returns the ServerCertificateId field value if set, zero value otherwise.
func (o *ListenerForCreation) GetServerCertificateId() string {
	if o == nil || o.ServerCertificateId == nil {
		var ret string
		return ret
	}
	return *o.ServerCertificateId
}

// GetServerCertificateIdOk returns a tuple with the ServerCertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListenerForCreation) GetServerCertificateIdOk() (*string, bool) {
	if o == nil || o.ServerCertificateId == nil {
		return nil, false
	}
	return o.ServerCertificateId, true
}

// HasServerCertificateId returns a boolean if a field has been set.
func (o *ListenerForCreation) HasServerCertificateId() bool {
	if o != nil && o.ServerCertificateId != nil {
		return true
	}

	return false
}

// SetServerCertificateId gets a reference to the given string and assigns it to the ServerCertificateId field.
func (o *ListenerForCreation) SetServerCertificateId(v string) {
	o.ServerCertificateId = &v
}

func (o ListenerForCreation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["BackendPort"] = o.BackendPort
	}
	if o.BackendProtocol != nil {
		toSerialize["BackendProtocol"] = o.BackendProtocol
	}
	if true {
		toSerialize["LoadBalancerPort"] = o.LoadBalancerPort
	}
	if true {
		toSerialize["LoadBalancerProtocol"] = o.LoadBalancerProtocol
	}
	if o.ServerCertificateId != nil {
		toSerialize["ServerCertificateId"] = o.ServerCertificateId
	}
	return json.Marshal(toSerialize)
}

type NullableListenerForCreation struct {
	value *ListenerForCreation
	isSet bool
}

func (v NullableListenerForCreation) Get() *ListenerForCreation {
	return v.value
}

func (v *NullableListenerForCreation) Set(val *ListenerForCreation) {
	v.value = val
	v.isSet = true
}

func (v NullableListenerForCreation) IsSet() bool {
	return v.isSet
}

func (v *NullableListenerForCreation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListenerForCreation(val *ListenerForCreation) *NullableListenerForCreation {
	return &NullableListenerForCreation{value: val, isSet: true}
}

func (v NullableListenerForCreation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListenerForCreation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
