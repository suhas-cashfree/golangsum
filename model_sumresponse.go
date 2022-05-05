/*
POC of OpenAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Sumresponse struct for Sumresponse
type Sumresponse struct {
	Answer *string `json:"answer,omitempty"`
}

// NewSumresponse instantiates a new Sumresponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSumresponse() *Sumresponse {
	this := Sumresponse{}
	return &this
}

// NewSumresponseWithDefaults instantiates a new Sumresponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSumresponseWithDefaults() *Sumresponse {
	this := Sumresponse{}
	return &this
}

// GetAnswer returns the Answer field value if set, zero value otherwise.
func (o *Sumresponse) GetAnswer() string {
	if o == nil || o.Answer == nil {
		var ret string
		return ret
	}
	return *o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Sumresponse) GetAnswerOk() (*string, bool) {
	if o == nil || o.Answer == nil {
		return nil, false
	}
	return o.Answer, true
}

// HasAnswer returns a boolean if a field has been set.
func (o *Sumresponse) HasAnswer() bool {
	if o != nil && o.Answer != nil {
		return true
	}

	return false
}

// SetAnswer gets a reference to the given string and assigns it to the Answer field.
func (o *Sumresponse) SetAnswer(v string) {
	o.Answer = &v
}

func (o Sumresponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Answer != nil {
		toSerialize["answer"] = o.Answer
	}
	return json.Marshal(toSerialize)
}

type NullableSumresponse struct {
	value *Sumresponse
	isSet bool
}

func (v NullableSumresponse) Get() *Sumresponse {
	return v.value
}

func (v *NullableSumresponse) Set(val *Sumresponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSumresponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSumresponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSumresponse(val *Sumresponse) *NullableSumresponse {
	return &NullableSumresponse{value: val, isSet: true}
}

func (v NullableSumresponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSumresponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


