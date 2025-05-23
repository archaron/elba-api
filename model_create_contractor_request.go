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

// checks if the CreateContractorRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateContractorRequest{}

// CreateContractorRequest struct for CreateContractorRequest
type CreateContractorRequest struct {
	// Полное название контрагента
	Name string `json:"name"`
	// ИНН
	Inn NullableString `json:"inn,omitempty"`
	// КПП
	Kpp NullableString `json:"kpp,omitempty"`
	// Краткое название контрагента
	ShortName NullableString `json:"shortName,omitempty"`
	// Юридический адрес
	Address NullableString `json:"address,omitempty"`
	// Фактический адрес
	PostAddress NullableString `json:"postAddress,omitempty"`
	// ОКПО
	Okpo NullableString `json:"okpo,omitempty"`
	// ОГРН
	Ogrn NullableString `json:"ogrn,omitempty"`
	// Контакты
	Contacts []ContractorContact `json:"contacts,omitempty"`
}

type _CreateContractorRequest CreateContractorRequest

// NewCreateContractorRequest instantiates a new CreateContractorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateContractorRequest(name string) *CreateContractorRequest {
	this := CreateContractorRequest{}
	this.Name = name
	return &this
}

// NewCreateContractorRequestWithDefaults instantiates a new CreateContractorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateContractorRequestWithDefaults() *CreateContractorRequest {
	this := CreateContractorRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateContractorRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateContractorRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateContractorRequest) SetName(v string) {
	o.Name = v
}

// GetInn returns the Inn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateContractorRequest) GetInn() string {
	if o == nil || IsNil(o.Inn.Get()) {
		var ret string
		return ret
	}
	return *o.Inn.Get()
}

// GetInnOk returns a tuple with the Inn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateContractorRequest) GetInnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inn.Get(), o.Inn.IsSet()
}

// HasInn returns a boolean if a field has been set.
func (o *CreateContractorRequest) HasInn() bool {
	if o != nil && o.Inn.IsSet() {
		return true
	}

	return false
}

// SetInn gets a reference to the given NullableString and assigns it to the Inn field.
func (o *CreateContractorRequest) SetInn(v string) {
	o.Inn.Set(&v)
}
// SetInnNil sets the value for Inn to be an explicit nil
func (o *CreateContractorRequest) SetInnNil() {
	o.Inn.Set(nil)
}

// UnsetInn ensures that no value is present for Inn, not even an explicit nil
func (o *CreateContractorRequest) UnsetInn() {
	o.Inn.Unset()
}

// GetKpp returns the Kpp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateContractorRequest) GetKpp() string {
	if o == nil || IsNil(o.Kpp.Get()) {
		var ret string
		return ret
	}
	return *o.Kpp.Get()
}

// GetKppOk returns a tuple with the Kpp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateContractorRequest) GetKppOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Kpp.Get(), o.Kpp.IsSet()
}

// HasKpp returns a boolean if a field has been set.
func (o *CreateContractorRequest) HasKpp() bool {
	if o != nil && o.Kpp.IsSet() {
		return true
	}

	return false
}

// SetKpp gets a reference to the given NullableString and assigns it to the Kpp field.
func (o *CreateContractorRequest) SetKpp(v string) {
	o.Kpp.Set(&v)
}
// SetKppNil sets the value for Kpp to be an explicit nil
func (o *CreateContractorRequest) SetKppNil() {
	o.Kpp.Set(nil)
}

// UnsetKpp ensures that no value is present for Kpp, not even an explicit nil
func (o *CreateContractorRequest) UnsetKpp() {
	o.Kpp.Unset()
}

// GetShortName returns the ShortName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateContractorRequest) GetShortName() string {
	if o == nil || IsNil(o.ShortName.Get()) {
		var ret string
		return ret
	}
	return *o.ShortName.Get()
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateContractorRequest) GetShortNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShortName.Get(), o.ShortName.IsSet()
}

// HasShortName returns a boolean if a field has been set.
func (o *CreateContractorRequest) HasShortName() bool {
	if o != nil && o.ShortName.IsSet() {
		return true
	}

	return false
}

// SetShortName gets a reference to the given NullableString and assigns it to the ShortName field.
func (o *CreateContractorRequest) SetShortName(v string) {
	o.ShortName.Set(&v)
}
// SetShortNameNil sets the value for ShortName to be an explicit nil
func (o *CreateContractorRequest) SetShortNameNil() {
	o.ShortName.Set(nil)
}

// UnsetShortName ensures that no value is present for ShortName, not even an explicit nil
func (o *CreateContractorRequest) UnsetShortName() {
	o.ShortName.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateContractorRequest) GetAddress() string {
	if o == nil || IsNil(o.Address.Get()) {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateContractorRequest) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *CreateContractorRequest) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *CreateContractorRequest) SetAddress(v string) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *CreateContractorRequest) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *CreateContractorRequest) UnsetAddress() {
	o.Address.Unset()
}

// GetPostAddress returns the PostAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateContractorRequest) GetPostAddress() string {
	if o == nil || IsNil(o.PostAddress.Get()) {
		var ret string
		return ret
	}
	return *o.PostAddress.Get()
}

