# ProductUnitToUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Идентификатор единицы измерения | [optional] 
**IsMain** | **bool** | Единица измерения является основной | 
**Name** | **string** | Название единицы измерения | 
**Ratio** | **int32** | Соотношение единицы измерения с основной | 

## Methods

### NewProductUnitToUpdate

`func NewProductUnitToUpdate(isMain bool, name string, ratio int32, ) *ProductUnitToUpdate`

NewProductUnitToUpdate instantiates a new ProductUnitToUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductUnitToUpdateWithDefaults

`func NewProductUnitToUpdateWithDefaults() *ProductUnitToUpdate`

NewProductUnitToUpdateWithDefaults instantiates a new ProductUnitToUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductUnitToUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductUnitToUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductUnitToUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProductUnitToUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProductUnitToUpdate) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProductUnitToUpdate) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsMain

`func (o *ProductUnitToUpdate) GetIsMain() bool`

GetIsMain returns the IsMain field if non-nil, zero value otherwise.

### GetIsMainOk

`func (o *ProductUnitToUpdate) GetIsMainOk() (*bool, bool)`

GetIsMainOk returns a tuple with the IsMain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMain

`func (o *ProductUnitToUpdate) SetIsMain(v bool)`

SetIsMain sets IsMain field to given value.


### GetName

`func (o *ProductUnitToUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductUnitToUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductUnitToUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetRatio

`func (o *ProductUnitToUpdate) GetRatio() int32`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *ProductUnitToUpdate) GetRatioOk() (*int32, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *ProductUnitToUpdate) SetRatio(v int32)`

SetRatio sets Ratio field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


