/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product_v202312

import (
    "encoding/json"
    "tiktokshop/open/sdk_golang/utils"
)

            // checks if the Product202312SearchProductsRequestBody type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202312SearchProductsRequestBody{}

// Product202312SearchProductsRequestBody struct for Product202312SearchProductsRequestBody
type Product202312SearchProductsRequestBody struct {
    // Filter products by their audit status for TikTok Shop. Possible values:  - AUDITING: The product is currently being audited. - FAILED: The product failed the audit, or the audit was cancelled. - APPROVED: The product passed the audit and has been listed on the platform.
    AuditStatus []string `json:"audit_status,omitempty"`
    // Filter products by the category tree version. Possible values based on region: - US: `v2`, represents the 7-level category tree. - Other regions: `v1`, represents the 3-level category tree. Default: Return all products from both `v1` and `v2` category trees.
    CategoryVersion *string `json:"category_version,omitempty"`
    // Filter products to show only those that are created on or after the specified date and time. Unix timestamp.  **Note**: `create_time_ge` and `create_time_le` together constitute the creation time filter condition. - If `create_time_ge` is filled but `create_time_le` is empty, `create_time_le` will default to the current time. - If `create_time_le` is filled but `create_time_ge` is empty, `create_time_ge` will default to the earliest shop time.
    CreateTimeGe *int64 `json:"create_time_ge,omitempty"`
    // Filter products to show only those that are created on or before the specified date and time. Unix timestamp. Refer to notes in `create_time_ge` for more usage information.
    CreateTimeLe *int64 `json:"create_time_le,omitempty"`
    // Filter products by the listing platforms. Possible values: - TOKOPEDIA - TIKTOK_SHOP Default: Return all products regardless of their listing platform.  Applicable only for sellers that migrated from Tokopedia. **Note**: - You must also specify a `status` value other than `ALL` when filtering by listing platforms. Returning all statuses is not supported. - If you pass in one platform, the search will return products that are listed on that platform, including those that are listed on both platforms. - If you pass in `[\"TIKTOK_SHOP\", \"TOKOPEDIA\"]`, only products listed on both platforms will be returned, not those listed on just one.
    ListingPlatforms []string `json:"listing_platforms,omitempty"`
    // Filter products by their listing quality tier. Possible values: - POOR - FAIR - GOOD  **Note**: Available only for the US market.
    ListingQualityTier *string `json:"listing_quality_tier,omitempty"`
    // Filter products by these seller SKU codes.
    SellerSkus []string `json:"seller_skus,omitempty"`
    // Filter products by SKU IDs. Max count: 10
    SkuIds []string `json:"sku_ids,omitempty"`
    // Filter products by their status. Default: ALL Possible values:  - ALL - DRAFT - PENDING - FAILED - ACTIVATE - SELLER_DEACTIVATED - PLATFORM_DEACTIVATED - FREEZE - DELETED 
    Status *string `json:"status,omitempty"`
    // Filter products to show only those that are updated on or after the specified date and time. Unix timestamp.  **Note**: `update_time_ge` and `update_time_le` together define the update time filter condition. - If `update_time_ge` is filled but `update_time_le` is empty, `update_time_le` will default to the current time. - If `update_time_le` is filled but `update_time_ge` is empty, `update_time_ge` will default to the earliest shop time.
    UpdateTimeGe *int64 `json:"update_time_ge,omitempty"`
    // Filter products to show only those that are updated on or before the specified date and time. Unix timestamp. Refer to notes in `update_time_ge` for more usage information.
    UpdateTimeLe *int64 `json:"update_time_le,omitempty"`
}

// NewProduct202312SearchProductsRequestBody instantiates a new Product202312SearchProductsRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202312SearchProductsRequestBody() *Product202312SearchProductsRequestBody {
    this := Product202312SearchProductsRequestBody{}
    return &this
}

// NewProduct202312SearchProductsRequestBodyWithDefaults instantiates a new Product202312SearchProductsRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202312SearchProductsRequestBodyWithDefaults() *Product202312SearchProductsRequestBody {
    this := Product202312SearchProductsRequestBody{}
    return &this
}

// GetAuditStatus returns the AuditStatus field value if set, zero value otherwise.
func (o *Product202312SearchProductsRequestBody) GetAuditStatus() []string {
    if o == nil || utils.IsNil(o.AuditStatus) {
        var ret []string
        return ret
    }
    return o.AuditStatus
}

