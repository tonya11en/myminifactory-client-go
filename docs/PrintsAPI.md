# \PrintsAPI

All URIs are relative to *https://www.myminifactory.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ObjectsObjectIdPrintsGet_0**](PrintsAPI.md#ObjectsObjectIdPrintsGet_0) | **Get** /objects/{object_id}/prints | 



## ObjectsObjectIdPrintsGet_0

> ObjectsObjectIdPrintsGet200Response ObjectsObjectIdPrintsGet_0(ctx, objectId).Page(page).PerPage(perPage).Execute()





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
	resp, r, err := apiClient.PrintsAPI.ObjectsObjectIdPrintsGet_0(context.Background(), objectId).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrintsAPI.ObjectsObjectIdPrintsGet_0``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObjectsObjectIdPrintsGet_0`: ObjectsObjectIdPrintsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `PrintsAPI.ObjectsObjectIdPrintsGet_0`: %v\n", resp)
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

