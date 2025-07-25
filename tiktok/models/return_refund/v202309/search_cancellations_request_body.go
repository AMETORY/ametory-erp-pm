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

            // checks if the ReturnRefund202309SearchCancellationsRequestBody type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ReturnRefund202309SearchCancellationsRequestBody{}

// ReturnRefund202309SearchCancellationsRequestBody struct for ReturnRefund202309SearchCancellationsRequestBody
type ReturnRefund202309SearchCancellationsRequestBody struct {
    // List of TikTok Shop buyer user IDs.
    BuyerUserIds []string `json:"buyer_user_ids,omitempty"`
    // List of order cancellations IDs.
    CancelIds []string `json:"cancel_ids,omitempty"`
    // List of order cancellation statuses. Possible values: - CANCELLATION_REQUEST_PENDING - CANCELLATION_REQUEST_SUCCESS - CANCELLATION_REQUEST_CANCEL - CANCELLATION_REQUEST_COMPLETE  Please see \"API Overview\" for more information about these statuses.
    CancelStatus []string `json:"cancel_status,omitempty"`
    // List of order cancellation types. Possible values: - CANCEL: Cancel by seller or system. - BUYER_CANCEL: Cancel by buyer. Need to be approved by seller or system.
    CancelTypes []string `json:"cancel_types,omitempty"`
    // Filter cancellations to show only orders that have been created after a specified date and time. Unix timestamp. 
    CreateTimeGe *int64 `json:"create_time_ge,omitempty"`
    // Filter cancellations to show only orders that have been created before a specified date and time. Unix timestamp. 
    CreateTimeLt *int64 `json:"create_time_lt,omitempty"`
    // The BCP-47 locale codes for displaying the order, delimited by commas. Default: en-US Refer to [Locale codes](678e3a47bae28f030a8c7523) for the list of supported locale codes.
    Locale *string `json:"locale,omitempty"`
    // List of TikTok Shop order IDs.
    OrderIds []string `json:"order_ids,omitempty"`
    // Filter cancellations to show only orders that have been updated after a specified date and time. Unix timestamp.
    UpdateTimeGe *int64 `json:"update_time_ge,omitempty"`
    // Filter cancellations to show only orders that have been updated before a specified date and time. Unix timestamp.
    UpdateTimeLt *int64 `json:"update_time_lt,omitempty"`
}

// NewReturnRefund202309SearchCancellationsRequestBody instantiates a new ReturnRefund202309SearchCancellationsRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnRefund202309SearchCancellationsRequestBody() *ReturnRefund202309SearchCancellationsRequestBody {
    this := ReturnRefund202309SearchCancellationsRequestBody{}
    return &this
}

// NewReturnRefund202309SearchCancellationsRequestBodyWithDefaults instantiates a new ReturnRefund202309SearchCancellationsRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnRefund202309SearchCancellationsRequestBodyWithDefaults() *ReturnRefund202309SearchCancellationsRequestBody {
    this := ReturnRefund202309SearchCancellationsRequestBody{}
    return &this
}

// GetBuyerUserIds returns the BuyerUserIds field value if set, zero value otherwise.
func (o *ReturnRefund202309SearchCancellationsRequestBody) GetBuyerUserIds() []string {
    if o == nil || utils.IsNil(o.BuyerUserIds) {
        var ret []string
        return ret
    }
    return o.BuyerUserIds
}

// GetBuyerUserIdsOk returns a tuple with the BuyerUserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309SearchCancellationsRequestBody) GetBuyerUserIdsOk() ([]string, bool) {
    if o == nil || utils.IsNil(o.BuyerUserIds) {
        return nil, false
    }
    return o.BuyerUserIds, true
}

// HasBuyerUserIds returns a boolean if a field has been set.
func (o *ReturnRefund202309SearchCancellationsRequestBody) HasBuyerUserIds() bool {
    if o != nil && !utils.IsNil(o.BuyerUserIds) {
        return true
    }

    return false
}

