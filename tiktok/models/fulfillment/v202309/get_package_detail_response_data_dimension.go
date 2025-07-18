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

            // checks if the Fulfillment202309GetPackageDetailResponseDataDimension type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Fulfillment202309GetPackageDetailResponseDataDimension{}

// Fulfillment202309GetPackageDetailResponseDataDimension struct for Fulfillment202309GetPackageDetailResponseDataDimension
type Fulfillment202309GetPackageDetailResponseDataDimension struct {
    // The height of the scheduled package. 
    Height *string `json:"height,omitempty"`
    // The length of the scheduled package. 
    Length *string `json:"length,omitempty"`
    // The unit of measurement used to measure the length. Possible values: - `CM` - `INCH`
    Unit *string `json:"unit,omitempty"`
    // The width of the scheduled package. 
    Width *string `json:"width,omitempty"`
}

// NewFulfillment202309GetPackageDetailResponseDataDimension instantiates a new Fulfillment202309GetPackageDetailResponseDataDimension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillment202309GetPackageDetailResponseDataDimension() *Fulfillment202309GetPackageDetailResponseDataDimension {
    this := Fulfillment202309GetPackageDetailResponseDataDimension{}
    return &this
}

// NewFulfillment202309GetPackageDetailResponseDataDimensionWithDefaults instantiates a new Fulfillment202309GetPackageDetailResponseDataDimension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillment202309GetPackageDetailResponseDataDimensionWithDefaults() *Fulfillment202309GetPackageDetailResponseDataDimension {
    this := Fulfillment202309GetPackageDetailResponseDataDimension{}
    return &this
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *Fulfillment202309GetPackageDetailResponseDataDimension) GetHeight() string {
    if o == nil || utils.IsNil(o.Height) {
        var ret string
        return ret
    }
    return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309GetPackageDetailResponseDataDimension) GetHeightOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Height) {
        return nil, false
    }
    return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *Fulfillment202309GetPackageDetailResponseDataDimension) HasHeight() bool {
    if o != nil && !utils.IsNil(o.Height) {
        return true
    }

    return false
}

// SetHeight gets a reference to the given string and assigns it to the Height field.
func (o *Fulfillment202309GetPackageDetailResponseDataDimension) SetHeight(v string) {
    o.Height = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *Fulfillment202309GetPackageDetailResponseDataDimension) GetLength() string {
    if o == nil || utils.IsNil(o.Length) {
        var ret string
        return ret
    }
    return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309GetPackageDetailResponseDataDimension) GetLengthOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Length) {
        return nil, false
    }
    return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *Fulfillment202309GetPackageDetailResponseDataDimension) HasLength() bool {
    if o != nil && !utils.IsNil(o.Length) {
        return true
    }

    return false
}

// SetLength gets a reference to the given string and assigns it to the Length field.
func (o *Fulfillment202309GetPackageDetailResponseDataDimension) SetLength(v string) {
    o.Length = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *Fulfillment202309GetPackageDetailResponseDataDimension) GetUnit() string {
    if o == nil || utils.IsNil(o.Unit) {
        var ret string
        return ret
    }
    return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309GetPackageDetailResponseDataDimension) GetUnitOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Unit) {
        return nil, false
    }
    return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *Fulfillment202309GetPackageDetailResponseDataDimension) HasUnit() bool {
    if o != nil && !utils.IsNil(o.Unit) {
        return true
    }

    return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *Fulfillment202309GetPackageDetailResponseDataDimension) SetUnit(v string) {
    o.Unit = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *Fulfillment202309GetPackageDetailResponseDataDimension) GetWidth() string {
    if o == nil || utils.IsNil(o.Width) {
        var ret string
        return ret
    }
    return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fulfillment202309GetPackageDetailResponseDataDimension) GetWidthOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Width) {
        return nil, false
    }
    return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *Fulfillment202309GetPackageDetailResponseDataDimension) HasWidth() bool {
    if o != nil && !utils.IsNil(o.Width) {
        return true
    }

    return false
}

// SetWidth gets a reference to the given string and assigns it to the Width field.
func (o *Fulfillment202309GetPackageDetailResponseDataDimension) SetWidth(v string) {
    o.Width = &v
}

func (o Fulfillment202309GetPackageDetailResponseDataDimension) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Fulfillment202309GetPackageDetailResponseDataDimension) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Height) {
        toSerialize["height"] = o.Height
    }
    if !utils.IsNil(o.Length) {
        toSerialize["length"] = o.Length
    }
    if !utils.IsNil(o.Unit) {
        toSerialize["unit"] = o.Unit
    }
    if !utils.IsNil(o.Width) {
        toSerialize["width"] = o.Width
    }
    return toSerialize, nil
}

type NullableFulfillment202309GetPackageDetailResponseDataDimension struct {
	value *Fulfillment202309GetPackageDetailResponseDataDimension
	isSet bool
}

func (v NullableFulfillment202309GetPackageDetailResponseDataDimension) Get() *Fulfillment202309GetPackageDetailResponseDataDimension {
	return v.value
}

func (v *NullableFulfillment202309GetPackageDetailResponseDataDimension) Set(val *Fulfillment202309GetPackageDetailResponseDataDimension) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillment202309GetPackageDetailResponseDataDimension) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillment202309GetPackageDetailResponseDataDimension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillment202309GetPackageDetailResponseDataDimension(val *Fulfillment202309GetPackageDetailResponseDataDimension) *NullableFulfillment202309GetPackageDetailResponseDataDimension {
	return &NullableFulfillment202309GetPackageDetailResponseDataDimension{value: val, isSet: true}
}

func (v NullableFulfillment202309GetPackageDetailResponseDataDimension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFulfillment202309GetPackageDetailResponseDataDimension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


