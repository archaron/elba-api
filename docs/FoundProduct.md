# FoundProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор товара | 
**ProductMainName** | **string** | Главное название товара | 
**Article** | Pointer to **NullableString** | Артикул | [optional] 

## Methods

### NewFoundProduct

`func NewFoundProduct(id string, productMainName string, ) *FoundProduct`

NewFoundProduct instantiates a new FoundProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFoundProductWithDefaults

`func NewFoundProductWithDefaults() *FoundProduct`

NewFoundProductWithDefaults instantiates a new FoundProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FoundProduct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FoundProduct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FoundProduct) SetId(v string)`

SetId sets Id field to given value.


### GetProductMainName

`func (o *FoundProduct) GetProductMainName() string`

GetProductMainName returns the ProductMainName field if non-nil, zero value otherwise.

### GetProductMainNameOk

`func (o *FoundProduct) GetProductMainNameOk() (*string, bool)`

GetProductMainNameOk returns a tuple with the ProductMainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductMainName

`func (o *FoundProduct) SetProductMainName(v string)`

SetProductMainName sets ProductMainName field to given value.


### GetArticle

`func (o *FoundProduct) GetArticle() string`

GetArticle returns the Article field if non-nil, zero value otherwise.

### GetArticleOk

`func (o *FoundProduct) GetArticleOk() (*string, bool)`

GetArticleOk returns a tuple with the Article field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticle

`func (o *FoundProduct) SetArticle(v string)`

SetArticle sets Article field to given value.

### HasArticle

`func (o *FoundProduct) HasArticle() bool`

HasArticle returns a boolean if a field has been set.

### SetArticleNil

`func (o *FoundProduct) SetArticleNil(b bool)`

 SetArticleNil sets the value for Article to be an explicit nil

### UnsetArticle
`func (o *FoundProduct) UnsetArticle()`

UnsetArticle ensures that no value is present for Article, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


