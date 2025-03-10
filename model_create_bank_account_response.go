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

// checks if the CreateBankAccountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBankAccountResponse{}

// CreateBankAccountResponse struct for CreateBankAccountResponse
type CreateBankAccountResponse struct {
	// Идентификатор созданного банковского счёта
	Id string `json:"id"`
}

type _CreateBankAccountResponse CreateBankAccountResponse

// NewCreateBankAccountResponse instantiates a new CreateBankAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBankAccountResponse(id string) *CreateBankAccountResponse {
	this := CreateBankAccountResponse{}
	this.Id = id
	return &this
}

// NewCreateBankAccountResponseWithDefaults instantiates a new CreateBankAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBankAccountResponseWithDefaults() *CreateBankAccountResponse {
	this := CreateBankAccountResponse{}
	return &this
}

// GetId returns the Id field value
func (o *CreateBankAccountResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateBankAccountResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateBankAccountResponse) SetId(v string) {
	o.Id = v
}

func (o CreateBankAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBankAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *CreateBankAccountResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varCreateBankAccountResponse := _CreateBankAccountResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateBankAccountResponse)

	if err != nil {
		return err
	}

	*o = CreateBankAccountResponse(varCreateBankAccountResponse)

	return err
}

type NullableCreateBankAccountResponse struct {
	value *CreateBankAccountResponse
	isSet bool
}

func (v NullableCreateBankAccountResponse) Get() *CreateBankAccountResponse {
	return v.value
}

func (v *NullableCreateBankAccountResponse) Set(val *CreateBankAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBankAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBankAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBankAccountResponse(val *CreateBankAccountResponse) *NullableCreateBankAccountResponse {
	return &NullableCreateBankAccountResponse{value: val, isSet: true}
}

func (v NullableCreateBankAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBankAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


