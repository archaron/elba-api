# UniversalTransferDocumentItemToCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductName** | **string** | Наименование товара или услуги | 
**ProductType** | [**ProductType**](ProductType.md) | Тип товара  product (Товар)  service (Услуга) | 
**UnitName** | **string** | Единица измерения | 
**Quantity** | **float64** | Количество | 
**Price** | Pointer to **NullableFloat64** | Цена за единицу товара | [optional] 
**NdsRate** | Pointer to [**NullableNDSRateToSave**](NDSRateToSave.md) | НДС  withoutNds (Без НДС)  nds0 (0%)  nds5 (5%)  nds10 (10%)  nds20 (20%) | [optional] 
**Rnpt** | Pointer to **NullableString** | РНПТ | [optional] 
**OriginOfProductCountry** | Pointer to **NullableString** | Страна происхождения товара | [optional] 
**Discount** | Pointer to **NullableFloat64** | Скидка | [optional] 

## Methods

### NewUniversalTransferDocumentItemToCreate

`func NewUniversalTransferDocumentItemToCreate(productName string, productType ProductType, unitName string, quantity float64, ) *UniversalTransferDocumentItemToCreate`

NewUniversalTransferDocumentItemToCreate instantiates a new UniversalTransferDocumentItemToCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniversalTransferDocumentItemToCreateWithDefaults

`func NewUniversalTransferDocumentItemToCreateWithDefaults() *UniversalTransferDocumentItemToCreate`

NewUniversalTransferDocumentItemToCreateWithDefaults instantiates a new UniversalTransferDocumentItemToCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductName

`func (o *UniversalTransferDocumentItemToCreate) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *UniversalTransferDocumentItemToCreate) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *UniversalTransferDocumentItemToCreate) SetProductName(v string)`

SetProductName sets ProductName field to given value.


### GetProductType

`func (o *UniversalTransferDocumentItemToCreate) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *UniversalTransferDocumentItemToCreate) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *UniversalTransferDocumentItemToCreate) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.


### GetUnitName

`func (o *UniversalTransferDocumentItemToCreate) GetUnitName() string`

GetUnitName returns the UnitName field if non-nil, zero value otherwise.

### GetUnitNameOk

`func (o *UniversalTransferDocumentItemToCreate) GetUnitNameOk() (*string, bool)`

GetUnitNameOk returns a tuple with the UnitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitName

`func (o *UniversalTransferDocumentItemToCreate) SetUnitName(v string)`

SetUnitName sets UnitName field to given value.


### GetQuantity

`func (o *UniversalTransferDocumentItemToCreate) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UniversalTransferDocumentItemToCreate) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UniversalTransferDocumentItemToCreate) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.


### GetPrice

`func (o *UniversalTransferDocumentItemToCreate) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UniversalTransferDocumentItemToCreate) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UniversalTransferDocumentItemToCreate) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *UniversalTransferDocumentItemToCreate) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *UniversalTransferDocumentItemToCreate) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *UniversalTransferDocumentItemToCreate) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetNdsRate

`func (o *UniversalTransferDocumentItemToCreate) GetNdsRate() NDSRateToSave`

GetNdsRate returns the NdsRate field if non-nil, zero value otherwise.

### GetNdsRateOk

`func (o *UniversalTransferDocumentItemToCreate) GetNdsRateOk() (*NDSRateToSave, bool)`

GetNdsRateOk returns a tuple with the NdsRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNdsRate

`func (o *UniversalTransferDocumentItemToCreate) SetNdsRate(v NDSRateToSave)`

SetNdsRate sets NdsRate field to given value.

### HasNdsRate

`func (o *UniversalTransferDocumentItemToCreate) HasNdsRate() bool`

HasNdsRate returns a boolean if a field has been set.

### SetNdsRateNil

`func (o *UniversalTransferDocumentItemToCreate) SetNdsRateNil(b bool)`

 SetNdsRateNil sets the value for NdsRate to be an explicit nil

### UnsetNdsRate
`func (o *UniversalTransferDocumentItemToCreate) UnsetNdsRate()`

UnsetNdsRate ensures that no value is present for NdsRate, not even an explicit nil
### GetRnpt

`func (o *UniversalTransferDocumentItemToCreate) GetRnpt() string`

GetRnpt returns the Rnpt field if non-nil, zero value otherwise.

### GetRnptOk

`func (o *UniversalTransferDocumentItemToCreate) GetRnptOk() (*string, bool)`

GetRnptOk returns a tuple with the Rnpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRnpt

`func (o *UniversalTransferDocumentItemToCreate) SetRnpt(v string)`

SetRnpt sets Rnpt field to given value.

### HasRnpt

`func (o *UniversalTransferDocumentItemToCreate) HasRnpt() bool`

HasRnpt returns a boolean if a field has been set.

### SetRnptNil

`func (o *UniversalTransferDocumentItemToCreate) SetRnptNil(b bool)`

 SetRnptNil sets the value for Rnpt to be an explicit nil

### UnsetRnpt
`func (o *UniversalTransferDocumentItemToCreate) UnsetRnpt()`

UnsetRnpt ensures that no value is present for Rnpt, not even an explicit nil
### GetOriginOfProductCountry

`func (o *UniversalTransferDocumentItemToCreate) GetOriginOfProductCountry() string`

GetOriginOfProductCountry returns the OriginOfProductCountry field if non-nil, zero value otherwise.

### GetOriginOfProductCountryOk

`func (o *UniversalTransferDocumentItemToCreate) GetOriginOfProductCountryOk() (*string, bool)`

GetOriginOfProductCountryOk returns a tuple with the OriginOfProductCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginOfProductCountry

`func (o *UniversalTransferDocumentItemToCreate) SetOriginOfProductCountry(v string)`

SetOriginOfProductCountry sets OriginOfProductCountry field to given value.

### HasOriginOfProductCountry

`func (o *UniversalTransferDocumentItemToCreate) HasOriginOfProductCountry() bool`

HasOriginOfProductCountry returns a boolean if a field has been set.

### SetOriginOfProductCountryNil

`func (o *UniversalTransferDocumentItemToCreate) SetOriginOfProductCountryNil(b bool)`

 SetOriginOfProductCountryNil sets the value for OriginOfProductCountry to be an explicit nil

### UnsetOriginOfProductCountry
`func (o *UniversalTransferDocumentItemToCreate) UnsetOriginOfProductCountry()`

UnsetOriginOfProductCountry ensures that no value is present for OriginOfProductCountry, not even an explicit nil
### GetDiscount

`func (o *UniversalTransferDocumentItemToCreate) GetDiscount() float64`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *UniversalTransferDocumentItemToCreate) GetDiscountOk() (*float64, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *UniversalTransferDocumentItemToCreate) SetDiscount(v float64)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *UniversalTransferDocumentItemToCreate) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### SetDiscountNil

`func (o *UniversalTransferDocumentItemToCreate) SetDiscountNil(b bool)`

 SetDiscountNil sets the value for Discount to be an explicit nil

### UnsetDiscount
`func (o *UniversalTransferDocumentItemToCreate) UnsetDiscount()`

UnsetDiscount ensures that no value is present for Discount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


