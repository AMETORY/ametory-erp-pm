/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package return_refund_v202309

import (
    "encoding/json"
    "tiktokshop/open/sdk_golang/utils"
)

            // checks if the ReturnRefund202309CreateReturnResponseData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ReturnRefund202309CreateReturnResponseData{}

// ReturnRefund202309CreateReturnResponseData struct for ReturnRefund202309CreateReturnResponseData
type ReturnRefund202309CreateReturnResponseData struct {
    // The identifier of a specific return request.
    ReturnId *string `json:"return_id,omitempty"`
    // Return status, available values: - RETURN_OR_REFUND_REQUEST_PENDING: Request is pending, needs to be approved by seller or platform - REFUND_OR_RETURN_REQUEST_REJECT: Seller rejected the request - AWAITING_BUYER_SHIP: Waiting buyer to ship items to seller, if exceed the deadline, request will be closed by platform - BUYER_SHIPPED_ITEM: Buyer has shipped items to seller. - REJECT_RECEIVE_PACKAGE: Seller reject return package - RETURN_OR_REFUND_REQUEST_SUCCESS: The refund/return request is successful, buyer will be refunded. - RETURN_OR_REFUND_REQUEST_CANCEL: The request has been cancelled by buyer or system - RETURN_OR_REFUND_REQUEST_COMPLETE: The request is successful, and the amount has been refunded.
    ReturnStatus *string `json:"return_status,omitempty"`
}

// NewReturnRefund202309CreateReturnResponseData instantiates a new ReturnRefund202309CreateReturnResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnRefund202309CreateReturnResponseData() *ReturnRefund202309CreateReturnResponseData {
    this := ReturnRefund202309CreateReturnResponseData{}
    return &this
}

// NewReturnRefund202309CreateReturnResponseDataWithDefaults instantiates a new ReturnRefund202309CreateReturnResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnRefund202309CreateReturnResponseDataWithDefaults() *ReturnRefund202309CreateReturnResponseData {
    this := ReturnRefund202309CreateReturnResponseData{}
    return &this
}

// GetReturnId returns the ReturnId field value if set, zero value otherwise.
func (o *ReturnRefund202309CreateReturnResponseData) GetReturnId() string {
    if o == nil || utils.IsNil(o.ReturnId) {
        var ret string
        return ret
    }
    return *o.ReturnId
}

// GetReturnIdOk returns a tuple with the ReturnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309CreateReturnResponseData) GetReturnIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ReturnId) {
        return nil, false
    }
    return o.ReturnId, true
}

// HasReturnId returns a boolean if a field has been set.
func (o *ReturnRefund202309CreateReturnResponseData) HasReturnId() bool {
    if o != nil && !utils.IsNil(o.ReturnId) {
        return true
    }

    return false
}

// SetReturnId gets a reference to the given string and assigns it to the ReturnId field.
func (o *ReturnRefund202309CreateReturnResponseData) SetReturnId(v string) {
    o.ReturnId = &v
}

// GetReturnStatus returns the ReturnStatus field value if set, zero value otherwise.
func (o *ReturnRefund202309CreateReturnResponseData) GetReturnStatus() string {
    if o == nil || utils.IsNil(o.ReturnStatus) {
        var ret string
        return ret
    }
    return *o.ReturnStatus
}

// GetReturnStatusOk returns a tuple with the ReturnStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309CreateReturnResponseData) GetReturnStatusOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ReturnStatus) {
        return nil, false
    }
    return o.ReturnStatus, true
}

// HasReturnStatus returns a boolean if a field has been set.
func (o *ReturnRefund202309CreateReturnResponseData) HasReturnStatus() bool {
    if o != nil && !utils.IsNil(o.ReturnStatus) {
        return true
    }

    return false
}

// SetReturnStatus gets a reference to the given string and assigns it to the ReturnStatus field.
func (o *ReturnRefund202309CreateReturnResponseData) SetReturnStatus(v string) {
    o.ReturnStatus = &v
}

func (o ReturnRefund202309CreateReturnResponseData) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o ReturnRefund202309CreateReturnResponseData) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.ReturnId) {
        toSerialize["return_id"] = o.ReturnId
    }
    if !utils.IsNil(o.ReturnStatus) {
        toSerialize["return_status"] = o.ReturnStatus
    }
    return toSerialize, nil
}

type NullableReturnRefund202309CreateReturnResponseData struct {
	value *ReturnRefund202309CreateReturnResponseData
	isSet bool
}

func (v NullableReturnRefund202309CreateReturnResponseData) Get() *ReturnRefund202309CreateReturnResponseData {
	return v.value
}

func (v *NullableReturnRefund202309CreateReturnResponseData) Set(val *ReturnRefund202309CreateReturnResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnRefund202309CreateReturnResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnRefund202309CreateReturnResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnRefund202309CreateReturnResponseData(val *ReturnRefund202309CreateReturnResponseData) *NullableReturnRefund202309CreateReturnResponseData {
	return &NullableReturnRefund202309CreateReturnResponseData{value: val, isSet: true}
}

func (v NullableReturnRefund202309CreateReturnResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnRefund202309CreateReturnResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


