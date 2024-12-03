// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type EcommerceOrdersFilter struct {
	// Customer email address to filter on
	Email *string `queryParam:"name=email"`
	// Customer id to filter on
	CustomerID *string `queryParam:"name=customer_id"`
	// Minimum date the order was last modified
	UpdatedSince *string `queryParam:"name=updated_since"`
	// Minimum date the order was created
	CreatedSince *string `queryParam:"name=created_since"`
}

func (o *EcommerceOrdersFilter) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *EcommerceOrdersFilter) GetCustomerID() *string {
	if o == nil {
		return nil
	}
	return o.CustomerID
}

func (o *EcommerceOrdersFilter) GetUpdatedSince() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedSince
}

func (o *EcommerceOrdersFilter) GetCreatedSince() *string {
	if o == nil {
		return nil
	}
	return o.CreatedSince
}
