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

            // checks if the Product202309UpdatePriceRequestBodySkusPrice type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309UpdatePriceRequestBodySkusPrice{}

// Product202309UpdatePriceRequestBodySkusPrice struct for Product202309UpdatePriceRequestBodySkusPrice
type Product202309UpdatePriceRequestBodySkusPrice struct {
    // **Local sellers/Intra-EU sellers** The SKU's **local display price** shown on the product page before any discounts. Refer to [Product Pricing](https://partner.tiktokshop.com/docv2/page/67e1288d76cfee049d9af858) for the allowed price ranges in each market.   **Global sellers** The SKU's **local pre-tax price**. This excludes any applicable charges such as cross-border shipping costs, taxes, and other fees, and therefore does not appear on the product page. Refer to [Product Pricing](https://partner.tiktokshop.com/docv2/page/67e1288d76cfee049d9af858) for the allowed price ranges in each market. - **Note**: Not applicable for JP and US shops using China warehouses, please use `price.sale_price` instead.
    Amount *string `json:"amount,omitempty"`
    // The currency. Possible values based on the region: - BRL: Brazil - EUR: France, Germany, Ireland, Italy, Spain - GBP: United Kingdom - IDR: Indonesia - JPY: Japan - MXN: Mexico - MYR: Malaysia - PHP: Philippines - SGD: Singapore - THB: Thailand - USD: United States - VND: Vietnam
    Currency *string `json:"currency,omitempty"`
    // **Global sellers** The SKU's **local display price** shown on the product page before any discounts.  Refer to [Product Pricing](https://partner.tiktokshop.com/docv2/page/67e1288d76cfee049d9af858) for the allowed price ranges in each market.  **Note**:  - Applicable only for global sellers. -  Required for the JP and US shops using China warehouses, optional for others. - This is the definitive final price shown on the product page, all other prices will be ignored.
    SalePrice *string `json:"sale_price,omitempty"`
}

// NewProduct202309UpdatePriceRequestBodySkusPrice instantiates a new Product202309UpdatePriceRequestBodySkusPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309UpdatePriceRequestBodySkusPrice() *Product202309UpdatePriceRequestBodySkusPrice {
    this := Product202309UpdatePriceRequestBodySkusPrice{}
    return &this
}

// NewProduct202309UpdatePriceRequestBodySkusPriceWithDefaults instantiates a new Product202309UpdatePriceRequestBodySkusPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309UpdatePriceRequestBodySkusPriceWithDefaults() *Product202309UpdatePriceRequestBodySkusPrice {
    this := Product202309UpdatePriceRequestBodySkusPrice{}
    return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Product202309UpdatePriceRequestBodySkusPrice) GetAmount() string {
    if o == nil || utils.IsNil(o.Amount) {
        var ret string
        return ret
    }
    return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309UpdatePriceRequestBodySkusPrice) GetAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Amount) {
        return nil, false
    }
    return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Product202309UpdatePriceRequestBodySkusPrice) HasAmount() bool {
    if o != nil && !utils.IsNil(o.Amount) {
        return true
    }

    return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *Product202309UpdatePriceRequestBodySkusPrice) SetAmount(v string) {
    o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Product202309UpdatePriceRequestBodySkusPrice) GetCurrency() string {
    if o == nil || utils.IsNil(o.Currency) {
        var ret string
        return ret
    }
    return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309UpdatePriceRequestBodySkusPrice) GetCurrencyOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Currency) {
        return nil, false
    }
    return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Product202309UpdatePriceRequestBodySkusPrice) HasCurrency() bool {
    if o != nil && !utils.IsNil(o.Currency) {
        return true
    }

    return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Product202309UpdatePriceRequestBodySkusPrice) SetCurrency(v string) {
    o.Currency = &v
}

// GetSalePrice returns the SalePrice field value if set, zero value otherwise.
func (o *Product202309UpdatePriceRequestBodySkusPrice) GetSalePrice() string {
    if o == nil || utils.IsNil(o.SalePrice) {
        var ret string
        return ret
    }
    return *o.SalePrice
}

// GetSalePriceOk returns a tuple with the SalePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309UpdatePriceRequestBodySkusPrice) GetSalePriceOk() (*string, bool) {
    if o == nil || utils.IsNil(o.SalePrice) {
        return nil, false
    }
    return o.SalePrice, true
}

// HasSalePrice returns a boolean if a field has been set.
func (o *Product202309UpdatePriceRequestBodySkusPrice) HasSalePrice() bool {
    if o != nil && !utils.IsNil(o.SalePrice) {
        return true
    }

    return false
}

// SetSalePrice gets a reference to the given string and assigns it to the SalePrice field.
func (o *Product202309UpdatePriceRequestBodySkusPrice) SetSalePrice(v string) {
    o.SalePrice = &v
}

func (o Product202309UpdatePriceRequestBodySkusPrice) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309UpdatePriceRequestBodySkusPrice) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Amount) {
        toSerialize["amount"] = o.Amount
    }
    if !utils.IsNil(o.Currency) {
        toSerialize["currency"] = o.Currency
    }
    if !utils.IsNil(o.SalePrice) {
        toSerialize["sale_price"] = o.SalePrice
    }
    return toSerialize, nil
}

type NullableProduct202309UpdatePriceRequestBodySkusPrice struct {
	value *Product202309UpdatePriceRequestBodySkusPrice
	isSet bool
}

func (v NullableProduct202309UpdatePriceRequestBodySkusPrice) Get() *Product202309UpdatePriceRequestBodySkusPrice {
	return v.value
}

func (v *NullableProduct202309UpdatePriceRequestBodySkusPrice) Set(val *Product202309UpdatePriceRequestBodySkusPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309UpdatePriceRequestBodySkusPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309UpdatePriceRequestBodySkusPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309UpdatePriceRequestBodySkusPrice(val *Product202309UpdatePriceRequestBodySkusPrice) *NullableProduct202309UpdatePriceRequestBodySkusPrice {
	return &NullableProduct202309UpdatePriceRequestBodySkusPrice{value: val, isSet: true}
}

func (v NullableProduct202309UpdatePriceRequestBodySkusPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309UpdatePriceRequestBodySkusPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


