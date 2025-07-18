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

            // checks if the Promotion202309UpdateActivityProductRequestBodyProductsSkus type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Promotion202309UpdateActivityProductRequestBodyProductsSkus{}

// Promotion202309UpdateActivityProductRequestBodyProductsSkus struct for Promotion202309UpdateActivityProductRequestBodyProductsSkus
type Promotion202309UpdateActivityProductRequestBodyProductsSkus struct {
    // Deal price.  You must specify the value when `product_level==VARIATION` and `activity_type==FIXED_PRICE / FLASHSALE`. The currency of activity price is the same as that of SKU price.
    ActivityPriceAmount *string `json:"activity_price_amount,omitempty"`
    // Discount value. If the SKU is 10% off, the value is `10`. You must specify the value when `product_level==VARIATION` and `activity_type==DIRECT_DISCOUNT`.
    Discount *string `json:"discount,omitempty"`
    // SKU ID
    Id *string `json:"id,omitempty"`
    // The quantity limit of the SKU involved in the activity. The range is `[1, 99]`, or you can use `-1` for unlimited. If you are updating the value of an existing SKU, the value cannot be decreased.
    QuantityLimit *int32 `json:"quantity_limit,omitempty"`
    // Limit of SKU purchase per buyer. The range is `[1, 99]`, or you can use `-1` for unlimited. If you are updating the value of an existing SKU, the value cannot be decreased.
    QuantityPerUser *int32 `json:"quantity_per_user,omitempty"`
}

// NewPromotion202309UpdateActivityProductRequestBodyProductsSkus instantiates a new Promotion202309UpdateActivityProductRequestBodyProductsSkus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotion202309UpdateActivityProductRequestBodyProductsSkus() *Promotion202309UpdateActivityProductRequestBodyProductsSkus {
    this := Promotion202309UpdateActivityProductRequestBodyProductsSkus{}
    return &this
}

// NewPromotion202309UpdateActivityProductRequestBodyProductsSkusWithDefaults instantiates a new Promotion202309UpdateActivityProductRequestBodyProductsSkus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotion202309UpdateActivityProductRequestBodyProductsSkusWithDefaults() *Promotion202309UpdateActivityProductRequestBodyProductsSkus {
    this := Promotion202309UpdateActivityProductRequestBodyProductsSkus{}
    return &this
}

// GetActivityPriceAmount returns the ActivityPriceAmount field value if set, zero value otherwise.
func (o *Promotion202309UpdateActivityProductRequestBodyProductsSkus) GetActivityPriceAmount() string {
    if o == nil || utils.IsNil(o.ActivityPriceAmount) {
        var ret string
        return ret
    }
    return *o.ActivityPriceAmount
}

// GetActivityPriceAmountOk returns a tuple with the ActivityPriceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202309UpdateActivityProductRequestBodyProductsSkus) GetActivityPriceAmountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ActivityPriceAmount) {
        return nil, false
    }
    return o.ActivityPriceAmount, true
}

// HasActivityPriceAmount returns a boolean if a field has been set.
func (o *Promotion202309UpdateActivityProductRequestBodyProductsSkus) HasActivityPriceAmount() bool {
    if o != nil && !utils.IsNil(o.ActivityPriceAmount) {
        return true
    }

    return false
}

// SetActivityPriceAmount gets a reference to the given string and assigns it to the ActivityPriceAmount field.
func (o *Promotion202309UpdateActivityProductRequestBodyProductsSkus) SetActivityPriceAmount(v string) {
    o.ActivityPriceAmount = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *Promotion202309UpdateActivityProductRequestBodyProductsSkus) GetDiscount() string {
    if o == nil || utils.IsNil(o.Discount) {
        var ret string
        return ret
    }
    return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202309UpdateActivityProductRequestBodyProductsSkus) GetDiscountOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Discount) {
        return nil, false
    }
    return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *Promotion202309UpdateActivityProductRequestBodyProductsSkus) HasDiscount() bool {
    if o != nil && !utils.IsNil(o.Discount) {
        return true
    }

    return false
}

