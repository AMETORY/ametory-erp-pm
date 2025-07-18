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

            // checks if the Product202309SearchProductsResponseDataProductsSkusPrice type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309SearchProductsResponseDataProductsSkusPrice{}

// Product202309SearchProductsResponseDataProductsSkusPrice struct for Product202309SearchProductsResponseDataProductsSkusPrice
type Product202309SearchProductsResponseDataProductsSkusPrice struct {
    // The currency of the SKU price. Possible values: - EUR: France, Germany, Ireland, Italy, Spain - GBP: United Kingdom - IDR: Indonesia - JPY: Japan - MXN: Mexico - MYR: Malaysia - PHP: Philippines - SGD: Singapore - THB: Thailand - USD: United States - VND: Vietnam
    Currency *string `json:"currency,omitempty"`
    // The SKU's selling price, inclusive of tax. Applicable only for cross-border sellers from China.
    SalePrice *string `json:"sale_price,omitempty"`
    // The SKU's selling price, exclusive of tax.
    TaxExclusivePrice *string `json:"tax_exclusive_price,omitempty"`
}

// NewProduct202309SearchProductsResponseDataProductsSkusPrice instantiates a new Product202309SearchProductsResponseDataProductsSkusPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309SearchProductsResponseDataProductsSkusPrice() *Product202309SearchProductsResponseDataProductsSkusPrice {
    this := Product202309SearchProductsResponseDataProductsSkusPrice{}
    return &this
}

// NewProduct202309SearchProductsResponseDataProductsSkusPriceWithDefaults instantiates a new Product202309SearchProductsResponseDataProductsSkusPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309SearchProductsResponseDataProductsSkusPriceWithDefaults() *Product202309SearchProductsResponseDataProductsSkusPrice {
    this := Product202309SearchProductsResponseDataProductsSkusPrice{}
    return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Product202309SearchProductsResponseDataProductsSkusPrice) GetCurrency() string {
    if o == nil || utils.IsNil(o.Currency) {
        var ret string
        return ret
    }
    return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309SearchProductsResponseDataProductsSkusPrice) GetCurrencyOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Currency) {
        return nil, false
    }
    return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Product202309SearchProductsResponseDataProductsSkusPrice) HasCurrency() bool {
    if o != nil && !utils.IsNil(o.Currency) {
        return true
    }

    return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Product202309SearchProductsResponseDataProductsSkusPrice) SetCurrency(v string) {
    o.Currency = &v
}

// GetSalePrice returns the SalePrice field value if set, zero value otherwise.
func (o *Product202309SearchProductsResponseDataProductsSkusPrice) GetSalePrice() string {
    if o == nil || utils.IsNil(o.SalePrice) {
        var ret string
        return ret
    }
    return *o.SalePrice
}

// GetSalePriceOk returns a tuple with the SalePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309SearchProductsResponseDataProductsSkusPrice) GetSalePriceOk() (*string, bool) {
    if o == nil || utils.IsNil(o.SalePrice) {
        return nil, false
    }
    return o.SalePrice, true
}

// HasSalePrice returns a boolean if a field has been set.
func (o *Product202309SearchProductsResponseDataProductsSkusPrice) HasSalePrice() bool {
    if o != nil && !utils.IsNil(o.SalePrice) {
        return true
    }

    return false
}

// SetSalePrice gets a reference to the given string and assigns it to the SalePrice field.
func (o *Product202309SearchProductsResponseDataProductsSkusPrice) SetSalePrice(v string) {
    o.SalePrice = &v
}

// GetTaxExclusivePrice returns the TaxExclusivePrice field value if set, zero value otherwise.
func (o *Product202309SearchProductsResponseDataProductsSkusPrice) GetTaxExclusivePrice() string {
    if o == nil || utils.IsNil(o.TaxExclusivePrice) {
        var ret string
        return ret
    }
    return *o.TaxExclusivePrice
}

// GetTaxExclusivePriceOk returns a tuple with the TaxExclusivePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309SearchProductsResponseDataProductsSkusPrice) GetTaxExclusivePriceOk() (*string, bool) {
    if o == nil || utils.IsNil(o.TaxExclusivePrice) {
        return nil, false
    }
    return o.TaxExclusivePrice, true
}

// HasTaxExclusivePrice returns a boolean if a field has been set.
func (o *Product202309SearchProductsResponseDataProductsSkusPrice) HasTaxExclusivePrice() bool {
    if o != nil && !utils.IsNil(o.TaxExclusivePrice) {
        return true
    }

    return false
}

// SetTaxExclusivePrice gets a reference to the given string and assigns it to the TaxExclusivePrice field.
func (o *Product202309SearchProductsResponseDataProductsSkusPrice) SetTaxExclusivePrice(v string) {
    o.TaxExclusivePrice = &v
}

func (o Product202309SearchProductsResponseDataProductsSkusPrice) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309SearchProductsResponseDataProductsSkusPrice) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Currency) {
        toSerialize["currency"] = o.Currency
    }
    if !utils.IsNil(o.SalePrice) {
        toSerialize["sale_price"] = o.SalePrice
    }
    if !utils.IsNil(o.TaxExclusivePrice) {
        toSerialize["tax_exclusive_price"] = o.TaxExclusivePrice
    }
    return toSerialize, nil
}

type NullableProduct202309SearchProductsResponseDataProductsSkusPrice struct {
	value *Product202309SearchProductsResponseDataProductsSkusPrice
	isSet bool
}

func (v NullableProduct202309SearchProductsResponseDataProductsSkusPrice) Get() *Product202309SearchProductsResponseDataProductsSkusPrice {
	return v.value
}

func (v *NullableProduct202309SearchProductsResponseDataProductsSkusPrice) Set(val *Product202309SearchProductsResponseDataProductsSkusPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309SearchProductsResponseDataProductsSkusPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309SearchProductsResponseDataProductsSkusPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309SearchProductsResponseDataProductsSkusPrice(val *Product202309SearchProductsResponseDataProductsSkusPrice) *NullableProduct202309SearchProductsResponseDataProductsSkusPrice {
	return &NullableProduct202309SearchProductsResponseDataProductsSkusPrice{value: val, isSet: true}
}

func (v NullableProduct202309SearchProductsResponseDataProductsSkusPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309SearchProductsResponseDataProductsSkusPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


