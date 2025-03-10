# GetDocumentNewsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**News** | [**[]DocumentNews**](DocumentNews.md) | Новости | 

## Methods

### NewGetDocumentNewsResponse

`func NewGetDocumentNewsResponse(news []DocumentNews, ) *GetDocumentNewsResponse`

NewGetDocumentNewsResponse instantiates a new GetDocumentNewsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDocumentNewsResponseWithDefaults

`func NewGetDocumentNewsResponseWithDefaults() *GetDocumentNewsResponse`

NewGetDocumentNewsResponseWithDefaults instantiates a new GetDocumentNewsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNews

`func (o *GetDocumentNewsResponse) GetNews() []DocumentNews`

GetNews returns the News field if non-nil, zero value otherwise.

### GetNewsOk

`func (o *GetDocumentNewsResponse) GetNewsOk() (*[]DocumentNews, bool)`

GetNewsOk returns a tuple with the News field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNews

`func (o *GetDocumentNewsResponse) SetNews(v []DocumentNews)`

SetNews sets News field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


