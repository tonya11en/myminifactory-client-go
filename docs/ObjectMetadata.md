# ObjectMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **int32** | 2: Public, 0: Private | [optional] 
**HowTo** | Pointer to **string** |  | [optional] 
**Dimensions** | Pointer to **string** |  | [optional] 
**TimeToDoFrom** | Pointer to **int32** |  | [optional] 
**TimeToDoTo** | Pointer to **int32** |  | [optional] 
**SupportFree** | Pointer to **bool** |  | [optional] 
**FilamentQuantity** | Pointer to **string** |  | [optional] 
**ClientUrl** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**Licenses** | Pointer to [**[]License**](License.md) |  | [optional] 
**Files** | Pointer to [**[]FileUploadRequest**](FileUploadRequest.md) |  | [optional] 
**Images** | Pointer to [**[]ImageUploadRequest**](ImageUploadRequest.md) |  | [optional] 

## Methods

### NewObjectMetadata

`func NewObjectMetadata() *ObjectMetadata`

NewObjectMetadata instantiates a new ObjectMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectMetadataWithDefaults

`func NewObjectMetadataWithDefaults() *ObjectMetadata`

NewObjectMetadataWithDefaults instantiates a new ObjectMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ObjectMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObjectMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ObjectMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ObjectMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ObjectMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ObjectMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *ObjectMetadata) GetVisibility() int32`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ObjectMetadata) GetVisibilityOk() (*int32, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ObjectMetadata) SetVisibility(v int32)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ObjectMetadata) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetHowTo

`func (o *ObjectMetadata) GetHowTo() string`

GetHowTo returns the HowTo field if non-nil, zero value otherwise.

### GetHowToOk

`func (o *ObjectMetadata) GetHowToOk() (*string, bool)`

GetHowToOk returns a tuple with the HowTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHowTo

`func (o *ObjectMetadata) SetHowTo(v string)`

SetHowTo sets HowTo field to given value.

### HasHowTo

`func (o *ObjectMetadata) HasHowTo() bool`

HasHowTo returns a boolean if a field has been set.

### GetDimensions

`func (o *ObjectMetadata) GetDimensions() string`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *ObjectMetadata) GetDimensionsOk() (*string, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *ObjectMetadata) SetDimensions(v string)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *ObjectMetadata) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetTimeToDoFrom

`func (o *ObjectMetadata) GetTimeToDoFrom() int32`

GetTimeToDoFrom returns the TimeToDoFrom field if non-nil, zero value otherwise.

### GetTimeToDoFromOk

`func (o *ObjectMetadata) GetTimeToDoFromOk() (*int32, bool)`

GetTimeToDoFromOk returns a tuple with the TimeToDoFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToDoFrom

`func (o *ObjectMetadata) SetTimeToDoFrom(v int32)`

SetTimeToDoFrom sets TimeToDoFrom field to given value.

### HasTimeToDoFrom

`func (o *ObjectMetadata) HasTimeToDoFrom() bool`

HasTimeToDoFrom returns a boolean if a field has been set.

### GetTimeToDoTo

`func (o *ObjectMetadata) GetTimeToDoTo() int32`

GetTimeToDoTo returns the TimeToDoTo field if non-nil, zero value otherwise.

### GetTimeToDoToOk

`func (o *ObjectMetadata) GetTimeToDoToOk() (*int32, bool)`

GetTimeToDoToOk returns a tuple with the TimeToDoTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToDoTo

`func (o *ObjectMetadata) SetTimeToDoTo(v int32)`

SetTimeToDoTo sets TimeToDoTo field to given value.

### HasTimeToDoTo

`func (o *ObjectMetadata) HasTimeToDoTo() bool`

HasTimeToDoTo returns a boolean if a field has been set.

### GetSupportFree

`func (o *ObjectMetadata) GetSupportFree() bool`

GetSupportFree returns the SupportFree field if non-nil, zero value otherwise.

### GetSupportFreeOk

`func (o *ObjectMetadata) GetSupportFreeOk() (*bool, bool)`

GetSupportFreeOk returns a tuple with the SupportFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportFree

`func (o *ObjectMetadata) SetSupportFree(v bool)`

SetSupportFree sets SupportFree field to given value.

### HasSupportFree

`func (o *ObjectMetadata) HasSupportFree() bool`

HasSupportFree returns a boolean if a field has been set.

### GetFilamentQuantity

`func (o *ObjectMetadata) GetFilamentQuantity() string`

GetFilamentQuantity returns the FilamentQuantity field if non-nil, zero value otherwise.

### GetFilamentQuantityOk

`func (o *ObjectMetadata) GetFilamentQuantityOk() (*string, bool)`

GetFilamentQuantityOk returns a tuple with the FilamentQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilamentQuantity

`func (o *ObjectMetadata) SetFilamentQuantity(v string)`

SetFilamentQuantity sets FilamentQuantity field to given value.

### HasFilamentQuantity

`func (o *ObjectMetadata) HasFilamentQuantity() bool`

HasFilamentQuantity returns a boolean if a field has been set.

### GetClientUrl

`func (o *ObjectMetadata) GetClientUrl() string`

GetClientUrl returns the ClientUrl field if non-nil, zero value otherwise.

### GetClientUrlOk

`func (o *ObjectMetadata) GetClientUrlOk() (*string, bool)`

GetClientUrlOk returns a tuple with the ClientUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUrl

`func (o *ObjectMetadata) SetClientUrl(v string)`

SetClientUrl sets ClientUrl field to given value.

### HasClientUrl

`func (o *ObjectMetadata) HasClientUrl() bool`

HasClientUrl returns a boolean if a field has been set.

### GetTags

`func (o *ObjectMetadata) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ObjectMetadata) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ObjectMetadata) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ObjectMetadata) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLicenses

`func (o *ObjectMetadata) GetLicenses() []License`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *ObjectMetadata) GetLicensesOk() (*[]License, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *ObjectMetadata) SetLicenses(v []License)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *ObjectMetadata) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetFiles

`func (o *ObjectMetadata) GetFiles() []FileUploadRequest`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ObjectMetadata) GetFilesOk() (*[]FileUploadRequest, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ObjectMetadata) SetFiles(v []FileUploadRequest)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ObjectMetadata) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetImages

`func (o *ObjectMetadata) GetImages() []ImageUploadRequest`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ObjectMetadata) GetImagesOk() (*[]ImageUploadRequest, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ObjectMetadata) SetImages(v []ImageUploadRequest)`

SetImages sets Images field to given value.

### HasImages

`func (o *ObjectMetadata) HasImages() bool`

HasImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


