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

// checks if the UsersUsernameCollectionsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsersUsernameCollectionsGet200Response{}

// UsersUsernameCollectionsGet200Response struct for UsersUsernameCollectionsGet200Response
type UsersUsernameCollectionsGet200Response struct {
	TotalCount *int32 `json:"total_count,omitempty"`
	Items []OneCollection `json:"items,omitempty"`
}

// NewUsersUsernameCollectionsGet200Response instantiates a new UsersUsernameCollectionsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersUsernameCollectionsGet200Response() *UsersUsernameCollectionsGet200Response {
	this := UsersUsernameCollectionsGet200Response{}
	return &this
}

// NewUsersUsernameCollectionsGet200ResponseWithDefaults instantiates a new UsersUsernameCollectionsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersUsernameCollectionsGet200ResponseWithDefaults() *UsersUsernameCollectionsGet200Response {
	this := UsersUsernameCollectionsGet200Response{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *UsersUsernameCollectionsGet200Response) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersUsernameCollectionsGet200Response) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *UsersUsernameCollectionsGet200Response) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *UsersUsernameCollectionsGet200Response) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *UsersUsernameCollectionsGet200Response) GetItems() []OneCollection {
	if o == nil || IsNil(o.Items) {
		var ret []OneCollection
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersUsernameCollectionsGet200Response) GetItemsOk() ([]OneCollection, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *UsersUsernameCollectionsGet200Response) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []OneCollection and assigns it to the Items field.
func (o *UsersUsernameCollectionsGet200Response) SetItems(v []OneCollection) {
	o.Items = v
}

func (o UsersUsernameCollectionsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsersUsernameCollectionsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalCount) {
		toSerialize["total_count"] = o.TotalCount
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableUsersUsernameCollectionsGet200Response struct {
	value *UsersUsernameCollectionsGet200Response
	isSet bool
}

func (v NullableUsersUsernameCollectionsGet200Response) Get() *UsersUsernameCollectionsGet200Response {
	return v.value
}

func (v *NullableUsersUsernameCollectionsGet200Response) Set(val *UsersUsernameCollectionsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersUsernameCollectionsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersUsernameCollectionsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersUsernameCollectionsGet200Response(val *UsersUsernameCollectionsGet200Response) *NullableUsersUsernameCollectionsGet200Response {
	return &NullableUsersUsernameCollectionsGet200Response{value: val, isSet: true}
}

func (v NullableUsersUsernameCollectionsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersUsernameCollectionsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

