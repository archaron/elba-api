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

// LongRunningTaskStatus   inProcess (Операция в обработке)  completed (Операция успешно обработана)  failed (Ошибка при выполнении операции)
type LongRunningTaskStatus string

// List of LongRunningTaskStatus
const (
	InProcess LongRunningTaskStatus = "inProcess"
	Completed LongRunningTaskStatus = "completed"
	Failed LongRunningTaskStatus = "failed"
)

// All allowed values of LongRunningTaskStatus enum
var AllowedLongRunningTaskStatusEnumValues = []LongRunningTaskStatus{
	"inProcess",
	"completed",
	"failed",
}

func (v *LongRunningTaskStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LongRunningTaskStatus(value)
	for _, existing := range AllowedLongRunningTaskStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LongRunningTaskStatus", value)
}

// NewLongRunningTaskStatusFromValue returns a pointer to a valid LongRunningTaskStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLongRunningTaskStatusFromValue(v string) (*LongRunningTaskStatus, error) {
	ev := LongRunningTaskStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LongRunningTaskStatus: valid values are %v", v, AllowedLongRunningTaskStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LongRunningTaskStatus) IsValid() bool {
	for _, existing := range AllowedLongRunningTaskStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LongRunningTaskStatus value
func (v LongRunningTaskStatus) Ptr() *LongRunningTaskStatus {
	return &v
}

type NullableLongRunningTaskStatus struct {
	value *LongRunningTaskStatus
	isSet bool
}

func (v NullableLongRunningTaskStatus) Get() *LongRunningTaskStatus {
	return v.value
}

func (v *NullableLongRunningTaskStatus) Set(val *LongRunningTaskStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableLongRunningTaskStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableLongRunningTaskStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLongRunningTaskStatus(val *LongRunningTaskStatus) *NullableLongRunningTaskStatus {
	return &NullableLongRunningTaskStatus{value: val, isSet: true}
}

func (v NullableLongRunningTaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLongRunningTaskStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

