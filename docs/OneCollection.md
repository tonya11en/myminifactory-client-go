# OneCollection

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

## Methods

### NewOneCollection

`func NewOneCollection() *OneCollection`

NewOneCollection instantiates a new OneCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOneCollectionWithDefaults

`func NewOneCollectionWithDefaults() *OneCollection`

NewOneCollectionWithDefaults instantiates a new OneCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OneCollection) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OneCollection) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OneCollection) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OneCollection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OneCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OneCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OneCollection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OneCollection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *OneCollection) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *OneCollection) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *OneCollection) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *OneCollection) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetUrl

`func (o *OneCollection) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OneCollection) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OneCollection) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OneCollection) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetOwner

`func (o *OneCollection) GetOwner() User`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *OneCollection) GetOwnerOk() (*User, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *OneCollection) SetOwner(v User)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *OneCollection) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetFeatured

`func (o *OneCollection) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *OneCollection) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *OneCollection) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *OneCollection) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetPublic

`func (o *OneCollection) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *OneCollection) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *OneCollection) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *OneCollection) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OneCollection) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OneCollection) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OneCollection) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OneCollection) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *OneCollection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OneCollection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OneCollection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OneCollection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCoverObject

`func (o *OneCollection) GetCoverObject() Object`

GetCoverObject returns the CoverObject field if non-nil, zero value otherwise.

### GetCoverObjectOk

`func (o *OneCollection) GetCoverObjectOk() (*Object, bool)`

GetCoverObjectOk returns a tuple with the CoverObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverObject

`func (o *OneCollection) SetCoverObject(v Object)`

SetCoverObject sets CoverObject field to given value.

### HasCoverObject

`func (o *OneCollection) HasCoverObject() bool`

HasCoverObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


