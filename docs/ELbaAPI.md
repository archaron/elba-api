# \ELbaAPI

All URIs are relative to *https://elba-api.kontur.ru/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationsGet**](ELbaAPI.md#OrganizationsGet) | **Get** /organizations | Получение списка доступных организаций



## OrganizationsGet

> GetOrganizationsResponse OrganizationsGet(ctx).Offset(offset).Limit(limit).Execute()

Получение списка доступных организаций

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
	offset := int32(56) // int32 | Позиция, начиная с которой будут вычитываться организации. По-умолчанию 0 (optional)
	limit := int32(56) // int32 | Максимальное количество организаций для получения. По-умолчанию 100 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ELbaAPI.OrganizationsGet(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ELbaAPI.OrganizationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsGet`: GetOrganizationsResponse
	fmt.Fprintf(os.Stdout, "Response from `ELbaAPI.OrganizationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Позиция, начиная с которой будут вычитываться организации. По-умолчанию 0 | 
 **limit** | **int32** | Максимальное количество организаций для получения. По-умолчанию 100 | 

### Return type

[**GetOrganizationsResponse**](GetOrganizationsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

