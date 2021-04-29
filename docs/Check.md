# Check

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier for the check. Read-only. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | The date and time when this check was created. Read-only. | [optional] [readonly] 
**Href** | Pointer to **string** | The uri of this resource. Read-only. | [optional] [readonly] 
**Status** | Pointer to **string** | The current state of the check in the checking process. Read-only. | [optional] [readonly] 
**Result** | Pointer to **string** | The overall result of the check, based on the results of the constituent reports. Read-only. | [optional] [readonly] 
**DownloadUri** | Pointer to **string** | A link to a PDF output of the check results. Append &#x60;.pdf&#x60; to get the pdf file. Read-only. | [optional] [readonly] 
**FormUri** | Pointer to **string** | A link to the applicant form, if &#x60;applicant_provides_data&#x60; is &#x60;true&#x60;. Read-only. | [optional] [readonly] 
**RedirectUri** | Pointer to **string** | For checks where &#x60;applicant_provides_data&#x60; is &#x60;true&#x60;, redirect to this URI when the applicant has submitted their data. Read-only. | [optional] [readonly] 
**ResultsUri** | Pointer to **string** | A link to the corresponding results page on the Onfido dashboard. | [optional] [readonly] 
**ReportNames** | Pointer to **[]string** | An array of report names (strings). | [optional] 
**ApplicantId** | Pointer to **string** | The ID of the applicant to do the check on. | [optional] 
**PrivacyNoticesReadConsentGiven** | Pointer to **bool** | Indicates whether the privacy notices and terms of service have been read and, where specific laws require, that consent has been given for Onfido. | [optional] 
**Tags** | Pointer to **[]string** | Array of tags being assigned to this check. | [optional] 
**ApplicantProvidesData** | Pointer to **bool** | Send an applicant form to applicant to complete to proceed with check. Defaults to false.  | [optional] 
**SuppressFormEmails** | Pointer to **bool** | For checks where &#x60;applicant_provides_data&#x60; is &#x60;true&#x60;, applicant form will not be automatically sent if &#x60;suppress_form_emails&#x60; is set to &#x60;true&#x60;. You can manually send the form at any time after the check has been created, using the link found in the form_uri attribute of the check object. Write-only. Defaults to false.  | [optional] 
**Asynchronous** | Pointer to **bool** | Defaults to &#x60;true&#x60;. Write-only. If set to &#x60;false&#x60;, you will only receive a response when all reports in your check have completed.  | [optional] 
**WebhookIds** | Pointer to **[]string** | Optional. An array of strings describing which webhooks to trigger for this check. By default, all webhooks registered in the account will be triggered and this value will be null in the responses. | [optional] 
**ReportIds** | Pointer to **[]string** | An array of report ids. | [optional] 
**DocumentIds** | Pointer to **[]string** | Optional. An array of document ids, for use with Document reports only. If omitted, the Document report will use the most recently uploaded document by default. | [optional] 
**Consider** | Pointer to **[]string** | Returns a pre-determined consider sub-result in sandbox for the specific reports in the consider array. | [optional] 
**SubResult** | Pointer to **string** | Triggers a pre-determined sub-result response for sandbox Document reports. | [optional] 

## Methods

### NewCheck

`func NewCheck() *Check`

NewCheck instantiates a new Check object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckWithDefaults

`func NewCheckWithDefaults() *Check`

NewCheckWithDefaults instantiates a new Check object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Check) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Check) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Check) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Check) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Check) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Check) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Check) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Check) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHref

`func (o *Check) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Check) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Check) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Check) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetStatus

`func (o *Check) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Check) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Check) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Check) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *Check) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Check) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Check) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *Check) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetDownloadUri

`func (o *Check) GetDownloadUri() string`

GetDownloadUri returns the DownloadUri field if non-nil, zero value otherwise.

### GetDownloadUriOk

`func (o *Check) GetDownloadUriOk() (*string, bool)`

GetDownloadUriOk returns a tuple with the DownloadUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUri

`func (o *Check) SetDownloadUri(v string)`

SetDownloadUri sets DownloadUri field to given value.

### HasDownloadUri

`func (o *Check) HasDownloadUri() bool`

HasDownloadUri returns a boolean if a field has been set.

### GetFormUri

`func (o *Check) GetFormUri() string`

GetFormUri returns the FormUri field if non-nil, zero value otherwise.

### GetFormUriOk

`func (o *Check) GetFormUriOk() (*string, bool)`

GetFormUriOk returns a tuple with the FormUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormUri

`func (o *Check) SetFormUri(v string)`

SetFormUri sets FormUri field to given value.

### HasFormUri

`func (o *Check) HasFormUri() bool`

HasFormUri returns a boolean if a field has been set.

### GetRedirectUri

`func (o *Check) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *Check) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *Check) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *Check) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetResultsUri

`func (o *Check) GetResultsUri() string`

GetResultsUri returns the ResultsUri field if non-nil, zero value otherwise.

### GetResultsUriOk

`func (o *Check) GetResultsUriOk() (*string, bool)`

GetResultsUriOk returns a tuple with the ResultsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsUri

`func (o *Check) SetResultsUri(v string)`

SetResultsUri sets ResultsUri field to given value.

### HasResultsUri

`func (o *Check) HasResultsUri() bool`

HasResultsUri returns a boolean if a field has been set.

### GetReportNames

`func (o *Check) GetReportNames() []string`

GetReportNames returns the ReportNames field if non-nil, zero value otherwise.

### GetReportNamesOk

`func (o *Check) GetReportNamesOk() (*[]string, bool)`

GetReportNamesOk returns a tuple with the ReportNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportNames

`func (o *Check) SetReportNames(v []string)`

SetReportNames sets ReportNames field to given value.

### HasReportNames

`func (o *Check) HasReportNames() bool`

HasReportNames returns a boolean if a field has been set.

### GetApplicantId

`func (o *Check) GetApplicantId() string`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *Check) GetApplicantIdOk() (*string, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *Check) SetApplicantId(v string)`

