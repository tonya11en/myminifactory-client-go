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


// FilesAPIService FilesAPI service
type FilesAPIService service

type ApiFilePostRequest struct {
	ctx context.Context
	ApiService *FilesAPIService
	uploadId *string
}

// Temporary identifier to upload a file
func (r ApiFilePostRequest) UploadId(uploadId string) ApiFilePostRequest {
	r.uploadId = &uploadId
	return r
}

func (r ApiFilePostRequest) Execute() (*FilePost201Response, *http.Response, error) {
	return r.ApiService.FilePostExecute(r)
}

/*
FilePost Method for FilePost

Upload a file

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFilePostRequest
*/
func (a *FilesAPIService) FilePost(ctx context.Context) ApiFilePostRequest {
	return ApiFilePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FilePost201Response
func (a *FilesAPIService) FilePostExecute(r ApiFilePostRequest) (*FilePost201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FilePost201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesAPIService.FilePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/file"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.uploadId == nil {
		return localVarReturnValue, nil, reportError("uploadId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "upload_id", r.uploadId, "", "")
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

type ApiImagePostRequest struct {
	ctx context.Context
	ApiService *FilesAPIService
	uploadId *string
}

// Temporary identifier to upload a file
func (r ApiImagePostRequest) UploadId(uploadId string) ApiImagePostRequest {
	r.uploadId = &uploadId
	return r
}

func (r ApiImagePostRequest) Execute() (*FilePost201Response, *http.Response, error) {
	return r.ApiService.ImagePostExecute(r)
}

/*
ImagePost Method for ImagePost

Upload an image

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiImagePostRequest
*/
func (a *FilesAPIService) ImagePost(ctx context.Context) ApiImagePostRequest {
	return ApiImagePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FilePost201Response
func (a *FilesAPIService) ImagePostExecute(r ApiImagePostRequest) (*FilePost201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FilePost201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesAPIService.ImagePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/image"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.uploadId == nil {
		return localVarReturnValue, nil, reportError("uploadId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "upload_id", r.uploadId, "", "")
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

type ApiObjectsObjectIdFilesGet_0Request struct {
	ctx context.Context
	ApiService *FilesAPIService
	objectId float32
	page *string
	perPage *string
}

// Page number. Default is 1
func (r ApiObjectsObjectIdFilesGet_0Request) Page(page string) ApiObjectsObjectIdFilesGet_0Request {
	r.page = &page
	return r
}

// Number of results per page. Default is 20
func (r ApiObjectsObjectIdFilesGet_0Request) PerPage(perPage string) ApiObjectsObjectIdFilesGet_0Request {
	r.perPage = &perPage
	return r
}

func (r ApiObjectsObjectIdFilesGet_0Request) Execute() (*ObjectsObjectIdFilesGet200Response, *http.Response, error) {
	return r.ApiService.ObjectsObjectIdFilesGetExecute(r)
}

/*
ObjectsObjectIdFilesGet_0 Method for ObjectsObjectIdFilesGet_0

Get the list of files of the object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectId The object identifier number
 @return ApiObjectsObjectIdFilesGet_0Request
*/
func (a *FilesAPIService) ObjectsObjectIdFilesGet(ctx context.Context, objectId float32) ApiObjectsObjectIdFilesGet_0Request {
	return ApiObjectsObjectIdFilesGet_0Request{
		ApiService: a,
		ctx: ctx,
		objectId: objectId,
	}
}

// Execute executes the request
//  @return ObjectsObjectIdFilesGet200Response
func (a *FilesAPIService) ObjectsObjectIdFilesGetExecute(r ApiObjectsObjectIdFilesGet_0Request) (*ObjectsObjectIdFilesGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ObjectsObjectIdFilesGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesAPIService.ObjectsObjectIdFilesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/objects/{object_id}/files"
	localVarPath = strings.Replace(localVarPath, "{"+"object_id"+"}", url.PathEscape(parameterValueToString(r.objectId, "objectId")), -1)

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
