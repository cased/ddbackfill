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

// WidgetServiceSummaryDisplayFormat Number of columns to display.
type WidgetServiceSummaryDisplayFormat string

// List of WidgetServiceSummaryDisplayFormat
const (
	WIDGETSERVICESUMMARYDISPLAYFORMAT_ONE_COLUMN   WidgetServiceSummaryDisplayFormat = "one_column"
	WIDGETSERVICESUMMARYDISPLAYFORMAT_TWO_COLUMN   WidgetServiceSummaryDisplayFormat = "two_column"
	WIDGETSERVICESUMMARYDISPLAYFORMAT_THREE_COLUMN WidgetServiceSummaryDisplayFormat = "three_column"
)

var allowedWidgetServiceSummaryDisplayFormatEnumValues = []WidgetServiceSummaryDisplayFormat{
	"one_column",
	"two_column",
	"three_column",
}

func (w *WidgetServiceSummaryDisplayFormat) GetAllowedValues() []WidgetServiceSummaryDisplayFormat {
	return allowedWidgetServiceSummaryDisplayFormatEnumValues
}

func (v *WidgetServiceSummaryDisplayFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WidgetServiceSummaryDisplayFormat(value)
	for _, existing := range allowedWidgetServiceSummaryDisplayFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WidgetServiceSummaryDisplayFormat", value)
}

// NewWidgetServiceSummaryDisplayFormatFromValue returns a pointer to a valid WidgetServiceSummaryDisplayFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWidgetServiceSummaryDisplayFormatFromValue(v string) (*WidgetServiceSummaryDisplayFormat, error) {
	ev := WidgetServiceSummaryDisplayFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WidgetServiceSummaryDisplayFormat: valid values are %v", v, allowedWidgetServiceSummaryDisplayFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WidgetServiceSummaryDisplayFormat) IsValid() bool {
	for _, existing := range allowedWidgetServiceSummaryDisplayFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetServiceSummaryDisplayFormat value
func (v WidgetServiceSummaryDisplayFormat) Ptr() *WidgetServiceSummaryDisplayFormat {
	return &v
}

type NullableWidgetServiceSummaryDisplayFormat struct {
	value *WidgetServiceSummaryDisplayFormat
	isSet bool
}

func (v NullableWidgetServiceSummaryDisplayFormat) Get() *WidgetServiceSummaryDisplayFormat {
	return v.value
}

func (v *NullableWidgetServiceSummaryDisplayFormat) Set(val *WidgetServiceSummaryDisplayFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetServiceSummaryDisplayFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetServiceSummaryDisplayFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetServiceSummaryDisplayFormat(val *WidgetServiceSummaryDisplayFormat) *NullableWidgetServiceSummaryDisplayFormat {
	return &NullableWidgetServiceSummaryDisplayFormat{value: val, isSet: true}
}

func (v NullableWidgetServiceSummaryDisplayFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetServiceSummaryDisplayFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
