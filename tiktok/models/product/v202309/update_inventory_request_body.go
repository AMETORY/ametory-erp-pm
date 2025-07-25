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

            // checks if the Product202309UpdateInventoryRequestBody type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309UpdateInventoryRequestBody{}

// Product202309UpdateInventoryRequestBody struct for Product202309UpdateInventoryRequestBody
type Product202309UpdateInventoryRequestBody struct {
    // A list of Stock Keeping Units (SKUs) used to identify distinct variants of the product.
    Skus []Product202309UpdateInventoryRequestBodySkus `json:"skus,omitempty"`
}

// NewProduct202309UpdateInventoryRequestBody instantiates a new Product202309UpdateInventoryRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309UpdateInventoryRequestBody() *Product202309UpdateInventoryRequestBody {
    this := Product202309UpdateInventoryRequestBody{}
    return &this
}

// NewProduct202309UpdateInventoryRequestBodyWithDefaults instantiates a new Product202309UpdateInventoryRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309UpdateInventoryRequestBodyWithDefaults() *Product202309UpdateInventoryRequestBody {
    this := Product202309UpdateInventoryRequestBody{}
    return &this
}

// GetSkus returns the Skus field value if set, zero value otherwise.
func (o *Product202309UpdateInventoryRequestBody) GetSkus() []Product202309UpdateInventoryRequestBodySkus {
    if o == nil || utils.IsNil(o.Skus) {
        var ret []Product202309UpdateInventoryRequestBodySkus
        return ret
    }
    return o.Skus
}

// GetSkusOk returns a tuple with the Skus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309UpdateInventoryRequestBody) GetSkusOk() ([]Product202309UpdateInventoryRequestBodySkus, bool) {
    if o == nil || utils.IsNil(o.Skus) {
        return nil, false
    }
    return o.Skus, true
}

// HasSkus returns a boolean if a field has been set.
func (o *Product202309UpdateInventoryRequestBody) HasSkus() bool {
    if o != nil && !utils.IsNil(o.Skus) {
        return true
    }

    return false
}

// SetSkus gets a reference to the given []Product202309UpdateInventoryRequestBodySkus and assigns it to the Skus field.
func (o *Product202309UpdateInventoryRequestBody) SetSkus(v []Product202309UpdateInventoryRequestBodySkus) {
    o.Skus = v
}

func (o Product202309UpdateInventoryRequestBody) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309UpdateInventoryRequestBody) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Skus) {
        toSerialize["skus"] = o.Skus
    }
    return toSerialize, nil
}

type NullableProduct202309UpdateInventoryRequestBody struct {
	value *Product202309UpdateInventoryRequestBody
	isSet bool
}

func (v NullableProduct202309UpdateInventoryRequestBody) Get() *Product202309UpdateInventoryRequestBody {
	return v.value
}

func (v *NullableProduct202309UpdateInventoryRequestBody) Set(val *Product202309UpdateInventoryRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309UpdateInventoryRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309UpdateInventoryRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309UpdateInventoryRequestBody(val *Product202309UpdateInventoryRequestBody) *NullableProduct202309UpdateInventoryRequestBody {
	return &NullableProduct202309UpdateInventoryRequestBody{value: val, isSet: true}
}

func (v NullableProduct202309UpdateInventoryRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309UpdateInventoryRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


