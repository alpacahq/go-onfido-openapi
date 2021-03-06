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
	"time"
)

// Report struct for Report
type Report struct {
	// The unique identifier for the report. Read-only.
	Id *string `json:"id,omitempty"`
	// The date and time at which the report was first initiated. Read-only.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The API endpoint to retrieve the report. Read-only.
	Href *string `json:"href,omitempty"`
	// The current state of the report in the checking process. Read-only.
	Status *string `json:"status,omitempty"`
	// The result of the report. Read-only.
	Result *string `json:"result,omitempty"`
	// The sub_result of the report. It gives a more detailed result for document reports only, and will be null otherwise. Read-only.
	SubResult *string `json:"sub_result,omitempty"`
	// The details of the report. This is specific to each type of report. Read-only.
	Breakdown *map[string]interface{} `json:"breakdown,omitempty"`
	// The properties associated with the report, if any. Read-only.
	Properties *map[string]interface{} `json:"properties,omitempty"`
	// The ID of the check to which the report belongs. Read-only.
	CheckId *string `json:"check_id,omitempty"`
	// The name of the report type.
	Name string `json:"name"`
	// Array of objects with document ids that were used in the Onfido engine. [ONLY USED IN A DOCUMENT CHECK]
	Documents *[]ReportDocument `json:"documents,omitempty"`
}

// NewReport instantiates a new Report object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReport(name string) *Report {
	this := Report{}
	this.Name = name
	return &this
}

// NewReportWithDefaults instantiates a new Report object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportWithDefaults() *Report {
	this := Report{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Report) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Report) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Report) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Report) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Report) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Report) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Report) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Report) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Report) SetHref(v string) {
	o.Href = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Report) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Report) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Report) SetStatus(v string) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *Report) GetResult() string {
	if o == nil || o.Result == nil {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetResultOk() (*string, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *Report) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *Report) SetResult(v string) {
	o.Result = &v
}

// GetSubResult returns the SubResult field value if set, zero value otherwise.
func (o *Report) GetSubResult() string {
	if o == nil || o.SubResult == nil {
		var ret string
		return ret
	}
	return *o.SubResult
}

// GetSubResultOk returns a tuple with the SubResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetSubResultOk() (*string, bool) {
	if o == nil || o.SubResult == nil {
		return nil, false
	}
	return o.SubResult, true
}

// HasSubResult returns a boolean if a field has been set.
func (o *Report) HasSubResult() bool {
	if o != nil && o.SubResult != nil {
		return true
	}

	return false
}

// SetSubResult gets a reference to the given string and assigns it to the SubResult field.
func (o *Report) SetSubResult(v string) {
	o.SubResult = &v
}

// GetBreakdown returns the Breakdown field value if set, zero value otherwise.
func (o *Report) GetBreakdown() map[string]interface{} {
	if o == nil || o.Breakdown == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Breakdown
}

// GetBreakdownOk returns a tuple with the Breakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetBreakdownOk() (*map[string]interface{}, bool) {
	if o == nil || o.Breakdown == nil {
		return nil, false
	}
	return o.Breakdown, true
}

// HasBreakdown returns a boolean if a field has been set.
func (o *Report) HasBreakdown() bool {
	if o != nil && o.Breakdown != nil {
		return true
	}

	return false
}

// SetBreakdown gets a reference to the given map[string]interface{} and assigns it to the Breakdown field.
func (o *Report) SetBreakdown(v map[string]interface{}) {
	o.Breakdown = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *Report) GetProperties() map[string]interface{} {
	if o == nil || o.Properties == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetPropertiesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *Report) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *Report) SetProperties(v map[string]interface{}) {
	o.Properties = &v
}

// GetCheckId returns the CheckId field value if set, zero value otherwise.
func (o *Report) GetCheckId() string {
	if o == nil || o.CheckId == nil {
		var ret string
		return ret
	}
	return *o.CheckId
}

// GetCheckIdOk returns a tuple with the CheckId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetCheckIdOk() (*string, bool) {
	if o == nil || o.CheckId == nil {
		return nil, false
	}
	return o.CheckId, true
}

// HasCheckId returns a boolean if a field has been set.
func (o *Report) HasCheckId() bool {
	if o != nil && o.CheckId != nil {
		return true
	}

	return false
}

// SetCheckId gets a reference to the given string and assigns it to the CheckId field.
func (o *Report) SetCheckId(v string) {
	o.CheckId = &v
}

// GetName returns the Name field value
func (o *Report) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Report) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Report) SetName(v string) {
	o.Name = v
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *Report) GetDocuments() []ReportDocument {
	if o == nil || o.Documents == nil {
		var ret []ReportDocument
		return ret
	}
	return *o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetDocumentsOk() (*[]ReportDocument, bool) {
	if o == nil || o.Documents == nil {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *Report) HasDocuments() bool {
	if o != nil && o.Documents != nil {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []ReportDocument and assigns it to the Documents field.
func (o *Report) SetDocuments(v []ReportDocument) {
	o.Documents = &v
}

func (o Report) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.SubResult != nil {
		toSerialize["sub_result"] = o.SubResult
	}
	if o.Breakdown != nil {
		toSerialize["breakdown"] = o.Breakdown
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.CheckId != nil {
		toSerialize["check_id"] = o.CheckId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Documents != nil {
		toSerialize["documents"] = o.Documents
	}
	return json.Marshal(toSerialize)
}

type NullableReport struct {
	value *Report
	isSet bool
}

func (v NullableReport) Get() *Report {
	return v.value
}

func (v *NullableReport) Set(val *Report) {
	v.value = val
	v.isSet = true
}

func (v NullableReport) IsSet() bool {
	return v.isSet
}

func (v *NullableReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReport(val *Report) *NullableReport {
	return &NullableReport{value: val, isSet: true}
}

func (v NullableReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


