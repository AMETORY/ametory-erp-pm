/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package customer_service_v202309

import (
    "encoding/json"
    "tiktokshop/open/sdk_golang/utils"
)

            // checks if the CustomerService202309UpdateAgentSettingsRequestBody type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CustomerService202309UpdateAgentSettingsRequestBody{}

// CustomerService202309UpdateAgentSettingsRequestBody struct for CustomerService202309UpdateAgentSettingsRequestBody
type CustomerService202309UpdateAgentSettingsRequestBody struct {
    // If true, the agent will receive auto-assigned chats. The agent can manually select chats to respond.  If false, the agent will receive manually assigned chats only. When using IM API, we recommend setting this field to true.
    CanAcceptChat *bool `json:"can_accept_chat,omitempty"`
}

// NewCustomerService202309UpdateAgentSettingsRequestBody instantiates a new CustomerService202309UpdateAgentSettingsRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerService202309UpdateAgentSettingsRequestBody() *CustomerService202309UpdateAgentSettingsRequestBody {
    this := CustomerService202309UpdateAgentSettingsRequestBody{}
    return &this
}

// NewCustomerService202309UpdateAgentSettingsRequestBodyWithDefaults instantiates a new CustomerService202309UpdateAgentSettingsRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerService202309UpdateAgentSettingsRequestBodyWithDefaults() *CustomerService202309UpdateAgentSettingsRequestBody {
    this := CustomerService202309UpdateAgentSettingsRequestBody{}
    return &this
}

// GetCanAcceptChat returns the CanAcceptChat field value if set, zero value otherwise.
func (o *CustomerService202309UpdateAgentSettingsRequestBody) GetCanAcceptChat() bool {
    if o == nil || utils.IsNil(o.CanAcceptChat) {
        var ret bool
        return ret
    }
    return *o.CanAcceptChat
}

// GetCanAcceptChatOk returns a tuple with the CanAcceptChat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerService202309UpdateAgentSettingsRequestBody) GetCanAcceptChatOk() (*bool, bool) {
    if o == nil || utils.IsNil(o.CanAcceptChat) {
        return nil, false
    }
    return o.CanAcceptChat, true
}

// HasCanAcceptChat returns a boolean if a field has been set.
func (o *CustomerService202309UpdateAgentSettingsRequestBody) HasCanAcceptChat() bool {
    if o != nil && !utils.IsNil(o.CanAcceptChat) {
        return true
    }

    return false
}

// SetCanAcceptChat gets a reference to the given bool and assigns it to the CanAcceptChat field.
func (o *CustomerService202309UpdateAgentSettingsRequestBody) SetCanAcceptChat(v bool) {
    o.CanAcceptChat = &v
}

func (o CustomerService202309UpdateAgentSettingsRequestBody) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o CustomerService202309UpdateAgentSettingsRequestBody) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.CanAcceptChat) {
        toSerialize["can_accept_chat"] = o.CanAcceptChat
    }
    return toSerialize, nil
}

type NullableCustomerService202309UpdateAgentSettingsRequestBody struct {
	value *CustomerService202309UpdateAgentSettingsRequestBody
	isSet bool
}

func (v NullableCustomerService202309UpdateAgentSettingsRequestBody) Get() *CustomerService202309UpdateAgentSettingsRequestBody {
	return v.value
}

func (v *NullableCustomerService202309UpdateAgentSettingsRequestBody) Set(val *CustomerService202309UpdateAgentSettingsRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerService202309UpdateAgentSettingsRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerService202309UpdateAgentSettingsRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerService202309UpdateAgentSettingsRequestBody(val *CustomerService202309UpdateAgentSettingsRequestBody) *NullableCustomerService202309UpdateAgentSettingsRequestBody {
	return &NullableCustomerService202309UpdateAgentSettingsRequestBody{value: val, isSet: true}
}

func (v NullableCustomerService202309UpdateAgentSettingsRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerService202309UpdateAgentSettingsRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


