/*
Elba Public API

Testing ELbaAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package elba

import (
	"context"
	openapiclient "github.com/archaron/e
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_elba_ELbaAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ELbaAPIService OrganizationsGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ELbaAPI.OrganizationsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