// SetBuyerUserIds gets a reference to the given []string and assigns it to the BuyerUserIds field.
func (o *ReturnRefund202309SearchCancellationsRequestBody) SetBuyerUserIds(v []string) {
    o.BuyerUserIds = v
}

// GetCancelIds returns the CancelIds field value if set, zero value otherwise.
func (o *ReturnRefund202309SearchCancellationsRequestBody) GetCancelIds() []string {
    if o == nil || utils.IsNil(o.CancelIds) {
        var ret []string
        return ret
    }
    return o.CancelIds
}

// GetCancelIdsOk returns a tuple with the CancelIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309SearchCancellationsRequestBody) GetCancelIdsOk() ([]string, bool) {
    if o == nil || utils.IsNil(o.CancelIds) {
        return nil, false
    }
    return o.CancelIds, true
}

// HasCancelIds returns a boolean if a field has been set.
func (o *ReturnRefund202309SearchCancellationsRequestBody) HasCancelIds() bool {
    if o != nil && !utils.IsNil(o.CancelIds) {
        return true
    }

    return false
}

// SetCancelIds gets a reference to the given []string and assigns it to the CancelIds field.
func (o *ReturnRefund202309SearchCancellationsRequestBody) SetCancelIds(v []string) {
    o.CancelIds = v
}

// GetCancelStatus returns the CancelStatus field value if set, zero value otherwise.
func (o *ReturnRefund202309SearchCancellationsRequestBody) GetCancelStatus() []string {
    if o == nil || utils.IsNil(o.CancelStatus) {
        var ret []string
        return ret
    }
    return o.CancelStatus
}

// GetCancelStatusOk returns a tuple with the CancelStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309SearchCancellationsRequestBody) GetCancelStatusOk() ([]string, bool) {
    if o == nil || utils.IsNil(o.CancelStatus) {
        return nil, false
    }
    return o.CancelStatus, true
}

// HasCancelStatus returns a boolean if a field has been set.
func (o *ReturnRefund202309SearchCancellationsRequestBody) HasCancelStatus() bool {
    if o != nil && !utils.IsNil(o.CancelStatus) {
        return true
    }

    return false
}

// SetCancelStatus gets a reference to the given []string and assigns it to the CancelStatus field.
func (o *ReturnRefund202309SearchCancellationsRequestBody) SetCancelStatus(v []string) {
    o.CancelStatus = v
}

// GetCancelTypes returns the CancelTypes field value if set, zero value otherwise.
func (o *ReturnRefund202309SearchCancellationsRequestBody) GetCancelTypes() []string {
    if o == nil || utils.IsNil(o.CancelTypes) {
        var ret []string
        return ret
    }
    return o.CancelTypes
}

// GetCancelTypesOk returns a tuple with the CancelTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309SearchCancellationsRequestBody) GetCancelTypesOk() ([]string, bool) {
    if o == nil || utils.IsNil(o.CancelTypes) {
        return nil, false
    }
    return o.CancelTypes, true
}

// HasCancelTypes returns a boolean if a field has been set.
func (o *ReturnRefund202309SearchCancellationsRequestBody) HasCancelTypes() bool {
    if o != nil && !utils.IsNil(o.CancelTypes) {
        return true
    }

    return false
}

// SetCancelTypes gets a reference to the given []string and assigns it to the CancelTypes field.
func (o *ReturnRefund202309SearchCancellationsRequestBody) SetCancelTypes(v []string) {
    o.CancelTypes = v
}

// GetCreateTimeGe returns the CreateTimeGe field value if set, zero value otherwise.
func (o *ReturnRefund202309SearchCancellationsRequestBody) GetCreateTimeGe() int64 {
    if o == nil || utils.IsNil(o.CreateTimeGe) {
        var ret int64
        return ret
    }
    return *o.CreateTimeGe
}

// GetCreateTimeGeOk returns a tuple with the CreateTimeGe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309SearchCancellationsRequestBody) GetCreateTimeGeOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.CreateTimeGe) {
        return nil, false
    }
    return o.CreateTimeGe, true
}

