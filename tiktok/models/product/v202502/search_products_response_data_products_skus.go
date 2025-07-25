/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product_v202502

import (
    "encoding/json"
    "tiktokshop/open/sdk_golang/utils"
)

            // checks if the Product202502SearchProductsResponseDataProductsSkus type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202502SearchProductsResponseDataProductsSkus{}

// Product202502SearchProductsResponseDataProductsSkus struct for Product202502SearchProductsResponseDataProductsSkus
type Product202502SearchProductsResponseDataProductsSkus struct {
    // The SKU list price (e.g. MSRP, RRP) or original price information on external ecommerce platforms. Applicable only for selected sellers in the US market.  **Note**: This value may appear as the strikethrough price on the product page. However, whether the strikethrough price is shown and the amount shown are subject to the audit team's review and decision based on various pricing information.
    ExternalListPrices []Product202502SearchProductsResponseDataProductsSkusExternalListPrices `json:"external_list_prices,omitempty"`
    // The SKU ID generated by TikTok Shop.
    Id *string `json:"id,omitempty"`
    // SKU inventory information.
    Inventory []Product202502SearchProductsResponseDataProductsSkusInventory `json:"inventory,omitempty"`
    ListPrice *Product202502SearchProductsResponseDataProductsSkusListPrice `json:"list_price,omitempty"`
    PreSale *Product202502SearchProductsResponseDataProductsSkusPreSale `json:"pre_sale,omitempty"`
    Price *Product202502SearchProductsResponseDataProductsSkusPrice `json:"price,omitempty"`
    // An internal code/name for managing SKUs, not visible to buyers. 
    SellerSku *string `json:"seller_sku,omitempty"`
}

// NewProduct202502SearchProductsResponseDataProductsSkus instantiates a new Product202502SearchProductsResponseDataProductsSkus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202502SearchProductsResponseDataProductsSkus() *Product202502SearchProductsResponseDataProductsSkus {
    this := Product202502SearchProductsResponseDataProductsSkus{}
    return &this
}

// NewProduct202502SearchProductsResponseDataProductsSkusWithDefaults instantiates a new Product202502SearchProductsResponseDataProductsSkus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202502SearchProductsResponseDataProductsSkusWithDefaults() *Product202502SearchProductsResponseDataProductsSkus {
    this := Product202502SearchProductsResponseDataProductsSkus{}
    return &this
}

// GetExternalListPrices returns the ExternalListPrices field value if set, zero value otherwise.
func (o *Product202502SearchProductsResponseDataProductsSkus) GetExternalListPrices() []Product202502SearchProductsResponseDataProductsSkusExternalListPrices {
    if o == nil || utils.IsNil(o.ExternalListPrices) {
        var ret []Product202502SearchProductsResponseDataProductsSkusExternalListPrices
        return ret
    }
    return o.ExternalListPrices
}

// GetExternalListPricesOk returns a tuple with the ExternalListPrices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202502SearchProductsResponseDataProductsSkus) GetExternalListPricesOk() ([]Product202502SearchProductsResponseDataProductsSkusExternalListPrices, bool) {
    if o == nil || utils.IsNil(o.ExternalListPrices) {
        return nil, false
    }
    return o.ExternalListPrices, true
}

// HasExternalListPrices returns a boolean if a field has been set.
func (o *Product202502SearchProductsResponseDataProductsSkus) HasExternalListPrices() bool {
    if o != nil && !utils.IsNil(o.ExternalListPrices) {
        return true
    }

    return false
}

