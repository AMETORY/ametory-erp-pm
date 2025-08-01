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

            // checks if the Fulfillment202309MarkPackageAsShippedResponseData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Fulfillment202309MarkPackageAsShippedResponseData{}

// Fulfillment202309MarkPackageAsShippedResponseData struct for Fulfillment202309MarkPackageAsShippedResponseData
type Fulfillment202309MarkPackageAsShippedResponseData struct {
    // TikTok Shop order ID.
    OrderId *string `json:"order_id,omitempty"`
    // List of order line item IDs.
    OrderLineItemIds []string `json:"order_line_item_ids,omitempty"`
    // Package ID.
    PackageId *string `json:"package_id,omitempty"`
    Warning *Fulfillment202309MarkPackageAsShippedResponseDataWarning `json:"warning,omitempty"`
}

// NewFulfillment202309MarkPackageAsShippedResponseData instantiates a new Fulfillment202309MarkPackageAsShippedResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillment202309MarkPackageAsShippedResponseData() *Fulfillment202309MarkPackageAsShippedResponseData {
    this := Fulfillment202309MarkPackageAsShippedResponseData{}
    return &this
}

// NewFulfillment202309MarkPackageAsShippedResponseDataWithDefaults instantiates a new Fulfillment202309MarkPackageAsShippedResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillment202309MarkPackageAsShippedResponseDataWithDefaults() *Fulfillment202309MarkPackageAsShippedResponseData {
    this := Fulfillment202309MarkPackageAsShippedResponseData{}
    return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *Fulfillment202309MarkPackageAsShippedResponseData) GetOrderId() string {
    if o == nil || utils.IsNil(o.OrderId) {
        var ret string
        return ret
    }
    return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309MarkPackageAsShippedResponseData) GetOrderIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.OrderId) {
        return nil, false
    }
    return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *Fulfillment202309MarkPackageAsShippedResponseData) HasOrderId() bool {
    if o != nil && !utils.IsNil(o.OrderId) {
        return true
    }

    return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *Fulfillment202309MarkPackageAsShippedResponseData) SetOrderId(v string) {
    o.OrderId = &v
}

// GetOrderLineItemIds returns the OrderLineItemIds field value if set, zero value otherwise.
func (o *Fulfillment202309MarkPackageAsShippedResponseData) GetOrderLineItemIds() []string {
    if o == nil || utils.IsNil(o.OrderLineItemIds) {
        var ret []string
        return ret
    }
    return o.OrderLineItemIds
}

// GetOrderLineItemIdsOk returns a tuple with the OrderLineItemIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309MarkPackageAsShippedResponseData) GetOrderLineItemIdsOk() ([]string, bool) {
    if o == nil || utils.IsNil(o.OrderLineItemIds) {
        return nil, false
    }
    return o.OrderLineItemIds, true
}

// HasOrderLineItemIds returns a boolean if a field has been set.
func (o *Fulfillment202309MarkPackageAsShippedResponseData) HasOrderLineItemIds() bool {
    if o != nil && !utils.IsNil(o.OrderLineItemIds) {
        return true
    }

    return false
}

// SetOrderLineItemIds gets a reference to the given []string and assigns it to the OrderLineItemIds field.
func (o *Fulfillment202309MarkPackageAsShippedResponseData) SetOrderLineItemIds(v []string) {
    o.OrderLineItemIds = v
}

// GetPackageId returns the PackageId field value if set, zero value otherwise.
func (o *Fulfillment202309MarkPackageAsShippedResponseData) GetPackageId() string {
    if o == nil || utils.IsNil(o.PackageId) {
        var ret string
        return ret
    }
    return *o.PackageId
}

// GetPackageIdOk returns a tuple with the PackageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309MarkPackageAsShippedResponseData) GetPackageIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.PackageId) {
        return nil, false
    }
    return o.PackageId, true
}

// HasPackageId returns a boolean if a field has been set.
func (o *Fulfillment202309MarkPackageAsShippedResponseData) HasPackageId() bool {
    if o != nil && !utils.IsNil(o.PackageId) {
        return true
    }

    return false
}

// SetPackageId gets a reference to the given string and assigns it to the PackageId field.
func (o *Fulfillment202309MarkPackageAsShippedResponseData) SetPackageId(v string) {
    o.PackageId = &v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *Fulfillment202309MarkPackageAsShippedResponseData) GetWarning() Fulfillment202309MarkPackageAsShippedResponseDataWarning {
    if o == nil || utils.IsNil(o.Warning) {
        var ret Fulfillment202309MarkPackageAsShippedResponseDataWarning
        return ret
    }
    return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309MarkPackageAsShippedResponseData) GetWarningOk() (*Fulfillment202309MarkPackageAsShippedResponseDataWarning, bool) {
    if o == nil || utils.IsNil(o.Warning) {
        return nil, false
    }
    return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *Fulfillment202309MarkPackageAsShippedResponseData) HasWarning() bool {
    if o != nil && !utils.IsNil(o.Warning) {
        return true
    }

    return false
}

// SetWarning gets a reference to the given Fulfillment202309MarkPackageAsShippedResponseDataWarning and assigns it to the Warning field.
func (o *Fulfillment202309MarkPackageAsShippedResponseData) SetWarning(v Fulfillment202309MarkPackageAsShippedResponseDataWarning) {
    o.Warning = &v
}

func (o Fulfillment202309MarkPackageAsShippedResponseData) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Fulfillment202309MarkPackageAsShippedResponseData) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.OrderId) {
        toSerialize["order_id"] = o.OrderId
    }
    if !utils.IsNil(o.OrderLineItemIds) {
        toSerialize["order_line_item_ids"] = o.OrderLineItemIds
    }
    if !utils.IsNil(o.PackageId) {
        toSerialize["package_id"] = o.PackageId
    }
    if !utils.IsNil(o.Warning) {
        toSerialize["warning"] = o.Warning
    }
    return toSerialize, nil
}

type NullableFulfillment202309MarkPackageAsShippedResponseData struct {
	value *Fulfillment202309MarkPackageAsShippedResponseData
	isSet bool
}

func (v NullableFulfillment202309MarkPackageAsShippedResponseData) Get() *Fulfillment202309MarkPackageAsShippedResponseData {
	return v.value
}

func (v *NullableFulfillment202309MarkPackageAsShippedResponseData) Set(val *Fulfillment202309MarkPackageAsShippedResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillment202309MarkPackageAsShippedResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillment202309MarkPackageAsShippedResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillment202309MarkPackageAsShippedResponseData(val *Fulfillment202309MarkPackageAsShippedResponseData) *NullableFulfillment202309MarkPackageAsShippedResponseData {
	return &NullableFulfillment202309MarkPackageAsShippedResponseData{value: val, isSet: true}
}

func (v NullableFulfillment202309MarkPackageAsShippedResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFulfillment202309MarkPackageAsShippedResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


