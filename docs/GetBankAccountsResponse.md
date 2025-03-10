# GetBankAccountsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankAccounts** | [**[]BankAccount**](BankAccount.md) | Список банковских счетов | 

## Methods

### NewGetBankAccountsResponse

`func NewGetBankAccountsResponse(bankAccounts []BankAccount, ) *GetBankAccountsResponse`

NewGetBankAccountsResponse instantiates a new GetBankAccountsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBankAccountsResponseWithDefaults

`func NewGetBankAccountsResponseWithDefaults() *GetBankAccountsResponse`

NewGetBankAccountsResponseWithDefaults instantiates a new GetBankAccountsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankAccounts

`func (o *GetBankAccountsResponse) GetBankAccounts() []BankAccount`

GetBankAccounts returns the BankAccounts field if non-nil, zero value otherwise.

### GetBankAccountsOk

`func (o *GetBankAccountsResponse) GetBankAccountsOk() (*[]BankAccount, bool)`

GetBankAccountsOk returns a tuple with the BankAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccounts

`func (o *GetBankAccountsResponse) SetBankAccounts(v []BankAccount)`

SetBankAccounts sets BankAccounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