// GetAuditStatusOk returns a tuple with the AuditStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202312SearchProductsRequestBody) GetAuditStatusOk() ([]string, bool) {
    if o == nil || utils.IsNil(o.AuditStatus) {
        return nil, false
    }
    return o.AuditStatus, true
}

// HasAuditStatus returns a boolean if a field has been set.
func (o *Product202312SearchProductsRequestBody) HasAuditStatus() bool {
    if o != nil && !utils.IsNil(o.AuditStatus) {
        return true
    }

    return false
}

// SetAuditStatus gets a reference to the given []string and assigns it to the AuditStatus field.
func (o *Product202312SearchProductsRequestBody) SetAuditStatus(v []string) {
    o.AuditStatus = v
}

// GetCategoryVersion returns the CategoryVersion field value if set, zero value otherwise.
func (o *Product202312SearchProductsRequestBody) GetCategoryVersion() string {
    if o == nil || utils.IsNil(o.CategoryVersion) {
        var ret string
        return ret
    }
    return *o.CategoryVersion
}

// GetCategoryVersionOk returns a tuple with the CategoryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202312SearchProductsRequestBody) GetCategoryVersionOk() (*string, bool) {
    if o == nil || utils.IsNil(o.CategoryVersion) {
        return nil, false
    }
    return o.CategoryVersion, true
}

// HasCategoryVersion returns a boolean if a field has been set.
func (o *Product202312SearchProductsRequestBody) HasCategoryVersion() bool {
    if o != nil && !utils.IsNil(o.CategoryVersion) {
        return true
    }

    return false
}

// SetCategoryVersion gets a reference to the given string and assigns it to the CategoryVersion field.
func (o *Product202312SearchProductsRequestBody) SetCategoryVersion(v string) {
    o.CategoryVersion = &v
}

// GetCreateTimeGe returns the CreateTimeGe field value if set, zero value otherwise.
func (o *Product202312SearchProductsRequestBody) GetCreateTimeGe() int64 {
    if o == nil || utils.IsNil(o.CreateTimeGe) {
        var ret int64
        return ret
    }
    return *o.CreateTimeGe
}

// GetCreateTimeGeOk returns a tuple with the CreateTimeGe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202312SearchProductsRequestBody) GetCreateTimeGeOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.CreateTimeGe) {
        return nil, false
    }
    return o.CreateTimeGe, true
}

// HasCreateTimeGe returns a boolean if a field has been set.
func (o *Product202312SearchProductsRequestBody) HasCreateTimeGe() bool {
    if o != nil && !utils.IsNil(o.CreateTimeGe) {
        return true
    }

    return false
}

// SetCreateTimeGe gets a reference to the given int64 and assigns it to the CreateTimeGe field.
func (o *Product202312SearchProductsRequestBody) SetCreateTimeGe(v int64) {
    o.CreateTimeGe = &v
}

// GetCreateTimeLe returns the CreateTimeLe field value if set, zero value otherwise.
func (o *Product202312SearchProductsRequestBody) GetCreateTimeLe() int64 {
    if o == nil || utils.IsNil(o.CreateTimeLe) {
        var ret int64
        return ret
    }
    return *o.CreateTimeLe
}

// GetCreateTimeLeOk returns a tuple with the CreateTimeLe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202312SearchProductsRequestBody) GetCreateTimeLeOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.CreateTimeLe) {
        return nil, false
    }
    return o.CreateTimeLe, true
}

// HasCreateTimeLe returns a boolean if a field has been set.
func (o *Product202312SearchProductsRequestBody) HasCreateTimeLe() bool {
    if o != nil && !utils.IsNil(o.CreateTimeLe) {
        return true
    }

    return false
}

// SetCreateTimeLe gets a reference to the given int64 and assigns it to the CreateTimeLe field.
func (o *Product202312SearchProductsRequestBody) SetCreateTimeLe(v int64) {
    o.CreateTimeLe = &v
}

// GetListingPlatforms returns the ListingPlatforms field value if set, zero value otherwise.
func (o *Product202312SearchProductsRequestBody) GetListingPlatforms() []string {
    if o == nil || utils.IsNil(o.ListingPlatforms) {
        var ret []string
        return ret
    }
    return o.ListingPlatforms
}

