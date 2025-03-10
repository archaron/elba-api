# \ElbaAPI

All URIs are relative to *https://elba-api.kontur.ru/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LongRunningTasksTaskIdGet**](ElbaAPI.md#LongRunningTasksTaskIdGet) | **Get** /long-running-tasks/{taskId} | Получить состояние операции
[**OrganizationsOrganizationIdActsPost**](ElbaAPI.md#OrganizationsOrganizationIdActsPost) | **Post** /organizations/{organizationId}/acts | Создание исходящего акта
[**OrganizationsOrganizationIdBankAccountsGet**](ElbaAPI.md#OrganizationsOrganizationIdBankAccountsGet) | **Get** /organizations/{organizationId}/bank-accounts | Получение списка банковских счетов
[**OrganizationsOrganizationIdBankAccountsPost**](ElbaAPI.md#OrganizationsOrganizationIdBankAccountsPost) | **Post** /organizations/{organizationId}/bank-accounts | Создание банковского счёта
[**OrganizationsOrganizationIdBillsBillIdGet**](ElbaAPI.md#OrganizationsOrganizationIdBillsBillIdGet) | **Get** /organizations/{organizationId}/bills/{billId} | Получение исходящего счёта
[**OrganizationsOrganizationIdBillsPost**](ElbaAPI.md#OrganizationsOrganizationIdBillsPost) | **Post** /organizations/{organizationId}/bills | Создание исходящего счёта
[**OrganizationsOrganizationIdContractorsContractorIdGet**](ElbaAPI.md#OrganizationsOrganizationIdContractorsContractorIdGet) | **Get** /organizations/{organizationId}/contractors/{contractorId} | Получение контрагента
[**OrganizationsOrganizationIdContractorsPost**](ElbaAPI.md#OrganizationsOrganizationIdContractorsPost) | **Post** /organizations/{organizationId}/contractors | Создание контрагента
[**OrganizationsOrganizationIdContractorsSearchPost**](ElbaAPI.md#OrganizationsOrganizationIdContractorsSearchPost) | **Post** /organizations/{organizationId}/contractors/search | Поиск контрагентов
[**OrganizationsOrganizationIdDeliveryNotesPost**](ElbaAPI.md#OrganizationsOrganizationIdDeliveryNotesPost) | **Post** /organizations/{organizationId}/delivery-notes | Создание исходящей накладной
[**OrganizationsOrganizationIdDocumentNewsGet**](ElbaAPI.md#OrganizationsOrganizationIdDocumentNewsGet) | **Get** /organizations/{organizationId}/document-news | Получение списка новостей по документам
[**OrganizationsOrganizationIdDocumentNewsSubscriptionPost**](ElbaAPI.md#OrganizationsOrganizationIdDocumentNewsSubscriptionPost) | **Post** /organizations/{organizationId}/document-news/subscription | Подписаться на новости по документам
[**OrganizationsOrganizationIdProductsPost**](ElbaAPI.md#OrganizationsOrganizationIdProductsPost) | **Post** /organizations/{organizationId}/products | Создание товара
[**OrganizationsOrganizationIdProductsProductIdGet**](ElbaAPI.md#OrganizationsOrganizationIdProductsProductIdGet) | **Get** /organizations/{organizationId}/products/{productId} | Получение товара
[**OrganizationsOrganizationIdProductsProductIdUpdatePost**](ElbaAPI.md#OrganizationsOrganizationIdProductsProductIdUpdatePost) | **Post** /organizations/{organizationId}/products/{productId}/update | Обновление товара
[**OrganizationsOrganizationIdProductsSearchPost**](ElbaAPI.md#OrganizationsOrganizationIdProductsSearchPost) | **Post** /organizations/{organizationId}/products/search | Поиск товаров
[**OrganizationsOrganizationIdUniversalTransferDocumentsPost**](ElbaAPI.md#OrganizationsOrganizationIdUniversalTransferDocumentsPost) | **Post** /organizations/{organizationId}/universal-transfer-documents | Создание исходящего универсального передаточного документа



## LongRunningTasksTaskIdGet

> LongRunningTaskState LongRunningTasksTaskIdGet(ctx, taskId).Execute()

Получить состояние операции

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
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Идентификатор длительной операции

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElbaAPI.LongRunningTasksTaskIdGet(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElbaAPI.LongRunningTasksTaskIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LongRunningTasksTaskIdGet`: LongRunningTaskState
	fmt.Fprintf(os.Stdout, "Response from `ElbaAPI.LongRunningTasksTaskIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | Идентификатор длительной операции | 

### Other Parameters

Other parameters are passed through a pointer to a apiLongRunningTasksTaskIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LongRunningTaskState**](LongRunningTaskState.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationIdActsPost

> CreateActResponse OrganizationsOrganizationIdActsPost(ctx, organizationId).CreateActRequest(createActRequest).Execute()

Создание исходящего акта

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "972252c5-6551-4ec4-90ad-cb5b296490cb" // string | Идентификатор организации
	createActRequest := *openapiclient.NewCreateActRequest(time.Now()) // CreateActRequest | Данные акта

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElbaAPI.OrganizationsOrganizationIdActsPost(context.Background(), organizationId).CreateActRequest(createActRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElbaAPI.OrganizationsOrganizationIdActsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsOrganizationIdActsPost`: CreateActResponse
	fmt.Fprintf(os.Stdout, "Response from `ElbaAPI.OrganizationsOrganizationIdActsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Идентификатор организации | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationIdActsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createActRequest** | [**CreateActRequest**](CreateActRequest.md) | Данные акта | 

### Return type

[**CreateActResponse**](CreateActResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationIdBankAccountsGet

> GetBankAccountsResponse OrganizationsOrganizationIdBankAccountsGet(ctx, organizationId).Offset(offset).Limit(limit).Execute()

Получение списка банковских счетов

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
	organizationId := "972252c5-6551-4ec4-90ad-cb5b296490cb" // string | 
	offset := int32(56) // int32 | Позиция, начиная с которой будут вычитываться счета. По-умолчанию 0 (optional)
	limit := int32(56) // int32 | Максимальное количество счетов для получения. По-умолчанию 100 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElbaAPI.OrganizationsOrganizationIdBankAccountsGet(context.Background(), organizationId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElbaAPI.OrganizationsOrganizationIdBankAccountsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsOrganizationIdBankAccountsGet`: GetBankAccountsResponse
	fmt.Fprintf(os.Stdout, "Response from `ElbaAPI.OrganizationsOrganizationIdBankAccountsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationIdBankAccountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | Позиция, начиная с которой будут вычитываться счета. По-умолчанию 0 | 
 **limit** | **int32** | Максимальное количество счетов для получения. По-умолчанию 100 | 

### Return type

[**GetBankAccountsResponse**](GetBankAccountsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationIdBankAccountsPost

> CreateBankAccountResponse OrganizationsOrganizationIdBankAccountsPost(ctx, organizationId).CreateBankAccountRequest(createBankAccountRequest).Execute()

Создание банковского счёта

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
	organizationId := "972252c5-6551-4ec4-90ad-cb5b296490cb" // string | Идентификатор организации
	createBankAccountRequest := *openapiclient.NewCreateBankAccountRequest("70719810591646992616", "046577904") // CreateBankAccountRequest | Данные банковского счёта

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElbaAPI.OrganizationsOrganizationIdBankAccountsPost(context.Background(), organizationId).CreateBankAccountRequest(createBankAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElbaAPI.OrganizationsOrganizationIdBankAccountsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsOrganizationIdBankAccountsPost`: CreateBankAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `ElbaAPI.OrganizationsOrganizationIdBankAccountsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Идентификатор организации | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationIdBankAccountsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createBankAccountRequest** | [**CreateBankAccountRequest**](CreateBankAccountRequest.md) | Данные банковского счёта | 

### Return type

[**CreateBankAccountResponse**](CreateBankAccountResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationIdBillsBillIdGet

> Bill OrganizationsOrganizationIdBillsBillIdGet(ctx, organizationId, billId).Execute()

Получение исходящего счёта

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
	organizationId := "972252c5-6551-4ec4-90ad-cb5b296490cb" // string | Идентификатор организации
	billId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Идентификатор счёта

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElbaAPI.OrganizationsOrganizationIdBillsBillIdGet(context.Background(), organizationId, billId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElbaAPI.OrganizationsOrganizationIdBillsBillIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsOrganizationIdBillsBillIdGet`: Bill
	fmt.Fprintf(os.Stdout, "Response from `ElbaAPI.OrganizationsOrganizationIdBillsBillIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Идентификатор организации | 
**billId** | **string** | Идентификатор счёта | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationIdBillsBillIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Bill**](Bill.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationIdBillsPost

> CreateBillResponse OrganizationsOrganizationIdBillsPost(ctx, organizationId).CreateBillRequest(createBillRequest).Execute()

Создание исходящего счёта

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "972252c5-6551-4ec4-90ad-cb5b296490cb" // string | Идентификатор организации
	createBillRequest := *openapiclient.NewCreateBillRequest(time.Now()) // CreateBillRequest | Данные счёта

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElbaAPI.OrganizationsOrganizationIdBillsPost(context.Background(), organizationId).CreateBillRequest(createBillRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElbaAPI.OrganizationsOrganizationIdBillsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsOrganizationIdBillsPost`: CreateBillResponse
	fmt.Fprintf(os.Stdout, "Response from `ElbaAPI.OrganizationsOrganizationIdBillsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Идентификатор организации | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationIdBillsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createBillRequest** | [**CreateBillRequest**](CreateBillRequest.md) | Данные счёта | 

### Return type

[**CreateBillResponse**](CreateBillResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationIdContractorsContractorIdGet

> Contractor OrganizationsOrganizationIdContractorsContractorIdGet(ctx, organizationId, contractorId).Execute()

Получение контрагента

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
	organizationId := "972252c5-6551-4ec4-90ad-cb5b296490cb" // string | Идентификатор организации
	contractorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Идентификатор контагента

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElbaAPI.OrganizationsOrganizationIdContractorsContractorIdGet(context.Background(), organizationId, contractorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElbaAPI.OrganizationsOrganizationIdContractorsContractorIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsOrganizationIdContractorsContractorIdGet`: Contractor
	fmt.Fprintf(os.Stdout, "Response from `ElbaAPI.OrganizationsOrganizationIdContractorsContractorIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Идентификатор организации | 
**contractorId** | **string** | Идентификатор контагента | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationIdContractorsContractorIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Contractor**](Contractor.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationIdContractorsPost

> CreateContractorResponse OrganizationsOrganizationIdContractorsPost(ctx, organizationId).CreateContractorRequest(createContractorRequest).Execute()

Создание контрагента

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
	organizationId := "972252c5-6551-4ec4-90ad-cb5b296490cb" // string | Идентификатор организации
	createContractorRequest := *openapiclient.NewCreateContractorRequest("Name_example") // CreateContractorRequest | Данные контрагента

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElbaAPI.OrganizationsOrganizationIdContractorsPost(context.Background(), organizationId).CreateContractorRequest(createContractorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElbaAPI.OrganizationsOrganizationIdContractorsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsOrganizationIdContractorsPost`: CreateContractorResponse
	fmt.Fprintf(os.Stdout, "Response from `ElbaAPI.OrganizationsOrganizationIdContractorsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Идентификатор организации | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationIdContractorsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createContractorRequest** | [**CreateContractorRequest**](CreateContractorRequest.md) | Данные контрагента | 

### Return type

[**CreateContractorResponse**](CreateContractorResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationIdContractorsSearchPost

> SearchContractorsTaskResult OrganizationsOrganizationIdContractorsSearchPost(ctx, organizationId).SearchContractorsRequest(searchContractorsRequest).Execute()

Поиск контрагентов

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
	organizationId := "972252c5-6551-4ec4-90ad-cb5b296490cb" // string | Идентификатор организации
	searchContractorsRequest := *openapiclient.NewSearchContractorsRequest() // SearchContractorsRequest | Параметры поиска

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElbaAPI.OrganizationsOrganizationIdContractorsSearchPost(context.Background(), organizationId).SearchContractorsRequest(searchContractorsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElbaAPI.OrganizationsOrganizationIdContractorsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsOrganizationIdContractorsSearchPost`: SearchContractorsTaskResult
	fmt.Fprintf(os.Stdout, "Response from `ElbaAPI.OrganizationsOrganizationIdContractorsSearchPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Идентификатор организации | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationIdContractorsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchContractorsRequest** | [**SearchContractorsRequest**](SearchContractorsRequest.md) | Параметры поиска | 

### Return type

[**SearchContractorsTaskResult**](SearchContractorsTaskResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationIdDeliveryNotesPost

> CreateDeliveryNoteResponse OrganizationsOrganizationIdDeliveryNotesPost(ctx, organizationId).CreateDeliveryNoteRequest(createDeliveryNoteRequest).Execute()

Создание исходящей накладной

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "972252c5-6551-4ec4-90ad-cb5b296490cb" // string | Идентификатор организации
	createDeliveryNoteRequest := *openapiclient.NewCreateDeliveryNoteRequest(time.Now()) // CreateDeliveryNoteRequest | Данные накладной

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElbaAPI.OrganizationsOrganizationIdDeliveryNotesPost(context.Background(), organizationId).CreateDeliveryNoteRequest(createDeliveryNoteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElbaAPI.OrganizationsOrganizationIdDeliveryNotesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsOrganizationIdDeliveryNotesPost`: CreateDeliveryNoteResponse
	fmt.Fprintf(os.Stdout, "Response from `ElbaAPI.OrganizationsOrganizationIdDeliveryNotesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Идентификатор организации | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationIdDeliveryNotesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeliveryNoteRequest** | [**CreateDeliveryNoteRequest**](CreateDeliveryNoteRequest.md) | Данные накладной | 

### Return type

[**CreateDeliveryNoteResponse**](CreateDeliveryNoteResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationIdDocumentNewsGet

> GetDocumentNewsResponse OrganizationsOrganizationIdDocumentNewsGet(ctx, organizationId).Offset(offset).Limit(limit).Execute()

Получение списка новостей по документам

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
	organizationId := "972252c5-6551-4ec4-90ad-cb5b296490cb" // string | Идентификатор организации
	offset := int64(789) // int64 | Позиция, начиная с которой будут вычитываться новости
	limit := int32(56) // int32 | Максимальное количество новостей для получения. По-умолчанию 100 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElbaAPI.OrganizationsOrganizationIdDocumentNewsGet(context.Background(), organizationId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElbaAPI.OrganizationsOrganizationIdDocumentNewsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsOrganizationIdDocumentNewsGet`: GetDocumentNewsResponse
	fmt.Fprintf(os.Stdout, "Response from `ElbaAPI.OrganizationsOrganizationIdDocumentNewsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Идентификатор организации | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationIdDocumentNewsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int64** | Позиция, начиная с которой будут вычитываться новости | 
 **limit** | **int32** | Максимальное количество новостей для получения. По-умолчанию 100 | 

### Return type

[**GetDocumentNewsResponse**](GetDocumentNewsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationIdDocumentNewsSubscriptionPost

> OrganizationsOrganizationIdDocumentNewsSubscriptionPost(ctx, organizationId).Execute()

Подписаться на новости по документам

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
	organizationId := "972252c5-6551-4ec4-90ad-cb5b296490cb" // string | Идентификатор организации

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ElbaAPI.OrganizationsOrganizationIdDocumentNewsSubscriptionPost(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElbaAPI.OrganizationsOrganizationIdDocumentNewsSubscriptionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Идентификатор организации | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationIdDocumentNewsSubscriptionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationIdProductsPost

> CreateProductResponse OrganizationsOrganizationIdProductsPost(ctx, organizationId).CreateProductRequest(createProductRequest).Execute()

Создание товара

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
	organizationId := "972252c5-6551-4ec4-90ad-cb5b296490cb" // string | Идентификатор организации
	createProductRequest := *openapiclient.NewCreateProductRequest([]openapiclient.ProductNameToCreate{*openapiclient.NewProductNameToCreate("Бумага", true)}, []openapiclient.ProductUnitToCreate{*openapiclient.NewProductUnitToCreate(true, "шт.", float64(1))}) // CreateProductRequest | Данные товара

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElbaAPI.OrganizationsOrganizationIdProductsPost(context.Background(), organizationId).CreateProductRequest(createProductRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElbaAPI.OrganizationsOrganizationIdProductsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsOrganizationIdProductsPost`: CreateProductResponse
	fmt.Fprintf(os.Stdout, "Response from `ElbaAPI.OrganizationsOrganizationIdProductsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Идентификатор организации | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationIdProductsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createProductRequest** | [**CreateProductRequest**](CreateProductRequest.md) | Данные товара | 

### Return type

[**CreateProductResponse**](CreateProductResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationIdProductsProductIdGet

> Product OrganizationsOrganizationIdProductsProductIdGet(ctx, organizationId, productId).Execute()

Получение товара

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
	organizationId := "972252c5-6551-4ec4-90ad-cb5b296490cb" // string | Идентификатор организации
	productId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Идентификатор товара

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElbaAPI.OrganizationsOrganizationIdProductsProductIdGet(context.Background(), organizationId, productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElbaAPI.OrganizationsOrganizationIdProductsProductIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsOrganizationIdProductsProductIdGet`: Product
	fmt.Fprintf(os.Stdout, "Response from `ElbaAPI.OrganizationsOrganizationIdProductsProductIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Идентификатор организации | 
**productId** | **string** | Идентификатор товара | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationIdProductsProductIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Product**](Product.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationIdProductsProductIdUpdatePost

> AcceptedLongRunningTask OrganizationsOrganizationIdProductsProductIdUpdatePost(ctx, organizationId, productId).UpdateProductRequest(updateProductRequest).Execute()

Обновление товара

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
	organizationId := "972252c5-6551-4ec4-90ad-cb5b296490cb" // string | Идентификатор организации
	productId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Индентификатор товара
	updateProductRequest := *openapiclient.NewUpdateProductRequest([]openapiclient.ProductNameToUpdate{*openapiclient.NewProductNameToUpdate("Бумага", true)}, []openapiclient.ProductUnitToUpdate{*openapiclient.NewProductUnitToUpdate(true, "шт.", int32(1))}) // UpdateProductRequest | Данные товара

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElbaAPI.OrganizationsOrganizationIdProductsProductIdUpdatePost(context.Background(), organizationId, productId).UpdateProductRequest(updateProductRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElbaAPI.OrganizationsOrganizationIdProductsProductIdUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsOrganizationIdProductsProductIdUpdatePost`: AcceptedLongRunningTask
	fmt.Fprintf(os.Stdout, "Response from `ElbaAPI.OrganizationsOrganizationIdProductsProductIdUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Идентификатор организации | 
**productId** | **string** | Индентификатор товара | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationIdProductsProductIdUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateProductRequest** | [**UpdateProductRequest**](UpdateProductRequest.md) | Данные товара | 

### Return type

[**AcceptedLongRunningTask**](AcceptedLongRunningTask.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationIdProductsSearchPost

> SearchProductsTaskResult OrganizationsOrganizationIdProductsSearchPost(ctx, organizationId).SearchProductsRequest(searchProductsRequest).Execute()

Поиск товаров

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
	organizationId := "972252c5-6551-4ec4-90ad-cb5b296490cb" // string | Идентификатор организации
	searchProductsRequest := *openapiclient.NewSearchProductsRequest() // SearchProductsRequest | Параметры поиска. Должен быть заполнен хотябы один

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElbaAPI.OrganizationsOrganizationIdProductsSearchPost(context.Background(), organizationId).SearchProductsRequest(searchProductsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElbaAPI.OrganizationsOrganizationIdProductsSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsOrganizationIdProductsSearchPost`: SearchProductsTaskResult
	fmt.Fprintf(os.Stdout, "Response from `ElbaAPI.OrganizationsOrganizationIdProductsSearchPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Идентификатор организации | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationIdProductsSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchProductsRequest** | [**SearchProductsRequest**](SearchProductsRequest.md) | Параметры поиска. Должен быть заполнен хотябы один | 

### Return type

[**SearchProductsTaskResult**](SearchProductsTaskResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationsOrganizationIdUniversalTransferDocumentsPost

> CreateUniversalTransferDocumentResponse OrganizationsOrganizationIdUniversalTransferDocumentsPost(ctx, organizationId).CreateUniversalTransferDocumentRequest(createUniversalTransferDocumentRequest).Execute()

Создание исходящего универсального передаточного документа

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "972252c5-6551-4ec4-90ad-cb5b296490cb" // string | Идентификатор организации
	createUniversalTransferDocumentRequest := *openapiclient.NewCreateUniversalTransferDocumentRequest(time.Now(), openapiclient.UniversalTransferDocumentStatus("status1")) // CreateUniversalTransferDocumentRequest | Данные универсального передаточного документа

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElbaAPI.OrganizationsOrganizationIdUniversalTransferDocumentsPost(context.Background(), organizationId).CreateUniversalTransferDocumentRequest(createUniversalTransferDocumentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElbaAPI.OrganizationsOrganizationIdUniversalTransferDocumentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationsOrganizationIdUniversalTransferDocumentsPost`: CreateUniversalTransferDocumentResponse
	fmt.Fprintf(os.Stdout, "Response from `ElbaAPI.OrganizationsOrganizationIdUniversalTransferDocumentsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Идентификатор организации | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationsOrganizationIdUniversalTransferDocumentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createUniversalTransferDocumentRequest** | [**CreateUniversalTransferDocumentRequest**](CreateUniversalTransferDocumentRequest.md) | Данные универсального передаточного документа | 

### Return type

[**CreateUniversalTransferDocumentResponse**](CreateUniversalTransferDocumentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

