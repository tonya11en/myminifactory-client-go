/*
MyMiniFactory API

3D printable object API

API version: 2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// UsersAPIService UsersAPI service
type UsersAPIService service

type ApiUserGetRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
}

func (r ApiUserGetRequest) Execute() (*User, *http.Response, error) {
	return r.ApiService.UserGetExecute(r)
}

/*
UserGet Method for UserGet

Get the current user basic information.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUserGetRequest
*/
func (a *UsersAPIService) UserGet(ctx context.Context) ApiUserGetRequest {
	return ApiUserGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return User
func (a *UsersAPIService) UserGetExecute(r ApiUserGetRequest) (*User, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *User
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.UserGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUsersUsernameCollectionsCollectionSlugGetRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	username string
	collectionSlug string
	page *string
	perPage *string
}

// Page number. Default is 1
func (r ApiUsersUsernameCollectionsCollectionSlugGetRequest) Page(page string) ApiUsersUsernameCollectionsCollectionSlugGetRequest {
	r.page = &page
	return r
}

// Number of results per page. Default is 20
func (r ApiUsersUsernameCollectionsCollectionSlugGetRequest) PerPage(perPage string) ApiUsersUsernameCollectionsCollectionSlugGetRequest {
	r.perPage = &perPage
	return r
}

func (r ApiUsersUsernameCollectionsCollectionSlugGetRequest) Execute() (*Collection, *http.Response, error) {
	return r.ApiService.UsersUsernameCollectionsCollectionSlugGetExecute(r)
}

/*
UsersUsernameCollectionsCollectionSlugGet Method for UsersUsernameCollectionsCollectionSlugGet

Get a user's collections.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username The user's username
 @param collectionSlug The collection slug name
 @return ApiUsersUsernameCollectionsCollectionSlugGetRequest
*/
func (a *UsersAPIService) UsersUsernameCollectionsCollectionSlugGet(ctx context.Context, username string, collectionSlug string) ApiUsersUsernameCollectionsCollectionSlugGetRequest {
	return ApiUsersUsernameCollectionsCollectionSlugGetRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
		collectionSlug: collectionSlug,
	}
}

