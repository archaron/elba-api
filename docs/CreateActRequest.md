# CreateActRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **NullableString** | Номер акта. Если номер не указан, то номер будет сгенерирован Эльбой | [optional] 
**Date** | **string** | Дата акта | 
**WithNDS** | Pointer to **NullableBool** | Ндс | [optional] 
**ReasonName** | Pointer to **NullableString** | Основание | [optional] 
**Comment** | Pointer to **NullableString** | Комментарий | [optional] 
**ReasonDate** | Pointer to **NullableString** | Дата основания | [optional] 
**BankAccountId** | Pointer to **NullableString** | Идентификатор банковского счёта | [optional] 
**ContractorId** | Pointer to **NullableString** | Идентификатор контрагента | [optional] 
**SumsWithNDS** | Pointer to **NullableBool** | Цена за единицу товара включает в себя НДС | [optional] 
**WithDiscount** | Pointer to **NullableBool** | Выставить акт со скидкой | [optional] 
**WarehouseItems** | Pointer to [**[]ActItemToCreate**](ActItemToCreate.md) | Фактурная часть | [optional] 

## Methods

### NewCreateActRequest

`func NewCreateActRequest(date string, ) *CreateActRequest`

NewCreateActRequest instantiates a new CreateActRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateActRequestWithDefaults

`func NewCreateActRequestWithDefaults() *CreateActRequest`

NewCreateActRequestWithDefaults instantiates a new CreateActRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *CreateActRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CreateActRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CreateActRequest) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CreateActRequest) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *CreateActRequest) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *CreateActRequest) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetDate

`func (o *CreateActRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CreateActRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CreateActRequest) SetDate(v string)`

SetDate sets Date field to given value.


### GetWithNDS

`func (o *CreateActRequest) GetWithNDS() bool`

GetWithNDS returns the WithNDS field if non-nil, zero value otherwise.

### GetWithNDSOk

`func (o *CreateActRequest) GetWithNDSOk() (*bool, bool)`

GetWithNDSOk returns a tuple with the WithNDS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithNDS

`func (o *CreateActRequest) SetWithNDS(v bool)`

SetWithNDS sets WithNDS field to given value.

### HasWithNDS

`func (o *CreateActRequest) HasWithNDS() bool`

HasWithNDS returns a boolean if a field has been set.

### SetWithNDSNil

`func (o *CreateActRequest) SetWithNDSNil(b bool)`

 SetWithNDSNil sets the value for WithNDS to be an explicit nil

### UnsetWithNDS
`func (o *CreateActRequest) UnsetWithNDS()`

UnsetWithNDS ensures that no value is present for WithNDS, not even an explicit nil
### GetReasonName

`func (o *CreateActRequest) GetReasonName() string`

GetReasonName returns the ReasonName field if non-nil, zero value otherwise.

### GetReasonNameOk

`func (o *CreateActRequest) GetReasonNameOk() (*string, bool)`

GetReasonNameOk returns a tuple with the ReasonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonName

`func (o *CreateActRequest) SetReasonName(v string)`

SetReasonName sets ReasonName field to given value.

### HasReasonName

`func (o *CreateActRequest) HasReasonName() bool`

HasReasonName returns a boolean if a field has been set.

### SetReasonNameNil

`func (o *CreateActRequest) SetReasonNameNil(b bool)`

 SetReasonNameNil sets the value for ReasonName to be an explicit nil

### UnsetReasonName
`func (o *CreateActRequest) UnsetReasonName()`

UnsetReasonName ensures that no value is present for ReasonName, not even an explicit nil
### GetComment

`func (o *CreateActRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateActRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateActRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateActRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *CreateActRequest) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *CreateActRequest) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetReasonDate

`func (o *CreateActRequest) GetReasonDate() string`

GetReasonDate returns the ReasonDate field if non-nil, zero value otherwise.

### GetReasonDateOk

`func (o *CreateActRequest) GetReasonDateOk() (*string, bool)`

GetReasonDateOk returns a tuple with the ReasonDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonDate

`func (o *CreateActRequest) SetReasonDate(v string)`

SetReasonDate sets ReasonDate field to given value.

### HasReasonDate

`func (o *CreateActRequest) HasReasonDate() bool`

HasReasonDate returns a boolean if a field has been set.

### SetReasonDateNil

`func (o *CreateActRequest) SetReasonDateNil(b bool)`

 SetReasonDateNil sets the value for ReasonDate to be an explicit nil

### UnsetReasonDate
`func (o *CreateActRequest) UnsetReasonDate()`

UnsetReasonDate ensures that no value is present for ReasonDate, not even an explicit nil
### GetBankAccountId

`func (o *CreateActRequest) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *CreateActRequest) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *CreateActRequest) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *CreateActRequest) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### SetBankAccountIdNil

