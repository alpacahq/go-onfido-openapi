# WorkflowRunResponseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of error. | [optional] 
**Message** | Pointer to **string** | A textual description of the error. | [optional] 

## Methods

### NewWorkflowRunResponseError

`func NewWorkflowRunResponseError() *WorkflowRunResponseError`

NewWorkflowRunResponseError instantiates a new WorkflowRunResponseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRunResponseErrorWithDefaults

`func NewWorkflowRunResponseErrorWithDefaults() *WorkflowRunResponseError`

NewWorkflowRunResponseErrorWithDefaults instantiates a new WorkflowRunResponseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WorkflowRunResponseError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowRunResponseError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowRunResponseError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowRunResponseError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *WorkflowRunResponseError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WorkflowRunResponseError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WorkflowRunResponseError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *WorkflowRunResponseError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


