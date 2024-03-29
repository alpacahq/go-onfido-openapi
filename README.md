# Go API client for onfido_openapi

The Onfido API is used to submit check requests.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./onfido_openapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.eu.onfido.com/v3.1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**CancelReport**](docs/DefaultApi.md#cancelreport) | **Post** /reports/{report_id}/cancel | This endpoint is for cancelling individual paused reports.
*DefaultApi* | [**CreateApplicant**](docs/DefaultApi.md#createapplicant) | **Post** /applicants | Create Applicant
*DefaultApi* | [**CreateCheck**](docs/DefaultApi.md#createcheck) | **Post** /checks | Create a check
*DefaultApi* | [**CreateWebhook**](docs/DefaultApi.md#createwebhook) | **Post** /webhooks | Create a webhook
*DefaultApi* | [**DeleteWebhook**](docs/DefaultApi.md#deletewebhook) | **Delete** /webhooks/{webhook_id} | Delete a webhook
*DefaultApi* | [**DestroyApplicant**](docs/DefaultApi.md#destroyapplicant) | **Delete** /applicants/{applicant_id} | Delete Applicant
*DefaultApi* | [**DownloadCheck**](docs/DefaultApi.md#downloadcheck) | **Get** /checks/{check_id}/download | Download a PDF of the check
*DefaultApi* | [**DownloadDocument**](docs/DefaultApi.md#downloaddocument) | **Get** /documents/{document_id}/download | Download a documents raw data
*DefaultApi* | [**DownloadLivePhoto**](docs/DefaultApi.md#downloadlivephoto) | **Get** /live_photos/{live_photo_id}/download | Download live photo
*DefaultApi* | [**DownloadLiveVideo**](docs/DefaultApi.md#downloadlivevideo) | **Get** /live_videos/{live_video_id}/download | Download live video
*DefaultApi* | [**DownloadLiveVideoFrame**](docs/DefaultApi.md#downloadlivevideoframe) | **Get** /live_videos/{live_video_id}/frame | Download live video frame
*DefaultApi* | [**EditWebhook**](docs/DefaultApi.md#editwebhook) | **Put** /webhooks/{webhook_id} | Edit a webhook
*DefaultApi* | [**FindAddresses**](docs/DefaultApi.md#findaddresses) | **Get** /addresses/pick | Search for addresses by postcode
*DefaultApi* | [**FindApplicant**](docs/DefaultApi.md#findapplicant) | **Get** /applicants/{applicant_id} | Retrieve Applicant
*DefaultApi* | [**FindCheck**](docs/DefaultApi.md#findcheck) | **Get** /checks/{check_id} | Retrieve a Check
*DefaultApi* | [**FindDocument**](docs/DefaultApi.md#finddocument) | **Get** /documents/{document_id} | A single document can be retrieved by calling this endpoint with the document’s unique identifier.
*DefaultApi* | [**FindLivePhoto**](docs/DefaultApi.md#findlivephoto) | **Get** /live_photos/{live_photo_id} | Retrieve live photo
*DefaultApi* | [**FindLiveVideo**](docs/DefaultApi.md#findlivevideo) | **Get** /live_videos/{live_video_id} | Retrieve live video
*DefaultApi* | [**FindReport**](docs/DefaultApi.md#findreport) | **Get** /reports/{report_id} | A single report can be retrieved using this endpoint with the corresponding unique identifier.
*DefaultApi* | [**FindWebhook**](docs/DefaultApi.md#findwebhook) | **Get** /webhooks/{webhook_id} | Retrieve a Webhook
*DefaultApi* | [**GenerateSdkToken**](docs/DefaultApi.md#generatesdktoken) | **Post** /sdk_token | Generate a SDK token
*DefaultApi* | [**ListApplicants**](docs/DefaultApi.md#listapplicants) | **Get** /applicants | List Applicants
*DefaultApi* | [**ListChecks**](docs/DefaultApi.md#listchecks) | **Get** /checks | Retrieve Checks
*DefaultApi* | [**ListDocuments**](docs/DefaultApi.md#listdocuments) | **Get** /documents | List documents
*DefaultApi* | [**ListLivePhotos**](docs/DefaultApi.md#listlivephotos) | **Get** /live_photos | List live photos
*DefaultApi* | [**ListLiveVideos**](docs/DefaultApi.md#listlivevideos) | **Get** /live_videos | List live videos
*DefaultApi* | [**ListReports**](docs/DefaultApi.md#listreports) | **Get** /reports | All the reports belonging to a particular check can be listed from this endpoint.
*DefaultApi* | [**ListWebhooks**](docs/DefaultApi.md#listwebhooks) | **Get** /webhooks | List webhooks
*DefaultApi* | [**RestoreApplicant**](docs/DefaultApi.md#restoreapplicant) | **Post** /applicants/{applicant_id}/restore | Restore Applicant
*DefaultApi* | [**ResumeCheck**](docs/DefaultApi.md#resumecheck) | **Post** /checks/{check_id}/resume | Resume a Check
*DefaultApi* | [**ResumeReport**](docs/DefaultApi.md#resumereport) | **Post** /reports/{report_id}/resume | This endpoint is for resuming individual paused reports.
*DefaultApi* | [**UpdateApplicant**](docs/DefaultApi.md#updateapplicant) | **Put** /applicants/{applicant_id} | Update Applicant
*DefaultApi* | [**UploadDocument**](docs/DefaultApi.md#uploaddocument) | **Post** /documents | Upload a document
*DefaultApi* | [**UploadLivePhoto**](docs/DefaultApi.md#uploadlivephoto) | **Post** /live_photos | Upload live photo


## Documentation For Models

 - [Address](docs/Address.md)
 - [AddressesList](docs/AddressesList.md)
 - [Applicant](docs/Applicant.md)
 - [ApplicantsList](docs/ApplicantsList.md)
 - [Check](docs/Check.md)
 - [ChecksList](docs/ChecksList.md)
 - [Document](docs/Document.md)
 - [DocumentsList](docs/DocumentsList.md)
 - [Error](docs/Error.md)
 - [ErrorProperties](docs/ErrorProperties.md)
 - [IdNumber](docs/IdNumber.md)
 - [LivePhoto](docs/LivePhoto.md)
 - [LivePhotosList](docs/LivePhotosList.md)
 - [LiveVideo](docs/LiveVideo.md)
 - [LiveVideosList](docs/LiveVideosList.md)
 - [Report](docs/Report.md)
 - [ReportDocument](docs/ReportDocument.md)
 - [ReportsList](docs/ReportsList.md)
 - [SdkToken](docs/SdkToken.md)
 - [Webhook](docs/Webhook.md)
 - [WebhooksList](docs/WebhooksList.md)


## Documentation For Authorization



### Token

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



