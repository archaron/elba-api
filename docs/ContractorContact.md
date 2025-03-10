# ContractorContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Имя | [optional] 
**Post** | Pointer to **NullableString** | Должность | [optional] 
**Phone** | Pointer to **NullableString** | Телефон | [optional] 
**Emails** | **[]string** | Почта | 
**Comment** | Pointer to **NullableString** | Комментарий | [optional] 

## Methods

### NewContractorContact

`func NewContractorContact(emails []string, ) *ContractorContact`

NewContractorContact instantiates a new ContractorContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractorContactWithDefaults

`func NewContractorContactWithDefaults() *ContractorContact`

NewContractorContactWithDefaults instantiates a new ContractorContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContractorContact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContractorContact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContractorContact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContractorContact) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ContractorContact) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ContractorContact) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPost

`func (o *ContractorContact) GetPost() string`

GetPost returns the Post field if non-nil, zero value otherwise.

### GetPostOk

`func (o *ContractorContact) GetPostOk() (*string, bool)`

GetPostOk returns a tuple with the Post field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPost

`func (o *ContractorContact) SetPost(v string)`

SetPost sets Post field to given value.

### HasPost

`func (o *ContractorContact) HasPost() bool`

HasPost returns a boolean if a field has been set.

### SetPostNil

`func (o *ContractorContact) SetPostNil(b bool)`

 SetPostNil sets the value for Post to be an explicit nil

### UnsetPost
`func (o *ContractorContact) UnsetPost()`

UnsetPost ensures that no value is present for Post, not even an explicit nil
### GetPhone

`func (o *ContractorContact) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ContractorContact) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ContractorContact) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ContractorContact) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *ContractorContact) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *ContractorContact) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetEmails

`func (o *ContractorContact) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *ContractorContact) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *ContractorContact) SetEmails(v []string)`

SetEmails sets Emails field to given value.


### SetEmailsNil

`func (o *ContractorContact) SetEmailsNil(b bool)`

 SetEmailsNil sets the value for Emails to be an explicit nil

### UnsetEmails
`func (o *ContractorContact) UnsetEmails()`

UnsetEmails ensures that no value is present for Emails, not even an explicit nil
### GetComment

`func (o *ContractorContact) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ContractorContact) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ContractorContact) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ContractorContact) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *ContractorContact) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *ContractorContact) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


