# ImageUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 

## Methods

### NewImageUploadRequest

`func NewImageUploadRequest() *ImageUploadRequest`

NewImageUploadRequest instantiates a new ImageUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageUploadRequestWithDefaults

`func NewImageUploadRequestWithDefaults() *ImageUploadRequest`

NewImageUploadRequestWithDefaults instantiates a new ImageUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *ImageUploadRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ImageUploadRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ImageUploadRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *ImageUploadRequest) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetSize

`func (o *ImageUploadRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ImageUploadRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ImageUploadRequest) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ImageUploadRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


