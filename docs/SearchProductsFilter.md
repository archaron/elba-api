# SearchProductsFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductName** | Pointer to **NullableString** | Название товара | [optional] 
**Article** | Pointer to **NullableString** | Артикул | [optional] 

## Methods

### NewSearchProductsFilter

`func NewSearchProductsFilter() *SearchProductsFilter`

NewSearchProductsFilter instantiates a new SearchProductsFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchProductsFilterWithDefaults

`func NewSearchProductsFilterWithDefaults() *SearchProductsFilter`

NewSearchProductsFilterWithDefaults instantiates a new SearchProductsFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductName

`func (o *SearchProductsFilter) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *SearchProductsFilter) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *SearchProductsFilter) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *SearchProductsFilter) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *SearchProductsFilter) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *SearchProductsFilter) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetArticle

`func (o *SearchProductsFilter) GetArticle() string`

GetArticle returns the Article field if non-nil, zero value otherwise.

### GetArticleOk

`func (o *SearchProductsFilter) GetArticleOk() (*string, bool)`

GetArticleOk returns a tuple with the Article field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticle

`func (o *SearchProductsFilter) SetArticle(v string)`

SetArticle sets Article field to given value.

### HasArticle

`func (o *SearchProductsFilter) HasArticle() bool`

HasArticle returns a boolean if a field has been set.

### SetArticleNil

`func (o *SearchProductsFilter) SetArticleNil(b bool)`

 SetArticleNil sets the value for Article to be an explicit nil

### UnsetArticle
`func (o *SearchProductsFilter) UnsetArticle()`

UnsetArticle ensures that no value is present for Article, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


