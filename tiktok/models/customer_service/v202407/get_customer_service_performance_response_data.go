/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package customer_service_v202407

import (
    "encoding/json"
    "tiktokshop/open/sdk_golang/utils"
)

            // checks if the CustomerService202407GetCustomerServicePerformanceResponseData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CustomerService202407GetCustomerServicePerformanceResponseData{}

// CustomerService202407GetCustomerServicePerformanceResponseData struct for CustomerService202407GetCustomerServicePerformanceResponseData
type CustomerService202407GetCustomerServicePerformanceResponseData struct {
    Performance *CustomerService202407GetCustomerServicePerformanceResponseDataPerformance `json:"performance,omitempty"`
}

// NewCustomerService202407GetCustomerServicePerformanceResponseData instantiates a new CustomerService202407GetCustomerServicePerformanceResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerService202407GetCustomerServicePerformanceResponseData() *CustomerService202407GetCustomerServicePerformanceResponseData {
    this := CustomerService202407GetCustomerServicePerformanceResponseData{}
    return &this
}

// NewCustomerService202407GetCustomerServicePerformanceResponseDataWithDefaults instantiates a new CustomerService202407GetCustomerServicePerformanceResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerService202407GetCustomerServicePerformanceResponseDataWithDefaults() *CustomerService202407GetCustomerServicePerformanceResponseData {
    this := CustomerService202407GetCustomerServicePerformanceResponseData{}
    return &this
}

// GetPerformance returns the Performance field value if set, zero value otherwise.
func (o *CustomerService202407GetCustomerServicePerformanceResponseData) GetPerformance() CustomerService202407GetCustomerServicePerformanceResponseDataPerformance {
    if o == nil || utils.IsNil(o.Performance) {
        var ret CustomerService202407GetCustomerServicePerformanceResponseDataPerformance
        return ret
    }
    return *o.Performance
}

// GetPerformanceOk returns a tuple with the Performance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerService202407GetCustomerServicePerformanceResponseData) GetPerformanceOk() (*CustomerService202407GetCustomerServicePerformanceResponseDataPerformance, bool) {
    if o == nil || utils.IsNil(o.Performance) {
        return nil, false
    }
    return o.Performance, true
}

// HasPerformance returns a boolean if a field has been set.
func (o *CustomerService202407GetCustomerServicePerformanceResponseData) HasPerformance() bool {
    if o != nil && !utils.IsNil(o.Performance) {
        return true
    }

    return false
}

// SetPerformance gets a reference to the given CustomerService202407GetCustomerServicePerformanceResponseDataPerformance and assigns it to the Performance field.
func (o *CustomerService202407GetCustomerServicePerformanceResponseData) SetPerformance(v CustomerService202407GetCustomerServicePerformanceResponseDataPerformance) {
    o.Performance = &v
}

func (o CustomerService202407GetCustomerServicePerformanceResponseData) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o CustomerService202407GetCustomerServicePerformanceResponseData) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Performance) {
        toSerialize["performance"] = o.Performance
    }
    return toSerialize, nil
}

type NullableCustomerService202407GetCustomerServicePerformanceResponseData struct {
	value *CustomerService202407GetCustomerServicePerformanceResponseData
	isSet bool
}

func (v NullableCustomerService202407GetCustomerServicePerformanceResponseData) Get() *CustomerService202407GetCustomerServicePerformanceResponseData {
	return v.value
}

func (v *NullableCustomerService202407GetCustomerServicePerformanceResponseData) Set(val *CustomerService202407GetCustomerServicePerformanceResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerService202407GetCustomerServicePerformanceResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerService202407GetCustomerServicePerformanceResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerService202407GetCustomerServicePerformanceResponseData(val *CustomerService202407GetCustomerServicePerformanceResponseData) *NullableCustomerService202407GetCustomerServicePerformanceResponseData {
	return &NullableCustomerService202407GetCustomerServicePerformanceResponseData{value: val, isSet: true}
}

func (v NullableCustomerService202407GetCustomerServicePerformanceResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerService202407GetCustomerServicePerformanceResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


