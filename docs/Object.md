# Object

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**ArchiveDownloadUrl** | Pointer to **string** | Available ONLY with Oauth connected User. Not with API key. | [optional] 
**ParentId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** | 0: Private, 2: Public | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Views** | Pointer to **int32** |  | [optional] 
**Likes** | Pointer to **int32** |  | [optional] 
**PublishedAt** | Pointer to **string** | The value is specified in ISO 8601 (YYYY-MM-DDThh:mm:ss.sZ) format. | [optional] 
**Featured** | Pointer to **string** |  | [optional] 
**Complexity** | Pointer to **int32** |  | [optional] 
**Dimensions** | Pointer to **string** |  | [optional] 
**MaterialQuantity** | Pointer to **string** |  | [optional] 
**Designer** | Pointer to [**User**](User.md) |  | [optional] 
**Images** | Pointer to [**[]Image**](Image.md) |  | [optional] 
**Files** | Pointer to [**[]File**](File.md) |  | [optional] 
**Categories** | Pointer to [**[]OneCategory**](OneCategory.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Licenses** | Pointer to [**[]License**](License.md) |  | [optional] 

## Methods

### NewObject

`func NewObject() *Object`

NewObject instantiates a new Object object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectWithDefaults

`func NewObjectWithDefaults() *Object`

NewObjectWithDefaults instantiates a new Object object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Object) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Object) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Object) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Object) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *Object) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Object) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Object) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Object) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetArchiveDownloadUrl

`func (o *Object) GetArchiveDownloadUrl() string`

GetArchiveDownloadUrl returns the ArchiveDownloadUrl field if non-nil, zero value otherwise.

### GetArchiveDownloadUrlOk

`func (o *Object) GetArchiveDownloadUrlOk() (*string, bool)`

GetArchiveDownloadUrlOk returns a tuple with the ArchiveDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveDownloadUrl

`func (o *Object) SetArchiveDownloadUrl(v string)`

SetArchiveDownloadUrl sets ArchiveDownloadUrl field to given value.

### HasArchiveDownloadUrl

`func (o *Object) HasArchiveDownloadUrl() bool`

HasArchiveDownloadUrl returns a boolean if a field has been set.

### GetParentId

`func (o *Object) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Object) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Object) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Object) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetName

`func (o *Object) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Object) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Object) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Object) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVisibility

`func (o *Object) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Object) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Object) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Object) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *Object) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Object) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Object) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Object) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetViews

`func (o *Object) GetViews() int32`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *Object) GetViewsOk() (*int32, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *Object) SetViews(v int32)`

SetViews sets Views field to given value.

### HasViews

`func (o *Object) HasViews() bool`

HasViews returns a boolean if a field has been set.

### GetLikes

`func (o *Object) GetLikes() int32`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *Object) GetLikesOk() (*int32, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *Object) SetLikes(v int32)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *Object) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetPublishedAt

`func (o *Object) GetPublishedAt() string`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *Object) GetPublishedAtOk() (*string, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *Object) SetPublishedAt(v string)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *Object) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetFeatured

`func (o *Object) GetFeatured() string`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *Object) GetFeaturedOk() (*string, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *Object) SetFeatured(v string)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *Object) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetComplexity

`func (o *Object) GetComplexity() int32`

GetComplexity returns the Complexity field if non-nil, zero value otherwise.

### GetComplexityOk

`func (o *Object) GetComplexityOk() (*int32, bool)`

GetComplexityOk returns a tuple with the Complexity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexity

`func (o *Object) SetComplexity(v int32)`

SetComplexity sets Complexity field to given value.

### HasComplexity

`func (o *Object) HasComplexity() bool`

HasComplexity returns a boolean if a field has been set.

### GetDimensions

`func (o *Object) GetDimensions() string`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *Object) GetDimensionsOk() (*string, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *Object) SetDimensions(v string)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *Object) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetMaterialQuantity

`func (o *Object) GetMaterialQuantity() string`

GetMaterialQuantity returns the MaterialQuantity field if non-nil, zero value otherwise.

### GetMaterialQuantityOk

`func (o *Object) GetMaterialQuantityOk() (*string, bool)`

GetMaterialQuantityOk returns a tuple with the MaterialQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterialQuantity

`func (o *Object) SetMaterialQuantity(v string)`

SetMaterialQuantity sets MaterialQuantity field to given value.

### HasMaterialQuantity

`func (o *Object) HasMaterialQuantity() bool`

HasMaterialQuantity returns a boolean if a field has been set.

### GetDesigner

`func (o *Object) GetDesigner() User`

GetDesigner returns the Designer field if non-nil, zero value otherwise.

### GetDesignerOk

`func (o *Object) GetDesignerOk() (*User, bool)`

GetDesignerOk returns a tuple with the Designer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesigner

`func (o *Object) SetDesigner(v User)`

SetDesigner sets Designer field to given value.

### HasDesigner

`func (o *Object) HasDesigner() bool`

HasDesigner returns a boolean if a field has been set.

### GetImages

`func (o *Object) GetImages() []Image`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *Object) GetImagesOk() (*[]Image, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *Object) SetImages(v []Image)`

SetImages sets Images field to given value.

### HasImages

`func (o *Object) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetFiles

`func (o *Object) GetFiles() []File`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *Object) GetFilesOk() (*[]File, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *Object) SetFiles(v []File)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *Object) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetCategories

`func (o *Object) GetCategories() []OneCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Object) GetCategoriesOk() (*[]OneCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Object) SetCategories(v []OneCategory)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *Object) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetTags

`func (o *Object) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Object) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Object) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Object) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLicenses

`func (o *Object) GetLicenses() []License`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *Object) GetLicensesOk() (*[]License, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *Object) SetLicenses(v []License)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *Object) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


