# CreateContractorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Полное название контрагента | 
**Inn** | Pointer to **NullableString** | ИНН | [optional] 
**Kpp** | Pointer to **NullableString** | КПП | [optional] 
**ShortName** | Pointer to **NullableString** | Краткое название контрагента | [optional] 
**Address** | Pointer to **NullableString** | Юридический адрес | [optional] 
**PostAddress** | Pointer to **NullableString** | Фактический адрес | [optional] 
**Okpo** | Pointer to **NullableString** | ОКПО | [optional] 
**Ogrn** | Pointer to **NullableString** | ОГРН | [optional] 
**Contacts** | Pointer to [**[]ContractorContact**](ContractorContact.md) | Контакты | [optional] 

## Methods

### NewCreateContractorRequest

`func NewCreateContractorRequest(name string, ) *CreateContractorRequest`

NewCreateContractorRequest instantiates a new CreateContractorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContractorRequestWithDefaults

`func NewCreateContractorRequestWithDefaults() *CreateContractorRequest`

NewCreateContractorRequestWithDefaults instantiates a new CreateContractorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateContractorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateContractorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateContractorRequest) SetName(v string)`

SetName sets Name field to given value.


### GetInn

`func (o *CreateContractorRequest) GetInn() string`

GetInn returns the Inn field if non-nil, zero value otherwise.

### GetInnOk

`func (o *CreateContractorRequest) GetInnOk() (*string, bool)`

GetInnOk returns a tuple with the Inn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInn

`func (o *CreateContractorRequest) SetInn(v string)`

SetInn sets Inn field to given value.

### HasInn

`func (o *CreateContractorRequest) HasInn() bool`

HasInn returns a boolean if a field has been set.

### SetInnNil

`func (o *CreateContractorRequest) SetInnNil(b bool)`

 SetInnNil sets the value for Inn to be an explicit nil

### UnsetInn
`func (o *CreateContractorRequest) UnsetInn()`

UnsetInn ensures that no value is present for Inn, not even an explicit nil
### GetKpp

`func (o *CreateContractorRequest) GetKpp() string`

GetKpp returns the Kpp field if non-nil, zero value otherwise.

### GetKppOk

`func (o *CreateContractorRequest) GetKppOk() (*string, bool)`

GetKppOk returns a tuple with the Kpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKpp

`func (o *CreateContractorRequest) SetKpp(v string)`

SetKpp sets Kpp field to given value.

### HasKpp

`func (o *CreateContractorRequest) HasKpp() bool`

HasKpp returns a boolean if a field has been set.

### SetKppNil

`func (o *CreateContractorRequest) SetKppNil(b bool)`

 SetKppNil sets the value for Kpp to be an explicit nil

### UnsetKpp
`func (o *CreateContractorRequest) UnsetKpp()`

UnsetKpp ensures that no value is present for Kpp, not even an explicit nil
### GetShortName

`func (o *CreateContractorRequest) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *CreateContractorRequest) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *CreateContractorRequest) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *CreateContractorRequest) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### SetShortNameNil

`func (o *CreateContractorRequest) SetShortNameNil(b bool)`

 SetShortNameNil sets the value for ShortName to be an explicit nil

### UnsetShortName
`func (o *CreateContractorRequest) UnsetShortName()`

UnsetShortName ensures that no value is present for ShortName, not even an explicit nil
### GetAddress

`func (o *CreateContractorRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateContractorRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateContractorRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CreateContractorRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *CreateContractorRequest) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *CreateContractorRequest) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetPostAddress

`func (o *CreateContractorRequest) GetPostAddress() string`

GetPostAddress returns the PostAddress field if non-nil, zero value otherwise.

### GetPostAddressOk

`func (o *CreateContractorRequest) GetPostAddressOk() (*string, bool)`

GetPostAddressOk returns a tuple with the PostAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostAddress

`func (o *CreateContractorRequest) SetPostAddress(v string)`

SetPostAddress sets PostAddress field to given value.

### HasPostAddress

`func (o *CreateContractorRequest) HasPostAddress() bool`

HasPostAddress returns a boolean if a field has been set.

### SetPostAddressNil

`func (o *CreateContractorRequest) SetPostAddressNil(b bool)`

 SetPostAddressNil sets the value for PostAddress to be an explicit nil

### UnsetPostAddress
`func (o *CreateContractorRequest) UnsetPostAddress()`

UnsetPostAddress ensures that no value is present for PostAddress, not even an explicit nil
### GetOkpo

`func (o *CreateContractorRequest) GetOkpo() string`

GetOkpo returns the Okpo field if non-nil, zero value otherwise.

### GetOkpoOk

`func (o *CreateContractorRequest) GetOkpoOk() (*string, bool)`

GetOkpoOk returns a tuple with the Okpo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOkpo

`func (o *CreateContractorRequest) SetOkpo(v string)`

SetOkpo sets Okpo field to given value.

### HasOkpo

`func (o *CreateContractorRequest) HasOkpo() bool`

HasOkpo returns a boolean if a field has been set.

### SetOkpoNil

`func (o *CreateContractorRequest) SetOkpoNil(b bool)`

 SetOkpoNil sets the value for Okpo to be an explicit nil

### UnsetOkpo
`func (o *CreateContractorRequest) UnsetOkpo()`

UnsetOkpo ensures that no value is present for Okpo, not even an explicit nil
### GetOgrn

`func (o *CreateContractorRequest) GetOgrn() string`

GetOgrn returns the Ogrn field if non-nil, zero value otherwise.

### GetOgrnOk

`func (o *CreateContractorRequest) GetOgrnOk() (*string, bool)`

GetOgrnOk returns a tuple with the Ogrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOgrn

`func (o *CreateContractorRequest) SetOgrn(v string)`

SetOgrn sets Ogrn field to given value.

### HasOgrn

`func (o *CreateContractorRequest) HasOgrn() bool`

HasOgrn returns a boolean if a field has been set.

### SetOgrnNil

`func (o *CreateContractorRequest) SetOgrnNil(b bool)`

 SetOgrnNil sets the value for Ogrn to be an explicit nil

### UnsetOgrn
`func (o *CreateContractorRequest) UnsetOgrn()`

UnsetOgrn ensures that no value is present for Ogrn, not even an explicit nil
### GetContacts

`func (o *CreateContractorRequest) GetContacts() []ContractorContact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *CreateContractorRequest) GetContactsOk() (*[]ContractorContact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *CreateContractorRequest) SetContacts(v []ContractorContact)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *CreateContractorRequest) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### SetContactsNil

`func (o *CreateContractorRequest) SetContactsNil(b bool)`

 SetContactsNil sets the value for Contacts to be an explicit nil

### UnsetContacts
`func (o *CreateContractorRequest) UnsetContacts()`

UnsetContacts ensures that no value is present for Contacts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


