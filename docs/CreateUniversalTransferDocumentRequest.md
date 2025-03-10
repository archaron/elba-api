# CreateUniversalTransferDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **NullableString** | Номер УПД. Если номер не указан, то номер будет сгенерирован Эльбой | [optional] 
**Date** | **string** | Дата УПД | 
**BankAccountId** | Pointer to **NullableString** | Идентификатор банковского счёта | [optional] 
**ContractorId** | Pointer to **NullableString** | Идентификатор контрагента | [optional] 
**SumsWithNDS** | Pointer to **NullableBool** | Цена за единицу товара включает в себя НДС | [optional] 
**ConsigneeId** | Pointer to **NullableString** | Идентификатор грузополучателя | [optional] 
**WithProductTracing** | Pointer to **NullableBool** | Продаю прослеживаемый товар | [optional] 
**Status** | [**UniversalTransferDocumentStatus**](UniversalTransferDocumentStatus.md) | Статус  status1 (Со счётом-фактурой, то есть с НДС)  status2 (Без счёта-фактуры) | 
**StateContractIdentifier** | Pointer to **NullableString** | Идентификатор госконтракта | [optional] 
**ReasonName** | Pointer to **NullableString** | Основание | [optional] 
**ReasonDate** | Pointer to **NullableString** | Дата основания | [optional] 
**WithDiscount** | Pointer to **NullableBool** | Выставить УПД со скидкой | [optional] 
**WarehouseItems** | Pointer to [**[]UniversalTransferDocumentItemToCreate**](UniversalTransferDocumentItemToCreate.md) | Фактурная часть | [optional] 

## Methods

### NewCreateUniversalTransferDocumentRequest

`func NewCreateUniversalTransferDocumentRequest(date string, status UniversalTransferDocumentStatus, ) *CreateUniversalTransferDocumentRequest`

NewCreateUniversalTransferDocumentRequest instantiates a new CreateUniversalTransferDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUniversalTransferDocumentRequestWithDefaults

`func NewCreateUniversalTransferDocumentRequestWithDefaults() *CreateUniversalTransferDocumentRequest`

NewCreateUniversalTransferDocumentRequestWithDefaults instantiates a new CreateUniversalTransferDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *CreateUniversalTransferDocumentRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CreateUniversalTransferDocumentRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CreateUniversalTransferDocumentRequest) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CreateUniversalTransferDocumentRequest) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *CreateUniversalTransferDocumentRequest) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *CreateUniversalTransferDocumentRequest) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetDate

`func (o *CreateUniversalTransferDocumentRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CreateUniversalTransferDocumentRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CreateUniversalTransferDocumentRequest) SetDate(v string)`

SetDate sets Date field to given value.


### GetBankAccountId

