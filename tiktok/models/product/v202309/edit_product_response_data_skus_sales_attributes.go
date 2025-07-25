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

            // checks if the Product202309EditProductResponseDataSkusSalesAttributes type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309EditProductResponseDataSkusSalesAttributes{}

// Product202309EditProductResponseDataSkusSalesAttributes struct for Product202309EditProductResponseDataSkusSalesAttributes
type Product202309EditProductResponseDataSkusSalesAttributes struct {
    // The sales attribute ID.  If you included the custom sales attribute name in the request, this is a newly generated ID.
    Id *string `json:"id,omitempty"`
    // The sales attribute value ID.  If you included the custom sales attribute value name in the request, this is a newly generated ID.
    ValueId *string `json:"value_id,omitempty"`
}

// NewProduct202309EditProductResponseDataSkusSalesAttributes instantiates a new Product202309EditProductResponseDataSkusSalesAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309EditProductResponseDataSkusSalesAttributes() *Product202309EditProductResponseDataSkusSalesAttributes {
    this := Product202309EditProductResponseDataSkusSalesAttributes{}
    return &this
}

// NewProduct202309EditProductResponseDataSkusSalesAttributesWithDefaults instantiates a new Product202309EditProductResponseDataSkusSalesAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309EditProductResponseDataSkusSalesAttributesWithDefaults() *Product202309EditProductResponseDataSkusSalesAttributes {
    this := Product202309EditProductResponseDataSkusSalesAttributes{}
    return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Product202309EditProductResponseDataSkusSalesAttributes) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309EditProductResponseDataSkusSalesAttributes) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Product202309EditProductResponseDataSkusSalesAttributes) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Product202309EditProductResponseDataSkusSalesAttributes) SetId(v string) {
    o.Id = &v
}

// GetValueId returns the ValueId field value if set, zero value otherwise.
func (o *Product202309EditProductResponseDataSkusSalesAttributes) GetValueId() string {
    if o == nil || utils.IsNil(o.ValueId) {
        var ret string
        return ret
    }
    return *o.ValueId
}

// GetValueIdOk returns a tuple with the ValueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309EditProductResponseDataSkusSalesAttributes) GetValueIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ValueId) {
        return nil, false
    }
    return o.ValueId, true
}

// HasValueId returns a boolean if a field has been set.
func (o *Product202309EditProductResponseDataSkusSalesAttributes) HasValueId() bool {
    if o != nil && !utils.IsNil(o.ValueId) {
        return true
    }

    return false
}

// SetValueId gets a reference to the given string and assigns it to the ValueId field.
func (o *Product202309EditProductResponseDataSkusSalesAttributes) SetValueId(v string) {
    o.ValueId = &v
}

func (o Product202309EditProductResponseDataSkusSalesAttributes) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309EditProductResponseDataSkusSalesAttributes) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    if !utils.IsNil(o.ValueId) {
        toSerialize["value_id"] = o.ValueId
    }
    return toSerialize, nil
}

type NullableProduct202309EditProductResponseDataSkusSalesAttributes struct {
	value *Product202309EditProductResponseDataSkusSalesAttributes
	isSet bool
}

func (v NullableProduct202309EditProductResponseDataSkusSalesAttributes) Get() *Product202309EditProductResponseDataSkusSalesAttributes {
	return v.value
}

func (v *NullableProduct202309EditProductResponseDataSkusSalesAttributes) Set(val *Product202309EditProductResponseDataSkusSalesAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309EditProductResponseDataSkusSalesAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309EditProductResponseDataSkusSalesAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309EditProductResponseDataSkusSalesAttributes(val *Product202309EditProductResponseDataSkusSalesAttributes) *NullableProduct202309EditProductResponseDataSkusSalesAttributes {
	return &NullableProduct202309EditProductResponseDataSkusSalesAttributes{value: val, isSet: true}
}

func (v NullableProduct202309EditProductResponseDataSkusSalesAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309EditProductResponseDataSkusSalesAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


