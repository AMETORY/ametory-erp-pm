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

            // checks if the Promotion202309GetActivityResponseDataProductsSkus type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Promotion202309GetActivityResponseDataProductsSkus{}

// Promotion202309GetActivityResponseDataProductsSkus struct for Promotion202309GetActivityResponseDataProductsSkus
type Promotion202309GetActivityResponseDataProductsSkus struct {
    ActivityPrice *Promotion202309GetActivityResponseDataProductsSkusActivityPrice `json:"activity_price,omitempty"`
    // Discount value. If the SKU is 10% off, the value is `10`. Available only if `activity_type==DIRECT_DISCOUNT`.
    Discount *string `json:"discount,omitempty"`
    // TikTok Shop SKU ID.
    Id *string `json:"id,omitempty"`
    // The quantity limit of the SKU involved in the activity. The range is `[1, 99]`, or `-1` for unlimited.
    QuantityLimit *int32 `json:"quantity_limit,omitempty"`
    // Limit of SKU purchase per buyer. The range is `[1, 99]`, or `-1` for unlimited.
    QuantityPerUser *int32 `json:"quantity_per_user,omitempty"`
}

// NewPromotion202309GetActivityResponseDataProductsSkus instantiates a new Promotion202309GetActivityResponseDataProductsSkus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotion202309GetActivityResponseDataProductsSkus() *Promotion202309GetActivityResponseDataProductsSkus {
    this := Promotion202309GetActivityResponseDataProductsSkus{}
    return &this
}

// NewPromotion202309GetActivityResponseDataProductsSkusWithDefaults instantiates a new Promotion202309GetActivityResponseDataProductsSkus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotion202309GetActivityResponseDataProductsSkusWithDefaults() *Promotion202309GetActivityResponseDataProductsSkus {
    this := Promotion202309GetActivityResponseDataProductsSkus{}
    return &this
}

// GetActivityPrice returns the ActivityPrice field value if set, zero value otherwise.
func (o *Promotion202309GetActivityResponseDataProductsSkus) GetActivityPrice() Promotion202309GetActivityResponseDataProductsSkusActivityPrice {
    if o == nil || utils.IsNil(o.ActivityPrice) {
        var ret Promotion202309GetActivityResponseDataProductsSkusActivityPrice
        return ret
    }
    return *o.ActivityPrice
}

// GetActivityPriceOk returns a tuple with the ActivityPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202309GetActivityResponseDataProductsSkus) GetActivityPriceOk() (*Promotion202309GetActivityResponseDataProductsSkusActivityPrice, bool) {
    if o == nil || utils.IsNil(o.ActivityPrice) {
        return nil, false
    }
    return o.ActivityPrice, true
}

// HasActivityPrice returns a boolean if a field has been set.
func (o *Promotion202309GetActivityResponseDataProductsSkus) HasActivityPrice() bool {
    if o != nil && !utils.IsNil(o.ActivityPrice) {
        return true
    }

    return false
}

// SetActivityPrice gets a reference to the given Promotion202309GetActivityResponseDataProductsSkusActivityPrice and assigns it to the ActivityPrice field.
func (o *Promotion202309GetActivityResponseDataProductsSkus) SetActivityPrice(v Promotion202309GetActivityResponseDataProductsSkusActivityPrice) {
    o.ActivityPrice = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *Promotion202309GetActivityResponseDataProductsSkus) GetDiscount() string {
    if o == nil || utils.IsNil(o.Discount) {
        var ret string
        return ret
    }
    return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202309GetActivityResponseDataProductsSkus) GetDiscountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Discount) {
        return nil, false
    }
    return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *Promotion202309GetActivityResponseDataProductsSkus) HasDiscount() bool {
    if o != nil && !utils.IsNil(o.Discount) {
        return true
    }

    return false
}

// SetDiscount gets a reference to the given string and assigns it to the Discount field.
func (o *Promotion202309GetActivityResponseDataProductsSkus) SetDiscount(v string) {
    o.Discount = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Promotion202309GetActivityResponseDataProductsSkus) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202309GetActivityResponseDataProductsSkus) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Promotion202309GetActivityResponseDataProductsSkus) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Promotion202309GetActivityResponseDataProductsSkus) SetId(v string) {
    o.Id = &v
}

