/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product_v202404

import (
    "encoding/json"
    "tiktokshop/open/sdk_golang/utils"
)

            // checks if the Product202404OptimizedImagesResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202404OptimizedImagesResponse{}

// Product202404OptimizedImagesResponse struct for Product202404OptimizedImagesResponse
type Product202404OptimizedImagesResponse struct {
    // The success or failure status code returned in API response.
    Code *int32 `json:"code,omitempty"`
    Data *Product202404OptimizedImagesResponseData `json:"data,omitempty"`
    // The success or failure messages returned in API response. Reasons of failure will be described in the message.
    Message *string `json:"message,omitempty"`
    // Request log.
    RequestId *string `json:"request_id,omitempty"`
}

// NewProduct202404OptimizedImagesResponse instantiates a new Product202404OptimizedImagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202404OptimizedImagesResponse() *Product202404OptimizedImagesResponse {
    this := Product202404OptimizedImagesResponse{}
    return &this
}

// NewProduct202404OptimizedImagesResponseWithDefaults instantiates a new Product202404OptimizedImagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202404OptimizedImagesResponseWithDefaults() *Product202404OptimizedImagesResponse {
    this := Product202404OptimizedImagesResponse{}
    return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Product202404OptimizedImagesResponse) GetCode() int32 {
    if o == nil || utils.IsNil(o.Code) {
        var ret int32
        return ret
    }
    return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202404OptimizedImagesResponse) GetCodeOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Code) {
        return nil, false
    }
    return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Product202404OptimizedImagesResponse) HasCode() bool {
    if o != nil && !utils.IsNil(o.Code) {
        return true
    }

    return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *Product202404OptimizedImagesResponse) SetCode(v int32) {
    o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Product202404OptimizedImagesResponse) GetData() Product202404OptimizedImagesResponseData {
    if o == nil || utils.IsNil(o.Data) {
        var ret Product202404OptimizedImagesResponseData
        return ret
    }
    return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202404OptimizedImagesResponse) GetDataOk() (*Product202404OptimizedImagesResponseData, bool) {
    if o == nil || utils.IsNil(o.Data) {
        return nil, false
    }
    return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Product202404OptimizedImagesResponse) HasData() bool {
    if o != nil && !utils.IsNil(o.Data) {
        return true
    }

    return false
}

// SetData gets a reference to the given Product202404OptimizedImagesResponseData and assigns it to the Data field.
func (o *Product202404OptimizedImagesResponse) SetData(v Product202404OptimizedImagesResponseData) {
    o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Product202404OptimizedImagesResponse) GetMessage() string {
    if o == nil || utils.IsNil(o.Message) {
        var ret string
        return ret
    }
    return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202404OptimizedImagesResponse) GetMessageOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Message) {
        return nil, false
    }
    return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Product202404OptimizedImagesResponse) HasMessage() bool {
    if o != nil && !utils.IsNil(o.Message) {
        return true
    }

    return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Product202404OptimizedImagesResponse) SetMessage(v string) {
    o.Message = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *Product202404OptimizedImagesResponse) GetRequestId() string {
    if o == nil || utils.IsNil(o.RequestId) {
        var ret string
        return ret
    }
    return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202404OptimizedImagesResponse) GetRequestIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.RequestId) {
        return nil, false
    }
    return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *Product202404OptimizedImagesResponse) HasRequestId() bool {
    if o != nil && !utils.IsNil(o.RequestId) {
        return true
    }

    return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *Product202404OptimizedImagesResponse) SetRequestId(v string) {
    o.RequestId = &v
}

func (o Product202404OptimizedImagesResponse) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202404OptimizedImagesResponse) ToMap() (map[string]interface{}, error) {
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

type NullableProduct202404OptimizedImagesResponse struct {
	value *Product202404OptimizedImagesResponse
	isSet bool
}

func (v NullableProduct202404OptimizedImagesResponse) Get() *Product202404OptimizedImagesResponse {
	return v.value
}

func (v *NullableProduct202404OptimizedImagesResponse) Set(val *Product202404OptimizedImagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202404OptimizedImagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202404OptimizedImagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202404OptimizedImagesResponse(val *Product202404OptimizedImagesResponse) *NullableProduct202404OptimizedImagesResponse {
	return &NullableProduct202404OptimizedImagesResponse{value: val, isSet: true}
}

func (v NullableProduct202404OptimizedImagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202404OptimizedImagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


