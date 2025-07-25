/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product_v202405

import (
    "encoding/json"
    "tiktokshop/open/sdk_golang/utils"
)

            // checks if the Product202405GetRecommendedProductTitleAndDescriptionResponseData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202405GetRecommendedProductTitleAndDescriptionResponseData{}

// Product202405GetRecommendedProductTitleAndDescriptionResponseData struct for Product202405GetRecommendedProductTitleAndDescriptionResponseData
type Product202405GetRecommendedProductTitleAndDescriptionResponseData struct {
    // The list of requested products and the corresponding suggestions.
    Products []Product202405GetRecommendedProductTitleAndDescriptionResponseDataProducts `json:"products,omitempty"`
}

// NewProduct202405GetRecommendedProductTitleAndDescriptionResponseData instantiates a new Product202405GetRecommendedProductTitleAndDescriptionResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202405GetRecommendedProductTitleAndDescriptionResponseData() *Product202405GetRecommendedProductTitleAndDescriptionResponseData {
    this := Product202405GetRecommendedProductTitleAndDescriptionResponseData{}
    return &this
}

// NewProduct202405GetRecommendedProductTitleAndDescriptionResponseDataWithDefaults instantiates a new Product202405GetRecommendedProductTitleAndDescriptionResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202405GetRecommendedProductTitleAndDescriptionResponseDataWithDefaults() *Product202405GetRecommendedProductTitleAndDescriptionResponseData {
    this := Product202405GetRecommendedProductTitleAndDescriptionResponseData{}
    return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *Product202405GetRecommendedProductTitleAndDescriptionResponseData) GetProducts() []Product202405GetRecommendedProductTitleAndDescriptionResponseDataProducts {
    if o == nil || utils.IsNil(o.Products) {
        var ret []Product202405GetRecommendedProductTitleAndDescriptionResponseDataProducts
        return ret
    }
    return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202405GetRecommendedProductTitleAndDescriptionResponseData) GetProductsOk() ([]Product202405GetRecommendedProductTitleAndDescriptionResponseDataProducts, bool) {
    if o == nil || utils.IsNil(o.Products) {
        return nil, false
    }
    return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *Product202405GetRecommendedProductTitleAndDescriptionResponseData) HasProducts() bool {
    if o != nil && !utils.IsNil(o.Products) {
        return true
    }

    return false
}

// SetProducts gets a reference to the given []Product202405GetRecommendedProductTitleAndDescriptionResponseDataProducts and assigns it to the Products field.
func (o *Product202405GetRecommendedProductTitleAndDescriptionResponseData) SetProducts(v []Product202405GetRecommendedProductTitleAndDescriptionResponseDataProducts) {
    o.Products = v
}

func (o Product202405GetRecommendedProductTitleAndDescriptionResponseData) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202405GetRecommendedProductTitleAndDescriptionResponseData) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Products) {
        toSerialize["products"] = o.Products
    }
    return toSerialize, nil
}

type NullableProduct202405GetRecommendedProductTitleAndDescriptionResponseData struct {
	value *Product202405GetRecommendedProductTitleAndDescriptionResponseData
	isSet bool
}

func (v NullableProduct202405GetRecommendedProductTitleAndDescriptionResponseData) Get() *Product202405GetRecommendedProductTitleAndDescriptionResponseData {
	return v.value
}

func (v *NullableProduct202405GetRecommendedProductTitleAndDescriptionResponseData) Set(val *Product202405GetRecommendedProductTitleAndDescriptionResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202405GetRecommendedProductTitleAndDescriptionResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202405GetRecommendedProductTitleAndDescriptionResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202405GetRecommendedProductTitleAndDescriptionResponseData(val *Product202405GetRecommendedProductTitleAndDescriptionResponseData) *NullableProduct202405GetRecommendedProductTitleAndDescriptionResponseData {
	return &NullableProduct202405GetRecommendedProductTitleAndDescriptionResponseData{value: val, isSet: true}
}

func (v NullableProduct202405GetRecommendedProductTitleAndDescriptionResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202405GetRecommendedProductTitleAndDescriptionResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


