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

// Log Information about the log.
type Log struct {
	// The account ID.
	AccountId *string `json:"AccountId,omitempty"`
	// The duration of the call (in milliseconds).
	CallDuration *int32 `json:"CallDuration,omitempty"`
	// The API key used for the query.
	QueryAccessKey *string `json:"QueryAccessKey,omitempty"`
	// The name of the API service used by the call.
	QueryApiName *string `json:"QueryApiName,omitempty"`
	// The version of the API service used by the call.
	QueryApiVersion *string `json:"QueryApiVersion,omitempty"`
	// The name of the call.
	QueryCallName *string `json:"QueryCallName,omitempty"`
	// The date of the query, in ISO 8601 format.
	QueryDate *string `json:"QueryDate,omitempty"`
	// The query header raw.
	QueryHeaderRaw *string `json:"QueryHeaderRaw,omitempty"`
	// The query header size.
	QueryHeaderSize *int32 `json:"QueryHeaderSize,omitempty"`
	// The IP address used for the query.
	QueryIpAddress *string `json:"QueryIpAddress,omitempty"`
	// The query payload raw.
	QueryPayloadRaw *string `json:"QueryPayloadRaw,omitempty"`
	// The query payload size.
	QueryPayloadSize *int32 `json:"QueryPayloadSize,omitempty"`
	// The user agent used for the HTTP request.
	QueryUserAgent *string `json:"QueryUserAgent,omitempty"`
	// The ID provided in the response.
	RequestId *string `json:"RequestId,omitempty"`
	// The size of the response (in bytes).
	ResponseSize *int32 `json:"ResponseSize,omitempty"`
	// The HTTP code provided in the response.
	ResponseStatusCode *int32 `json:"ResponseStatusCode,omitempty"`
}

// NewLog instantiates a new Log object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLog() *Log {
	this := Log{}
	return &this
}

// NewLogWithDefaults instantiates a new Log object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogWithDefaults() *Log {
	this := Log{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *Log) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *Log) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *Log) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCallDuration returns the CallDuration field value if set, zero value otherwise.
func (o *Log) GetCallDuration() int32 {
	if o == nil || o.CallDuration == nil {
		var ret int32
		return ret
	}
	return *o.CallDuration
}

// GetCallDurationOk returns a tuple with the CallDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetCallDurationOk() (*int32, bool) {
	if o == nil || o.CallDuration == nil {
		return nil, false
	}
	return o.CallDuration, true
}

// HasCallDuration returns a boolean if a field has been set.
func (o *Log) HasCallDuration() bool {
	if o != nil && o.CallDuration != nil {
		return true
	}

	return false
}

// SetCallDuration gets a reference to the given int32 and assigns it to the CallDuration field.
func (o *Log) SetCallDuration(v int32) {
	o.CallDuration = &v
}

// GetQueryAccessKey returns the QueryAccessKey field value if set, zero value otherwise.
func (o *Log) GetQueryAccessKey() string {
	if o == nil || o.QueryAccessKey == nil {
		var ret string
		return ret
	}
	return *o.QueryAccessKey
}

// GetQueryAccessKeyOk returns a tuple with the QueryAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetQueryAccessKeyOk() (*string, bool) {
	if o == nil || o.QueryAccessKey == nil {
		return nil, false
	}
	return o.QueryAccessKey, true
}

// HasQueryAccessKey returns a boolean if a field has been set.
func (o *Log) HasQueryAccessKey() bool {
	if o != nil && o.QueryAccessKey != nil {
		return true
	}

	return false
}

// SetQueryAccessKey gets a reference to the given string and assigns it to the QueryAccessKey field.
func (o *Log) SetQueryAccessKey(v string) {
	o.QueryAccessKey = &v
}

// GetQueryApiName returns the QueryApiName field value if set, zero value otherwise.
func (o *Log) GetQueryApiName() string {
	if o == nil || o.QueryApiName == nil {
		var ret string
		return ret
	}
	return *o.QueryApiName
}

// GetQueryApiNameOk returns a tuple with the QueryApiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetQueryApiNameOk() (*string, bool) {
	if o == nil || o.QueryApiName == nil {
		return nil, false
	}
	return o.QueryApiName, true
}

