# UpdateProductTaskResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateProductError** | Pointer to [**NullableUpdateProductError**](UpdateProductError.md) | Ошибка при обновлении товара  productWithSameNameExists (Уже есть товар с таким названием)  deletingProductNameThatUsedInDocuments (Название товара нельзя удалить, так как оно используется в документах)  renamingProductNameThatUsedInDocuments (Название товара нельзя изменить, так как оно используется в документах)  deletingProductUnitThatUsedInDocuments (Единицу измерения нельзя удалить, так как она используется в документах)  renamingProductUnitThatUsedInDocuments (Единицу измерения нельзя переименовать, так как она используется в документах) | [optional] 

## Methods

### NewUpdateProductTaskResult

`func NewUpdateProductTaskResult() *UpdateProductTaskResult`

NewUpdateProductTaskResult instantiates a new UpdateProductTaskResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProductTaskResultWithDefaults

`func NewUpdateProductTaskResultWithDefaults() *UpdateProductTaskResult`

NewUpdateProductTaskResultWithDefaults instantiates a new UpdateProductTaskResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateProductError

`func (o *UpdateProductTaskResult) GetUpdateProductError() UpdateProductError`

GetUpdateProductError returns the UpdateProductError field if non-nil, zero value otherwise.

### GetUpdateProductErrorOk

`func (o *UpdateProductTaskResult) GetUpdateProductErrorOk() (*UpdateProductError, bool)`

GetUpdateProductErrorOk returns a tuple with the UpdateProductError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateProductError

`func (o *UpdateProductTaskResult) SetUpdateProductError(v UpdateProductError)`

SetUpdateProductError sets UpdateProductError field to given value.

### HasUpdateProductError

`func (o *UpdateProductTaskResult) HasUpdateProductError() bool`

HasUpdateProductError returns a boolean if a field has been set.

### SetUpdateProductErrorNil

`func (o *UpdateProductTaskResult) SetUpdateProductErrorNil(b bool)`

 SetUpdateProductErrorNil sets the value for UpdateProductError to be an explicit nil

### UnsetUpdateProductError
`func (o *UpdateProductTaskResult) UnsetUpdateProductError()`

UnsetUpdateProductError ensures that no value is present for UpdateProductError, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


