# Contractor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор созданного контрагента | 
**Name** | **string** | Полное название контрагента | 
**ShortName** | Pointer to **NullableString** | Краткое название контрагента | [optional] 
**Inn** | Pointer to **NullableString** | ИНН | [optional] 
**Kpp** | Pointer to **NullableString** | КПП | [optional] 
**Address** | Pointer to **NullableString** | Юридический адрес | [optional] 
**PostAddress** | Pointer to **NullableString** | Фактический адрес | [optional] 
**Okpo** | Pointer to **NullableString** | ОКПО | [optional] 
**Ogrn** | Pointer to **NullableString** | ОГРН | [optional] 
**Contacts** | Pointer to [**[]ContractorContact**](ContractorContact.md) | Контакты | [optional] 

## Methods

### NewContractor

`func NewContractor(id string, name string, ) *Contractor`

NewContractor instantiates a new Contractor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractorWithDefaults

`func NewContractorWithDefaults() *Contractor`

NewContractorWithDefaults instantiates a new Contractor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Contractor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Contractor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Contractor) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Contractor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Contractor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Contractor) SetName(v string)`

SetName sets Name field to given value.


### GetShortName

`func (o *Contractor) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *Contractor) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *Contractor) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *Contractor) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### SetShortNameNil

`func (o *Contractor) SetShortNameNil(b bool)`

 SetShortNameNil sets the value for ShortName to be an explicit nil

### UnsetShortName
`func (o *Contractor) UnsetShortName()`

UnsetShortName ensures that no value is present for ShortName, not even an explicit nil
### GetInn

`func (o *Contractor) GetInn() string`

GetInn returns the Inn field if non-nil, zero value otherwise.

### GetInnOk

`func (o *Contractor) GetInnOk() (*string, bool)`

GetInnOk returns a tuple with the Inn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInn

`func (o *Contractor) SetInn(v string)`

SetInn sets Inn field to given value.

### HasInn

`func (o *Contractor) HasInn() bool`

HasInn returns a boolean if a field has been set.

### SetInnNil

`func (o *Contractor) SetInnNil(b bool)`

 SetInnNil sets the value for Inn to be an explicit nil

### UnsetInn
`func (o *Contractor) UnsetInn()`

UnsetInn ensures that no value is present for Inn, not even an explicit nil
### GetKpp

`func (o *Contractor) GetKpp() string`

GetKpp returns the Kpp field if non-nil, zero value otherwise.

### GetKppOk

`func (o *Contractor) GetKppOk() (*string, bool)`

GetKppOk returns a tuple with the Kpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKpp

`func (o *Contractor) SetKpp(v string)`

SetKpp sets Kpp field to given value.

### HasKpp

`func (o *Contractor) HasKpp() bool`

HasKpp returns a boolean if a field has been set.

### SetKppNil

`func (o *Contractor) SetKppNil(b bool)`

 SetKppNil sets the value for Kpp to be an explicit nil

### UnsetKpp
`func (o *Contractor) UnsetKpp()`

UnsetKpp ensures that no value is present for Kpp, not even an explicit nil
### GetAddress

`func (o *Contractor) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Contractor) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Contractor) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Contractor) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *Contractor) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *Contractor) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetPostAddress

`func (o *Contractor) GetPostAddress() string`

GetPostAddress returns the PostAddress field if non-nil, zero value otherwise.

### GetPostAddressOk

`func (o *Contractor) GetPostAddressOk() (*string, bool)`

GetPostAddressOk returns a tuple with the PostAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostAddress

`func (o *Contractor) SetPostAddress(v string)`

SetPostAddress sets PostAddress field to given value.

### HasPostAddress

`func (o *Contractor) HasPostAddress() bool`

HasPostAddress returns a boolean if a field has been set.

### SetPostAddressNil

`func (o *Contractor) SetPostAddressNil(b bool)`

 SetPostAddressNil sets the value for PostAddress to be an explicit nil

### UnsetPostAddress
`func (o *Contractor) UnsetPostAddress()`

UnsetPostAddress ensures that no value is present for PostAddress, not even an explicit nil
### GetOkpo

`func (o *Contractor) GetOkpo() string`

GetOkpo returns the Okpo field if non-nil, zero value otherwise.

### GetOkpoOk

`func (o *Contractor) GetOkpoOk() (*string, bool)`

GetOkpoOk returns a tuple with the Okpo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOkpo

`func (o *Contractor) SetOkpo(v string)`

SetOkpo sets Okpo field to given value.

### HasOkpo

`func (o *Contractor) HasOkpo() bool`

HasOkpo returns a boolean if a field has been set.

### SetOkpoNil

`func (o *Contractor) SetOkpoNil(b bool)`

 SetOkpoNil sets the value for Okpo to be an explicit nil

### UnsetOkpo
`func (o *Contractor) UnsetOkpo()`

UnsetOkpo ensures that no value is present for Okpo, not even an explicit nil
### GetOgrn

`func (o *Contractor) GetOgrn() string`

GetOgrn returns the Ogrn field if non-nil, zero value otherwise.

### GetOgrnOk

`func (o *Contractor) GetOgrnOk() (*string, bool)`

GetOgrnOk returns a tuple with the Ogrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgrn

`func (o *Contractor) SetOgrn(v string)`

SetOgrn sets Ogrn field to given value.

### HasOgrn

`func (o *Contractor) HasOgrn() bool`

HasOgrn returns a boolean if a field has been set.

### SetOgrnNil

`func (o *Contractor) SetOgrnNil(b bool)`

 SetOgrnNil sets the value for Ogrn to be an explicit nil

### UnsetOgrn
`func (o *Contractor) UnsetOgrn()`

UnsetOgrn ensures that no value is present for Ogrn, not even an explicit nil
### GetContacts

`func (o *Contractor) GetContacts() []ContractorContact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *Contractor) GetContactsOk() (*[]ContractorContact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *Contractor) SetContacts(v []ContractorContact)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *Contractor) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### SetContactsNil

`func (o *Contractor) SetContactsNil(b bool)`

 SetContactsNil sets the value for Contacts to be an explicit nil

### UnsetContacts
`func (o *Contractor) UnsetContacts()`

UnsetContacts ensures that no value is present for Contacts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


