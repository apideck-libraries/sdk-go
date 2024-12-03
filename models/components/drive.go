// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"time"
)

type Drive struct {
	// A unique identifier for an object.
	ID string `json:"id"`
	// The name of the drive
	Name string `json:"name"`
	// A description of the object.
	Description *string `json:"description,omitempty"`
	// When custom mappings are configured on the resource, the result is included here.
	CustomMappings *CustomMappings `json:"custom_mappings,omitempty"`
	// The user who last updated the object.
	UpdatedBy *string `json:"updated_by,omitempty"`
	// The user who created the object.
	CreatedBy *string `json:"created_by,omitempty"`
	// The date and time when the object was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The date and time when the object was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The pass_through property allows passing service-specific, custom data or structured modifications in request body when creating or updating resources.
	PassThrough []PassThroughBody `json:"pass_through,omitempty"`
}

func (d Drive) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *Drive) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Drive) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Drive) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Drive) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Drive) GetCustomMappings() *CustomMappings {
	if o == nil {
		return nil
	}
	return o.CustomMappings
}

func (o *Drive) GetUpdatedBy() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedBy
}

func (o *Drive) GetCreatedBy() *string {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *Drive) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *Drive) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Drive) GetPassThrough() []PassThroughBody {
	if o == nil {
		return nil
	}
	return o.PassThrough
}
