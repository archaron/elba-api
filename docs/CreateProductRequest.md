# CreateProductRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Names** | [**[]ProductNameToCreate**](ProductNameToCreate.md) | Список названий товара | 
**Units** | [**[]ProductUnitToCreate**](ProductUnitToCreate.md) | Список единиц измерений товара | 
**Article** | Pointer to **NullableString** | Артикул | [optional] 
**Comment** | Pointer to **NullableString** | Комментарий | [optional] 
**Amount** | Pointer to **NullableFloat64** | Количество на складе | [optional] 

## Methods

### NewCreateProductRequest

`func NewCreateProductRequest(names []ProductNameToCreate, units []ProductUnitToCreate, ) *CreateProductRequest`

NewCreateProductRequest instantiates a new CreateProductRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProductRequestWithDefaults

`func NewCreateProductRequestWithDefaults() *CreateProductRequest`

NewCreateProductRequestWithDefaults instantiates a new CreateProductRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNames

`func (o *CreateProductRequest) GetNames() []ProductNameToCreate`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *CreateProductRequest) GetNamesOk() (*[]ProductNameToCreate, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *CreateProductRequest) SetNames(v []ProductNameToCreate)`

SetNames sets Names field to given value.


### GetUnits

`func (o *CreateProductRequest) GetUnits() []ProductUnitToCreate`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *CreateProductRequest) GetUnitsOk() (*[]ProductUnitToCreate, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *CreateProductRequest) SetUnits(v []ProductUnitToCreate)`

SetUnits sets Units field to given value.


### GetArticle

`func (o *CreateProductRequest) GetArticle() string`

GetArticle returns the Article field if non-nil, zero value otherwise.

### GetArticleOk

`func (o *CreateProductRequest) GetArticleOk() (*string, bool)`

GetArticleOk returns a tuple with the Article field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticle

`func (o *CreateProductRequest) SetArticle(v string)`

SetArticle sets Article field to given value.

### HasArticle

`func (o *CreateProductRequest) HasArticle() bool`

HasArticle returns a boolean if a field has been set.

### SetArticleNil

`func (o *CreateProductRequest) SetArticleNil(b bool)`

 SetArticleNil sets the value for Article to be an explicit nil

### UnsetArticle
`func (o *CreateProductRequest) UnsetArticle()`

UnsetArticle ensures that no value is present for Article, not even an explicit nil
### GetComment

`func (o *CreateProductRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateProductRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateProductRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateProductRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *CreateProductRequest) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *CreateProductRequest) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetAmount

`func (o *CreateProductRequest) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateProductRequest) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateProductRequest) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CreateProductRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *CreateProductRequest) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *CreateProductRequest) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