// SetExternalListPrices gets a reference to the given []Product202502SearchProductsResponseDataProductsSkusExternalListPrices and assigns it to the ExternalListPrices field.
func (o *Product202502SearchProductsResponseDataProductsSkus) SetExternalListPrices(v []Product202502SearchProductsResponseDataProductsSkusExternalListPrices) {
    o.ExternalListPrices = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Product202502SearchProductsResponseDataProductsSkus) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202502SearchProductsResponseDataProductsSkus) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Product202502SearchProductsResponseDataProductsSkus) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Product202502SearchProductsResponseDataProductsSkus) SetId(v string) {
    o.Id = &v
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *Product202502SearchProductsResponseDataProductsSkus) GetInventory() []Product202502SearchProductsResponseDataProductsSkusInventory {
    if o == nil || utils.IsNil(o.Inventory) {
        var ret []Product202502SearchProductsResponseDataProductsSkusInventory
        return ret
    }
    return o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202502SearchProductsResponseDataProductsSkus) GetInventoryOk() ([]Product202502SearchProductsResponseDataProductsSkusInventory, bool) {
    if o == nil || utils.IsNil(o.Inventory) {
        return nil, false
    }
    return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *Product202502SearchProductsResponseDataProductsSkus) HasInventory() bool {
    if o != nil && !utils.IsNil(o.Inventory) {
        return true
    }

    return false
}

// SetInventory gets a reference to the given []Product202502SearchProductsResponseDataProductsSkusInventory and assigns it to the Inventory field.
func (o *Product202502SearchProductsResponseDataProductsSkus) SetInventory(v []Product202502SearchProductsResponseDataProductsSkusInventory) {
    o.Inventory = v
}

// GetListPrice returns the ListPrice field value if set, zero value otherwise.
func (o *Product202502SearchProductsResponseDataProductsSkus) GetListPrice() Product202502SearchProductsResponseDataProductsSkusListPrice {
    if o == nil || utils.IsNil(o.ListPrice) {
        var ret Product202502SearchProductsResponseDataProductsSkusListPrice
        return ret
    }
    return *o.ListPrice
}

// GetListPriceOk returns a tuple with the ListPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202502SearchProductsResponseDataProductsSkus) GetListPriceOk() (*Product202502SearchProductsResponseDataProductsSkusListPrice, bool) {
    if o == nil || utils.IsNil(o.ListPrice) {
        return nil, false
    }
    return o.ListPrice, true
}

// HasListPrice returns a boolean if a field has been set.
func (o *Product202502SearchProductsResponseDataProductsSkus) HasListPrice() bool {
    if o != nil && !utils.IsNil(o.ListPrice) {
        return true
    }

    return false
}

// SetListPrice gets a reference to the given Product202502SearchProductsResponseDataProductsSkusListPrice and assigns it to the ListPrice field.
func (o *Product202502SearchProductsResponseDataProductsSkus) SetListPrice(v Product202502SearchProductsResponseDataProductsSkusListPrice) {
    o.ListPrice = &v
}

// GetPreSale returns the PreSale field value if set, zero value otherwise.
func (o *Product202502SearchProductsResponseDataProductsSkus) GetPreSale() Product202502SearchProductsResponseDataProductsSkusPreSale {
    if o == nil || utils.IsNil(o.PreSale) {
        var ret Product202502SearchProductsResponseDataProductsSkusPreSale
        return ret
    }
    return *o.PreSale
}

// GetPreSaleOk returns a tuple with the PreSale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202502SearchProductsResponseDataProductsSkus) GetPreSaleOk() (*Product202502SearchProductsResponseDataProductsSkusPreSale, bool) {
    if o == nil || utils.IsNil(o.PreSale) {
        return nil, false
    }
    return o.PreSale, true
}

// HasPreSale returns a boolean if a field has been set.
func (o *Product202502SearchProductsResponseDataProductsSkus) HasPreSale() bool {
    if o != nil && !utils.IsNil(o.PreSale) {
        return true
    }

    return false
}

