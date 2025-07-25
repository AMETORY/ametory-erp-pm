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

            // checks if the Product202309DeleteGlobalProductsResponseDataErrors type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309DeleteGlobalProductsResponseDataErrors{}

// Product202309DeleteGlobalProductsResponseDataErrors struct for Product202309DeleteGlobalProductsResponseDataErrors
type Product202309DeleteGlobalProductsResponseDataErrors struct {
    // The error code.
    Code *int32 `json:"code,omitempty"`
    Detail *Product202309DeleteGlobalProductsResponseDataErrorsDetail `json:"detail,omitempty"`
    // The error message.
    Message *string `json:"message,omitempty"`
}

// NewProduct202309DeleteGlobalProductsResponseDataErrors instantiates a new Product202309DeleteGlobalProductsResponseDataErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309DeleteGlobalProductsResponseDataErrors() *Product202309DeleteGlobalProductsResponseDataErrors {
    this := Product202309DeleteGlobalProductsResponseDataErrors{}
    return &this
}

// NewProduct202309DeleteGlobalProductsResponseDataErrorsWithDefaults instantiates a new Product202309DeleteGlobalProductsResponseDataErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309DeleteGlobalProductsResponseDataErrorsWithDefaults() *Product202309DeleteGlobalProductsResponseDataErrors {
    this := Product202309DeleteGlobalProductsResponseDataErrors{}
    return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Product202309DeleteGlobalProductsResponseDataErrors) GetCode() int32 {
    if o == nil || utils.IsNil(o.Code) {
        var ret int32
        return ret
    }
    return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309DeleteGlobalProductsResponseDataErrors) GetCodeOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Code) {
        return nil, false
    }
    return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Product202309DeleteGlobalProductsResponseDataErrors) HasCode() bool {
    if o != nil && !utils.IsNil(o.Code) {
        return true
    }

    return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *Product202309DeleteGlobalProductsResponseDataErrors) SetCode(v int32) {
    o.Code = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *Product202309DeleteGlobalProductsResponseDataErrors) GetDetail() Product202309DeleteGlobalProductsResponseDataErrorsDetail {
    if o == nil || utils.IsNil(o.Detail) {
        var ret Product202309DeleteGlobalProductsResponseDataErrorsDetail
        return ret
    }
    return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309DeleteGlobalProductsResponseDataErrors) GetDetailOk() (*Product202309DeleteGlobalProductsResponseDataErrorsDetail, bool) {
    if o == nil || utils.IsNil(o.Detail) {
        return nil, false
    }
    return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *Product202309DeleteGlobalProductsResponseDataErrors) HasDetail() bool {
    if o != nil && !utils.IsNil(o.Detail) {
        return true
    }

    return false
}

// SetDetail gets a reference to the given Product202309DeleteGlobalProductsResponseDataErrorsDetail and assigns it to the Detail field.
func (o *Product202309DeleteGlobalProductsResponseDataErrors) SetDetail(v Product202309DeleteGlobalProductsResponseDataErrorsDetail) {
    o.Detail = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Product202309DeleteGlobalProductsResponseDataErrors) GetMessage() string {
    if o == nil || utils.IsNil(o.Message) {
        var ret string
        return ret
    }
    return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309DeleteGlobalProductsResponseDataErrors) GetMessageOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Message) {
        return nil, false
    }
    return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Product202309DeleteGlobalProductsResponseDataErrors) HasMessage() bool {
    if o != nil && !utils.IsNil(o.Message) {
        return true
    }

    return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Product202309DeleteGlobalProductsResponseDataErrors) SetMessage(v string) {
    o.Message = &v
}

func (o Product202309DeleteGlobalProductsResponseDataErrors) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309DeleteGlobalProductsResponseDataErrors) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Code) {
        toSerialize["code"] = o.Code
    }
    if !utils.IsNil(o.Detail) {
        toSerialize["detail"] = o.Detail
    }
    if !utils.IsNil(o.Message) {
        toSerialize["message"] = o.Message
    }
    return toSerialize, nil
}

type NullableProduct202309DeleteGlobalProductsResponseDataErrors struct {
	value *Product202309DeleteGlobalProductsResponseDataErrors
	isSet bool
}

func (v NullableProduct202309DeleteGlobalProductsResponseDataErrors) Get() *Product202309DeleteGlobalProductsResponseDataErrors {
	return v.value
}

func (v *NullableProduct202309DeleteGlobalProductsResponseDataErrors) Set(val *Product202309DeleteGlobalProductsResponseDataErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309DeleteGlobalProductsResponseDataErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309DeleteGlobalProductsResponseDataErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309DeleteGlobalProductsResponseDataErrors(val *Product202309DeleteGlobalProductsResponseDataErrors) *NullableProduct202309DeleteGlobalProductsResponseDataErrors {
	return &NullableProduct202309DeleteGlobalProductsResponseDataErrors{value: val, isSet: true}
}

func (v NullableProduct202309DeleteGlobalProductsResponseDataErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309DeleteGlobalProductsResponseDataErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


