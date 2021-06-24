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

// SyntheticsBasicAuth Object to handle basic authentication when performing the test.
type SyntheticsBasicAuth struct {
	// Password to use for the basic authentication.
	Password string `json:"password"`
	// Username to use for the basic authentication.
	Username string `json:"username"`
}

// NewSyntheticsBasicAuth instantiates a new SyntheticsBasicAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsBasicAuth(password string, username string) *SyntheticsBasicAuth {
	this := SyntheticsBasicAuth{}
	this.Password = password
	this.Username = username
	return &this
}

// NewSyntheticsBasicAuthWithDefaults instantiates a new SyntheticsBasicAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsBasicAuthWithDefaults() *SyntheticsBasicAuth {
	this := SyntheticsBasicAuth{}
	return &this
}

// GetPassword returns the Password field value
func (o *SyntheticsBasicAuth) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *SyntheticsBasicAuth) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *SyntheticsBasicAuth) SetPassword(v string) {
	o.Password = v
}

// GetUsername returns the Username field value
func (o *SyntheticsBasicAuth) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *SyntheticsBasicAuth) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *SyntheticsBasicAuth) SetUsername(v string) {
	o.Username = v
}

func (o SyntheticsBasicAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

func (o *SyntheticsBasicAuth) UnmarshalJSON(bytes []byte) (err error) {
	required := struct {
		Password *string `json:"password"`
		Username *string `json:"username"`
	}{}
	all := struct {
		Password string `json:"password"`
		Username string `json:"username"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Password == nil {
		return fmt.Errorf("Required field password missing")
	}
	if required.Username == nil {
		return fmt.Errorf("Required field username missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		return err
	}
	o.Password = all.Password
	o.Username = all.Username
	return nil
}

type NullableSyntheticsBasicAuth struct {
	value *SyntheticsBasicAuth
	isSet bool
}

func (v NullableSyntheticsBasicAuth) Get() *SyntheticsBasicAuth {
	return v.value
}

func (v *NullableSyntheticsBasicAuth) Set(val *SyntheticsBasicAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsBasicAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsBasicAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsBasicAuth(val *SyntheticsBasicAuth) *NullableSyntheticsBasicAuth {
	return &NullableSyntheticsBasicAuth{value: val, isSet: true}
}

func (v NullableSyntheticsBasicAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsBasicAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
