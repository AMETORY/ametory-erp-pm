/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product_v202407

import (
    "encoding/json"
    "tiktokshop/open/sdk_golang/utils"
)

            // checks if the Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions{}

// Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions struct for Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions
type Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions struct {
    // The id of option
    Id *string `json:"id,omitempty"`
    // The name of option
    Name *string `json:"name,omitempty"`
}

// NewProduct202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions instantiates a new Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions() *Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions {
    this := Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions{}
    return &this
}

// NewProduct202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptionsWithDefaults instantiates a new Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptionsWithDefaults() *Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions {
    this := Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions{}
    return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions) SetId(v string) {
    o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions) GetName() string {
    if o == nil || utils.IsNil(o.Name) {
        var ret string
        return ret
    }
    return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions) GetNameOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Name) {
        return nil, false
    }
    return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions) HasName() bool {
    if o != nil && !utils.IsNil(o.Name) {
        return true
    }

    return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions) SetName(v string) {
    o.Name = &v
}

func (o Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    if !utils.IsNil(o.Name) {
        toSerialize["name"] = o.Name
    }
    return toSerialize, nil
}

type NullableProduct202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions struct {
	value *Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions
	isSet bool
}

func (v NullableProduct202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions) Get() *Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions {
	return v.value
}

func (v *NullableProduct202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions) Set(val *Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions(val *Product202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions) *NullableProduct202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions {
	return &NullableProduct202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions{value: val, isSet: true}
}

func (v NullableProduct202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202407ListingSchemasResponseDataListingSchemasFieldsComplexValuesOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


