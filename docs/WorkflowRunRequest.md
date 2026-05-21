# WorkflowRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomData** | Pointer to **map[string]interface{}** | Object with Custom Input Data to be used in the Workflow Run. | [optional] 

## Methods

### NewWorkflowRunRequest

`func NewWorkflowRunRequest() *WorkflowRunRequest`

NewWorkflowRunRequest instantiates a new WorkflowRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRunRequestWithDefaults

`func NewWorkflowRunRequestWithDefaults() *WorkflowRunRequest`

NewWorkflowRunRequestWithDefaults instantiates a new WorkflowRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomData

`func (o *WorkflowRunRequest) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *WorkflowRunRequest) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *WorkflowRunRequest) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *WorkflowRunRequest) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


