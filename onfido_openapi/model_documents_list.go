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

// DocumentsList struct for DocumentsList
type DocumentsList struct {
	Documents *[]Document `json:"documents,omitempty"`
}

// NewDocumentsList instantiates a new DocumentsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentsList() *DocumentsList {
	this := DocumentsList{}
	return &this
}

// NewDocumentsListWithDefaults instantiates a new DocumentsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentsListWithDefaults() *DocumentsList {
	this := DocumentsList{}
	return &this
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *DocumentsList) GetDocuments() []Document {
	if o == nil || o.Documents == nil {
		var ret []Document
		return ret
	}
	return *o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentsList) GetDocumentsOk() (*[]Document, bool) {
	if o == nil || o.Documents == nil {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *DocumentsList) HasDocuments() bool {
	if o != nil && o.Documents != nil {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []Document and assigns it to the Documents field.
func (o *DocumentsList) SetDocuments(v []Document) {
	o.Documents = &v
}

func (o DocumentsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Documents != nil {
		toSerialize["documents"] = o.Documents
	}
	return json.Marshal(toSerialize)
}

type NullableDocumentsList struct {
	value *DocumentsList
	isSet bool
}

func (v NullableDocumentsList) Get() *DocumentsList {
	return v.value
}

func (v *NullableDocumentsList) Set(val *DocumentsList) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentsList) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentsList(val *DocumentsList) *NullableDocumentsList {
	return &NullableDocumentsList{value: val, isSet: true}
}

func (v NullableDocumentsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


