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

// ProductType   product (Товар)  service (Услуга)
type ProductType string

// List of ProductType
const (
	Product ProductType = "product"
	Service ProductType = "service"
)

// All allowed values of ProductType enum
var AllowedProductTypeEnumValues = []ProductType{
	"product",
	"service",
}

func (v *ProductType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProductType(value)
	for _, existing := range AllowedProductTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProductType", value)
}

// NewProductTypeFromValue returns a pointer to a valid ProductType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProductTypeFromValue(v string) (*ProductType, error) {
	ev := ProductType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProductType: valid values are %v", v, AllowedProductTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProductType) IsValid() bool {
	for _, existing := range AllowedProductTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductType value
func (v ProductType) Ptr() *ProductType {
	return &v
}

type NullableProductType struct {
	value *ProductType
	isSet bool
}

func (v NullableProductType) Get() *ProductType {
	return v.value
}

func (v *NullableProductType) Set(val *ProductType) {
	v.value = val
	v.isSet = true
}

func (v NullableProductType) IsSet() bool {
	return v.isSet
}

func (v *NullableProductType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductType(val *ProductType) *NullableProductType {
	return &NullableProductType{value: val, isSet: true}
}

func (v NullableProductType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

