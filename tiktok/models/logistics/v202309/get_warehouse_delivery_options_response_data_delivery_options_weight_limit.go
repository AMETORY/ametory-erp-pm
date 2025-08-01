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

            // checks if the Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit{}

// Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit struct for Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit
type Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit struct {
    // Maximum weight limit.
    MaxWeight *int64 `json:"max_weight,omitempty"`
    // Minimum weight limit.
    MinWeight *int64 `json:"min_weight,omitempty"`
    // The unit of measurement for the weight, with possible values: - GRAM - POUND
    Unit *string `json:"unit,omitempty"`
}

// NewLogistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit instantiates a new Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit() *Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit {
    this := Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit{}
    return &this
}

// NewLogistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimitWithDefaults instantiates a new Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimitWithDefaults() *Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit {
    this := Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit{}
    return &this
}

// GetMaxWeight returns the MaxWeight field value if set, zero value otherwise.
func (o *Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit) GetMaxWeight() int64 {
    if o == nil || utils.IsNil(o.MaxWeight) {
        var ret int64
        return ret
    }
    return *o.MaxWeight
}

// GetMaxWeightOk returns a tuple with the MaxWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit) GetMaxWeightOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.MaxWeight) {
        return nil, false
    }
    return o.MaxWeight, true
}

// HasMaxWeight returns a boolean if a field has been set.
func (o *Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit) HasMaxWeight() bool {
    if o != nil && !utils.IsNil(o.MaxWeight) {
        return true
    }

    return false
}

// SetMaxWeight gets a reference to the given int64 and assigns it to the MaxWeight field.
func (o *Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit) SetMaxWeight(v int64) {
    o.MaxWeight = &v
}

// GetMinWeight returns the MinWeight field value if set, zero value otherwise.
func (o *Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit) GetMinWeight() int64 {
    if o == nil || utils.IsNil(o.MinWeight) {
        var ret int64
        return ret
    }
    return *o.MinWeight
}

// GetMinWeightOk returns a tuple with the MinWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit) GetMinWeightOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.MinWeight) {
        return nil, false
    }
    return o.MinWeight, true
}

// HasMinWeight returns a boolean if a field has been set.
func (o *Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit) HasMinWeight() bool {
    if o != nil && !utils.IsNil(o.MinWeight) {
        return true
    }

    return false
}

// SetMinWeight gets a reference to the given int64 and assigns it to the MinWeight field.
func (o *Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit) SetMinWeight(v int64) {
    o.MinWeight = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit) GetUnit() string {
    if o == nil || utils.IsNil(o.Unit) {
        var ret string
        return ret
    }
    return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit) GetUnitOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Unit) {
        return nil, false
    }
    return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit) HasUnit() bool {
    if o != nil && !utils.IsNil(o.Unit) {
        return true
    }

    return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit) SetUnit(v string) {
    o.Unit = &v
}

func (o Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.MaxWeight) {
        toSerialize["max_weight"] = o.MaxWeight
    }
    if !utils.IsNil(o.MinWeight) {
        toSerialize["min_weight"] = o.MinWeight
    }
    if !utils.IsNil(o.Unit) {
        toSerialize["unit"] = o.Unit
    }
    return toSerialize, nil
}

type NullableLogistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit struct {
	value *Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit
	isSet bool
}

func (v NullableLogistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit) Get() *Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit {
	return v.value
}

func (v *NullableLogistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit) Set(val *Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableLogistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableLogistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit(val *Logistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit) *NullableLogistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit {
	return &NullableLogistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit{value: val, isSet: true}
}

func (v NullableLogistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogistics202309GetWarehouseDeliveryOptionsResponseDataDeliveryOptionsWeightLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


