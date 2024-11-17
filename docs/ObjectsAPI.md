# \ObjectsAPI

All URIs are relative to *https://www.myminifactory.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesFileIdGet**](ObjectsAPI.md#FilesFileIdGet) | **Get** /files/{file_id} | 
[**ObjectPatch**](ObjectsAPI.md#ObjectPatch) | **Patch** /object | 
[**ObjectPost**](ObjectsAPI.md#ObjectPost) | **Post** /object | 
[**ObjectsObjectIdFilesGet**](ObjectsAPI.md#ObjectsObjectIdFilesGet) | **Get** /objects/{object_id}/files | 
[**ObjectsObjectIdGet**](ObjectsAPI.md#ObjectsObjectIdGet) | **Get** /objects/{object_id} | 
[**ObjectsObjectIdPrintsGet**](ObjectsAPI.md#ObjectsObjectIdPrintsGet) | **Get** /objects/{object_id}/prints | 
[**ObjectsObjectIdUploadStatusGet**](ObjectsAPI.md#ObjectsObjectIdUploadStatusGet) | **Get** /objects/{object_id}/upload_status | 
[**UsersUsernameCollectionsGet_0**](ObjectsAPI.md#UsersUsernameCollectionsGet_0) | **Get** /users/{username}/collections | 
[**UsersUsernameObjectsGet_0**](ObjectsAPI.md#UsersUsernameObjectsGet_0) | **Get** /users/{username}/objects | 
[**UsersUsernameObjectsLikedGet_0**](ObjectsAPI.md#UsersUsernameObjectsLikedGet_0) | **Get** /users/{username}/objects_liked | 



## FilesFileIdGet

> File FilesFileIdGet(ctx, fileId).Execute()





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
	fileId := "fileId_example" // string | The file identifier number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectsAPI.FilesFileIdGet(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectsAPI.FilesFileIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FilesFileIdGet`: File
	fmt.Fprintf(os.Stdout, "Response from `ObjectsAPI.FilesFileIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** | The file identifier number | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesFileIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**File**](File.md)

### Authorization

[oauth_implicit](../README.md#oauth_implicit), [key](../README.md#key), [oauth_authorization_code](../README.md#oauth_authorization_code)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObjectPatch

> ObjectUpload ObjectPatch(ctx).PatchMetadata(patchMetadata).Execute()





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
	resp, r, err := apiClient.ObjectsAPI.ObjectPatch(context.Background()).PatchMetadata(patchMetadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectsAPI.ObjectPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObjectPatch`: ObjectUpload
	fmt.Fprintf(os.Stdout, "Response from `ObjectsAPI.ObjectPatch`: %v\n", resp)
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


## ObjectPost

> ObjectUpload ObjectPost(ctx).Metadata(metadata).Execute()





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
	resp, r, err := apiClient.ObjectsAPI.ObjectPost(context.Background()).Metadata(metadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectsAPI.ObjectPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObjectPost`: ObjectUpload
	fmt.Fprintf(os.Stdout, "Response from `ObjectsAPI.ObjectPost`: %v\n", resp)
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


## ObjectsObjectIdFilesGet

> ObjectsObjectIdFilesGet200Response ObjectsObjectIdFilesGet(ctx, objectId).Page(page).PerPage(perPage).Execute()





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
	resp, r, err := apiClient.ObjectsAPI.ObjectsObjectIdFilesGet(context.Background(), objectId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectsAPI.ObjectsObjectIdFilesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObjectsObjectIdFilesGet`: ObjectsObjectIdFilesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectsAPI.ObjectsObjectIdFilesGet`: %v\n", resp)
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


## ObjectsObjectIdGet

> Object ObjectsObjectIdGet(ctx, objectId).Page(page).PerPage(perPage).Execute()





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
	resp, r, err := apiClient.ObjectsAPI.ObjectsObjectIdGet(context.Background(), objectId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectsAPI.ObjectsObjectIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObjectsObjectIdGet`: Object
	fmt.Fprintf(os.Stdout, "Response from `ObjectsAPI.ObjectsObjectIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **float32** | The object identifier number | 

### Other Parameters

Other parameters are passed through a pointer to a apiObjectsObjectIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Page number. Default is 1 | 
 **perPage** | **string** | Number of results per page. Default is 20 | 

### Return type

[**Object**](Object.md)

### Authorization

[oauth_implicit](../README.md#oauth_implicit), [key](../README.md#key), [oauth_authorization_code](../README.md#oauth_authorization_code)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObjectsObjectIdPrintsGet

> ObjectsObjectIdPrintsGet200Response ObjectsObjectIdPrintsGet(ctx, objectId).Page(page).PerPage(perPage).Execute()





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
	resp, r, err := apiClient.ObjectsAPI.ObjectsObjectIdPrintsGet(context.Background(), objectId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectsAPI.ObjectsObjectIdPrintsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObjectsObjectIdPrintsGet`: ObjectsObjectIdPrintsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectsAPI.ObjectsObjectIdPrintsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **float32** | The object identifier number | 

### Other Parameters

Other parameters are passed through a pointer to a apiObjectsObjectIdPrintsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Page number. Default is 1 | 
 **perPage** | **string** | Number of results per page. Default is 20 | 

### Return type

[**ObjectsObjectIdPrintsGet200Response**](ObjectsObjectIdPrintsGet200Response.md)

### Authorization

[oauth_implicit](../README.md#oauth_implicit), [key](../README.md#key), [oauth_authorization_code](../README.md#oauth_authorization_code)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObjectsObjectIdUploadStatusGet

> ObjectUploadStatus ObjectsObjectIdUploadStatusGet(ctx, objectId).Execute()





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
	resp, r, err := apiClient.ObjectsAPI.ObjectsObjectIdUploadStatusGet(context.Background(), objectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectsAPI.ObjectsObjectIdUploadStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObjectsObjectIdUploadStatusGet`: ObjectUploadStatus
	fmt.Fprintf(os.Stdout, "Response from `ObjectsAPI.ObjectsObjectIdUploadStatusGet`: %v\n", resp)
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


## UsersUsernameCollectionsGet_0

> UsersUsernameCollectionsGet200Response UsersUsernameCollectionsGet_0(ctx, username).Page(page).PerPage(perPage).Execute()





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
	username := "username_example" // string | The user's username
	page := "page_example" // string | Page number. Default is 1 (optional)
	perPage := "perPage_example" // string | Number of results per page. Default is 20 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectsAPI.UsersUsernameCollectionsGet_0(context.Background(), username).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectsAPI.UsersUsernameCollectionsGet_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsernameCollectionsGet_0`: UsersUsernameCollectionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectsAPI.UsersUsernameCollectionsGet_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The user&#39;s username | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsernameCollectionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Page number. Default is 1 | 
 **perPage** | **string** | Number of results per page. Default is 20 | 

### Return type

[**UsersUsernameCollectionsGet200Response**](UsersUsernameCollectionsGet200Response.md)

### Authorization

[oauth_implicit](../README.md#oauth_implicit), [key](../README.md#key), [oauth_authorization_code](../README.md#oauth_authorization_code)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsernameObjectsGet_0

> Objects UsersUsernameObjectsGet_0(ctx, username).Page(page).PerPage(perPage).Execute()





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
	username := "username_example" // string | The user's username
	page := "page_example" // string | Page number. Default is 1 (optional)
	perPage := "perPage_example" // string | Number of results per page. Default is 20 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectsAPI.UsersUsernameObjectsGet_0(context.Background(), username).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectsAPI.UsersUsernameObjectsGet_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsernameObjectsGet_0`: Objects
	fmt.Fprintf(os.Stdout, "Response from `ObjectsAPI.UsersUsernameObjectsGet_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The user&#39;s username | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsernameObjectsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Page number. Default is 1 | 
 **perPage** | **string** | Number of results per page. Default is 20 | 

### Return type

[**Objects**](Objects.md)

### Authorization

[oauth_implicit](../README.md#oauth_implicit), [key](../README.md#key), [oauth_authorization_code](../README.md#oauth_authorization_code)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsernameObjectsLikedGet_0

> SearchGet200Response UsersUsernameObjectsLikedGet_0(ctx, username).Page(page).PerPage(perPage).Execute()





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
	username := "username_example" // string | The user's username
	page := "page_example" // string | Page number. Default is 1 (optional)
	perPage := "perPage_example" // string | Number of results per page. Default is 20 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectsAPI.UsersUsernameObjectsLikedGet_0(context.Background(), username).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectsAPI.UsersUsernameObjectsLikedGet_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsernameObjectsLikedGet_0`: SearchGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectsAPI.UsersUsernameObjectsLikedGet_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The user&#39;s username | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsernameObjectsLikedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Page number. Default is 1 | 
 **perPage** | **string** | Number of results per page. Default is 20 | 

### Return type

[**SearchGet200Response**](SearchGet200Response.md)

### Authorization

[oauth_implicit](../README.md#oauth_implicit), [key](../README.md#key), [oauth_authorization_code](../README.md#oauth_authorization_code)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

