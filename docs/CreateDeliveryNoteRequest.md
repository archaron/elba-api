# CreateDeliveryNoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **NullableString** | Номер накладной. Если номер не указан, то номер будет сгенерирован Эльбой | [optional] 
**Date** | **string** | Дата накладной | 
**WithNDS** | Pointer to **NullableBool** | Ндс | [optional] 
**BankAccountId** | Pointer to **NullableString** | Идентификатор банковского счёта | [optional] 
**ContractorId** | Pointer to **NullableString** | Идентификатор контрагента | [optional] 
**SumsWithNDS** | Pointer to **NullableBool** | Цена за единицу товара включает в себя НДС | [optional] 
**ConsigneeId** | Pointer to **NullableString** | Идентификатор грузополучателя | [optional] 
**ReasonName** | Pointer to **NullableString** | Основание | [optional] 
**ReasonDate** | Pointer to **NullableString** | Дата основания | [optional] 
**WithDiscount** | Pointer to **NullableBool** | Выставить накладную со скидкой | [optional] 
**WarehouseItems** | Pointer to [**[]DeliveryNoteItemToCreate**](DeliveryNoteItemToCreate.md) | Фактурная часть | [optional] 

## Methods

### NewCreateDeliveryNoteRequest

`func NewCreateDeliveryNoteRequest(date string, ) *CreateDeliveryNoteRequest`

NewCreateDeliveryNoteRequest instantiates a new CreateDeliveryNoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeliveryNoteRequestWithDefaults

`func NewCreateDeliveryNoteRequestWithDefaults() *CreateDeliveryNoteRequest`

NewCreateDeliveryNoteRequestWithDefaults instantiates a new CreateDeliveryNoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *CreateDeliveryNoteRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CreateDeliveryNoteRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CreateDeliveryNoteRequest) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CreateDeliveryNoteRequest) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *CreateDeliveryNoteRequest) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *CreateDeliveryNoteRequest) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetDate

`func (o *CreateDeliveryNoteRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CreateDeliveryNoteRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CreateDeliveryNoteRequest) SetDate(v string)`

SetDate sets Date field to given value.


### GetWithNDS

`func (o *CreateDeliveryNoteRequest) GetWithNDS() bool`

GetWithNDS returns the WithNDS field if non-nil, zero value otherwise.

### GetWithNDSOk

`func (o *CreateDeliveryNoteRequest) GetWithNDSOk() (*bool, bool)`

GetWithNDSOk returns a tuple with the WithNDS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithNDS

`func (o *CreateDeliveryNoteRequest) SetWithNDS(v bool)`

SetWithNDS sets WithNDS field to given value.

### HasWithNDS

`func (o *CreateDeliveryNoteRequest) HasWithNDS() bool`

HasWithNDS returns a boolean if a field has been set.

### SetWithNDSNil

`func (o *CreateDeliveryNoteRequest) SetWithNDSNil(b bool)`

 SetWithNDSNil sets the value for WithNDS to be an explicit nil

### UnsetWithNDS
`func (o *CreateDeliveryNoteRequest) UnsetWithNDS()`

UnsetWithNDS ensures that no value is present for WithNDS, not even an explicit nil
### GetBankAccountId

