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

// checks if the User type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &User{}

// User struct for User
type User struct {
	Username *string `json:"username,omitempty"`
	Name *string `json:"name,omitempty"`
	ProfileUrl *string `json:"profile_url,omitempty"`
	AvatarUrl *string `json:"avatar_url,omitempty"`
	AvatarThumbnailUrl *string `json:"avatar_thumbnail_url,omitempty"`
	Bio *string `json:"bio,omitempty"`
	Website *string `json:"website,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *User) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *User) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *User) SetUsername(v string) {
	o.Username = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *User) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *User) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *User) SetName(v string) {
	o.Name = &v
}

// GetProfileUrl returns the ProfileUrl field value if set, zero value otherwise.
func (o *User) GetProfileUrl() string {
	if o == nil || IsNil(o.ProfileUrl) {
		var ret string
		return ret
	}
	return *o.ProfileUrl
}

// GetProfileUrlOk returns a tuple with the ProfileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetProfileUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileUrl) {
		return nil, false
	}
	return o.ProfileUrl, true
}

// HasProfileUrl returns a boolean if a field has been set.
func (o *User) HasProfileUrl() bool {
	if o != nil && !IsNil(o.ProfileUrl) {
		return true
	}

	return false
}

// SetProfileUrl gets a reference to the given string and assigns it to the ProfileUrl field.
func (o *User) SetProfileUrl(v string) {
	o.ProfileUrl = &v
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *User) GetAvatarUrl() string {
	if o == nil || IsNil(o.AvatarUrl) {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAvatarUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AvatarUrl) {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *User) HasAvatarUrl() bool {
	if o != nil && !IsNil(o.AvatarUrl) {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *User) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

// GetAvatarThumbnailUrl returns the AvatarThumbnailUrl field value if set, zero value otherwise.
func (o *User) GetAvatarThumbnailUrl() string {
	if o == nil || IsNil(o.AvatarThumbnailUrl) {
		var ret string
		return ret
	}
	return *o.AvatarThumbnailUrl
}

// GetAvatarThumbnailUrlOk returns a tuple with the AvatarThumbnailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAvatarThumbnailUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AvatarThumbnailUrl) {
		return nil, false
	}
	return o.AvatarThumbnailUrl, true
}

// HasAvatarThumbnailUrl returns a boolean if a field has been set.
func (o *User) HasAvatarThumbnailUrl() bool {
	if o != nil && !IsNil(o.AvatarThumbnailUrl) {
		return true
	}

	return false
}

// SetAvatarThumbnailUrl gets a reference to the given string and assigns it to the AvatarThumbnailUrl field.
func (o *User) SetAvatarThumbnailUrl(v string) {
	o.AvatarThumbnailUrl = &v
}

// GetBio returns the Bio field value if set, zero value otherwise.
func (o *User) GetBio() string {
	if o == nil || IsNil(o.Bio) {
		var ret string
		return ret
	}
	return *o.Bio
}

// GetBioOk returns a tuple with the Bio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetBioOk() (*string, bool) {
	if o == nil || IsNil(o.Bio) {
		return nil, false
	}
	return o.Bio, true
}

// HasBio returns a boolean if a field has been set.
func (o *User) HasBio() bool {
	if o != nil && !IsNil(o.Bio) {
		return true
	}

	return false
}

// SetBio gets a reference to the given string and assigns it to the Bio field.
func (o *User) SetBio(v string) {
	o.Bio = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *User) GetWebsite() string {
	if o == nil || IsNil(o.Website) {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *User) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *User) SetWebsite(v string) {
	o.Website = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o User) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ProfileUrl) {
		toSerialize["profile_url"] = o.ProfileUrl
	}
	if !IsNil(o.AvatarUrl) {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	if !IsNil(o.AvatarThumbnailUrl) {
		toSerialize["avatar_thumbnail_url"] = o.AvatarThumbnailUrl
	}
	if !IsNil(o.Bio) {
		toSerialize["bio"] = o.Bio
	}
	if !IsNil(o.Website) {
		toSerialize["website"] = o.Website
	}
	return toSerialize, nil
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


