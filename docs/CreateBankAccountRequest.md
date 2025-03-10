# CreateBankAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | Номер банковского счёта | 
**Bik** | **string** | БИК | 

## Methods

### NewCreateBankAccountRequest

`func NewCreateBankAccountRequest(accountNumber string, bik string, ) *CreateBankAccountRequest`

NewCreateBankAccountRequest instantiates a new CreateBankAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBankAccountRequestWithDefaults

`func NewCreateBankAccountRequestWithDefaults() *CreateBankAccountRequest`

NewCreateBankAccountRequestWithDefaults instantiates a new CreateBankAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *CreateBankAccountRequest) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *CreateBankAccountRequest) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *CreateBankAccountRequest) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetBik

`func (o *CreateBankAccountRequest) GetBik() string`

GetBik returns the Bik field if non-nil, zero value otherwise.

### GetBikOk

`func (o *CreateBankAccountRequest) GetBikOk() (*string, bool)`

GetBikOk returns a tuple with the Bik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBik

`func (o *CreateBankAccountRequest) SetBik(v string)`

SetBik sets Bik field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