`func (o *CreateDeliveryNoteRequest) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *CreateDeliveryNoteRequest) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *CreateDeliveryNoteRequest) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *CreateDeliveryNoteRequest) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### SetBankAccountIdNil

`func (o *CreateDeliveryNoteRequest) SetBankAccountIdNil(b bool)`

 SetBankAccountIdNil sets the value for BankAccountId to be an explicit nil

### UnsetBankAccountId
`func (o *CreateDeliveryNoteRequest) UnsetBankAccountId()`

UnsetBankAccountId ensures that no value is present for BankAccountId, not even an explicit nil
### GetContractorId

`func (o *CreateDeliveryNoteRequest) GetContractorId() string`

GetContractorId returns the ContractorId field if non-nil, zero value otherwise.

### GetContractorIdOk

`func (o *CreateDeliveryNoteRequest) GetContractorIdOk() (*string, bool)`

GetContractorIdOk returns a tuple with the ContractorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractorId

`func (o *CreateDeliveryNoteRequest) SetContractorId(v string)`

SetContractorId sets ContractorId field to given value.

### HasContractorId

`func (o *CreateDeliveryNoteRequest) HasContractorId() bool`

HasContractorId returns a boolean if a field has been set.

### SetContractorIdNil

`func (o *CreateDeliveryNoteRequest) SetContractorIdNil(b bool)`

 SetContractorIdNil sets the value for ContractorId to be an explicit nil

### UnsetContractorId
`func (o *CreateDeliveryNoteRequest) UnsetContractorId()`

UnsetContractorId ensures that no value is present for ContractorId, not even an explicit nil
### GetSumsWithNDS

`func (o *CreateDeliveryNoteRequest) GetSumsWithNDS() bool`

GetSumsWithNDS returns the SumsWithNDS field if non-nil, zero value otherwise.

### GetSumsWithNDSOk

`func (o *CreateDeliveryNoteRequest) GetSumsWithNDSOk() (*bool, bool)`

GetSumsWithNDSOk returns a tuple with the SumsWithNDS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumsWithNDS

`func (o *CreateDeliveryNoteRequest) SetSumsWithNDS(v bool)`

SetSumsWithNDS sets SumsWithNDS field to given value.

### HasSumsWithNDS

`func (o *CreateDeliveryNoteRequest) HasSumsWithNDS() bool`

HasSumsWithNDS returns a boolean if a field has been set.

### SetSumsWithNDSNil

`func (o *CreateDeliveryNoteRequest) SetSumsWithNDSNil(b bool)`

 SetSumsWithNDSNil sets the value for SumsWithNDS to be an explicit nil

### UnsetSumsWithNDS
`func (o *CreateDeliveryNoteRequest) UnsetSumsWithNDS()`

UnsetSumsWithNDS ensures that no value is present for SumsWithNDS, not even an explicit nil
### GetConsigneeId

`func (o *CreateDeliveryNoteRequest) GetConsigneeId() string`

GetConsigneeId returns the ConsigneeId field if non-nil, zero value otherwise.

### GetConsigneeIdOk

`func (o *CreateDeliveryNoteRequest) GetConsigneeIdOk() (*string, bool)`

GetConsigneeIdOk returns a tuple with the ConsigneeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsigneeId

`func (o *CreateDeliveryNoteRequest) SetConsigneeId(v string)`

SetConsigneeId sets ConsigneeId field to given value.

### HasConsigneeId

`func (o *CreateDeliveryNoteRequest) HasConsigneeId() bool`

HasConsigneeId returns a boolean if a field has been set.

### SetConsigneeIdNil

`func (o *CreateDeliveryNoteRequest) SetConsigneeIdNil(b bool)`

 SetConsigneeIdNil sets the value for ConsigneeId to be an explicit nil

### UnsetConsigneeId
`func (o *CreateDeliveryNoteRequest) UnsetConsigneeId()`

UnsetConsigneeId ensures that no value is present for ConsigneeId, not even an explicit nil
### GetReasonName

`func (o *CreateDeliveryNoteRequest) GetReasonName() string`

GetReasonName returns the ReasonName field if non-nil, zero value otherwise.

### GetReasonNameOk

`func (o *CreateDeliveryNoteRequest) GetReasonNameOk() (*string, bool)`

GetReasonNameOk returns a tuple with the ReasonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonName

`func (o *CreateDeliveryNoteRequest) SetReasonName(v string)`

SetReasonName sets ReasonName field to given value.

### HasReasonName

`func (o *CreateDeliveryNoteRequest) HasReasonName() bool`

HasReasonName returns a boolean if a field has been set.

### SetReasonNameNil

`func (o *CreateDeliveryNoteRequest) SetReasonNameNil(b bool)`

 SetReasonNameNil sets the value for ReasonName to be an explicit nil

### UnsetReasonName
`func (o *CreateDeliveryNoteRequest) UnsetReasonName()`

UnsetReasonName ensures that no value is present for ReasonName, not even an explicit nil
### GetReasonDate

`func (o *CreateDeliveryNoteRequest) GetReasonDate() string`

GetReasonDate returns the ReasonDate field if non-nil, zero value otherwise.

### GetReasonDateOk

`func (o *CreateDeliveryNoteRequest) GetReasonDateOk() (*string, bool)`

GetReasonDateOk returns a tuple with the ReasonDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonDate

`func (o *CreateDeliveryNoteRequest) SetReasonDate(v string)`

SetReasonDate sets ReasonDate field to given value.

### HasReasonDate

`func (o *CreateDeliveryNoteRequest) HasReasonDate() bool`

HasReasonDate returns a boolean if a field has been set.

### SetReasonDateNil

`func (o *CreateDeliveryNoteRequest) SetReasonDateNil(b bool)`

 SetReasonDateNil sets the value for ReasonDate to be an explicit nil

### UnsetReasonDate
`func (o *CreateDeliveryNoteRequest) UnsetReasonDate()`

UnsetReasonDate ensures that no value is present for ReasonDate, not even an explicit nil
### GetWithDiscount

`func (o *CreateDeliveryNoteRequest) GetWithDiscount() bool`

GetWithDiscount returns the WithDiscount field if non-nil, zero value otherwise.

### GetWithDiscountOk

`func (o *CreateDeliveryNoteRequest) GetWithDiscountOk() (*bool, bool)`

GetWithDiscountOk returns a tuple with the WithDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithDiscount

`func (o *CreateDeliveryNoteRequest) SetWithDiscount(v bool)`

SetWithDiscount sets WithDiscount field to given value.

### HasWithDiscount

`func (o *CreateDeliveryNoteRequest) HasWithDiscount() bool`

HasWithDiscount returns a boolean if a field has been set.

### SetWithDiscountNil

`func (o *CreateDeliveryNoteRequest) SetWithDiscountNil(b bool)`

 SetWithDiscountNil sets the value for WithDiscount to be an explicit nil

### UnsetWithDiscount
`func (o *CreateDeliveryNoteRequest) UnsetWithDiscount()`

UnsetWithDiscount ensures that no value is present for WithDiscount, not even an explicit nil
### GetWarehouseItems

`func (o *CreateDeliveryNoteRequest) GetWarehouseItems() []DeliveryNoteItemToCreate`

GetWarehouseItems returns the WarehouseItems field if non-nil, zero value otherwise.

### GetWarehouseItemsOk

`func (o *CreateDeliveryNoteRequest) GetWarehouseItemsOk() (*[]DeliveryNoteItemToCreate, bool)`

GetWarehouseItemsOk returns a tuple with the WarehouseItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseItems

`func (o *CreateDeliveryNoteRequest) SetWarehouseItems(v []DeliveryNoteItemToCreate)`

SetWarehouseItems sets WarehouseItems field to given value.

### HasWarehouseItems

`func (o *CreateDeliveryNoteRequest) HasWarehouseItems() bool`

HasWarehouseItems returns a boolean if a field has been set.

### SetWarehouseItemsNil

`func (o *CreateDeliveryNoteRequest) SetWarehouseItemsNil(b bool)`

 SetWarehouseItemsNil sets the value for WarehouseItems to be an explicit nil

### UnsetWarehouseItems
`func (o *CreateDeliveryNoteRequest) UnsetWarehouseItems()`

UnsetWarehouseItems ensures that no value is present for WarehouseItems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


