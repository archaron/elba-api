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

// NDSRate   withoutNds (Без НДС)  nds0 (0%)  nds5 (5%)  nds10 (10%)  nds18 (18%)  nds20 (20%)
type NDSRate string

// List of NDSRate
const (
	WithoutNds NDSRate = "withoutNds"
	Nds0 NDSRate = "nds0"
	Nds5 NDSRate = "nds5"
	Nds10 NDSRate = "nds10"
	Nds18 NDSRate = "nds18"
	Nds20 NDSRate = "nds20"
)

// All allowed values of NDSRate enum
var AllowedNDSRateEnumValues = []NDSRate{
	"withoutNds",
	"nds0",
	"nds5",
	"nds10",
	"nds18",
	"nds20",
}

func (v *NDSRate) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NDSRate(value)
	for _, existing := range AllowedNDSRateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NDSRate", value)
}

// NewNDSRateFromValue returns a pointer to a valid NDSRate
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNDSRateFromValue(v string) (*NDSRate, error) {
	ev := NDSRate(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NDSRate: valid values are %v", v, AllowedNDSRateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NDSRate) IsValid() bool {
	for _, existing := range AllowedNDSRateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NDSRate value
func (v NDSRate) Ptr() *NDSRate {
	return &v
}

type NullableNDSRate struct {
	value *NDSRate
	isSet bool
}

func (v NullableNDSRate) Get() *NDSRate {
	return v.value
}

func (v *NullableNDSRate) Set(val *NDSRate) {
	v.value = val
	v.isSet = true
}

func (v NullableNDSRate) IsSet() bool {
	return v.isSet
}

func (v *NullableNDSRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNDSRate(val *NDSRate) *NullableNDSRate {
	return &NullableNDSRate{value: val, isSet: true}
}

func (v NullableNDSRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNDSRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

