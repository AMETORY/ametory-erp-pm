/*
tiktok shop openapi

sdk for apis

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product_v202407

import (
    "encoding/json"
    "tiktokshop/open/sdk_golang/utils"
)

            // checks if the Product202407ListingSchemasResponseDataListingSchemasFieldsRules type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Product202407ListingSchemasResponseDataListingSchemasFieldsRules{}

// Product202407ListingSchemasResponseDataListingSchemasFieldsRules struct for Product202407ListingSchemasResponseDataListingSchemasFieldsRules
type Product202407ListingSchemasResponseDataListingSchemasFieldsRules struct {
    // The type of rule, with detailed explanation, such as: - VALUE_TYPE(required field): The field values for the following data types need to be satisfied, including: Types: string (text type), For example: Title, SellerSKU, custom properties integer (integer type), For example: Inventory amount. date (date type), For example: Creation time, update time, etc. uri (media resource ID),For example: Main image ID html (text supporting HTML markup syntax) For example: Product description - REQUIRED(required field): Is the field a required field. - SUPPORTED(optional field): Is the field a supported field. - DISABLE(required field): The rule description field is a deprecated field. - MAX_LENGTH/ MIN_LENGTH(optional field):The maximum/minimum length generally refers to the character length limit. - MAX_VALUE/MIN_VALUE(optional field): The maximum/minimum value generally refers to the numerical limit. - MAX_INPUT_NUM/MIN_INPUT_NUM(optional field): The maximum/minimum number of selections generally refers to the number of options that can be selected in a multiple-choice scenario. - MAX_TARGE_TSIZE/MIN_TARGE_TSIZE(optional field):The maximum/minimum target file size generally refers to the size of the resource. - REGX(optional field)：Regular expression matching refers to input rules for input classes. - TIP(optional field):Provide an explanation for filling in the rule description field. - SAMPLE(optional field):Provide an example for filling in the rule description field. - CUSTOM(optional field):Explain whether the rule supports customization for fields. By default, customization is not supported. - MULTI_INPUT(optional field):Explain whether the rule supports multiple inputs for fields. By default, it is single input. - AVAILABLE(optional field)：The rule is used to express whether it is in an available state. For example, the category is available. - AUTHORIZED(optional field)：The rule is used to express whether it is in an authorized state, for example, scheduled category authorization.
    Type *string `json:"type,omitempty"`
    // The values of the rules
    Value *string `json:"value,omitempty"`
}

// NewProduct202407ListingSchemasResponseDataListingSchemasFieldsRules instantiates a new Product202407ListingSchemasResponseDataListingSchemasFieldsRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProduct202407ListingSchemasResponseDataListingSchemasFieldsRules() *Product202407ListingSchemasResponseDataListingSchemasFieldsRules {
    this := Product202407ListingSchemasResponseDataListingSchemasFieldsRules{}
    return &this
}

// NewProduct202407ListingSchemasResponseDataListingSchemasFieldsRulesWithDefaults instantiates a new Product202407ListingSchemasResponseDataListingSchemasFieldsRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProduct202407ListingSchemasResponseDataListingSchemasFieldsRulesWithDefaults() *Product202407ListingSchemasResponseDataListingSchemasFieldsRules {
    this := Product202407ListingSchemasResponseDataListingSchemasFieldsRules{}
    return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Product202407ListingSchemasResponseDataListingSchemasFieldsRules) GetType() string {
    if o == nil || utils.IsNil(o.Type) {
        var ret string
        return ret
    }
    return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202407ListingSchemasResponseDataListingSchemasFieldsRules) GetTypeOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Type) {
        return nil, false
    }
    return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Product202407ListingSchemasResponseDataListingSchemasFieldsRules) HasType() bool {
    if o != nil && !utils.IsNil(o.Type) {
        return true
    }

    return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Product202407ListingSchemasResponseDataListingSchemasFieldsRules) SetType(v string) {
    o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Product202407ListingSchemasResponseDataListingSchemasFieldsRules) GetValue() string {
    if o == nil || utils.IsNil(o.Value) {
        var ret string
        return ret
    }
    return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Product202407ListingSchemasResponseDataListingSchemasFieldsRules) GetValueOk() (*string, bool) {
    if o == nil || utils.IsNil(o.Value) {
        return nil, false
    }
    return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Product202407ListingSchemasResponseDataListingSchemasFieldsRules) HasValue() bool {
    if o != nil && !utils.IsNil(o.Value) {
        return true
    }

    return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *Product202407ListingSchemasResponseDataListingSchemasFieldsRules) SetValue(v string) {
    o.Value = &v
}

func (o Product202407ListingSchemasResponseDataListingSchemasFieldsRules) MarshalJSON() ([]byte, error) {
    toSerialize,err := o.ToMap()
    if err != nil {
        return []byte{}, err
    }
    return json.Marshal(toSerialize)
}

func (o Product202407ListingSchemasResponseDataListingSchemasFieldsRules) ToMap() (map[string]interface{}, error) {
    toSerialize := map[string]interface{}{}
    if !utils.IsNil(o.Type) {
        toSerialize["type"] = o.Type
    }
    if !utils.IsNil(o.Value) {
        toSerialize["value"] = o.Value
    }
    return toSerialize, nil
}

type NullableProduct202407ListingSchemasResponseDataListingSchemasFieldsRules struct {
	value *Product202407ListingSchemasResponseDataListingSchemasFieldsRules
	isSet bool
}

func (v NullableProduct202407ListingSchemasResponseDataListingSchemasFieldsRules) Get() *Product202407ListingSchemasResponseDataListingSchemasFieldsRules {
	return v.value
}

func (v *NullableProduct202407ListingSchemasResponseDataListingSchemasFieldsRules) Set(val *Product202407ListingSchemasResponseDataListingSchemasFieldsRules) {
	v.value = val
	v.isSet = true
}

func (v NullableProduct202407ListingSchemasResponseDataListingSchemasFieldsRules) IsSet() bool {
	return v.isSet
}

func (v *NullableProduct202407ListingSchemasResponseDataListingSchemasFieldsRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProduct202407ListingSchemasResponseDataListingSchemasFieldsRules(val *Product202407ListingSchemasResponseDataListingSchemasFieldsRules) *NullableProduct202407ListingSchemasResponseDataListingSchemasFieldsRules {
	return &NullableProduct202407ListingSchemasResponseDataListingSchemasFieldsRules{value: val, isSet: true}
}

func (v NullableProduct202407ListingSchemasResponseDataListingSchemasFieldsRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProduct202407ListingSchemasResponseDataListingSchemasFieldsRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