`func (o *CreateUniversalTransferDocumentRequest) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *CreateUniversalTransferDocumentRequest) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *CreateUniversalTransferDocumentRequest) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *CreateUniversalTransferDocumentRequest) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### SetBankAccountIdNil

`func (o *CreateUniversalTransferDocumentRequest) SetBankAccountIdNil(b bool)`

 SetBankAccountIdNil sets the value for BankAccountId to be an explicit nil

### UnsetBankAccountId
`func (o *CreateUniversalTransferDocumentRequest) UnsetBankAccountId()`

UnsetBankAccountId ensures that no value is present for BankAccountId, not even an explicit nil
### GetContractorId

`func (o *CreateUniversalTransferDocumentRequest) GetContractorId() string`

GetContractorId returns the ContractorId field if non-nil, zero value otherwise.

### GetContractorIdOk

`func (o *CreateUniversalTransferDocumentRequest) GetContractorIdOk() (*string, bool)`

GetContractorIdOk returns a tuple with the ContractorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractorId

`func (o *CreateUniversalTransferDocumentRequest) SetContractorId(v string)`

SetContractorId sets ContractorId field to given value.

### HasContractorId

`func (o *CreateUniversalTransferDocumentRequest) HasContractorId() bool`

HasContractorId returns a boolean if a field has been set.

### SetContractorIdNil

`func (o *CreateUniversalTransferDocumentRequest) SetContractorIdNil(b bool)`

 SetContractorIdNil sets the value for ContractorId to be an explicit nil

### UnsetContractorId
`func (o *CreateUniversalTransferDocumentRequest) UnsetContractorId()`

UnsetContractorId ensures that no value is present for ContractorId, not even an explicit nil
### GetSumsWithNDS

`func (o *CreateUniversalTransferDocumentRequest) GetSumsWithNDS() bool`

GetSumsWithNDS returns the SumsWithNDS field if non-nil, zero value otherwise.

### GetSumsWithNDSOk

`func (o *CreateUniversalTransferDocumentRequest) GetSumsWithNDSOk() (*bool, bool)`

GetSumsWithNDSOk returns a tuple with the SumsWithNDS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumsWithNDS

`func (o *CreateUniversalTransferDocumentRequest) SetSumsWithNDS(v bool)`

SetSumsWithNDS sets SumsWithNDS field to given value.

### HasSumsWithNDS

`func (o *CreateUniversalTransferDocumentRequest) HasSumsWithNDS() bool`

HasSumsWithNDS returns a boolean if a field has been set.

### SetSumsWithNDSNil

`func (o *CreateUniversalTransferDocumentRequest) SetSumsWithNDSNil(b bool)`

 SetSumsWithNDSNil sets the value for SumsWithNDS to be an explicit nil

### UnsetSumsWithNDS
`func (o *CreateUniversalTransferDocumentRequest) UnsetSumsWithNDS()`

UnsetSumsWithNDS ensures that no value is present for SumsWithNDS, not even an explicit nil
### GetConsigneeId

`func (o *CreateUniversalTransferDocumentRequest) GetConsigneeId() string`

GetConsigneeId returns the ConsigneeId field if non-nil, zero value otherwise.

### GetConsigneeIdOk

`func (o *CreateUniversalTransferDocumentRequest) GetConsigneeIdOk() (*string, bool)`

GetConsigneeIdOk returns a tuple with the ConsigneeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsigneeId

`func (o *CreateUniversalTransferDocumentRequest) SetConsigneeId(v string)`

SetConsigneeId sets ConsigneeId field to given value.

### HasConsigneeId

`func (o *CreateUniversalTransferDocumentRequest) HasConsigneeId() bool`

HasConsigneeId returns a boolean if a field has been set.

### SetConsigneeIdNil

`func (o *CreateUniversalTransferDocumentRequest) SetConsigneeIdNil(b bool)`

 SetConsigneeIdNil sets the value for ConsigneeId to be an explicit nil

### UnsetConsigneeId
`func (o *CreateUniversalTransferDocumentRequest) UnsetConsigneeId()`

UnsetConsigneeId ensures that no value is present for ConsigneeId, not even an explicit nil
### GetWithProductTracing

`func (o *CreateUniversalTransferDocumentRequest) GetWithProductTracing() bool`

GetWithProductTracing returns the WithProductTracing field if non-nil, zero value otherwise.

### GetWithProductTracingOk

`func (o *CreateUniversalTransferDocumentRequest) GetWithProductTracingOk() (*bool, bool)`

GetWithProductTracingOk returns a tuple with the WithProductTracing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithProductTracing

`func (o *CreateUniversalTransferDocumentRequest) SetWithProductTracing(v bool)`

SetWithProductTracing sets WithProductTracing field to given value.

### HasWithProductTracing

`func (o *CreateUniversalTransferDocumentRequest) HasWithProductTracing() bool`

HasWithProductTracing returns a boolean if a field has been set.

### SetWithProductTracingNil

`func (o *CreateUniversalTransferDocumentRequest) SetWithProductTracingNil(b bool)`

 SetWithProductTracingNil sets the value for WithProductTracing to be an explicit nil

### UnsetWithProductTracing
`func (o *CreateUniversalTransferDocumentRequest) UnsetWithProductTracing()`

UnsetWithProductTracing ensures that no value is present for WithProductTracing, not even an explicit nil
### GetStatus

`func (o *CreateUniversalTransferDocumentRequest) GetStatus() UniversalTransferDocumentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateUniversalTransferDocumentRequest) GetStatusOk() (*UniversalTransferDocumentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateUniversalTransferDocumentRequest) SetStatus(v UniversalTransferDocumentStatus)`

SetStatus sets Status field to given value.


### GetStateContractIdentifier

`func (o *CreateUniversalTransferDocumentRequest) GetStateContractIdentifier() string`

GetStateContractIdentifier returns the StateContractIdentifier field if non-nil, zero value otherwise.

### GetStateContractIdentifierOk

`func (o *CreateUniversalTransferDocumentRequest) GetStateContractIdentifierOk() (*string, bool)`

GetStateContractIdentifierOk returns a tuple with the StateContractIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateContractIdentifier

`func (o *CreateUniversalTransferDocumentRequest) SetStateContractIdentifier(v string)`

SetStateContractIdentifier sets StateContractIdentifier field to given value.

### HasStateContractIdentifier

`func (o *CreateUniversalTransferDocumentRequest) HasStateContractIdentifier() bool`

HasStateContractIdentifier returns a boolean if a field has been set.

### SetStateContractIdentifierNil

`func (o *CreateUniversalTransferDocumentRequest) SetStateContractIdentifierNil(b bool)`

 SetStateContractIdentifierNil sets the value for StateContractIdentifier to be an explicit nil

### UnsetStateContractIdentifier
`func (o *CreateUniversalTransferDocumentRequest) UnsetStateContractIdentifier()`

UnsetStateContractIdentifier ensures that no value is present for StateContractIdentifier, not even an explicit nil
### GetReasonName

`func (o *CreateUniversalTransferDocumentRequest) GetReasonName() string`

GetReasonName returns the ReasonName field if non-nil, zero value otherwise.

### GetReasonNameOk

`func (o *CreateUniversalTransferDocumentRequest) GetReasonNameOk() (*string, bool)`

GetReasonNameOk returns a tuple with the ReasonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonName

`func (o *CreateUniversalTransferDocumentRequest) SetReasonName(v string)`

SetReasonName sets ReasonName field to given value.

### HasReasonName

`func (o *CreateUniversalTransferDocumentRequest) HasReasonName() bool`

HasReasonName returns a boolean if a field has been set.

### SetReasonNameNil

`func (o *CreateUniversalTransferDocumentRequest) SetReasonNameNil(b bool)`

 SetReasonNameNil sets the value for ReasonName to be an explicit nil

### UnsetReasonName
`func (o *CreateUniversalTransferDocumentRequest) UnsetReasonName()`

UnsetReasonName ensures that no value is present for ReasonName, not even an explicit nil
### GetReasonDate

`func (o *CreateUniversalTransferDocumentRequest) GetReasonDate() string`

GetReasonDate returns the ReasonDate field if non-nil, zero value otherwise.

### GetReasonDateOk

`func (o *CreateUniversalTransferDocumentRequest) GetReasonDateOk() (*string, bool)`

GetReasonDateOk returns a tuple with the ReasonDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonDate

`func (o *CreateUniversalTransferDocumentRequest) SetReasonDate(v string)`

SetReasonDate sets ReasonDate field to given value.

### HasReasonDate

`func (o *CreateUniversalTransferDocumentRequest) HasReasonDate() bool`

HasReasonDate returns a boolean if a field has been set.

### SetReasonDateNil

`func (o *CreateUniversalTransferDocumentRequest) SetReasonDateNil(b bool)`

 SetReasonDateNil sets the value for ReasonDate to be an explicit nil

### UnsetReasonDate
`func (o *CreateUniversalTransferDocumentRequest) UnsetReasonDate()`

UnsetReasonDate ensures that no value is present for ReasonDate, not even an explicit nil
### GetWithDiscount

`func (o *CreateUniversalTransferDocumentRequest) GetWithDiscount() bool`

GetWithDiscount returns the WithDiscount field if non-nil, zero value otherwise.

### GetWithDiscountOk

`func (o *CreateUniversalTransferDocumentRequest) GetWithDiscountOk() (*bool, bool)`

GetWithDiscountOk returns a tuple with the WithDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithDiscount

`func (o *CreateUniversalTransferDocumentRequest) SetWithDiscount(v bool)`

SetWithDiscount sets WithDiscount field to given value.

### HasWithDiscount

`func (o *CreateUniversalTransferDocumentRequest) HasWithDiscount() bool`

HasWithDiscount returns a boolean if a field has been set.

### SetWithDiscountNil

`func (o *CreateUniversalTransferDocumentRequest) SetWithDiscountNil(b bool)`

 SetWithDiscountNil sets the value for WithDiscount to be an explicit nil

### UnsetWithDiscount
`func (o *CreateUniversalTransferDocumentRequest) UnsetWithDiscount()`

UnsetWithDiscount ensures that no value is present for WithDiscount, not even an explicit nil
### GetWarehouseItems

`func (o *CreateUniversalTransferDocumentRequest) GetWarehouseItems() []UniversalTransferDocumentItemToCreate`

GetWarehouseItems returns the WarehouseItems field if non-nil, zero value otherwise.

### GetWarehouseItemsOk

`func (o *CreateUniversalTransferDocumentRequest) GetWarehouseItemsOk() (*[]UniversalTransferDocumentItemToCreate, bool)`

GetWarehouseItemsOk returns a tuple with the WarehouseItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseItems

`func (o *CreateUniversalTransferDocumentRequest) SetWarehouseItems(v []UniversalTransferDocumentItemToCreate)`

SetWarehouseItems sets WarehouseItems field to given value.

### HasWarehouseItems

`func (o *CreateUniversalTransferDocumentRequest) HasWarehouseItems() bool`

HasWarehouseItems returns a boolean if a field has been set.

### SetWarehouseItemsNil

`func (o *CreateUniversalTransferDocumentRequest) SetWarehouseItemsNil(b bool)`

 SetWarehouseItemsNil sets the value for WarehouseItems to be an explicit nil

### UnsetWarehouseItems
`func (o *CreateUniversalTransferDocumentRequest) UnsetWarehouseItems()`

UnsetWarehouseItems ensures that no value is present for WarehouseItems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


