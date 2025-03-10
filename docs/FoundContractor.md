# FoundContractor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор созданного контрагента | 
**Name** | **string** | Полное название контрагента | 
**ShortName** | Pointer to **NullableString** | Краткое название контрагента | [optional] 
**Inn** | Pointer to **NullableString** | ИНН | [optional] 
**Kpp** | Pointer to **NullableString** | КПП | [optional] 
**Address** | Pointer to **NullableString** | Юридический адрес | [optional] 

## Methods

### NewFoundContractor

`func NewFoundContractor(id string, name string, ) *FoundContractor`

NewFoundContractor instantiates a new FoundContractor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFoundContractorWithDefaults

`func NewFoundContractorWithDefaults() *FoundContractor`

NewFoundContractorWithDefaults instantiates a new FoundContractor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FoundContractor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FoundContractor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FoundContractor) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *FoundContractor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FoundContractor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FoundContractor) SetName(v string)`

SetName sets Name field to given value.


### GetShortName

`func (o *FoundContractor) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *FoundContractor) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *FoundContractor) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *FoundContractor) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### SetShortNameNil

`func (o *FoundContractor) SetShortNameNil(b bool)`

 SetShortNameNil sets the value for ShortName to be an explicit nil

### UnsetShortName
`func (o *FoundContractor) UnsetShortName()`

UnsetShortName ensures that no value is present for ShortName, not even an explicit nil
### GetInn

`func (o *FoundContractor) GetInn() string`

GetInn returns the Inn field if non-nil, zero value otherwise.

### GetInnOk

`func (o *FoundContractor) GetInnOk() (*string, bool)`

GetInnOk returns a tuple with the Inn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInn

`func (o *FoundContractor) SetInn(v string)`

SetInn sets Inn field to given value.

### HasInn

`func (o *FoundContractor) HasInn() bool`

HasInn returns a boolean if a field has been set.

### SetInnNil

`func (o *FoundContractor) SetInnNil(b bool)`

 SetInnNil sets the value for Inn to be an explicit nil

### UnsetInn
`func (o *FoundContractor) UnsetInn()`

UnsetInn ensures that no value is present for Inn, not even an explicit nil
### GetKpp

`func (o *FoundContractor) GetKpp() string`

GetKpp returns the Kpp field if non-nil, zero value otherwise.

### GetKppOk

`func (o *FoundContractor) GetKppOk() (*string, bool)`

GetKppOk returns a tuple with the Kpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKpp

`func (o *FoundContractor) SetKpp(v string)`

SetKpp sets Kpp field to given value.

### HasKpp

`func (o *FoundContractor) HasKpp() bool`

HasKpp returns a boolean if a field has been set.

### SetKppNil

`func (o *FoundContractor) SetKppNil(b bool)`

 SetKppNil sets the value for Kpp to be an explicit nil

### UnsetKpp
`func (o *FoundContractor) UnsetKpp()`

UnsetKpp ensures that no value is present for Kpp, not even an explicit nil
### GetAddress

`func (o *FoundContractor) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FoundContractor) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FoundContractor) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *FoundContractor) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *FoundContractor) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *FoundContractor) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


