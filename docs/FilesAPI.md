# \FilesAPI

All URIs are relative to *https://www.myminifactory.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilePost**](FilesAPI.md#FilePost) | **Post** /file | 
[**ImagePost**](FilesAPI.md#ImagePost) | **Post** /image | 
[**ObjectsObjectIdFilesGet_0**](FilesAPI.md#ObjectsObjectIdFilesGet_0) | **Get** /objects/{object_id}/files | 



## FilePost

> FilePost201Response FilePost(ctx).UploadId(uploadId).Execute()





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
	resp, r, err := apiClient.FilesAPI.FilePost(context.Background()).UploadId(uploadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.FilePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilePost`: FilePost201Response
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.FilePost`: %v\n", resp)
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


## ImagePost

> FilePost201Response ImagePost(ctx).UploadId(uploadId).Execute()





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
	resp, r, err := apiClient.FilesAPI.ImagePost(context.Background()).UploadId(uploadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.ImagePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImagePost`: FilePost201Response
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.ImagePost`: %v\n", resp)
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


## ObjectsObjectIdFilesGet_0

> ObjectsObjectIdFilesGet200Response ObjectsObjectIdFilesGet_0(ctx, objectId).Page(page).PerPage(perPage).Execute()





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
	page := "page_example" // string | Page number. Default is 1 (optional)
	perPage := "perPage_example" // string | Number of results per page. Default is 20 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.ObjectsObjectIdFilesGet_0(context.Background(), objectId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.ObjectsObjectIdFilesGet_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObjectsObjectIdFilesGet_0`: ObjectsObjectIdFilesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.ObjectsObjectIdFilesGet_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **float32** | The object identifier number | 

### Other Parameters

Other parameters are passed through a pointer to a apiObjectsObjectIdFilesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Page number. Default is 1 | 
 **perPage** | **string** | Number of results per page. Default is 20 | 

### Return type

[**ObjectsObjectIdFilesGet200Response**](ObjectsObjectIdFilesGet200Response.md)

### Authorization

[oauth_implicit](../README.md#oauth_implicit), [key](../README.md#key), [oauth_authorization_code](../README.md#oauth_authorization_code)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

