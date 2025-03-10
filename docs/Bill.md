# Bill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор счёта | 
**Number** | **string** | Номер счёта | 
**Date** | **string** | Дата счёта | 
**WithNDS** | Pointer to **NullableBool** | Ндс | [optional] 
**WithDiscount** | Pointer to **NullableBool** | Cчёт выставлен со скидкой | [optional] 
**Comment** | Pointer to **NullableString** | Комментарий | [optional] 
**BankAccountId** | Pointer to **NullableString** | Идентификатор банковского счёта | [optional] 
**ContractorId** | Pointer to **NullableString** | Идентификатор контрагента | [optional] 
**WarehouseItems** | Pointer to [**[]BillItem**](BillItem.md) | Фактурная часть | [optional] 
**SumsWithNds** | Pointer to **NullableBool** | Цена за единицу товара включает в себя НДС | [optional] 
**Status** | [**BillStatus**](BillStatus.md) | Статус оплаты.  notPaid (Не оплачен)  paid (Оплачен)  partiallyPaid (Частично оплачен)  rejected (Отменён)  overdue (Истёк срок оплаты) | 

## Methods

### NewBill

`func NewBill(id string, number string, date string, status BillStatus, ) *Bill`

NewBill instantiates a new Bill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillWithDefaults

`func NewBillWithDefaults() *Bill`

NewBillWithDefaults instantiates a new Bill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Bill) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bill) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bill) SetId(v string)`

SetId sets Id field to given value.


### GetNumber

`func (o *Bill) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Bill) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Bill) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetDate

`func (o *Bill) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Bill) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Bill) SetDate(v string)`

SetDate sets Date field to given value.


### GetWithNDS

`func (o *Bill) GetWithNDS() bool`

GetWithNDS returns the WithNDS field if non-nil, zero value otherwise.

### GetWithNDSOk

`func (o *Bill) GetWithNDSOk() (*bool, bool)`

GetWithNDSOk returns a tuple with the WithNDS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithNDS

`func (o *Bill) SetWithNDS(v bool)`

SetWithNDS sets WithNDS field to given value.

### HasWithNDS

`func (o *Bill) HasWithNDS() bool`

HasWithNDS returns a boolean if a field has been set.

### SetWithNDSNil

`func (o *Bill) SetWithNDSNil(b bool)`

 SetWithNDSNil sets the value for WithNDS to be an explicit nil

### UnsetWithNDS
`func (o *Bill) UnsetWithNDS()`

UnsetWithNDS ensures that no value is present for WithNDS, not even an explicit nil
### GetWithDiscount

`func (o *Bill) GetWithDiscount() bool`

GetWithDiscount returns the WithDiscount field if non-nil, zero value otherwise.

### GetWithDiscountOk

`func (o *Bill) GetWithDiscountOk() (*bool, bool)`

GetWithDiscountOk returns a tuple with the WithDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithDiscount

`func (o *Bill) SetWithDiscount(v bool)`

SetWithDiscount sets WithDiscount field to given value.

### HasWithDiscount

`func (o *Bill) HasWithDiscount() bool`

HasWithDiscount returns a boolean if a field has been set.

### SetWithDiscountNil

`func (o *Bill) SetWithDiscountNil(b bool)`

 SetWithDiscountNil sets the value for WithDiscount to be an explicit nil

### UnsetWithDiscount
`func (o *Bill) UnsetWithDiscount()`

UnsetWithDiscount ensures that no value is present for WithDiscount, not even an explicit nil
### GetComment

`func (o *Bill) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Bill) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Bill) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Bill) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *Bill) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *Bill) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetBankAccountId

`func (o *Bill) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *Bill) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *Bill) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *Bill) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### SetBankAccountIdNil

`func (o *Bill) SetBankAccountIdNil(b bool)`

 SetBankAccountIdNil sets the value for BankAccountId to be an explicit nil

### UnsetBankAccountId
`func (o *Bill) UnsetBankAccountId()`

UnsetBankAccountId ensures that no value is present for BankAccountId, not even an explicit nil
### GetContractorId

`func (o *Bill) GetContractorId() string`

GetContractorId returns the ContractorId field if non-nil, zero value otherwise.

### GetContractorIdOk

`func (o *Bill) GetContractorIdOk() (*string, bool)`

GetContractorIdOk returns a tuple with the ContractorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractorId

`func (o *Bill) SetContractorId(v string)`

SetContractorId sets ContractorId field to given value.

### HasContractorId

`func (o *Bill) HasContractorId() bool`

HasContractorId returns a boolean if a field has been set.

### SetContractorIdNil

`func (o *Bill) SetContractorIdNil(b bool)`

 SetContractorIdNil sets the value for ContractorId to be an explicit nil

### UnsetContractorId
`func (o *Bill) UnsetContractorId()`

UnsetContractorId ensures that no value is present for ContractorId, not even an explicit nil
### GetWarehouseItems

`func (o *Bill) GetWarehouseItems() []BillItem`

GetWarehouseItems returns the WarehouseItems field if non-nil, zero value otherwise.

### GetWarehouseItemsOk

`func (o *Bill) GetWarehouseItemsOk() (*[]BillItem, bool)`

GetWarehouseItemsOk returns a tuple with the WarehouseItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseItems

`func (o *Bill) SetWarehouseItems(v []BillItem)`

SetWarehouseItems sets WarehouseItems field to given value.

### HasWarehouseItems

`func (o *Bill) HasWarehouseItems() bool`

HasWarehouseItems returns a boolean if a field has been set.

### SetWarehouseItemsNil

`func (o *Bill) SetWarehouseItemsNil(b bool)`

 SetWarehouseItemsNil sets the value for WarehouseItems to be an explicit nil

### UnsetWarehouseItems
`func (o *Bill) UnsetWarehouseItems()`

UnsetWarehouseItems ensures that no value is present for WarehouseItems, not even an explicit nil
### GetSumsWithNds

`func (o *Bill) GetSumsWithNds() bool`

GetSumsWithNds returns the SumsWithNds field if non-nil, zero value otherwise.

### GetSumsWithNdsOk

`func (o *Bill) GetSumsWithNdsOk() (*bool, bool)`

GetSumsWithNdsOk returns a tuple with the SumsWithNds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumsWithNds

`func (o *Bill) SetSumsWithNds(v bool)`

SetSumsWithNds sets SumsWithNds field to given value.

### HasSumsWithNds

`func (o *Bill) HasSumsWithNds() bool`

HasSumsWithNds returns a boolean if a field has been set.

### SetSumsWithNdsNil

`func (o *Bill) SetSumsWithNdsNil(b bool)`

 SetSumsWithNdsNil sets the value for SumsWithNds to be an explicit nil

### UnsetSumsWithNds
`func (o *Bill) UnsetSumsWithNds()`

UnsetSumsWithNds ensures that no value is present for SumsWithNds, not even an explicit nil
### GetStatus

`func (o *Bill) GetStatus() BillStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Bill) GetStatusOk() (*BillStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Bill) SetStatus(v BillStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


