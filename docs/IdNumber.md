# IdNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type of ID number. Valid values are &#x60;ssn&#x60;, &#x60;social_insurance&#x60;, &#x60;tax_id&#x60;, &#x60;identity_card&#x60;, &#x60;passport&#x60; and &#x60;driving_license&#x60; | [optional] 
**Value** | Pointer to **string** | Value of ID number | [optional] 
**StateCode** | Pointer to **string** | Two letter code of issuing state (state-issued driving licenses only) | [optional] 

## Methods

### NewIdNumber

`func NewIdNumber() *IdNumber`

NewIdNumber instantiates a new IdNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdNumberWithDefaults

`func NewIdNumberWithDefaults() *IdNumber`

NewIdNumberWithDefaults instantiates a new IdNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IdNumber) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdNumber) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdNumber) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IdNumber) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *IdNumber) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IdNumber) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IdNumber) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IdNumber) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetStateCode

`func (o *IdNumber) GetStateCode() string`

GetStateCode returns the StateCode field if non-nil, zero value otherwise.

### GetStateCodeOk

`func (o *IdNumber) GetStateCodeOk() (*string, bool)`

GetStateCodeOk returns a tuple with the StateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCode

`func (o *IdNumber) SetStateCode(v string)`

SetStateCode sets StateCode field to given value.

### HasStateCode

`func (o *IdNumber) HasStateCode() bool`

HasStateCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


