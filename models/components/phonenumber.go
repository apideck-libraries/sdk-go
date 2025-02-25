// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// PhoneNumberType - The type of phone number
type PhoneNumberType string

const (
	PhoneNumberTypePrimary      PhoneNumberType = "primary"
	PhoneNumberTypeSecondary    PhoneNumberType = "secondary"
	PhoneNumberTypeHome         PhoneNumberType = "home"
	PhoneNumberTypeWork         PhoneNumberType = "work"
	PhoneNumberTypeOffice       PhoneNumberType = "office"
	PhoneNumberTypeMobile       PhoneNumberType = "mobile"
	PhoneNumberTypeAssistant    PhoneNumberType = "assistant"
	PhoneNumberTypeFax          PhoneNumberType = "fax"
	PhoneNumberTypeDirectDialIn PhoneNumberType = "direct-dial-in"
	PhoneNumberTypePersonal     PhoneNumberType = "personal"
	PhoneNumberTypeOther        PhoneNumberType = "other"
)

func (e PhoneNumberType) ToPointer() *PhoneNumberType {
	return &e
}
func (e *PhoneNumberType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "primary":
		fallthrough
	case "secondary":
		fallthrough
	case "home":
		fallthrough
	case "work":
		fallthrough
	case "office":
		fallthrough
	case "mobile":
		fallthrough
	case "assistant":
		fallthrough
	case "fax":
		fallthrough
	case "direct-dial-in":
		fallthrough
	case "personal":
		fallthrough
	case "other":
		*e = PhoneNumberType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PhoneNumberType: %v", v)
	}
}

type PhoneNumber struct {
	// Unique identifier of the phone number
	ID *string `json:"id,omitempty"`
	// The country code of the phone number, e.g. +1
	CountryCode *string `json:"country_code,omitempty"`
	// The area code of the phone number, e.g. 323
	AreaCode *string `json:"area_code,omitempty"`
	// The phone number
	Number string `json:"number"`
	// The extension of the phone number
	Extension *string `json:"extension,omitempty"`
	// The type of phone number
	Type *PhoneNumberType `json:"type,omitempty"`
}

func (o *PhoneNumber) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PhoneNumber) GetCountryCode() *string {
	if o == nil {
		return nil
	}
	return o.CountryCode
}

func (o *PhoneNumber) GetAreaCode() *string {
	if o == nil {
		return nil
	}
	return o.AreaCode
}

func (o *PhoneNumber) GetNumber() string {
	if o == nil {
		return ""
	}
	return o.Number
}

func (o *PhoneNumber) GetExtension() *string {
	if o == nil {
		return nil
	}
	return o.Extension
}

func (o *PhoneNumber) GetType() *PhoneNumberType {
	if o == nil {
		return nil
	}
	return o.Type
}
