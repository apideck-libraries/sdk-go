// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"time"
)

type Department struct {
	// A unique identifier for an object.
	ID *string `json:"id,omitempty"`
	// Parent ID
	ParentID *string `json:"parent_id,omitempty"`
	// Department name
	Name        *string `json:"name,omitempty"`
	Code        *string `json:"code,omitempty"`
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

func (d Department) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *Department) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Department) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Department) GetParentID() *string {
	if o == nil {
		return nil
	}
	return o.ParentID
}

func (o *Department) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Department) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *Department) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Department) GetCustomMappings() *CustomMappings {
	if o == nil {
		return nil
	}
	return o.CustomMappings
}

func (o *Department) GetUpdatedBy() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedBy
}

func (o *Department) GetCreatedBy() *string {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *Department) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *Department) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Department) GetPassThrough() []PassThroughBody {
	if o == nil {
		return nil
	}
	return o.PassThrough
}
