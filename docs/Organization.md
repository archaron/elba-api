# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор организации | 
**Inn** | Pointer to **NullableString** | ИНН | [optional] 
**Kpp** | Pointer to **NullableString** | КПП | [optional] 

## Methods

### NewOrganization

`func NewOrganization(id string, ) *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Organization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Organization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Organization) SetId(v string)`

SetId sets Id field to given value.


### GetInn

`func (o *Organization) GetInn() string`

GetInn returns the Inn field if non-nil, zero value otherwise.

### GetInnOk

`func (o *Organization) GetInnOk() (*string, bool)`

GetInnOk returns a tuple with the Inn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInn

`func (o *Organization) SetInn(v string)`

SetInn sets Inn field to given value.

### HasInn

`func (o *Organization) HasInn() bool`

HasInn returns a boolean if a field has been set.

### SetInnNil

`func (o *Organization) SetInnNil(b bool)`

 SetInnNil sets the value for Inn to be an explicit nil

### UnsetInn
`func (o *Organization) UnsetInn()`

UnsetInn ensures that no value is present for Inn, not even an explicit nil
### GetKpp

`func (o *Organization) GetKpp() string`

GetKpp returns the Kpp field if non-nil, zero value otherwise.

### GetKppOk

`func (o *Organization) GetKppOk() (*string, bool)`

GetKppOk returns a tuple with the Kpp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKpp

`func (o *Organization) SetKpp(v string)`

SetKpp sets Kpp field to given value.

### HasKpp

`func (o *Organization) HasKpp() bool`

HasKpp returns a boolean if a field has been set.

### SetKppNil

`func (o *Organization) SetKppNil(b bool)`

 SetKppNil sets the value for Kpp to be an explicit nil

### UnsetKpp
`func (o *Organization) UnsetKpp()`

UnsetKpp ensures that no value is present for Kpp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


