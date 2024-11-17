# \UploadAPI

All URIs are relative to *https://www.myminifactory.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilePost_0**](UploadAPI.md#FilePost_0) | **Post** /file | 
[**ImagePost_0**](UploadAPI.md#ImagePost_0) | **Post** /image | 
[**ObjectPatch_0**](UploadAPI.md#ObjectPatch_0) | **Patch** /object | 
[**ObjectPost_0**](UploadAPI.md#ObjectPost_0) | **Post** /object | 
[**ObjectsObjectIdUploadStatusGet_0**](UploadAPI.md#ObjectsObjectIdUploadStatusGet_0) | **Get** /objects/{object_id}/upload_status | 



## FilePost_0

> FilePost201Response FilePost_0(ctx).UploadId(uploadId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uploadId := "uploadId_example" // string | Temporary identifier to upload a file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadAPI.FilePost_0(context.Background()).UploadId(uploadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadAPI.FilePost_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilePost_0`: FilePost201Response
	fmt.Fprintf(os.Stdout, "Response from `UploadAPI.FilePost_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadId** | **string** | Temporary identifier to upload a file | 

### Return type

[**FilePost201Response**](FilePost201Response.md)

### Authorization

[oauth_implicit](../README.md#oauth_implicit), [oauth_authorization_code](../README.md#oauth_authorization_code)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImagePost_0

> FilePost201Response ImagePost_0(ctx).UploadId(uploadId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uploadId := "uploadId_example" // string | Temporary identifier to upload a file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadAPI.ImagePost_0(context.Background()).UploadId(uploadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadAPI.ImagePost_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImagePost_0`: FilePost201Response
	fmt.Fprintf(os.Stdout, "Response from `UploadAPI.ImagePost_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImagePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadId** | **string** | Temporary identifier to upload a file | 

### Return type

[**FilePost201Response**](FilePost201Response.md)

### Authorization

[oauth_implicit](../README.md#oauth_implicit), [oauth_authorization_code](../README.md#oauth_authorization_code)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObjectPatch_0

> ObjectUpload ObjectPatch_0(ctx).PatchMetadata(patchMetadata).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	patchMetadata := *openapiclient.NewObjectPatchMetadata() // ObjectPatchMetadata | Edit the fields of an object. Must be the owner.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadAPI.ObjectPatch_0(context.Background()).PatchMetadata(patchMetadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadAPI.ObjectPatch_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObjectPatch_0`: ObjectUpload
	fmt.Fprintf(os.Stdout, "Response from `UploadAPI.ObjectPatch_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiObjectPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchMetadata** | [**ObjectPatchMetadata**](ObjectPatchMetadata.md) | Edit the fields of an object. Must be the owner. | 

### Return type

[**ObjectUpload**](ObjectUpload.md)

### Authorization

[oauth_implicit](../README.md#oauth_implicit), [oauth_authorization_code](../README.md#oauth_authorization_code)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObjectPost_0

> ObjectUpload ObjectPost_0(ctx).Metadata(metadata).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	metadata := *openapiclient.NewObjectMetadata() // ObjectMetadata | Metadata for submitting a for object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadAPI.ObjectPost_0(context.Background()).Metadata(metadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadAPI.ObjectPost_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObjectPost_0`: ObjectUpload
	fmt.Fprintf(os.Stdout, "Response from `UploadAPI.ObjectPost_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiObjectPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metadata** | [**ObjectMetadata**](ObjectMetadata.md) | Metadata for submitting a for object | 

### Return type

[**ObjectUpload**](ObjectUpload.md)

### Authorization

[oauth_implicit](../README.md#oauth_implicit), [oauth_authorization_code](../README.md#oauth_authorization_code)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObjectsObjectIdUploadStatusGet_0

> ObjectUploadStatus ObjectsObjectIdUploadStatusGet_0(ctx, objectId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	objectId := float32(8.14) // float32 | The object identifier number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadAPI.ObjectsObjectIdUploadStatusGet_0(context.Background(), objectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadAPI.ObjectsObjectIdUploadStatusGet_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObjectsObjectIdUploadStatusGet_0`: ObjectUploadStatus
	fmt.Fprintf(os.Stdout, "Response from `UploadAPI.ObjectsObjectIdUploadStatusGet_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **float32** | The object identifier number | 

### Other Parameters

Other parameters are passed through a pointer to a apiObjectsObjectIdUploadStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectUploadStatus**](ObjectUploadStatus.md)

### Authorization

[oauth_implicit](../README.md#oauth_implicit), [oauth_authorization_code](../README.md#oauth_authorization_code)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

