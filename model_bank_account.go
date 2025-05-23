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

// checks if the BankAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BankAccount{}

// BankAccount struct for BankAccount
type BankAccount struct {
	// Идентификатор банковского счёта
	Id string `json:"id"`
	// Номер банковского счёта
	AccountNumber NullableString `json:"accountNumber,omitempty"`
	// Информация о банке
	Bank Bank `json:"bank"`
	// Закрытость счета
	Closed bool `json:"closed"`
}

type _BankAccount BankAccount

// NewBankAccount instantiates a new BankAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankAccount(id string, bank Bank, closed bool) *BankAccount {
	this := BankAccount{}
	this.Id = id
	this.Bank = bank
	this.Closed = closed
	return &this
}

// NewBankAccountWithDefaults instantiates a new BankAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankAccountWithDefaults() *BankAccount {
	this := BankAccount{}
	return &this
}

// GetId returns the Id field value
func (o *BankAccount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BankAccount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BankAccount) SetId(v string) {
	o.Id = v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BankAccount) GetAccountNumber() string {
	if o == nil || IsNil(o.AccountNumber.Get()) {
		var ret string
		return ret
	}
	return *o.AccountNumber.Get()
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankAccount) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountNumber.Get(), o.AccountNumber.IsSet()
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *BankAccount) HasAccountNumber() bool {
	if o != nil && o.AccountNumber.IsSet() {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given NullableString and assigns it to the AccountNumber field.
func (o *BankAccount) SetAccountNumber(v string) {
	o.AccountNumber.Set(&v)
}
// SetAccountNumberNil sets the value for AccountNumber to be an explicit nil
func (o *BankAccount) SetAccountNumberNil() {
	o.AccountNumber.Set(nil)
}

// UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil
func (o *BankAccount) UnsetAccountNumber() {
	o.AccountNumber.Unset()
}

// GetBank returns the Bank field value
func (o *BankAccount) GetBank() Bank {
	if o == nil {
		var ret Bank
		return ret
	}

	return o.Bank
}

// GetBankOk returns a tuple with the Bank field value
// and a boolean to check if the value has been set.
func (o *BankAccount) GetBankOk() (*Bank, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bank, true
}

// SetBank sets field value
func (o *BankAccount) SetBank(v Bank) {
	o.Bank = v
}

// GetClosed returns the Closed field value
func (o *BankAccount) GetClosed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Closed
}

// GetClosedOk returns a tuple with the Closed field value
// and a boolean to check if the value has been set.
func (o *BankAccount) GetClosedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Closed, true
}

// SetClosed sets field value
func (o *BankAccount) SetClosed(v bool) {
	o.Closed = v
}

func (o BankAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BankAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.AccountNumber.IsSet() {
		toSerialize["accountNumber"] = o.AccountNumber.Get()
	}
	toSerialize["bank"] = o.Bank
	toSerialize["closed"] = o.Closed
	return toSerialize, nil
}

func (o *BankAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"bank",
		"closed",
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

	varBankAccount := _BankAccount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBankAccount)

	if err != nil {
		return err
	}

	*o = BankAccount(varBankAccount)

	return err
}

type NullableBankAccount struct {
	value *BankAccount
	isSet bool
}

func (v NullableBankAccount) Get() *BankAccount {
	return v.value
}

func (v *NullableBankAccount) Set(val *BankAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableBankAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableBankAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankAccount(val *BankAccount) *NullableBankAccount {
	return &NullableBankAccount{value: val, isSet: true}
}

func (v NullableBankAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