// Execute executes the request
//  @return Collection
func (a *UsersAPIService) UsersUsernameCollectionsCollectionSlugGetExecute(r ApiUsersUsernameCollectionsCollectionSlugGetRequest) (*Collection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Collection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.UsersUsernameCollectionsCollectionSlugGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{username}/collections/{collection_slug}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"collection_slug"+"}", url.PathEscape(parameterValueToString(r.collectionSlug, "collectionSlug")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "", "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("key", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUsersUsernameCollectionsGetRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	username string
	page *string
	perPage *string
}

// Page number. Default is 1
func (r ApiUsersUsernameCollectionsGetRequest) Page(page string) ApiUsersUsernameCollectionsGetRequest {
	r.page = &page
	return r
}

// Number of results per page. Default is 20
func (r ApiUsersUsernameCollectionsGetRequest) PerPage(perPage string) ApiUsersUsernameCollectionsGetRequest {
	r.perPage = &perPage
	return r
}

func (r ApiUsersUsernameCollectionsGetRequest) Execute() (*UsersUsernameCollectionsGet200Response, *http.Response, error) {
	return r.ApiService.UsersUsernameCollectionsGetExecute(r)
}

/*
UsersUsernameCollectionsGet Method for UsersUsernameCollectionsGet

Get a user's collections.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username The user's username
 @return ApiUsersUsernameCollectionsGetRequest
*/
func (a *UsersAPIService) UsersUsernameCollectionsGet(ctx context.Context, username string) ApiUsersUsernameCollectionsGetRequest {
	return ApiUsersUsernameCollectionsGetRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

// Execute executes the request
//  @return UsersUsernameCollectionsGet200Response
func (a *UsersAPIService) UsersUsernameCollectionsGetExecute(r ApiUsersUsernameCollectionsGetRequest) (*UsersUsernameCollectionsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UsersUsernameCollectionsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.UsersUsernameCollectionsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{username}/collections"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "", "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("key", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUsersUsernameFollowersGetRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	username string
	page *string
	perPage *string
}

// Page number. Default is 1
func (r ApiUsersUsernameFollowersGetRequest) Page(page string) ApiUsersUsernameFollowersGetRequest {
	r.page = &page
	return r
}

// Number of results per page. Default is 20
func (r ApiUsersUsernameFollowersGetRequest) PerPage(perPage string) ApiUsersUsernameFollowersGetRequest {
	r.perPage = &perPage
	return r
}

func (r ApiUsersUsernameFollowersGetRequest) Execute() (*UsersUsernameFollowingGet200Response, *http.Response, error) {
	return r.ApiService.UsersUsernameFollowersGetExecute(r)
}

/*
UsersUsernameFollowersGet Method for UsersUsernameFollowersGet

Get a user's followers.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username The user's username
 @return ApiUsersUsernameFollowersGetRequest
*/
func (a *UsersAPIService) UsersUsernameFollowersGet(ctx context.Context, username string) ApiUsersUsernameFollowersGetRequest {
	return ApiUsersUsernameFollowersGetRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

// Execute executes the request
//  @return UsersUsernameFollowingGet200Response
func (a *UsersAPIService) UsersUsernameFollowersGetExecute(r ApiUsersUsernameFollowersGetRequest) (*UsersUsernameFollowingGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UsersUsernameFollowingGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.UsersUsernameFollowersGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{username}/followers"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "", "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("key", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUsersUsernameFollowingGetRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	username string
	page *string
	perPage *string
}

// Page number. Default is 1
func (r ApiUsersUsernameFollowingGetRequest) Page(page string) ApiUsersUsernameFollowingGetRequest {
	r.page = &page
	return r
}

// Number of results per page. Default is 20
func (r ApiUsersUsernameFollowingGetRequest) PerPage(perPage string) ApiUsersUsernameFollowingGetRequest {
	r.perPage = &perPage
	return r
}

func (r ApiUsersUsernameFollowingGetRequest) Execute() (*UsersUsernameFollowingGet200Response, *http.Response, error) {
	return r.ApiService.UsersUsernameFollowingGetExecute(r)
}

/*
UsersUsernameFollowingGet Method for UsersUsernameFollowingGet

Get a user's followings.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username The user's username
 @return ApiUsersUsernameFollowingGetRequest
*/
func (a *UsersAPIService) UsersUsernameFollowingGet(ctx context.Context, username string) ApiUsersUsernameFollowingGetRequest {
	return ApiUsersUsernameFollowingGetRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

// Execute executes the request
//  @return UsersUsernameFollowingGet200Response
func (a *UsersAPIService) UsersUsernameFollowingGetExecute(r ApiUsersUsernameFollowingGetRequest) (*UsersUsernameFollowingGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UsersUsernameFollowingGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.UsersUsernameFollowingGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{username}/following"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "", "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("key", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUsersUsernameGetRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	username string
}

func (r ApiUsersUsernameGetRequest) Execute() (*User, *http.Response, error) {
	return r.ApiService.UsersUsernameGetExecute(r)
}

/*
UsersUsernameGet Method for UsersUsernameGet

Get basic information about a user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username The user's username
 @return ApiUsersUsernameGetRequest
*/
func (a *UsersAPIService) UsersUsernameGet(ctx context.Context, username string) ApiUsersUsernameGetRequest {
	return ApiUsersUsernameGetRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

// Execute executes the request
//  @return User
func (a *UsersAPIService) UsersUsernameGetExecute(r ApiUsersUsernameGetRequest) (*User, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *User
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.UsersUsernameGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{username}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("key", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUsersUsernameObjectsGetRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	username string
	page *string
	perPage *string
}

// Page number. Default is 1
func (r ApiUsersUsernameObjectsGetRequest) Page(page string) ApiUsersUsernameObjectsGetRequest {
	r.page = &page
	return r
}

// Number of results per page. Default is 20
func (r ApiUsersUsernameObjectsGetRequest) PerPage(perPage string) ApiUsersUsernameObjectsGetRequest {
	r.perPage = &perPage
	return r
}

func (r ApiUsersUsernameObjectsGetRequest) Execute() (*Objects, *http.Response, error) {
	return r.ApiService.UsersUsernameObjectsGetExecute(r)
}

/*
UsersUsernameObjectsGet Method for UsersUsernameObjectsGet

Get a user's objects.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username The user's username
 @return ApiUsersUsernameObjectsGetRequest
*/
func (a *UsersAPIService) UsersUsernameObjectsGet(ctx context.Context, username string) ApiUsersUsernameObjectsGetRequest {
	return ApiUsersUsernameObjectsGetRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

// Execute executes the request
//  @return Objects
func (a *UsersAPIService) UsersUsernameObjectsGetExecute(r ApiUsersUsernameObjectsGetRequest) (*Objects, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Objects
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.UsersUsernameObjectsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{username}/objects"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "", "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("key", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUsersUsernameObjectsLikedGetRequest struct {
	ctx context.Context
	ApiService *UsersAPIService
	username string
	page *string
	perPage *string
}

// Page number. Default is 1
func (r ApiUsersUsernameObjectsLikedGetRequest) Page(page string) ApiUsersUsernameObjectsLikedGetRequest {
	r.page = &page
	return r
}

// Number of results per page. Default is 20
func (r ApiUsersUsernameObjectsLikedGetRequest) PerPage(perPage string) ApiUsersUsernameObjectsLikedGetRequest {
	r.perPage = &perPage
	return r
}

func (r ApiUsersUsernameObjectsLikedGetRequest) Execute() (*SearchGet200Response, *http.Response, error) {
	return r.ApiService.UsersUsernameObjectsLikedGetExecute(r)
}

/*
UsersUsernameObjectsLikedGet Method for UsersUsernameObjectsLikedGet

Get a user's liked objects.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username The user's username
 @return ApiUsersUsernameObjectsLikedGetRequest
*/
func (a *UsersAPIService) UsersUsernameObjectsLikedGet(ctx context.Context, username string) ApiUsersUsernameObjectsLikedGetRequest {
	return ApiUsersUsernameObjectsLikedGetRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

// Execute executes the request
//  @return SearchGet200Response
func (a *UsersAPIService) UsersUsernameObjectsLikedGetExecute(r ApiUsersUsernameObjectsLikedGetRequest) (*SearchGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SearchGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.UsersUsernameObjectsLikedGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{username}/objects_liked"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "", "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("key", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
