# \CollectionsAPI

All URIs are relative to *https://www.myminifactory.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CollectionsCollectionIdGet**](CollectionsAPI.md#CollectionsCollectionIdGet) | **Get** /collections/{collection_id} | 
[**CollectionsGet**](CollectionsAPI.md#CollectionsGet) | **Get** /collections | 
[**UsersUsernameCollectionsGet_1**](CollectionsAPI.md#UsersUsernameCollectionsGet_1) | **Get** /users/{username}/collections | 



## CollectionsCollectionIdGet

> Collection CollectionsCollectionIdGet(ctx, collectionId).Page(page).PerPage(perPage).Execute()





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
	collectionId := "collectionId_example" // string | The collection identifier number
	page := "page_example" // string | Page number. Default is 1 (optional)
	perPage := "perPage_example" // string | Number of results per page. Default is 20 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.CollectionsCollectionIdGet(context.Background(), collectionId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.CollectionsCollectionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CollectionsCollectionIdGet`: Collection
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.CollectionsCollectionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The collection identifier number | 

### Other Parameters

Other parameters are passed through a pointer to a apiCollectionsCollectionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Page number. Default is 1 | 
 **perPage** | **string** | Number of results per page. Default is 20 | 

### Return type

[**Collection**](Collection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CollectionsGet

> UsersUsernameCollectionsGet200Response CollectionsGet(ctx).Page(page).PerPage(perPage).Execute()





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
	page := "page_example" // string | Page number. Default is 1 (optional)
	perPage := "perPage_example" // string | Number of results per page. Default is 20 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.CollectionsGet(context.Background()).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.CollectionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CollectionsGet`: UsersUsernameCollectionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.CollectionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCollectionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Page number. Default is 1 | 
 **perPage** | **string** | Number of results per page. Default is 20 | 

### Return type

[**UsersUsernameCollectionsGet200Response**](UsersUsernameCollectionsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsernameCollectionsGet_1

> UsersUsernameCollectionsGet200Response UsersUsernameCollectionsGet_1(ctx, username).Page(page).PerPage(perPage).Execute()





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
	resp, r, err := apiClient.CollectionsAPI.UsersUsernameCollectionsGet_1(context.Background(), username).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.UsersUsernameCollectionsGet_1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsernameCollectionsGet_1`: UsersUsernameCollectionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.UsersUsernameCollectionsGet_1`: %v\n", resp)
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