`func (o *CreateActRequest) SetBankAccountIdNil(b bool)`

 SetBankAccountIdNil sets the value for BankAccountId to be an explicit nil

### UnsetBankAccountId
`func (o *CreateActRequest) UnsetBankAccountId()`

UnsetBankAccountId ensures that no value is present for BankAccountId, not even an explicit nil
### GetContractorId

`func (o *CreateActRequest) GetContractorId() string`

GetContractorId returns the ContractorId field if non-nil, zero value otherwise.

### GetContractorIdOk

`func (o *CreateActRequest) GetContractorIdOk() (*string, bool)`

GetContractorIdOk returns a tuple with the ContractorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractorId

`func (o *CreateActRequest) SetContractorId(v string)`

SetContractorId sets ContractorId field to given value.

### HasContractorId

`func (o *CreateActRequest) HasContractorId() bool`

HasContractorId returns a boolean if a field has been set.

### SetContractorIdNil

`func (o *CreateActRequest) SetContractorIdNil(b bool)`

 SetContractorIdNil sets the value for ContractorId to be an explicit nil

### UnsetContractorId
`func (o *CreateActRequest) UnsetContractorId()`

UnsetContractorId ensures that no value is present for ContractorId, not even an explicit nil
### GetSumsWithNDS

`func (o *CreateActRequest) GetSumsWithNDS() bool`

GetSumsWithNDS returns the SumsWithNDS field if non-nil, zero value otherwise.

### GetSumsWithNDSOk

`func (o *CreateActRequest) GetSumsWithNDSOk() (*bool, bool)`

GetSumsWithNDSOk returns a tuple with the SumsWithNDS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumsWithNDS

`func (o *CreateActRequest) SetSumsWithNDS(v bool)`

SetSumsWithNDS sets SumsWithNDS field to given value.

### HasSumsWithNDS

`func (o *CreateActRequest) HasSumsWithNDS() bool`

HasSumsWithNDS returns a boolean if a field has been set.

### SetSumsWithNDSNil

`func (o *CreateActRequest) SetSumsWithNDSNil(b bool)`

 SetSumsWithNDSNil sets the value for SumsWithNDS to be an explicit nil

### UnsetSumsWithNDS
`func (o *CreateActRequest) UnsetSumsWithNDS()`

UnsetSumsWithNDS ensures that no value is present for SumsWithNDS, not even an explicit nil
### GetWithDiscount

`func (o *CreateActRequest) GetWithDiscount() bool`

GetWithDiscount returns the WithDiscount field if non-nil, zero value otherwise.

### GetWithDiscountOk

`func (o *CreateActRequest) GetWithDiscountOk() (*bool, bool)`

GetWithDiscountOk returns a tuple with the WithDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithDiscount

`func (o *CreateActRequest) SetWithDiscount(v bool)`

SetWithDiscount sets WithDiscount field to given value.

### HasWithDiscount

`func (o *CreateActRequest) HasWithDiscount() bool`

HasWithDiscount returns a boolean if a field has been set.

### SetWithDiscountNil

`func (o *CreateActRequest) SetWithDiscountNil(b bool)`

 SetWithDiscountNil sets the value for WithDiscount to be an explicit nil

### UnsetWithDiscount
`func (o *CreateActRequest) UnsetWithDiscount()`

UnsetWithDiscount ensures that no value is present for WithDiscount, not even an explicit nil
### GetWarehouseItems

`func (o *CreateActRequest) GetWarehouseItems() []ActItemToCreate`

GetWarehouseItems returns the WarehouseItems field if non-nil, zero value otherwise.

### GetWarehouseItemsOk

`func (o *CreateActRequest) GetWarehouseItemsOk() (*[]ActItemToCreate, bool)`

GetWarehouseItemsOk returns a tuple with the WarehouseItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseItems

`func (o *CreateActRequest) SetWarehouseItems(v []ActItemToCreate)`

SetWarehouseItems sets WarehouseItems field to given value.

### HasWarehouseItems

`func (o *CreateActRequest) HasWarehouseItems() bool`

HasWarehouseItems returns a boolean if a field has been set.

### SetWarehouseItemsNil

`func (o *CreateActRequest) SetWarehouseItemsNil(b bool)`

 SetWarehouseItemsNil sets the value for WarehouseItems to be an explicit nil

### UnsetWarehouseItems
`func (o *CreateActRequest) UnsetWarehouseItems()`

UnsetWarehouseItems ensures that no value is present for WarehouseItems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


