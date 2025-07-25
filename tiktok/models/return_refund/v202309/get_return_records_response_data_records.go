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

            // checks if the ReturnRefund202309GetReturnRecordsResponseDataRecords type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ReturnRefund202309GetReturnRecordsResponseDataRecords{}

// ReturnRefund202309GetReturnRecordsResponseDataRecords struct for ReturnRefund202309GetReturnRecordsResponseDataRecords
type ReturnRefund202309GetReturnRecordsResponseDataRecords struct {
    // The creation time for the return. Unix timestamp.
    CreateTime *int64 `json:"create_time,omitempty"`
    // Description of the return record.
    Description *string `json:"description,omitempty"`
    // The type of order event. In this case, it will always be ORDER_RETURN.
    Event *string `json:"event,omitempty"`
    // Images provided by the seller or buyer. You can use the role field to differentiate whether the images are from the seller or buyer.
    Images []ReturnRefund202309GetReturnRecordsResponseDataRecordsImages `json:"images,omitempty"`
    // A note provided by the seller or buyer. You can use the role field to differentiate whether the note is from the seller or buyer.
    Note *string `json:"note,omitempty"`
    // The corresponding text for a return reason, localized based on the locale input parameter. 
    ReasonText *string `json:"reason_text,omitempty"`
    // The role that initiated the order return request. Possible values: - BUYER - SELLER - OPERATOR: If the order is canceled by the customer service agent manually, then the cancel initiator will be 'OPERATOR'. - SYSTEM: If the order is automatically canceled due to a policy reason, then the cancel initiator will be 'SYSTEM'.
    Role *string `json:"role,omitempty"`
    // Videos uploaded by the buyer. Only buyers are allowed to upload videos.
    Videos []ReturnRefund202309GetReturnRecordsResponseDataRecordsVideos `json:"videos,omitempty"`
}

// NewReturnRefund202309GetReturnRecordsResponseDataRecords instantiates a new ReturnRefund202309GetReturnRecordsResponseDataRecords object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnRefund202309GetReturnRecordsResponseDataRecords() *ReturnRefund202309GetReturnRecordsResponseDataRecords {
    this := ReturnRefund202309GetReturnRecordsResponseDataRecords{}
    return &this
}

// NewReturnRefund202309GetReturnRecordsResponseDataRecordsWithDefaults instantiates a new ReturnRefund202309GetReturnRecordsResponseDataRecords object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnRefund202309GetReturnRecordsResponseDataRecordsWithDefaults() *ReturnRefund202309GetReturnRecordsResponseDataRecords {
    this := ReturnRefund202309GetReturnRecordsResponseDataRecords{}
    return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) GetCreateTime() int64 {
    if o == nil || utils.IsNil(o.CreateTime) {
        var ret int64
        return ret
    }
    return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) GetCreateTimeOk() (*int64, bool) {
    if o == nil || utils.IsNil(o.CreateTime) {
        return nil, false
    }
    return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) HasCreateTime() bool {
    if o != nil && !utils.IsNil(o.CreateTime) {
        return true
    }

    return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) SetCreateTime(v int64) {
    o.CreateTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) GetDescription() string {
    if o == nil || utils.IsNil(o.Description) {
        var ret string
        return ret
    }
    return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) GetDescriptionOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Description) {
        return nil, false
    }
    return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) HasDescription() bool {
    if o != nil && !utils.IsNil(o.Description) {
        return true
    }

    return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) SetDescription(v string) {
    o.Description = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) GetEvent() string {
    if o == nil || utils.IsNil(o.Event) {
        var ret string
        return ret
    }
    return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) GetEventOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Event) {
        return nil, false
    }
    return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) HasEvent() bool {
    if o != nil && !utils.IsNil(o.Event) {
        return true
    }

    return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) SetEvent(v string) {
    o.Event = &v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) GetImages() []ReturnRefund202309GetReturnRecordsResponseDataRecordsImages {
    if o == nil || utils.IsNil(o.Images) {
        var ret []ReturnRefund202309GetReturnRecordsResponseDataRecordsImages
        return ret
    }
    return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) GetImagesOk() ([]ReturnRefund202309GetReturnRecordsResponseDataRecordsImages, bool) {
    if o == nil || utils.IsNil(o.Images) {
        return nil, false
    }
    return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) HasImages() bool {
    if o != nil && !utils.IsNil(o.Images) {
        return true
    }

    return false
}

