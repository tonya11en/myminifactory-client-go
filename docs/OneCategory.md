# OneCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Children** | Pointer to [**CategoryChildren**](CategoryChildren.md) |  | [optional] 
**Parent** | Pointer to [**CategoryChildrenItemsInner**](CategoryChildrenItemsInner.md) |  | [optional] 

## Methods

### NewOneCategory

`func NewOneCategory() *OneCategory`

NewOneCategory instantiates a new OneCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOneCategoryWithDefaults

`func NewOneCategoryWithDefaults() *OneCategory`

NewOneCategoryWithDefaults instantiates a new OneCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OneCategory) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OneCategory) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OneCategory) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OneCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSlug

`func (o *OneCategory) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *OneCategory) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *OneCategory) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *OneCategory) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetUrl

`func (o *OneCategory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OneCategory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OneCategory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OneCategory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *OneCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OneCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OneCategory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OneCategory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetChildren

`func (o *OneCategory) GetChildren() CategoryChildren`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *OneCategory) GetChildrenOk() (*CategoryChildren, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *OneCategory) SetChildren(v CategoryChildren)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *OneCategory) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetParent

`func (o *OneCategory) GetParent() CategoryChildrenItemsInner`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *OneCategory) GetParentOk() (*CategoryChildrenItemsInner, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *OneCategory) SetParent(v CategoryChildrenItemsInner)`

SetParent sets Parent field to given value.

### HasParent

`func (o *OneCategory) HasParent() bool`

HasParent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


