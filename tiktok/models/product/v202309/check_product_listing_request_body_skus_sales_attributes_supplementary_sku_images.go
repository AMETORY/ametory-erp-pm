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

            // checks if the Product202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages{}

// Product202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages struct for Product202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages
type Product202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages struct {
    // The URI of the image. Obtain this URI by uploading the images through the [Upload Product Image API](6509df95defece02be598a22)  with `use_case=ATTRIBUTE_IMAGE`. 
    Uri *string `json:"uri,omitempty"`
}

// NewProduct202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages instantiates a new Product202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages() *Product202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages {
    this := Product202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages{}
    return &this
}

// NewProduct202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImagesWithDefaults instantiates a new Product202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImagesWithDefaults() *Product202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages {
    this := Product202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages{}
    return &this
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *Product202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages) GetUri() string {
    if o == nil || utils.IsNil(o.Uri) {
        var ret string
        return ret
    }
    return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages) GetUriOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Uri) {
        return nil, false
    }
    return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *Product202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages) HasUri() bool {
    if o != nil && !utils.IsNil(o.Uri) {
        return true
    }

    return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *Product202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages) SetUri(v string) {
    o.Uri = &v
}

func (o Product202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Uri) {
        toSerialize["uri"] = o.Uri
    }
    return toSerialize, nil
}

type NullableProduct202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages struct {
	value *Product202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages
	isSet bool
}

func (v NullableProduct202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages) Get() *Product202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages {
	return v.value
}

func (v *NullableProduct202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages) Set(val *Product202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages(val *Product202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages) *NullableProduct202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages {
	return &NullableProduct202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages{value: val, isSet: true}
}

func (v NullableProduct202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309CheckProductListingRequestBodySkusSalesAttributesSupplementarySkuImages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


