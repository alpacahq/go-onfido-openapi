# WorkflowRunBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicantId** | **string** | The unique identifier for the Applicant. | 
**WorkflowId** | **string** | The unique identifier for the Workflow. | 
**Tags** | Pointer to **[]string** | Tags or labels assigned to the workflow run. | [optional] 
**CustomerUserId** | Pointer to **string** | Customer-provided user identifier. | [optional] 
**Link** | Pointer to [**WorkflowRunSharedLink**](WorkflowRunSharedLink.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when the Workflow Run was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time when the Workflow Run was last updated. | [optional] 
**CustomData** | Pointer to **map[string]interface{}** | Object with Custom Input Data to be used in the Workflow Run. | [optional] 

## Methods

### NewWorkflowRunBuilder

`func NewWorkflowRunBuilder(applicantId string, workflowId string, ) *WorkflowRunBuilder`

NewWorkflowRunBuilder instantiates a new WorkflowRunBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRunBuilderWithDefaults

`func NewWorkflowRunBuilderWithDefaults() *WorkflowRunBuilder`

NewWorkflowRunBuilderWithDefaults instantiates a new WorkflowRunBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicantId

`func (o *WorkflowRunBuilder) GetApplicantId() string`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *WorkflowRunBuilder) GetApplicantIdOk() (*string, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *WorkflowRunBuilder) SetApplicantId(v string)`

SetApplicantId sets ApplicantId field to given value.


### GetWorkflowId

`func (o *WorkflowRunBuilder) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowRunBuilder) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowRunBuilder) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetTags

`func (o *WorkflowRunBuilder) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WorkflowRunBuilder) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WorkflowRunBuilder) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WorkflowRunBuilder) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *WorkflowRunBuilder) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *WorkflowRunBuilder) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetCustomerUserId

`func (o *WorkflowRunBuilder) GetCustomerUserId() string`

GetCustomerUserId returns the CustomerUserId field if non-nil, zero value otherwise.

### GetCustomerUserIdOk

`func (o *WorkflowRunBuilder) GetCustomerUserIdOk() (*string, bool)`

GetCustomerUserIdOk returns a tuple with the CustomerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUserId

`func (o *WorkflowRunBuilder) SetCustomerUserId(v string)`

SetCustomerUserId sets CustomerUserId field to given value.

### HasCustomerUserId

`func (o *WorkflowRunBuilder) HasCustomerUserId() bool`

HasCustomerUserId returns a boolean if a field has been set.

### GetLink

`func (o *WorkflowRunBuilder) GetLink() WorkflowRunSharedLink`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *WorkflowRunBuilder) GetLinkOk() (*WorkflowRunSharedLink, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *WorkflowRunBuilder) SetLink(v WorkflowRunSharedLink)`

SetLink sets Link field to given value.

### HasLink

`func (o *WorkflowRunBuilder) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WorkflowRunBuilder) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkflowRunBuilder) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkflowRunBuilder) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WorkflowRunBuilder) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WorkflowRunBuilder) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkflowRunBuilder) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkflowRunBuilder) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WorkflowRunBuilder) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCustomData

`func (o *WorkflowRunBuilder) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *WorkflowRunBuilder) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *WorkflowRunBuilder) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *WorkflowRunBuilder) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


