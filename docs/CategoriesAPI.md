# \CategoriesAPI

All URIs are relative to *https://www.myminifactory.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CategoriesCategoryIdGet**](CategoriesAPI.md#CategoriesCategoryIdGet) | **Get** /categories/{category_id} | 
[**CategoriesGet**](CategoriesAPI.md#CategoriesGet) | **Get** /categories | 



## CategoriesCategoryIdGet

> Category CategoriesCategoryIdGet(ctx, categoryId).Execute()





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
	categoryId := "categoryId_example" // string | The category identifier number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoriesAPI.CategoriesCategoryIdGet(context.Background(), categoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.CategoriesCategoryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoriesCategoryIdGet`: Category
	fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.CategoriesCategoryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **string** | The category identifier number | 

### Other Parameters

Other parameters are passed through a pointer to a apiCategoriesCategoryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Category**](Category.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoriesGet

> CategoriesGet200Response CategoriesGet(ctx).Top(top).Page(page).PerPage(perPage).Execute()





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
	top := true // bool | Filter top categories (optional)
	page := "page_example" // string | Page number. Default is 1 (optional)
	perPage := "perPage_example" // string | Number of results per page. Default is 20 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CategoriesAPI.CategoriesGet(context.Background()).Top(top).Page(page).PerPage(perPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CategoriesAPI.CategoriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CategoriesGet`: CategoriesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CategoriesAPI.CategoriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCategoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **bool** | Filter top categories | 
 **page** | **string** | Page number. Default is 1 | 
 **perPage** | **string** | Number of results per page. Default is 20 | 

### Return type

[**CategoriesGet200Response**](CategoriesGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

