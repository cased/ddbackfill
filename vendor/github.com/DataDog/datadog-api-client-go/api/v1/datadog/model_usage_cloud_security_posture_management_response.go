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

// UsageCloudSecurityPostureManagementResponse The response containing the Cloud Security Posture Management usage for each hour for a given organization.
type UsageCloudSecurityPostureManagementResponse struct {
	// Get hourly usage for Cloud Security Posture Management.
	Usage *[]UsageCloudSecurityPostureManagementHour `json:"usage,omitempty"`
}

// NewUsageCloudSecurityPostureManagementResponse instantiates a new UsageCloudSecurityPostureManagementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageCloudSecurityPostureManagementResponse() *UsageCloudSecurityPostureManagementResponse {
	this := UsageCloudSecurityPostureManagementResponse{}
	return &this
}

// NewUsageCloudSecurityPostureManagementResponseWithDefaults instantiates a new UsageCloudSecurityPostureManagementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageCloudSecurityPostureManagementResponseWithDefaults() *UsageCloudSecurityPostureManagementResponse {
	this := UsageCloudSecurityPostureManagementResponse{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *UsageCloudSecurityPostureManagementResponse) GetUsage() []UsageCloudSecurityPostureManagementHour {
	if o == nil || o.Usage == nil {
		var ret []UsageCloudSecurityPostureManagementHour
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCloudSecurityPostureManagementResponse) GetUsageOk() (*[]UsageCloudSecurityPostureManagementHour, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *UsageCloudSecurityPostureManagementResponse) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []UsageCloudSecurityPostureManagementHour and assigns it to the Usage field.
func (o *UsageCloudSecurityPostureManagementResponse) SetUsage(v []UsageCloudSecurityPostureManagementHour) {
	o.Usage = &v
}

func (o UsageCloudSecurityPostureManagementResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableUsageCloudSecurityPostureManagementResponse struct {
	value *UsageCloudSecurityPostureManagementResponse
	isSet bool
}

func (v NullableUsageCloudSecurityPostureManagementResponse) Get() *UsageCloudSecurityPostureManagementResponse {
	return v.value
}

func (v *NullableUsageCloudSecurityPostureManagementResponse) Set(val *UsageCloudSecurityPostureManagementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageCloudSecurityPostureManagementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageCloudSecurityPostureManagementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageCloudSecurityPostureManagementResponse(val *UsageCloudSecurityPostureManagementResponse) *NullableUsageCloudSecurityPostureManagementResponse {
	return &NullableUsageCloudSecurityPostureManagementResponse{value: val, isSet: true}
}

func (v NullableUsageCloudSecurityPostureManagementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageCloudSecurityPostureManagementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
