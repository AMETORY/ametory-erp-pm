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

            // checks if the Product202309CreateProductRequestBodySkusListPrice type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309CreateProductRequestBodySkusListPrice{}

// Product202309CreateProductRequestBodySkusListPrice struct for Product202309CreateProductRequestBodySkusListPrice
type Product202309CreateProductRequestBodySkusListPrice struct {
    // The price amount. Valid range: [0.01, 7600]   **Note**:  - The value must be equal to or greater than `skus.price.amount`. Otherwise, it will be discarded. - If the value is verified to be legitimate by the audit team, it will be stored and returned in the [Get Product API](6509d85b4a0bb702c057fdda).
    Amount *string `json:"amount,omitempty"`
    // The currency. Possible values: USD
    Currency *string `json:"currency,omitempty"`
}

// NewProduct202309CreateProductRequestBodySkusListPrice instantiates a new Product202309CreateProductRequestBodySkusListPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309CreateProductRequestBodySkusListPrice() *Product202309CreateProductRequestBodySkusListPrice {
    this := Product202309CreateProductRequestBodySkusListPrice{}
    return &this
}

// NewProduct202309CreateProductRequestBodySkusListPriceWithDefaults instantiates a new Product202309CreateProductRequestBodySkusListPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309CreateProductRequestBodySkusListPriceWithDefaults() *Product202309CreateProductRequestBodySkusListPrice {
    this := Product202309CreateProductRequestBodySkusListPrice{}
    return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Product202309CreateProductRequestBodySkusListPrice) GetAmount() string {
    if o == nil || utils.IsNil(o.Amount) {
        var ret string
        return ret
    }
    return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309CreateProductRequestBodySkusListPrice) GetAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Amount) {
        return nil, false
    }
    return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Product202309CreateProductRequestBodySkusListPrice) HasAmount() bool {
    if o != nil && !utils.IsNil(o.Amount) {
        return true
    }

    return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *Product202309CreateProductRequestBodySkusListPrice) SetAmount(v string) {
    o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Product202309CreateProductRequestBodySkusListPrice) GetCurrency() string {
    if o == nil || utils.IsNil(o.Currency) {
        var ret string
        return ret
    }
    return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309CreateProductRequestBodySkusListPrice) GetCurrencyOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Currency) {
        return nil, false
    }
    return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Product202309CreateProductRequestBodySkusListPrice) HasCurrency() bool {
    if o != nil && !utils.IsNil(o.Currency) {
        return true
    }

    return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Product202309CreateProductRequestBodySkusListPrice) SetCurrency(v string) {
    o.Currency = &v
}

func (o Product202309CreateProductRequestBodySkusListPrice) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309CreateProductRequestBodySkusListPrice) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Amount) {
        toSerialize["amount"] = o.Amount
    }
    if !utils.IsNil(o.Currency) {
        toSerialize["currency"] = o.Currency
    }
    return toSerialize, nil
}

type NullableProduct202309CreateProductRequestBodySkusListPrice struct {
	value *Product202309CreateProductRequestBodySkusListPrice
	isSet bool
}

func (v NullableProduct202309CreateProductRequestBodySkusListPrice) Get() *Product202309CreateProductRequestBodySkusListPrice {
	return v.value
}

func (v *NullableProduct202309CreateProductRequestBodySkusListPrice) Set(val *Product202309CreateProductRequestBodySkusListPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309CreateProductRequestBodySkusListPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309CreateProductRequestBodySkusListPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309CreateProductRequestBodySkusListPrice(val *Product202309CreateProductRequestBodySkusListPrice) *NullableProduct202309CreateProductRequestBodySkusListPrice {
	return &NullableProduct202309CreateProductRequestBodySkusListPrice{value: val, isSet: true}
}

func (v NullableProduct202309CreateProductRequestBodySkusListPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309CreateProductRequestBodySkusListPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


