/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package return_refund_v202309

import (
    "encoding/json"
    "tiktokshop/open/sdk_golang/utils"
)

            // checks if the ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage{}

// ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage struct for ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage
type ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage struct {
    // Product image height. Unit: px
    Height *int32 `json:"height,omitempty"`
    // Product image URL.
    Url *string `json:"url,omitempty"`
    // Product image width. Unit: px
    Width *int32 `json:"width,omitempty"`
}

// NewReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage instantiates a new ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage() *ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage {
    this := ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage{}
    return &this
}

// NewReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImageWithDefaults instantiates a new ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImageWithDefaults() *ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage {
    this := ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage{}
    return &this
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage) GetHeight() int32 {
    if o == nil || utils.IsNil(o.Height) {
        var ret int32
        return ret
    }
    return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage) GetHeightOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Height) {
        return nil, false
    }
    return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage) HasHeight() bool {
    if o != nil && !utils.IsNil(o.Height) {
        return true
    }

    return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage) SetHeight(v int32) {
    o.Height = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage) GetUrl() string {
    if o == nil || utils.IsNil(o.Url) {
        var ret string
        return ret
    }
    return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage) GetUrlOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Url) {
        return nil, false
    }
    return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage) HasUrl() bool {
    if o != nil && !utils.IsNil(o.Url) {
        return true
    }

    return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage) SetUrl(v string) {
    o.Url = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage) GetWidth() int32 {
    if o == nil || utils.IsNil(o.Width) {
        var ret int32
        return ret
    }
    return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage) GetWidthOk() (*int32, bool) {
    if o == nil || utils.IsNil(o.Width) {
        return nil, false
    }
    return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage) HasWidth() bool {
    if o != nil && !utils.IsNil(o.Width) {
        return true
    }

    return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage) SetWidth(v int32) {
    o.Width = &v
}

func (o ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Height) {
        toSerialize["height"] = o.Height
    }
    if !utils.IsNil(o.Url) {
        toSerialize["url"] = o.Url
    }
    if !utils.IsNil(o.Width) {
        toSerialize["width"] = o.Width
    }
    return toSerialize, nil
}

type NullableReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage struct {
	value *ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage
	isSet bool
}

func (v NullableReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage) Get() *ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage {
	return v.value
}

func (v *NullableReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage) Set(val *ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage(val *ReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage) *NullableReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage {
	return &NullableReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage{value: val, isSet: true}
}

func (v NullableReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnRefund202309SearchCancellationsResponseDataCancellationsCancelLineItemsProductImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


