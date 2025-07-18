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

            // checks if the Product202309EditGlobalProductRequestBodyPackageWeight type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309EditGlobalProductRequestBodyPackageWeight{}

// Product202309EditGlobalProductRequestBodyPackageWeight struct for Product202309EditGlobalProductRequestBodyPackageWeight
type Product202309EditGlobalProductRequestBodyPackageWeight struct {
    // The unit for the package weight. Only `KILOGRAM` is supported.
    Unit *string `json:"unit,omitempty"`
    // The package weight, which must be a positive number with up to 3 decimal places.
    Value *string `json:"value,omitempty"`
}

// NewProduct202309EditGlobalProductRequestBodyPackageWeight instantiates a new Product202309EditGlobalProductRequestBodyPackageWeight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309EditGlobalProductRequestBodyPackageWeight() *Product202309EditGlobalProductRequestBodyPackageWeight {
    this := Product202309EditGlobalProductRequestBodyPackageWeight{}
    return &this
}

// NewProduct202309EditGlobalProductRequestBodyPackageWeightWithDefaults instantiates a new Product202309EditGlobalProductRequestBodyPackageWeight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309EditGlobalProductRequestBodyPackageWeightWithDefaults() *Product202309EditGlobalProductRequestBodyPackageWeight {
    this := Product202309EditGlobalProductRequestBodyPackageWeight{}
    return &this
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *Product202309EditGlobalProductRequestBodyPackageWeight) GetUnit() string {
    if o == nil || utils.IsNil(o.Unit) {
        var ret string
        return ret
    }
    return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309EditGlobalProductRequestBodyPackageWeight) GetUnitOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Unit) {
        return nil, false
    }
    return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *Product202309EditGlobalProductRequestBodyPackageWeight) HasUnit() bool {
    if o != nil && !utils.IsNil(o.Unit) {
        return true
    }

    return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *Product202309EditGlobalProductRequestBodyPackageWeight) SetUnit(v string) {
    o.Unit = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Product202309EditGlobalProductRequestBodyPackageWeight) GetValue() string {
    if o == nil || utils.IsNil(o.Value) {
        var ret string
        return ret
    }
    return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309EditGlobalProductRequestBodyPackageWeight) GetValueOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Value) {
        return nil, false
    }
    return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Product202309EditGlobalProductRequestBodyPackageWeight) HasValue() bool {
    if o != nil && !utils.IsNil(o.Value) {
        return true
    }

    return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *Product202309EditGlobalProductRequestBodyPackageWeight) SetValue(v string) {
    o.Value = &v
}

func (o Product202309EditGlobalProductRequestBodyPackageWeight) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309EditGlobalProductRequestBodyPackageWeight) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Unit) {
        toSerialize["unit"] = o.Unit
    }
    if !utils.IsNil(o.Value) {
        toSerialize["value"] = o.Value
    }
    return toSerialize, nil
}

type NullableProduct202309EditGlobalProductRequestBodyPackageWeight struct {
	value *Product202309EditGlobalProductRequestBodyPackageWeight
	isSet bool
}

func (v NullableProduct202309EditGlobalProductRequestBodyPackageWeight) Get() *Product202309EditGlobalProductRequestBodyPackageWeight {
	return v.value
}

func (v *NullableProduct202309EditGlobalProductRequestBodyPackageWeight) Set(val *Product202309EditGlobalProductRequestBodyPackageWeight) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309EditGlobalProductRequestBodyPackageWeight) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309EditGlobalProductRequestBodyPackageWeight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309EditGlobalProductRequestBodyPackageWeight(val *Product202309EditGlobalProductRequestBodyPackageWeight) *NullableProduct202309EditGlobalProductRequestBodyPackageWeight {
	return &NullableProduct202309EditGlobalProductRequestBodyPackageWeight{value: val, isSet: true}
}

func (v NullableProduct202309EditGlobalProductRequestBodyPackageWeight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309EditGlobalProductRequestBodyPackageWeight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


