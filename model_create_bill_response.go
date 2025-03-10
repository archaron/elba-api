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

// checks if the CreateBillResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBillResponse{}

// CreateBillResponse struct for CreateBillResponse
type CreateBillResponse struct {
	// Идентификатор созданного документа
	Id string `json:"id"`
}

type _CreateBillResponse CreateBillResponse

// NewCreateBillResponse instantiates a new CreateBillResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBillResponse(id string) *CreateBillResponse {
	this := CreateBillResponse{}
	this.Id = id
	return &this
}

// NewCreateBillResponseWithDefaults instantiates a new CreateBillResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBillResponseWithDefaults() *CreateBillResponse {
	this := CreateBillResponse{}
	return &this
}

// GetId returns the Id field value
func (o *CreateBillResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateBillResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateBillResponse) SetId(v string) {
	o.Id = v
}

func (o CreateBillResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBillResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *CreateBillResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCreateBillResponse := _CreateBillResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateBillResponse)

	if err != nil {
		return err
	}

	*o = CreateBillResponse(varCreateBillResponse)

	return err
}

type NullableCreateBillResponse struct {
	value *CreateBillResponse
	isSet bool
}

func (v NullableCreateBillResponse) Get() *CreateBillResponse {
	return v.value
}

func (v *NullableCreateBillResponse) Set(val *CreateBillResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBillResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBillResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBillResponse(val *CreateBillResponse) *NullableCreateBillResponse {
	return &NullableCreateBillResponse{value: val, isSet: true}
}

func (v NullableCreateBillResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBillResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


