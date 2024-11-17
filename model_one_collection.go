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

// checks if the OneCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OneCollection{}

// OneCollection struct for OneCollection
type OneCollection struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Slug *string `json:"slug,omitempty"`
	Url *string `json:"url,omitempty"`
	Owner *User `json:"owner,omitempty"`
	Featured *bool `json:"featured,omitempty"`
	Public *bool `json:"public,omitempty"`
	// The value is specified in ISO 8601 (YYYY-MM-DDThh:mm:ss.sZ) format.
	CreatedAt *string `json:"created_at,omitempty"`
	Description *string `json:"description,omitempty"`
	CoverObject *Object `json:"cover_object,omitempty"`
}

// NewOneCollection instantiates a new OneCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOneCollection() *OneCollection {
	this := OneCollection{}
	return &this
}

// NewOneCollectionWithDefaults instantiates a new OneCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOneCollectionWithDefaults() *OneCollection {
	this := OneCollection{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OneCollection) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneCollection) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OneCollection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OneCollection) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OneCollection) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneCollection) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OneCollection) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OneCollection) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *OneCollection) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneCollection) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *OneCollection) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *OneCollection) SetSlug(v string) {
	o.Slug = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *OneCollection) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneCollection) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *OneCollection) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *OneCollection) SetUrl(v string) {
	o.Url = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *OneCollection) GetOwner() User {
	if o == nil || IsNil(o.Owner) {
		var ret User
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneCollection) GetOwnerOk() (*User, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *OneCollection) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given User and assigns it to the Owner field.
func (o *OneCollection) SetOwner(v User) {
	o.Owner = &v
}

// GetFeatured returns the Featured field value if set, zero value otherwise.
func (o *OneCollection) GetFeatured() bool {
	if o == nil || IsNil(o.Featured) {
		var ret bool
		return ret
	}
	return *o.Featured
}

// GetFeaturedOk returns a tuple with the Featured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneCollection) GetFeaturedOk() (*bool, bool) {
	if o == nil || IsNil(o.Featured) {
		return nil, false
	}
	return o.Featured, true
}

// HasFeatured returns a boolean if a field has been set.
func (o *OneCollection) HasFeatured() bool {
	if o != nil && !IsNil(o.Featured) {
		return true
	}

	return false
}

// SetFeatured gets a reference to the given bool and assigns it to the Featured field.
func (o *OneCollection) SetFeatured(v bool) {
	o.Featured = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *OneCollection) GetPublic() bool {
	if o == nil || IsNil(o.Public) {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneCollection) GetPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.Public) {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *OneCollection) HasPublic() bool {
	if o != nil && !IsNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *OneCollection) SetPublic(v bool) {
	o.Public = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OneCollection) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneCollection) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OneCollection) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *OneCollection) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OneCollection) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneCollection) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OneCollection) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OneCollection) SetDescription(v string) {
	o.Description = &v
}

// GetCoverObject returns the CoverObject field value if set, zero value otherwise.
func (o *OneCollection) GetCoverObject() Object {
	if o == nil || IsNil(o.CoverObject) {
		var ret Object
		return ret
	}
	return *o.CoverObject
}

// GetCoverObjectOk returns a tuple with the CoverObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneCollection) GetCoverObjectOk() (*Object, bool) {
	if o == nil || IsNil(o.CoverObject) {
		return nil, false
	}
	return o.CoverObject, true
}

// HasCoverObject returns a boolean if a field has been set.
func (o *OneCollection) HasCoverObject() bool {
	if o != nil && !IsNil(o.CoverObject) {
		return true
	}

	return false
}

// SetCoverObject gets a reference to the given Object and assigns it to the CoverObject field.
func (o *OneCollection) SetCoverObject(v Object) {
	o.CoverObject = &v
}

func (o OneCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OneCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Featured) {
		toSerialize["featured"] = o.Featured
	}
	if !IsNil(o.Public) {
		toSerialize["public"] = o.Public
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CoverObject) {
		toSerialize["cover_object"] = o.CoverObject
	}
	return toSerialize, nil
}

type NullableOneCollection struct {
	value *OneCollection
	isSet bool
}

func (v NullableOneCollection) Get() *OneCollection {
	return v.value
}

func (v *NullableOneCollection) Set(val *OneCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableOneCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableOneCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOneCollection(val *OneCollection) *NullableOneCollection {
	return &NullableOneCollection{value: val, isSet: true}
}

func (v NullableOneCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOneCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

