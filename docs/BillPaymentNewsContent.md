# BillPaymentNewsContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillId** | **string** | Идентификатор счёта | 
**BillStatus** | [**BillStatus**](BillStatus.md) | Статус оплаченности счёта.  notPaid (Не оплачен)  paid (Оплачен)  partiallyPaid (Частично оплачен)  rejected (Отменён)  overdue (Истёк срок оплаты) | 
**RelatedPaymentsSum** | **float64** | Сумма привязанных к счёту операций | 
**BillSum** | **float64** | Сумма счёта | 

## Methods

### NewBillPaymentNewsContent

`func NewBillPaymentNewsContent(billId string, billStatus BillStatus, relatedPaymentsSum float64, billSum float64, ) *BillPaymentNewsContent`

NewBillPaymentNewsContent instantiates a new BillPaymentNewsContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillPaymentNewsContentWithDefaults

`func NewBillPaymentNewsContentWithDefaults() *BillPaymentNewsContent`

NewBillPaymentNewsContentWithDefaults instantiates a new BillPaymentNewsContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillId

`func (o *BillPaymentNewsContent) GetBillId() string`

GetBillId returns the BillId field if non-nil, zero value otherwise.

### GetBillIdOk

`func (o *BillPaymentNewsContent) GetBillIdOk() (*string, bool)`

GetBillIdOk returns a tuple with the BillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillId

`func (o *BillPaymentNewsContent) SetBillId(v string)`

SetBillId sets BillId field to given value.


### GetBillStatus

`func (o *BillPaymentNewsContent) GetBillStatus() BillStatus`

GetBillStatus returns the BillStatus field if non-nil, zero value otherwise.

### GetBillStatusOk

`func (o *BillPaymentNewsContent) GetBillStatusOk() (*BillStatus, bool)`

GetBillStatusOk returns a tuple with the BillStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillStatus

`func (o *BillPaymentNewsContent) SetBillStatus(v BillStatus)`

SetBillStatus sets BillStatus field to given value.


### GetRelatedPaymentsSum

`func (o *BillPaymentNewsContent) GetRelatedPaymentsSum() float64`

GetRelatedPaymentsSum returns the RelatedPaymentsSum field if non-nil, zero value otherwise.

### GetRelatedPaymentsSumOk

`func (o *BillPaymentNewsContent) GetRelatedPaymentsSumOk() (*float64, bool)`

GetRelatedPaymentsSumOk returns a tuple with the RelatedPaymentsSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedPaymentsSum

`func (o *BillPaymentNewsContent) SetRelatedPaymentsSum(v float64)`

SetRelatedPaymentsSum sets RelatedPaymentsSum field to given value.


### GetBillSum

`func (o *BillPaymentNewsContent) GetBillSum() float64`

GetBillSum returns the BillSum field if non-nil, zero value otherwise.

### GetBillSumOk

`func (o *BillPaymentNewsContent) GetBillSumOk() (*float64, bool)`

GetBillSumOk returns a tuple with the BillSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillSum

`func (o *BillPaymentNewsContent) SetBillSum(v float64)`

SetBillSum sets BillSum field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