// HasCreateTimeGe returns a boolean if a field has been set.
func (o *ReturnRefund202309SearchCancellationsRequestBody) HasCreateTimeGe() bool {
    if o != nil && !utils.IsNil(o.CreateTimeGe) {
        return true
    }

    return false
}

// SetCreateTimeGe gets a reference to the given int64 and assigns it to the CreateTimeGe field.
func (o *ReturnRefund202309SearchCancellationsRequestBody) SetCreateTimeGe(v int64) {
    o.CreateTimeGe = &v
}

// GetCreateTimeLt returns the CreateTimeLt field value if set, zero value otherwise.
func (o *ReturnRefund202309SearchCancellationsRequestBody) GetCreateTimeLt() int64 {
    if o == nil || utils.IsNil(o.CreateTimeLt) {
        var ret int64
        return ret
    }
    return *o.CreateTimeLt
}

// GetCreateTimeLtOk returns a tuple with the CreateTimeLt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309SearchCancellationsRequestBody) GetCreateTimeLtOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.CreateTimeLt) {
        return nil, false
    }
    return o.CreateTimeLt, true
}

// HasCreateTimeLt returns a boolean if a field has been set.
func (o *ReturnRefund202309SearchCancellationsRequestBody) HasCreateTimeLt() bool {
    if o != nil && !utils.IsNil(o.CreateTimeLt) {
        return true
    }

    return false
}

// SetCreateTimeLt gets a reference to the given int64 and assigns it to the CreateTimeLt field.
func (o *ReturnRefund202309SearchCancellationsRequestBody) SetCreateTimeLt(v int64) {
    o.CreateTimeLt = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *ReturnRefund202309SearchCancellationsRequestBody) GetLocale() string {
    if o == nil || utils.IsNil(o.Locale) {
        var ret string
        return ret
    }
    return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309SearchCancellationsRequestBody) GetLocaleOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Locale) {
        return nil, false
    }
    return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *ReturnRefund202309SearchCancellationsRequestBody) HasLocale() bool {
    if o != nil && !utils.IsNil(o.Locale) {
        return true
    }

    return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *ReturnRefund202309SearchCancellationsRequestBody) SetLocale(v string) {
    o.Locale = &v
}

// GetOrderIds returns the OrderIds field value if set, zero value otherwise.
func (o *ReturnRefund202309SearchCancellationsRequestBody) GetOrderIds() []string {
    if o == nil || utils.IsNil(o.OrderIds) {
        var ret []string
        return ret
    }
    return o.OrderIds
}

// GetOrderIdsOk returns a tuple with the OrderIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309SearchCancellationsRequestBody) GetOrderIdsOk() ([]string, bool) {
    if o == nil || utils.IsNil(o.OrderIds) {
        return nil, false
    }
    return o.OrderIds, true
}

// HasOrderIds returns a boolean if a field has been set.
func (o *ReturnRefund202309SearchCancellationsRequestBody) HasOrderIds() bool {
    if o != nil && !utils.IsNil(o.OrderIds) {
        return true
    }

    return false
}

// SetOrderIds gets a reference to the given []string and assigns it to the OrderIds field.
func (o *ReturnRefund202309SearchCancellationsRequestBody) SetOrderIds(v []string) {
    o.OrderIds = v
}

// GetUpdateTimeGe returns the UpdateTimeGe field value if set, zero value otherwise.
func (o *ReturnRefund202309SearchCancellationsRequestBody) GetUpdateTimeGe() int64 {
    if o == nil || utils.IsNil(o.UpdateTimeGe) {
        var ret int64
        return ret
    }
    return *o.UpdateTimeGe
}

// GetUpdateTimeGeOk returns a tuple with the UpdateTimeGe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309SearchCancellationsRequestBody) GetUpdateTimeGeOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.UpdateTimeGe) {
        return nil, false
    }
    return o.UpdateTimeGe, true
}

