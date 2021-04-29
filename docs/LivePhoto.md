# LivePhoto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier for the photo. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time at which the photo was uploaded. | [optional] 
**Href** | Pointer to **string** | The uri of this resource. | [optional] 
**DownloadHref** | Pointer to **string** | The uri that can be used to download the photo. | [optional] 
**FileName** | Pointer to **string** | The name of the uploaded file. | [optional] 
**FileSize** | Pointer to **int32** | The size of the file in bytes. | [optional] 
**FileType** | Pointer to **string** | The file type of the uploaded file. | [optional] 

## Methods

### NewLivePhoto

`func NewLivePhoto() *LivePhoto`

NewLivePhoto instantiates a new LivePhoto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLivePhotoWithDefaults

`func NewLivePhotoWithDefaults() *LivePhoto`

NewLivePhotoWithDefaults instantiates a new LivePhoto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LivePhoto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LivePhoto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LivePhoto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LivePhoto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LivePhoto) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LivePhoto) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LivePhoto) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LivePhoto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHref

`func (o *LivePhoto) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LivePhoto) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LivePhoto) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *LivePhoto) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetDownloadHref

`func (o *LivePhoto) GetDownloadHref() string`

GetDownloadHref returns the DownloadHref field if non-nil, zero value otherwise.

### GetDownloadHrefOk

`func (o *LivePhoto) GetDownloadHrefOk() (*string, bool)`

GetDownloadHrefOk returns a tuple with the DownloadHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadHref

`func (o *LivePhoto) SetDownloadHref(v string)`

SetDownloadHref sets DownloadHref field to given value.

### HasDownloadHref

`func (o *LivePhoto) HasDownloadHref() bool`

HasDownloadHref returns a boolean if a field has been set.

### GetFileName

`func (o *LivePhoto) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *LivePhoto) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *LivePhoto) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *LivePhoto) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileSize

`func (o *LivePhoto) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *LivePhoto) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *LivePhoto) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *LivePhoto) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetFileType

`func (o *LivePhoto) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *LivePhoto) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *LivePhoto) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *LivePhoto) HasFileType() bool`

HasFileType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


