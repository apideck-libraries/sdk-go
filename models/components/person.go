// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/types"
)

type Person struct {
	// A unique identifier for an object.
	ID *string `json:"id,omitempty"`
	// The first name of the person.
	FirstName *string `json:"first_name,omitempty"`
	// The last name of the person.
	LastName *string `json:"last_name,omitempty"`
	// Middle name of the person.
	MiddleName *string `json:"middle_name,omitempty"`
	// The gender represents the gender identity of a person.
	Gender *Gender `json:"gender,omitempty"`
	// Initials of the person
	Initials *string `json:"initials,omitempty"`
	// Date of birth
	Birthday *types.Date `json:"birthday,omitempty"`
	// Date of death
	DeceasedOn *types.Date `json:"deceased_on,omitempty"`
	// When custom mappings are configured on the resource, the result is included here.
	CustomMappings *CustomMappings `json:"custom_mappings,omitempty"`
}

func (p Person) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *Person) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Person) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Person) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *Person) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *Person) GetMiddleName() *string {
	if o == nil {
		return nil
	}
	return o.MiddleName
}

func (o *Person) GetGender() *Gender {
	if o == nil {
		return nil
	}
	return o.Gender
}

func (o *Person) GetInitials() *string {
	if o == nil {
		return nil
	}
	return o.Initials
}

func (o *Person) GetBirthday() *types.Date {
	if o == nil {
		return nil
	}
	return o.Birthday
}

func (o *Person) GetDeceasedOn() *types.Date {
	if o == nil {
		return nil
	}
	return o.DeceasedOn
}

func (o *Person) GetCustomMappings() *CustomMappings {
	if o == nil {
		return nil
	}
	return o.CustomMappings
}
