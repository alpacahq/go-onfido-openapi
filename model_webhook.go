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

// Webhook struct for Webhook
type Webhook struct {
	// The unique identifier of the webhook. Read-only.
	Id *string `json:"id,omitempty"`
	// Webhook secret token used to sign the webhook's payload. Read-only.
	Token *string `json:"token,omitempty"`
	// The API endpoint to retrieve the webhook. Read-only.
	Href *string `json:"href,omitempty"`
	// The url that will listen to notifications (must be https).
	Url string `json:"url"`
	// Determine if the webhook is active.
	Enabled *bool `json:"enabled,omitempty"`
	// The environments from which the webhook will receive events. Allowed values are “sandbox” and “live”. If the environments parameter is omitted the webhook will receive events from both environments. 
	Environments *[]string `json:"environments,omitempty"`
	// The events that will be published to the webhook. The supported events are: `report.completed`, `report.withdrawn`, `check.completed`, `check.started`, `check.form_opened`, `check.form_completed`. If the events parameter is omitted all the events will be subscribed. 
	Events *[]string `json:"events,omitempty"`
}

// NewWebhook instantiates a new Webhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhook(url string) *Webhook {
	this := Webhook{}
	this.Url = url
	return &this
}

// NewWebhookWithDefaults instantiates a new Webhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookWithDefaults() *Webhook {
	this := Webhook{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Webhook) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Webhook) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Webhook) SetId(v string) {
	o.Id = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *Webhook) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *Webhook) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *Webhook) SetToken(v string) {
	o.Token = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Webhook) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Webhook) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Webhook) SetHref(v string) {
	o.Href = &v
}

// GetUrl returns the Url field value
func (o *Webhook) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Webhook) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Webhook) SetUrl(v string) {
	o.Url = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Webhook) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Webhook) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Webhook) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *Webhook) GetEnvironments() []string {
	if o == nil || o.Environments == nil {
		var ret []string
		return ret
	}
	return *o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetEnvironmentsOk() (*[]string, bool) {
	if o == nil || o.Environments == nil {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *Webhook) HasEnvironments() bool {
	if o != nil && o.Environments != nil {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given []string and assigns it to the Environments field.
func (o *Webhook) SetEnvironments(v []string) {
	o.Environments = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *Webhook) GetEvents() []string {
	if o == nil || o.Events == nil {
		var ret []string
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetEventsOk() (*[]string, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *Webhook) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []string and assigns it to the Events field.
func (o *Webhook) SetEvents(v []string) {
	o.Events = &v
}

func (o Webhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Environments != nil {
		toSerialize["environments"] = o.Environments
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableWebhook struct {
	value *Webhook
	isSet bool
}

func (v NullableWebhook) Get() *Webhook {
	return v.value
}

func (v *NullableWebhook) Set(val *Webhook) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhook(val *Webhook) *NullableWebhook {
	return &NullableWebhook{value: val, isSet: true}
}

func (v NullableWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


