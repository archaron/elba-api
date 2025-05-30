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
	"fmt"
)

// UpdateProductError   productWithSameNameExists (Уже есть товар с таким названием)  deletingProductNameThatUsedInDocuments (Название товара нельзя удалить, так как оно используется в документах)  renamingProductNameThatUsedInDocuments (Название товара нельзя изменить, так как оно используется в документах)  deletingProductUnitThatUsedInDocuments (Единицу измерения нельзя удалить, так как она используется в документах)  renamingProductUnitThatUsedInDocuments (Единицу измерения нельзя переименовать, так как она используется в документах)
type UpdateProductError string

// List of UpdateProductError
const (
	UpdatingProductWithSameNameExists      UpdateProductError = "productWithSameNameExists"
	DeletingProductNameThatUsedInDocuments UpdateProductError = "deletingProductNameThatUsedInDocuments"
	RenamingProductNameThatUsedInDocuments UpdateProductError = "renamingProductNameThatUsedInDocuments"
	DeletingProductUnitThatUsedInDocuments UpdateProductError = "deletingProductUnitThatUsedInDocuments"
	RenamingProductUnitThatUsedInDocuments UpdateProductError = "renamingProductUnitThatUsedInDocuments"
)

// All allowed values of UpdateProductError enum
var AllowedUpdateProductErrorEnumValues = []UpdateProductError{
	"productWithSameNameExists",
	"deletingProductNameThatUsedInDocuments",
	"renamingProductNameThatUsedInDocuments",
	"deletingProductUnitThatUsedInDocuments",
	"renamingProductUnitThatUsedInDocuments",
}

func (v *UpdateProductError) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UpdateProductError(value)
	for _, existing := range AllowedUpdateProductErrorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UpdateProductError", value)
}

// NewUpdateProductErrorFromValue returns a pointer to a valid UpdateProductError
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUpdateProductErrorFromValue(v string) (*UpdateProductError, error) {
	ev := UpdateProductError(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UpdateProductError: valid values are %v", v, AllowedUpdateProductErrorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UpdateProductError) IsValid() bool {
	for _, existing := range AllowedUpdateProductErrorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UpdateProductError value
func (v UpdateProductError) Ptr() *UpdateProductError {
	return &v
}

type NullableUpdateProductError struct {
	value *UpdateProductError
	isSet bool
}

func (v NullableUpdateProductError) Get() *UpdateProductError {
	return v.value
}

func (v *NullableUpdateProductError) Set(val *UpdateProductError) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProductError) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProductError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProductError(val *UpdateProductError) *NullableUpdateProductError {
	return &NullableUpdateProductError{value: val, isSet: true}
}

func (v NullableUpdateProductError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProductError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