// HasQueryApiName returns a boolean if a field has been set.
func (o *Log) HasQueryApiName() bool {
	if o != nil && o.QueryApiName != nil {
		return true
	}

	return false
}

// SetQueryApiName gets a reference to the given string and assigns it to the QueryApiName field.
func (o *Log) SetQueryApiName(v string) {
	o.QueryApiName = &v
}

// GetQueryApiVersion returns the QueryApiVersion field value if set, zero value otherwise.
func (o *Log) GetQueryApiVersion() string {
	if o == nil || o.QueryApiVersion == nil {
		var ret string
		return ret
	}
	return *o.QueryApiVersion
}

// GetQueryApiVersionOk returns a tuple with the QueryApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetQueryApiVersionOk() (*string, bool) {
	if o == nil || o.QueryApiVersion == nil {
		return nil, false
	}
	return o.QueryApiVersion, true
}

// HasQueryApiVersion returns a boolean if a field has been set.
func (o *Log) HasQueryApiVersion() bool {
	if o != nil && o.QueryApiVersion != nil {
		return true
	}

	return false
}

// SetQueryApiVersion gets a reference to the given string and assigns it to the QueryApiVersion field.
func (o *Log) SetQueryApiVersion(v string) {
	o.QueryApiVersion = &v
}

// GetQueryCallName returns the QueryCallName field value if set, zero value otherwise.
func (o *Log) GetQueryCallName() string {
	if o == nil || o.QueryCallName == nil {
		var ret string
		return ret
	}
	return *o.QueryCallName
}

// GetQueryCallNameOk returns a tuple with the QueryCallName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetQueryCallNameOk() (*string, bool) {
	if o == nil || o.QueryCallName == nil {
		return nil, false
	}
	return o.QueryCallName, true
}

// HasQueryCallName returns a boolean if a field has been set.
func (o *Log) HasQueryCallName() bool {
	if o != nil && o.QueryCallName != nil {
		return true
	}

	return false
}

// SetQueryCallName gets a reference to the given string and assigns it to the QueryCallName field.
func (o *Log) SetQueryCallName(v string) {
	o.QueryCallName = &v
}

// GetQueryDate returns the QueryDate field value if set, zero value otherwise.
func (o *Log) GetQueryDate() string {
	if o == nil || o.QueryDate == nil {
		var ret string
		return ret
	}
	return *o.QueryDate
}

// GetQueryDateOk returns a tuple with the QueryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetQueryDateOk() (*string, bool) {
	if o == nil || o.QueryDate == nil {
		return nil, false
	}
	return o.QueryDate, true
}

// HasQueryDate returns a boolean if a field has been set.
func (o *Log) HasQueryDate() bool {
	if o != nil && o.QueryDate != nil {
		return true
	}

	return false
}

// SetQueryDate gets a reference to the given string and assigns it to the QueryDate field.
func (o *Log) SetQueryDate(v string) {
	o.QueryDate = &v
}

// GetQueryHeaderRaw returns the QueryHeaderRaw field value if set, zero value otherwise.
func (o *Log) GetQueryHeaderRaw() string {
	if o == nil || o.QueryHeaderRaw == nil {
		var ret string
		return ret
	}
	return *o.QueryHeaderRaw
}

// GetQueryHeaderRawOk returns a tuple with the QueryHeaderRaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetQueryHeaderRawOk() (*string, bool) {
	if o == nil || o.QueryHeaderRaw == nil {
		return nil, false
	}
	return o.QueryHeaderRaw, true
}

// HasQueryHeaderRaw returns a boolean if a field has been set.
func (o *Log) HasQueryHeaderRaw() bool {
	if o != nil && o.QueryHeaderRaw != nil {
		return true
	}

	return false
}

// SetQueryHeaderRaw gets a reference to the given string and assigns it to the QueryHeaderRaw field.
func (o *Log) SetQueryHeaderRaw(v string) {
	o.QueryHeaderRaw = &v
}

// GetQueryHeaderSize returns the QueryHeaderSize field value if set, zero value otherwise.
func (o *Log) GetQueryHeaderSize() int32 {
	if o == nil || o.QueryHeaderSize == nil {
		var ret int32
		return ret
	}
	return *o.QueryHeaderSize
}

