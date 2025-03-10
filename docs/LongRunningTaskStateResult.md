# LongRunningTaskStateResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateProductError** | Pointer to [**NullableUpdateProductError**](UpdateProductError.md) | Ошибка при обновлении товара  productWithSameNameExists (Уже есть товар с таким названием)  deletingProductNameThatUsedInDocuments (Название товара нельзя удалить, так как оно используется в документах)  renamingProductNameThatUsedInDocuments (Название товара нельзя изменить, так как оно используется в документах)  deletingProductUnitThatUsedInDocuments (Единицу измерения нельзя удалить, так как она используется в документах)  renamingProductUnitThatUsedInDocuments (Единицу измерения нельзя переименовать, так как она используется в документах) | [optional] 
**Products** | [**[]FoundProduct**](FoundProduct.md) | Список найденных товаров | 
**Id** | Pointer to **NullableString** | Идентификатор созданного товара. Заполнен, если нет ошибки создания товара | [optional] 
**CreateProductError** | Pointer to [**NullableCreateProductError**](CreateProductError.md) | Ошибка при создании товара  productWithSameNameExists (Уже есть товар с таким названием) | [optional] 
**Contractors** | [**[]FoundContractor**](FoundContractor.md) | Список найденных контрагентов | 

## Methods

### NewLongRunningTaskStateResult

`func NewLongRunningTaskStateResult(products []FoundProduct, contractors []FoundContractor, ) *LongRunningTaskStateResult`

NewLongRunningTaskStateResult instantiates a new LongRunningTaskStateResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLongRunningTaskStateResultWithDefaults

`func NewLongRunningTaskStateResultWithDefaults() *LongRunningTaskStateResult`

NewLongRunningTaskStateResultWithDefaults instantiates a new LongRunningTaskStateResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateProductError

`func (o *LongRunningTaskStateResult) GetUpdateProductError() UpdateProductError`

GetUpdateProductError returns the UpdateProductError field if non-nil, zero value otherwise.

### GetUpdateProductErrorOk

`func (o *LongRunningTaskStateResult) GetUpdateProductErrorOk() (*UpdateProductError, bool)`

GetUpdateProductErrorOk returns a tuple with the UpdateProductError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateProductError

`func (o *LongRunningTaskStateResult) SetUpdateProductError(v UpdateProductError)`

SetUpdateProductError sets UpdateProductError field to given value.

### HasUpdateProductError

`func (o *LongRunningTaskStateResult) HasUpdateProductError() bool`

HasUpdateProductError returns a boolean if a field has been set.

### SetUpdateProductErrorNil

`func (o *LongRunningTaskStateResult) SetUpdateProductErrorNil(b bool)`

 SetUpdateProductErrorNil sets the value for UpdateProductError to be an explicit nil

### UnsetUpdateProductError
`func (o *LongRunningTaskStateResult) UnsetUpdateProductError()`

UnsetUpdateProductError ensures that no value is present for UpdateProductError, not even an explicit nil
### GetProducts

`func (o *LongRunningTaskStateResult) GetProducts() []FoundProduct`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *LongRunningTaskStateResult) GetProductsOk() (*[]FoundProduct, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *LongRunningTaskStateResult) SetProducts(v []FoundProduct)`

SetProducts sets Products field to given value.


### GetId

`func (o *LongRunningTaskStateResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LongRunningTaskStateResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LongRunningTaskStateResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LongRunningTaskStateResult) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *LongRunningTaskStateResult) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *LongRunningTaskStateResult) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCreateProductError

`func (o *LongRunningTaskStateResult) GetCreateProductError() CreateProductError`

GetCreateProductError returns the CreateProductError field if non-nil, zero value otherwise.

### GetCreateProductErrorOk

`func (o *LongRunningTaskStateResult) GetCreateProductErrorOk() (*CreateProductError, bool)`

GetCreateProductErrorOk returns a tuple with the CreateProductError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateProductError

`func (o *LongRunningTaskStateResult) SetCreateProductError(v CreateProductError)`

SetCreateProductError sets CreateProductError field to given value.

### HasCreateProductError

`func (o *LongRunningTaskStateResult) HasCreateProductError() bool`

HasCreateProductError returns a boolean if a field has been set.

### SetCreateProductErrorNil

`func (o *LongRunningTaskStateResult) SetCreateProductErrorNil(b bool)`

 SetCreateProductErrorNil sets the value for CreateProductError to be an explicit nil

### UnsetCreateProductError
`func (o *LongRunningTaskStateResult) UnsetCreateProductError()`

UnsetCreateProductError ensures that no value is present for CreateProductError, not even an explicit nil
### GetContractors

`func (o *LongRunningTaskStateResult) GetContractors() []FoundContractor`

GetContractors returns the Contractors field if non-nil, zero value otherwise.

### GetContractorsOk

`func (o *LongRunningTaskStateResult) GetContractorsOk() (*[]FoundContractor, bool)`

GetContractorsOk returns a tuple with the Contractors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractors

`func (o *LongRunningTaskStateResult) SetContractors(v []FoundContractor)`

SetContractors sets Contractors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


