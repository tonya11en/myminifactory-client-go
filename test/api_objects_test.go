/*
MyMiniFactory API

Testing ObjectsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_ObjectsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectsAPIService FilesFileIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var fileId string

		resp, httpRes, err := apiClient.ObjectsAPI.FilesFileIdGet(context.Background(), fileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectsAPIService ObjectPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectsAPI.ObjectPatch(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectsAPIService ObjectPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ObjectsAPI.ObjectPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectsAPIService ObjectsObjectIdFilesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var objectId float32

		resp, httpRes, err := apiClient.ObjectsAPI.ObjectsObjectIdFilesGet(context.Background(), objectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectsAPIService ObjectsObjectIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var objectId float32

		resp, httpRes, err := apiClient.ObjectsAPI.ObjectsObjectIdGet(context.Background(), objectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectsAPIService ObjectsObjectIdPrintsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var objectId float32

		resp, httpRes, err := apiClient.ObjectsAPI.ObjectsObjectIdPrintsGet(context.Background(), objectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectsAPIService ObjectsObjectIdUploadStatusGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var objectId float32

		resp, httpRes, err := apiClient.ObjectsAPI.ObjectsObjectIdUploadStatusGet(context.Background(), objectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectsAPIService UsersUsernameCollectionsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		resp, httpRes, err := apiClient.ObjectsAPI.UsersUsernameCollectionsGet_0(context.Background(), username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectsAPIService UsersUsernameObjectsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		resp, httpRes, err := apiClient.ObjectsAPI.UsersUsernameObjectsGet_0(context.Background(), username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectsAPIService UsersUsernameObjectsLikedGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		resp, httpRes, err := apiClient.ObjectsAPI.UsersUsernameObjectsLikedGet_0(context.Background(), username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
