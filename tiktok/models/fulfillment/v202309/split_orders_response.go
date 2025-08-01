/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fulfillment_v202309

import (
    "encoding/json"
    "tiktokshop/open/sdk_golang/utils"
)

            // checks if the Fulfillment202309SplitOrdersResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Fulfillment202309SplitOrdersResponse{}

// Fulfillment202309SplitOrdersResponse struct for Fulfillment202309SplitOrdersResponse
type Fulfillment202309SplitOrdersResponse struct {
    // The success or failure status code returned in API response.
    Code *int32 `json:"code,omitempty"`
    Data *Fulfillment202309SplitOrdersResponseData `json:"data,omitempty"`
    // The success or failure messages returned in API response. Reasons of failure will be described in the message.
    Message *string `json:"message,omitempty"`
    // Request log.
    RequestId *string `json:"request_id,omitempty"`
}

// NewFulfillment202309SplitOrdersResponse instantiates a new Fulfillment202309SplitOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillment202309SplitOrdersResponse() *Fulfillment202309SplitOrdersResponse {
    this := Fulfillment202309SplitOrdersResponse{}
    return &this
}

// NewFulfillment202309SplitOrdersResponseWithDefaults instantiates a new Fulfillment202309SplitOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillment202309SplitOrdersResponseWithDefaults() *Fulfillment202309SplitOrdersResponse {
    this := Fulfillment202309SplitOrdersResponse{}
    return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Fulfillment202309SplitOrdersResponse) GetCode() int32 {
    if o == nil || utils.IsNil(o.Code) {
        var ret int32
        return ret
    }
    return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309SplitOrdersResponse) GetCodeOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Code) {
        return nil, false
    }
    return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Fulfillment202309SplitOrdersResponse) HasCode() bool {
    if o != nil && !utils.IsNil(o.Code) {
        return true
    }

    return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *Fulfillment202309SplitOrdersResponse) SetCode(v int32) {
    o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Fulfillment202309SplitOrdersResponse) GetData() Fulfillment202309SplitOrdersResponseData {
    if o == nil || utils.IsNil(o.Data) {
        var ret Fulfillment202309SplitOrdersResponseData
        return ret
    }
    return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309SplitOrdersResponse) GetDataOk() (*Fulfillment202309SplitOrdersResponseData, bool) {
    if o == nil || utils.IsNil(o.Data) {
        return nil, false
    }
    return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Fulfillment202309SplitOrdersResponse) HasData() bool {
    if o != nil && !utils.IsNil(o.Data) {
        return true
    }

    return false
}

// SetData gets a reference to the given Fulfillment202309SplitOrdersResponseData and assigns it to the Data field.
func (o *Fulfillment202309SplitOrdersResponse) SetData(v Fulfillment202309SplitOrdersResponseData) {
    o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Fulfillment202309SplitOrdersResponse) GetMessage() string {
    if o == nil || utils.IsNil(o.Message) {
        var ret string
        return ret
    }
    return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309SplitOrdersResponse) GetMessageOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Message) {
        return nil, false
    }
    return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Fulfillment202309SplitOrdersResponse) HasMessage() bool {
    if o != nil && !utils.IsNil(o.Message) {
        return true
    }

    return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Fulfillment202309SplitOrdersResponse) SetMessage(v string) {
    o.Message = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *Fulfillment202309SplitOrdersResponse) GetRequestId() string {
    if o == nil || utils.IsNil(o.RequestId) {
        var ret string
        return ret
    }
    return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309SplitOrdersResponse) GetRequestIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.RequestId) {
        return nil, false
    }
    return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *Fulfillment202309SplitOrdersResponse) HasRequestId() bool {
    if o != nil && !utils.IsNil(o.RequestId) {
        return true
    }

    return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *Fulfillment202309SplitOrdersResponse) SetRequestId(v string) {
    o.RequestId = &v
}

func (o Fulfillment202309SplitOrdersResponse) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Fulfillment202309SplitOrdersResponse) ToMap() (map[string]interface{}, error) {
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

type NullableFulfillment202309SplitOrdersResponse struct {
	value *Fulfillment202309SplitOrdersResponse
	isSet bool
}

func (v NullableFulfillment202309SplitOrdersResponse) Get() *Fulfillment202309SplitOrdersResponse {
	return v.value
}

func (v *NullableFulfillment202309SplitOrdersResponse) Set(val *Fulfillment202309SplitOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillment202309SplitOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillment202309SplitOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillment202309SplitOrdersResponse(val *Fulfillment202309SplitOrdersResponse) *NullableFulfillment202309SplitOrdersResponse {
	return &NullableFulfillment202309SplitOrdersResponse{value: val, isSet: true}
}

func (v NullableFulfillment202309SplitOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFulfillment202309SplitOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


