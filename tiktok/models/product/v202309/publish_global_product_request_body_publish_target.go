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

            // checks if the Product202309PublishGlobalProductRequestBodyPublishTarget type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202309PublishGlobalProductRequestBodyPublishTarget{}

// Product202309PublishGlobalProductRequestBodyPublishTarget struct for Product202309PublishGlobalProductRequestBodyPublishTarget
type Product202309PublishGlobalProductRequestBodyPublishTarget struct {
    // A comma-delimited list of manufacturer IDs. Retrieve the IDs from the [Search Manufacturers API](67066a580dcee902fa03ccf9). Default: The IDs provided when the global product was created. **Note**: Applicable only for the EU market in certain categories. Use the [Get Global Category Rules API](650a056df1fd3102b91b5b8e) to check the requirements.
    ManufacturerIds []string `json:"manufacturer_ids,omitempty"`
    // The new market where you want to publish the global product. Possible values: - DE: Germany - ES: Spain - FR: France - GB: United Kingdom - ID: Indonesia - IE: Ireland - IT: Italy - JP: Japan - MY: Malaysia - PH: Philippines - SG: Singapore - TH: Thailand - US: United States - VN: Vietnam  **Note**: You can only publish in each market once.
    Region *string `json:"region,omitempty"`
    // A comma-delimited list of responsible person IDs. Retrieve the IDs from the [Search Responsible Persons API](67066a55f17b7d02f95d2fb1). Default: The IDs provided when the global product was created. **Note**: Applicable only for the EU market in certain categories. Use the [Get Global Category Rules API](650a056df1fd3102b91b5b8e) to check the requirements.
    ResponsiblePersonIds []string `json:"responsible_person_ids,omitempty"`
    // The SKUs to be published in the specified market. - Max SKUs for EU, JP, UK, US: 300 - Max SKUs for other regions: 100
    Skus []Product202309PublishGlobalProductRequestBodyPublishTargetSkus `json:"skus,omitempty"`
}

// NewProduct202309PublishGlobalProductRequestBodyPublishTarget instantiates a new Product202309PublishGlobalProductRequestBodyPublishTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202309PublishGlobalProductRequestBodyPublishTarget() *Product202309PublishGlobalProductRequestBodyPublishTarget {
    this := Product202309PublishGlobalProductRequestBodyPublishTarget{}
    return &this
}

// NewProduct202309PublishGlobalProductRequestBodyPublishTargetWithDefaults instantiates a new Product202309PublishGlobalProductRequestBodyPublishTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202309PublishGlobalProductRequestBodyPublishTargetWithDefaults() *Product202309PublishGlobalProductRequestBodyPublishTarget {
    this := Product202309PublishGlobalProductRequestBodyPublishTarget{}
    return &this
}

// GetManufacturerIds returns the ManufacturerIds field value if set, zero value otherwise.
func (o *Product202309PublishGlobalProductRequestBodyPublishTarget) GetManufacturerIds() []string {
    if o == nil || utils.IsNil(o.ManufacturerIds) {
        var ret []string
        return ret
    }
    return o.ManufacturerIds
}

// GetManufacturerIdsOk returns a tuple with the ManufacturerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309PublishGlobalProductRequestBodyPublishTarget) GetManufacturerIdsOk() ([]string, bool) {
    if o == nil || utils.IsNil(o.ManufacturerIds) {
        return nil, false
    }
    return o.ManufacturerIds, true
}

// HasManufacturerIds returns a boolean if a field has been set.
func (o *Product202309PublishGlobalProductRequestBodyPublishTarget) HasManufacturerIds() bool {
    if o != nil && !utils.IsNil(o.ManufacturerIds) {
        return true
    }

    return false
}

// SetManufacturerIds gets a reference to the given []string and assigns it to the ManufacturerIds field.
func (o *Product202309PublishGlobalProductRequestBodyPublishTarget) SetManufacturerIds(v []string) {
    o.ManufacturerIds = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *Product202309PublishGlobalProductRequestBodyPublishTarget) GetRegion() string {
    if o == nil || utils.IsNil(o.Region) {
        var ret string
        return ret
    }
    return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309PublishGlobalProductRequestBodyPublishTarget) GetRegionOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Region) {
        return nil, false
    }
    return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *Product202309PublishGlobalProductRequestBodyPublishTarget) HasRegion() bool {
    if o != nil && !utils.IsNil(o.Region) {
        return true
    }

    return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *Product202309PublishGlobalProductRequestBodyPublishTarget) SetRegion(v string) {
    o.Region = &v
}

