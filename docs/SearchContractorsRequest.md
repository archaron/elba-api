# SearchContractorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | Pointer to **NullableInt32** | Позиция, начиная с которой будут вычитываться результаты поиска контрагентов. По-умолчанию - 0 | [optional] 
**Limit** | Pointer to **NullableInt32** | Количество контрагентов, которое необходимо получить при поиске. По-умолчанию - 100 | [optional] 
**Filter** | Pointer to [**NullableSearchContractorsFilter**](SearchContractorsFilter.md) | Фильтр поиска. Хотябы один из параметров фильтра должен быть заполнен | [optional] 

## Methods

### NewSearchContractorsRequest

`func NewSearchContractorsRequest() *SearchContractorsRequest`

NewSearchContractorsRequest instantiates a new SearchContractorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchContractorsRequestWithDefaults

`func NewSearchContractorsRequestWithDefaults() *SearchContractorsRequest`

NewSearchContractorsRequestWithDefaults instantiates a new SearchContractorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *SearchContractorsRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SearchContractorsRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SearchContractorsRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SearchContractorsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *SearchContractorsRequest) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *SearchContractorsRequest) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetLimit

`func (o *SearchContractorsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchContractorsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchContractorsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchContractorsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *SearchContractorsRequest) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *SearchContractorsRequest) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetFilter

`func (o *SearchContractorsRequest) GetFilter() SearchContractorsFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *SearchContractorsRequest) GetFilterOk() (*SearchContractorsFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *SearchContractorsRequest) SetFilter(v SearchContractorsFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *SearchContractorsRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *SearchContractorsRequest) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *SearchContractorsRequest) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


