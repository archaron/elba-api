# UpdateProductRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Names** | [**[]ProductNameToUpdate**](ProductNameToUpdate.md) | Список названий | 
**Units** | [**[]ProductUnitToUpdate**](ProductUnitToUpdate.md) | Список единиц измерений | 
**Article** | Pointer to **NullableString** | Артикул | [optional] 
**Comment** | Pointer to **NullableString** | Комментарий | [optional] 

## Methods

### NewUpdateProductRequest

`func NewUpdateProductRequest(names []ProductNameToUpdate, units []ProductUnitToUpdate, ) *UpdateProductRequest`

NewUpdateProductRequest instantiates a new UpdateProductRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProductRequestWithDefaults

`func NewUpdateProductRequestWithDefaults() *UpdateProductRequest`

NewUpdateProductRequestWithDefaults instantiates a new UpdateProductRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNames

`func (o *UpdateProductRequest) GetNames() []ProductNameToUpdate`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *UpdateProductRequest) GetNamesOk() (*[]ProductNameToUpdate, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *UpdateProductRequest) SetNames(v []ProductNameToUpdate)`

SetNames sets Names field to given value.


### GetUnits

`func (o *UpdateProductRequest) GetUnits() []ProductUnitToUpdate`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *UpdateProductRequest) GetUnitsOk() (*[]ProductUnitToUpdate, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *UpdateProductRequest) SetUnits(v []ProductUnitToUpdate)`

SetUnits sets Units field to given value.


### GetArticle

`func (o *UpdateProductRequest) GetArticle() string`

GetArticle returns the Article field if non-nil, zero value otherwise.

### GetArticleOk

`func (o *UpdateProductRequest) GetArticleOk() (*string, bool)`

GetArticleOk returns a tuple with the Article field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticle

`func (o *UpdateProductRequest) SetArticle(v string)`

SetArticle sets Article field to given value.

### HasArticle

`func (o *UpdateProductRequest) HasArticle() bool`

HasArticle returns a boolean if a field has been set.

### SetArticleNil

`func (o *UpdateProductRequest) SetArticleNil(b bool)`

 SetArticleNil sets the value for Article to be an explicit nil

### UnsetArticle
`func (o *UpdateProductRequest) UnsetArticle()`

UnsetArticle ensures that no value is present for Article, not even an explicit nil
### GetComment

`func (o *UpdateProductRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateProductRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateProductRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateProductRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *UpdateProductRequest) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *UpdateProductRequest) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


