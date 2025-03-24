// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/apideck-libraries/sdk-go/internal/utils"
)

type FiveType string

const (
	FiveTypeStr     FiveType = "str"
	FiveTypeInteger FiveType = "integer"
	FiveTypeNumber  FiveType = "number"
)

type Five struct {
	Str     *string  `queryParam:"inline"`
	Integer *int64   `queryParam:"inline"`
	Number  *float64 `queryParam:"inline"`

	Type FiveType
}

func CreateFiveStr(str string) Five {
	typ := FiveTypeStr

	return Five{
		Str:  &str,
		Type: typ,
	}
}

func CreateFiveInteger(integer int64) Five {
	typ := FiveTypeInteger

	return Five{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateFiveNumber(number float64) Five {
	typ := FiveTypeNumber

	return Five{
		Number: &number,
		Type:   typ,
	}
}

func (u *Five) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = FiveTypeStr
		return nil
	}

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = FiveTypeInteger
		return nil
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, true); err == nil {
		u.Number = &number
		u.Type = FiveTypeNumber
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Five", string(data))
}

func (u Five) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	return nil, errors.New("could not marshal union type Five: all fields are null")
}

type SimpleFormFieldOptionValueType string

const (
	SimpleFormFieldOptionValueTypeStr      SimpleFormFieldOptionValueType = "str"
	SimpleFormFieldOptionValueTypeInteger  SimpleFormFieldOptionValueType = "integer"
	SimpleFormFieldOptionValueTypeNumber   SimpleFormFieldOptionValueType = "number"
	SimpleFormFieldOptionValueTypeBoolean  SimpleFormFieldOptionValueType = "boolean"
	SimpleFormFieldOptionValueTypeArrayOf5 SimpleFormFieldOptionValueType = "arrayOf5"
)

type SimpleFormFieldOptionValue struct {
	Str      *string  `queryParam:"inline"`
	Integer  *int64   `queryParam:"inline"`
	Number   *float64 `queryParam:"inline"`
	Boolean  *bool    `queryParam:"inline"`
	ArrayOf5 []Five   `queryParam:"inline"`

	Type SimpleFormFieldOptionValueType
}

func CreateSimpleFormFieldOptionValueStr(str string) SimpleFormFieldOptionValue {
	typ := SimpleFormFieldOptionValueTypeStr

	return SimpleFormFieldOptionValue{
		Str:  &str,
		Type: typ,
	}
}

func CreateSimpleFormFieldOptionValueInteger(integer int64) SimpleFormFieldOptionValue {
	typ := SimpleFormFieldOptionValueTypeInteger

	return SimpleFormFieldOptionValue{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateSimpleFormFieldOptionValueNumber(number float64) SimpleFormFieldOptionValue {
	typ := SimpleFormFieldOptionValueTypeNumber

	return SimpleFormFieldOptionValue{
		Number: &number,
		Type:   typ,
	}
}

func CreateSimpleFormFieldOptionValueBoolean(boolean bool) SimpleFormFieldOptionValue {
	typ := SimpleFormFieldOptionValueTypeBoolean

	return SimpleFormFieldOptionValue{
		Boolean: &boolean,
		Type:    typ,
	}
}

func CreateSimpleFormFieldOptionValueArrayOf5(arrayOf5 []Five) SimpleFormFieldOptionValue {
	typ := SimpleFormFieldOptionValueTypeArrayOf5

	return SimpleFormFieldOptionValue{
		ArrayOf5: arrayOf5,
		Type:     typ,
	}
}

func (u *SimpleFormFieldOptionValue) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = SimpleFormFieldOptionValueTypeStr
		return nil
	}

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = SimpleFormFieldOptionValueTypeInteger
		return nil
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, true); err == nil {
		u.Number = &number
		u.Type = SimpleFormFieldOptionValueTypeNumber
		return nil
	}

	var boolean bool = false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, true); err == nil {
		u.Boolean = &boolean
		u.Type = SimpleFormFieldOptionValueTypeBoolean
		return nil
	}

	var arrayOf5 []Five = []Five{}
	if err := utils.UnmarshalJSON(data, &arrayOf5, "", true, true); err == nil {
		u.ArrayOf5 = arrayOf5
		u.Type = SimpleFormFieldOptionValueTypeArrayOf5
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SimpleFormFieldOptionValue", string(data))
}

func (u SimpleFormFieldOptionValue) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	if u.ArrayOf5 != nil {
		return utils.MarshalJSON(u.ArrayOf5, "", true)
	}

	return nil, errors.New("could not marshal union type SimpleFormFieldOptionValue: all fields are null")
}

type OptionType string

const (
	OptionTypeSimple OptionType = "simple"
)

func (e OptionType) ToPointer() *OptionType {
	return &e
}
func (e *OptionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "simple":
		*e = OptionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OptionType: %v", v)
	}
}

type SimpleFormFieldOption struct {
	Label      string                      `json:"label"`
	Value      *SimpleFormFieldOptionValue `json:"value,omitempty"`
	OptionType OptionType                  `json:"option_type"`
}

func (o *SimpleFormFieldOption) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *SimpleFormFieldOption) GetValue() *SimpleFormFieldOptionValue {
	if o == nil {
		return nil
	}
	return o.Value
}

func (o *SimpleFormFieldOption) GetOptionType() OptionType {
	if o == nil {
		return OptionType("")
	}
	return o.OptionType
}
