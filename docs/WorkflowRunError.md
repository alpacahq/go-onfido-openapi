# WorkflowRunError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of error. | [optional] 
**Message** | Pointer to **string** | A textual description of the error. | [optional] 

## Methods

### NewWorkflowRunError

`func NewWorkflowRunError() *WorkflowRunError`

NewWorkflowRunError instantiates a new WorkflowRunError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRunErrorWithDefaults

`func NewWorkflowRunErrorWithDefaults() *WorkflowRunError`

NewWorkflowRunErrorWithDefaults instantiates a new WorkflowRunError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WorkflowRunError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowRunError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowRunError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowRunError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *WorkflowRunError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WorkflowRunError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WorkflowRunError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *WorkflowRunError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


