# Bank

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Название банка | [optional] 
**Bik** | Pointer to **NullableString** | БИК | [optional] 
**CorrAccount** | Pointer to **NullableString** | Корреспондентский счёт | [optional] 
**City** | Pointer to **NullableString** | Населенный пункт | [optional] 
**CityType** | Pointer to **NullableString** | Тип населенного пункта | [optional] 

## Methods

### NewBank

`func NewBank() *Bank`

NewBank instantiates a new Bank object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankWithDefaults

`func NewBankWithDefaults() *Bank`

NewBankWithDefaults instantiates a new Bank object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Bank) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Bank) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Bank) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Bank) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Bank) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Bank) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetBik

`func (o *Bank) GetBik() string`

GetBik returns the Bik field if non-nil, zero value otherwise.

### GetBikOk

`func (o *Bank) GetBikOk() (*string, bool)`

GetBikOk returns a tuple with the Bik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBik

`func (o *Bank) SetBik(v string)`

SetBik sets Bik field to given value.

### HasBik

`func (o *Bank) HasBik() bool`

HasBik returns a boolean if a field has been set.

### SetBikNil

`func (o *Bank) SetBikNil(b bool)`

 SetBikNil sets the value for Bik to be an explicit nil

### UnsetBik
`func (o *Bank) UnsetBik()`

UnsetBik ensures that no value is present for Bik, not even an explicit nil
### GetCorrAccount

`func (o *Bank) GetCorrAccount() string`

GetCorrAccount returns the CorrAccount field if non-nil, zero value otherwise.

### GetCorrAccountOk

`func (o *Bank) GetCorrAccountOk() (*string, bool)`

GetCorrAccountOk returns a tuple with the CorrAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrAccount

`func (o *Bank) SetCorrAccount(v string)`

SetCorrAccount sets CorrAccount field to given value.

### HasCorrAccount

`func (o *Bank) HasCorrAccount() bool`

HasCorrAccount returns a boolean if a field has been set.

### SetCorrAccountNil

`func (o *Bank) SetCorrAccountNil(b bool)`

 SetCorrAccountNil sets the value for CorrAccount to be an explicit nil

### UnsetCorrAccount
`func (o *Bank) UnsetCorrAccount()`

UnsetCorrAccount ensures that no value is present for CorrAccount, not even an explicit nil
### GetCity

`func (o *Bank) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Bank) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Bank) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Bank) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *Bank) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *Bank) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCityType

`func (o *Bank) GetCityType() string`

GetCityType returns the CityType field if non-nil, zero value otherwise.

### GetCityTypeOk

`func (o *Bank) GetCityTypeOk() (*string, bool)`

GetCityTypeOk returns a tuple with the CityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityType

`func (o *Bank) SetCityType(v string)`

SetCityType sets CityType field to given value.

### HasCityType

`func (o *Bank) HasCityType() bool`

HasCityType returns a boolean if a field has been set.

### SetCityTypeNil

`func (o *Bank) SetCityTypeNil(b bool)`

 SetCityTypeNil sets the value for CityType to be an explicit nil

### UnsetCityType
`func (o *Bank) UnsetCityType()`

UnsetCityType ensures that no value is present for CityType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