// HasUpdateTimeGe returns a boolean if a field has been set.
func (o *ReturnRefund202309SearchCancellationsRequestBody) HasUpdateTimeGe() bool {
    if o != nil && !utils.IsNil(o.UpdateTimeGe) {
        return true
    }

    return false
}

// SetUpdateTimeGe gets a reference to the given int64 and assigns it to the UpdateTimeGe field.
func (o *ReturnRefund202309SearchCancellationsRequestBody) SetUpdateTimeGe(v int64) {
    o.UpdateTimeGe = &v
}

// GetUpdateTimeLt returns the UpdateTimeLt field value if set, zero value otherwise.
func (o *ReturnRefund202309SearchCancellationsRequestBody) GetUpdateTimeLt() int64 {
    if o == nil || utils.IsNil(o.UpdateTimeLt) {
        var ret int64
        return ret
    }
    return *o.UpdateTimeLt
}

// GetUpdateTimeLtOk returns a tuple with the UpdateTimeLt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309SearchCancellationsRequestBody) GetUpdateTimeLtOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.UpdateTimeLt) {
        return nil, false
    }
    return o.UpdateTimeLt, true
}

// HasUpdateTimeLt returns a boolean if a field has been set.
func (o *ReturnRefund202309SearchCancellationsRequestBody) HasUpdateTimeLt() bool {
    if o != nil && !utils.IsNil(o.UpdateTimeLt) {
        return true
    }

    return false
}

// SetUpdateTimeLt gets a reference to the given int64 and assigns it to the UpdateTimeLt field.
func (o *ReturnRefund202309SearchCancellationsRequestBody) SetUpdateTimeLt(v int64) {
    o.UpdateTimeLt = &v
}

func (o ReturnRefund202309SearchCancellationsRequestBody) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o ReturnRefund202309SearchCancellationsRequestBody) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.BuyerUserIds) {
        toSerialize["buyer_user_ids"] = o.BuyerUserIds
    }
    if !utils.IsNil(o.CancelIds) {
        toSerialize["cancel_ids"] = o.CancelIds
    }
    if !utils.IsNil(o.CancelStatus) {
        toSerialize["cancel_status"] = o.CancelStatus
    }
    if !utils.IsNil(o.CancelTypes) {
        toSerialize["cancel_types"] = o.CancelTypes
    }
    if !utils.IsNil(o.CreateTimeGe) {
        toSerialize["create_time_ge"] = o.CreateTimeGe
    }
    if !utils.IsNil(o.CreateTimeLt) {
        toSerialize["create_time_lt"] = o.CreateTimeLt
    }
    if !utils.IsNil(o.Locale) {
        toSerialize["locale"] = o.Locale
    }
    if !utils.IsNil(o.OrderIds) {
        toSerialize["order_ids"] = o.OrderIds
    }
    if !utils.IsNil(o.UpdateTimeGe) {
        toSerialize["update_time_ge"] = o.UpdateTimeGe
    }
    if !utils.IsNil(o.UpdateTimeLt) {
        toSerialize["update_time_lt"] = o.UpdateTimeLt
    }
    return toSerialize, nil
}

type NullableReturnRefund202309SearchCancellationsRequestBody struct {
	value *ReturnRefund202309SearchCancellationsRequestBody
	isSet bool
}

func (v NullableReturnRefund202309SearchCancellationsRequestBody) Get() *ReturnRefund202309SearchCancellationsRequestBody {
	return v.value
}

func (v *NullableReturnRefund202309SearchCancellationsRequestBody) Set(val *ReturnRefund202309SearchCancellationsRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnRefund202309SearchCancellationsRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnRefund202309SearchCancellationsRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnRefund202309SearchCancellationsRequestBody(val *ReturnRefund202309SearchCancellationsRequestBody) *NullableReturnRefund202309SearchCancellationsRequestBody {
	return &NullableReturnRefund202309SearchCancellationsRequestBody{value: val, isSet: true}
}

func (v NullableReturnRefund202309SearchCancellationsRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnRefund202309SearchCancellationsRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


