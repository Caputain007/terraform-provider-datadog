/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// FreeTextWidgetDefinitionType Type of the free text widget.
type FreeTextWidgetDefinitionType string

// List of FreeTextWidgetDefinitionType
const (
	FREETEXTWIDGETDEFINITIONTYPE_FREE_TEXT FreeTextWidgetDefinitionType = "free_text"
)

var allowedFreeTextWidgetDefinitionTypeEnumValues = []FreeTextWidgetDefinitionType{
	"free_text",
}

func (v *FreeTextWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FreeTextWidgetDefinitionType(value)
	for _, existing := range allowedFreeTextWidgetDefinitionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FreeTextWidgetDefinitionType", value)
}

// NewFreeTextWidgetDefinitionTypeFromValue returns a pointer to a valid FreeTextWidgetDefinitionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFreeTextWidgetDefinitionTypeFromValue(v string) (*FreeTextWidgetDefinitionType, error) {
	ev := FreeTextWidgetDefinitionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FreeTextWidgetDefinitionType: valid values are %v", v, allowedFreeTextWidgetDefinitionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FreeTextWidgetDefinitionType) IsValid() bool {
	for _, existing := range allowedFreeTextWidgetDefinitionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FreeTextWidgetDefinitionType value
func (v FreeTextWidgetDefinitionType) Ptr() *FreeTextWidgetDefinitionType {
	return &v
}

type NullableFreeTextWidgetDefinitionType struct {
	value *FreeTextWidgetDefinitionType
	isSet bool
}

func (v NullableFreeTextWidgetDefinitionType) Get() *FreeTextWidgetDefinitionType {
	return v.value
}

func (v *NullableFreeTextWidgetDefinitionType) Set(val *FreeTextWidgetDefinitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableFreeTextWidgetDefinitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableFreeTextWidgetDefinitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreeTextWidgetDefinitionType(val *FreeTextWidgetDefinitionType) *NullableFreeTextWidgetDefinitionType {
	return &NullableFreeTextWidgetDefinitionType{value: val, isSet: true}
}

func (v NullableFreeTextWidgetDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreeTextWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
