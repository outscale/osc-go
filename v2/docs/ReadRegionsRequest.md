# ReadRegionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | If true, checks whether you have the required permissions to perform the action. | [optional] 

## Methods

### NewReadRegionsRequest

`func NewReadRegionsRequest() *ReadRegionsRequest`

NewReadRegionsRequest instantiates a new ReadRegionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadRegionsRequestWithDefaults

`func NewReadRegionsRequestWithDefaults() *ReadRegionsRequest`

NewReadRegionsRequestWithDefaults instantiates a new ReadRegionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *ReadRegionsRequest) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *ReadRegionsRequest) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *ReadRegionsRequest) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *ReadRegionsRequest) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


