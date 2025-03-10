# ActItemToCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductName** | **string** | Наименование товара или услуги | 
**UnitName** | **string** | Единицы измерения | 
**Quantity** | **float64** | Количество | 
**Price** | Pointer to **NullableFloat64** | Цена за единицу товара | [optional] 
**NdsRate** | Pointer to [**NullableNDSRateToSave**](NDSRateToSave.md) | НДС  withoutNds (Без НДС)  nds0 (0%)  nds5 (5%)  nds10 (10%)  nds20 (20%) | [optional] 
**Discount** | Pointer to **NullableFloat64** | Скидка | [optional] 

## Methods

### NewActItemToCreate

`func NewActItemToCreate(productName string, unitName string, quantity float64, ) *ActItemToCreate`

NewActItemToCreate instantiates a new ActItemToCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActItemToCreateWithDefaults

`func NewActItemToCreateWithDefaults() *ActItemToCreate`

NewActItemToCreateWithDefaults instantiates a new ActItemToCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductName

`func (o *ActItemToCreate) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *ActItemToCreate) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *ActItemToCreate) SetProductName(v string)`

SetProductName sets ProductName field to given value.


### GetUnitName

`func (o *ActItemToCreate) GetUnitName() string`

GetUnitName returns the UnitName field if non-nil, zero value otherwise.

### GetUnitNameOk

`func (o *ActItemToCreate) GetUnitNameOk() (*string, bool)`

GetUnitNameOk returns a tuple with the UnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitName

`func (o *ActItemToCreate) SetUnitName(v string)`

SetUnitName sets UnitName field to given value.


### GetQuantity

`func (o *ActItemToCreate) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ActItemToCreate) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ActItemToCreate) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.


### GetPrice

`func (o *ActItemToCreate) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ActItemToCreate) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ActItemToCreate) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ActItemToCreate) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *ActItemToCreate) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *ActItemToCreate) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetNdsRate

`func (o *ActItemToCreate) GetNdsRate() NDSRateToSave`

GetNdsRate returns the NdsRate field if non-nil, zero value otherwise.

### GetNdsRateOk

`func (o *ActItemToCreate) GetNdsRateOk() (*NDSRateToSave, bool)`

GetNdsRateOk returns a tuple with the NdsRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNdsRate

`func (o *ActItemToCreate) SetNdsRate(v NDSRateToSave)`

SetNdsRate sets NdsRate field to given value.

### HasNdsRate

`func (o *ActItemToCreate) HasNdsRate() bool`

HasNdsRate returns a boolean if a field has been set.

### SetNdsRateNil

`func (o *ActItemToCreate) SetNdsRateNil(b bool)`

 SetNdsRateNil sets the value for NdsRate to be an explicit nil

### UnsetNdsRate
`func (o *ActItemToCreate) UnsetNdsRate()`

UnsetNdsRate ensures that no value is present for NdsRate, not even an explicit nil
### GetDiscount

`func (o *ActItemToCreate) GetDiscount() float64`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *ActItemToCreate) GetDiscountOk() (*float64, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *ActItemToCreate) SetDiscount(v float64)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *ActItemToCreate) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### SetDiscountNil

`func (o *ActItemToCreate) SetDiscountNil(b bool)`

 SetDiscountNil sets the value for Discount to be an explicit nil

### UnsetDiscount
`func (o *ActItemToCreate) UnsetDiscount()`

UnsetDiscount ensures that no value is present for Discount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


