# DocumentNews

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Идентификатор новости | 
**NewsType** | [**NewsType**](NewsType.md) | Тип новости  billPaymentNews | 
**NewsContent** | [**DocumentNewsNewsContent**](DocumentNewsNewsContent.md) |  | 

## Methods

### NewDocumentNews

`func NewDocumentNews(id int64, newsType NewsType, newsContent DocumentNewsNewsContent, ) *DocumentNews`

NewDocumentNews instantiates a new DocumentNews object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentNewsWithDefaults

`func NewDocumentNewsWithDefaults() *DocumentNews`

NewDocumentNewsWithDefaults instantiates a new DocumentNews object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocumentNews) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentNews) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentNews) SetId(v int64)`

SetId sets Id field to given value.


### GetNewsType

`func (o *DocumentNews) GetNewsType() NewsType`

GetNewsType returns the NewsType field if non-nil, zero value otherwise.

### GetNewsTypeOk

`func (o *DocumentNews) GetNewsTypeOk() (*NewsType, bool)`

GetNewsTypeOk returns a tuple with the NewsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewsType

`func (o *DocumentNews) SetNewsType(v NewsType)`

SetNewsType sets NewsType field to given value.


### GetNewsContent

`func (o *DocumentNews) GetNewsContent() DocumentNewsNewsContent`

GetNewsContent returns the NewsContent field if non-nil, zero value otherwise.

### GetNewsContentOk

`func (o *DocumentNews) GetNewsContentOk() (*DocumentNewsNewsContent, bool)`

GetNewsContentOk returns a tuple with the NewsContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewsContent

`func (o *DocumentNews) SetNewsContent(v DocumentNewsNewsContent)`

SetNewsContent sets NewsContent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


