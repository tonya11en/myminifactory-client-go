# \SearchAPI

All URIs are relative to *https://www.myminifactory.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchGet**](SearchAPI.md#SearchGet) | **Get** /search | 



## SearchGet

> SearchGet200Response SearchGet(ctx).Q(q).Page(page).PerPage(perPage).Sort(sort).Order(order).Cat(cat).Support(support).Tech(tech).Complexity(complexity).Featured(featured).Remix(remix).CommercialUse(commercialUse).Exclusive(exclusive).Execute()





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
	q := "q_example" // string | Search query. See the details query syntax document
	page := "page_example" // string | Page number. Default is 1 (optional)
	perPage := "perPage_example" // string | Number of results per page. Default is 20 (optional)
	sort := "sort_example" // string | Sort results by: 'visits', 'date', 'popularity'. (optional)
	order := "order_example" // string | Sorting order: 'desc', 'asc' (Defaults to desc). (optional)
	cat := "cat_example" // string | Filter object of a certain category, by the category id. eg. cat={category_id} (optional)
	support := "support_example" // string | Filter printing support-free objects: 1: support-free (optional)
	tech := "tech_example" // string | Filter printing technology recommanded for the object: eg. 'DLP/SLA' (optional)
	complexity := "complexity_example" // string | Filter object difficulty (How hard to build) : 'eas', 'med', 'diff' (optional)
	featured := "featured_example" // string | Filter featured object: 0: Non-featured, 1: featured (optional)
	remix := "remix_example" // string | License filter: the designer accepted his object the remixed. (optional)
	commercialUse := "commercialUse_example" // string | License filter: the designer accepted commercial use of his object. (optional)
	exclusive := "exclusive_example" // string | License filter: this object should be shared exclusivly on MyMiniFactory. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchGet(context.Background()).Q(q).Page(page).PerPage(perPage).Sort(sort).Order(order).Cat(cat).Support(support).Tech(tech).Complexity(complexity).Featured(featured).Remix(remix).CommercialUse(commercialUse).Exclusive(exclusive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGet`: SearchGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Search query. See the details query syntax document | 
 **page** | **string** | Page number. Default is 1 | 
 **perPage** | **string** | Number of results per page. Default is 20 | 
 **sort** | **string** | Sort results by: &#39;visits&#39;, &#39;date&#39;, &#39;popularity&#39;. | 
 **order** | **string** | Sorting order: &#39;desc&#39;, &#39;asc&#39; (Defaults to desc). | 
 **cat** | **string** | Filter object of a certain category, by the category id. eg. cat&#x3D;{category_id} | 
 **support** | **string** | Filter printing support-free objects: 1: support-free | 
 **tech** | **string** | Filter printing technology recommanded for the object: eg. &#39;DLP/SLA&#39; | 
 **complexity** | **string** | Filter object difficulty (How hard to build) : &#39;eas&#39;, &#39;med&#39;, &#39;diff&#39; | 
 **featured** | **string** | Filter featured object: 0: Non-featured, 1: featured | 
 **remix** | **string** | License filter: the designer accepted his object the remixed. | 
 **commercialUse** | **string** | License filter: the designer accepted commercial use of his object. | 
 **exclusive** | **string** | License filter: this object should be shared exclusivly on MyMiniFactory. | 

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

