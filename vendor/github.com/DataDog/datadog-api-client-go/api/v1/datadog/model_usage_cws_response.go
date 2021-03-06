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

// UsageCWSResponse Response containing the Cloud Workload Security usage for each hour for a given organization.
type UsageCWSResponse struct {
	// Get hourly usage for Cloud Workload Security.
	Usage *[]UsageCWSHour `json:"usage,omitempty"`
}

// NewUsageCWSResponse instantiates a new UsageCWSResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageCWSResponse() *UsageCWSResponse {
	this := UsageCWSResponse{}
	return &this
}

// NewUsageCWSResponseWithDefaults instantiates a new UsageCWSResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageCWSResponseWithDefaults() *UsageCWSResponse {
	this := UsageCWSResponse{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *UsageCWSResponse) GetUsage() []UsageCWSHour {
	if o == nil || o.Usage == nil {
		var ret []UsageCWSHour
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCWSResponse) GetUsageOk() (*[]UsageCWSHour, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *UsageCWSResponse) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []UsageCWSHour and assigns it to the Usage field.
func (o *UsageCWSResponse) SetUsage(v []UsageCWSHour) {
	o.Usage = &v
}

func (o UsageCWSResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableUsageCWSResponse struct {
	value *UsageCWSResponse
	isSet bool
}

func (v NullableUsageCWSResponse) Get() *UsageCWSResponse {
	return v.value
}

func (v *NullableUsageCWSResponse) Set(val *UsageCWSResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageCWSResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageCWSResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageCWSResponse(val *UsageCWSResponse) *NullableUsageCWSResponse {
	return &NullableUsageCWSResponse{value: val, isSet: true}
}

func (v NullableUsageCWSResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageCWSResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
