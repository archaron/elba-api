# Product

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор товара | 
**Names** | [**[]ProductName**](ProductName.md) | Список названий товара | 
**Units** | [**[]ProductUnit**](ProductUnit.md) | Список единиц измерений товара | 
**Article** | Pointer to **NullableString** | Артикул | [optional] 
**Comment** | Pointer to **NullableString** | Комментарий | [optional] 

## Methods

### NewProduct

`func NewProduct(id string, names []ProductName, units []ProductUnit, ) *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Product) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Product) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Product) SetId(v string)`

SetId sets Id field to given value.


### GetNames

`func (o *Product) GetNames() []ProductName`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *Product) GetNamesOk() (*[]ProductName, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *Product) SetNames(v []ProductName)`

SetNames sets Names field to given value.


### GetUnits

`func (o *Product) GetUnits() []ProductUnit`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *Product) GetUnitsOk() (*[]ProductUnit, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *Product) SetUnits(v []ProductUnit)`

SetUnits sets Units field to given value.


### GetArticle

`func (o *Product) GetArticle() string`

GetArticle returns the Article field if non-nil, zero value otherwise.

### GetArticleOk

`func (o *Product) GetArticleOk() (*string, bool)`

GetArticleOk returns a tuple with the Article field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticle

`func (o *Product) SetArticle(v string)`

SetArticle sets Article field to given value.

### HasArticle

`func (o *Product) HasArticle() bool`

HasArticle returns a boolean if a field has been set.

### SetArticleNil

`func (o *Product) SetArticleNil(b bool)`

 SetArticleNil sets the value for Article to be an explicit nil

### UnsetArticle
`func (o *Product) UnsetArticle()`

UnsetArticle ensures that no value is present for Article, not even an explicit nil
### GetComment

`func (o *Product) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Product) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Product) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Product) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *Product) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *Product) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


