# \UsersAPI

All URIs are relative to *https://www.myminifactory.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserGet**](UsersAPI.md#UserGet) | **Get** /user | 
[**UsersUsernameCollectionsCollectionSlugGet**](UsersAPI.md#UsersUsernameCollectionsCollectionSlugGet) | **Get** /users/{username}/collections/{collection_slug} | 
[**UsersUsernameCollectionsGet**](UsersAPI.md#UsersUsernameCollectionsGet) | **Get** /users/{username}/collections | 
[**UsersUsernameFollowersGet**](UsersAPI.md#UsersUsernameFollowersGet) | **Get** /users/{username}/followers | 
[**UsersUsernameFollowingGet**](UsersAPI.md#UsersUsernameFollowingGet) | **Get** /users/{username}/following | 
[**UsersUsernameGet**](UsersAPI.md#UsersUsernameGet) | **Get** /users/{username} | 
[**UsersUsernameObjectsGet**](UsersAPI.md#UsersUsernameObjectsGet) | **Get** /users/{username}/objects | 
[**UsersUsernameObjectsLikedGet**](UsersAPI.md#UsersUsernameObjectsLikedGet) | **Get** /users/{username}/objects_liked | 



## UserGet

> User UserGet(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UserGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserGet`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetRequest struct via the builder pattern


### Return type

[**User**](User.md)

### Authorization

[oauth_implicit](../README.md#oauth_implicit), [oauth_authorization_code](../README.md#oauth_authorization_code)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsernameCollectionsCollectionSlugGet

> Collection UsersUsernameCollectionsCollectionSlugGet(ctx, username, collectionSlug).Page(page).PerPage(perPage).Execute()





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
	collectionSlug := "collectionSlug_example" // string | The collection slug name
	page := "page_example" // string | Page number. Default is 1 (optional)
	perPage := "perPage_example" // string | Number of results per page. Default is 20 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUsernameCollectionsCollectionSlugGet(context.Background(), username, collectionSlug).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsernameCollectionsCollectionSlugGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsernameCollectionsCollectionSlugGet`: Collection
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUsernameCollectionsCollectionSlugGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The user&#39;s username | 
**collectionSlug** | **string** | The collection slug name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsernameCollectionsCollectionSlugGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **string** | Page number. Default is 1 | 
 **perPage** | **string** | Number of results per page. Default is 20 | 

### Return type

[**Collection**](Collection.md)

### Authorization

[oauth_implicit](../README.md#oauth_implicit), [key](../README.md#key), [oauth_authorization_code](../README.md#oauth_authorization_code)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsernameCollectionsGet

> UsersUsernameCollectionsGet200Response UsersUsernameCollectionsGet(ctx, username).Page(page).PerPage(perPage).Execute()





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
	resp, r, err := apiClient.UsersAPI.UsersUsernameCollectionsGet(context.Background(), username).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsernameCollectionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsernameCollectionsGet`: UsersUsernameCollectionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUsernameCollectionsGet`: %v\n", resp)
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


## UsersUsernameFollowersGet

> UsersUsernameFollowingGet200Response UsersUsernameFollowersGet(ctx, username).Page(page).PerPage(perPage).Execute()





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
	resp, r, err := apiClient.UsersAPI.UsersUsernameFollowersGet(context.Background(), username).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsernameFollowersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsernameFollowersGet`: UsersUsernameFollowingGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUsernameFollowersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The user&#39;s username | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsernameFollowersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Page number. Default is 1 | 
 **perPage** | **string** | Number of results per page. Default is 20 | 

### Return type

[**UsersUsernameFollowingGet200Response**](UsersUsernameFollowingGet200Response.md)

### Authorization

[oauth_implicit](../README.md#oauth_implicit), [key](../README.md#key), [oauth_authorization_code](../README.md#oauth_authorization_code)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsernameFollowingGet

> UsersUsernameFollowingGet200Response UsersUsernameFollowingGet(ctx, username).Page(page).PerPage(perPage).Execute()





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
	resp, r, err := apiClient.UsersAPI.UsersUsernameFollowingGet(context.Background(), username).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsernameFollowingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsernameFollowingGet`: UsersUsernameFollowingGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUsernameFollowingGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The user&#39;s username | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsernameFollowingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Page number. Default is 1 | 
 **perPage** | **string** | Number of results per page. Default is 20 | 

### Return type

[**UsersUsernameFollowingGet200Response**](UsersUsernameFollowingGet200Response.md)

### Authorization

[oauth_implicit](../README.md#oauth_implicit), [key](../README.md#key), [oauth_authorization_code](../README.md#oauth_authorization_code)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsernameGet

> User UsersUsernameGet(ctx, username).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUsernameGet(context.Background(), username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsernameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsernameGet`: User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUsernameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The user&#39;s username | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUsernameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[oauth_implicit](../README.md#oauth_implicit), [key](../README.md#key), [oauth_authorization_code](../README.md#oauth_authorization_code)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUsernameObjectsGet

> Objects UsersUsernameObjectsGet(ctx, username).Page(page).PerPage(perPage).Execute()





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
	resp, r, err := apiClient.UsersAPI.UsersUsernameObjectsGet(context.Background(), username).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsernameObjectsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsernameObjectsGet`: Objects
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUsernameObjectsGet`: %v\n", resp)
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


## UsersUsernameObjectsLikedGet

> SearchGet200Response UsersUsernameObjectsLikedGet(ctx, username).Page(page).PerPage(perPage).Execute()





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
	resp, r, err := apiClient.UsersAPI.UsersUsernameObjectsLikedGet(context.Background(), username).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUsernameObjectsLikedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUsernameObjectsLikedGet`: SearchGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUsernameObjectsLikedGet`: %v\n", resp)
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

