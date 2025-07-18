/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product_v202409

import (
    "encoding/json"
    "tiktokshop/open/sdk_golang/utils"
)

            // checks if the Product202409SearchResponsiblePersonsRequestBody type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202409SearchResponsiblePersonsRequestBody{}

// Product202409SearchResponsiblePersonsRequestBody struct for Product202409SearchResponsiblePersonsRequestBody
type Product202409SearchResponsiblePersonsRequestBody struct {
    // Filter results to show those that contain this keyword. Search scope: name, local_number, email Max length: 200 characters  **Note**: Provide either the `responsible_person_ids` or `keyword`; if both are provided, `responsible_person_ids` will take priority.
    Keyword *string `json:"keyword,omitempty"`
    // Filter results by these responsible person IDs. Max IDs: The value of `page_size`
    ResponsiblePersonIds []string `json:"responsible_person_ids,omitempty"`
}

// NewProduct202409SearchResponsiblePersonsRequestBody instantiates a new Product202409SearchResponsiblePersonsRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202409SearchResponsiblePersonsRequestBody() *Product202409SearchResponsiblePersonsRequestBody {
    this := Product202409SearchResponsiblePersonsRequestBody{}
    return &this
}

// NewProduct202409SearchResponsiblePersonsRequestBodyWithDefaults instantiates a new Product202409SearchResponsiblePersonsRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202409SearchResponsiblePersonsRequestBodyWithDefaults() *Product202409SearchResponsiblePersonsRequestBody {
    this := Product202409SearchResponsiblePersonsRequestBody{}
    return &this
}

// GetKeyword returns the Keyword field value if set, zero value otherwise.
func (o *Product202409SearchResponsiblePersonsRequestBody) GetKeyword() string {
    if o == nil || utils.IsNil(o.Keyword) {
        var ret string
        return ret
    }
    return *o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202409SearchResponsiblePersonsRequestBody) GetKeywordOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Keyword) {
        return nil, false
    }
    return o.Keyword, true
}

// HasKeyword returns a boolean if a field has been set.
func (o *Product202409SearchResponsiblePersonsRequestBody) HasKeyword() bool {
    if o != nil && !utils.IsNil(o.Keyword) {
        return true
    }

    return false
}

// SetKeyword gets a reference to the given string and assigns it to the Keyword field.
func (o *Product202409SearchResponsiblePersonsRequestBody) SetKeyword(v string) {
    o.Keyword = &v
}

// GetResponsiblePersonIds returns the ResponsiblePersonIds field value if set, zero value otherwise.
func (o *Product202409SearchResponsiblePersonsRequestBody) GetResponsiblePersonIds() []string {
    if o == nil || utils.IsNil(o.ResponsiblePersonIds) {
        var ret []string
        return ret
    }
    return o.ResponsiblePersonIds
}

// GetResponsiblePersonIdsOk returns a tuple with the ResponsiblePersonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202409SearchResponsiblePersonsRequestBody) GetResponsiblePersonIdsOk() ([]string, bool) {
    if o == nil || utils.IsNil(o.ResponsiblePersonIds) {
        return nil, false
    }
    return o.ResponsiblePersonIds, true
}

// HasResponsiblePersonIds returns a boolean if a field has been set.
func (o *Product202409SearchResponsiblePersonsRequestBody) HasResponsiblePersonIds() bool {
    if o != nil && !utils.IsNil(o.ResponsiblePersonIds) {
        return true
    }

    return false
}

// SetResponsiblePersonIds gets a reference to the given []string and assigns it to the ResponsiblePersonIds field.
func (o *Product202409SearchResponsiblePersonsRequestBody) SetResponsiblePersonIds(v []string) {
    o.ResponsiblePersonIds = v
}

func (o Product202409SearchResponsiblePersonsRequestBody) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202409SearchResponsiblePersonsRequestBody) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Keyword) {
        toSerialize["keyword"] = o.Keyword
    }
    if !utils.IsNil(o.ResponsiblePersonIds) {
        toSerialize["responsible_person_ids"] = o.ResponsiblePersonIds
    }
    return toSerialize, nil
}

type NullableProduct202409SearchResponsiblePersonsRequestBody struct {
	value *Product202409SearchResponsiblePersonsRequestBody
	isSet bool
}

func (v NullableProduct202409SearchResponsiblePersonsRequestBody) Get() *Product202409SearchResponsiblePersonsRequestBody {
	return v.value
}

func (v *NullableProduct202409SearchResponsiblePersonsRequestBody) Set(val *Product202409SearchResponsiblePersonsRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202409SearchResponsiblePersonsRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202409SearchResponsiblePersonsRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202409SearchResponsiblePersonsRequestBody(val *Product202409SearchResponsiblePersonsRequestBody) *NullableProduct202409SearchResponsiblePersonsRequestBody {
	return &NullableProduct202409SearchResponsiblePersonsRequestBody{value: val, isSet: true}
}

func (v NullableProduct202409SearchResponsiblePersonsRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202409SearchResponsiblePersonsRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