// SetDiscount gets a reference to the given string and assigns it to the Discount field.
func (o *Promotion202309UpdateActivityProductRequestBodyProductsSkus) SetDiscount(v string) {
    o.Discount = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Promotion202309UpdateActivityProductRequestBodyProductsSkus) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202309UpdateActivityProductRequestBodyProductsSkus) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Promotion202309UpdateActivityProductRequestBodyProductsSkus) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Promotion202309UpdateActivityProductRequestBodyProductsSkus) SetId(v string) {
    o.Id = &v
}

// GetQuantityLimit returns the QuantityLimit field value if set, zero value otherwise.
func (o *Promotion202309UpdateActivityProductRequestBodyProductsSkus) GetQuantityLimit() int32 {
    if o == nil || utils.IsNil(o.QuantityLimit) {
        var ret int32
        return ret
    }
    return *o.QuantityLimit
}

// GetQuantityLimitOk returns a tuple with the QuantityLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202309UpdateActivityProductRequestBodyProductsSkus) GetQuantityLimitOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.QuantityLimit) {
        return nil, false
    }
    return o.QuantityLimit, true
}

// HasQuantityLimit returns a boolean if a field has been set.
func (o *Promotion202309UpdateActivityProductRequestBodyProductsSkus) HasQuantityLimit() bool {
    if o != nil && !utils.IsNil(o.QuantityLimit) {
        return true
    }

    return false
}

// SetQuantityLimit gets a reference to the given int32 and assigns it to the QuantityLimit field.
func (o *Promotion202309UpdateActivityProductRequestBodyProductsSkus) SetQuantityLimit(v int32) {
    o.QuantityLimit = &v
}

// GetQuantityPerUser returns the QuantityPerUser field value if set, zero value otherwise.
func (o *Promotion202309UpdateActivityProductRequestBodyProductsSkus) GetQuantityPerUser() int32 {
    if o == nil || utils.IsNil(o.QuantityPerUser) {
        var ret int32
        return ret
    }
    return *o.QuantityPerUser
}

// GetQuantityPerUserOk returns a tuple with the QuantityPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202309UpdateActivityProductRequestBodyProductsSkus) GetQuantityPerUserOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.QuantityPerUser) {
        return nil, false
    }
    return o.QuantityPerUser, true
}

// HasQuantityPerUser returns a boolean if a field has been set.
func (o *Promotion202309UpdateActivityProductRequestBodyProductsSkus) HasQuantityPerUser() bool {
    if o != nil && !utils.IsNil(o.QuantityPerUser) {
        return true
    }

    return false
}

// SetQuantityPerUser gets a reference to the given int32 and assigns it to the QuantityPerUser field.
func (o *Promotion202309UpdateActivityProductRequestBodyProductsSkus) SetQuantityPerUser(v int32) {
    o.QuantityPerUser = &v
}

func (o Promotion202309UpdateActivityProductRequestBodyProductsSkus) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Promotion202309UpdateActivityProductRequestBodyProductsSkus) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.ActivityPriceAmount) {
        toSerialize["activity_price_amount"] = o.ActivityPriceAmount
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

type NullablePromotion202309UpdateActivityProductRequestBodyProductsSkus struct {
	value *Promotion202309UpdateActivityProductRequestBodyProductsSkus
	isSet bool
}

func (v NullablePromotion202309UpdateActivityProductRequestBodyProductsSkus) Get() *Promotion202309UpdateActivityProductRequestBodyProductsSkus {
	return v.value
}

func (v *NullablePromotion202309UpdateActivityProductRequestBodyProductsSkus) Set(val *Promotion202309UpdateActivityProductRequestBodyProductsSkus) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotion202309UpdateActivityProductRequestBodyProductsSkus) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotion202309UpdateActivityProductRequestBodyProductsSkus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotion202309UpdateActivityProductRequestBodyProductsSkus(val *Promotion202309UpdateActivityProductRequestBodyProductsSkus) *NullablePromotion202309UpdateActivityProductRequestBodyProductsSkus {
	return &NullablePromotion202309UpdateActivityProductRequestBodyProductsSkus{value: val, isSet: true}
}

func (v NullablePromotion202309UpdateActivityProductRequestBodyProductsSkus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotion202309UpdateActivityProductRequestBodyProductsSkus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