SetApplicantId sets ApplicantId field to given value.

### HasApplicantId

`func (o *Check) HasApplicantId() bool`

HasApplicantId returns a boolean if a field has been set.

### GetPrivacyNoticesReadConsentGiven

`func (o *Check) GetPrivacyNoticesReadConsentGiven() bool`

GetPrivacyNoticesReadConsentGiven returns the PrivacyNoticesReadConsentGiven field if non-nil, zero value otherwise.

### GetPrivacyNoticesReadConsentGivenOk

`func (o *Check) GetPrivacyNoticesReadConsentGivenOk() (*bool, bool)`

GetPrivacyNoticesReadConsentGivenOk returns a tuple with the PrivacyNoticesReadConsentGiven field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyNoticesReadConsentGiven

`func (o *Check) SetPrivacyNoticesReadConsentGiven(v bool)`

SetPrivacyNoticesReadConsentGiven sets PrivacyNoticesReadConsentGiven field to given value.

### HasPrivacyNoticesReadConsentGiven

`func (o *Check) HasPrivacyNoticesReadConsentGiven() bool`

HasPrivacyNoticesReadConsentGiven returns a boolean if a field has been set.

### GetTags

`func (o *Check) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Check) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Check) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Check) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetApplicantProvidesData

`func (o *Check) GetApplicantProvidesData() bool`

GetApplicantProvidesData returns the ApplicantProvidesData field if non-nil, zero value otherwise.

### GetApplicantProvidesDataOk

`func (o *Check) GetApplicantProvidesDataOk() (*bool, bool)`

GetApplicantProvidesDataOk returns a tuple with the ApplicantProvidesData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantProvidesData

`func (o *Check) SetApplicantProvidesData(v bool)`

SetApplicantProvidesData sets ApplicantProvidesData field to given value.

### HasApplicantProvidesData

`func (o *Check) HasApplicantProvidesData() bool`

HasApplicantProvidesData returns a boolean if a field has been set.

### GetSuppressFormEmails

`func (o *Check) GetSuppressFormEmails() bool`

GetSuppressFormEmails returns the SuppressFormEmails field if non-nil, zero value otherwise.

### GetSuppressFormEmailsOk

`func (o *Check) GetSuppressFormEmailsOk() (*bool, bool)`

GetSuppressFormEmailsOk returns a tuple with the SuppressFormEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressFormEmails

`func (o *Check) SetSuppressFormEmails(v bool)`

SetSuppressFormEmails sets SuppressFormEmails field to given value.

### HasSuppressFormEmails

`func (o *Check) HasSuppressFormEmails() bool`

HasSuppressFormEmails returns a boolean if a field has been set.

### GetAsynchronous

`func (o *Check) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *Check) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *Check) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *Check) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetWebhookIds

`func (o *Check) GetWebhookIds() []string`

GetWebhookIds returns the WebhookIds field if non-nil, zero value otherwise.

### GetWebhookIdsOk

`func (o *Check) GetWebhookIdsOk() (*[]string, bool)`

GetWebhookIdsOk returns a tuple with the WebhookIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookIds

`func (o *Check) SetWebhookIds(v []string)`

SetWebhookIds sets WebhookIds field to given value.

### HasWebhookIds

`func (o *Check) HasWebhookIds() bool`

HasWebhookIds returns a boolean if a field has been set.

### GetReportIds

`func (o *Check) GetReportIds() []string`

GetReportIds returns the ReportIds field if non-nil, zero value otherwise.

### GetReportIdsOk

`func (o *Check) GetReportIdsOk() (*[]string, bool)`

GetReportIdsOk returns a tuple with the ReportIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportIds

`func (o *Check) SetReportIds(v []string)`

SetReportIds sets ReportIds field to given value.

### HasReportIds

`func (o *Check) HasReportIds() bool`

HasReportIds returns a boolean if a field has been set.

### GetDocumentIds

`func (o *Check) GetDocumentIds() []string`

GetDocumentIds returns the DocumentIds field if non-nil, zero value otherwise.

### GetDocumentIdsOk

`func (o *Check) GetDocumentIdsOk() (*[]string, bool)`

GetDocumentIdsOk returns a tuple with the DocumentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIds

`func (o *Check) SetDocumentIds(v []string)`

SetDocumentIds sets DocumentIds field to given value.

### HasDocumentIds

`func (o *Check) HasDocumentIds() bool`

HasDocumentIds returns a boolean if a field has been set.

### GetConsider

`func (o *Check) GetConsider() []string`

GetConsider returns the Consider field if non-nil, zero value otherwise.

### GetConsiderOk

`func (o *Check) GetConsiderOk() (*[]string, bool)`

GetConsiderOk returns a tuple with the Consider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsider

`func (o *Check) SetConsider(v []string)`

SetConsider sets Consider field to given value.

### HasConsider

`func (o *Check) HasConsider() bool`

HasConsider returns a boolean if a field has been set.

### GetSubResult

`func (o *Check) GetSubResult() string`

GetSubResult returns the SubResult field if non-nil, zero value otherwise.

### GetSubResultOk

`func (o *Check) GetSubResultOk() (*string, bool)`

GetSubResultOk returns a tuple with the SubResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubResult

`func (o *Check) SetSubResult(v string)`

SetSubResult sets SubResult field to given value.

### HasSubResult

`func (o *Check) HasSubResult() bool`

HasSubResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


