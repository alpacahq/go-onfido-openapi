# WorkflowRunShared

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

## Methods

### NewWorkflowRunShared

`func NewWorkflowRunShared(applicantId string, workflowId string, ) *WorkflowRunShared`

NewWorkflowRunShared instantiates a new WorkflowRunShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRunSharedWithDefaults

`func NewWorkflowRunSharedWithDefaults() *WorkflowRunShared`

NewWorkflowRunSharedWithDefaults instantiates a new WorkflowRunShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicantId

`func (o *WorkflowRunShared) GetApplicantId() string`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *WorkflowRunShared) GetApplicantIdOk() (*string, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *WorkflowRunShared) SetApplicantId(v string)`

SetApplicantId sets ApplicantId field to given value.


### GetWorkflowId

`func (o *WorkflowRunShared) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowRunShared) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowRunShared) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetTags

`func (o *WorkflowRunShared) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WorkflowRunShared) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WorkflowRunShared) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WorkflowRunShared) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *WorkflowRunShared) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *WorkflowRunShared) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetCustomerUserId

`func (o *WorkflowRunShared) GetCustomerUserId() string`

GetCustomerUserId returns the CustomerUserId field if non-nil, zero value otherwise.

### GetCustomerUserIdOk

`func (o *WorkflowRunShared) GetCustomerUserIdOk() (*string, bool)`

GetCustomerUserIdOk returns a tuple with the CustomerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerUserId

`func (o *WorkflowRunShared) SetCustomerUserId(v string)`

SetCustomerUserId sets CustomerUserId field to given value.

### HasCustomerUserId

`func (o *WorkflowRunShared) HasCustomerUserId() bool`

HasCustomerUserId returns a boolean if a field has been set.

### GetLink

`func (o *WorkflowRunShared) GetLink() WorkflowRunSharedLink`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *WorkflowRunShared) GetLinkOk() (*WorkflowRunSharedLink, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *WorkflowRunShared) SetLink(v WorkflowRunSharedLink)`

SetLink sets Link field to given value.

### HasLink

`func (o *WorkflowRunShared) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WorkflowRunShared) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkflowRunShared) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkflowRunShared) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WorkflowRunShared) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WorkflowRunShared) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkflowRunShared) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkflowRunShared) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WorkflowRunShared) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


