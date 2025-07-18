/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package finance_v202309

import (
    "encoding/json"
    "tiktokshop/open/sdk_golang/utils"
)

            // checks if the Finance202309GetPaymentsResponseData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Finance202309GetPaymentsResponseData{}

// Finance202309GetPaymentsResponseData struct for Finance202309GetPaymentsResponseData
type Finance202309GetPaymentsResponseData struct {
    // An opaque token used to retrieve the next page of a paginated result set. Provide this value in the `page_token` parameter of your request if the current response does not return all the results.
    NextPageToken *string `json:"next_page_token,omitempty"`
    // The list of payments that meet the query conditions.
    Payments []Finance202309GetPaymentsResponseDataPayments `json:"payments,omitempty"`
}

// NewFinance202309GetPaymentsResponseData instantiates a new Finance202309GetPaymentsResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinance202309GetPaymentsResponseData() *Finance202309GetPaymentsResponseData {
    this := Finance202309GetPaymentsResponseData{}
    return &this
}

// NewFinance202309GetPaymentsResponseDataWithDefaults instantiates a new Finance202309GetPaymentsResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinance202309GetPaymentsResponseDataWithDefaults() *Finance202309GetPaymentsResponseData {
    this := Finance202309GetPaymentsResponseData{}
    return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *Finance202309GetPaymentsResponseData) GetNextPageToken() string {
    if o == nil || utils.IsNil(o.NextPageToken) {
        var ret string
        return ret
    }
    return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202309GetPaymentsResponseData) GetNextPageTokenOk() (*string, bool) {
    if o == nil || utils.IsNil(o.NextPageToken) {
        return nil, false
    }
    return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *Finance202309GetPaymentsResponseData) HasNextPageToken() bool {
    if o != nil && !utils.IsNil(o.NextPageToken) {
        return true
    }

    return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *Finance202309GetPaymentsResponseData) SetNextPageToken(v string) {
    o.NextPageToken = &v
}

// GetPayments returns the Payments field value if set, zero value otherwise.
func (o *Finance202309GetPaymentsResponseData) GetPayments() []Finance202309GetPaymentsResponseDataPayments {
    if o == nil || utils.IsNil(o.Payments) {
        var ret []Finance202309GetPaymentsResponseDataPayments
        return ret
    }
    return o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Finance202309GetPaymentsResponseData) GetPaymentsOk() ([]Finance202309GetPaymentsResponseDataPayments, bool) {
    if o == nil || utils.IsNil(o.Payments) {
        return nil, false
    }
    return o.Payments, true
}

// HasPayments returns a boolean if a field has been set.
func (o *Finance202309GetPaymentsResponseData) HasPayments() bool {
    if o != nil && !utils.IsNil(o.Payments) {
        return true
    }

    return false
}

// SetPayments gets a reference to the given []Finance202309GetPaymentsResponseDataPayments and assigns it to the Payments field.
func (o *Finance202309GetPaymentsResponseData) SetPayments(v []Finance202309GetPaymentsResponseDataPayments) {
    o.Payments = v
}

func (o Finance202309GetPaymentsResponseData) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Finance202309GetPaymentsResponseData) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.NextPageToken) {
        toSerialize["next_page_token"] = o.NextPageToken
    }
    if !utils.IsNil(o.Payments) {
        toSerialize["payments"] = o.Payments
    }
    return toSerialize, nil
}

type NullableFinance202309GetPaymentsResponseData struct {
	value *Finance202309GetPaymentsResponseData
	isSet bool
}

func (v NullableFinance202309GetPaymentsResponseData) Get() *Finance202309GetPaymentsResponseData {
	return v.value
}

func (v *NullableFinance202309GetPaymentsResponseData) Set(val *Finance202309GetPaymentsResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableFinance202309GetPaymentsResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableFinance202309GetPaymentsResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinance202309GetPaymentsResponseData(val *Finance202309GetPaymentsResponseData) *NullableFinance202309GetPaymentsResponseData {
	return &NullableFinance202309GetPaymentsResponseData{value: val, isSet: true}
}

func (v NullableFinance202309GetPaymentsResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinance202309GetPaymentsResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


