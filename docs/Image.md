# Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**IsPrimary** | Pointer to **bool** |  | [optional] 
**Original** | Pointer to [**PrintOriginal**](PrintOriginal.md) |  | [optional] 
**Thumbnail** | Pointer to [**PrintOriginal**](PrintOriginal.md) |  | [optional] 
**Standard** | Pointer to [**PrintOriginal**](PrintOriginal.md) |  | [optional] 

## Methods

### NewImage

`func NewImage() *Image`

NewImage instantiates a new Image object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageWithDefaults

`func NewImageWithDefaults() *Image`

NewImageWithDefaults instantiates a new Image object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Image) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Image) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Image) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Image) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPrimary

`func (o *Image) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *Image) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *Image) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *Image) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetOriginal

`func (o *Image) GetOriginal() PrintOriginal`

GetOriginal returns the Original field if non-nil, zero value otherwise.

### GetOriginalOk

`func (o *Image) GetOriginalOk() (*PrintOriginal, bool)`

GetOriginalOk returns a tuple with the Original field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginal

`func (o *Image) SetOriginal(v PrintOriginal)`

SetOriginal sets Original field to given value.

### HasOriginal

`func (o *Image) HasOriginal() bool`

HasOriginal returns a boolean if a field has been set.

### GetThumbnail

`func (o *Image) GetThumbnail() PrintOriginal`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *Image) GetThumbnailOk() (*PrintOriginal, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *Image) SetThumbnail(v PrintOriginal)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *Image) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetStandard

`func (o *Image) GetStandard() PrintOriginal`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *Image) GetStandardOk() (*PrintOriginal, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *Image) SetStandard(v PrintOriginal)`

SetStandard sets Standard field to given value.

### HasStandard

`func (o *Image) HasStandard() bool`

HasStandard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


