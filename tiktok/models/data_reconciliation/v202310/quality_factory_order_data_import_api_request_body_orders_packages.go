/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package data_reconciliation_v202310

import (
    "encoding/json"
    "tiktokshop/open/sdk_golang/utils"
)

            // checks if the DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages{}

// DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages struct for DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages
type DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages struct {
    // The tracking corresponding Tiktok shop package id 
    PackageId *string `json:"package_id,omitempty"`
    // The provider name of tracking info
    ShippingProviderName *string `json:"shipping_provider_name,omitempty"`
    // Tracking number of tracking info 
    TrackingNumber *string `json:"tracking_number,omitempty"`
}

// NewDataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages instantiates a new DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages() *DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages {
    this := DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages{}
    return &this
}

// NewDataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackagesWithDefaults instantiates a new DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackagesWithDefaults() *DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages {
    this := DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages{}
    return &this
}

// GetPackageId returns the PackageId field value if set, zero value otherwise.
func (o *DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages) GetPackageId() string {
    if o == nil || utils.IsNil(o.PackageId) {
        var ret string
        return ret
    }
    return *o.PackageId
}

// GetPackageIdOk returns a tuple with the PackageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages) GetPackageIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.PackageId) {
        return nil, false
    }
    return o.PackageId, true
}

// HasPackageId returns a boolean if a field has been set.
func (o *DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages) HasPackageId() bool {
    if o != nil && !utils.IsNil(o.PackageId) {
        return true
    }

    return false
}

// SetPackageId gets a reference to the given string and assigns it to the PackageId field.
func (o *DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages) SetPackageId(v string) {
    o.PackageId = &v
}

// GetShippingProviderName returns the ShippingProviderName field value if set, zero value otherwise.
func (o *DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages) GetShippingProviderName() string {
    if o == nil || utils.IsNil(o.ShippingProviderName) {
        var ret string
        return ret
    }
    return *o.ShippingProviderName
}

// GetShippingProviderNameOk returns a tuple with the ShippingProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages) GetShippingProviderNameOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ShippingProviderName) {
        return nil, false
    }
    return o.ShippingProviderName, true
}

// HasShippingProviderName returns a boolean if a field has been set.
func (o *DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages) HasShippingProviderName() bool {
    if o != nil && !utils.IsNil(o.ShippingProviderName) {
        return true
    }

    return false
}

// SetShippingProviderName gets a reference to the given string and assigns it to the ShippingProviderName field.
func (o *DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages) SetShippingProviderName(v string) {
    o.ShippingProviderName = &v
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise.
func (o *DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages) GetTrackingNumber() string {
    if o == nil || utils.IsNil(o.TrackingNumber) {
        var ret string
        return ret
    }
    return *o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages) GetTrackingNumberOk() (*string, bool) {
    if o == nil || utils.IsNil(o.TrackingNumber) {
        return nil, false
    }
    return o.TrackingNumber, true
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages) HasTrackingNumber() bool {
    if o != nil && !utils.IsNil(o.TrackingNumber) {
        return true
    }

    return false
}

// SetTrackingNumber gets a reference to the given string and assigns it to the TrackingNumber field.
func (o *DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages) SetTrackingNumber(v string) {
    o.TrackingNumber = &v
}

func (o DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.PackageId) {
        toSerialize["package_id"] = o.PackageId
    }
    if !utils.IsNil(o.ShippingProviderName) {
        toSerialize["shipping_provider_name"] = o.ShippingProviderName
    }
    if !utils.IsNil(o.TrackingNumber) {
        toSerialize["tracking_number"] = o.TrackingNumber
    }
    return toSerialize, nil
}

type NullableDataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages struct {
	value *DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages
	isSet bool
}

func (v NullableDataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages) Get() *DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages {
	return v.value
}

func (v *NullableDataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages) Set(val *DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages) {
	v.value = val
	v.isSet = true
}

func (v NullableDataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages) IsSet() bool {
	return v.isSet
}

func (v *NullableDataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages(val *DataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages) *NullableDataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages {
	return &NullableDataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages{value: val, isSet: true}
}

func (v NullableDataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataReconciliation202310QualityFactoryOrderDataImportAPIRequestBodyOrdersPackages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