// GetQuantityLimit returns the QuantityLimit field value if set, zero value otherwise.
func (o *Promotion202309GetActivityResponseDataProductsSkus) GetQuantityLimit() int32 {
    if o == nil || utils.IsNil(o.QuantityLimit) {
        var ret int32
        return ret
    }
    return *o.QuantityLimit
}

// GetQuantityLimitOk returns a tuple with the QuantityLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202309GetActivityResponseDataProductsSkus) GetQuantityLimitOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.QuantityLimit) {
        return nil, false
    }
    return o.QuantityLimit, true
}

// HasQuantityLimit returns a boolean if a field has been set.
func (o *Promotion202309GetActivityResponseDataProductsSkus) HasQuantityLimit() bool {
    if o != nil && !utils.IsNil(o.QuantityLimit) {
        return true
    }

    return false
}

// SetQuantityLimit gets a reference to the given int32 and assigns it to the QuantityLimit field.
func (o *Promotion202309GetActivityResponseDataProductsSkus) SetQuantityLimit(v int32) {
    o.QuantityLimit = &v
}

// GetQuantityPerUser returns the QuantityPerUser field value if set, zero value otherwise.
func (o *Promotion202309GetActivityResponseDataProductsSkus) GetQuantityPerUser() int32 {
    if o == nil || utils.IsNil(o.QuantityPerUser) {
        var ret int32
        return ret
    }
    return *o.QuantityPerUser
}

// GetQuantityPerUserOk returns a tuple with the QuantityPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202309GetActivityResponseDataProductsSkus) GetQuantityPerUserOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.QuantityPerUser) {
        return nil, false
    }
    return o.QuantityPerUser, true
}

// HasQuantityPerUser returns a boolean if a field has been set.
func (o *Promotion202309GetActivityResponseDataProductsSkus) HasQuantityPerUser() bool {
    if o != nil && !utils.IsNil(o.QuantityPerUser) {
        return true
    }

    return false
}

// SetQuantityPerUser gets a reference to the given int32 and assigns it to the QuantityPerUser field.
func (o *Promotion202309GetActivityResponseDataProductsSkus) SetQuantityPerUser(v int32) {
    o.QuantityPerUser = &v
}

func (o Promotion202309GetActivityResponseDataProductsSkus) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Promotion202309GetActivityResponseDataProductsSkus) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.ActivityPrice) {
        toSerialize["activity_price"] = o.ActivityPrice
    }
    if !utils.IsNil(o.Discount) {
        toSerialize["discount"] = o.Discount
    }
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    if !utils.IsNil(o.QuantityLimit) {
        toSerialize["quantity_limit"] = o.QuantityLimit
    }
    if !utils.IsNil(o.QuantityPerUser) {
        toSerialize["quantity_per_user"] = o.QuantityPerUser
    }
    return toSerialize, nil
}

type NullablePromotion202309GetActivityResponseDataProductsSkus struct {
	value *Promotion202309GetActivityResponseDataProductsSkus
	isSet bool
}

func (v NullablePromotion202309GetActivityResponseDataProductsSkus) Get() *Promotion202309GetActivityResponseDataProductsSkus {
	return v.value
}

func (v *NullablePromotion202309GetActivityResponseDataProductsSkus) Set(val *Promotion202309GetActivityResponseDataProductsSkus) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotion202309GetActivityResponseDataProductsSkus) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotion202309GetActivityResponseDataProductsSkus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotion202309GetActivityResponseDataProductsSkus(val *Promotion202309GetActivityResponseDataProductsSkus) *NullablePromotion202309GetActivityResponseDataProductsSkus {
	return &NullablePromotion202309GetActivityResponseDataProductsSkus{value: val, isSet: true}
}

func (v NullablePromotion202309GetActivityResponseDataProductsSkus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotion202309GetActivityResponseDataProductsSkus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