// GetQueryHeaderSizeOk returns a tuple with the QueryHeaderSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetQueryHeaderSizeOk() (*int32, bool) {
	if o == nil || o.QueryHeaderSize == nil {
		return nil, false
	}
	return o.QueryHeaderSize, true
}

// HasQueryHeaderSize returns a boolean if a field has been set.
func (o *Log) HasQueryHeaderSize() bool {
	if o != nil && o.QueryHeaderSize != nil {
		return true
	}

	return false
}

// SetQueryHeaderSize gets a reference to the given int32 and assigns it to the QueryHeaderSize field.
func (o *Log) SetQueryHeaderSize(v int32) {
	o.QueryHeaderSize = &v
}

// GetQueryIpAddress returns the QueryIpAddress field value if set, zero value otherwise.
func (o *Log) GetQueryIpAddress() string {
	if o == nil || o.QueryIpAddress == nil {
		var ret string
		return ret
	}
	return *o.QueryIpAddress
}

// GetQueryIpAddressOk returns a tuple with the QueryIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetQueryIpAddressOk() (*string, bool) {
	if o == nil || o.QueryIpAddress == nil {
		return nil, false
	}
	return o.QueryIpAddress, true
}

// HasQueryIpAddress returns a boolean if a field has been set.
func (o *Log) HasQueryIpAddress() bool {
	if o != nil && o.QueryIpAddress != nil {
		return true
	}

	return false
}

// SetQueryIpAddress gets a reference to the given string and assigns it to the QueryIpAddress field.
func (o *Log) SetQueryIpAddress(v string) {
	o.QueryIpAddress = &v
}

// GetQueryPayloadRaw returns the QueryPayloadRaw field value if set, zero value otherwise.
func (o *Log) GetQueryPayloadRaw() string {
	if o == nil || o.QueryPayloadRaw == nil {
		var ret string
		return ret
	}
	return *o.QueryPayloadRaw
}

// GetQueryPayloadRawOk returns a tuple with the QueryPayloadRaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetQueryPayloadRawOk() (*string, bool) {
	if o == nil || o.QueryPayloadRaw == nil {
		return nil, false
	}
	return o.QueryPayloadRaw, true
}

// HasQueryPayloadRaw returns a boolean if a field has been set.
func (o *Log) HasQueryPayloadRaw() bool {
	if o != nil && o.QueryPayloadRaw != nil {
		return true
	}

	return false
}

// SetQueryPayloadRaw gets a reference to the given string and assigns it to the QueryPayloadRaw field.
func (o *Log) SetQueryPayloadRaw(v string) {
	o.QueryPayloadRaw = &v
}

// GetQueryPayloadSize returns the QueryPayloadSize field value if set, zero value otherwise.
func (o *Log) GetQueryPayloadSize() int32 {
	if o == nil || o.QueryPayloadSize == nil {
		var ret int32
		return ret
	}
	return *o.QueryPayloadSize
}

// GetQueryPayloadSizeOk returns a tuple with the QueryPayloadSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetQueryPayloadSizeOk() (*int32, bool) {
	if o == nil || o.QueryPayloadSize == nil {
		return nil, false
	}
	return o.QueryPayloadSize, true
}

// HasQueryPayloadSize returns a boolean if a field has been set.
func (o *Log) HasQueryPayloadSize() bool {
	if o != nil && o.QueryPayloadSize != nil {
		return true
	}

	return false
}

// SetQueryPayloadSize gets a reference to the given int32 and assigns it to the QueryPayloadSize field.
func (o *Log) SetQueryPayloadSize(v int32) {
	o.QueryPayloadSize = &v
}

// GetQueryUserAgent returns the QueryUserAgent field value if set, zero value otherwise.
func (o *Log) GetQueryUserAgent() string {
	if o == nil || o.QueryUserAgent == nil {
		var ret string
		return ret
	}
	return *o.QueryUserAgent
}

// GetQueryUserAgentOk returns a tuple with the QueryUserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetQueryUserAgentOk() (*string, bool) {
	if o == nil || o.QueryUserAgent == nil {
		return nil, false
	}
	return o.QueryUserAgent, true
}

// HasQueryUserAgent returns a boolean if a field has been set.
func (o *Log) HasQueryUserAgent() bool {
	if o != nil && o.QueryUserAgent != nil {
		return true
	}

	return false
}

