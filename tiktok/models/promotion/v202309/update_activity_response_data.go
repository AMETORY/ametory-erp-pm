/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package promotion_v202309

import (
    "encoding/json"
    "tiktokshop/open/sdk_golang/utils"
)

            // checks if the Promotion202309UpdateActivityResponseData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Promotion202309UpdateActivityResponseData{}

// Promotion202309UpdateActivityResponseData struct for Promotion202309UpdateActivityResponseData
type Promotion202309UpdateActivityResponseData struct {
    // A unique ID that identifies different activities.
    ActivityId *string `json:"activity_id,omitempty"`
    // Activity name set by the seller.
    Title *string `json:"title,omitempty"`
    // Last update time. UNIX timestamp.
    UpdateTime *int64 `json:"update_time,omitempty"`
}

// NewPromotion202309UpdateActivityResponseData instantiates a new Promotion202309UpdateActivityResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotion202309UpdateActivityResponseData() *Promotion202309UpdateActivityResponseData {
    this := Promotion202309UpdateActivityResponseData{}
    return &this
}

// NewPromotion202309UpdateActivityResponseDataWithDefaults instantiates a new Promotion202309UpdateActivityResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotion202309UpdateActivityResponseDataWithDefaults() *Promotion202309UpdateActivityResponseData {
    this := Promotion202309UpdateActivityResponseData{}
    return &this
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise.
func (o *Promotion202309UpdateActivityResponseData) GetActivityId() string {
    if o == nil || utils.IsNil(o.ActivityId) {
        var ret string
        return ret
    }
    return *o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202309UpdateActivityResponseData) GetActivityIdOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ActivityId) {
        return nil, false
    }
    return o.ActivityId, true
}

// HasActivityId returns a boolean if a field has been set.
func (o *Promotion202309UpdateActivityResponseData) HasActivityId() bool {
    if o != nil && !utils.IsNil(o.ActivityId) {
        return true
    }

    return false
}

// SetActivityId gets a reference to the given string and assigns it to the ActivityId field.
func (o *Promotion202309UpdateActivityResponseData) SetActivityId(v string) {
    o.ActivityId = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Promotion202309UpdateActivityResponseData) GetTitle() string {
    if o == nil || utils.IsNil(o.Title) {
        var ret string
        return ret
    }
    return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202309UpdateActivityResponseData) GetTitleOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Title) {
        return nil, false
    }
    return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Promotion202309UpdateActivityResponseData) HasTitle() bool {
    if o != nil && !utils.IsNil(o.Title) {
        return true
    }

    return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Promotion202309UpdateActivityResponseData) SetTitle(v string) {
    o.Title = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *Promotion202309UpdateActivityResponseData) GetUpdateTime() int64 {
    if o == nil || utils.IsNil(o.UpdateTime) {
        var ret int64
        return ret
    }
    return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion202309UpdateActivityResponseData) GetUpdateTimeOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.UpdateTime) {
        return nil, false
    }
    return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *Promotion202309UpdateActivityResponseData) HasUpdateTime() bool {
    if o != nil && !utils.IsNil(o.UpdateTime) {
        return true
    }

    return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *Promotion202309UpdateActivityResponseData) SetUpdateTime(v int64) {
    o.UpdateTime = &v
}

func (o Promotion202309UpdateActivityResponseData) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Promotion202309UpdateActivityResponseData) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.ActivityId) {
        toSerialize["activity_id"] = o.ActivityId
    }
    if !utils.IsNil(o.Title) {
        toSerialize["title"] = o.Title
    }
    if !utils.IsNil(o.UpdateTime) {
        toSerialize["update_time"] = o.UpdateTime
    }
    return toSerialize, nil
}

type NullablePromotion202309UpdateActivityResponseData struct {
	value *Promotion202309UpdateActivityResponseData
	isSet bool
}

func (v NullablePromotion202309UpdateActivityResponseData) Get() *Promotion202309UpdateActivityResponseData {
	return v.value
}

func (v *NullablePromotion202309UpdateActivityResponseData) Set(val *Promotion202309UpdateActivityResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotion202309UpdateActivityResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotion202309UpdateActivityResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotion202309UpdateActivityResponseData(val *Promotion202309UpdateActivityResponseData) *NullablePromotion202309UpdateActivityResponseData {
	return &NullablePromotion202309UpdateActivityResponseData{value: val, isSet: true}
}

func (v NullablePromotion202309UpdateActivityResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotion202309UpdateActivityResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


