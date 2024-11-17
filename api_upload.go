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


// UploadAPIService UploadAPI service
type UploadAPIService service

type ApiFilePost_0Request struct {
	ctx context.Context
	ApiService *UploadAPIService
	uploadId *string
}

// Temporary identifier to upload a file
func (r ApiFilePost_0Request) UploadId(uploadId string) ApiFilePost_0Request {
	r.uploadId = &uploadId
	return r
}

func (r ApiFilePost_0Request) Execute() (*FilePost201Response, *http.Response, error) {
	return r.ApiService.FilePostExecute(r)
}

/*
FilePost_0 Method for FilePost_0

Upload a file

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFilePost_0Request
*/
func (a *UploadAPIService) FilePost(ctx context.Context) ApiFilePost_0Request {
	return ApiFilePost_0Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FilePost201Response
func (a *UploadAPIService) FilePostExecute(r ApiFilePost_0Request) (*FilePost201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FilePost201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UploadAPIService.FilePost")
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

type ApiImagePost_0Request struct {
	ctx context.Context
	ApiService *UploadAPIService
	uploadId *string
}

// Temporary identifier to upload a file
func (r ApiImagePost_0Request) UploadId(uploadId string) ApiImagePost_0Request {
	r.uploadId = &uploadId
	return r
}

func (r ApiImagePost_0Request) Execute() (*FilePost201Response, *http.Response, error) {
	return r.ApiService.ImagePostExecute(r)
}

/*
ImagePost_0 Method for ImagePost_0

Upload an image

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiImagePost_0Request
*/
func (a *UploadAPIService) ImagePost(ctx context.Context) ApiImagePost_0Request {
	return ApiImagePost_0Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FilePost201Response
func (a *UploadAPIService) ImagePostExecute(r ApiImagePost_0Request) (*FilePost201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FilePost201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UploadAPIService.ImagePost")
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

type ApiObjectPatch_0Request struct {
	ctx context.Context
	ApiService *UploadAPIService
	patchMetadata *ObjectPatchMetadata
}

// Edit the fields of an object. Must be the owner.
func (r ApiObjectPatch_0Request) PatchMetadata(patchMetadata ObjectPatchMetadata) ApiObjectPatch_0Request {
	r.patchMetadata = &patchMetadata
	return r
}

func (r ApiObjectPatch_0Request) Execute() (*ObjectUpload, *http.Response, error) {
	return r.ApiService.ObjectPatchExecute(r)
}

/*
ObjectPatch_0 Method for ObjectPatch_0

Sent object metadata.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiObjectPatch_0Request
*/
func (a *UploadAPIService) ObjectPatch(ctx context.Context) ApiObjectPatch_0Request {
	return ApiObjectPatch_0Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ObjectUpload
func (a *UploadAPIService) ObjectPatchExecute(r ApiObjectPatch_0Request) (*ObjectUpload, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ObjectUpload
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UploadAPIService.ObjectPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchMetadata == nil {
		return localVarReturnValue, nil, reportError("patchMetadata is required and must be specified")
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
	// body params
	localVarPostBody = r.patchMetadata
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

type ApiObjectPost_0Request struct {
	ctx context.Context
	ApiService *UploadAPIService
	metadata *ObjectMetadata
}

// Metadata for submitting a for object
func (r ApiObjectPost_0Request) Metadata(metadata ObjectMetadata) ApiObjectPost_0Request {
	r.metadata = &metadata
	return r
}

func (r ApiObjectPost_0Request) Execute() (*ObjectUpload, *http.Response, error) {
	return r.ApiService.ObjectPostExecute(r)
}

/*
ObjectPost_0 Method for ObjectPost_0

Sent object metadata.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiObjectPost_0Request
*/
func (a *UploadAPIService) ObjectPost(ctx context.Context) ApiObjectPost_0Request {
	return ApiObjectPost_0Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ObjectUpload
func (a *UploadAPIService) ObjectPostExecute(r ApiObjectPost_0Request) (*ObjectUpload, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ObjectUpload
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UploadAPIService.ObjectPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.metadata == nil {
		return localVarReturnValue, nil, reportError("metadata is required and must be specified")
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
	// body params
	localVarPostBody = r.metadata
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

type ApiObjectsObjectIdUploadStatusGet_0Request struct {
	ctx context.Context
	ApiService *UploadAPIService
	objectId float32
}

func (r ApiObjectsObjectIdUploadStatusGet_0Request) Execute() (*ObjectUploadStatus, *http.Response, error) {
	return r.ApiService.ObjectsObjectIdUploadStatusGetExecute(r)
}

/*
ObjectsObjectIdUploadStatusGet_0 Method for ObjectsObjectIdUploadStatusGet_0

Check the status of an object and its files.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param objectId The object identifier number
 @return ApiObjectsObjectIdUploadStatusGet_0Request
*/
func (a *UploadAPIService) ObjectsObjectIdUploadStatusGet(ctx context.Context, objectId float32) ApiObjectsObjectIdUploadStatusGet_0Request {
	return ApiObjectsObjectIdUploadStatusGet_0Request{
		ApiService: a,
		ctx: ctx,
		objectId: objectId,
	}
}

// Execute executes the request
//  @return ObjectUploadStatus
func (a *UploadAPIService) ObjectsObjectIdUploadStatusGetExecute(r ApiObjectsObjectIdUploadStatusGet_0Request) (*ObjectUploadStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ObjectUploadStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UploadAPIService.ObjectsObjectIdUploadStatusGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/objects/{object_id}/upload_status"
	localVarPath = strings.Replace(localVarPath, "{"+"object_id"+"}", url.PathEscape(parameterValueToString(r.objectId, "objectId")), -1)

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
