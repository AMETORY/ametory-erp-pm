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

            // checks if the Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes{}

// Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes struct for Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes
type Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes struct {
    // The newly generated local sales attribute ID.
    Id *string `json:"id,omitempty"`
    // The newly generated local sales attribute value ID.
    ValueId *string `json:"value_id,omitempty"`
}

// NewProduct202309PublishGlobalProductResponseDataProductsSkusSaleAttributes instantiates a new Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309PublishGlobalProductResponseDataProductsSkusSaleAttributes() *Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes {
    this := Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes{}
    return &this
}

// NewProduct202309PublishGlobalProductResponseDataProductsSkusSaleAttributesWithDefaults instantiates a new Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309PublishGlobalProductResponseDataProductsSkusSaleAttributesWithDefaults() *Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes {
    this := Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes{}
    return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes) SetId(v string) {
    o.Id = &v
}

// GetValueId returns the ValueId field value if set, zero value otherwise.
func (o *Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes) GetValueId() string {
    if o == nil || utils.IsNil(o.ValueId) {
        var ret string
        return ret
    }
    return *o.ValueId
}

// GetValueIdOk returns a tuple with the ValueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes) GetValueIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ValueId) {
        return nil, false
    }
    return o.ValueId, true
}

// HasValueId returns a boolean if a field has been set.
func (o *Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes) HasValueId() bool {
    if o != nil && !utils.IsNil(o.ValueId) {
        return true
    }

    return false
}

// SetValueId gets a reference to the given string and assigns it to the ValueId field.
func (o *Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes) SetValueId(v string) {
    o.ValueId = &v
}

func (o Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    if !utils.IsNil(o.ValueId) {
        toSerialize["value_id"] = o.ValueId
    }
    return toSerialize, nil
}

type NullableProduct202309PublishGlobalProductResponseDataProductsSkusSaleAttributes struct {
	value *Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes
	isSet bool
}

func (v NullableProduct202309PublishGlobalProductResponseDataProductsSkusSaleAttributes) Get() *Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes {
	return v.value
}

func (v *NullableProduct202309PublishGlobalProductResponseDataProductsSkusSaleAttributes) Set(val *Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309PublishGlobalProductResponseDataProductsSkusSaleAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309PublishGlobalProductResponseDataProductsSkusSaleAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309PublishGlobalProductResponseDataProductsSkusSaleAttributes(val *Product202309PublishGlobalProductResponseDataProductsSkusSaleAttributes) *NullableProduct202309PublishGlobalProductResponseDataProductsSkusSaleAttributes {
	return &NullableProduct202309PublishGlobalProductResponseDataProductsSkusSaleAttributes{value: val, isSet: true}
}

func (v NullableProduct202309PublishGlobalProductResponseDataProductsSkusSaleAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309PublishGlobalProductResponseDataProductsSkusSaleAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


