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

            // checks if the Product202309SearchGlobalProductsRequestBody type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309SearchGlobalProductsRequestBody{}

// Product202309SearchGlobalProductsRequestBody struct for Product202309SearchGlobalProductsRequestBody
type Product202309SearchGlobalProductsRequestBody struct {
    // The fields \"create_time_ge\" and \"create_time_le\" together constitute the filter condition for the creation time of the global product.  - If you only fill in the \"create_time_le\", and the \"create_time_ge\" is empty , then we will set the earliest time of the shop to the field \"create_time_ge\" by default.  - If you only fill in the \"create_time_ge\", and the \"create_time_le\" is empty , then we will set the current time to the field \"create_time_le\" by default.  The time search condition uses Unix timestamp in GMT (UTC+00:00). 
    CreateTimeGe *int64 `json:"create_time_ge,omitempty"`
    // Refer to the description of \"create_time_ge\".
    CreateTimeLe *int64 `json:"create_time_le,omitempty"`
    // Seller SKUs, a filtering condition used for global product search. This field allows you to search for all global products that contain these Seller SKUs.
    SellerSkus []string `json:"seller_skus,omitempty"`
    // Global Product status, used as a filtering criterion for global product search. including  PUBLISHED,UNPUBLISHED,DRAFT,DELETED
    Status *string `json:"status,omitempty"`
    // The fields \"update_time_ge\" and \"update_time_le\" together constitute the filter condition for the update time of the global product.  -  If you only fill in the \"update_time_le\", and the \"update_time_ge\" is empty , then we will set the earliest time of the shop to the field \"update_time_ge\" by default.  - If you only fill in the \"update_time_ge\", and the \"update_time_le\" is empty , then we will set the current time to the field \"update_time_le\" by default.
    UpdateTimeGe *int64 `json:"update_time_ge,omitempty"`
    // Refer to the description of \"update_time_ge\".
    UpdateTimeLe *int64 `json:"update_time_le,omitempty"`
}

// NewProduct202309SearchGlobalProductsRequestBody instantiates a new Product202309SearchGlobalProductsRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309SearchGlobalProductsRequestBody() *Product202309SearchGlobalProductsRequestBody {
    this := Product202309SearchGlobalProductsRequestBody{}
    return &this
}

// NewProduct202309SearchGlobalProductsRequestBodyWithDefaults instantiates a new Product202309SearchGlobalProductsRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309SearchGlobalProductsRequestBodyWithDefaults() *Product202309SearchGlobalProductsRequestBody {
    this := Product202309SearchGlobalProductsRequestBody{}
    return &this
}

// GetCreateTimeGe returns the CreateTimeGe field value if set, zero value otherwise.
func (o *Product202309SearchGlobalProductsRequestBody) GetCreateTimeGe() int64 {
    if o == nil || utils.IsNil(o.CreateTimeGe) {
        var ret int64
        return ret
    }
    return *o.CreateTimeGe
}

// GetCreateTimeGeOk returns a tuple with the CreateTimeGe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309SearchGlobalProductsRequestBody) GetCreateTimeGeOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.CreateTimeGe) {
        return nil, false
    }
    return o.CreateTimeGe, true
}

// HasCreateTimeGe returns a boolean if a field has been set.
func (o *Product202309SearchGlobalProductsRequestBody) HasCreateTimeGe() bool {
    if o != nil && !utils.IsNil(o.CreateTimeGe) {
        return true
    }

    return false
}

// SetCreateTimeGe gets a reference to the given int64 and assigns it to the CreateTimeGe field.
func (o *Product202309SearchGlobalProductsRequestBody) SetCreateTimeGe(v int64) {
    o.CreateTimeGe = &v
}

// GetCreateTimeLe returns the CreateTimeLe field value if set, zero value otherwise.
func (o *Product202309SearchGlobalProductsRequestBody) GetCreateTimeLe() int64 {
    if o == nil || utils.IsNil(o.CreateTimeLe) {
        var ret int64
        return ret
    }
    return *o.CreateTimeLe
}

// GetCreateTimeLeOk returns a tuple with the CreateTimeLe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309SearchGlobalProductsRequestBody) GetCreateTimeLeOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.CreateTimeLe) {
        return nil, false
    }
    return o.CreateTimeLe, true
}

// HasCreateTimeLe returns a boolean if a field has been set.
func (o *Product202309SearchGlobalProductsRequestBody) HasCreateTimeLe() bool {
    if o != nil && !utils.IsNil(o.CreateTimeLe) {
        return true
    }

    return false
}

// SetCreateTimeLe gets a reference to the given int64 and assigns it to the CreateTimeLe field.
func (o *Product202309SearchGlobalProductsRequestBody) SetCreateTimeLe(v int64) {
    o.CreateTimeLe = &v
}

// GetSellerSkus returns the SellerSkus field value if set, zero value otherwise.
func (o *Product202309SearchGlobalProductsRequestBody) GetSellerSkus() []string {
    if o == nil || utils.IsNil(o.SellerSkus) {
        var ret []string
        return ret
    }
    return o.SellerSkus
}

