# BillItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор позиции | 
**ProductName** | Pointer to **NullableString** | Наименование товара или услуги | [optional] 
**UnitName** | Pointer to **NullableString** | Единицы измерения | [optional] 
**Quantity** | **float64** | Количество | 
**Price** | Pointer to **NullableFloat64** | Цена за единицу товара | [optional] 
**Sum** | Pointer to **NullableFloat64** | Сумма | [optional] 
**NdsRate** | [**NDSRate**](NDSRate.md) | НДС  withoutNds (Без НДС)  nds0 (0%)  nds5 (5%)  nds10 (10%)  nds18 (18%)  nds20 (20%) | 
**Discount** | Pointer to **NullableFloat64** | Скидка | [optional] 

## Methods

### NewBillItem

`func NewBillItem(id string, quantity float64, ndsRate NDSRate, ) *BillItem`

NewBillItem instantiates a new BillItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillItemWithDefaults

`func NewBillItemWithDefaults() *BillItem`

NewBillItemWithDefaults instantiates a new BillItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillItem) SetId(v string)`

SetId sets Id field to given value.


### GetProductName

`func (o *BillItem) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *BillItem) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *BillItem) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *BillItem) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *BillItem) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *BillItem) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetUnitName

`func (o *BillItem) GetUnitName() string`

GetUnitName returns the UnitName field if non-nil, zero value otherwise.

### GetUnitNameOk

`func (o *BillItem) GetUnitNameOk() (*string, bool)`

GetUnitNameOk returns a tuple with the UnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitName

`func (o *BillItem) SetUnitName(v string)`

SetUnitName sets UnitName field to given value.

### HasUnitName

`func (o *BillItem) HasUnitName() bool`

HasUnitName returns a boolean if a field has been set.

### SetUnitNameNil

`func (o *BillItem) SetUnitNameNil(b bool)`

 SetUnitNameNil sets the value for UnitName to be an explicit nil

### UnsetUnitName
`func (o *BillItem) UnsetUnitName()`

UnsetUnitName ensures that no value is present for UnitName, not even an explicit nil
### GetQuantity

`func (o *BillItem) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BillItem) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BillItem) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.


### GetPrice

`func (o *BillItem) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillItem) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillItem) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *BillItem) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *BillItem) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetSum

`func (o *BillItem) GetSum() float64`

GetSum returns the Sum field if non-nil, zero value otherwise.

### GetSumOk

`func (o *BillItem) GetSumOk() (*float64, bool)`

GetSumOk returns a tuple with the Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSum

`func (o *BillItem) SetSum(v float64)`

SetSum sets Sum field to given value.

### HasSum

`func (o *BillItem) HasSum() bool`

HasSum returns a boolean if a field has been set.

### SetSumNil

`func (o *BillItem) SetSumNil(b bool)`

 SetSumNil sets the value for Sum to be an explicit nil

### UnsetSum
`func (o *BillItem) UnsetSum()`

UnsetSum ensures that no value is present for Sum, not even an explicit nil
### GetNdsRate

`func (o *BillItem) GetNdsRate() NDSRate`

GetNdsRate returns the NdsRate field if non-nil, zero value otherwise.

### GetNdsRateOk

`func (o *BillItem) GetNdsRateOk() (*NDSRate, bool)`

GetNdsRateOk returns a tuple with the NdsRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNdsRate

`func (o *BillItem) SetNdsRate(v NDSRate)`

SetNdsRate sets NdsRate field to given value.


### GetDiscount

`func (o *BillItem) GetDiscount() float64`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *BillItem) GetDiscountOk() (*float64, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *BillItem) SetDiscount(v float64)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *BillItem) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### SetDiscountNil

`func (o *BillItem) SetDiscountNil(b bool)`

 SetDiscountNil sets the value for Discount to be an explicit nil

### UnsetDiscount
`func (o *BillItem) UnsetDiscount()`

UnsetDiscount ensures that no value is present for Discount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