// SetQueryUserAgent gets a reference to the given string and assigns it to the QueryUserAgent field.
func (o *Log) SetQueryUserAgent(v string) {
	o.QueryUserAgent = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *Log) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *Log) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *Log) SetRequestId(v string) {
	o.RequestId = &v
}

// GetResponseSize returns the ResponseSize field value if set, zero value otherwise.
func (o *Log) GetResponseSize() int32 {
	if o == nil || o.ResponseSize == nil {
		var ret int32
		return ret
	}
	return *o.ResponseSize
}

// GetResponseSizeOk returns a tuple with the ResponseSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetResponseSizeOk() (*int32, bool) {
	if o == nil || o.ResponseSize == nil {
		return nil, false
	}
	return o.ResponseSize, true
}

// HasResponseSize returns a boolean if a field has been set.
func (o *Log) HasResponseSize() bool {
	if o != nil && o.ResponseSize != nil {
		return true
	}

	return false
}

// SetResponseSize gets a reference to the given int32 and assigns it to the ResponseSize field.
func (o *Log) SetResponseSize(v int32) {
	o.ResponseSize = &v
}

// GetResponseStatusCode returns the ResponseStatusCode field value if set, zero value otherwise.
func (o *Log) GetResponseStatusCode() int32 {
	if o == nil || o.ResponseStatusCode == nil {
		var ret int32
		return ret
	}
	return *o.ResponseStatusCode
}

// GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Log) GetResponseStatusCodeOk() (*int32, bool) {
	if o == nil || o.ResponseStatusCode == nil {
		return nil, false
	}
	return o.ResponseStatusCode, true
}

// HasResponseStatusCode returns a boolean if a field has been set.
func (o *Log) HasResponseStatusCode() bool {
	if o != nil && o.ResponseStatusCode != nil {
		return true
	}

	return false
}

// SetResponseStatusCode gets a reference to the given int32 and assigns it to the ResponseStatusCode field.
func (o *Log) SetResponseStatusCode(v int32) {
	o.ResponseStatusCode = &v
}

func (o Log) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["AccountId"] = o.AccountId
	}
	if o.CallDuration != nil {
		toSerialize["CallDuration"] = o.CallDuration
	}
	if o.QueryAccessKey != nil {
		toSerialize["QueryAccessKey"] = o.QueryAccessKey
	}
	if o.QueryApiName != nil {
		toSerialize["QueryApiName"] = o.QueryApiName
	}
	if o.QueryApiVersion != nil {
		toSerialize["QueryApiVersion"] = o.QueryApiVersion
	}
	if o.QueryCallName != nil {
		toSerialize["QueryCallName"] = o.QueryCallName
	}
	if o.QueryDate != nil {
		toSerialize["QueryDate"] = o.QueryDate
	}
	if o.QueryHeaderRaw != nil {
		toSerialize["QueryHeaderRaw"] = o.QueryHeaderRaw
	}
	if o.QueryHeaderSize != nil {
		toSerialize["QueryHeaderSize"] = o.QueryHeaderSize
	}
	if o.QueryIpAddress != nil {
		toSerialize["QueryIpAddress"] = o.QueryIpAddress
	}
	if o.QueryPayloadRaw != nil {
		toSerialize["QueryPayloadRaw"] = o.QueryPayloadRaw
	}
	if o.QueryPayloadSize != nil {
		toSerialize["QueryPayloadSize"] = o.QueryPayloadSize
	}
	if o.QueryUserAgent != nil {
		toSerialize["QueryUserAgent"] = o.QueryUserAgent
	}
	if o.RequestId != nil {
		toSerialize["RequestId"] = o.RequestId
	}
	if o.ResponseSize != nil {
		toSerialize["ResponseSize"] = o.ResponseSize
	}
	if o.ResponseStatusCode != nil {
		toSerialize["ResponseStatusCode"] = o.ResponseStatusCode
	}
	return json.Marshal(toSerialize)
}

type NullableLog struct {
	value *Log
	isSet bool
}

func (v NullableLog) Get() *Log {
	return v.value
}

func (v *NullableLog) Set(val *Log) {
	v.value = val
	v.isSet = true
}

func (v NullableLog) IsSet() bool {
	return v.isSet
}

func (v *NullableLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLog(val *Log) *NullableLog {
	return &NullableLog{value: val, isSet: true}
}

func (v NullableLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
