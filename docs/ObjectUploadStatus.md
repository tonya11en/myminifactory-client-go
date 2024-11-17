# ObjectUploadStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Files** | Pointer to [**[]FileUploadStatus**](FileUploadStatus.md) |  | [optional] 
**Images** | Pointer to [**[]ImageUploadStatus**](ImageUploadStatus.md) |  | [optional] 

## Methods

### NewObjectUploadStatus

`func NewObjectUploadStatus() *ObjectUploadStatus`

NewObjectUploadStatus instantiates a new ObjectUploadStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectUploadStatusWithDefaults

`func NewObjectUploadStatusWithDefaults() *ObjectUploadStatus`

NewObjectUploadStatusWithDefaults instantiates a new ObjectUploadStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObjectUploadStatus) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectUploadStatus) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectUploadStatus) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectUploadStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ObjectUploadStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectUploadStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectUploadStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectUploadStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFiles

`func (o *ObjectUploadStatus) GetFiles() []FileUploadStatus`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ObjectUploadStatus) GetFilesOk() (*[]FileUploadStatus, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ObjectUploadStatus) SetFiles(v []FileUploadStatus)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ObjectUploadStatus) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetImages

`func (o *ObjectUploadStatus) GetImages() []ImageUploadStatus`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ObjectUploadStatus) GetImagesOk() (*[]ImageUploadStatus, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ObjectUploadStatus) SetImages(v []ImageUploadStatus)`

SetImages sets Images field to given value.

### HasImages

`func (o *ObjectUploadStatus) HasImages() bool`

HasImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


