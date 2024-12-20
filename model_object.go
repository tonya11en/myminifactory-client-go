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

// checks if the Object type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Object{}

// Object struct for Object
type Object struct {
	Id *int32 `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	// Available ONLY with Oauth connected User. Not with API key.
	ArchiveDownloadUrl *string `json:"archive_download_url,omitempty"`
	ParentId *int32 `json:"parent_id,omitempty"`
	Name *string `json:"name,omitempty"`
	// 0: Private, 2: Public
	Visibility *string `json:"visibility,omitempty"`
	Description *string `json:"description,omitempty"`
	Views *int32 `json:"views,omitempty"`
	Likes *int32 `json:"likes,omitempty"`
	// The value is specified in ISO 8601 (YYYY-MM-DDThh:mm:ss.sZ) format.
	PublishedAt *string `json:"published_at,omitempty"`
	Featured *string `json:"featured,omitempty"`
	Complexity *int32 `json:"complexity,omitempty"`
	Dimensions *string `json:"dimensions,omitempty"`
	MaterialQuantity *string `json:"material_quantity,omitempty"`
	Designer *User `json:"designer,omitempty"`
	Images []Image `json:"images,omitempty"`
	Files []File `json:"files,omitempty"`
	Categories []OneCategory `json:"categories,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Licenses []License `json:"licenses,omitempty"`
}

// NewObject instantiates a new Object object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObject() *Object {
	this := Object{}
	return &this
}

// NewObjectWithDefaults instantiates a new Object object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectWithDefaults() *Object {
	this := Object{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Object) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Object) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Object) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Object) SetId(v int32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Object) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Object) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Object) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Object) SetUrl(v string) {
	o.Url = &v
}

// GetArchiveDownloadUrl returns the ArchiveDownloadUrl field value if set, zero value otherwise.
func (o *Object) GetArchiveDownloadUrl() string {
	if o == nil || IsNil(o.ArchiveDownloadUrl) {
		var ret string
		return ret
	}
	return *o.ArchiveDownloadUrl
}

// GetArchiveDownloadUrlOk returns a tuple with the ArchiveDownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Object) GetArchiveDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ArchiveDownloadUrl) {
		return nil, false
	}
	return o.ArchiveDownloadUrl, true
}

// HasArchiveDownloadUrl returns a boolean if a field has been set.
func (o *Object) HasArchiveDownloadUrl() bool {
	if o != nil && !IsNil(o.ArchiveDownloadUrl) {
		return true
	}

	return false
}

// SetArchiveDownloadUrl gets a reference to the given string and assigns it to the ArchiveDownloadUrl field.
func (o *Object) SetArchiveDownloadUrl(v string) {
	o.ArchiveDownloadUrl = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *Object) GetParentId() int32 {
	if o == nil || IsNil(o.ParentId) {
		var ret int32
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Object) GetParentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *Object) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given int32 and assigns it to the ParentId field.
func (o *Object) SetParentId(v int32) {
	o.ParentId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Object) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Object) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Object) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Object) SetName(v string) {
	o.Name = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *Object) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Object) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *Object) HasVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *Object) SetVisibility(v string) {
	o.Visibility = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Object) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Object) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Object) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Object) SetDescription(v string) {
	o.Description = &v
}

// GetViews returns the Views field value if set, zero value otherwise.
func (o *Object) GetViews() int32 {
	if o == nil || IsNil(o.Views) {
		var ret int32
		return ret
	}
	return *o.Views
}

// GetViewsOk returns a tuple with the Views field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Object) GetViewsOk() (*int32, bool) {
	if o == nil || IsNil(o.Views) {
		return nil, false
	}
	return o.Views, true
}

// HasViews returns a boolean if a field has been set.
func (o *Object) HasViews() bool {
	if o != nil && !IsNil(o.Views) {
		return true
	}

	return false
}

// SetViews gets a reference to the given int32 and assigns it to the Views field.
func (o *Object) SetViews(v int32) {
	o.Views = &v
}

// GetLikes returns the Likes field value if set, zero value otherwise.
func (o *Object) GetLikes() int32 {
	if o == nil || IsNil(o.Likes) {
		var ret int32
		return ret
	}
	return *o.Likes
}

