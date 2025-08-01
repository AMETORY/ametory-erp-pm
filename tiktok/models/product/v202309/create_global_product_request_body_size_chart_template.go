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

            // checks if the Product202309CreateGlobalProductRequestBodySizeChartTemplate type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309CreateGlobalProductRequestBodySizeChartTemplate{}

// Product202309CreateGlobalProductRequestBodySizeChartTemplate struct for Product202309CreateGlobalProductRequestBodySizeChartTemplate
type Product202309CreateGlobalProductRequestBodySizeChartTemplate struct {
    // The size chart template ID.
    Id *string `json:"id,omitempty"`
}

// NewProduct202309CreateGlobalProductRequestBodySizeChartTemplate instantiates a new Product202309CreateGlobalProductRequestBodySizeChartTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309CreateGlobalProductRequestBodySizeChartTemplate() *Product202309CreateGlobalProductRequestBodySizeChartTemplate {
    this := Product202309CreateGlobalProductRequestBodySizeChartTemplate{}
    return &this
}

// NewProduct202309CreateGlobalProductRequestBodySizeChartTemplateWithDefaults instantiates a new Product202309CreateGlobalProductRequestBodySizeChartTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309CreateGlobalProductRequestBodySizeChartTemplateWithDefaults() *Product202309CreateGlobalProductRequestBodySizeChartTemplate {
    this := Product202309CreateGlobalProductRequestBodySizeChartTemplate{}
    return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Product202309CreateGlobalProductRequestBodySizeChartTemplate) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309CreateGlobalProductRequestBodySizeChartTemplate) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Product202309CreateGlobalProductRequestBodySizeChartTemplate) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Product202309CreateGlobalProductRequestBodySizeChartTemplate) SetId(v string) {
    o.Id = &v
}

func (o Product202309CreateGlobalProductRequestBodySizeChartTemplate) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309CreateGlobalProductRequestBodySizeChartTemplate) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    return toSerialize, nil
}

type NullableProduct202309CreateGlobalProductRequestBodySizeChartTemplate struct {
	value *Product202309CreateGlobalProductRequestBodySizeChartTemplate
	isSet bool
}

func (v NullableProduct202309CreateGlobalProductRequestBodySizeChartTemplate) Get() *Product202309CreateGlobalProductRequestBodySizeChartTemplate {
	return v.value
}

func (v *NullableProduct202309CreateGlobalProductRequestBodySizeChartTemplate) Set(val *Product202309CreateGlobalProductRequestBodySizeChartTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309CreateGlobalProductRequestBodySizeChartTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309CreateGlobalProductRequestBodySizeChartTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309CreateGlobalProductRequestBodySizeChartTemplate(val *Product202309CreateGlobalProductRequestBodySizeChartTemplate) *NullableProduct202309CreateGlobalProductRequestBodySizeChartTemplate {
	return &NullableProduct202309CreateGlobalProductRequestBodySizeChartTemplate{value: val, isSet: true}
}

func (v NullableProduct202309CreateGlobalProductRequestBodySizeChartTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309CreateGlobalProductRequestBodySizeChartTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


