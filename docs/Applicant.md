# Applicant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier for the applicant. Read-only. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this applicant was created. Read-only. | [optional] [readonly] 
**DeleteAt** | Pointer to **time.Time** | The date and time when this applicant is scheduled to be deleted. Read-only. | [optional] [readonly] 
**Href** | Pointer to **string** | The uri of this resource. Read-only. | [optional] [readonly] 
**Sandbox** | Pointer to **bool** | Read-only. | [optional] [readonly] 
**FirstName** | Pointer to **string** | The applicant’s first name | [optional] 
**LastName** | Pointer to **string** | The applicant’s surname | [optional] 
**Email** | Pointer to **string** | The applicant’s email address. Required if doing a US check, or a UK check for which &#x60;applicant_provides_data&#x60; is &#x60;true&#x60;. | [optional] 
**Dob** | Pointer to **string** | The applicant’s date of birth | [optional] 
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**IdNumbers** | Pointer to [**[]IdNumber**](IdNumber.md) |  | [optional] 

## Methods

### NewApplicant

`func NewApplicant() *Applicant`

NewApplicant instantiates a new Applicant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicantWithDefaults

`func NewApplicantWithDefaults() *Applicant`

NewApplicantWithDefaults instantiates a new Applicant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Applicant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Applicant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Applicant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Applicant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Applicant) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Applicant) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Applicant) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Applicant) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeleteAt

`func (o *Applicant) GetDeleteAt() time.Time`

GetDeleteAt returns the DeleteAt field if non-nil, zero value otherwise.

### GetDeleteAtOk

`func (o *Applicant) GetDeleteAtOk() (*time.Time, bool)`

GetDeleteAtOk returns a tuple with the DeleteAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAt

`func (o *Applicant) SetDeleteAt(v time.Time)`

SetDeleteAt sets DeleteAt field to given value.

### HasDeleteAt

`func (o *Applicant) HasDeleteAt() bool`

HasDeleteAt returns a boolean if a field has been set.

### GetHref

`func (o *Applicant) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Applicant) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Applicant) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Applicant) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSandbox

`func (o *Applicant) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *Applicant) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *Applicant) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.

### HasSandbox

`func (o *Applicant) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.

### GetFirstName

`func (o *Applicant) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Applicant) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Applicant) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Applicant) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *Applicant) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Applicant) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Applicant) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Applicant) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *Applicant) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Applicant) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Applicant) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Applicant) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDob

`func (o *Applicant) GetDob() string`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *Applicant) GetDobOk() (*string, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *Applicant) SetDob(v string)`

SetDob sets Dob field to given value.

### HasDob

`func (o *Applicant) HasDob() bool`

HasDob returns a boolean if a field has been set.

### GetAddress

`func (o *Applicant) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Applicant) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Applicant) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Applicant) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetIdNumbers

`func (o *Applicant) GetIdNumbers() []IdNumber`

GetIdNumbers returns the IdNumbers field if non-nil, zero value otherwise.

### GetIdNumbersOk

`func (o *Applicant) GetIdNumbersOk() (*[]IdNumber, bool)`

GetIdNumbersOk returns a tuple with the IdNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdNumbers

`func (o *Applicant) SetIdNumbers(v []IdNumber)`

SetIdNumbers sets IdNumbers field to given value.

### HasIdNumbers

`func (o *Applicant) HasIdNumbers() bool`

HasIdNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


