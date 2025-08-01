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

            // checks if the Product202309EditProductRequestBodyCertifications type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309EditProductRequestBodyCertifications{}

// Product202309EditProductRequestBodyCertifications struct for Product202309EditProductRequestBodyCertifications
type Product202309EditProductRequestBodyCertifications struct {
    // The expiration date of this certification expressed in unix timestamp (seconds) UTC+0. This field may be required for certain certifications. Use the [Get Category Rules API](https://partner.tiktokshop.com/docv2/page/6509c0febace3e02b74594a9) to find out the requirements.
    ExpirationDate *int64 `json:"expiration_date,omitempty"`
    // A list of certification related files.
    Files []Product202309EditProductRequestBodyCertificationsFiles `json:"files,omitempty"`
    // The ID to identify the type of certification required for the product category. Retrieve this value from the [Get Category Rules API](https://partner.tiktokshop.com/docv2/page/6509c0febace3e02b74594a9). 
    Id *string `json:"id,omitempty"`
    // A list of certification related images.
    Images []Product202309EditProductRequestBodyCertificationsImages `json:"images,omitempty"`
}

// NewProduct202309EditProductRequestBodyCertifications instantiates a new Product202309EditProductRequestBodyCertifications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309EditProductRequestBodyCertifications() *Product202309EditProductRequestBodyCertifications {
    this := Product202309EditProductRequestBodyCertifications{}
    return &this
}

// NewProduct202309EditProductRequestBodyCertificationsWithDefaults instantiates a new Product202309EditProductRequestBodyCertifications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309EditProductRequestBodyCertificationsWithDefaults() *Product202309EditProductRequestBodyCertifications {
    this := Product202309EditProductRequestBodyCertifications{}
    return &this
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *Product202309EditProductRequestBodyCertifications) GetExpirationDate() int64 {
    if o == nil || utils.IsNil(o.ExpirationDate) {
        var ret int64
        return ret
    }
    return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309EditProductRequestBodyCertifications) GetExpirationDateOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.ExpirationDate) {
        return nil, false
    }
    return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *Product202309EditProductRequestBodyCertifications) HasExpirationDate() bool {
    if o != nil && !utils.IsNil(o.ExpirationDate) {
        return true
    }

    return false
}

// SetExpirationDate gets a reference to the given int64 and assigns it to the ExpirationDate field.
func (o *Product202309EditProductRequestBodyCertifications) SetExpirationDate(v int64) {
    o.ExpirationDate = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *Product202309EditProductRequestBodyCertifications) GetFiles() []Product202309EditProductRequestBodyCertificationsFiles {
    if o == nil || utils.IsNil(o.Files) {
        var ret []Product202309EditProductRequestBodyCertificationsFiles
        return ret
    }
    return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309EditProductRequestBodyCertifications) GetFilesOk() ([]Product202309EditProductRequestBodyCertificationsFiles, bool) {
    if o == nil || utils.IsNil(o.Files) {
        return nil, false
    }
    return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *Product202309EditProductRequestBodyCertifications) HasFiles() bool {
    if o != nil && !utils.IsNil(o.Files) {
        return true
    }

    return false
}

// SetFiles gets a reference to the given []Product202309EditProductRequestBodyCertificationsFiles and assigns it to the Files field.
func (o *Product202309EditProductRequestBodyCertifications) SetFiles(v []Product202309EditProductRequestBodyCertificationsFiles) {
    o.Files = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Product202309EditProductRequestBodyCertifications) GetId() string {
    if o == nil || utils.IsNil(o.Id) {
        var ret string
        return ret
    }
    return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309EditProductRequestBodyCertifications) GetIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Id) {
        return nil, false
    }
    return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Product202309EditProductRequestBodyCertifications) HasId() bool {
    if o != nil && !utils.IsNil(o.Id) {
        return true
    }

    return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Product202309EditProductRequestBodyCertifications) SetId(v string) {
    o.Id = &v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *Product202309EditProductRequestBodyCertifications) GetImages() []Product202309EditProductRequestBodyCertificationsImages {
    if o == nil || utils.IsNil(o.Images) {
        var ret []Product202309EditProductRequestBodyCertificationsImages
        return ret
    }
    return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309EditProductRequestBodyCertifications) GetImagesOk() ([]Product202309EditProductRequestBodyCertificationsImages, bool) {
    if o == nil || utils.IsNil(o.Images) {
        return nil, false
    }
    return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *Product202309EditProductRequestBodyCertifications) HasImages() bool {
    if o != nil && !utils.IsNil(o.Images) {
        return true
    }

    return false
}

// SetImages gets a reference to the given []Product202309EditProductRequestBodyCertificationsImages and assigns it to the Images field.
func (o *Product202309EditProductRequestBodyCertifications) SetImages(v []Product202309EditProductRequestBodyCertificationsImages) {
    o.Images = v
}

func (o Product202309EditProductRequestBodyCertifications) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309EditProductRequestBodyCertifications) ToMap() (map[string]interface{}, error) {
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

type NullableProduct202309EditProductRequestBodyCertifications struct {
	value *Product202309EditProductRequestBodyCertifications
	isSet bool
}

func (v NullableProduct202309EditProductRequestBodyCertifications) Get() *Product202309EditProductRequestBodyCertifications {
	return v.value
}

func (v *NullableProduct202309EditProductRequestBodyCertifications) Set(val *Product202309EditProductRequestBodyCertifications) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309EditProductRequestBodyCertifications) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309EditProductRequestBodyCertifications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309EditProductRequestBodyCertifications(val *Product202309EditProductRequestBodyCertifications) *NullableProduct202309EditProductRequestBodyCertifications {
	return &NullableProduct202309EditProductRequestBodyCertifications{value: val, isSet: true}
}

func (v NullableProduct202309EditProductRequestBodyCertifications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309EditProductRequestBodyCertifications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