// GetListingPlatformsOk returns a tuple with the ListingPlatforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202312SearchProductsRequestBody) GetListingPlatformsOk() ([]string, bool) {
    if o == nil || utils.IsNil(o.ListingPlatforms) {
        return nil, false
    }
    return o.ListingPlatforms, true
}

// HasListingPlatforms returns a boolean if a field has been set.
func (o *Product202312SearchProductsRequestBody) HasListingPlatforms() bool {
    if o != nil && !utils.IsNil(o.ListingPlatforms) {
        return true
    }

    return false
}

// SetListingPlatforms gets a reference to the given []string and assigns it to the ListingPlatforms field.
func (o *Product202312SearchProductsRequestBody) SetListingPlatforms(v []string) {
    o.ListingPlatforms = v
}

// GetListingQualityTier returns the ListingQualityTier field value if set, zero value otherwise.
func (o *Product202312SearchProductsRequestBody) GetListingQualityTier() string {
    if o == nil || utils.IsNil(o.ListingQualityTier) {
        var ret string
        return ret
    }
    return *o.ListingQualityTier
}

// GetListingQualityTierOk returns a tuple with the ListingQualityTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202312SearchProductsRequestBody) GetListingQualityTierOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ListingQualityTier) {
        return nil, false
    }
    return o.ListingQualityTier, true
}

// HasListingQualityTier returns a boolean if a field has been set.
func (o *Product202312SearchProductsRequestBody) HasListingQualityTier() bool {
    if o != nil && !utils.IsNil(o.ListingQualityTier) {
        return true
    }

    return false
}

// SetListingQualityTier gets a reference to the given string and assigns it to the ListingQualityTier field.
func (o *Product202312SearchProductsRequestBody) SetListingQualityTier(v string) {
    o.ListingQualityTier = &v
}

// GetSellerSkus returns the SellerSkus field value if set, zero value otherwise.
func (o *Product202312SearchProductsRequestBody) GetSellerSkus() []string {
    if o == nil || utils.IsNil(o.SellerSkus) {
        var ret []string
        return ret
    }
    return o.SellerSkus
}

// GetSellerSkusOk returns a tuple with the SellerSkus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202312SearchProductsRequestBody) GetSellerSkusOk() ([]string, bool) {
    if o == nil || utils.IsNil(o.SellerSkus) {
        return nil, false
    }
    return o.SellerSkus, true
}

// HasSellerSkus returns a boolean if a field has been set.
func (o *Product202312SearchProductsRequestBody) HasSellerSkus() bool {
    if o != nil && !utils.IsNil(o.SellerSkus) {
        return true
    }

    return false
}

// SetSellerSkus gets a reference to the given []string and assigns it to the SellerSkus field.
func (o *Product202312SearchProductsRequestBody) SetSellerSkus(v []string) {
    o.SellerSkus = v
}

// GetSkuIds returns the SkuIds field value if set, zero value otherwise.
func (o *Product202312SearchProductsRequestBody) GetSkuIds() []string {
    if o == nil || utils.IsNil(o.SkuIds) {
        var ret []string
        return ret
    }
    return o.SkuIds
}

// GetSkuIdsOk returns a tuple with the SkuIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202312SearchProductsRequestBody) GetSkuIdsOk() ([]string, bool) {
    if o == nil || utils.IsNil(o.SkuIds) {
        return nil, false
    }
    return o.SkuIds, true
}

// HasSkuIds returns a boolean if a field has been set.
func (o *Product202312SearchProductsRequestBody) HasSkuIds() bool {
    if o != nil && !utils.IsNil(o.SkuIds) {
        return true
    }

    return false
}

// SetSkuIds gets a reference to the given []string and assigns it to the SkuIds field.
func (o *Product202312SearchProductsRequestBody) SetSkuIds(v []string) {
    o.SkuIds = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Product202312SearchProductsRequestBody) GetStatus() string {
    if o == nil || utils.IsNil(o.Status) {
        var ret string
        return ret
    }
    return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202312SearchProductsRequestBody) GetStatusOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Status) {
        return nil, false
    }
    return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Product202312SearchProductsRequestBody) HasStatus() bool {
    if o != nil && !utils.IsNil(o.Status) {
        return true
    }

    return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Product202312SearchProductsRequestBody) SetStatus(v string) {
    o.Status = &v
}

