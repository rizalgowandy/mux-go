/*
 * Mux API
 *
 * Mux is how developers build online video. This API encompasses both Mux Video and Mux Data functionality to help you build your video-related projects better and faster than ever before. 
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package muxgo

import (
	"encoding/json"
)

// DimensionValue struct for DimensionValue
type DimensionValue struct {
	Value *string `json:"value,omitempty"`
	TotalCount *int64 `json:"total_count,omitempty"`
}

// NewDimensionValue instantiates a new DimensionValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDimensionValue() *DimensionValue {
	this := DimensionValue{}
	return &this
}

// NewDimensionValueWithDefaults instantiates a new DimensionValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDimensionValueWithDefaults() *DimensionValue {
	this := DimensionValue{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DimensionValue) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DimensionValue) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DimensionValue) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DimensionValue) SetValue(v string) {
	o.Value = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *DimensionValue) GetTotalCount() int64 {
	if o == nil || o.TotalCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DimensionValue) GetTotalCountOk() (*int64, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *DimensionValue) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *DimensionValue) SetTotalCount(v int64) {
	o.TotalCount = &v
}

func (o DimensionValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.TotalCount != nil {
		toSerialize["total_count"] = o.TotalCount
	}
	return json.Marshal(toSerialize)
}

type NullableDimensionValue struct {
	value *DimensionValue
	isSet bool
}

func (v NullableDimensionValue) Get() *DimensionValue {
	return v.value
}

func (v *NullableDimensionValue) Set(val *DimensionValue) {
	v.value = val
	v.isSet = true
}

func (v NullableDimensionValue) IsSet() bool {
	return v.isSet
}

func (v *NullableDimensionValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDimensionValue(val *DimensionValue) *NullableDimensionValue {
	return &NullableDimensionValue{value: val, isSet: true}
}

func (v NullableDimensionValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDimensionValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

