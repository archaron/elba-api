# LongRunningTaskExecutionError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TraceId** | **string** | Идентификатор ошибки | 
**Message** | **string** | Текст ошибки | 

## Methods

### NewLongRunningTaskExecutionError

`func NewLongRunningTaskExecutionError(traceId string, message string, ) *LongRunningTaskExecutionError`

NewLongRunningTaskExecutionError instantiates a new LongRunningTaskExecutionError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLongRunningTaskExecutionErrorWithDefaults

`func NewLongRunningTaskExecutionErrorWithDefaults() *LongRunningTaskExecutionError`

NewLongRunningTaskExecutionErrorWithDefaults instantiates a new LongRunningTaskExecutionError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraceId

`func (o *LongRunningTaskExecutionError) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *LongRunningTaskExecutionError) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *LongRunningTaskExecutionError) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetMessage

`func (o *LongRunningTaskExecutionError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LongRunningTaskExecutionError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LongRunningTaskExecutionError) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