// GetUpdateTimeGe returns the UpdateTimeGe field value if set, zero value otherwise.
func (o *Product202312SearchProductsRequestBody) GetUpdateTimeGe() int64 {
    if o == nil || utils.IsNil(o.UpdateTimeGe) {
        var ret int64
        return ret
    }
    return *o.UpdateTimeGe
}

// GetUpdateTimeGeOk returns a tuple with the UpdateTimeGe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202312SearchProductsRequestBody) GetUpdateTimeGeOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.UpdateTimeGe) {
        return nil, false
    }
    return o.UpdateTimeGe, true
}

// HasUpdateTimeGe returns a boolean if a field has been set.
func (o *Product202312SearchProductsRequestBody) HasUpdateTimeGe() bool {
    if o != nil && !utils.IsNil(o.UpdateTimeGe) {
        return true
    }

    return false
}

// SetUpdateTimeGe gets a reference to the given int64 and assigns it to the UpdateTimeGe field.
func (o *Product202312SearchProductsRequestBody) SetUpdateTimeGe(v int64) {
    o.UpdateTimeGe = &v
}

// GetUpdateTimeLe returns the UpdateTimeLe field value if set, zero value otherwise.
func (o *Product202312SearchProductsRequestBody) GetUpdateTimeLe() int64 {
    if o == nil || utils.IsNil(o.UpdateTimeLe) {
        var ret int64
        return ret
    }
    return *o.UpdateTimeLe
}

// GetUpdateTimeLeOk returns a tuple with the UpdateTimeLe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202312SearchProductsRequestBody) GetUpdateTimeLeOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.UpdateTimeLe) {
        return nil, false
    }
    return o.UpdateTimeLe, true
}

// HasUpdateTimeLe returns a boolean if a field has been set.
func (o *Product202312SearchProductsRequestBody) HasUpdateTimeLe() bool {
    if o != nil && !utils.IsNil(o.UpdateTimeLe) {
        return true
    }

    return false
}

// SetUpdateTimeLe gets a reference to the given int64 and assigns it to the UpdateTimeLe field.
func (o *Product202312SearchProductsRequestBody) SetUpdateTimeLe(v int64) {
    o.UpdateTimeLe = &v
}

func (o Product202312SearchProductsRequestBody) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202312SearchProductsRequestBody) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.AuditStatus) {
        toSerialize["audit_status"] = o.AuditStatus
    }
    if !utils.IsNil(o.CategoryVersion) {
        toSerialize["category_version"] = o.CategoryVersion
    }
    if !utils.IsNil(o.CreateTimeGe) {
        toSerialize["create_time_ge"] = o.CreateTimeGe
    }
    if !utils.IsNil(o.CreateTimeLe) {
        toSerialize["create_time_le"] = o.CreateTimeLe
    }
    if !utils.IsNil(o.ListingPlatforms) {
        toSerialize["listing_platforms"] = o.ListingPlatforms
    }
    if !utils.IsNil(o.ListingQualityTier) {
        toSerialize["listing_quality_tier"] = o.ListingQualityTier
    }
    if !utils.IsNil(o.SellerSkus) {
        toSerialize["seller_skus"] = o.SellerSkus
    }
    if !utils.IsNil(o.SkuIds) {
        toSerialize["sku_ids"] = o.SkuIds
    }
    if !utils.IsNil(o.Status) {
        toSerialize["status"] = o.Status
    }
    if !utils.IsNil(o.UpdateTimeGe) {
        toSerialize["update_time_ge"] = o.UpdateTimeGe
    }
    if !utils.IsNil(o.UpdateTimeLe) {
        toSerialize["update_time_le"] = o.UpdateTimeLe
    }
    return toSerialize, nil
}

type NullableProduct202312SearchProductsRequestBody struct {
	value *Product202312SearchProductsRequestBody
	isSet bool
}

func (v NullableProduct202312SearchProductsRequestBody) Get() *Product202312SearchProductsRequestBody {
	return v.value
}

func (v *NullableProduct202312SearchProductsRequestBody) Set(val *Product202312SearchProductsRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202312SearchProductsRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202312SearchProductsRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202312SearchProductsRequestBody(val *Product202312SearchProductsRequestBody) *NullableProduct202312SearchProductsRequestBody {
	return &NullableProduct202312SearchProductsRequestBody{value: val, isSet: true}
}

func (v NullableProduct202312SearchProductsRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202312SearchProductsRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


