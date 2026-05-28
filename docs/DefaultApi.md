# \DefaultApi

All URIs are relative to *https://api.eu.onfido.com/v3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelReport**](DefaultApi.md#CancelReport) | **Post** /reports/{report_id}/cancel | This endpoint is for cancelling individual paused reports.
[**CreateApplicant**](DefaultApi.md#CreateApplicant) | **Post** /applicants | Create Applicant
[**CreateCheck**](DefaultApi.md#CreateCheck) | **Post** /checks | Create a check
[**CreateWebhook**](DefaultApi.md#CreateWebhook) | **Post** /webhooks | Create a webhook
[**CreateWorkflowRun**](DefaultApi.md#CreateWorkflowRun) | **Post** /workflow_runs | Create a Workflow Run
[**DeleteWebhook**](DefaultApi.md#DeleteWebhook) | **Delete** /webhooks/{webhook_id} | Delete a webhook
[**DestroyApplicant**](DefaultApi.md#DestroyApplicant) | **Delete** /applicants/{applicant_id} | Delete Applicant
[**DownloadCheck**](DefaultApi.md#DownloadCheck) | **Get** /checks/{check_id}/download | Download a PDF of the check
[**DownloadDocument**](DefaultApi.md#DownloadDocument) | **Get** /documents/{document_id}/download | Download a documents raw data
[**DownloadEvidenceFolder**](DefaultApi.md#DownloadEvidenceFolder) | **Get** /workflow_runs/{workflow_run_id}/evidence_folder | Retrieve Workflow Run Evidence Folder
[**DownloadLivePhoto**](DefaultApi.md#DownloadLivePhoto) | **Get** /live_photos/{live_photo_id}/download | Download live photo
[**DownloadLiveVideo**](DefaultApi.md#DownloadLiveVideo) | **Get** /live_videos/{live_video_id}/download | Download live video
[**DownloadLiveVideoFrame**](DefaultApi.md#DownloadLiveVideoFrame) | **Get** /live_videos/{live_video_id}/frame | Download live video frame
[**DownloadSignedEvidenceFile**](DefaultApi.md#DownloadSignedEvidenceFile) | **Get** /workflow_runs/{workflow_run_id}/signed_evidence_file | Retrieve Workflow Run Evidence Summary File
[**EditWebhook**](DefaultApi.md#EditWebhook) | **Put** /webhooks/{webhook_id} | Edit a webhook
[**FindAddresses**](DefaultApi.md#FindAddresses) | **Get** /addresses/pick | Search for addresses by postcode
[**FindApplicant**](DefaultApi.md#FindApplicant) | **Get** /applicants/{applicant_id} | Retrieve Applicant
[**FindCheck**](DefaultApi.md#FindCheck) | **Get** /checks/{check_id} | Retrieve a Check
[**FindDocument**](DefaultApi.md#FindDocument) | **Get** /documents/{document_id} | A single document can be retrieved by calling this endpoint with the document’s unique identifier.
[**FindLivePhoto**](DefaultApi.md#FindLivePhoto) | **Get** /live_photos/{live_photo_id} | Retrieve live photo
[**FindLiveVideo**](DefaultApi.md#FindLiveVideo) | **Get** /live_videos/{live_video_id} | Retrieve live video
[**FindReport**](DefaultApi.md#FindReport) | **Get** /reports/{report_id} | A single report can be retrieved using this endpoint with the corresponding unique identifier.
[**FindWebhook**](DefaultApi.md#FindWebhook) | **Get** /webhooks/{webhook_id} | Retrieve a Webhook
[**FindWorkflowRun**](DefaultApi.md#FindWorkflowRun) | **Get** /workflow_runs/{workflow_run_id} | Retrieve Workflow Run
[**GenerateSdkToken**](DefaultApi.md#GenerateSdkToken) | **Post** /sdk_token | Generate a SDK token
[**ListApplicants**](DefaultApi.md#ListApplicants) | **Get** /applicants | List Applicants
[**ListChecks**](DefaultApi.md#ListChecks) | **Get** /checks | Retrieve Checks
[**ListDocuments**](DefaultApi.md#ListDocuments) | **Get** /documents | List documents
[**ListLivePhotos**](DefaultApi.md#ListLivePhotos) | **Get** /live_photos | List live photos
[**ListLiveVideos**](DefaultApi.md#ListLiveVideos) | **Get** /live_videos | List live videos
[**ListReports**](DefaultApi.md#ListReports) | **Get** /reports | All the reports belonging to a particular check can be listed from this endpoint.
[**ListWebhooks**](DefaultApi.md#ListWebhooks) | **Get** /webhooks | List webhooks
[**ListWorkflowRuns**](DefaultApi.md#ListWorkflowRuns) | **Get** /workflow_runs | List Workflow Runs
[**RestoreApplicant**](DefaultApi.md#RestoreApplicant) | **Post** /applicants/{applicant_id}/restore | Restore Applicant
[**ResumeCheck**](DefaultApi.md#ResumeCheck) | **Post** /checks/{check_id}/resume | Resume a Check
[**ResumeReport**](DefaultApi.md#ResumeReport) | **Post** /reports/{report_id}/resume | This endpoint is for resuming individual paused reports.
[**UpdateApplicant**](DefaultApi.md#UpdateApplicant) | **Put** /applicants/{applicant_id} | Update Applicant
[**UploadDocument**](DefaultApi.md#UploadDocument) | **Post** /documents | Upload a document
[**UploadLivePhoto**](DefaultApi.md#UploadLivePhoto) | **Post** /live_photos | Upload live photo



## CancelReport

> CancelReport(ctx, reportId).Execute()

This endpoint is for cancelling individual paused reports.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    reportId := "reportId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CancelReport(context.Background(), reportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CancelReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicant

> Applicant CreateApplicant(ctx).Applicant(applicant).Execute()

Create Applicant

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicant := *openapiclient.NewApplicant() // Applicant | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateApplicant(context.Background()).Applicant(applicant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateApplicant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicant`: Applicant
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateApplicant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicant** | [**Applicant**](Applicant.md) |  | 

### Return type

[**Applicant**](Applicant.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCheck

> Check CreateCheck(ctx).Check(check).Execute()

Create a check

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    check := *openapiclient.NewCheck() // Check | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateCheck(context.Background()).Check(check).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCheck`: Check
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **check** | [**Check**](Check.md) |  | 

### Return type

[**Check**](Check.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebhook

> Webhook CreateWebhook(ctx).Webhook(webhook).Execute()

Create a webhook

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    webhook := *openapiclient.NewWebhook("Url_example") // Webhook | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateWebhook(context.Background()).Webhook(webhook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebhook`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhook** | [**Webhook**](Webhook.md) |  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkflowRun

> WorkflowRun CreateWorkflowRun(ctx).WorkflowRunBuilder(workflowRunBuilder).Execute()

Create a Workflow Run



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowRunBuilder := *openapiclient.NewWorkflowRunBuilder("ApplicantId_example", "WorkflowId_example") // WorkflowRunBuilder | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateWorkflowRun(context.Background()).WorkflowRunBuilder(workflowRunBuilder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateWorkflowRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowRun`: WorkflowRun
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateWorkflowRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowRunBuilder** | [**WorkflowRunBuilder**](WorkflowRunBuilder.md) |  | 

### Return type

[**WorkflowRun**](WorkflowRun.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhook

> DeleteWebhook(ctx, webhookId).Execute()

Delete a webhook

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    webhookId := "webhookId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteWebhook(context.Background(), webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyApplicant

> DestroyApplicant(ctx, applicantId).Execute()

Delete Applicant

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicantId := "applicantId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DestroyApplicant(context.Background(), applicantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DestroyApplicant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyApplicantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadCheck

> DownloadCheck(ctx, checkId).Execute()

Download a PDF of the check

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    checkId := "checkId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DownloadCheck(context.Background(), checkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DownloadCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadDocument

> *os.File DownloadDocument(ctx, documentId).Execute()

Download a documents raw data

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    documentId := "documentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DownloadDocument(context.Background(), documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DownloadDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadDocument`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DownloadDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadEvidenceFolder

> *os.File DownloadEvidenceFolder(ctx, workflowRunId).Execute()

Retrieve Workflow Run Evidence Folder



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowRunId := TODO // string | Workflow Run ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DownloadEvidenceFolder(context.Background(), workflowRunId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DownloadEvidenceFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadEvidenceFolder`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DownloadEvidenceFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowRunId** | [**string**](.md) | Workflow Run ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadEvidenceFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadLivePhoto

> *os.File DownloadLivePhoto(ctx, livePhotoId).Execute()

Download live photo



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    livePhotoId := "livePhotoId_example" // string | The live photo’s unique identifier.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DownloadLivePhoto(context.Background(), livePhotoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DownloadLivePhoto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadLivePhoto`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DownloadLivePhoto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**livePhotoId** | **string** | The live photo’s unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLivePhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadLiveVideo

> *os.File DownloadLiveVideo(ctx, liveVideoId).Execute()

Download live video



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    liveVideoId := "liveVideoId_example" // string | The live video’s unique identifier.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DownloadLiveVideo(context.Background(), liveVideoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DownloadLiveVideo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadLiveVideo`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DownloadLiveVideo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liveVideoId** | **string** | The live video’s unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLiveVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadLiveVideoFrame

> *os.File DownloadLiveVideoFrame(ctx, liveVideoId).Execute()

Download live video frame



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    liveVideoId := "liveVideoId_example" // string | The live video’s unique identifier.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DownloadLiveVideoFrame(context.Background(), liveVideoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DownloadLiveVideoFrame``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadLiveVideoFrame`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DownloadLiveVideoFrame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liveVideoId** | **string** | The live video’s unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLiveVideoFrameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadSignedEvidenceFile

> *os.File DownloadSignedEvidenceFile(ctx, workflowRunId).Execute()

Retrieve Workflow Run Evidence Summary File



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowRunId := TODO // string | Workflow Run ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DownloadSignedEvidenceFile(context.Background(), workflowRunId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DownloadSignedEvidenceFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadSignedEvidenceFile`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DownloadSignedEvidenceFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowRunId** | [**string**](.md) | Workflow Run ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadSignedEvidenceFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditWebhook

> Webhook EditWebhook(ctx, webhookId).Webhook(webhook).Execute()

Edit a webhook

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    webhookId := "webhookId_example" // string | 
    webhook := *openapiclient.NewWebhook("Url_example") // Webhook | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.EditWebhook(context.Background(), webhookId).Webhook(webhook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EditWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditWebhook`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.EditWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webhook** | [**Webhook**](Webhook.md) |  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAddresses

> AddressesList FindAddresses(ctx).Postcode(postcode).Execute()

Search for addresses by postcode

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    postcode := "postcode_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FindAddresses(context.Background()).Postcode(postcode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FindAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAddresses`: AddressesList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FindAddresses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postcode** | **string** |  | 

### Return type

[**AddressesList**](AddressesList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindApplicant

> Applicant FindApplicant(ctx, applicantId).Execute()

Retrieve Applicant

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicantId := "applicantId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FindApplicant(context.Background(), applicantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FindApplicant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindApplicant`: Applicant
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FindApplicant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindApplicantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Applicant**](Applicant.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindCheck

> Check FindCheck(ctx, checkId).Execute()

Retrieve a Check

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    checkId := "checkId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FindCheck(context.Background(), checkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FindCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindCheck`: Check
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FindCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Check**](Check.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindDocument

> Document FindDocument(ctx, documentId).Execute()

A single document can be retrieved by calling this endpoint with the document’s unique identifier.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    documentId := "documentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FindDocument(context.Background(), documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FindDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindDocument`: Document
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FindDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Document**](Document.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindLivePhoto

> LivePhoto FindLivePhoto(ctx, livePhotoId).Execute()

Retrieve live photo

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    livePhotoId := "livePhotoId_example" // string | The live photo’s unique identifier.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FindLivePhoto(context.Background(), livePhotoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FindLivePhoto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindLivePhoto`: LivePhoto
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FindLivePhoto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**livePhotoId** | **string** | The live photo’s unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindLivePhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LivePhoto**](LivePhoto.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindLiveVideo

> LiveVideo FindLiveVideo(ctx, liveVideoId).Execute()

Retrieve live video

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    liveVideoId := "liveVideoId_example" // string | The live video’s unique identifier.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FindLiveVideo(context.Background(), liveVideoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FindLiveVideo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindLiveVideo`: LiveVideo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FindLiveVideo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liveVideoId** | **string** | The live video’s unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindLiveVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LiveVideo**](LiveVideo.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindReport

> Report FindReport(ctx, reportId).Execute()

A single report can be retrieved using this endpoint with the corresponding unique identifier.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    reportId := "reportId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FindReport(context.Background(), reportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FindReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindReport`: Report
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FindReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Report**](Report.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindWebhook

> Webhook FindWebhook(ctx, webhookId).Execute()

Retrieve a Webhook

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    webhookId := "webhookId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FindWebhook(context.Background(), webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FindWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindWebhook`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FindWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Webhook**](Webhook.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindWorkflowRun

> WorkflowRun FindWorkflowRun(ctx, workflowRunId).Execute()

Retrieve Workflow Run



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workflowRunId := TODO // string | The unique identifier of the Workflow Run.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FindWorkflowRun(context.Background(), workflowRunId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FindWorkflowRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindWorkflowRun`: WorkflowRun
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FindWorkflowRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowRunId** | [**string**](.md) | The unique identifier of the Workflow Run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindWorkflowRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowRun**](WorkflowRun.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateSdkToken

> SdkToken GenerateSdkToken(ctx).SdkToken(sdkToken).Execute()

Generate a SDK token

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sdkToken := *openapiclient.NewSdkToken() // SdkToken | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GenerateSdkToken(context.Background()).SdkToken(sdkToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GenerateSdkToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateSdkToken`: SdkToken
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GenerateSdkToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateSdkTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sdkToken** | [**SdkToken**](SdkToken.md) |  | 

### Return type

[**SdkToken**](SdkToken.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicants

> ApplicantsList ListApplicants(ctx).Page(page).PerPage(perPage).IncludeDeleted(includeDeleted).Execute()

List Applicants

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 | The page to return. The first page is `page=1` (optional) (default to 1)
    perPage := int32(56) // int32 | The number of objects per page. (optional) (default to 20)
    includeDeleted := true // bool | Whether to also include applicants scheduled for deletion. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListApplicants(context.Background()).Page(page).PerPage(perPage).IncludeDeleted(includeDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListApplicants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicants`: ApplicantsList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListApplicants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApplicantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page to return. The first page is &#x60;page&#x3D;1&#x60; | [default to 1]
 **perPage** | **int32** | The number of objects per page. | [default to 20]
 **includeDeleted** | **bool** | Whether to also include applicants scheduled for deletion. | [default to false]

### Return type

[**ApplicantsList**](ApplicantsList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChecks

> ChecksList ListChecks(ctx).ApplicantId(applicantId).Execute()

Retrieve Checks

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicantId := "applicantId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListChecks(context.Background()).ApplicantId(applicantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListChecks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListChecks`: ChecksList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListChecks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicantId** | **string** |  | 

### Return type

[**ChecksList**](ChecksList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDocuments

> DocumentsList ListDocuments(ctx).ApplicantId(applicantId).Execute()

List documents



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicantId := "applicantId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListDocuments(context.Background()).ApplicantId(applicantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDocuments`: DocumentsList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicantId** | **string** |  | 

### Return type

[**DocumentsList**](DocumentsList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLivePhotos

> LivePhotosList ListLivePhotos(ctx).ApplicantId(applicantId).Execute()

List live photos

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicantId := "applicantId_example" // string | The id of the applicant the live photos belong to.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListLivePhotos(context.Background()).ApplicantId(applicantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListLivePhotos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLivePhotos`: LivePhotosList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListLivePhotos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLivePhotosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicantId** | **string** | The id of the applicant the live photos belong to. | 

### Return type

[**LivePhotosList**](LivePhotosList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLiveVideos

> LiveVideosList ListLiveVideos(ctx).ApplicantId(applicantId).Execute()

List live videos

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicantId := "applicantId_example" // string | The id of the applicant the live videos belong to.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListLiveVideos(context.Background()).ApplicantId(applicantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListLiveVideos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLiveVideos`: LiveVideosList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListLiveVideos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLiveVideosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicantId** | **string** | The id of the applicant the live videos belong to. | 

### Return type

[**LiveVideosList**](LiveVideosList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReports

> ReportsList ListReports(ctx).CheckId(checkId).Execute()

All the reports belonging to a particular check can be listed from this endpoint.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    checkId := "checkId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListReports(context.Background()).CheckId(checkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListReports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListReports`: ReportsList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListReports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkId** | **string** |  | 

### Return type

[**ReportsList**](ReportsList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhooks

> WebhooksList ListWebhooks(ctx).Execute()

List webhooks

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListWebhooks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWebhooks`: WebhooksList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListWebhooks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWebhooksRequest struct via the builder pattern


### Return type

[**WebhooksList**](WebhooksList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflowRuns

> []WorkflowRun ListWorkflowRuns(ctx).Page(page).Status(status).CreatedAtGt(createdAtGt).CreatedAtLt(createdAtLt).Sort(sort).ApplicantId(applicantId).Tags(tags).Execute()

List Workflow Runs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 | The number of the page to be retrieved. If not specified, defaults to 1. (optional) (default to 1)
    status := "status_example" // string | A list of comma separated status values to filter the results. Possible values are 'processing', 'awaiting_input', 'approved', 'declined', 'review', 'abandoned' and 'error'. (optional)
    createdAtGt := time.Now() // string | A ISO-8601 date to filter results with a created date greater than (after) the one provided. (optional)
    createdAtLt := time.Now() // string | A ISO-8601 date to filter results with a created date less than (before) the one provided. (optional)
    sort := "sort_example" // string | A string with the value 'desc' or 'asc' that allows to sort the returned list by the completed datetime either descending or ascending, respectively. If not specified, defaults to 'desc'. (optional) (default to "desc")
    applicantId := TODO // string | the applicant's id. (optional)
    tags := []string{"Inner_example"} // []string | A list of tags to filter the results. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListWorkflowRuns(context.Background()).Page(page).Status(status).CreatedAtGt(createdAtGt).CreatedAtLt(createdAtLt).Sort(sort).ApplicantId(applicantId).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListWorkflowRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflowRuns`: []WorkflowRun
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListWorkflowRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The number of the page to be retrieved. If not specified, defaults to 1. | [default to 1]
 **status** | **string** | A list of comma separated status values to filter the results. Possible values are &#39;processing&#39;, &#39;awaiting_input&#39;, &#39;approved&#39;, &#39;declined&#39;, &#39;review&#39;, &#39;abandoned&#39; and &#39;error&#39;. | 
 **createdAtGt** | **string** | A ISO-8601 date to filter results with a created date greater than (after) the one provided. | 
 **createdAtLt** | **string** | A ISO-8601 date to filter results with a created date less than (before) the one provided. | 
 **sort** | **string** | A string with the value &#39;desc&#39; or &#39;asc&#39; that allows to sort the returned list by the completed datetime either descending or ascending, respectively. If not specified, defaults to &#39;desc&#39;. | [default to &quot;desc&quot;]
 **applicantId** | [**string**](string.md) | the applicant&#39;s id. | 
 **tags** | **[]string** | A list of tags to filter the results. | 

### Return type

[**[]WorkflowRun**](WorkflowRun.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreApplicant

> RestoreApplicant(ctx, applicantId).Execute()

Restore Applicant

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicantId := "applicantId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RestoreApplicant(context.Background(), applicantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RestoreApplicant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreApplicantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeCheck

> ResumeCheck(ctx, checkId).Execute()

Resume a Check

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    checkId := "checkId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ResumeCheck(context.Background(), checkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ResumeCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeReport

> ResumeReport(ctx, reportId).Execute()

This endpoint is for resuming individual paused reports.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    reportId := "reportId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ResumeReport(context.Background(), reportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ResumeReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicant

> Applicant UpdateApplicant(ctx, applicantId).Applicant(applicant).Execute()

Update Applicant



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicantId := "applicantId_example" // string | 
    applicant := *openapiclient.NewApplicant() // Applicant | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateApplicant(context.Background(), applicantId).Applicant(applicant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateApplicant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplicant`: Applicant
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateApplicant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicant** | [**Applicant**](Applicant.md) |  | 

### Return type

[**Applicant**](Applicant.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadDocument

> Document UploadDocument(ctx).InlineObject(inlineObject).Execute()

Upload a document



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inlineObject := TODO // InlineObject | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UploadDocument(context.Background()).InlineObject(inlineObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UploadDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadDocument`: Document
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UploadDocument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject** | [**InlineObject**](InlineObject.md) |  | 

### Return type

[**Document**](Document.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadLivePhoto

> LivePhoto UploadLivePhoto(ctx).InlineObject1(inlineObject1).Execute()

Upload live photo



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inlineObject1 := TODO // InlineObject1 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UploadLivePhoto(context.Background()).InlineObject1(inlineObject1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UploadLivePhoto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadLivePhoto`: LivePhoto
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UploadLivePhoto`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadLivePhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject1** | [**InlineObject1**](InlineObject1.md) |  | 

### Return type

[**LivePhoto**](LivePhoto.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

