/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product_v202407

import (
    "encoding/json"
    "tiktokshop/open/sdk_golang/utils"
)

            // checks if the Product202407SearchSizeChartsResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202407SearchSizeChartsResponse{}

// Product202407SearchSizeChartsResponse struct for Product202407SearchSizeChartsResponse
type Product202407SearchSizeChartsResponse struct {
    // The success or failure status code returned in API response.
    Code *int32 `json:"code,omitempty"`
    Data *Product202407SearchSizeChartsResponseData `json:"data,omitempty"`
    // The success or failure messages returned in API response. Reasons of failure will be described in the message.
    Message *string `json:"message,omitempty"`
    // Request log.
    RequestId *string `json:"request_id,omitempty"`
}

// NewProduct202407SearchSizeChartsResponse instantiates a new Product202407SearchSizeChartsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202407SearchSizeChartsResponse() *Product202407SearchSizeChartsResponse {
    this := Product202407SearchSizeChartsResponse{}
    return &this
}

// NewProduct202407SearchSizeChartsResponseWithDefaults instantiates a new Product202407SearchSizeChartsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202407SearchSizeChartsResponseWithDefaults() *Product202407SearchSizeChartsResponse {
    this := Product202407SearchSizeChartsResponse{}
    return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Product202407SearchSizeChartsResponse) GetCode() int32 {
    if o == nil || utils.IsNil(o.Code) {
        var ret int32
        return ret
    }
    return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202407SearchSizeChartsResponse) GetCodeOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Code) {
        return nil, false
    }
    return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Product202407SearchSizeChartsResponse) HasCode() bool {
    if o != nil && !utils.IsNil(o.Code) {
        return true
    }

    return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *Product202407SearchSizeChartsResponse) SetCode(v int32) {
    o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Product202407SearchSizeChartsResponse) GetData() Product202407SearchSizeChartsResponseData {
    if o == nil || utils.IsNil(o.Data) {
        var ret Product202407SearchSizeChartsResponseData
        return ret
    }
    return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202407SearchSizeChartsResponse) GetDataOk() (*Product202407SearchSizeChartsResponseData, bool) {
    if o == nil || utils.IsNil(o.Data) {
        return nil, false
    }
    return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Product202407SearchSizeChartsResponse) HasData() bool {
    if o != nil && !utils.IsNil(o.Data) {
        return true
    }

    return false
}

// SetData gets a reference to the given Product202407SearchSizeChartsResponseData and assigns it to the Data field.
func (o *Product202407SearchSizeChartsResponse) SetData(v Product202407SearchSizeChartsResponseData) {
    o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Product202407SearchSizeChartsResponse) GetMessage() string {
    if o == nil || utils.IsNil(o.Message) {
        var ret string
        return ret
    }
    return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202407SearchSizeChartsResponse) GetMessageOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Message) {
        return nil, false
    }
    return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Product202407SearchSizeChartsResponse) HasMessage() bool {
    if o != nil && !utils.IsNil(o.Message) {
        return true
    }

    return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Product202407SearchSizeChartsResponse) SetMessage(v string) {
    o.Message = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *Product202407SearchSizeChartsResponse) GetRequestId() string {
    if o == nil || utils.IsNil(o.RequestId) {
        var ret string
        return ret
    }
    return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202407SearchSizeChartsResponse) GetRequestIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.RequestId) {
        return nil, false
    }
    return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *Product202407SearchSizeChartsResponse) HasRequestId() bool {
    if o != nil && !utils.IsNil(o.RequestId) {
        return true
    }

    return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *Product202407SearchSizeChartsResponse) SetRequestId(v string) {
    o.RequestId = &v
}

func (o Product202407SearchSizeChartsResponse) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202407SearchSizeChartsResponse) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Code) {
        toSerialize["code"] = o.Code
    }
    if !utils.IsNil(o.Data) {
        toSerialize["data"] = o.Data
    }
    if !utils.IsNil(o.Message) {
        toSerialize["message"] = o.Message
    }
    if !utils.IsNil(o.RequestId) {
        toSerialize["request_id"] = o.RequestId
    }
    return toSerialize, nil
}

type NullableProduct202407SearchSizeChartsResponse struct {
	value *Product202407SearchSizeChartsResponse
	isSet bool
}

func (v NullableProduct202407SearchSizeChartsResponse) Get() *Product202407SearchSizeChartsResponse {
	return v.value
}

func (v *NullableProduct202407SearchSizeChartsResponse) Set(val *Product202407SearchSizeChartsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202407SearchSizeChartsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202407SearchSizeChartsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202407SearchSizeChartsResponse(val *Product202407SearchSizeChartsResponse) *NullableProduct202407SearchSizeChartsResponse {
	return &NullableProduct202407SearchSizeChartsResponse{value: val, isSet: true}
}

func (v NullableProduct202407SearchSizeChartsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202407SearchSizeChartsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


