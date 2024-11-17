# SearchGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** |  | [optional] 
**Items** | Pointer to [**[]Object**](Object.md) |  | [optional] 

## Methods

### NewSearchGet200Response

`func NewSearchGet200Response() *SearchGet200Response`

NewSearchGet200Response instantiates a new SearchGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchGet200ResponseWithDefaults

`func NewSearchGet200ResponseWithDefaults() *SearchGet200Response`

NewSearchGet200ResponseWithDefaults instantiates a new SearchGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *SearchGet200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SearchGet200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SearchGet200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *SearchGet200Response) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetItems

`func (o *SearchGet200Response) GetItems() []Object`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SearchGet200Response) GetItemsOk() (*[]Object, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SearchGet200Response) SetItems(v []Object)`

SetItems sets Items field to given value.

### HasItems

`func (o *SearchGet200Response) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


