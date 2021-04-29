# SdkToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicantId** | Pointer to **string** | The unique identifier of the applicant | [optional] 
**Referrer** | Pointer to **string** | The referrer URL pattern | [optional] 
**ApplicationId** | Pointer to **string** | The application ID (iOS or Android) | [optional] 
**CrossDeviceUrl** | Pointer to **string** | Enterprise Feature - The URL to be used for the cross device flow of the Web SDK | [optional] 
**Token** | Pointer to **string** | The generated SDK token | [optional] [readonly] 

## Methods

### NewSdkToken

`func NewSdkToken() *SdkToken`

NewSdkToken instantiates a new SdkToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkTokenWithDefaults

`func NewSdkTokenWithDefaults() *SdkToken`

NewSdkTokenWithDefaults instantiates a new SdkToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicantId

`func (o *SdkToken) GetApplicantId() string`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *SdkToken) GetApplicantIdOk() (*string, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *SdkToken) SetApplicantId(v string)`

SetApplicantId sets ApplicantId field to given value.

### HasApplicantId

`func (o *SdkToken) HasApplicantId() bool`

HasApplicantId returns a boolean if a field has been set.

### GetReferrer

`func (o *SdkToken) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *SdkToken) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *SdkToken) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.

### HasReferrer

`func (o *SdkToken) HasReferrer() bool`

HasReferrer returns a boolean if a field has been set.

### GetApplicationId

`func (o *SdkToken) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *SdkToken) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *SdkToken) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *SdkToken) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetCrossDeviceUrl

`func (o *SdkToken) GetCrossDeviceUrl() string`

GetCrossDeviceUrl returns the CrossDeviceUrl field if non-nil, zero value otherwise.

### GetCrossDeviceUrlOk

`func (o *SdkToken) GetCrossDeviceUrlOk() (*string, bool)`

GetCrossDeviceUrlOk returns a tuple with the CrossDeviceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossDeviceUrl

`func (o *SdkToken) SetCrossDeviceUrl(v string)`

SetCrossDeviceUrl sets CrossDeviceUrl field to given value.

### HasCrossDeviceUrl

`func (o *SdkToken) HasCrossDeviceUrl() bool`

HasCrossDeviceUrl returns a boolean if a field has been set.

### GetToken

`func (o *SdkToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SdkToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SdkToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SdkToken) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