// SetImages gets a reference to the given []ReturnRefund202309GetReturnRecordsResponseDataRecordsImages and assigns it to the Images field.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) SetImages(v []ReturnRefund202309GetReturnRecordsResponseDataRecordsImages) {
    o.Images = v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) GetNote() string {
    if o == nil || utils.IsNil(o.Note) {
        var ret string
        return ret
    }
    return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) GetNoteOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Note) {
        return nil, false
    }
    return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) HasNote() bool {
    if o != nil && !utils.IsNil(o.Note) {
        return true
    }

    return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) SetNote(v string) {
    o.Note = &v
}

// GetReasonText returns the ReasonText field value if set, zero value otherwise.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) GetReasonText() string {
    if o == nil || utils.IsNil(o.ReasonText) {
        var ret string
        return ret
    }
    return *o.ReasonText
}

// GetReasonTextOk returns a tuple with the ReasonText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) GetReasonTextOk() (*string, bool) {
    if o == nil || utils.IsNil(o.ReasonText) {
        return nil, false
    }
    return o.ReasonText, true
}

// HasReasonText returns a boolean if a field has been set.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) HasReasonText() bool {
    if o != nil && !utils.IsNil(o.ReasonText) {
        return true
    }

    return false
}

// SetReasonText gets a reference to the given string and assigns it to the ReasonText field.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) SetReasonText(v string) {
    o.ReasonText = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) GetRole() string {
    if o == nil || utils.IsNil(o.Role) {
        var ret string
        return ret
    }
    return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) GetRoleOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Role) {
        return nil, false
    }
    return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) HasRole() bool {
    if o != nil && !utils.IsNil(o.Role) {
        return true
    }

    return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) SetRole(v string) {
    o.Role = &v
}

// GetVideos returns the Videos field value if set, zero value otherwise.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) GetVideos() []ReturnRefund202309GetReturnRecordsResponseDataRecordsVideos {
    if o == nil || utils.IsNil(o.Videos) {
        var ret []ReturnRefund202309GetReturnRecordsResponseDataRecordsVideos
        return ret
    }
    return o.Videos
}

// GetVideosOk returns a tuple with the Videos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) GetVideosOk() ([]ReturnRefund202309GetReturnRecordsResponseDataRecordsVideos, bool) {
    if o == nil || utils.IsNil(o.Videos) {
        return nil, false
    }
    return o.Videos, true
}

// HasVideos returns a boolean if a field has been set.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) HasVideos() bool {
    if o != nil && !utils.IsNil(o.Videos) {
        return true
    }

    return false
}

// SetVideos gets a reference to the given []ReturnRefund202309GetReturnRecordsResponseDataRecordsVideos and assigns it to the Videos field.
func (o *ReturnRefund202309GetReturnRecordsResponseDataRecords) SetVideos(v []ReturnRefund202309GetReturnRecordsResponseDataRecordsVideos) {
    o.Videos = v
}

func (o ReturnRefund202309GetReturnRecordsResponseDataRecords) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o ReturnRefund202309GetReturnRecordsResponseDataRecords) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.CreateTime) {
        toSerialize["create_time"] = o.CreateTime
    }
    if !utils.IsNil(o.Description) {
        toSerialize["description"] = o.Description
    }
    if !utils.IsNil(o.Event) {
        toSerialize["event"] = o.Event
    }
    if !utils.IsNil(o.Images) {
        toSerialize["images"] = o.Images
    }
    if !utils.IsNil(o.Note) {
        toSerialize["note"] = o.Note
    }
    if !utils.IsNil(o.ReasonText) {
        toSerialize["reason_text"] = o.ReasonText
    }
    if !utils.IsNil(o.Role) {
        toSerialize["role"] = o.Role
    }
    if !utils.IsNil(o.Videos) {
        toSerialize["videos"] = o.Videos
    }
    return toSerialize, nil
}

type NullableReturnRefund202309GetReturnRecordsResponseDataRecords struct {
	value *ReturnRefund202309GetReturnRecordsResponseDataRecords
	isSet bool
}

func (v NullableReturnRefund202309GetReturnRecordsResponseDataRecords) Get() *ReturnRefund202309GetReturnRecordsResponseDataRecords {
	return v.value
}

func (v *NullableReturnRefund202309GetReturnRecordsResponseDataRecords) Set(val *ReturnRefund202309GetReturnRecordsResponseDataRecords) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnRefund202309GetReturnRecordsResponseDataRecords) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnRefund202309GetReturnRecordsResponseDataRecords) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnRefund202309GetReturnRecordsResponseDataRecords(val *ReturnRefund202309GetReturnRecordsResponseDataRecords) *NullableReturnRefund202309GetReturnRecordsResponseDataRecords {
	return &NullableReturnRefund202309GetReturnRecordsResponseDataRecords{value: val, isSet: true}
}

func (v NullableReturnRefund202309GetReturnRecordsResponseDataRecords) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnRefund202309GetReturnRecordsResponseDataRecords) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