// GetResponsiblePersonIds returns the ResponsiblePersonIds field value if set, zero value otherwise.
func (o *Product202309PublishGlobalProductRequestBodyPublishTarget) GetResponsiblePersonIds() []string {
    if o == nil || utils.IsNil(o.ResponsiblePersonIds) {
        var ret []string
        return ret
    }
    return o.ResponsiblePersonIds
}

// GetResponsiblePersonIdsOk returns a tuple with the ResponsiblePersonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309PublishGlobalProductRequestBodyPublishTarget) GetResponsiblePersonIdsOk() ([]string, bool) {
    if o == nil || utils.IsNil(o.ResponsiblePersonIds) {
        return nil, false
    }
    return o.ResponsiblePersonIds, true
}

// HasResponsiblePersonIds returns a boolean if a field has been set.
func (o *Product202309PublishGlobalProductRequestBodyPublishTarget) HasResponsiblePersonIds() bool {
    if o != nil && !utils.IsNil(o.ResponsiblePersonIds) {
        return true
    }

    return false
}

// SetResponsiblePersonIds gets a reference to the given []string and assigns it to the ResponsiblePersonIds field.
func (o *Product202309PublishGlobalProductRequestBodyPublishTarget) SetResponsiblePersonIds(v []string) {
    o.ResponsiblePersonIds = v
}

// GetSkus returns the Skus field value if set, zero value otherwise.
func (o *Product202309PublishGlobalProductRequestBodyPublishTarget) GetSkus() []Product202309PublishGlobalProductRequestBodyPublishTargetSkus {
    if o == nil || utils.IsNil(o.Skus) {
        var ret []Product202309PublishGlobalProductRequestBodyPublishTargetSkus
        return ret
    }
    return o.Skus
}

// GetSkusOk returns a tuple with the Skus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202309PublishGlobalProductRequestBodyPublishTarget) GetSkusOk() ([]Product202309PublishGlobalProductRequestBodyPublishTargetSkus, bool) {
    if o == nil || utils.IsNil(o.Skus) {
        return nil, false
    }
    return o.Skus, true
}

// HasSkus returns a boolean if a field has been set.
func (o *Product202309PublishGlobalProductRequestBodyPublishTarget) HasSkus() bool {
    if o != nil && !utils.IsNil(o.Skus) {
        return true
    }

    return false
}

// SetSkus gets a reference to the given []Product202309PublishGlobalProductRequestBodyPublishTargetSkus and assigns it to the Skus field.
func (o *Product202309PublishGlobalProductRequestBodyPublishTarget) SetSkus(v []Product202309PublishGlobalProductRequestBodyPublishTargetSkus) {
    o.Skus = v
}

func (o Product202309PublishGlobalProductRequestBodyPublishTarget) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202309PublishGlobalProductRequestBodyPublishTarget) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.ManufacturerIds) {
        toSerialize["manufacturer_ids"] = o.ManufacturerIds
    }
    if !utils.IsNil(o.Region) {
        toSerialize["region"] = o.Region
    }
    if !utils.IsNil(o.ResponsiblePersonIds) {
        toSerialize["responsible_person_ids"] = o.ResponsiblePersonIds
    }
    if !utils.IsNil(o.Skus) {
        toSerialize["skus"] = o.Skus
    }
    return toSerialize, nil
}

type NullableProduct202309PublishGlobalProductRequestBodyPublishTarget struct {
	value *Product202309PublishGlobalProductRequestBodyPublishTarget
	isSet bool
}

func (v NullableProduct202309PublishGlobalProductRequestBodyPublishTarget) Get() *Product202309PublishGlobalProductRequestBodyPublishTarget {
	return v.value
}

func (v *NullableProduct202309PublishGlobalProductRequestBodyPublishTarget) Set(val *Product202309PublishGlobalProductRequestBodyPublishTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202309PublishGlobalProductRequestBodyPublishTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202309PublishGlobalProductRequestBodyPublishTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202309PublishGlobalProductRequestBodyPublishTarget(val *Product202309PublishGlobalProductRequestBodyPublishTarget) *NullableProduct202309PublishGlobalProductRequestBodyPublishTarget {
	return &NullableProduct202309PublishGlobalProductRequestBodyPublishTarget{value: val, isSet: true}
}

func (v NullableProduct202309PublishGlobalProductRequestBodyPublishTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202309PublishGlobalProductRequestBodyPublishTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


