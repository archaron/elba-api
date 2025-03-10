# CreateProductTaskResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Идентификатор созданного товара. Заполнен, если нет ошибки создания товара | [optional] 
**CreateProductError** | Pointer to [**NullableCreateProductError**](CreateProductError.md) | Ошибка при создании товара  productWithSameNameExists (Уже есть товар с таким названием) | [optional] 

## Methods

### NewCreateProductTaskResult

`func NewCreateProductTaskResult() *CreateProductTaskResult`

NewCreateProductTaskResult instantiates a new CreateProductTaskResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProductTaskResultWithDefaults

`func NewCreateProductTaskResultWithDefaults() *CreateProductTaskResult`

NewCreateProductTaskResultWithDefaults instantiates a new CreateProductTaskResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateProductTaskResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateProductTaskResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateProductTaskResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateProductTaskResult) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CreateProductTaskResult) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CreateProductTaskResult) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCreateProductError

`func (o *CreateProductTaskResult) GetCreateProductError() CreateProductError`

GetCreateProductError returns the CreateProductError field if non-nil, zero value otherwise.

### GetCreateProductErrorOk

`func (o *CreateProductTaskResult) GetCreateProductErrorOk() (*CreateProductError, bool)`

GetCreateProductErrorOk returns a tuple with the CreateProductError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateProductError

`func (o *CreateProductTaskResult) SetCreateProductError(v CreateProductError)`

SetCreateProductError sets CreateProductError field to given value.

### HasCreateProductError

`func (o *CreateProductTaskResult) HasCreateProductError() bool`

HasCreateProductError returns a boolean if a field has been set.

### SetCreateProductErrorNil

`func (o *CreateProductTaskResult) SetCreateProductErrorNil(b bool)`

 SetCreateProductErrorNil sets the value for CreateProductError to be an explicit nil

### UnsetCreateProductError
`func (o *CreateProductTaskResult) UnsetCreateProductError()`

UnsetCreateProductError ensures that no value is present for CreateProductError, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