// GetPostAddressOk returns a tuple with the PostAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateContractorRequest) GetPostAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostAddress.Get(), o.PostAddress.IsSet()
}

// HasPostAddress returns a boolean if a field has been set.
func (o *CreateContractorRequest) HasPostAddress() bool {
	if o != nil && o.PostAddress.IsSet() {
		return true
	}

	return false
}

// SetPostAddress gets a reference to the given NullableString and assigns it to the PostAddress field.
func (o *CreateContractorRequest) SetPostAddress(v string) {
	o.PostAddress.Set(&v)
}
// SetPostAddressNil sets the value for PostAddress to be an explicit nil
func (o *CreateContractorRequest) SetPostAddressNil() {
	o.PostAddress.Set(nil)
}

// UnsetPostAddress ensures that no value is present for PostAddress, not even an explicit nil
func (o *CreateContractorRequest) UnsetPostAddress() {
	o.PostAddress.Unset()
}

// GetOkpo returns the Okpo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateContractorRequest) GetOkpo() string {
	if o == nil || IsNil(o.Okpo.Get()) {
		var ret string
		return ret
	}
	return *o.Okpo.Get()
}

// GetOkpoOk returns a tuple with the Okpo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateContractorRequest) GetOkpoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Okpo.Get(), o.Okpo.IsSet()
}

// HasOkpo returns a boolean if a field has been set.
func (o *CreateContractorRequest) HasOkpo() bool {
	if o != nil && o.Okpo.IsSet() {
		return true
	}

	return false
}

// SetOkpo gets a reference to the given NullableString and assigns it to the Okpo field.
func (o *CreateContractorRequest) SetOkpo(v string) {
	o.Okpo.Set(&v)
}
// SetOkpoNil sets the value for Okpo to be an explicit nil
func (o *CreateContractorRequest) SetOkpoNil() {
	o.Okpo.Set(nil)
}

// UnsetOkpo ensures that no value is present for Okpo, not even an explicit nil
func (o *CreateContractorRequest) UnsetOkpo() {
	o.Okpo.Unset()
}

// GetOgrn returns the Ogrn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateContractorRequest) GetOgrn() string {
	if o == nil || IsNil(o.Ogrn.Get()) {
		var ret string
		return ret
	}
	return *o.Ogrn.Get()
}

// GetOgrnOk returns a tuple with the Ogrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateContractorRequest) GetOgrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ogrn.Get(), o.Ogrn.IsSet()
}

// HasOgrn returns a boolean if a field has been set.
func (o *CreateContractorRequest) HasOgrn() bool {
	if o != nil && o.Ogrn.IsSet() {
		return true
	}

	return false
}

// SetOgrn gets a reference to the given NullableString and assigns it to the Ogrn field.
func (o *CreateContractorRequest) SetOgrn(v string) {
	o.Ogrn.Set(&v)
}
// SetOgrnNil sets the value for Ogrn to be an explicit nil
func (o *CreateContractorRequest) SetOgrnNil() {
	o.Ogrn.Set(nil)
}

// UnsetOgrn ensures that no value is present for Ogrn, not even an explicit nil
func (o *CreateContractorRequest) UnsetOgrn() {
	o.Ogrn.Unset()
}

// GetContacts returns the Contacts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateContractorRequest) GetContacts() []ContractorContact {
	if o == nil {
		var ret []ContractorContact
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateContractorRequest) GetContactsOk() ([]ContractorContact, bool) {
	if o == nil || IsNil(o.Contacts) {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *CreateContractorRequest) HasContacts() bool {
	if o != nil && !IsNil(o.Contacts) {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []ContractorContact and assigns it to the Contacts field.
func (o *CreateContractorRequest) SetContacts(v []ContractorContact) {
	o.Contacts = v
}

func (o CreateContractorRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateContractorRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Inn.IsSet() {
		toSerialize["inn"] = o.Inn.Get()
	}
	if o.Kpp.IsSet() {
		toSerialize["kpp"] = o.Kpp.Get()
	}
	if o.ShortName.IsSet() {
		toSerialize["shortName"] = o.ShortName.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.PostAddress.IsSet() {
		toSerialize["postAddress"] = o.PostAddress.Get()
	}
	if o.Okpo.IsSet() {
		toSerialize["okpo"] = o.Okpo.Get()
	}
	if o.Ogrn.IsSet() {
		toSerialize["ogrn"] = o.Ogrn.Get()
	}
	if o.Contacts != nil {
		toSerialize["contacts"] = o.Contacts
	}
	return toSerialize, nil
}

func (o *CreateContractorRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCreateContractorRequest := _CreateContractorRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateContractorRequest)

	if err != nil {
		return err
	}

	*o = CreateContractorRequest(varCreateContractorRequest)

	return err
}

type NullableCreateContractorRequest struct {
	value *CreateContractorRequest
	isSet bool
}

func (v NullableCreateContractorRequest) Get() *CreateContractorRequest {
	return v.value
}

func (v *NullableCreateContractorRequest) Set(val *CreateContractorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateContractorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateContractorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateContractorRequest(val *CreateContractorRequest) *NullableCreateContractorRequest {
	return &NullableCreateContractorRequest{value: val, isSet: true}
}

func (v NullableCreateContractorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateContractorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


