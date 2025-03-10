# ProductNameToUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Идентификатор названия товара | [optional] 
**Name** | **string** | Название товара | 
**IsMain** | **bool** | Название товара является основным | 

## Methods

### NewProductNameToUpdate

`func NewProductNameToUpdate(name string, isMain bool, ) *ProductNameToUpdate`

NewProductNameToUpdate instantiates a new ProductNameToUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductNameToUpdateWithDefaults

`func NewProductNameToUpdateWithDefaults() *ProductNameToUpdate`

NewProductNameToUpdateWithDefaults instantiates a new ProductNameToUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductNameToUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductNameToUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductNameToUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProductNameToUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProductNameToUpdate) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProductNameToUpdate) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ProductNameToUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductNameToUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductNameToUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetIsMain

`func (o *ProductNameToUpdate) GetIsMain() bool`

GetIsMain returns the IsMain field if non-nil, zero value otherwise.

### GetIsMainOk

`func (o *ProductNameToUpdate) GetIsMainOk() (*bool, bool)`

GetIsMainOk returns a tuple with the IsMain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMain

`func (o *ProductNameToUpdate) SetIsMain(v bool)`

SetIsMain sets IsMain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


