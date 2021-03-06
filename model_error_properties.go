/*
 * Onfido API
 *
 * The Onfido API is used to submit check requests.
 *
 * API version: 3.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onfido_openapi

import (
	"encoding/json"
)

// ErrorProperties struct for ErrorProperties
type ErrorProperties struct {
	Type *string `json:"type,omitempty"`
	Message *string `json:"message,omitempty"`
	Fields *map[string]interface{} `json:"fields,omitempty"`
}

// NewErrorProperties instantiates a new ErrorProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorProperties() *ErrorProperties {
	this := ErrorProperties{}
	return &this
}

// NewErrorPropertiesWithDefaults instantiates a new ErrorProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorPropertiesWithDefaults() *ErrorProperties {
	this := ErrorProperties{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ErrorProperties) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorProperties) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ErrorProperties) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ErrorProperties) SetType(v string) {
	o.Type = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorProperties) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorProperties) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorProperties) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorProperties) SetMessage(v string) {
	o.Message = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *ErrorProperties) GetFields() map[string]interface{} {
	if o == nil || o.Fields == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorProperties) GetFieldsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *ErrorProperties) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string]interface{} and assigns it to the Fields field.
func (o *ErrorProperties) SetFields(v map[string]interface{}) {
	o.Fields = &v
}

func (o ErrorProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	return json.Marshal(toSerialize)
}

type NullableErrorProperties struct {
	value *ErrorProperties
	isSet bool
}

func (v NullableErrorProperties) Get() *ErrorProperties {
	return v.value
}

func (v *NullableErrorProperties) Set(val *ErrorProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorProperties(val *ErrorProperties) *NullableErrorProperties {
	return &NullableErrorProperties{value: val, isSet: true}
}

func (v NullableErrorProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


