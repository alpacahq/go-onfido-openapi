# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlatNumber** | Pointer to **string** | The flat number of this address | [optional] 
**BuildingNumber** | Pointer to **string** | The building number of this address | [optional] 
**BuildingName** | Pointer to **string** | The building name of this address | [optional] 
**Street** | Pointer to **string** | The street of the applicant’s address | [optional] 
**SubStreet** | Pointer to **string** | The sub-street of the applicant’s address | [optional] 
**Town** | Pointer to **string** | The town of the applicant’s address | [optional] 
**Postcode** | Pointer to **string** | The postcode or ZIP of the applicant’s address | [optional] 
**Country** | Pointer to **string** | The 3 character ISO country code of this address. For example, GBR is the country code for the United Kingdom | [optional] 
**State** | Pointer to **string** | The address state. US states must use the USPS abbreviation (see also ISO 3166-2:US), for example AK, CA, or TX. | [optional] 
**Line1** | Pointer to **string** | Line 1 of the applicant&#39;s address | [optional] 
**Line2** | Pointer to **string** | Line 2 of the applicant&#39;s address | [optional] 
**Line3** | Pointer to **string** | Line 3 of the applicant&#39;s address | [optional] 

## Methods

### NewAddress

`func NewAddress() *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlatNumber

`func (o *Address) GetFlatNumber() string`

GetFlatNumber returns the FlatNumber field if non-nil, zero value otherwise.

### GetFlatNumberOk

`func (o *Address) GetFlatNumberOk() (*string, bool)`

GetFlatNumberOk returns a tuple with the FlatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatNumber

`func (o *Address) SetFlatNumber(v string)`

SetFlatNumber sets FlatNumber field to given value.

### HasFlatNumber

`func (o *Address) HasFlatNumber() bool`

HasFlatNumber returns a boolean if a field has been set.

### GetBuildingNumber

`func (o *Address) GetBuildingNumber() string`

GetBuildingNumber returns the BuildingNumber field if non-nil, zero value otherwise.

### GetBuildingNumberOk

`func (o *Address) GetBuildingNumberOk() (*string, bool)`

GetBuildingNumberOk returns a tuple with the BuildingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingNumber

`func (o *Address) SetBuildingNumber(v string)`

SetBuildingNumber sets BuildingNumber field to given value.

### HasBuildingNumber

`func (o *Address) HasBuildingNumber() bool`

HasBuildingNumber returns a boolean if a field has been set.

### GetBuildingName

`func (o *Address) GetBuildingName() string`

GetBuildingName returns the BuildingName field if non-nil, zero value otherwise.

### GetBuildingNameOk

`func (o *Address) GetBuildingNameOk() (*string, bool)`

GetBuildingNameOk returns a tuple with the BuildingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingName

`func (o *Address) SetBuildingName(v string)`

SetBuildingName sets BuildingName field to given value.

### HasBuildingName

`func (o *Address) HasBuildingName() bool`

HasBuildingName returns a boolean if a field has been set.

### GetStreet

`func (o *Address) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *Address) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *Address) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *Address) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetSubStreet

`func (o *Address) GetSubStreet() string`

GetSubStreet returns the SubStreet field if non-nil, zero value otherwise.

### GetSubStreetOk

`func (o *Address) GetSubStreetOk() (*string, bool)`

GetSubStreetOk returns a tuple with the SubStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStreet

`func (o *Address) SetSubStreet(v string)`

SetSubStreet sets SubStreet field to given value.

### HasSubStreet

`func (o *Address) HasSubStreet() bool`

HasSubStreet returns a boolean if a field has been set.

### GetTown

`func (o *Address) GetTown() string`

GetTown returns the Town field if non-nil, zero value otherwise.

### GetTownOk

`func (o *Address) GetTownOk() (*string, bool)`

GetTownOk returns a tuple with the Town field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTown

`func (o *Address) SetTown(v string)`

SetTown sets Town field to given value.

### HasTown

`func (o *Address) HasTown() bool`

HasTown returns a boolean if a field has been set.

### GetPostcode

`func (o *Address) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *Address) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *Address) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.

### HasPostcode

`func (o *Address) HasPostcode() bool`

HasPostcode returns a boolean if a field has been set.

### GetCountry

`func (o *Address) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Address) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Address) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Address) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetState

`func (o *Address) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Address) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Address) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Address) HasState() bool`

HasState returns a boolean if a field has been set.

### GetLine1

`func (o *Address) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *Address) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *Address) SetLine1(v string)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *Address) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### GetLine2

`func (o *Address) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *Address) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *Address) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *Address) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### GetLine3

`func (o *Address) GetLine3() string`

GetLine3 returns the Line3 field if non-nil, zero value otherwise.

### GetLine3Ok

`func (o *Address) GetLine3Ok() (*string, bool)`

GetLine3Ok returns a tuple with the Line3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine3

`func (o *Address) SetLine3(v string)`

SetLine3 sets Line3 field to given value.

### HasLine3

`func (o *Address) HasLine3() bool`

HasLine3 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


