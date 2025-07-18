/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product_v202409

import (
    "encoding/json"
    "tiktokshop/open/sdk_golang/utils"
)

            // checks if the Product202409PartialEditResponsiblePersonRequestBodyAddress type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202409PartialEditResponsiblePersonRequestBodyAddress{}

// Product202409PartialEditResponsiblePersonRequestBodyAddress struct for Product202409PartialEditResponsiblePersonRequestBodyAddress
type Product202409PartialEditResponsiblePersonRequestBodyAddress struct {
    // (**Deprecated**: This field is deprecated and will be removed in a future API version. If provided, its value will be merged into `street_address_line1`. It is recommended to specify `street_address_line1` directly.) The city name. Max length: 500 characters
    City *string `json:"city,omitempty"`
    // The two letter ISO 3166 country code representing the country of the address. It must be an EU country.
    Country *string `json:"country,omitempty"`
    // (**Deprecated**: This field is deprecated and will be removed in a future API version. If provided, its value will be merged into `street_address_line1`. It is recommended to specify `street_address_line1` directly.) The district name. Max length: 500 characters
    District *string `json:"district,omitempty"`
    // The postal code. Max length: 500 characters
    PostalCode *string `json:"postal_code,omitempty"`
    // (**Deprecated**: This field is deprecated and will be removed in a future API version. If provided, its value will be merged into `street_address_line1`. It is recommended to specify `street_address_line1` directly.) The province, state, or region name. Max length: 500 characters
    Province *string `json:"province,omitempty"`
    // The detailed street address of the location, including the building number, street name, district, city, province, and any relevant details. Max length: 500 characters
    StreetAddressLine1 *string `json:"street_address_line1,omitempty"`
    // (**Deprecated**: This field is deprecated and will be removed in a future API version. If provided, its value will be merged into `street_address_line1`. It is recommended to specify `street_address_line1` directly.) An optional secondary line for additional address details, if necessary. Max length: 500 characters
    StreetAddressLine2 *string `json:"street_address_line2,omitempty"`
}

// NewProduct202409PartialEditResponsiblePersonRequestBodyAddress instantiates a new Product202409PartialEditResponsiblePersonRequestBodyAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202409PartialEditResponsiblePersonRequestBodyAddress() *Product202409PartialEditResponsiblePersonRequestBodyAddress {
    this := Product202409PartialEditResponsiblePersonRequestBodyAddress{}
    return &this
}

// NewProduct202409PartialEditResponsiblePersonRequestBodyAddressWithDefaults instantiates a new Product202409PartialEditResponsiblePersonRequestBodyAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202409PartialEditResponsiblePersonRequestBodyAddressWithDefaults() *Product202409PartialEditResponsiblePersonRequestBodyAddress {
    this := Product202409PartialEditResponsiblePersonRequestBodyAddress{}
    return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) GetCity() string {
    if o == nil || utils.IsNil(o.City) {
        var ret string
        return ret
    }
    return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) GetCityOk() (*string, bool) {
    if o == nil || utils.IsNil(o.City) {
        return nil, false
    }
    return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) HasCity() bool {
    if o != nil && !utils.IsNil(o.City) {
        return true
    }

    return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) SetCity(v string) {
    o.City = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) GetCountry() string {
    if o == nil || utils.IsNil(o.Country) {
        var ret string
        return ret
    }
    return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) GetCountryOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Country) {
        return nil, false
    }
    return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) HasCountry() bool {
    if o != nil && !utils.IsNil(o.Country) {
        return true
    }

    return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) SetCountry(v string) {
    o.Country = &v
}

// GetDistrict returns the District field value if set, zero value otherwise.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) GetDistrict() string {
    if o == nil || utils.IsNil(o.District) {
        var ret string
        return ret
    }
    return *o.District
}

// GetDistrictOk returns a tuple with the District field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) GetDistrictOk() (*string, bool) {
    if o == nil || utils.IsNil(o.District) {
        return nil, false
    }
    return o.District, true
}

// HasDistrict returns a boolean if a field has been set.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) HasDistrict() bool {
    if o != nil && !utils.IsNil(o.District) {
        return true
    }

    return false
}

// SetDistrict gets a reference to the given string and assigns it to the District field.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) SetDistrict(v string) {
    o.District = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) GetPostalCode() string {
    if o == nil || utils.IsNil(o.PostalCode) {
        var ret string
        return ret
    }
    return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) GetPostalCodeOk() (*string, bool) {
    if o == nil || utils.IsNil(o.PostalCode) {
        return nil, false
    }
    return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) HasPostalCode() bool {
    if o != nil && !utils.IsNil(o.PostalCode) {
        return true
    }

    return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) SetPostalCode(v string) {
    o.PostalCode = &v
}