// GetSellerSkusOk returns a tuple with the SellerSkus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309SearchGlobalProductsRequestBody) GetSellerSkusOk() ([]string, bool) {
    if o == nil || utils.IsNil(o.SellerSkus) {
        return nil, false
    }
    return o.SellerSkus, true
}

// HasSellerSkus returns a boolean if a field has been set.
func (o *Product202309SearchGlobalProductsRequestBody) HasSellerSkus() bool {
    if o != nil && !utils.IsNil(o.SellerSkus) {
        return true
    }

    return false
}

// SetSellerSkus gets a reference to the given []string and assigns it to the SellerSkus field.
func (o *Product202309SearchGlobalProductsRequestBody) SetSellerSkus(v []string) {
    o.SellerSkus = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Product202309SearchGlobalProductsRequestBody) GetStatus() string {
    if o == nil || utils.IsNil(o.Status) {
        var ret string
        return ret
    }
    return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309SearchGlobalProductsRequestBody) GetStatusOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Status) {
        return nil, false
    }
    return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Product202309SearchGlobalProductsRequestBody) HasStatus() bool {
    if o != nil && !utils.IsNil(o.Status) {
        return true
    }

    return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Product202309SearchGlobalProductsRequestBody) SetStatus(v string) {
    o.Status = &v
}

// GetUpdateTimeGe returns the UpdateTimeGe field value if set, zero value otherwise.
func (o *Product202309SearchGlobalProductsRequestBody) GetUpdateTimeGe() int64 {
    if o == nil || utils.IsNil(o.UpdateTimeGe) {
        var ret int64
        return ret
    }
    return *o.UpdateTimeGe
}

// GetUpdateTimeGeOk returns a tuple with the UpdateTimeGe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309SearchGlobalProductsRequestBody) GetUpdateTimeGeOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.UpdateTimeGe) {
        return nil, false
    }
    return o.UpdateTimeGe, true
}

// HasUpdateTimeGe returns a boolean if a field has been set.
func (o *Product202309SearchGlobalProductsRequestBody) HasUpdateTimeGe() bool {
    if o != nil && !utils.IsNil(o.UpdateTimeGe) {
        return true
    }

    return false
}

// SetUpdateTimeGe gets a reference to the given int64 and assigns it to the UpdateTimeGe field.
func (o *Product202309SearchGlobalProductsRequestBody) SetUpdateTimeGe(v int64) {
    o.UpdateTimeGe = &v
}

// GetUpdateTimeLe returns the UpdateTimeLe field value if set, zero value otherwise.
func (o *Product202309SearchGlobalProductsRequestBody) GetUpdateTimeLe() int64 {
    if o == nil || utils.IsNil(o.UpdateTimeLe) {
        var ret int64
        return ret
    }
    return *o.UpdateTimeLe
}

// GetUpdateTimeLeOk returns a tuple with the UpdateTimeLe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309SearchGlobalProductsRequestBody) GetUpdateTimeLeOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.UpdateTimeLe) {
        return nil, false
    }
    return o.UpdateTimeLe, true
}

// HasUpdateTimeLe returns a boolean if a field has been set.
func (o *Product202309SearchGlobalProductsRequestBody) HasUpdateTimeLe() bool {
    if o != nil && !utils.IsNil(o.UpdateTimeLe) {
        return true
    }

    return false
}

// SetUpdateTimeLe gets a reference to the given int64 and assigns it to the UpdateTimeLe field.
func (o *Product202309SearchGlobalProductsRequestBody) SetUpdateTimeLe(v int64) {
    o.UpdateTimeLe = &v
}

func (o Product202309SearchGlobalProductsRequestBody) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309SearchGlobalProductsRequestBody) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.CreateTimeGe) {
        toSerialize["create_time_ge"] = o.CreateTimeGe
    }
    if !utils.IsNil(o.CreateTimeLe) {
        toSerialize["create_time_le"] = o.CreateTimeLe
    }
    if !utils.IsNil(o.SellerSkus) {
        toSerialize["seller_skus"] = o.SellerSkus
    }
    if !utils.IsNil(o.Status) {
        toSerialize["status"] = o.Status
    }
    if !utils.IsNil(o.UpdateTimeGe) {
        toSerialize["update_time_ge"] = o.UpdateTimeGe
    }
    if !utils.IsNil(o.UpdateTimeLe) {
        toSerialize["update_time_le"] = o.UpdateTimeLe
    }
    return toSerialize, nil
}

type NullableProduct202309SearchGlobalProductsRequestBody struct {
	value *Product202309SearchGlobalProductsRequestBody
	isSet bool
}

func (v NullableProduct202309SearchGlobalProductsRequestBody) Get() *Product202309SearchGlobalProductsRequestBody {
	return v.value
}

func (v *NullableProduct202309SearchGlobalProductsRequestBody) Set(val *Product202309SearchGlobalProductsRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309SearchGlobalProductsRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309SearchGlobalProductsRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309SearchGlobalProductsRequestBody(val *Product202309SearchGlobalProductsRequestBody) *NullableProduct202309SearchGlobalProductsRequestBody {
	return &NullableProduct202309SearchGlobalProductsRequestBody{value: val, isSet: true}
}

func (v NullableProduct202309SearchGlobalProductsRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309SearchGlobalProductsRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


