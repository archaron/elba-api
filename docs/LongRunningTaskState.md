# LongRunningTaskState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор асинхронной операции | 
**Status** | [**LongRunningTaskStatus**](LongRunningTaskStatus.md) | Статус асинхронной операции  inProcess (Операция в обработке)  completed (Операция успешно обработана)  failed (Ошибка при выполнении операции) | 
**Type** | [**LongRunningTaskType**](LongRunningTaskType.md) | Тип операции  searchContractors (Операция поиска контрагентов)  searchProducts (Операция поиска товаров)  createProduct (Операция создания товара)  updateProduct (Операция обновления товара) | 
**Result** | Pointer to [**NullableLongRunningTaskStateResult**](LongRunningTaskStateResult.md) |  | [optional] 
**ErrorDescription** | Pointer to [**NullableLongRunningTaskExecutionError**](LongRunningTaskExecutionError.md) | Информация об ошибке, возникшей во время выполнения операции. Заполнена, если операция завершена с ошибкой. | [optional] 

## Methods

### NewLongRunningTaskState

`func NewLongRunningTaskState(id string, status LongRunningTaskStatus, type_ LongRunningTaskType, ) *LongRunningTaskState`

NewLongRunningTaskState instantiates a new LongRunningTaskState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLongRunningTaskStateWithDefaults

`func NewLongRunningTaskStateWithDefaults() *LongRunningTaskState`

NewLongRunningTaskStateWithDefaults instantiates a new LongRunningTaskState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LongRunningTaskState) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LongRunningTaskState) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LongRunningTaskState) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *LongRunningTaskState) GetStatus() LongRunningTaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LongRunningTaskState) GetStatusOk() (*LongRunningTaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LongRunningTaskState) SetStatus(v LongRunningTaskStatus)`

SetStatus sets Status field to given value.


### GetType

`func (o *LongRunningTaskState) GetType() LongRunningTaskType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LongRunningTaskState) GetTypeOk() (*LongRunningTaskType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LongRunningTaskState) SetType(v LongRunningTaskType)`

SetType sets Type field to given value.


### GetResult

`func (o *LongRunningTaskState) GetResult() LongRunningTaskStateResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *LongRunningTaskState) GetResultOk() (*LongRunningTaskStateResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *LongRunningTaskState) SetResult(v LongRunningTaskStateResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *LongRunningTaskState) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *LongRunningTaskState) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *LongRunningTaskState) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetErrorDescription

`func (o *LongRunningTaskState) GetErrorDescription() LongRunningTaskExecutionError`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *LongRunningTaskState) GetErrorDescriptionOk() (*LongRunningTaskExecutionError, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *LongRunningTaskState) SetErrorDescription(v LongRunningTaskExecutionError)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *LongRunningTaskState) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### SetErrorDescriptionNil

`func (o *LongRunningTaskState) SetErrorDescriptionNil(b bool)`

 SetErrorDescriptionNil sets the value for ErrorDescription to be an explicit nil

### UnsetErrorDescription
`func (o *LongRunningTaskState) UnsetErrorDescription()`

UnsetErrorDescription ensures that no value is present for ErrorDescription, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


