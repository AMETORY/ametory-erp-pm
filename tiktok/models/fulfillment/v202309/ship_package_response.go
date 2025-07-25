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

            // checks if the Fulfillment202309ShipPackageResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Fulfillment202309ShipPackageResponse{}

// Fulfillment202309ShipPackageResponse struct for Fulfillment202309ShipPackageResponse
type Fulfillment202309ShipPackageResponse struct {
    // The success or failure status code returned in API response.
    Code *int32 `json:"code,omitempty"`
    Data map[string]interface{} `json:"data,omitempty"`
    // The success or failure messages returned in API response. Reasons of failure will be described in the message.
    Message *string `json:"message,omitempty"`
    // Request log.
    RequestId *string `json:"request_id,omitempty"`
}

// NewFulfillment202309ShipPackageResponse instantiates a new Fulfillment202309ShipPackageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillment202309ShipPackageResponse() *Fulfillment202309ShipPackageResponse {
    this := Fulfillment202309ShipPackageResponse{}
    return &this
}

// NewFulfillment202309ShipPackageResponseWithDefaults instantiates a new Fulfillment202309ShipPackageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillment202309ShipPackageResponseWithDefaults() *Fulfillment202309ShipPackageResponse {
    this := Fulfillment202309ShipPackageResponse{}
    return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Fulfillment202309ShipPackageResponse) GetCode() int32 {
    if o == nil || utils.IsNil(o.Code) {
        var ret int32
        return ret
    }
    return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309ShipPackageResponse) GetCodeOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Code) {
        return nil, false
    }
    return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Fulfillment202309ShipPackageResponse) HasCode() bool {
    if o != nil && !utils.IsNil(o.Code) {
        return true
    }

    return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *Fulfillment202309ShipPackageResponse) SetCode(v int32) {
    o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Fulfillment202309ShipPackageResponse) GetData() map[string]interface{} {
    if o == nil || utils.IsNil(o.Data) {
        var ret map[string]interface{}
        return ret
    }
    return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309ShipPackageResponse) GetDataOk() (map[string]interface{}, bool) {
    if o == nil || utils.IsNil(o.Data) {
        return map[string]interface{}{}, false
    }
    return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Fulfillment202309ShipPackageResponse) HasData() bool {
    if o != nil && !utils.IsNil(o.Data) {
        return true
    }

    return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *Fulfillment202309ShipPackageResponse) SetData(v map[string]interface{}) {
    o.Data = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Fulfillment202309ShipPackageResponse) GetMessage() string {
    if o == nil || utils.IsNil(o.Message) {
        var ret string
        return ret
    }
    return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309ShipPackageResponse) GetMessageOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Message) {
        return nil, false
    }
    return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Fulfillment202309ShipPackageResponse) HasMessage() bool {
    if o != nil && !utils.IsNil(o.Message) {
        return true
    }

    return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Fulfillment202309ShipPackageResponse) SetMessage(v string) {
    o.Message = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *Fulfillment202309ShipPackageResponse) GetRequestId() string {
    if o == nil || utils.IsNil(o.RequestId) {
        var ret string
        return ret
    }
    return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309ShipPackageResponse) GetRequestIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.RequestId) {
        return nil, false
    }
    return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *Fulfillment202309ShipPackageResponse) HasRequestId() bool {
    if o != nil && !utils.IsNil(o.RequestId) {
        return true
    }

    return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *Fulfillment202309ShipPackageResponse) SetRequestId(v string) {
    o.RequestId = &v
}

func (o Fulfillment202309ShipPackageResponse) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Fulfillment202309ShipPackageResponse) ToMap() (map[string]interface{}, error) {
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

type NullableFulfillment202309ShipPackageResponse struct {
	value *Fulfillment202309ShipPackageResponse
	isSet bool
}

func (v NullableFulfillment202309ShipPackageResponse) Get() *Fulfillment202309ShipPackageResponse {
	return v.value
}

func (v *NullableFulfillment202309ShipPackageResponse) Set(val *Fulfillment202309ShipPackageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillment202309ShipPackageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillment202309ShipPackageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillment202309ShipPackageResponse(val *Fulfillment202309ShipPackageResponse) *NullableFulfillment202309ShipPackageResponse {
	return &NullableFulfillment202309ShipPackageResponse{value: val, isSet: true}
}

func (v NullableFulfillment202309ShipPackageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFulfillment202309ShipPackageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


