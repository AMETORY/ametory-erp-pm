/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package promotion_v202309

import (
    "encoding/json"
    "tiktokshop/open/sdk_golang/utils"
)

            // checks if the Promotion202309UpdateActivityRequestBodyDiscount type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Promotion202309UpdateActivityRequestBodyDiscount{}

// Promotion202309UpdateActivityRequestBodyDiscount struct for Promotion202309UpdateActivityRequestBodyDiscount
type Promotion202309UpdateActivityRequestBodyDiscount struct {
    BmsmDiscount *Promotion202309UpdateActivityRequestBodyDiscountBmsmDiscount `json:"bmsm_discount,omitempty"`
    ShippingDiscount *Promotion202309UpdateActivityRequestBodyDiscountShippingDiscount `json:"shipping_discount,omitempty"`
}

// NewPromotion202309UpdateActivityRequestBodyDiscount instantiates a new Promotion202309UpdateActivityRequestBodyDiscount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotion202309UpdateActivityRequestBodyDiscount() *Promotion202309UpdateActivityRequestBodyDiscount {
    this := Promotion202309UpdateActivityRequestBodyDiscount{}
    return &this
}

// NewPromotion202309UpdateActivityRequestBodyDiscountWithDefaults instantiates a new Promotion202309UpdateActivityRequestBodyDiscount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotion202309UpdateActivityRequestBodyDiscountWithDefaults() *Promotion202309UpdateActivityRequestBodyDiscount {
    this := Promotion202309UpdateActivityRequestBodyDiscount{}
    return &this
}

// GetBmsmDiscount returns the BmsmDiscount field value if set, zero value otherwise.
func (o *Promotion202309UpdateActivityRequestBodyDiscount) GetBmsmDiscount() Promotion202309UpdateActivityRequestBodyDiscountBmsmDiscount {
    if o == nil || utils.IsNil(o.BmsmDiscount) {
        var ret Promotion202309UpdateActivityRequestBodyDiscountBmsmDiscount
        return ret
    }
    return *o.BmsmDiscount
}

// GetBmsmDiscountOk returns a tuple with the BmsmDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202309UpdateActivityRequestBodyDiscount) GetBmsmDiscountOk() (*Promotion202309UpdateActivityRequestBodyDiscountBmsmDiscount, bool) {
    if o == nil || utils.IsNil(o.BmsmDiscount) {
        return nil, false
    }
    return o.BmsmDiscount, true
}

// HasBmsmDiscount returns a boolean if a field has been set.
func (o *Promotion202309UpdateActivityRequestBodyDiscount) HasBmsmDiscount() bool {
    if o != nil && !utils.IsNil(o.BmsmDiscount) {
        return true
    }

    return false
}

// SetBmsmDiscount gets a reference to the given Promotion202309UpdateActivityRequestBodyDiscountBmsmDiscount and assigns it to the BmsmDiscount field.
func (o *Promotion202309UpdateActivityRequestBodyDiscount) SetBmsmDiscount(v Promotion202309UpdateActivityRequestBodyDiscountBmsmDiscount) {
    o.BmsmDiscount = &v
}

// GetShippingDiscount returns the ShippingDiscount field value if set, zero value otherwise.
func (o *Promotion202309UpdateActivityRequestBodyDiscount) GetShippingDiscount() Promotion202309UpdateActivityRequestBodyDiscountShippingDiscount {
    if o == nil || utils.IsNil(o.ShippingDiscount) {
        var ret Promotion202309UpdateActivityRequestBodyDiscountShippingDiscount
        return ret
    }
    return *o.ShippingDiscount
}

// GetShippingDiscountOk returns a tuple with the ShippingDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202309UpdateActivityRequestBodyDiscount) GetShippingDiscountOk() (*Promotion202309UpdateActivityRequestBodyDiscountShippingDiscount, bool) {
    if o == nil || utils.IsNil(o.ShippingDiscount) {
        return nil, false
    }
    return o.ShippingDiscount, true
}

// HasShippingDiscount returns a boolean if a field has been set.
func (o *Promotion202309UpdateActivityRequestBodyDiscount) HasShippingDiscount() bool {
    if o != nil && !utils.IsNil(o.ShippingDiscount) {
        return true
    }

    return false
}

// SetShippingDiscount gets a reference to the given Promotion202309UpdateActivityRequestBodyDiscountShippingDiscount and assigns it to the ShippingDiscount field.
func (o *Promotion202309UpdateActivityRequestBodyDiscount) SetShippingDiscount(v Promotion202309UpdateActivityRequestBodyDiscountShippingDiscount) {
    o.ShippingDiscount = &v
}

func (o Promotion202309UpdateActivityRequestBodyDiscount) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Promotion202309UpdateActivityRequestBodyDiscount) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.BmsmDiscount) {
        toSerialize["bmsm_discount"] = o.BmsmDiscount
    }
    if !utils.IsNil(o.ShippingDiscount) {
        toSerialize["shipping_discount"] = o.ShippingDiscount
    }
    return toSerialize, nil
}

type NullablePromotion202309UpdateActivityRequestBodyDiscount struct {
	value *Promotion202309UpdateActivityRequestBodyDiscount
	isSet bool
}

func (v NullablePromotion202309UpdateActivityRequestBodyDiscount) Get() *Promotion202309UpdateActivityRequestBodyDiscount {
	return v.value
}

func (v *NullablePromotion202309UpdateActivityRequestBodyDiscount) Set(val *Promotion202309UpdateActivityRequestBodyDiscount) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotion202309UpdateActivityRequestBodyDiscount) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotion202309UpdateActivityRequestBodyDiscount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotion202309UpdateActivityRequestBodyDiscount(val *Promotion202309UpdateActivityRequestBodyDiscount) *NullablePromotion202309UpdateActivityRequestBodyDiscount {
	return &NullablePromotion202309UpdateActivityRequestBodyDiscount{value: val, isSet: true}
}

func (v NullablePromotion202309UpdateActivityRequestBodyDiscount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotion202309UpdateActivityRequestBodyDiscount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


