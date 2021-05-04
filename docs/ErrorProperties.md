# ErrorProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewErrorProperties

`func NewErrorProperties() *ErrorProperties`

NewErrorProperties instantiates a new ErrorProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorPropertiesWithDefaults

`func NewErrorPropertiesWithDefaults() *ErrorProperties`

NewErrorPropertiesWithDefaults instantiates a new ErrorProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ErrorProperties) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorProperties) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorProperties) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ErrorProperties) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorProperties) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorProperties) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorProperties) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorProperties) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFields

`func (o *ErrorProperties) GetFields() map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ErrorProperties) GetFieldsOk() (*map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ErrorProperties) SetFields(v map[string]interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *ErrorProperties) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


