# Go API client for elba


## С чего начать

Для работы с API нужно выпустить API-ключ — уникальный токен, позволяющий авторизовывать ваши запросы в API Контур.Эльбы.

#### Как получить API-ключ

1. Откройте Эльбу, в верхнем правом углу нажмите «Настройки и оплата» → «Настройки сервиса».
2. Перейдите на вкладку «API».
2. Нажмите на кнопку «Выпустить ключ». После этого откроется окно со сгенерированным API-ключом.
3. В открывшемся окне появится ваш API-ключ. Скопируйте и сохраните его в надежном месте, потому что он будет показан только один раз. Это сделано в целях безопасности — мы не храним ключи на своей стороне.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Generator version: 7.7.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import elba "github.com/archaron/elba-api"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `elba.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), elba.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `elba.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), elba.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `elba.ContextOperationServerIndices` and `elba.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), elba.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), elba.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://elba-api.kontur.ru/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ElbaAPI* | [**LongRunningTasksTaskIdGet**](docs/ElbaAPI.md#longrunningtaskstaskidget) | **Get** /long-running-tasks/{taskId} | Получить состояние операции
*ElbaAPI* | [**OrganizationsGet**](docs/ElbaAPI.md#organizationsget) | **Get** /organizations | Получение списка доступных организаций
*ElbaAPI* | [**OrganizationsOrganizationIdActsPost**](docs/ElbaAPI.md#organizationsorganizationidactspost) | **Post** /organizations/{organizationId}/acts | Создание исходящего акта
*ElbaAPI* | [**OrganizationsOrganizationIdBankAccountsGet**](docs/ElbaAPI.md#organizationsorganizationidbankaccountsget) | **Get** /organizations/{organizationId}/bank-accounts | Получение списка банковских счетов
*ElbaAPI* | [**OrganizationsOrganizationIdBankAccountsPost**](docs/ElbaAPI.md#organizationsorganizationidbankaccountspost) | **Post** /organizations/{organizationId}/bank-accounts | Создание банковского счёта
*ElbaAPI* | [**OrganizationsOrganizationIdBillsBillIdGet**](docs/ElbaAPI.md#organizationsorganizationidbillsbillidget) | **Get** /organizations/{organizationId}/bills/{billId} | Получение исходящего счёта
*ElbaAPI* | [**OrganizationsOrganizationIdBillsPost**](docs/ElbaAPI.md#organizationsorganizationidbillspost) | **Post** /organizations/{organizationId}/bills | Создание исходящего счёта
*ElbaAPI* | [**OrganizationsOrganizationIdContractorsContractorIdGet**](docs/ElbaAPI.md#organizationsorganizationidcontractorscontractoridget) | **Get** /organizations/{organizationId}/contractors/{contractorId} | Получение контрагента
*ElbaAPI* | [**OrganizationsOrganizationIdContractorsPost**](docs/ElbaAPI.md#organizationsorganizationidcontractorspost) | **Post** /organizations/{organizationId}/contractors | Создание контрагента
*ElbaAPI* | [**OrganizationsOrganizationIdContractorsSearchPost**](docs/ElbaAPI.md#organizationsorganizationidcontractorssearchpost) | **Post** /organizations/{organizationId}/contractors/search | Поиск контрагентов
*ElbaAPI* | [**OrganizationsOrganizationIdDeliveryNotesPost**](docs/ElbaAPI.md#organizationsorganizationiddeliverynotespost) | **Post** /organizations/{organizationId}/delivery-notes | Создание исходящей накладной
*ElbaAPI* | [**OrganizationsOrganizationIdDocumentNewsGet**](docs/ElbaAPI.md#organizationsorganizationiddocumentnewsget) | **Get** /organizations/{organizationId}/document-news | Получение списка новостей по документам
*ElbaAPI* | [**OrganizationsOrganizationIdDocumentNewsSubscriptionPost**](docs/ElbaAPI.md#organizationsorganizationiddocumentnewssubscriptionpost) | **Post** /organizations/{organizationId}/document-news/subscription | Подписаться на новости по документам
*ElbaAPI* | [**OrganizationsOrganizationIdProductsPost**](docs/ElbaAPI.md#organizationsorganizationidproductspost) | **Post** /organizations/{organizationId}/products | Создание товара
*ElbaAPI* | [**OrganizationsOrganizationIdProductsProductIdGet**](docs/ElbaAPI.md#organizationsorganizationidproductsproductidget) | **Get** /organizations/{organizationId}/products/{productId} | Получение товара
*ElbaAPI* | [**OrganizationsOrganizationIdProductsProductIdUpdatePost**](docs/ElbaAPI.md#organizationsorganizationidproductsproductidupdatepost) | **Post** /organizations/{organizationId}/products/{productId}/update | Обновление товара
*ElbaAPI* | [**OrganizationsOrganizationIdProductsSearchPost**](docs/ElbaAPI.md#organizationsorganizationidproductssearchpost) | **Post** /organizations/{organizationId}/products/search | Поиск товаров
*ElbaAPI* | [**OrganizationsOrganizationIdUniversalTransferDocumentsPost**](docs/ElbaAPI.md#organizationsorganizationiduniversaltransferdocumentspost) | **Post** /organizations/{organizationId}/universal-transfer-documents | Создание исходящего универсального передаточного документа


## Documentation For Models

 - [AcceptedLongRunningTask](docs/AcceptedLongRunningTask.md)
 - [ActItemToCreate](docs/ActItemToCreate.md)
 - [Bank](docs/Bank.md)
 - [BankAccount](docs/BankAccount.md)
 - [Bill](docs/Bill.md)
 - [BillItem](docs/BillItem.md)
 - [BillItemToCreate](docs/BillItemToCreate.md)
 - [BillPaymentNewsContent](docs/BillPaymentNewsContent.md)
 - [BillStatus](docs/BillStatus.md)
 - [Contractor](docs/Contractor.md)
 - [ContractorContact](docs/ContractorContact.md)
 - [CreateActRequest](docs/CreateActRequest.md)
 - [CreateActResponse](docs/CreateActResponse.md)
 - [CreateBankAccountRequest](docs/CreateBankAccountRequest.md)
 - [CreateBankAccountResponse](docs/CreateBankAccountResponse.md)
 - [CreateBillRequest](docs/CreateBillRequest.md)
 - [CreateBillResponse](docs/CreateBillResponse.md)
 - [CreateContractorRequest](docs/CreateContractorRequest.md)
 - [CreateContractorResponse](docs/CreateContractorResponse.md)
 - [CreateDeliveryNoteRequest](docs/CreateDeliveryNoteRequest.md)
 - [CreateDeliveryNoteResponse](docs/CreateDeliveryNoteResponse.md)
 - [CreateProductError](docs/CreateProductError.md)
 - [CreateProductRequest](docs/CreateProductRequest.md)
 - [CreateProductResponse](docs/CreateProductResponse.md)
 - [CreateProductTaskResult](docs/CreateProductTaskResult.md)
 - [CreateUniversalTransferDocumentRequest](docs/CreateUniversalTransferDocumentRequest.md)
 - [CreateUniversalTransferDocumentResponse](docs/CreateUniversalTransferDocumentResponse.md)
 - [DeliveryNoteItemToCreate](docs/DeliveryNoteItemToCreate.md)
 - [DocumentNews](docs/DocumentNews.md)
 - [DocumentNewsNewsContent](docs/DocumentNewsNewsContent.md)
 - [Error](docs/Error.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [FoundContractor](docs/FoundContractor.md)
 - [FoundProduct](docs/FoundProduct.md)
 - [GetBankAccountsResponse](docs/GetBankAccountsResponse.md)
 - [GetDocumentNewsResponse](docs/GetDocumentNewsResponse.md)
 - [GetOrganizationsResponse](docs/GetOrganizationsResponse.md)
 - [LongRunningTaskExecutionError](docs/LongRunningTaskExecutionError.md)
 - [LongRunningTaskState](docs/LongRunningTaskState.md)
 - [LongRunningTaskStateResult](docs/LongRunningTaskStateResult.md)
 - [LongRunningTaskStatus](docs/LongRunningTaskStatus.md)
 - [LongRunningTaskType](docs/LongRunningTaskType.md)
 - [NDSRate](docs/NDSRate.md)
 - [NDSRateToSave](docs/NDSRateToSave.md)
 - [NewsType](docs/NewsType.md)
 - [Organization](docs/Organization.md)
 - [Product](docs/Product.md)
 - [ProductName](docs/ProductName.md)
 - [ProductNameToCreate](docs/ProductNameToCreate.md)
 - [ProductNameToUpdate](docs/ProductNameToUpdate.md)
 - [ProductType](docs/ProductType.md)
 - [ProductUnit](docs/ProductUnit.md)
 - [ProductUnitToCreate](docs/ProductUnitToCreate.md)
 - [ProductUnitToUpdate](docs/ProductUnitToUpdate.md)
 - [SearchContractorsFilter](docs/SearchContractorsFilter.md)
 - [SearchContractorsRequest](docs/SearchContractorsRequest.md)
 - [SearchContractorsTaskResult](docs/SearchContractorsTaskResult.md)
 - [SearchProductsFilter](docs/SearchProductsFilter.md)
 - [SearchProductsRequest](docs/SearchProductsRequest.md)
 - [SearchProductsTaskResult](docs/SearchProductsTaskResult.md)
 - [UniversalTransferDocumentItemToCreate](docs/UniversalTransferDocumentItemToCreate.md)
 - [UniversalTransferDocumentStatus](docs/UniversalTransferDocumentStatus.md)
 - [UpdateProductError](docs/UpdateProductError.md)
 - [UpdateProductRequest](docs/UpdateProductRequest.md)
 - [UpdateProductTaskResult](docs/UpdateProductTaskResult.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: X-Kontur-ApiKey
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-Kontur-ApiKey and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		elba.ContextAPIKeys,
		map[string]elba.APIKey{
			"X-Kontur-ApiKey": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

e@kontur.ru

