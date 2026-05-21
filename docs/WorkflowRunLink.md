# WorkflowRunLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | Link to access the Workflow Run without the need to integrate with Onfido&#39;s SDKs. | [optional] 
**CompletedRedirectUrl** | Pointer to **NullableString** | When the interactive section of the Workflow Run has completed successfully, the user will be redirected to this URL instead of seeing the default Onfido &#39;thank you&#39; page. | [optional] 
**ExpiredRedirectUrl** | Pointer to **NullableString** | When the link has expired, the user will be immediately redirected to this URL instead of seeing the default Onfido error message. | [optional] 
**ExpiresAt** | Pointer to **NullableTime** | Date and time when the link will expire. | [optional] 
**Language** | Pointer to **NullableString** | The code for the language when the workflow run is acessed using the link. | [optional] 

## Methods

### NewWorkflowRunLink

`func NewWorkflowRunLink() *WorkflowRunLink`

NewWorkflowRunLink instantiates a new WorkflowRunLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRunLinkWithDefaults

`func NewWorkflowRunLinkWithDefaults() *WorkflowRunLink`

NewWorkflowRunLinkWithDefaults instantiates a new WorkflowRunLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *WorkflowRunLink) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WorkflowRunLink) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WorkflowRunLink) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WorkflowRunLink) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCompletedRedirectUrl

`func (o *WorkflowRunLink) GetCompletedRedirectUrl() string`

GetCompletedRedirectUrl returns the CompletedRedirectUrl field if non-nil, zero value otherwise.

### GetCompletedRedirectUrlOk

`func (o *WorkflowRunLink) GetCompletedRedirectUrlOk() (*string, bool)`

GetCompletedRedirectUrlOk returns a tuple with the CompletedRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedRedirectUrl

`func (o *WorkflowRunLink) SetCompletedRedirectUrl(v string)`

SetCompletedRedirectUrl sets CompletedRedirectUrl field to given value.

### HasCompletedRedirectUrl

`func (o *WorkflowRunLink) HasCompletedRedirectUrl() bool`

HasCompletedRedirectUrl returns a boolean if a field has been set.

### SetCompletedRedirectUrlNil

`func (o *WorkflowRunLink) SetCompletedRedirectUrlNil(b bool)`

 SetCompletedRedirectUrlNil sets the value for CompletedRedirectUrl to be an explicit nil

### UnsetCompletedRedirectUrl
`func (o *WorkflowRunLink) UnsetCompletedRedirectUrl()`

UnsetCompletedRedirectUrl ensures that no value is present for CompletedRedirectUrl, not even an explicit nil
### GetExpiredRedirectUrl

`func (o *WorkflowRunLink) GetExpiredRedirectUrl() string`

GetExpiredRedirectUrl returns the ExpiredRedirectUrl field if non-nil, zero value otherwise.

### GetExpiredRedirectUrlOk

`func (o *WorkflowRunLink) GetExpiredRedirectUrlOk() (*string, bool)`

GetExpiredRedirectUrlOk returns a tuple with the ExpiredRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredRedirectUrl

`func (o *WorkflowRunLink) SetExpiredRedirectUrl(v string)`

SetExpiredRedirectUrl sets ExpiredRedirectUrl field to given value.

### HasExpiredRedirectUrl

`func (o *WorkflowRunLink) HasExpiredRedirectUrl() bool`

HasExpiredRedirectUrl returns a boolean if a field has been set.

### SetExpiredRedirectUrlNil

`func (o *WorkflowRunLink) SetExpiredRedirectUrlNil(b bool)`

 SetExpiredRedirectUrlNil sets the value for ExpiredRedirectUrl to be an explicit nil

### UnsetExpiredRedirectUrl
`func (o *WorkflowRunLink) UnsetExpiredRedirectUrl()`

UnsetExpiredRedirectUrl ensures that no value is present for ExpiredRedirectUrl, not even an explicit nil
### GetExpiresAt

`func (o *WorkflowRunLink) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *WorkflowRunLink) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *WorkflowRunLink) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *WorkflowRunLink) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *WorkflowRunLink) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *WorkflowRunLink) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetLanguage

`func (o *WorkflowRunLink) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *WorkflowRunLink) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *WorkflowRunLink) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *WorkflowRunLink) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *WorkflowRunLink) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *WorkflowRunLink) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


