/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logistics_v202309

import (
    "encoding/json"
    "tiktokshop/open/sdk_golang/utils"
)

            // checks if the Logistics202309GetShippingProvidersResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Logistics202309GetShippingProvidersResponse{}

// Logistics202309GetShippingProvidersResponse struct for Logistics202309GetShippingProvidersResponse
type Logistics202309GetShippingProvidersResponse struct {
    // The success or failure status code returned in API response.
    Code *int32 `json:"code,omitempty"`
    Data *Logistics202309GetShippingProvidersResponseData `json:"data,omitempty"`
    // The success or failure messages returned in API response. Reasons of failure will be described in the message.
    Message *string `json:"message,omitempty"`
    // Request log.
    RequestId *string `json:"request_id,omitempty"`
}

// NewLogistics202309GetShippingProvidersResponse instantiates a new Logistics202309GetShippingProvidersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogistics202309GetShippingProvidersResponse() *Logistics202309GetShippingProvidersResponse {
    this := Logistics202309GetShippingProvidersResponse{}
    return &this
}

// NewLogistics202309GetShippingProvidersResponseWithDefaults instantiates a new Logistics202309GetShippingProvidersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogistics202309GetShippingProvidersResponseWithDefaults() *Logistics202309GetShippingProvidersResponse {
    this := Logistics202309GetShippingProvidersResponse{}
    return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Logistics202309GetShippingProvidersResponse) GetCode() int32 {
    if o == nil || utils.IsNil(o.Code) {
        var ret int32
        return ret
    }
    return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Logistics202309GetShippingProvidersResponse) GetCodeOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Code) {
        return nil, false
    }
    return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Logistics202309GetShippingProvidersResponse) HasCode() bool {
    if o != nil && !utils.IsNil(o.Code) {
        return true
    }

    return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *Logistics202309GetShippingProvidersResponse) SetCode(v int32) {
    o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Logistics202309GetShippingProvidersResponse) GetData() Logistics202309GetShippingProvidersResponseData {
    if o == nil || utils.IsNil(o.Data) {
        var ret Logistics202309GetShippingProvidersResponseData
        return ret
    }
    return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Logistics202309GetShippingProvidersResponse) GetDataOk() (*Logistics202309GetShippingProvidersResponseData, bool) {
    if o == nil || utils.IsNil(o.Data) {
        return nil, false
    }
    return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Logistics202309GetShippingProvidersResponse) HasData() bool {
    if o != nil && !utils.IsNil(o.Data) {
        return true
    }

    return false
}

// SetData gets a reference to the given Logistics202309GetShippingProvidersResponseData and assigns it to the Data field.
func (o *Logistics202309GetShippingProvidersResponse) SetData(v Logistics202309GetShippingProvidersResponseData) {
    o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Logistics202309GetShippingProvidersResponse) GetMessage() string {
    if o == nil || utils.IsNil(o.Message) {
        var ret string
        return ret
    }
    return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Logistics202309GetShippingProvidersResponse) GetMessageOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Message) {
        return nil, false
    }
    return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Logistics202309GetShippingProvidersResponse) HasMessage() bool {
    if o != nil && !utils.IsNil(o.Message) {
        return true
    }

    return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Logistics202309GetShippingProvidersResponse) SetMessage(v string) {
    o.Message = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *Logistics202309GetShippingProvidersResponse) GetRequestId() string {
    if o == nil || utils.IsNil(o.RequestId) {
        var ret string
        return ret
    }
    return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Logistics202309GetShippingProvidersResponse) GetRequestIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.RequestId) {
        return nil, false
    }
    return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *Logistics202309GetShippingProvidersResponse) HasRequestId() bool {
    if o != nil && !utils.IsNil(o.RequestId) {
        return true
    }

    return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *Logistics202309GetShippingProvidersResponse) SetRequestId(v string) {
    o.RequestId = &v
}

func (o Logistics202309GetShippingProvidersResponse) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Logistics202309GetShippingProvidersResponse) ToMap() (map[string]interface{}, error) {
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

type NullableLogistics202309GetShippingProvidersResponse struct {
	value *Logistics202309GetShippingProvidersResponse
	isSet bool
}

func (v NullableLogistics202309GetShippingProvidersResponse) Get() *Logistics202309GetShippingProvidersResponse {
	return v.value
}

func (v *NullableLogistics202309GetShippingProvidersResponse) Set(val *Logistics202309GetShippingProvidersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLogistics202309GetShippingProvidersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLogistics202309GetShippingProvidersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogistics202309GetShippingProvidersResponse(val *Logistics202309GetShippingProvidersResponse) *NullableLogistics202309GetShippingProvidersResponse {
	return &NullableLogistics202309GetShippingProvidersResponse{value: val, isSet: true}
}

func (v NullableLogistics202309GetShippingProvidersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogistics202309GetShippingProvidersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