// SetPreSale gets a reference to the given Product202502SearchProductsResponseDataProductsSkusPreSale and assigns it to the PreSale field.
func (o *Product202502SearchProductsResponseDataProductsSkus) SetPreSale(v Product202502SearchProductsResponseDataProductsSkusPreSale) {
    o.PreSale = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *Product202502SearchProductsResponseDataProductsSkus) GetPrice() Product202502SearchProductsResponseDataProductsSkusPrice {
    if o == nil || utils.IsNil(o.Price) {
        var ret Product202502SearchProductsResponseDataProductsSkusPrice
        return ret
    }
    return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202502SearchProductsResponseDataProductsSkus) GetPriceOk() (*Product202502SearchProductsResponseDataProductsSkusPrice, bool) {
    if o == nil || utils.IsNil(o.Price) {
        return nil, false
    }
    return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *Product202502SearchProductsResponseDataProductsSkus) HasPrice() bool {
    if o != nil && !utils.IsNil(o.Price) {
        return true
    }

    return false
}

// SetPrice gets a reference to the given Product202502SearchProductsResponseDataProductsSkusPrice and assigns it to the Price field.
func (o *Product202502SearchProductsResponseDataProductsSkus) SetPrice(v Product202502SearchProductsResponseDataProductsSkusPrice) {
    o.Price = &v
}

// GetSellerSku returns the SellerSku field value if set, zero value otherwise.
func (o *Product202502SearchProductsResponseDataProductsSkus) GetSellerSku() string {
    if o == nil || utils.IsNil(o.SellerSku) {
        var ret string
        return ret
    }
    return *o.SellerSku
}

// GetSellerSkuOk returns a tuple with the SellerSku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202502SearchProductsResponseDataProductsSkus) GetSellerSkuOk() (*string, bool) {
    if o == nil || utils.IsNil(o.SellerSku) {
        return nil, false
    }
    return o.SellerSku, true
}

// HasSellerSku returns a boolean if a field has been set.
func (o *Product202502SearchProductsResponseDataProductsSkus) HasSellerSku() bool {
    if o != nil && !utils.IsNil(o.SellerSku) {
        return true
    }

    return false
}

// SetSellerSku gets a reference to the given string and assigns it to the SellerSku field.
func (o *Product202502SearchProductsResponseDataProductsSkus) SetSellerSku(v string) {
    o.SellerSku = &v
}

func (o Product202502SearchProductsResponseDataProductsSkus) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202502SearchProductsResponseDataProductsSkus) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.ExternalListPrices) {
        toSerialize["external_list_prices"] = o.ExternalListPrices
    }
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    if !utils.IsNil(o.Inventory) {
        toSerialize["inventory"] = o.Inventory
    }
    if !utils.IsNil(o.ListPrice) {
        toSerialize["list_price"] = o.ListPrice
    }
    if !utils.IsNil(o.PreSale) {
        toSerialize["pre_sale"] = o.PreSale
    }
    if !utils.IsNil(o.Price) {
        toSerialize["price"] = o.Price
    }
    if !utils.IsNil(o.SellerSku) {
        toSerialize["seller_sku"] = o.SellerSku
    }
    return toSerialize, nil
}

type NullableProduct202502SearchProductsResponseDataProductsSkus struct {
	value *Product202502SearchProductsResponseDataProductsSkus
	isSet bool
}

func (v NullableProduct202502SearchProductsResponseDataProductsSkus) Get() *Product202502SearchProductsResponseDataProductsSkus {
	return v.value
}

func (v *NullableProduct202502SearchProductsResponseDataProductsSkus) Set(val *Product202502SearchProductsResponseDataProductsSkus) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202502SearchProductsResponseDataProductsSkus) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202502SearchProductsResponseDataProductsSkus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202502SearchProductsResponseDataProductsSkus(val *Product202502SearchProductsResponseDataProductsSkus) *NullableProduct202502SearchProductsResponseDataProductsSkus {
	return &NullableProduct202502SearchProductsResponseDataProductsSkus{value: val, isSet: true}
}

func (v NullableProduct202502SearchProductsResponseDataProductsSkus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202502SearchProductsResponseDataProductsSkus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


