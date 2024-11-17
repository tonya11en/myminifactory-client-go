# CategoriesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** |  | [optional] 
**Items** | Pointer to [**[]Category**](Category.md) |  | [optional] 

## Methods

### NewCategoriesGet200Response

`func NewCategoriesGet200Response() *CategoriesGet200Response`

NewCategoriesGet200Response instantiates a new CategoriesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoriesGet200ResponseWithDefaults

`func NewCategoriesGet200ResponseWithDefaults() *CategoriesGet200Response`

NewCategoriesGet200ResponseWithDefaults instantiates a new CategoriesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *CategoriesGet200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CategoriesGet200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CategoriesGet200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *CategoriesGet200Response) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetItems

`func (o *CategoriesGet200Response) GetItems() []Category`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CategoriesGet200Response) GetItemsOk() (*[]Category, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CategoriesGet200Response) SetItems(v []Category)`

SetItems sets Items field to given value.

### HasItems

`func (o *CategoriesGet200Response) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


