# ProductUnitToCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsMain** | **bool** | Единица измерения является основной | 
**Name** | **string** | Название единицы измерения | 
**Ratio** | **float64** | Соотношение единицы измерения с основной | 

## Methods

### NewProductUnitToCreate

`func NewProductUnitToCreate(isMain bool, name string, ratio float64, ) *ProductUnitToCreate`

NewProductUnitToCreate instantiates a new ProductUnitToCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductUnitToCreateWithDefaults

`func NewProductUnitToCreateWithDefaults() *ProductUnitToCreate`

NewProductUnitToCreateWithDefaults instantiates a new ProductUnitToCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsMain

`func (o *ProductUnitToCreate) GetIsMain() bool`

GetIsMain returns the IsMain field if non-nil, zero value otherwise.

### GetIsMainOk

`func (o *ProductUnitToCreate) GetIsMainOk() (*bool, bool)`

GetIsMainOk returns a tuple with the IsMain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMain

`func (o *ProductUnitToCreate) SetIsMain(v bool)`

SetIsMain sets IsMain field to given value.


### GetName

`func (o *ProductUnitToCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductUnitToCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductUnitToCreate) SetName(v string)`

SetName sets Name field to given value.


### GetRatio

`func (o *ProductUnitToCreate) GetRatio() float64`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *ProductUnitToCreate) GetRatioOk() (*float64, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *ProductUnitToCreate) SetRatio(v float64)`

SetRatio sets Ratio field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


