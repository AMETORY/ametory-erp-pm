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

            // checks if the Product202309GetAttributesResponseData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309GetAttributesResponseData{}

// Product202309GetAttributesResponseData struct for Product202309GetAttributesResponseData
type Product202309GetAttributesResponseData struct {
    // The list of standard built-in product and sales attributes that are bound to the specified category, based on your shop's location.
    Attributes []Product202309GetAttributesResponseDataAttributes `json:"attributes,omitempty"`
}

// NewProduct202309GetAttributesResponseData instantiates a new Product202309GetAttributesResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309GetAttributesResponseData() *Product202309GetAttributesResponseData {
    this := Product202309GetAttributesResponseData{}
    return &this
}

// NewProduct202309GetAttributesResponseDataWithDefaults instantiates a new Product202309GetAttributesResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309GetAttributesResponseDataWithDefaults() *Product202309GetAttributesResponseData {
    this := Product202309GetAttributesResponseData{}
    return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Product202309GetAttributesResponseData) GetAttributes() []Product202309GetAttributesResponseDataAttributes {
    if o == nil || utils.IsNil(o.Attributes) {
        var ret []Product202309GetAttributesResponseDataAttributes
        return ret
    }
    return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309GetAttributesResponseData) GetAttributesOk() ([]Product202309GetAttributesResponseDataAttributes, bool) {
    if o == nil || utils.IsNil(o.Attributes) {
        return nil, false
    }
    return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Product202309GetAttributesResponseData) HasAttributes() bool {
    if o != nil && !utils.IsNil(o.Attributes) {
        return true
    }

    return false
}

// SetAttributes gets a reference to the given []Product202309GetAttributesResponseDataAttributes and assigns it to the Attributes field.
func (o *Product202309GetAttributesResponseData) SetAttributes(v []Product202309GetAttributesResponseDataAttributes) {
    o.Attributes = v
}

func (o Product202309GetAttributesResponseData) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309GetAttributesResponseData) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Attributes) {
        toSerialize["attributes"] = o.Attributes
    }
    return toSerialize, nil
}

type NullableProduct202309GetAttributesResponseData struct {
	value *Product202309GetAttributesResponseData
	isSet bool
}

func (v NullableProduct202309GetAttributesResponseData) Get() *Product202309GetAttributesResponseData {
	return v.value
}

func (v *NullableProduct202309GetAttributesResponseData) Set(val *Product202309GetAttributesResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309GetAttributesResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309GetAttributesResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309GetAttributesResponseData(val *Product202309GetAttributesResponseData) *NullableProduct202309GetAttributesResponseData {
	return &NullableProduct202309GetAttributesResponseData{value: val, isSet: true}
}

func (v NullableProduct202309GetAttributesResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309GetAttributesResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


