// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/apideck-libraries/sdk-go/internal/utils"
)

type FormFieldOptionType string

const (
	FormFieldOptionTypeSimpleFormFieldOption FormFieldOptionType = "SimpleFormFieldOption"
	FormFieldOptionTypeFormFieldOptionGroup  FormFieldOptionType = "FormFieldOptionGroup"
)

type FormFieldOption struct {
	SimpleFormFieldOption *SimpleFormFieldOption `queryParam:"inline"`
	FormFieldOptionGroup  *FormFieldOptionGroup  `queryParam:"inline"`

	Type FormFieldOptionType
}

func CreateFormFieldOptionSimpleFormFieldOption(simpleFormFieldOption SimpleFormFieldOption) FormFieldOption {
	typ := FormFieldOptionTypeSimpleFormFieldOption

	return FormFieldOption{
		SimpleFormFieldOption: &simpleFormFieldOption,
		Type:                  typ,
	}
}

func CreateFormFieldOptionFormFieldOptionGroup(formFieldOptionGroup FormFieldOptionGroup) FormFieldOption {
	typ := FormFieldOptionTypeFormFieldOptionGroup

	return FormFieldOption{
		FormFieldOptionGroup: &formFieldOptionGroup,
		Type:                 typ,
	}
}

func (u *FormFieldOption) UnmarshalJSON(data []byte) error {

	var simpleFormFieldOption SimpleFormFieldOption = SimpleFormFieldOption{}
	if err := utils.UnmarshalJSON(data, &simpleFormFieldOption, "", true, true); err == nil {
		u.SimpleFormFieldOption = &simpleFormFieldOption
		u.Type = FormFieldOptionTypeSimpleFormFieldOption
		return nil
	}

	var formFieldOptionGroup FormFieldOptionGroup = FormFieldOptionGroup{}
	if err := utils.UnmarshalJSON(data, &formFieldOptionGroup, "", true, true); err == nil {
		u.FormFieldOptionGroup = &formFieldOptionGroup
		u.Type = FormFieldOptionTypeFormFieldOptionGroup
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for FormFieldOption", string(data))
}

func (u FormFieldOption) MarshalJSON() ([]byte, error) {
	if u.SimpleFormFieldOption != nil {
		return utils.MarshalJSON(u.SimpleFormFieldOption, "", true)
	}

	if u.FormFieldOptionGroup != nil {
		return utils.MarshalJSON(u.FormFieldOptionGroup, "", true)
	}

	return nil, errors.New("could not marshal union type FormFieldOption: all fields are null")
}
