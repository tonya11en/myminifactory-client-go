# ObjectUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ObjectStatusUrl** | Pointer to **string** |  | [optional] 
**ObjectUrl** | Pointer to **string** |  | [optional] 
**Files** | Pointer to [**[]FileUploadId**](FileUploadId.md) |  | [optional] 
**Images** | Pointer to [**[]ImageUploadId**](ImageUploadId.md) |  | [optional] 

## Methods

### NewObjectUpload

`func NewObjectUpload() *ObjectUpload`

NewObjectUpload instantiates a new ObjectUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectUploadWithDefaults

`func NewObjectUploadWithDefaults() *ObjectUpload`

NewObjectUploadWithDefaults instantiates a new ObjectUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObjectUpload) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectUpload) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectUpload) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectUpload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ObjectUpload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectUpload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectUpload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectUpload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectStatusUrl

`func (o *ObjectUpload) GetObjectStatusUrl() string`

GetObjectStatusUrl returns the ObjectStatusUrl field if non-nil, zero value otherwise.

### GetObjectStatusUrlOk

`func (o *ObjectUpload) GetObjectStatusUrlOk() (*string, bool)`

GetObjectStatusUrlOk returns a tuple with the ObjectStatusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStatusUrl

`func (o *ObjectUpload) SetObjectStatusUrl(v string)`

SetObjectStatusUrl sets ObjectStatusUrl field to given value.

### HasObjectStatusUrl

`func (o *ObjectUpload) HasObjectStatusUrl() bool`

HasObjectStatusUrl returns a boolean if a field has been set.

### GetObjectUrl

`func (o *ObjectUpload) GetObjectUrl() string`

GetObjectUrl returns the ObjectUrl field if non-nil, zero value otherwise.

### GetObjectUrlOk

`func (o *ObjectUpload) GetObjectUrlOk() (*string, bool)`

GetObjectUrlOk returns a tuple with the ObjectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectUrl

`func (o *ObjectUpload) SetObjectUrl(v string)`

SetObjectUrl sets ObjectUrl field to given value.

### HasObjectUrl

`func (o *ObjectUpload) HasObjectUrl() bool`

HasObjectUrl returns a boolean if a field has been set.

### GetFiles

`func (o *ObjectUpload) GetFiles() []FileUploadId`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ObjectUpload) GetFilesOk() (*[]FileUploadId, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ObjectUpload) SetFiles(v []FileUploadId)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ObjectUpload) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetImages

`func (o *ObjectUpload) GetImages() []ImageUploadId`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ObjectUpload) GetImagesOk() (*[]ImageUploadId, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ObjectUpload) SetImages(v []ImageUploadId)`

SetImages sets Images field to given value.

### HasImages

`func (o *ObjectUpload) HasImages() bool`

HasImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


