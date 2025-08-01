// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ExpenseLineItemInput struct {
	// A list of linked tracking categories.
	TrackingCategories []*LinkedTrackingCategory `json:"tracking_categories,omitempty"`
	// The unique identifier for the ledger account.
	AccountID *string `json:"account_id,omitempty"`
	// The ID of the customer this expense item is linked to.
	CustomerID *string `json:"customer_id,omitempty"`
	// The ID of the department
	DepartmentID *string `json:"department_id,omitempty"`
	// The ID of the location
	LocationID *string `json:"location_id,omitempty"`
	// The ID of the subsidiary
	SubsidiaryID *string             `json:"subsidiary_id,omitempty"`
	TaxRate      *LinkedTaxRateInput `json:"tax_rate,omitempty"`
	// The expense line item description
	Description *string `json:"description,omitempty"`
	// The total amount of the expense line item.
	TotalAmount *float64 `json:"total_amount"`
	// Boolean that indicates if the line item is billable or not.
	Billable *bool `json:"billable,omitempty"`
	// Line number of the resource
	LineNumber *int64 `json:"line_number,omitempty"`
	// Rebilling metadata for this line item.
	Rebilling *Rebilling `json:"rebilling,omitempty"`
}

func (o *ExpenseLineItemInput) GetTrackingCategories() []*LinkedTrackingCategory {
	if o == nil {
		return nil
	}
	return o.TrackingCategories
}

func (o *ExpenseLineItemInput) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *ExpenseLineItemInput) GetCustomerID() *string {
	if o == nil {
		return nil
	}
	return o.CustomerID
}

func (o *ExpenseLineItemInput) GetDepartmentID() *string {
	if o == nil {
		return nil
	}
	return o.DepartmentID
}

func (o *ExpenseLineItemInput) GetLocationID() *string {
	if o == nil {
		return nil
	}
	return o.LocationID
}

func (o *ExpenseLineItemInput) GetSubsidiaryID() *string {
	if o == nil {
		return nil
	}
	return o.SubsidiaryID
}

func (o *ExpenseLineItemInput) GetTaxRate() *LinkedTaxRateInput {
	if o == nil {
		return nil
	}
	return o.TaxRate
}

func (o *ExpenseLineItemInput) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ExpenseLineItemInput) GetTotalAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *ExpenseLineItemInput) GetBillable() *bool {
	if o == nil {
		return nil
	}
	return o.Billable
}

func (o *ExpenseLineItemInput) GetLineNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.LineNumber
}

func (o *ExpenseLineItemInput) GetRebilling() *Rebilling {
	if o == nil {
		return nil
	}
	return o.Rebilling
}
