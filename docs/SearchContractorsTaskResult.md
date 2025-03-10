# SearchContractorsTaskResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contractors** | [**[]FoundContractor**](FoundContractor.md) | Список найденных контрагентов | 

## Methods

### NewSearchContractorsTaskResult

`func NewSearchContractorsTaskResult(contractors []FoundContractor, ) *SearchContractorsTaskResult`

NewSearchContractorsTaskResult instantiates a new SearchContractorsTaskResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchContractorsTaskResultWithDefaults

`func NewSearchContractorsTaskResultWithDefaults() *SearchContractorsTaskResult`

NewSearchContractorsTaskResultWithDefaults instantiates a new SearchContractorsTaskResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractors

`func (o *SearchContractorsTaskResult) GetContractors() []FoundContractor`

GetContractors returns the Contractors field if non-nil, zero value otherwise.

### GetContractorsOk

`func (o *SearchContractorsTaskResult) GetContractorsOk() (*[]FoundContractor, bool)`

GetContractorsOk returns a tuple with the Contractors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractors

`func (o *SearchContractorsTaskResult) SetContractors(v []FoundContractor)`

SetContractors sets Contractors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


