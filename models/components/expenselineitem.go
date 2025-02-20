// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ExpenseLineItem struct {
	// A unique identifier for an object.
	ID *string `json:"id,omitempty"`
	// A list of linked tracking categories.
	TrackingCategories []*LinkedTrackingCategory `json:"tracking_categories,omitempty"`
	// The unique identifier for the ledger account.
	AccountID *string `json:"account_id,omitempty"`
	// The ID of the customer this expense item is linked to.
	CustomerID *string `json:"customer_id,omitempty"`
	// The ID of the department this expense item is linked to.
	DepartmentID *string `json:"department_id,omitempty"`
	// The ID of the location this expense item is linked to.
	LocationID *string        `json:"location_id,omitempty"`
	TaxRate    *LinkedTaxRate `json:"tax_rate,omitempty"`
	// The expense line item description
	Description *string `json:"description,omitempty"`
	// The total amount of the expense line item.
	TotalAmount *float64 `json:"total_amount"`
	// Boolean that indicates if the line item is billable or not.
	Billable *bool `json:"billable,omitempty"`
}

func (o *ExpenseLineItem) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ExpenseLineItem) GetTrackingCategories() []*LinkedTrackingCategory {
	if o == nil {
		return nil
	}
	return o.TrackingCategories
}

func (o *ExpenseLineItem) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *ExpenseLineItem) GetCustomerID() *string {
	if o == nil {
		return nil
	}
	return o.CustomerID
}

func (o *ExpenseLineItem) GetDepartmentID() *string {
	if o == nil {
		return nil
	}
	return o.DepartmentID
}

func (o *ExpenseLineItem) GetLocationID() *string {
	if o == nil {
		return nil
	}
	return o.LocationID
}

func (o *ExpenseLineItem) GetTaxRate() *LinkedTaxRate {
	if o == nil {
		return nil
	}
	return o.TaxRate
}

func (o *ExpenseLineItem) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ExpenseLineItem) GetTotalAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *ExpenseLineItem) GetBillable() *bool {
	if o == nil {
		return nil
	}
	return o.Billable
}
