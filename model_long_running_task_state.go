/*
Elba Public API

  ## С чего начать    Для работы с API нужно выпустить API-ключ — уникальный токен, позволяющий авторизовывать ваши запросы в API Контур.Эльбы.    #### Как получить API-ключ    1. Откройте Эльбу, в верхнем правом углу нажмите «Настройки и оплата» → «Настройки сервиса».  2. Перейдите на вкладку «API».  2. Нажмите на кнопку «Выпустить ключ». После этого откроется окно со сгенерированным API-ключом.  3. В открывшемся окне появится ваш API-ключ. Скопируйте и сохраните его в надежном месте, потому что он будет показан только один раз. Это сделано в целях безопасности — мы не храним ключи на своей стороне.

API version: v1
Contact: e@kontur.ru
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package elba

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the LongRunningTaskState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LongRunningTaskState{}

// LongRunningTaskState struct for LongRunningTaskState
type LongRunningTaskState struct {
	// Идентификатор асинхронной операции
	Id string `json:"id"`
	// Статус асинхронной операции  inProcess (Операция в обработке)  completed (Операция успешно обработана)  failed (Ошибка при выполнении операции)
	Status LongRunningTaskStatus `json:"status"`
	// Тип операции  searchContractors (Операция поиска контрагентов)  searchProducts (Операция поиска товаров)  createProduct (Операция создания товара)  updateProduct (Операция обновления товара)
	Type LongRunningTaskType `json:"type"`
	Result NullableLongRunningTaskStateResult `json:"result,omitempty"`
	// Информация об ошибке, возникшей во время выполнения операции. Заполнена, если операция завершена с ошибкой.
	ErrorDescription NullableLongRunningTaskExecutionError `json:"errorDescription,omitempty"`
}

type _LongRunningTaskState LongRunningTaskState

// NewLongRunningTaskState instantiates a new LongRunningTaskState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLongRunningTaskState(id string, status LongRunningTaskStatus, type_ LongRunningTaskType) *LongRunningTaskState {
	this := LongRunningTaskState{}
	this.Id = id
	this.Status = status
	this.Type = type_
	return &this
}

// NewLongRunningTaskStateWithDefaults instantiates a new LongRunningTaskState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLongRunningTaskStateWithDefaults() *LongRunningTaskState {
	this := LongRunningTaskState{}
	return &this
}

// GetId returns the Id field value
func (o *LongRunningTaskState) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LongRunningTaskState) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LongRunningTaskState) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *LongRunningTaskState) GetStatus() LongRunningTaskStatus {
	if o == nil {
		var ret LongRunningTaskStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LongRunningTaskState) GetStatusOk() (*LongRunningTaskStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *LongRunningTaskState) SetStatus(v LongRunningTaskStatus) {
	o.Status = v
}

// GetType returns the Type field value
func (o *LongRunningTaskState) GetType() LongRunningTaskType {
	if o == nil {
		var ret LongRunningTaskType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LongRunningTaskState) GetTypeOk() (*LongRunningTaskType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LongRunningTaskState) SetType(v LongRunningTaskType) {
	o.Type = v
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LongRunningTaskState) GetResult() LongRunningTaskStateResult {
	if o == nil || IsNil(o.Result.Get()) {
		var ret LongRunningTaskStateResult
		return ret
	}
	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LongRunningTaskState) GetResultOk() (*LongRunningTaskStateResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// HasResult returns a boolean if a field has been set.
func (o *LongRunningTaskState) HasResult() bool {
	if o != nil && o.Result.IsSet() {
		return true
	}

	return false
}

// SetResult gets a reference to the given NullableLongRunningTaskStateResult and assigns it to the Result field.
func (o *LongRunningTaskState) SetResult(v LongRunningTaskStateResult) {
	o.Result.Set(&v)
}
// SetResultNil sets the value for Result to be an explicit nil
func (o *LongRunningTaskState) SetResultNil() {
	o.Result.Set(nil)
}

// UnsetResult ensures that no value is present for Result, not even an explicit nil
func (o *LongRunningTaskState) UnsetResult() {
	o.Result.Unset()
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LongRunningTaskState) GetErrorDescription() LongRunningTaskExecutionError {
	if o == nil || IsNil(o.ErrorDescription.Get()) {
		var ret LongRunningTaskExecutionError
		return ret
	}
	return *o.ErrorDescription.Get()
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LongRunningTaskState) GetErrorDescriptionOk() (*LongRunningTaskExecutionError, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorDescription.Get(), o.ErrorDescription.IsSet()
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *LongRunningTaskState) HasErrorDescription() bool {
	if o != nil && o.ErrorDescription.IsSet() {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given NullableLongRunningTaskExecutionError and assigns it to the ErrorDescription field.
func (o *LongRunningTaskState) SetErrorDescription(v LongRunningTaskExecutionError) {
	o.ErrorDescription.Set(&v)
}
// SetErrorDescriptionNil sets the value for ErrorDescription to be an explicit nil
func (o *LongRunningTaskState) SetErrorDescriptionNil() {
	o.ErrorDescription.Set(nil)
}

// UnsetErrorDescription ensures that no value is present for ErrorDescription, not even an explicit nil
func (o *LongRunningTaskState) UnsetErrorDescription() {
	o.ErrorDescription.Unset()
}

func (o LongRunningTaskState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LongRunningTaskState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	if o.Result.IsSet() {
		toSerialize["result"] = o.Result.Get()
	}
	if o.ErrorDescription.IsSet() {
		toSerialize["errorDescription"] = o.ErrorDescription.Get()
	}
	return toSerialize, nil
}

func (o *LongRunningTaskState) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"status",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varLongRunningTaskState := _LongRunningTaskState{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLongRunningTaskState)

	if err != nil {
		return err
	}

	*o = LongRunningTaskState(varLongRunningTaskState)

	return err
}

type NullableLongRunningTaskState struct {
	value *LongRunningTaskState
	isSet bool
}

func (v NullableLongRunningTaskState) Get() *LongRunningTaskState {
	return v.value
}

func (v *NullableLongRunningTaskState) Set(val *LongRunningTaskState) {
	v.value = val
	v.isSet = true
}

func (v NullableLongRunningTaskState) IsSet() bool {
	return v.isSet
}

func (v *NullableLongRunningTaskState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLongRunningTaskState(val *LongRunningTaskState) *NullableLongRunningTaskState {
	return &NullableLongRunningTaskState{value: val, isSet: true}
}

func (v NullableLongRunningTaskState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLongRunningTaskState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


