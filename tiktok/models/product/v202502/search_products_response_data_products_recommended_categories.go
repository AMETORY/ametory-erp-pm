/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product_v202502

import (
    "encoding/json"
    "tiktokshop/open/sdk_golang/utils"
)

            // checks if the Product202502SearchProductsResponseDataProductsRecommendedCategories type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202502SearchProductsResponseDataProductsRecommendedCategories{}

// Product202502SearchProductsResponseDataProductsRecommendedCategories struct for Product202502SearchProductsResponseDataProductsRecommendedCategories
type Product202502SearchProductsResponseDataProductsRecommendedCategories struct {
    // The ID of the recommended category.
    Id *string `json:"id,omitempty"`
    // The name of the category in the country where the shop operates.
    LocalName *string `json:"local_name,omitempty"`
}

// NewProduct202502SearchProductsResponseDataProductsRecommendedCategories instantiates a new Product202502SearchProductsResponseDataProductsRecommendedCategories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202502SearchProductsResponseDataProductsRecommendedCategories() *Product202502SearchProductsResponseDataProductsRecommendedCategories {
    this := Product202502SearchProductsResponseDataProductsRecommendedCategories{}
    return &this
}

// NewProduct202502SearchProductsResponseDataProductsRecommendedCategoriesWithDefaults instantiates a new Product202502SearchProductsResponseDataProductsRecommendedCategories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202502SearchProductsResponseDataProductsRecommendedCategoriesWithDefaults() *Product202502SearchProductsResponseDataProductsRecommendedCategories {
    this := Product202502SearchProductsResponseDataProductsRecommendedCategories{}
    return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Product202502SearchProductsResponseDataProductsRecommendedCategories) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202502SearchProductsResponseDataProductsRecommendedCategories) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Product202502SearchProductsResponseDataProductsRecommendedCategories) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Product202502SearchProductsResponseDataProductsRecommendedCategories) SetId(v string) {
    o.Id = &v
}

// GetLocalName returns the LocalName field value if set, zero value otherwise.
func (o *Product202502SearchProductsResponseDataProductsRecommendedCategories) GetLocalName() string {
    if o == nil || utils.IsNil(o.LocalName) {
        var ret string
        return ret
    }
    return *o.LocalName
}

// GetLocalNameOk returns a tuple with the LocalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202502SearchProductsResponseDataProductsRecommendedCategories) GetLocalNameOk() (*string, bool) {
    if o == nil || utils.IsNil(o.LocalName) {
        return nil, false
    }
    return o.LocalName, true
}

// HasLocalName returns a boolean if a field has been set.
func (o *Product202502SearchProductsResponseDataProductsRecommendedCategories) HasLocalName() bool {
    if o != nil && !utils.IsNil(o.LocalName) {
        return true
    }

    return false
}

// SetLocalName gets a reference to the given string and assigns it to the LocalName field.
func (o *Product202502SearchProductsResponseDataProductsRecommendedCategories) SetLocalName(v string) {
    o.LocalName = &v
}

func (o Product202502SearchProductsResponseDataProductsRecommendedCategories) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202502SearchProductsResponseDataProductsRecommendedCategories) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    if !utils.IsNil(o.LocalName) {
        toSerialize["local_name"] = o.LocalName
    }
    return toSerialize, nil
}

type NullableProduct202502SearchProductsResponseDataProductsRecommendedCategories struct {
	value *Product202502SearchProductsResponseDataProductsRecommendedCategories
	isSet bool
}

func (v NullableProduct202502SearchProductsResponseDataProductsRecommendedCategories) Get() *Product202502SearchProductsResponseDataProductsRecommendedCategories {
	return v.value
}

func (v *NullableProduct202502SearchProductsResponseDataProductsRecommendedCategories) Set(val *Product202502SearchProductsResponseDataProductsRecommendedCategories) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202502SearchProductsResponseDataProductsRecommendedCategories) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202502SearchProductsResponseDataProductsRecommendedCategories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202502SearchProductsResponseDataProductsRecommendedCategories(val *Product202502SearchProductsResponseDataProductsRecommendedCategories) *NullableProduct202502SearchProductsResponseDataProductsRecommendedCategories {
	return &NullableProduct202502SearchProductsResponseDataProductsRecommendedCategories{value: val, isSet: true}
}

func (v NullableProduct202502SearchProductsResponseDataProductsRecommendedCategories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202502SearchProductsResponseDataProductsRecommendedCategories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


