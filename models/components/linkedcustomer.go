// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// LinkedCustomer - The customer this entity is linked to.
type LinkedCustomer struct {
	// The ID of the customer this entity is linked to.
	ID *string `json:"id,omitempty"`
	// The display ID of the customer.
	DisplayID *string `json:"display_id,omitempty"`
	// The display name of the customer.
	DisplayName *string `json:"display_name,omitempty"`
	// The name of the customer. Deprecated, use display_name instead.
	//
	// Deprecated: This will be removed in a future release, please migrate away from it as soon as possible.
	Name *string `json:"name,omitempty"`
	// The company name of the customer.
	CompanyName *string `json:"company_name,omitempty"`
	// The email address of the customer.
	Email *string `json:"email,omitempty"`
}

func (o *LinkedCustomer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *LinkedCustomer) GetDisplayID() *string {
	if o == nil {
		return nil
	}
	return o.DisplayID
}

func (o *LinkedCustomer) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *LinkedCustomer) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LinkedCustomer) GetCompanyName() *string {
	if o == nil {
		return nil
	}
	return o.CompanyName
}

func (o *LinkedCustomer) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}
