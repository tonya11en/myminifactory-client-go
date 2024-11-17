# Print

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Original** | Pointer to [**PrintOriginal**](PrintOriginal.md) |  | [optional] 
**Thumbnail** | Pointer to [**PrintOriginal**](PrintOriginal.md) |  | [optional] 
**Standard** | Pointer to [**PrintOriginal**](PrintOriginal.md) |  | [optional] 

## Methods

### NewPrint

`func NewPrint() *Print`

NewPrint instantiates a new Print object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrintWithDefaults

`func NewPrintWithDefaults() *Print`

NewPrintWithDefaults instantiates a new Print object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Print) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Print) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Print) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Print) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOriginal

`func (o *Print) GetOriginal() PrintOriginal`

GetOriginal returns the Original field if non-nil, zero value otherwise.

### GetOriginalOk

`func (o *Print) GetOriginalOk() (*PrintOriginal, bool)`

GetOriginalOk returns a tuple with the Original field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginal

`func (o *Print) SetOriginal(v PrintOriginal)`

SetOriginal sets Original field to given value.

### HasOriginal

`func (o *Print) HasOriginal() bool`

HasOriginal returns a boolean if a field has been set.

### GetThumbnail

`func (o *Print) GetThumbnail() PrintOriginal`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *Print) GetThumbnailOk() (*PrintOriginal, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *Print) SetThumbnail(v PrintOriginal)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *Print) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetStandard

`func (o *Print) GetStandard() PrintOriginal`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *Print) GetStandardOk() (*PrintOriginal, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *Print) SetStandard(v PrintOriginal)`

SetStandard sets Standard field to given value.

### HasStandard

`func (o *Print) HasStandard() bool`

HasStandard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