// GetProvince returns the Province field value if set, zero value otherwise.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) GetProvince() string {
    if o == nil || utils.IsNil(o.Province) {
        var ret string
        return ret
    }
    return *o.Province
}

// GetProvinceOk returns a tuple with the Province field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) GetProvinceOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Province) {
        return nil, false
    }
    return o.Province, true
}

// HasProvince returns a boolean if a field has been set.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) HasProvince() bool {
    if o != nil && !utils.IsNil(o.Province) {
        return true
    }

    return false
}

// SetProvince gets a reference to the given string and assigns it to the Province field.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) SetProvince(v string) {
    o.Province = &v
}

// GetStreetAddressLine1 returns the StreetAddressLine1 field value if set, zero value otherwise.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) GetStreetAddressLine1() string {
    if o == nil || utils.IsNil(o.StreetAddressLine1) {
        var ret string
        return ret
    }
    return *o.StreetAddressLine1
}

// GetStreetAddressLine1Ok returns a tuple with the StreetAddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) GetStreetAddressLine1Ok() (*string, bool) {
    if o == nil || utils.IsNil(o.StreetAddressLine1) {
        return nil, false
    }
    return o.StreetAddressLine1, true
}

// HasStreetAddressLine1 returns a boolean if a field has been set.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) HasStreetAddressLine1() bool {
    if o != nil && !utils.IsNil(o.StreetAddressLine1) {
        return true
    }

    return false
}

// SetStreetAddressLine1 gets a reference to the given string and assigns it to the StreetAddressLine1 field.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) SetStreetAddressLine1(v string) {
    o.StreetAddressLine1 = &v
}

// GetStreetAddressLine2 returns the StreetAddressLine2 field value if set, zero value otherwise.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) GetStreetAddressLine2() string {
    if o == nil || utils.IsNil(o.StreetAddressLine2) {
        var ret string
        return ret
    }
    return *o.StreetAddressLine2
}

// GetStreetAddressLine2Ok returns a tuple with the StreetAddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) GetStreetAddressLine2Ok() (*string, bool) {
    if o == nil || utils.IsNil(o.StreetAddressLine2) {
        return nil, false
    }
    return o.StreetAddressLine2, true
}

// HasStreetAddressLine2 returns a boolean if a field has been set.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) HasStreetAddressLine2() bool {
    if o != nil && !utils.IsNil(o.StreetAddressLine2) {
        return true
    }

    return false
}

// SetStreetAddressLine2 gets a reference to the given string and assigns it to the StreetAddressLine2 field.
func (o *Product202409PartialEditResponsiblePersonRequestBodyAddress) SetStreetAddressLine2(v string) {
    o.StreetAddressLine2 = &v
}

func (o Product202409PartialEditResponsiblePersonRequestBodyAddress) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202409PartialEditResponsiblePersonRequestBodyAddress) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.City) {
        toSerialize["city"] = o.City
    }
    if !utils.IsNil(o.Country) {
        toSerialize["country"] = o.Country
    }
    if !utils.IsNil(o.District) {
        toSerialize["district"] = o.District
    }
    if !utils.IsNil(o.PostalCode) {
        toSerialize["postal_code"] = o.PostalCode
    }
    if !utils.IsNil(o.Province) {
        toSerialize["province"] = o.Province
    }
    if !utils.IsNil(o.StreetAddressLine1) {
        toSerialize["street_address_line1"] = o.StreetAddressLine1
    }
    if !utils.IsNil(o.StreetAddressLine2) {
        toSerialize["street_address_line2"] = o.StreetAddressLine2
    }
    return toSerialize, nil
}

type NullableProduct202409PartialEditResponsiblePersonRequestBodyAddress struct {
	value *Product202409PartialEditResponsiblePersonRequestBodyAddress
	isSet bool
}

func (v NullableProduct202409PartialEditResponsiblePersonRequestBodyAddress) Get() *Product202409PartialEditResponsiblePersonRequestBodyAddress {
	return v.value
}

func (v *NullableProduct202409PartialEditResponsiblePersonRequestBodyAddress) Set(val *Product202409PartialEditResponsiblePersonRequestBodyAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202409PartialEditResponsiblePersonRequestBodyAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202409PartialEditResponsiblePersonRequestBodyAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202409PartialEditResponsiblePersonRequestBodyAddress(val *Product202409PartialEditResponsiblePersonRequestBodyAddress) *NullableProduct202409PartialEditResponsiblePersonRequestBodyAddress {
	return &NullableProduct202409PartialEditResponsiblePersonRequestBodyAddress{value: val, isSet: true}
}

func (v NullableProduct202409PartialEditResponsiblePersonRequestBodyAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202409PartialEditResponsiblePersonRequestBodyAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


