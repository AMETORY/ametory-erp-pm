/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product_v202309

import (
    "encoding/json"
    "tiktokshop/open/sdk_golang/utils"
)

            // checks if the Product202309GetProductResponseDataProductFamilies type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309GetProductResponseDataProductFamilies{}

// Product202309GetProductResponseDataProductFamilies struct for Product202309GetProductResponseDataProductFamilies
type Product202309GetProductResponseDataProductFamilies struct {
    // The product family ID.
    Id *string `json:"id,omitempty"`
    // A list of products that belong to the family.
    Products []Product202309GetProductResponseDataProductFamiliesProducts `json:"products,omitempty"`
}

// NewProduct202309GetProductResponseDataProductFamilies instantiates a new Product202309GetProductResponseDataProductFamilies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309GetProductResponseDataProductFamilies() *Product202309GetProductResponseDataProductFamilies {
    this := Product202309GetProductResponseDataProductFamilies{}
    return &this
}

// NewProduct202309GetProductResponseDataProductFamiliesWithDefaults instantiates a new Product202309GetProductResponseDataProductFamilies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309GetProductResponseDataProductFamiliesWithDefaults() *Product202309GetProductResponseDataProductFamilies {
    this := Product202309GetProductResponseDataProductFamilies{}
    return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Product202309GetProductResponseDataProductFamilies) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309GetProductResponseDataProductFamilies) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Product202309GetProductResponseDataProductFamilies) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Product202309GetProductResponseDataProductFamilies) SetId(v string) {
    o.Id = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *Product202309GetProductResponseDataProductFamilies) GetProducts() []Product202309GetProductResponseDataProductFamiliesProducts {
    if o == nil || utils.IsNil(o.Products) {
        var ret []Product202309GetProductResponseDataProductFamiliesProducts
        return ret
    }
    return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309GetProductResponseDataProductFamilies) GetProductsOk() ([]Product202309GetProductResponseDataProductFamiliesProducts, bool) {
    if o == nil || utils.IsNil(o.Products) {
        return nil, false
    }
    return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *Product202309GetProductResponseDataProductFamilies) HasProducts() bool {
    if o != nil && !utils.IsNil(o.Products) {
        return true
    }

    return false
}

// SetProducts gets a reference to the given []Product202309GetProductResponseDataProductFamiliesProducts and assigns it to the Products field.
func (o *Product202309GetProductResponseDataProductFamilies) SetProducts(v []Product202309GetProductResponseDataProductFamiliesProducts) {
    o.Products = v
}

func (o Product202309GetProductResponseDataProductFamilies) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309GetProductResponseDataProductFamilies) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    if !utils.IsNil(o.Products) {
        toSerialize["products"] = o.Products
    }
    return toSerialize, nil
}

type NullableProduct202309GetProductResponseDataProductFamilies struct {
	value *Product202309GetProductResponseDataProductFamilies
	isSet bool
}

func (v NullableProduct202309GetProductResponseDataProductFamilies) Get() *Product202309GetProductResponseDataProductFamilies {
	return v.value
}

func (v *NullableProduct202309GetProductResponseDataProductFamilies) Set(val *Product202309GetProductResponseDataProductFamilies) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309GetProductResponseDataProductFamilies) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309GetProductResponseDataProductFamilies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309GetProductResponseDataProductFamilies(val *Product202309GetProductResponseDataProductFamilies) *NullableProduct202309GetProductResponseDataProductFamilies {
	return &NullableProduct202309GetProductResponseDataProductFamilies{value: val, isSet: true}
}

func (v NullableProduct202309GetProductResponseDataProductFamilies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309GetProductResponseDataProductFamilies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


