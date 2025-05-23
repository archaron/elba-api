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
)

// checks if the CreateProductTaskResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProductTaskResult{}

// CreateProductTaskResult struct for CreateProductTaskResult
type CreateProductTaskResult struct {
	// Идентификатор созданного товара. Заполнен, если нет ошибки создания товара
	Id NullableString `json:"id,omitempty"`
	// Ошибка при создании товара  productWithSameNameExists (Уже есть товар с таким названием)
	CreateProductError NullableCreateProductError `json:"createProductError,omitempty"`
}

// NewCreateProductTaskResult instantiates a new CreateProductTaskResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProductTaskResult() *CreateProductTaskResult {
	this := CreateProductTaskResult{}
	return &this
}

// NewCreateProductTaskResultWithDefaults instantiates a new CreateProductTaskResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProductTaskResultWithDefaults() *CreateProductTaskResult {
	this := CreateProductTaskResult{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateProductTaskResult) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateProductTaskResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CreateProductTaskResult) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *CreateProductTaskResult) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *CreateProductTaskResult) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CreateProductTaskResult) UnsetId() {
	o.Id.Unset()
}

// GetCreateProductError returns the CreateProductError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateProductTaskResult) GetCreateProductError() CreateProductError {
	if o == nil || IsNil(o.CreateProductError.Get()) {
		var ret CreateProductError
		return ret
	}
	return *o.CreateProductError.Get()
}

// GetCreateProductErrorOk returns a tuple with the CreateProductError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateProductTaskResult) GetCreateProductErrorOk() (*CreateProductError, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreateProductError.Get(), o.CreateProductError.IsSet()
}

// HasCreateProductError returns a boolean if a field has been set.
func (o *CreateProductTaskResult) HasCreateProductError() bool {
	if o != nil && o.CreateProductError.IsSet() {
		return true
	}

	return false
}

// SetCreateProductError gets a reference to the given NullableCreateProductError and assigns it to the CreateProductError field.
func (o *CreateProductTaskResult) SetCreateProductError(v CreateProductError) {
	o.CreateProductError.Set(&v)
}
// SetCreateProductErrorNil sets the value for CreateProductError to be an explicit nil
func (o *CreateProductTaskResult) SetCreateProductErrorNil() {
	o.CreateProductError.Set(nil)
}

// UnsetCreateProductError ensures that no value is present for CreateProductError, not even an explicit nil
func (o *CreateProductTaskResult) UnsetCreateProductError() {
	o.CreateProductError.Unset()
}

func (o CreateProductTaskResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProductTaskResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.CreateProductError.IsSet() {
		toSerialize["createProductError"] = o.CreateProductError.Get()
	}
	return toSerialize, nil
}

type NullableCreateProductTaskResult struct {
	value *CreateProductTaskResult
	isSet bool
}

func (v NullableCreateProductTaskResult) Get() *CreateProductTaskResult {
	return v.value
}

func (v *NullableCreateProductTaskResult) Set(val *CreateProductTaskResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProductTaskResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProductTaskResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProductTaskResult(val *CreateProductTaskResult) *NullableCreateProductTaskResult {
	return &NullableCreateProductTaskResult{value: val, isSet: true}
}

func (v NullableCreateProductTaskResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProductTaskResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