// GetLikesOk returns a tuple with the Likes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Object) GetLikesOk() (*int32, bool) {
	if o == nil || IsNil(o.Likes) {
		return nil, false
	}
	return o.Likes, true
}

// HasLikes returns a boolean if a field has been set.
func (o *Object) HasLikes() bool {
	if o != nil && !IsNil(o.Likes) {
		return true
	}

	return false
}

// SetLikes gets a reference to the given int32 and assigns it to the Likes field.
func (o *Object) SetLikes(v int32) {
	o.Likes = &v
}

// GetPublishedAt returns the PublishedAt field value if set, zero value otherwise.
func (o *Object) GetPublishedAt() string {
	if o == nil || IsNil(o.PublishedAt) {
		var ret string
		return ret
	}
	return *o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Object) GetPublishedAtOk() (*string, bool) {
	if o == nil || IsNil(o.PublishedAt) {
		return nil, false
	}
	return o.PublishedAt, true
}

// HasPublishedAt returns a boolean if a field has been set.
func (o *Object) HasPublishedAt() bool {
	if o != nil && !IsNil(o.PublishedAt) {
		return true
	}

	return false
}

// SetPublishedAt gets a reference to the given string and assigns it to the PublishedAt field.
func (o *Object) SetPublishedAt(v string) {
	o.PublishedAt = &v
}

// GetFeatured returns the Featured field value if set, zero value otherwise.
func (o *Object) GetFeatured() string {
	if o == nil || IsNil(o.Featured) {
		var ret string
		return ret
	}
	return *o.Featured
}

// GetFeaturedOk returns a tuple with the Featured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Object) GetFeaturedOk() (*string, bool) {
	if o == nil || IsNil(o.Featured) {
		return nil, false
	}
	return o.Featured, true
}

// HasFeatured returns a boolean if a field has been set.
func (o *Object) HasFeatured() bool {
	if o != nil && !IsNil(o.Featured) {
		return true
	}

	return false
}

// SetFeatured gets a reference to the given string and assigns it to the Featured field.
func (o *Object) SetFeatured(v string) {
	o.Featured = &v
}

// GetComplexity returns the Complexity field value if set, zero value otherwise.
func (o *Object) GetComplexity() int32 {
	if o == nil || IsNil(o.Complexity) {
		var ret int32
		return ret
	}
	return *o.Complexity
}

// GetComplexityOk returns a tuple with the Complexity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Object) GetComplexityOk() (*int32, bool) {
	if o == nil || IsNil(o.Complexity) {
		return nil, false
	}
	return o.Complexity, true
}

// HasComplexity returns a boolean if a field has been set.
func (o *Object) HasComplexity() bool {
	if o != nil && !IsNil(o.Complexity) {
		return true
	}

	return false
}

// SetComplexity gets a reference to the given int32 and assigns it to the Complexity field.
func (o *Object) SetComplexity(v int32) {
	o.Complexity = &v
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *Object) GetDimensions() string {
	if o == nil || IsNil(o.Dimensions) {
		var ret string
		return ret
	}
	return *o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Object) GetDimensionsOk() (*string, bool) {
	if o == nil || IsNil(o.Dimensions) {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *Object) HasDimensions() bool {
	if o != nil && !IsNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given string and assigns it to the Dimensions field.
func (o *Object) SetDimensions(v string) {
	o.Dimensions = &v
}

// GetMaterialQuantity returns the MaterialQuantity field value if set, zero value otherwise.
func (o *Object) GetMaterialQuantity() string {
	if o == nil || IsNil(o.MaterialQuantity) {
		var ret string
		return ret
	}
	return *o.MaterialQuantity
}

// GetMaterialQuantityOk returns a tuple with the MaterialQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Object) GetMaterialQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.MaterialQuantity) {
		return nil, false
	}
	return o.MaterialQuantity, true
}

// HasMaterialQuantity returns a boolean if a field has been set.
func (o *Object) HasMaterialQuantity() bool {
	if o != nil && !IsNil(o.MaterialQuantity) {
		return true
	}

	return false
}

