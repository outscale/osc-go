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

// CreateDirectLinkResponse struct for CreateDirectLinkResponse
type CreateDirectLinkResponse struct {
	DirectLink      *DirectLink      `json:"DirectLink,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewCreateDirectLinkResponse instantiates a new CreateDirectLinkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDirectLinkResponse() *CreateDirectLinkResponse {
	this := CreateDirectLinkResponse{}
	return &this
}

// NewCreateDirectLinkResponseWithDefaults instantiates a new CreateDirectLinkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDirectLinkResponseWithDefaults() *CreateDirectLinkResponse {
	this := CreateDirectLinkResponse{}
	return &this
}

// GetDirectLink returns the DirectLink field value if set, zero value otherwise.
func (o *CreateDirectLinkResponse) GetDirectLink() DirectLink {
	if o == nil || o.DirectLink == nil {
		var ret DirectLink
		return ret
	}
	return *o.DirectLink
}

// GetDirectLinkOk returns a tuple with the DirectLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDirectLinkResponse) GetDirectLinkOk() (*DirectLink, bool) {
	if o == nil || o.DirectLink == nil {
		return nil, false
	}
	return o.DirectLink, true
}

// HasDirectLink returns a boolean if a field has been set.
func (o *CreateDirectLinkResponse) HasDirectLink() bool {
	if o != nil && o.DirectLink != nil {
		return true
	}

	return false
}

// SetDirectLink gets a reference to the given DirectLink and assigns it to the DirectLink field.
func (o *CreateDirectLinkResponse) SetDirectLink(v DirectLink) {
	o.DirectLink = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateDirectLinkResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDirectLinkResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateDirectLinkResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateDirectLinkResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o CreateDirectLinkResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DirectLink != nil {
		toSerialize["DirectLink"] = o.DirectLink
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDirectLinkResponse struct {
	value *CreateDirectLinkResponse
	isSet bool
}

func (v NullableCreateDirectLinkResponse) Get() *CreateDirectLinkResponse {
	return v.value
}

func (v *NullableCreateDirectLinkResponse) Set(val *CreateDirectLinkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDirectLinkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDirectLinkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDirectLinkResponse(val *CreateDirectLinkResponse) *NullableCreateDirectLinkResponse {
	return &NullableCreateDirectLinkResponse{value: val, isSet: true}
}

func (v NullableCreateDirectLinkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDirectLinkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
