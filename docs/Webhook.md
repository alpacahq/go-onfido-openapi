# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the webhook. Read-only. | [optional] [readonly] 
**Token** | Pointer to **string** | Webhook secret token used to sign the webhook&#39;s payload. Read-only. | [optional] [readonly] 
**Href** | Pointer to **string** | The API endpoint to retrieve the webhook. Read-only. | [optional] [readonly] 
**Url** | **string** | The url that will listen to notifications (must be https). | 
**Enabled** | Pointer to **bool** | Determine if the webhook is active. | [optional] 
**Environments** | Pointer to **[]string** | The environments from which the webhook will receive events. Allowed values are “sandbox” and “live”. If the environments parameter is omitted the webhook will receive events from both environments.  | [optional] 
**Events** | Pointer to **[]string** | The events that will be published to the webhook. The supported events are: &#x60;report.completed&#x60;, &#x60;report.withdrawn&#x60;, &#x60;check.completed&#x60;, &#x60;check.started&#x60;, &#x60;check.form_opened&#x60;, &#x60;check.form_completed&#x60;. If the events parameter is omitted all the events will be subscribed.  | [optional] 

## Methods

### NewWebhook

`func NewWebhook(url string, ) *Webhook`

NewWebhook instantiates a new Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithDefaults

`func NewWebhookWithDefaults() *Webhook`

NewWebhookWithDefaults instantiates a new Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Webhook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Webhook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Webhook) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Webhook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetToken

`func (o *Webhook) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Webhook) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Webhook) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *Webhook) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetHref

`func (o *Webhook) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Webhook) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Webhook) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Webhook) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUrl

`func (o *Webhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Webhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Webhook) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEnabled

`func (o *Webhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Webhook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Webhook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Webhook) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEnvironments

`func (o *Webhook) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *Webhook) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *Webhook) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *Webhook) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetEvents

`func (o *Webhook) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Webhook) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Webhook) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *Webhook) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


