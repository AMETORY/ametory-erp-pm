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

            // checks if the Product202502SearchProductsResponseDataProductsProductFamilies type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202502SearchProductsResponseDataProductsProductFamilies{}

// Product202502SearchProductsResponseDataProductsProductFamilies struct for Product202502SearchProductsResponseDataProductsProductFamilies
type Product202502SearchProductsResponseDataProductsProductFamilies struct {
    // The product family ID.
    Id *string `json:"id,omitempty"`
    // A list of products that belong to the family.
    Products []Product202502SearchProductsResponseDataProductsProductFamiliesProducts `json:"products,omitempty"`
}

// NewProduct202502SearchProductsResponseDataProductsProductFamilies instantiates a new Product202502SearchProductsResponseDataProductsProductFamilies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202502SearchProductsResponseDataProductsProductFamilies() *Product202502SearchProductsResponseDataProductsProductFamilies {
    this := Product202502SearchProductsResponseDataProductsProductFamilies{}
    return &this
}

// NewProduct202502SearchProductsResponseDataProductsProductFamiliesWithDefaults instantiates a new Product202502SearchProductsResponseDataProductsProductFamilies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202502SearchProductsResponseDataProductsProductFamiliesWithDefaults() *Product202502SearchProductsResponseDataProductsProductFamilies {
    this := Product202502SearchProductsResponseDataProductsProductFamilies{}
    return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Product202502SearchProductsResponseDataProductsProductFamilies) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202502SearchProductsResponseDataProductsProductFamilies) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Product202502SearchProductsResponseDataProductsProductFamilies) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Product202502SearchProductsResponseDataProductsProductFamilies) SetId(v string) {
    o.Id = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *Product202502SearchProductsResponseDataProductsProductFamilies) GetProducts() []Product202502SearchProductsResponseDataProductsProductFamiliesProducts {
    if o == nil || utils.IsNil(o.Products) {
        var ret []Product202502SearchProductsResponseDataProductsProductFamiliesProducts
        return ret
    }
    return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202502SearchProductsResponseDataProductsProductFamilies) GetProductsOk() ([]Product202502SearchProductsResponseDataProductsProductFamiliesProducts, bool) {
    if o == nil || utils.IsNil(o.Products) {
        return nil, false
    }
    return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *Product202502SearchProductsResponseDataProductsProductFamilies) HasProducts() bool {
    if o != nil && !utils.IsNil(o.Products) {
        return true
    }

    return false
}

// SetProducts gets a reference to the given []Product202502SearchProductsResponseDataProductsProductFamiliesProducts and assigns it to the Products field.
func (o *Product202502SearchProductsResponseDataProductsProductFamilies) SetProducts(v []Product202502SearchProductsResponseDataProductsProductFamiliesProducts) {
    o.Products = v
}

func (o Product202502SearchProductsResponseDataProductsProductFamilies) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202502SearchProductsResponseDataProductsProductFamilies) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    if !utils.IsNil(o.Products) {
        toSerialize["products"] = o.Products
    }
    return toSerialize, nil
}

type NullableProduct202502SearchProductsResponseDataProductsProductFamilies struct {
	value *Product202502SearchProductsResponseDataProductsProductFamilies
	isSet bool
}

func (v NullableProduct202502SearchProductsResponseDataProductsProductFamilies) Get() *Product202502SearchProductsResponseDataProductsProductFamilies {
	return v.value
}

func (v *NullableProduct202502SearchProductsResponseDataProductsProductFamilies) Set(val *Product202502SearchProductsResponseDataProductsProductFamilies) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202502SearchProductsResponseDataProductsProductFamilies) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202502SearchProductsResponseDataProductsProductFamilies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202502SearchProductsResponseDataProductsProductFamilies(val *Product202502SearchProductsResponseDataProductsProductFamilies) *NullableProduct202502SearchProductsResponseDataProductsProductFamilies {
	return &NullableProduct202502SearchProductsResponseDataProductsProductFamilies{value: val, isSet: true}
}

func (v NullableProduct202502SearchProductsResponseDataProductsProductFamilies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202502SearchProductsResponseDataProductsProductFamilies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


