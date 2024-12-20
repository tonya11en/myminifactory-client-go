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

// checks if the FileUploadId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileUploadId{}

// FileUploadId struct for FileUploadId
type FileUploadId struct {
	UploadId *string `json:"upload_id,omitempty"`
	Filename *string `json:"filename,omitempty"`
}

// NewFileUploadId instantiates a new FileUploadId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileUploadId() *FileUploadId {
	this := FileUploadId{}
	return &this
}

// NewFileUploadIdWithDefaults instantiates a new FileUploadId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileUploadIdWithDefaults() *FileUploadId {
	this := FileUploadId{}
	return &this
}

// GetUploadId returns the UploadId field value if set, zero value otherwise.
func (o *FileUploadId) GetUploadId() string {
	if o == nil || IsNil(o.UploadId) {
		var ret string
		return ret
	}
	return *o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUploadId) GetUploadIdOk() (*string, bool) {
	if o == nil || IsNil(o.UploadId) {
		return nil, false
	}
	return o.UploadId, true
}

// HasUploadId returns a boolean if a field has been set.
func (o *FileUploadId) HasUploadId() bool {
	if o != nil && !IsNil(o.UploadId) {
		return true
	}

	return false
}

// SetUploadId gets a reference to the given string and assigns it to the UploadId field.
func (o *FileUploadId) SetUploadId(v string) {
	o.UploadId = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *FileUploadId) GetFilename() string {
	if o == nil || IsNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileUploadId) GetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *FileUploadId) HasFilename() bool {
	if o != nil && !IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *FileUploadId) SetFilename(v string) {
	o.Filename = &v
}

func (o FileUploadId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileUploadId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UploadId) {
		toSerialize["upload_id"] = o.UploadId
	}
	if !IsNil(o.Filename) {
		toSerialize["filename"] = o.Filename
	}
	return toSerialize, nil
}

type NullableFileUploadId struct {
	value *FileUploadId
	isSet bool
}

func (v NullableFileUploadId) Get() *FileUploadId {
	return v.value
}

func (v *NullableFileUploadId) Set(val *FileUploadId) {
	v.value = val
	v.isSet = true
}

func (v NullableFileUploadId) IsSet() bool {
	return v.isSet
}

func (v *NullableFileUploadId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileUploadId(val *FileUploadId) *NullableFileUploadId {
	return &NullableFileUploadId{value: val, isSet: true}
}

func (v NullableFileUploadId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileUploadId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


