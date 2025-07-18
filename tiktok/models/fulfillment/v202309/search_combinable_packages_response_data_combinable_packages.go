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

            // checks if the Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages{}

// Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages struct for Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages
type Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages struct {
    // A set of pre-generated package IDs after calling the Search Draft Package API. These package IDs will be used when the package combine is confirmed.
    Id *string `json:"id,omitempty"`
    // List of order IDs for this package.
    OrderIds []string `json:"order_ids,omitempty"`
}

// NewFulfillment202309SearchCombinablePackagesResponseDataCombinablePackages instantiates a new Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillment202309SearchCombinablePackagesResponseDataCombinablePackages() *Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages {
    this := Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages{}
    return &this
}

// NewFulfillment202309SearchCombinablePackagesResponseDataCombinablePackagesWithDefaults instantiates a new Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillment202309SearchCombinablePackagesResponseDataCombinablePackagesWithDefaults() *Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages {
    this := Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages{}
    return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages) SetId(v string) {
    o.Id = &v
}

// GetOrderIds returns the OrderIds field value if set, zero value otherwise.
func (o *Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages) GetOrderIds() []string {
    if o == nil || utils.IsNil(o.OrderIds) {
        var ret []string
        return ret
    }
    return o.OrderIds
}

// GetOrderIdsOk returns a tuple with the OrderIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages) GetOrderIdsOk() ([]string, bool) {
    if o == nil || utils.IsNil(o.OrderIds) {
        return nil, false
    }
    return o.OrderIds, true
}

// HasOrderIds returns a boolean if a field has been set.
func (o *Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages) HasOrderIds() bool {
    if o != nil && !utils.IsNil(o.OrderIds) {
        return true
    }

    return false
}

// SetOrderIds gets a reference to the given []string and assigns it to the OrderIds field.
func (o *Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages) SetOrderIds(v []string) {
    o.OrderIds = v
}

func (o Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    if !utils.IsNil(o.OrderIds) {
        toSerialize["order_ids"] = o.OrderIds
    }
    return toSerialize, nil
}

type NullableFulfillment202309SearchCombinablePackagesResponseDataCombinablePackages struct {
	value *Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages
	isSet bool
}

func (v NullableFulfillment202309SearchCombinablePackagesResponseDataCombinablePackages) Get() *Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages {
	return v.value
}

func (v *NullableFulfillment202309SearchCombinablePackagesResponseDataCombinablePackages) Set(val *Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillment202309SearchCombinablePackagesResponseDataCombinablePackages) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillment202309SearchCombinablePackagesResponseDataCombinablePackages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillment202309SearchCombinablePackagesResponseDataCombinablePackages(val *Fulfillment202309SearchCombinablePackagesResponseDataCombinablePackages) *NullableFulfillment202309SearchCombinablePackagesResponseDataCombinablePackages {
	return &NullableFulfillment202309SearchCombinablePackagesResponseDataCombinablePackages{value: val, isSet: true}
}

func (v NullableFulfillment202309SearchCombinablePackagesResponseDataCombinablePackages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFulfillment202309SearchCombinablePackagesResponseDataCombinablePackages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


