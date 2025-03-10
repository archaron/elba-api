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

// checks if the ActItemToCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActItemToCreate{}

// ActItemToCreate struct for ActItemToCreate
type ActItemToCreate struct {
	// Наименование товара или услуги
	ProductName string `json:"productName"`
	// Единицы измерения
	UnitName string `json:"unitName"`
	// Количество
	Quantity float64 `json:"quantity"`
	// Цена за единицу товара
	Price NullableFloat64 `json:"price,omitempty"`
	// НДС  withoutNds (Без НДС)  nds0 (0%)  nds5 (5%)  nds10 (10%)  nds20 (20%)
	NdsRate NullableNDSRateToSave `json:"ndsRate,omitempty"`
	// Скидка
	Discount NullableFloat64 `json:"discount,omitempty"`
}

type _ActItemToCreate ActItemToCreate

// NewActItemToCreate instantiates a new ActItemToCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActItemToCreate(productName string, unitName string, quantity float64) *ActItemToCreate {
	this := ActItemToCreate{}
	this.ProductName = productName
	this.UnitName = unitName
	this.Quantity = quantity
	return &this
}

// NewActItemToCreateWithDefaults instantiates a new ActItemToCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActItemToCreateWithDefaults() *ActItemToCreate {
	this := ActItemToCreate{}
	return &this
}

// GetProductName returns the ProductName field value
func (o *ActItemToCreate) GetProductName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value
// and a boolean to check if the value has been set.
func (o *ActItemToCreate) GetProductNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductName, true
}

// SetProductName sets field value
func (o *ActItemToCreate) SetProductName(v string) {
	o.ProductName = v
}

// GetUnitName returns the UnitName field value
func (o *ActItemToCreate) GetUnitName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnitName
}

// GetUnitNameOk returns a tuple with the UnitName field value
// and a boolean to check if the value has been set.
func (o *ActItemToCreate) GetUnitNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitName, true
}

// SetUnitName sets field value
func (o *ActItemToCreate) SetUnitName(v string) {
	o.UnitName = v
}

// GetQuantity returns the Quantity field value
func (o *ActItemToCreate) GetQuantity() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *ActItemToCreate) GetQuantityOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *ActItemToCreate) SetQuantity(v float64) {
	o.Quantity = v
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActItemToCreate) GetPrice() float64 {
	if o == nil || IsNil(o.Price.Get()) {
		var ret float64
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActItemToCreate) GetPriceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *ActItemToCreate) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableFloat64 and assigns it to the Price field.
func (o *ActItemToCreate) SetPrice(v float64) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *ActItemToCreate) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *ActItemToCreate) UnsetPrice() {
	o.Price.Unset()
}

// GetNdsRate returns the NdsRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActItemToCreate) GetNdsRate() NDSRateToSave {
	if o == nil || IsNil(o.NdsRate.Get()) {
		var ret NDSRateToSave
		return ret
	}
	return *o.NdsRate.Get()
}

// GetNdsRateOk returns a tuple with the NdsRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActItemToCreate) GetNdsRateOk() (*NDSRateToSave, bool) {
	if o == nil {
		return nil, false
	}
	return o.NdsRate.Get(), o.NdsRate.IsSet()
}

// HasNdsRate returns a boolean if a field has been set.
func (o *ActItemToCreate) HasNdsRate() bool {
	if o != nil && o.NdsRate.IsSet() {
		return true
	}

	return false
}

// SetNdsRate gets a reference to the given NullableNDSRateToSave and assigns it to the NdsRate field.
func (o *ActItemToCreate) SetNdsRate(v NDSRateToSave) {
	o.NdsRate.Set(&v)
}
// SetNdsRateNil sets the value for NdsRate to be an explicit nil
func (o *ActItemToCreate) SetNdsRateNil() {
	o.NdsRate.Set(nil)
}

// UnsetNdsRate ensures that no value is present for NdsRate, not even an explicit nil
func (o *ActItemToCreate) UnsetNdsRate() {
	o.NdsRate.Unset()
}

// GetDiscount returns the Discount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActItemToCreate) GetDiscount() float64 {
	if o == nil || IsNil(o.Discount.Get()) {
		var ret float64
		return ret
	}
	return *o.Discount.Get()
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActItemToCreate) GetDiscountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Discount.Get(), o.Discount.IsSet()
}

// HasDiscount returns a boolean if a field has been set.
func (o *ActItemToCreate) HasDiscount() bool {
	if o != nil && o.Discount.IsSet() {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given NullableFloat64 and assigns it to the Discount field.
func (o *ActItemToCreate) SetDiscount(v float64) {
	o.Discount.Set(&v)
}
// SetDiscountNil sets the value for Discount to be an explicit nil
func (o *ActItemToCreate) SetDiscountNil() {
	o.Discount.Set(nil)
}

// UnsetDiscount ensures that no value is present for Discount, not even an explicit nil
func (o *ActItemToCreate) UnsetDiscount() {
	o.Discount.Unset()
}

func (o ActItemToCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActItemToCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productName"] = o.ProductName
	toSerialize["unitName"] = o.UnitName
	toSerialize["quantity"] = o.Quantity
	if o.Price.IsSet() {
		toSerialize["price"] = o.Price.Get()
	}
	if o.NdsRate.IsSet() {
		toSerialize["ndsRate"] = o.NdsRate.Get()
	}
	if o.Discount.IsSet() {
		toSerialize["discount"] = o.Discount.Get()
	}
	return toSerialize, nil
}

func (o *ActItemToCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"productName",
		"unitName",
		"quantity",
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

	varActItemToCreate := _ActItemToCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActItemToCreate)

	if err != nil {
		return err
	}

	*o = ActItemToCreate(varActItemToCreate)

	return err
}

type NullableActItemToCreate struct {
	value *ActItemToCreate
	isSet bool
}

func (v NullableActItemToCreate) Get() *ActItemToCreate {
	return v.value
}

func (v *NullableActItemToCreate) Set(val *ActItemToCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableActItemToCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableActItemToCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActItemToCreate(val *ActItemToCreate) *NullableActItemToCreate {
	return &NullableActItemToCreate{value: val, isSet: true}
}

func (v NullableActItemToCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActItemToCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


