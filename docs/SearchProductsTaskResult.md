# SearchProductsTaskResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | [**[]FoundProduct**](FoundProduct.md) | Список найденных товаров | 

## Methods

### NewSearchProductsTaskResult

`func NewSearchProductsTaskResult(products []FoundProduct, ) *SearchProductsTaskResult`

NewSearchProductsTaskResult instantiates a new SearchProductsTaskResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchProductsTaskResultWithDefaults

`func NewSearchProductsTaskResultWithDefaults() *SearchProductsTaskResult`

NewSearchProductsTaskResultWithDefaults instantiates a new SearchProductsTaskResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *SearchProductsTaskResult) GetProducts() []FoundProduct`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *SearchProductsTaskResult) GetProductsOk() (*[]FoundProduct, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *SearchProductsTaskResult) SetProducts(v []FoundProduct)`

SetProducts sets Products field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


