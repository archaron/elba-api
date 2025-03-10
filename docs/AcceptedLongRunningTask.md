# AcceptedLongRunningTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор асинхронной операции | 
**Status** | [**LongRunningTaskStatus**](LongRunningTaskStatus.md) | Статус асинхронной операции  inProcess (Операция в обработке)  completed (Операция успешно обработана)  failed (Ошибка при выполнении операции) | 
**Type** | [**LongRunningTaskType**](LongRunningTaskType.md) | Тип задачи  searchContractors (Операция поиска контрагентов)  searchProducts (Операция поиска товаров)  createProduct (Операция создания товара)  updateProduct (Операция обновления товара) | 

## Methods

### NewAcceptedLongRunningTask

`func NewAcceptedLongRunningTask(id string, status LongRunningTaskStatus, type_ LongRunningTaskType, ) *AcceptedLongRunningTask`

NewAcceptedLongRunningTask instantiates a new AcceptedLongRunningTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceptedLongRunningTaskWithDefaults

`func NewAcceptedLongRunningTaskWithDefaults() *AcceptedLongRunningTask`

NewAcceptedLongRunningTaskWithDefaults instantiates a new AcceptedLongRunningTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AcceptedLongRunningTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AcceptedLongRunningTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AcceptedLongRunningTask) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *AcceptedLongRunningTask) GetStatus() LongRunningTaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AcceptedLongRunningTask) GetStatusOk() (*LongRunningTaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AcceptedLongRunningTask) SetStatus(v LongRunningTaskStatus)`

SetStatus sets Status field to given value.


### GetType

`func (o *AcceptedLongRunningTask) GetType() LongRunningTaskType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AcceptedLongRunningTask) GetTypeOk() (*LongRunningTaskType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AcceptedLongRunningTask) SetType(v LongRunningTaskType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


