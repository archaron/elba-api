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

// checks if the Organization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Organization{}

// Organization struct for Organization
type Organization struct {
	// Идентификатор организации
	Id string `json:"id"`
	// ИНН
	Inn NullableString `json:"inn,omitempty"`
	// КПП
	Kpp NullableString `json:"kpp,omitempty"`
}

type _Organization Organization

// NewOrganization instantiates a new Organization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganization(id string) *Organization {
	this := Organization{}
	this.Id = id
	return &this
}

// NewOrganizationWithDefaults instantiates a new Organization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationWithDefaults() *Organization {
	this := Organization{}
	return &this
}

// GetId returns the Id field value
func (o *Organization) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Organization) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Organization) SetId(v string) {
	o.Id = v
}

// GetInn returns the Inn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Organization) GetInn() string {
	if o == nil || IsNil(o.Inn.Get()) {
		var ret string
		return ret
	}
	return *o.Inn.Get()
}

// GetInnOk returns a tuple with the Inn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Organization) GetInnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inn.Get(), o.Inn.IsSet()
}

// HasInn returns a boolean if a field has been set.
func (o *Organization) HasInn() bool {
	if o != nil && o.Inn.IsSet() {
		return true
	}

	return false
}

// SetInn gets a reference to the given NullableString and assigns it to the Inn field.
func (o *Organization) SetInn(v string) {
	o.Inn.Set(&v)
}
// SetInnNil sets the value for Inn to be an explicit nil
func (o *Organization) SetInnNil() {
	o.Inn.Set(nil)
}

// UnsetInn ensures that no value is present for Inn, not even an explicit nil
func (o *Organization) UnsetInn() {
	o.Inn.Unset()
}

// GetKpp returns the Kpp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Organization) GetKpp() string {
	if o == nil || IsNil(o.Kpp.Get()) {
		var ret string
		return ret
	}
	return *o.Kpp.Get()
}

// GetKppOk returns a tuple with the Kpp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Organization) GetKppOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Kpp.Get(), o.Kpp.IsSet()
}

// HasKpp returns a boolean if a field has been set.
func (o *Organization) HasKpp() bool {
	if o != nil && o.Kpp.IsSet() {
		return true
	}

	return false
}

// SetKpp gets a reference to the given NullableString and assigns it to the Kpp field.
func (o *Organization) SetKpp(v string) {
	o.Kpp.Set(&v)
}
// SetKppNil sets the value for Kpp to be an explicit nil
func (o *Organization) SetKppNil() {
	o.Kpp.Set(nil)
}

// UnsetKpp ensures that no value is present for Kpp, not even an explicit nil
func (o *Organization) UnsetKpp() {
	o.Kpp.Unset()
}

func (o Organization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Organization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Inn.IsSet() {
		toSerialize["inn"] = o.Inn.Get()
	}
	if o.Kpp.IsSet() {
		toSerialize["kpp"] = o.Kpp.Get()
	}
	return toSerialize, nil
}

func (o *Organization) UnmarshalJSON(data []byte) (err error) {
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

	varOrganization := _Organization{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrganization)

	if err != nil {
		return err
	}

	*o = Organization(varOrganization)

	return err
}

type NullableOrganization struct {
	value *Organization
	isSet bool
}

func (v NullableOrganization) Get() *Organization {
	return v.value
}

func (v *NullableOrganization) Set(val *Organization) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganization(val *Organization) *NullableOrganization {
	return &NullableOrganization{value: val, isSet: true}
}

func (v NullableOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


