# Collection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**User**](User.md) |  | [optional] 
**Featured** | Pointer to **bool** |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **string** | The value is specified in ISO 8601 (YYYY-MM-DDThh:mm:ss.sZ) format. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CoverObject** | Pointer to [**Object**](Object.md) |  | [optional] 
**Objects** | Pointer to [**Objects**](Objects.md) |  | [optional] 

## Methods

### NewCollection

`func NewCollection() *Collection`

NewCollection instantiates a new Collection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionWithDefaults

`func NewCollectionWithDefaults() *Collection`

NewCollectionWithDefaults instantiates a new Collection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Collection) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Collection) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Collection) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Collection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Collection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Collection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Collection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Collection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *Collection) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Collection) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Collection) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *Collection) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetUrl

`func (o *Collection) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Collection) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Collection) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Collection) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetOwner

`func (o *Collection) GetOwner() User`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Collection) GetOwnerOk() (*User, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Collection) SetOwner(v User)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Collection) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetFeatured

`func (o *Collection) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *Collection) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *Collection) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *Collection) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetPublic

`func (o *Collection) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *Collection) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *Collection) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *Collection) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Collection) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Collection) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Collection) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Collection) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *Collection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Collection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Collection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Collection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCoverObject

`func (o *Collection) GetCoverObject() Object`

GetCoverObject returns the CoverObject field if non-nil, zero value otherwise.

### GetCoverObjectOk

`func (o *Collection) GetCoverObjectOk() (*Object, bool)`

GetCoverObjectOk returns a tuple with the CoverObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverObject

`func (o *Collection) SetCoverObject(v Object)`

SetCoverObject sets CoverObject field to given value.

### HasCoverObject

`func (o *Collection) HasCoverObject() bool`

HasCoverObject returns a boolean if a field has been set.

### GetObjects

`func (o *Collection) GetObjects() Objects`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *Collection) GetObjectsOk() (*Objects, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *Collection) SetObjects(v Objects)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *Collection) HasObjects() bool`

HasObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


