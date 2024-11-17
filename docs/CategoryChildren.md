# CategoryChildren

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** |  | [optional] 
**Items** | Pointer to [**[]CategoryChildrenItemsInner**](CategoryChildrenItemsInner.md) |  | [optional] 

## Methods

### NewCategoryChildren

`func NewCategoryChildren() *CategoryChildren`

NewCategoryChildren instantiates a new CategoryChildren object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryChildrenWithDefaults

`func NewCategoryChildrenWithDefaults() *CategoryChildren`

NewCategoryChildrenWithDefaults instantiates a new CategoryChildren object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *CategoryChildren) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CategoryChildren) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CategoryChildren) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *CategoryChildren) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetItems

`func (o *CategoryChildren) GetItems() []CategoryChildrenItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CategoryChildren) GetItemsOk() (*[]CategoryChildrenItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CategoryChildren) SetItems(v []CategoryChildrenItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *CategoryChildren) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


