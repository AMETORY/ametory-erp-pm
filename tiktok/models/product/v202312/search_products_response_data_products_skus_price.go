/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product_v202312

import (
    "encoding/json"
    "tiktokshop/open/sdk_golang/utils"
)

            // checks if the Product202312SearchProductsResponseDataProductsSkusPrice type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202312SearchProductsResponseDataProductsSkusPrice{}

// Product202312SearchProductsResponseDataProductsSkusPrice struct for Product202312SearchProductsResponseDataProductsSkusPrice
type Product202312SearchProductsResponseDataProductsSkusPrice struct {
    // The currency. Possible values: - BRL: Brazil - EUR: France, Germany, Ireland, Italy, Spain - GBP: United Kingdom - IDR: Indonesia - JPY: Japan - MXN: Mexico - MYR: Malaysia - PHP: Philippines - SGD: Singapore - THB: Thailand - USD: United States - VND: Vietnam
    Currency *string `json:"currency,omitempty"`
    // **Global sellers** The SKU's **local display price** shown on the product page before any discounts.
    SalePrice *string `json:"sale_price,omitempty"`
    // **Local sellers/Intra-EU sellers** The SKU's **local display price** shown on the product page before any discounts.  **Global sellers** The SKU's **local pre-tax price**. This excludes any applicable charges such as cross-border shipping costs, taxes, and other fees, and therefore does not appear on the product page.  **Note**: Tax-exclusive pricing does not apply to the JP market, therefore this value is the same as `sale_price`.
    TaxExclusivePrice *string `json:"tax_exclusive_price,omitempty"`
}

// NewProduct202312SearchProductsResponseDataProductsSkusPrice instantiates a new Product202312SearchProductsResponseDataProductsSkusPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202312SearchProductsResponseDataProductsSkusPrice() *Product202312SearchProductsResponseDataProductsSkusPrice {
    this := Product202312SearchProductsResponseDataProductsSkusPrice{}
    return &this
}

// NewProduct202312SearchProductsResponseDataProductsSkusPriceWithDefaults instantiates a new Product202312SearchProductsResponseDataProductsSkusPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202312SearchProductsResponseDataProductsSkusPriceWithDefaults() *Product202312SearchProductsResponseDataProductsSkusPrice {
    this := Product202312SearchProductsResponseDataProductsSkusPrice{}
    return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Product202312SearchProductsResponseDataProductsSkusPrice) GetCurrency() string {
    if o == nil || utils.IsNil(o.Currency) {
        var ret string
        return ret
    }
    return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202312SearchProductsResponseDataProductsSkusPrice) GetCurrencyOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Currency) {
        return nil, false
    }
    return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Product202312SearchProductsResponseDataProductsSkusPrice) HasCurrency() bool {
    if o != nil && !utils.IsNil(o.Currency) {
        return true
    }

    return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Product202312SearchProductsResponseDataProductsSkusPrice) SetCurrency(v string) {
    o.Currency = &v
}

// GetSalePrice returns the SalePrice field value if set, zero value otherwise.
func (o *Product202312SearchProductsResponseDataProductsSkusPrice) GetSalePrice() string {
    if o == nil || utils.IsNil(o.SalePrice) {
        var ret string
        return ret
    }
    return *o.SalePrice
}

// GetSalePriceOk returns a tuple with the SalePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202312SearchProductsResponseDataProductsSkusPrice) GetSalePriceOk() (*string, bool) {
    if o == nil || utils.IsNil(o.SalePrice) {
        return nil, false
    }
    return o.SalePrice, true
}

// HasSalePrice returns a boolean if a field has been set.
func (o *Product202312SearchProductsResponseDataProductsSkusPrice) HasSalePrice() bool {
    if o != nil && !utils.IsNil(o.SalePrice) {
        return true
    }

    return false
}

// SetSalePrice gets a reference to the given string and assigns it to the SalePrice field.
func (o *Product202312SearchProductsResponseDataProductsSkusPrice) SetSalePrice(v string) {
    o.SalePrice = &v
}

// GetTaxExclusivePrice returns the TaxExclusivePrice field value if set, zero value otherwise.
func (o *Product202312SearchProductsResponseDataProductsSkusPrice) GetTaxExclusivePrice() string {
    if o == nil || utils.IsNil(o.TaxExclusivePrice) {
        var ret string
        return ret
    }
    return *o.TaxExclusivePrice
}

// GetTaxExclusivePriceOk returns a tuple with the TaxExclusivePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202312SearchProductsResponseDataProductsSkusPrice) GetTaxExclusivePriceOk() (*string, bool) {
    if o == nil || utils.IsNil(o.TaxExclusivePrice) {
        return nil, false
    }
    return o.TaxExclusivePrice, true
}

// HasTaxExclusivePrice returns a boolean if a field has been set.
func (o *Product202312SearchProductsResponseDataProductsSkusPrice) HasTaxExclusivePrice() bool {
    if o != nil && !utils.IsNil(o.TaxExclusivePrice) {
        return true
    }

    return false
}

// SetTaxExclusivePrice gets a reference to the given string and assigns it to the TaxExclusivePrice field.
func (o *Product202312SearchProductsResponseDataProductsSkusPrice) SetTaxExclusivePrice(v string) {
    o.TaxExclusivePrice = &v
}

func (o Product202312SearchProductsResponseDataProductsSkusPrice) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202312SearchProductsResponseDataProductsSkusPrice) ToMap() (map[string]interface{}, error) {
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

type NullableProduct202312SearchProductsResponseDataProductsSkusPrice struct {
	value *Product202312SearchProductsResponseDataProductsSkusPrice
	isSet bool
}

func (v NullableProduct202312SearchProductsResponseDataProductsSkusPrice) Get() *Product202312SearchProductsResponseDataProductsSkusPrice {
	return v.value
}

func (v *NullableProduct202312SearchProductsResponseDataProductsSkusPrice) Set(val *Product202312SearchProductsResponseDataProductsSkusPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202312SearchProductsResponseDataProductsSkusPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202312SearchProductsResponseDataProductsSkusPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202312SearchProductsResponseDataProductsSkusPrice(val *Product202312SearchProductsResponseDataProductsSkusPrice) *NullableProduct202312SearchProductsResponseDataProductsSkusPrice {
	return &NullableProduct202312SearchProductsResponseDataProductsSkusPrice{value: val, isSet: true}
}

func (v NullableProduct202312SearchProductsResponseDataProductsSkusPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202312SearchProductsResponseDataProductsSkusPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


