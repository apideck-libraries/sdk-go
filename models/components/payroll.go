// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Payroll struct {
	// A unique identifier for an object.
	ID *string `json:"id"`
	// The unique identifier of the company.
	CompanyID *string `json:"company_id,omitempty"`
	// Whether or not the payroll has been successfully processed. Note that processed payrolls cannot be updated.
	Processed *bool `json:"processed"`
	// The date the payroll was processed.
	ProcessedDate *string `json:"processed_date,omitempty"`
	// The date on which employees will be paid for the payroll.
	CheckDate *string `json:"check_date"`
	// The start date, inclusive, of the pay period.
	StartDate *string `json:"start_date"`
	// The end date, inclusive, of the pay period.
	EndDate *string `json:"end_date"`
	// The overview of the payroll totals.
	Totals *PayrollTotals `json:"totals,omitempty"`
	// An array of compensations for the payroll.
	Compensations []Compensation `json:"compensations,omitempty"`
	// When custom mappings are configured on the resource, the result is included here.
	CustomMappings *CustomMappings `json:"custom_mappings,omitempty"`
}

func (o *Payroll) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Payroll) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *Payroll) GetProcessed() *bool {
	if o == nil {
		return nil
	}
	return o.Processed
}

func (o *Payroll) GetProcessedDate() *string {
	if o == nil {
		return nil
	}
	return o.ProcessedDate
}

func (o *Payroll) GetCheckDate() *string {
	if o == nil {
		return nil
	}
	return o.CheckDate
}

func (o *Payroll) GetStartDate() *string {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *Payroll) GetEndDate() *string {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *Payroll) GetTotals() *PayrollTotals {
	if o == nil {
		return nil
	}
	return o.Totals
}

func (o *Payroll) GetCompensations() []Compensation {
	if o == nil {
		return nil
	}
	return o.Compensations
}

func (o *Payroll) GetCustomMappings() *CustomMappings {
	if o == nil {
		return nil
	}
	return o.CustomMappings
}
