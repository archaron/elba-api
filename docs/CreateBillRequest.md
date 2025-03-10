# CreateBillRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **NullableString** | Номер счёта. Если номер не указан, то номер будет сгенерирован Эльбой | [optional] 
**Date** | **string** | Дата счёта | 
**WithNDS** | Pointer to **NullableBool** | Ндс | [optional] 
**Comment** | Pointer to **NullableString** | Комментарий | [optional] 
**BankAccountId** | Pointer to **NullableString** | Идентификатор банковского счёта | [optional] 
**ContractorId** | Pointer to **NullableString** | Идентификатор контрагента | [optional] 
**SumsWithNDS** | Pointer to **NullableBool** | Цена за единицу товара включает в себя НДС | [optional] 
**WithDiscount** | Pointer to **NullableBool** | Выставить счёт со скидкой | [optional] 
**WarehouseItems** | Pointer to [**[]BillItemToCreate**](BillItemToCreate.md) | Фактурная часть | [optional] 

## Methods

### NewCreateBillRequest

`func NewCreateBillRequest(date string, ) *CreateBillRequest`

NewCreateBillRequest instantiates a new CreateBillRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBillRequestWithDefaults

`func NewCreateBillRequestWithDefaults() *CreateBillRequest`

NewCreateBillRequestWithDefaults instantiates a new CreateBillRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *CreateBillRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CreateBillRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CreateBillRequest) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CreateBillRequest) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *CreateBillRequest) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *CreateBillRequest) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetDate

`func (o *CreateBillRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CreateBillRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CreateBillRequest) SetDate(v string)`

SetDate sets Date field to given value.


### GetWithNDS

`func (o *CreateBillRequest) GetWithNDS() bool`

GetWithNDS returns the WithNDS field if non-nil, zero value otherwise.

### GetWithNDSOk

`func (o *CreateBillRequest) GetWithNDSOk() (*bool, bool)`

GetWithNDSOk returns a tuple with the WithNDS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithNDS

`func (o *CreateBillRequest) SetWithNDS(v bool)`

SetWithNDS sets WithNDS field to given value.

### HasWithNDS

`func (o *CreateBillRequest) HasWithNDS() bool`

HasWithNDS returns a boolean if a field has been set.

### SetWithNDSNil

`func (o *CreateBillRequest) SetWithNDSNil(b bool)`

 SetWithNDSNil sets the value for WithNDS to be an explicit nil

### UnsetWithNDS
`func (o *CreateBillRequest) UnsetWithNDS()`

UnsetWithNDS ensures that no value is present for WithNDS, not even an explicit nil
### GetComment

`func (o *CreateBillRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateBillRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateBillRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateBillRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *CreateBillRequest) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *CreateBillRequest) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetBankAccountId

`func (o *CreateBillRequest) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *CreateBillRequest) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *CreateBillRequest) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *CreateBillRequest) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### SetBankAccountIdNil

`func (o *CreateBillRequest) SetBankAccountIdNil(b bool)`

 SetBankAccountIdNil sets the value for BankAccountId to be an explicit nil

### UnsetBankAccountId
`func (o *CreateBillRequest) UnsetBankAccountId()`

UnsetBankAccountId ensures that no value is present for BankAccountId, not even an explicit nil
### GetContractorId

`func (o *CreateBillRequest) GetContractorId() string`

GetContractorId returns the ContractorId field if non-nil, zero value otherwise.

### GetContractorIdOk

`func (o *CreateBillRequest) GetContractorIdOk() (*string, bool)`

GetContractorIdOk returns a tuple with the ContractorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractorId

`func (o *CreateBillRequest) SetContractorId(v string)`

SetContractorId sets ContractorId field to given value.

### HasContractorId

`func (o *CreateBillRequest) HasContractorId() bool`

HasContractorId returns a boolean if a field has been set.

### SetContractorIdNil

`func (o *CreateBillRequest) SetContractorIdNil(b bool)`

 SetContractorIdNil sets the value for ContractorId to be an explicit nil

### UnsetContractorId
`func (o *CreateBillRequest) UnsetContractorId()`

UnsetContractorId ensures that no value is present for ContractorId, not even an explicit nil
### GetSumsWithNDS

`func (o *CreateBillRequest) GetSumsWithNDS() bool`

GetSumsWithNDS returns the SumsWithNDS field if non-nil, zero value otherwise.

### GetSumsWithNDSOk

`func (o *CreateBillRequest) GetSumsWithNDSOk() (*bool, bool)`

GetSumsWithNDSOk returns a tuple with the SumsWithNDS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumsWithNDS

`func (o *CreateBillRequest) SetSumsWithNDS(v bool)`

SetSumsWithNDS sets SumsWithNDS field to given value.

### HasSumsWithNDS

`func (o *CreateBillRequest) HasSumsWithNDS() bool`

HasSumsWithNDS returns a boolean if a field has been set.

### SetSumsWithNDSNil

`func (o *CreateBillRequest) SetSumsWithNDSNil(b bool)`

 SetSumsWithNDSNil sets the value for SumsWithNDS to be an explicit nil

### UnsetSumsWithNDS
`func (o *CreateBillRequest) UnsetSumsWithNDS()`

UnsetSumsWithNDS ensures that no value is present for SumsWithNDS, not even an explicit nil
### GetWithDiscount

`func (o *CreateBillRequest) GetWithDiscount() bool`

GetWithDiscount returns the WithDiscount field if non-nil, zero value otherwise.

### GetWithDiscountOk

`func (o *CreateBillRequest) GetWithDiscountOk() (*bool, bool)`

GetWithDiscountOk returns a tuple with the WithDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithDiscount

`func (o *CreateBillRequest) SetWithDiscount(v bool)`

SetWithDiscount sets WithDiscount field to given value.

### HasWithDiscount

`func (o *CreateBillRequest) HasWithDiscount() bool`

HasWithDiscount returns a boolean if a field has been set.

### SetWithDiscountNil

`func (o *CreateBillRequest) SetWithDiscountNil(b bool)`

 SetWithDiscountNil sets the value for WithDiscount to be an explicit nil

### UnsetWithDiscount
`func (o *CreateBillRequest) UnsetWithDiscount()`

UnsetWithDiscount ensures that no value is present for WithDiscount, not even an explicit nil
### GetWarehouseItems

`func (o *CreateBillRequest) GetWarehouseItems() []BillItemToCreate`

GetWarehouseItems returns the WarehouseItems field if non-nil, zero value otherwise.

### GetWarehouseItemsOk

`func (o *CreateBillRequest) GetWarehouseItemsOk() (*[]BillItemToCreate, bool)`

GetWarehouseItemsOk returns a tuple with the WarehouseItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseItems

`func (o *CreateBillRequest) SetWarehouseItems(v []BillItemToCreate)`

SetWarehouseItems sets WarehouseItems field to given value.

### HasWarehouseItems

`func (o *CreateBillRequest) HasWarehouseItems() bool`

HasWarehouseItems returns a boolean if a field has been set.

### SetWarehouseItemsNil

`func (o *CreateBillRequest) SetWarehouseItemsNil(b bool)`

 SetWarehouseItemsNil sets the value for WarehouseItems to be an explicit nil

### UnsetWarehouseItems
`func (o *CreateBillRequest) UnsetWarehouseItems()`

UnsetWarehouseItems ensures that no value is present for WarehouseItems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


