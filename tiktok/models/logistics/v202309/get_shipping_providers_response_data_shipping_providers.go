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

            // checks if the Logistics202309GetShippingProvidersResponseDataShippingProviders type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Logistics202309GetShippingProvidersResponseDataShippingProviders{}

// Logistics202309GetShippingProvidersResponseDataShippingProviders struct for Logistics202309GetShippingProvidersResponseDataShippingProviders
type Logistics202309GetShippingProvidersResponseDataShippingProviders struct {
    // shipping provider id
    Id *string `json:"id,omitempty"`
    // shipping provider name
    Name *string `json:"name,omitempty"`
}

// NewLogistics202309GetShippingProvidersResponseDataShippingProviders instantiates a new Logistics202309GetShippingProvidersResponseDataShippingProviders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogistics202309GetShippingProvidersResponseDataShippingProviders() *Logistics202309GetShippingProvidersResponseDataShippingProviders {
    this := Logistics202309GetShippingProvidersResponseDataShippingProviders{}
    return &this
}

// NewLogistics202309GetShippingProvidersResponseDataShippingProvidersWithDefaults instantiates a new Logistics202309GetShippingProvidersResponseDataShippingProviders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogistics202309GetShippingProvidersResponseDataShippingProvidersWithDefaults() *Logistics202309GetShippingProvidersResponseDataShippingProviders {
    this := Logistics202309GetShippingProvidersResponseDataShippingProviders{}
    return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Logistics202309GetShippingProvidersResponseDataShippingProviders) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Logistics202309GetShippingProvidersResponseDataShippingProviders) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Logistics202309GetShippingProvidersResponseDataShippingProviders) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Logistics202309GetShippingProvidersResponseDataShippingProviders) SetId(v string) {
    o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Logistics202309GetShippingProvidersResponseDataShippingProviders) GetName() string {
    if o == nil || utils.IsNil(o.Name) {
        var ret string
        return ret
    }
    return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Logistics202309GetShippingProvidersResponseDataShippingProviders) GetNameOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Name) {
        return nil, false
    }
    return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Logistics202309GetShippingProvidersResponseDataShippingProviders) HasName() bool {
    if o != nil && !utils.IsNil(o.Name) {
        return true
    }

    return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Logistics202309GetShippingProvidersResponseDataShippingProviders) SetName(v string) {
    o.Name = &v
}

func (o Logistics202309GetShippingProvidersResponseDataShippingProviders) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Logistics202309GetShippingProvidersResponseDataShippingProviders) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    if !utils.IsNil(o.Name) {
        toSerialize["name"] = o.Name
    }
    return toSerialize, nil
}

type NullableLogistics202309GetShippingProvidersResponseDataShippingProviders struct {
	value *Logistics202309GetShippingProvidersResponseDataShippingProviders
	isSet bool
}

func (v NullableLogistics202309GetShippingProvidersResponseDataShippingProviders) Get() *Logistics202309GetShippingProvidersResponseDataShippingProviders {
	return v.value
}

func (v *NullableLogistics202309GetShippingProvidersResponseDataShippingProviders) Set(val *Logistics202309GetShippingProvidersResponseDataShippingProviders) {
	v.value = val
	v.isSet = true
}

func (v NullableLogistics202309GetShippingProvidersResponseDataShippingProviders) IsSet() bool {
	return v.isSet
}

func (v *NullableLogistics202309GetShippingProvidersResponseDataShippingProviders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogistics202309GetShippingProvidersResponseDataShippingProviders(val *Logistics202309GetShippingProvidersResponseDataShippingProviders) *NullableLogistics202309GetShippingProvidersResponseDataShippingProviders {
	return &NullableLogistics202309GetShippingProvidersResponseDataShippingProviders{value: val, isSet: true}
}

func (v NullableLogistics202309GetShippingProvidersResponseDataShippingProviders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogistics202309GetShippingProvidersResponseDataShippingProviders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


