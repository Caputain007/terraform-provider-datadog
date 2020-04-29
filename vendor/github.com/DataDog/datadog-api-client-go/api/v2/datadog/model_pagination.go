/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// Pagination Pagination object.
type Pagination struct {
	// Total count of pages.
	TotalCount *int64 `json:"total_count,omitempty"`
	// Total count of filtered pages.
	TotalFilteredCount *int64 `json:"total_filtered_count,omitempty"`
}

// NewPagination instantiates a new Pagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagination() *Pagination {
	this := Pagination{}
	return &this
}

// NewPaginationWithDefaults instantiates a new Pagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationWithDefaults() *Pagination {
	this := Pagination{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *Pagination) GetTotalCount() int64 {
	if o == nil || o.TotalCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetTotalCountOk() (*int64, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *Pagination) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *Pagination) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// GetTotalFilteredCount returns the TotalFilteredCount field value if set, zero value otherwise.
func (o *Pagination) GetTotalFilteredCount() int64 {
	if o == nil || o.TotalFilteredCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalFilteredCount
}

// GetTotalFilteredCountOk returns a tuple with the TotalFilteredCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetTotalFilteredCountOk() (*int64, bool) {
	if o == nil || o.TotalFilteredCount == nil {
		return nil, false
	}
	return o.TotalFilteredCount, true
}

// HasTotalFilteredCount returns a boolean if a field has been set.
func (o *Pagination) HasTotalFilteredCount() bool {
	if o != nil && o.TotalFilteredCount != nil {
		return true
	}

	return false
}

// SetTotalFilteredCount gets a reference to the given int64 and assigns it to the TotalFilteredCount field.
func (o *Pagination) SetTotalFilteredCount(v int64) {
	o.TotalFilteredCount = &v
}

func (o Pagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalCount != nil {
		toSerialize["total_count"] = o.TotalCount
	}
	if o.TotalFilteredCount != nil {
		toSerialize["total_filtered_count"] = o.TotalFilteredCount
	}
	return json.Marshal(toSerialize)
}

type NullablePagination struct {
	value *Pagination
	isSet bool
}

func (v NullablePagination) Get() *Pagination {
	return v.value
}

func (v *NullablePagination) Set(val *Pagination) {
	v.value = val
	v.isSet = true
}

func (v NullablePagination) IsSet() bool {
	return v.isSet
}

func (v *NullablePagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagination(val *Pagination) *NullablePagination {
	return &NullablePagination{value: val, isSet: true}
}

func (v NullablePagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}