// SetMaterialQuantity gets a reference to the given string and assigns it to the MaterialQuantity field.
func (o *Object) SetMaterialQuantity(v string) {
	o.MaterialQuantity = &v
}

// GetDesigner returns the Designer field value if set, zero value otherwise.
func (o *Object) GetDesigner() User {
	if o == nil || IsNil(o.Designer) {
		var ret User
		return ret
	}
	return *o.Designer
}

// GetDesignerOk returns a tuple with the Designer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Object) GetDesignerOk() (*User, bool) {
	if o == nil || IsNil(o.Designer) {
		return nil, false
	}
	return o.Designer, true
}

// HasDesigner returns a boolean if a field has been set.
func (o *Object) HasDesigner() bool {
	if o != nil && !IsNil(o.Designer) {
		return true
	}

	return false
}

// SetDesigner gets a reference to the given User and assigns it to the Designer field.
func (o *Object) SetDesigner(v User) {
	o.Designer = &v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *Object) GetImages() []Image {
	if o == nil || IsNil(o.Images) {
		var ret []Image
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Object) GetImagesOk() ([]Image, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *Object) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []Image and assigns it to the Images field.
func (o *Object) SetImages(v []Image) {
	o.Images = v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *Object) GetFiles() []File {
	if o == nil || IsNil(o.Files) {
		var ret []File
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Object) GetFilesOk() ([]File, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *Object) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []File and assigns it to the Files field.
func (o *Object) SetFiles(v []File) {
	o.Files = v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *Object) GetCategories() []OneCategory {
	if o == nil || IsNil(o.Categories) {
		var ret []OneCategory
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Object) GetCategoriesOk() ([]OneCategory, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *Object) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []OneCategory and assigns it to the Categories field.
func (o *Object) SetCategories(v []OneCategory) {
	o.Categories = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Object) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Object) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Object) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Object) SetTags(v []string) {
	o.Tags = v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *Object) GetLicenses() []License {
	if o == nil || IsNil(o.Licenses) {
		var ret []License
		return ret
	}
	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Object) GetLicensesOk() ([]License, bool) {
	if o == nil || IsNil(o.Licenses) {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *Object) HasLicenses() bool {
	if o != nil && !IsNil(o.Licenses) {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given []License and assigns it to the Licenses field.
func (o *Object) SetLicenses(v []License) {
	o.Licenses = v
}

func (o Object) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Object) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.ArchiveDownloadUrl) {
		toSerialize["archive_download_url"] = o.ArchiveDownloadUrl
	}
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Views) {
		toSerialize["views"] = o.Views
	}
	if !IsNil(o.Likes) {
		toSerialize["likes"] = o.Likes
	}
	if !IsNil(o.PublishedAt) {
		toSerialize["published_at"] = o.PublishedAt
	}
	if !IsNil(o.Featured) {
		toSerialize["featured"] = o.Featured
	}
	if !IsNil(o.Complexity) {
		toSerialize["complexity"] = o.Complexity
	}
	if !IsNil(o.Dimensions) {
		toSerialize["dimensions"] = o.Dimensions
	}
	if !IsNil(o.MaterialQuantity) {
		toSerialize["material_quantity"] = o.MaterialQuantity
	}
	if !IsNil(o.Designer) {
		toSerialize["designer"] = o.Designer
	}
	if !IsNil(o.Images) {
		toSerialize["images"] = o.Images
	}
	if !IsNil(o.Files) {
		toSerialize["files"] = o.Files
	}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Licenses) {
		toSerialize["licenses"] = o.Licenses
	}
	return toSerialize, nil
}

type NullableObject struct {
	value *Object
	isSet bool
}

func (v NullableObject) Get() *Object {
	return v.value
}

func (v *NullableObject) Set(val *Object) {
	v.value = val
	v.isSet = true
}

func (v NullableObject) IsSet() bool {
	return v.isSet
}

func (v *NullableObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObject(val *Object) *NullableObject {
	return &NullableObject{value: val, isSet: true}
}

func (v NullableObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


