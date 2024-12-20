/*
MyMiniFactory API

3D printable object API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Image type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Image{}

// Image struct for Image
type Image struct {
	Id *int32 `json:"id,omitempty"`
	IsPrimary *bool `json:"is_primary,omitempty"`
	Original *PrintOriginal `json:"original,omitempty"`
	Thumbnail *PrintOriginal `json:"thumbnail,omitempty"`
	Standard *PrintOriginal `json:"standard,omitempty"`
}

// NewImage instantiates a new Image object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImage() *Image {
	this := Image{}
	return &this
}

// NewImageWithDefaults instantiates a new Image object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageWithDefaults() *Image {
	this := Image{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Image) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Image) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Image) SetId(v int32) {
	o.Id = &v
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *Image) GetIsPrimary() bool {
	if o == nil || IsNil(o.IsPrimary) {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrimary) {
		return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *Image) HasIsPrimary() bool {
	if o != nil && !IsNil(o.IsPrimary) {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *Image) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetOriginal returns the Original field value if set, zero value otherwise.
func (o *Image) GetOriginal() PrintOriginal {
	if o == nil || IsNil(o.Original) {
		var ret PrintOriginal
		return ret
	}
	return *o.Original
}

// GetOriginalOk returns a tuple with the Original field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetOriginalOk() (*PrintOriginal, bool) {
	if o == nil || IsNil(o.Original) {
		return nil, false
	}
	return o.Original, true
}

// HasOriginal returns a boolean if a field has been set.
func (o *Image) HasOriginal() bool {
	if o != nil && !IsNil(o.Original) {
		return true
	}

	return false
}

// SetOriginal gets a reference to the given PrintOriginal and assigns it to the Original field.
func (o *Image) SetOriginal(v PrintOriginal) {
	o.Original = &v
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *Image) GetThumbnail() PrintOriginal {
	if o == nil || IsNil(o.Thumbnail) {
		var ret PrintOriginal
		return ret
	}
	return *o.Thumbnail
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetThumbnailOk() (*PrintOriginal, bool) {
	if o == nil || IsNil(o.Thumbnail) {
		return nil, false
	}
	return o.Thumbnail, true
}

// HasThumbnail returns a boolean if a field has been set.
func (o *Image) HasThumbnail() bool {
	if o != nil && !IsNil(o.Thumbnail) {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given PrintOriginal and assigns it to the Thumbnail field.
func (o *Image) SetThumbnail(v PrintOriginal) {
	o.Thumbnail = &v
}

// GetStandard returns the Standard field value if set, zero value otherwise.
func (o *Image) GetStandard() PrintOriginal {
	if o == nil || IsNil(o.Standard) {
		var ret PrintOriginal
		return ret
	}
	return *o.Standard
}

// GetStandardOk returns a tuple with the Standard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetStandardOk() (*PrintOriginal, bool) {
	if o == nil || IsNil(o.Standard) {
		return nil, false
	}
	return o.Standard, true
}

// HasStandard returns a boolean if a field has been set.
func (o *Image) HasStandard() bool {
	if o != nil && !IsNil(o.Standard) {
		return true
	}

	return false
}

// SetStandard gets a reference to the given PrintOriginal and assigns it to the Standard field.
func (o *Image) SetStandard(v PrintOriginal) {
	o.Standard = &v
}

func (o Image) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Image) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsPrimary) {
		toSerialize["is_primary"] = o.IsPrimary
	}
	if !IsNil(o.Original) {
		toSerialize["original"] = o.Original
	}
	if !IsNil(o.Thumbnail) {
		toSerialize["thumbnail"] = o.Thumbnail
	}
	if !IsNil(o.Standard) {
		toSerialize["standard"] = o.Standard
	}
	return toSerialize, nil
}

type NullableImage struct {
	value *Image
	isSet bool
}

func (v NullableImage) Get() *Image {
	return v.value
}

func (v *NullableImage) Set(val *Image) {
	v.value = val
	v.isSet = true
}

func (v NullableImage) IsSet() bool {
	return v.isSet
}

func (v *NullableImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImage(val *Image) *NullableImage {
	return &NullableImage{value: val, isSet: true}
}

func (v NullableImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


