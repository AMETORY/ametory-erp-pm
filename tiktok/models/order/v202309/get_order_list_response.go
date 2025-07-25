/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package order_v202309

import (
    "encoding/json"
    "tiktokshop/open/sdk_golang/utils"
)

            // checks if the Order202309GetOrderListResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Order202309GetOrderListResponse{}

// Order202309GetOrderListResponse struct for Order202309GetOrderListResponse
type Order202309GetOrderListResponse struct {
    // The success or failure status code returned in API response.
    Code *int32 `json:"code,omitempty"`
    Data *Order202309GetOrderListResponseData `json:"data,omitempty"`
    // The success or failure messages returned in API response. Reasons of failure will be described in the message.
    Message *string `json:"message,omitempty"`
    // Request log.
    RequestId *string `json:"request_id,omitempty"`
}

// NewOrder202309GetOrderListResponse instantiates a new Order202309GetOrderListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrder202309GetOrderListResponse() *Order202309GetOrderListResponse {
    this := Order202309GetOrderListResponse{}
    return &this
}

// NewOrder202309GetOrderListResponseWithDefaults instantiates a new Order202309GetOrderListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrder202309GetOrderListResponseWithDefaults() *Order202309GetOrderListResponse {
    this := Order202309GetOrderListResponse{}
    return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Order202309GetOrderListResponse) GetCode() int32 {
    if o == nil || utils.IsNil(o.Code) {
        var ret int32
        return ret
    }
    return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order202309GetOrderListResponse) GetCodeOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Code) {
        return nil, false
    }
    return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Order202309GetOrderListResponse) HasCode() bool {
    if o != nil && !utils.IsNil(o.Code) {
        return true
    }

    return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *Order202309GetOrderListResponse) SetCode(v int32) {
    o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Order202309GetOrderListResponse) GetData() Order202309GetOrderListResponseData {
    if o == nil || utils.IsNil(o.Data) {
        var ret Order202309GetOrderListResponseData
        return ret
    }
    return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order202309GetOrderListResponse) GetDataOk() (*Order202309GetOrderListResponseData, bool) {
    if o == nil || utils.IsNil(o.Data) {
        return nil, false
    }
    return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Order202309GetOrderListResponse) HasData() bool {
    if o != nil && !utils.IsNil(o.Data) {
        return true
    }

    return false
}

// SetData gets a reference to the given Order202309GetOrderListResponseData and assigns it to the Data field.
func (o *Order202309GetOrderListResponse) SetData(v Order202309GetOrderListResponseData) {
    o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Order202309GetOrderListResponse) GetMessage() string {
    if o == nil || utils.IsNil(o.Message) {
        var ret string
        return ret
    }
    return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order202309GetOrderListResponse) GetMessageOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Message) {
        return nil, false
    }
    return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Order202309GetOrderListResponse) HasMessage() bool {
    if o != nil && !utils.IsNil(o.Message) {
        return true
    }

    return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Order202309GetOrderListResponse) SetMessage(v string) {
    o.Message = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *Order202309GetOrderListResponse) GetRequestId() string {
    if o == nil || utils.IsNil(o.RequestId) {
        var ret string
        return ret
    }
    return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order202309GetOrderListResponse) GetRequestIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.RequestId) {
        return nil, false
    }
    return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *Order202309GetOrderListResponse) HasRequestId() bool {
    if o != nil && !utils.IsNil(o.RequestId) {
        return true
    }

    return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *Order202309GetOrderListResponse) SetRequestId(v string) {
    o.RequestId = &v
}

func (o Order202309GetOrderListResponse) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Order202309GetOrderListResponse) ToMap() (map[string]interface{}, error) {
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

type NullableOrder202309GetOrderListResponse struct {
	value *Order202309GetOrderListResponse
	isSet bool
}

func (v NullableOrder202309GetOrderListResponse) Get() *Order202309GetOrderListResponse {
	return v.value
}

func (v *NullableOrder202309GetOrderListResponse) Set(val *Order202309GetOrderListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrder202309GetOrderListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrder202309GetOrderListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrder202309GetOrderListResponse(val *Order202309GetOrderListResponse) *NullableOrder202309GetOrderListResponse {
	return &NullableOrder202309GetOrderListResponse{value: val, isSet: true}
}

func (v NullableOrder202309GetOrderListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrder202309GetOrderListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


