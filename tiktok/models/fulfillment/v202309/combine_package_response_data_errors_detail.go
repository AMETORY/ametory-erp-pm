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

            // checks if the Fulfillment202309CombinePackageResponseDataErrorsDetail type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Fulfillment202309CombinePackageResponseDataErrorsDetail{}

// Fulfillment202309CombinePackageResponseDataErrorsDetail struct for Fulfillment202309CombinePackageResponseDataErrorsDetail
type Fulfillment202309CombinePackageResponseDataErrorsDetail struct {
    // Package ID.
    PackageId *string `json:"package_id,omitempty"`
}

// NewFulfillment202309CombinePackageResponseDataErrorsDetail instantiates a new Fulfillment202309CombinePackageResponseDataErrorsDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillment202309CombinePackageResponseDataErrorsDetail() *Fulfillment202309CombinePackageResponseDataErrorsDetail {
    this := Fulfillment202309CombinePackageResponseDataErrorsDetail{}
    return &this
}

// NewFulfillment202309CombinePackageResponseDataErrorsDetailWithDefaults instantiates a new Fulfillment202309CombinePackageResponseDataErrorsDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillment202309CombinePackageResponseDataErrorsDetailWithDefaults() *Fulfillment202309CombinePackageResponseDataErrorsDetail {
    this := Fulfillment202309CombinePackageResponseDataErrorsDetail{}
    return &this
}

// GetPackageId returns the PackageId field value if set, zero value otherwise.
func (o *Fulfillment202309CombinePackageResponseDataErrorsDetail) GetPackageId() string {
    if o == nil || utils.IsNil(o.PackageId) {
        var ret string
        return ret
    }
    return *o.PackageId
}

// GetPackageIdOk returns a tuple with the PackageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309CombinePackageResponseDataErrorsDetail) GetPackageIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.PackageId) {
        return nil, false
    }
    return o.PackageId, true
}

// HasPackageId returns a boolean if a field has been set.
func (o *Fulfillment202309CombinePackageResponseDataErrorsDetail) HasPackageId() bool {
    if o != nil && !utils.IsNil(o.PackageId) {
        return true
    }

    return false
}

// SetPackageId gets a reference to the given string and assigns it to the PackageId field.
func (o *Fulfillment202309CombinePackageResponseDataErrorsDetail) SetPackageId(v string) {
    o.PackageId = &v
}

func (o Fulfillment202309CombinePackageResponseDataErrorsDetail) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Fulfillment202309CombinePackageResponseDataErrorsDetail) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.PackageId) {
        toSerialize["package_id"] = o.PackageId
    }
    return toSerialize, nil
}

type NullableFulfillment202309CombinePackageResponseDataErrorsDetail struct {
	value *Fulfillment202309CombinePackageResponseDataErrorsDetail
	isSet bool
}

func (v NullableFulfillment202309CombinePackageResponseDataErrorsDetail) Get() *Fulfillment202309CombinePackageResponseDataErrorsDetail {
	return v.value
}

func (v *NullableFulfillment202309CombinePackageResponseDataErrorsDetail) Set(val *Fulfillment202309CombinePackageResponseDataErrorsDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillment202309CombinePackageResponseDataErrorsDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillment202309CombinePackageResponseDataErrorsDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillment202309CombinePackageResponseDataErrorsDetail(val *Fulfillment202309CombinePackageResponseDataErrorsDetail) *NullableFulfillment202309CombinePackageResponseDataErrorsDetail {
	return &NullableFulfillment202309CombinePackageResponseDataErrorsDetail{value: val, isSet: true}
}

func (v NullableFulfillment202309CombinePackageResponseDataErrorsDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFulfillment202309CombinePackageResponseDataErrorsDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


