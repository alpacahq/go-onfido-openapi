# LiveVideo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier for the video. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time at which the video was uploaded. | [optional] 
**Href** | Pointer to **string** | The uri of this resource. | [optional] 
**DownloadHref** | Pointer to **string** | The uri that can be used to download the video. | [optional] 
**FileName** | Pointer to **string** | The name of the uploaded file. | [optional] 
**FileSize** | Pointer to **int32** | The size of the file in bytes. | [optional] 
**FileType** | Pointer to **string** | The file type of the uploaded file. | [optional] 

## Methods

### NewLiveVideo

`func NewLiveVideo() *LiveVideo`

NewLiveVideo instantiates a new LiveVideo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveVideoWithDefaults

`func NewLiveVideoWithDefaults() *LiveVideo`

NewLiveVideoWithDefaults instantiates a new LiveVideo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LiveVideo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LiveVideo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LiveVideo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LiveVideo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LiveVideo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LiveVideo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LiveVideo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LiveVideo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHref

`func (o *LiveVideo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LiveVideo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LiveVideo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *LiveVideo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetDownloadHref

`func (o *LiveVideo) GetDownloadHref() string`

GetDownloadHref returns the DownloadHref field if non-nil, zero value otherwise.

### GetDownloadHrefOk

`func (o *LiveVideo) GetDownloadHrefOk() (*string, bool)`

GetDownloadHrefOk returns a tuple with the DownloadHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadHref

`func (o *LiveVideo) SetDownloadHref(v string)`

SetDownloadHref sets DownloadHref field to given value.

### HasDownloadHref

`func (o *LiveVideo) HasDownloadHref() bool`

HasDownloadHref returns a boolean if a field has been set.

### GetFileName

`func (o *LiveVideo) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *LiveVideo) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *LiveVideo) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *LiveVideo) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileSize

`func (o *LiveVideo) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *LiveVideo) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *LiveVideo) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *LiveVideo) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetFileType

`func (o *LiveVideo) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *LiveVideo) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *LiveVideo) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *LiveVideo) HasFileType() bool`

HasFileType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


