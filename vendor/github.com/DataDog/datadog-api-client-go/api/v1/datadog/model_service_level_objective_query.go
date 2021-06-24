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

// ServiceLevelObjectiveQuery A metric SLI query. **Required if type is `metric`**. Note that Datadog only allows the sum by aggregator to be used because this will sum up all request counts instead of averaging them, or taking the max or min of all of those requests.
type ServiceLevelObjectiveQuery struct {
	// A Datadog metric query for total (valid) events.
	Denominator string `json:"denominator"`
	// A Datadog metric query for good events.
	Numerator string `json:"numerator"`
}

// NewServiceLevelObjectiveQuery instantiates a new ServiceLevelObjectiveQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceLevelObjectiveQuery(denominator string, numerator string) *ServiceLevelObjectiveQuery {
	this := ServiceLevelObjectiveQuery{}
	this.Denominator = denominator
	this.Numerator = numerator
	return &this
}

// NewServiceLevelObjectiveQueryWithDefaults instantiates a new ServiceLevelObjectiveQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceLevelObjectiveQueryWithDefaults() *ServiceLevelObjectiveQuery {
	this := ServiceLevelObjectiveQuery{}
	return &this
}

// GetDenominator returns the Denominator field value
func (o *ServiceLevelObjectiveQuery) GetDenominator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Denominator
}

// GetDenominatorOk returns a tuple with the Denominator field value
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjectiveQuery) GetDenominatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Denominator, true
}

// SetDenominator sets field value
func (o *ServiceLevelObjectiveQuery) SetDenominator(v string) {
	o.Denominator = v
}

// GetNumerator returns the Numerator field value
func (o *ServiceLevelObjectiveQuery) GetNumerator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Numerator
}

// GetNumeratorOk returns a tuple with the Numerator field value
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjectiveQuery) GetNumeratorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Numerator, true
}

// SetNumerator sets field value
func (o *ServiceLevelObjectiveQuery) SetNumerator(v string) {
	o.Numerator = v
}

func (o ServiceLevelObjectiveQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["denominator"] = o.Denominator
	}
	if true {
		toSerialize["numerator"] = o.Numerator
	}
	return json.Marshal(toSerialize)
}

func (o *ServiceLevelObjectiveQuery) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Denominator *string `json:"denominator"`
		Numerator   *string `json:"numerator"`
	}{}
	all := struct {
		Denominator string `json:"denominator"`
		Numerator   string `json:"numerator"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Denominator == nil {
		return fmt.Errorf("Required field denominator missing")
	}
	if required.Numerator == nil {
		return fmt.Errorf("Required field numerator missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		return err
	}
	o.Denominator = all.Denominator
	o.Numerator = all.Numerator
	return nil
}

type NullableServiceLevelObjectiveQuery struct {
	value *ServiceLevelObjectiveQuery
	isSet bool
}

func (v NullableServiceLevelObjectiveQuery) Get() *ServiceLevelObjectiveQuery {
	return v.value
}

func (v *NullableServiceLevelObjectiveQuery) Set(val *ServiceLevelObjectiveQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceLevelObjectiveQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceLevelObjectiveQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceLevelObjectiveQuery(val *ServiceLevelObjectiveQuery) *NullableServiceLevelObjectiveQuery {
	return &NullableServiceLevelObjectiveQuery{value: val, isSet: true}
}

func (v NullableServiceLevelObjectiveQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceLevelObjectiveQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
