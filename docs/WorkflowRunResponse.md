# WorkflowRunResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the Workflow Run. | 
**WorkflowVersionId** | Pointer to **int32** | The identifier for the Workflow version. | [optional] 
**DashboardUrl** | Pointer to **string** | The URL for viewing the Workflow Run results on your Onfido Dashboard. | [optional] 
**Status** | Pointer to [**WorkflowRunStatus**](workflow_run_status.md) | The status of the Workflow Run. | [optional] 
**Output** | Pointer to **map[string]interface{}** | Output object contains all of the properties configured on the Workflow version. | [optional] 
**Reasons** | Pointer to **[]string** | The reasons the Workflow Run outcome was reached. Configurable when creating the Workflow version. | [optional] 
**Error** | Pointer to [**WorkflowRunResponseError**](WorkflowRunResponseError.md) |  | [optional] 
**SdkToken** | Pointer to **NullableString** | Client token to use when loading this workflow run in the Onfido SDK. | [optional] 

## Methods

### NewWorkflowRunResponse

`func NewWorkflowRunResponse(id string, ) *WorkflowRunResponse`

NewWorkflowRunResponse instantiates a new WorkflowRunResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRunResponseWithDefaults

`func NewWorkflowRunResponseWithDefaults() *WorkflowRunResponse`

NewWorkflowRunResponseWithDefaults instantiates a new WorkflowRunResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowRunResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowRunResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowRunResponse) SetId(v string)`

SetId sets Id field to given value.


### GetWorkflowVersionId

`func (o *WorkflowRunResponse) GetWorkflowVersionId() int32`

GetWorkflowVersionId returns the WorkflowVersionId field if non-nil, zero value otherwise.

### GetWorkflowVersionIdOk

`func (o *WorkflowRunResponse) GetWorkflowVersionIdOk() (*int32, bool)`

GetWorkflowVersionIdOk returns a tuple with the WorkflowVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowVersionId

`func (o *WorkflowRunResponse) SetWorkflowVersionId(v int32)`

SetWorkflowVersionId sets WorkflowVersionId field to given value.

### HasWorkflowVersionId

`func (o *WorkflowRunResponse) HasWorkflowVersionId() bool`

HasWorkflowVersionId returns a boolean if a field has been set.

### GetDashboardUrl

`func (o *WorkflowRunResponse) GetDashboardUrl() string`

GetDashboardUrl returns the DashboardUrl field if non-nil, zero value otherwise.

### GetDashboardUrlOk

`func (o *WorkflowRunResponse) GetDashboardUrlOk() (*string, bool)`

GetDashboardUrlOk returns a tuple with the DashboardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardUrl

`func (o *WorkflowRunResponse) SetDashboardUrl(v string)`

SetDashboardUrl sets DashboardUrl field to given value.

### HasDashboardUrl

`func (o *WorkflowRunResponse) HasDashboardUrl() bool`

HasDashboardUrl returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowRunResponse) GetStatus() WorkflowRunStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowRunResponse) GetStatusOk() (*WorkflowRunStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowRunResponse) SetStatus(v WorkflowRunStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowRunResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOutput

`func (o *WorkflowRunResponse) GetOutput() map[string]interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *WorkflowRunResponse) GetOutputOk() (*map[string]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *WorkflowRunResponse) SetOutput(v map[string]interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *WorkflowRunResponse) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetReasons

`func (o *WorkflowRunResponse) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *WorkflowRunResponse) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *WorkflowRunResponse) SetReasons(v []string)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *WorkflowRunResponse) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetError

`func (o *WorkflowRunResponse) GetError() WorkflowRunResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WorkflowRunResponse) GetErrorOk() (*WorkflowRunResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WorkflowRunResponse) SetError(v WorkflowRunResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *WorkflowRunResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetSdkToken

`func (o *WorkflowRunResponse) GetSdkToken() string`

GetSdkToken returns the SdkToken field if non-nil, zero value otherwise.

### GetSdkTokenOk

`func (o *WorkflowRunResponse) GetSdkTokenOk() (*string, bool)`

GetSdkTokenOk returns a tuple with the SdkToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkToken

`func (o *WorkflowRunResponse) SetSdkToken(v string)`

SetSdkToken sets SdkToken field to given value.

### HasSdkToken

`func (o *WorkflowRunResponse) HasSdkToken() bool`

HasSdkToken returns a boolean if a field has been set.

### SetSdkTokenNil

`func (o *WorkflowRunResponse) SetSdkTokenNil(b bool)`

 SetSdkTokenNil sets the value for SdkToken to be an explicit nil

### UnsetSdkToken
`func (o *WorkflowRunResponse) UnsetSdkToken()`

UnsetSdkToken ensures that no value is present for SdkToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


