/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product_v202309

import (
    "encoding/json"
    "tiktokshop/open/sdk_golang/utils"
)

            // checks if the Product202309GetGlobalProductResponseDataMainImages type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309GetGlobalProductResponseDataMainImages{}

// Product202309GetGlobalProductResponseDataMainImages struct for Product202309GetGlobalProductResponseDataMainImages
type Product202309GetGlobalProductResponseDataMainImages struct {
    // The image height. Unit: px
    Height *int32 `json:"height,omitempty"`
    // The URI of the image.
    Uri *string `json:"uri,omitempty"`
    // The image width. Unit: px
    Width *int32 `json:"width,omitempty"`
}

// NewProduct202309GetGlobalProductResponseDataMainImages instantiates a new Product202309GetGlobalProductResponseDataMainImages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309GetGlobalProductResponseDataMainImages() *Product202309GetGlobalProductResponseDataMainImages {
    this := Product202309GetGlobalProductResponseDataMainImages{}
    return &this
}

// NewProduct202309GetGlobalProductResponseDataMainImagesWithDefaults instantiates a new Product202309GetGlobalProductResponseDataMainImages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309GetGlobalProductResponseDataMainImagesWithDefaults() *Product202309GetGlobalProductResponseDataMainImages {
    this := Product202309GetGlobalProductResponseDataMainImages{}
    return &this
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *Product202309GetGlobalProductResponseDataMainImages) GetHeight() int32 {
    if o == nil || utils.IsNil(o.Height) {
        var ret int32
        return ret
    }
    return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309GetGlobalProductResponseDataMainImages) GetHeightOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Height) {
        return nil, false
    }
    return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *Product202309GetGlobalProductResponseDataMainImages) HasHeight() bool {
    if o != nil && !utils.IsNil(o.Height) {
        return true
    }

    return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *Product202309GetGlobalProductResponseDataMainImages) SetHeight(v int32) {
    o.Height = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *Product202309GetGlobalProductResponseDataMainImages) GetUri() string {
    if o == nil || utils.IsNil(o.Uri) {
        var ret string
        return ret
    }
    return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309GetGlobalProductResponseDataMainImages) GetUriOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Uri) {
        return nil, false
    }
    return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *Product202309GetGlobalProductResponseDataMainImages) HasUri() bool {
    if o != nil && !utils.IsNil(o.Uri) {
        return true
    }

    return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *Product202309GetGlobalProductResponseDataMainImages) SetUri(v string) {
    o.Uri = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *Product202309GetGlobalProductResponseDataMainImages) GetWidth() int32 {
    if o == nil || utils.IsNil(o.Width) {
        var ret int32
        return ret
    }
    return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309GetGlobalProductResponseDataMainImages) GetWidthOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Width) {
        return nil, false
    }
    return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *Product202309GetGlobalProductResponseDataMainImages) HasWidth() bool {
    if o != nil && !utils.IsNil(o.Width) {
        return true
    }

    return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *Product202309GetGlobalProductResponseDataMainImages) SetWidth(v int32) {
    o.Width = &v
}

func (o Product202309GetGlobalProductResponseDataMainImages) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309GetGlobalProductResponseDataMainImages) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Height) {
        toSerialize["height"] = o.Height
    }
    if !utils.IsNil(o.Uri) {
        toSerialize["uri"] = o.Uri
    }
    if !utils.IsNil(o.Width) {
        toSerialize["width"] = o.Width
    }
    return toSerialize, nil
}

type NullableProduct202309GetGlobalProductResponseDataMainImages struct {
	value *Product202309GetGlobalProductResponseDataMainImages
	isSet bool
}

func (v NullableProduct202309GetGlobalProductResponseDataMainImages) Get() *Product202309GetGlobalProductResponseDataMainImages {
	return v.value
}

func (v *NullableProduct202309GetGlobalProductResponseDataMainImages) Set(val *Product202309GetGlobalProductResponseDataMainImages) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309GetGlobalProductResponseDataMainImages) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309GetGlobalProductResponseDataMainImages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309GetGlobalProductResponseDataMainImages(val *Product202309GetGlobalProductResponseDataMainImages) *NullableProduct202309GetGlobalProductResponseDataMainImages {
	return &NullableProduct202309GetGlobalProductResponseDataMainImages{value: val, isSet: true}
}

func (v NullableProduct202309GetGlobalProductResponseDataMainImages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309GetGlobalProductResponseDataMainImages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


