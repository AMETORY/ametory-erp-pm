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

            // checks if the Product202309CreateProductRequestBodyCertifications type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309CreateProductRequestBodyCertifications{}

// Product202309CreateProductRequestBodyCertifications struct for Product202309CreateProductRequestBodyCertifications
type Product202309CreateProductRequestBodyCertifications struct {
    // The expiration date of this certification expressed in unix timestamp (seconds) UTC+0. This field may be required for certain certifications. Use the [Get Category Rules API](https://partner.tiktokshop.com/docv2/page/6509c0febace3e02b74594a9) to find out the requirements.
    ExpirationDate *int64 `json:"expiration_date,omitempty"`
    // A list of certification related files.
    Files []Product202309CreateProductRequestBodyCertificationsFiles `json:"files,omitempty"`
    // The ID to identify the type of certification required for the product category. Retrieve this value from the [Get Category Rules API](https://partner.tiktokshop.com/docv2/page/6509c0febace3e02b74594a9). 
    Id *string `json:"id,omitempty"`
    // A list of certification related images.
    Images []Product202309CreateProductRequestBodyCertificationsImages `json:"images,omitempty"`
}

// NewProduct202309CreateProductRequestBodyCertifications instantiates a new Product202309CreateProductRequestBodyCertifications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309CreateProductRequestBodyCertifications() *Product202309CreateProductRequestBodyCertifications {
    this := Product202309CreateProductRequestBodyCertifications{}
    return &this
}

// NewProduct202309CreateProductRequestBodyCertificationsWithDefaults instantiates a new Product202309CreateProductRequestBodyCertifications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309CreateProductRequestBodyCertificationsWithDefaults() *Product202309CreateProductRequestBodyCertifications {
    this := Product202309CreateProductRequestBodyCertifications{}
    return &this
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *Product202309CreateProductRequestBodyCertifications) GetExpirationDate() int64 {
    if o == nil || utils.IsNil(o.ExpirationDate) {
        var ret int64
        return ret
    }
    return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309CreateProductRequestBodyCertifications) GetExpirationDateOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.ExpirationDate) {
        return nil, false
    }
    return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *Product202309CreateProductRequestBodyCertifications) HasExpirationDate() bool {
    if o != nil && !utils.IsNil(o.ExpirationDate) {
        return true
    }

    return false
}

// SetExpirationDate gets a reference to the given int64 and assigns it to the ExpirationDate field.
func (o *Product202309CreateProductRequestBodyCertifications) SetExpirationDate(v int64) {
    o.ExpirationDate = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *Product202309CreateProductRequestBodyCertifications) GetFiles() []Product202309CreateProductRequestBodyCertificationsFiles {
    if o == nil || utils.IsNil(o.Files) {
        var ret []Product202309CreateProductRequestBodyCertificationsFiles
        return ret
    }
    return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309CreateProductRequestBodyCertifications) GetFilesOk() ([]Product202309CreateProductRequestBodyCertificationsFiles, bool) {
    if o == nil || utils.IsNil(o.Files) {
        return nil, false
    }
    return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *Product202309CreateProductRequestBodyCertifications) HasFiles() bool {
    if o != nil && !utils.IsNil(o.Files) {
        return true
    }

    return false
}

// SetFiles gets a reference to the given []Product202309CreateProductRequestBodyCertificationsFiles and assigns it to the Files field.
func (o *Product202309CreateProductRequestBodyCertifications) SetFiles(v []Product202309CreateProductRequestBodyCertificationsFiles) {
    o.Files = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Product202309CreateProductRequestBodyCertifications) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309CreateProductRequestBodyCertifications) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Product202309CreateProductRequestBodyCertifications) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Product202309CreateProductRequestBodyCertifications) SetId(v string) {
    o.Id = &v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *Product202309CreateProductRequestBodyCertifications) GetImages() []Product202309CreateProductRequestBodyCertificationsImages {
    if o == nil || utils.IsNil(o.Images) {
        var ret []Product202309CreateProductRequestBodyCertificationsImages
        return ret
    }
    return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309CreateProductRequestBodyCertifications) GetImagesOk() ([]Product202309CreateProductRequestBodyCertificationsImages, bool) {
    if o == nil || utils.IsNil(o.Images) {
        return nil, false
    }
    return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *Product202309CreateProductRequestBodyCertifications) HasImages() bool {
    if o != nil && !utils.IsNil(o.Images) {
        return true
    }

    return false
}

// SetImages gets a reference to the given []Product202309CreateProductRequestBodyCertificationsImages and assigns it to the Images field.
func (o *Product202309CreateProductRequestBodyCertifications) SetImages(v []Product202309CreateProductRequestBodyCertificationsImages) {
    o.Images = v
}

func (o Product202309CreateProductRequestBodyCertifications) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309CreateProductRequestBodyCertifications) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.ExpirationDate) {
        toSerialize["expiration_date"] = o.ExpirationDate
    }
    if !utils.IsNil(o.Files) {
        toSerialize["files"] = o.Files
    }
    if !utils.IsNil(o.Id) {
        toSerialize["id"] = o.Id
    }
    if !utils.IsNil(o.Images) {
        toSerialize["images"] = o.Images
    }
    return toSerialize, nil
}

type NullableProduct202309CreateProductRequestBodyCertifications struct {
	value *Product202309CreateProductRequestBodyCertifications
	isSet bool
}

func (v NullableProduct202309CreateProductRequestBodyCertifications) Get() *Product202309CreateProductRequestBodyCertifications {
	return v.value
}

func (v *NullableProduct202309CreateProductRequestBodyCertifications) Set(val *Product202309CreateProductRequestBodyCertifications) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309CreateProductRequestBodyCertifications) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309CreateProductRequestBodyCertifications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309CreateProductRequestBodyCertifications(val *Product202309CreateProductRequestBodyCertifications) *NullableProduct202309CreateProductRequestBodyCertifications {
	return &NullableProduct202309CreateProductRequestBodyCertifications{value: val, isSet: true}
}

func (v NullableProduct202309CreateProductRequestBodyCertifications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309CreateProductRequestBodyCertifications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


