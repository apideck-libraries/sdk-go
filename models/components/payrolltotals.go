// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// PayrollTotals - The overview of the payroll totals.
type PayrollTotals struct {
	// The total company debit for the payroll.
	CompanyDebit *float64 `json:"company_debit,omitempty"`
	// The total tax debit for the payroll.
	TaxDebit *float64 `json:"tax_debit,omitempty"`
	// The total check amount for the payroll.
	CheckAmount *float64 `json:"check_amount,omitempty"`
	// The net pay amount for the payroll.
	NetPay *float64 `json:"net_pay,omitempty"`
	// The gross pay amount for the payroll.
	GrossPay *float64 `json:"gross_pay,omitempty"`
	// The total amount of employer paid taxes for the payroll.
	EmployerTaxes *float64 `json:"employer_taxes,omitempty"`
	// The total amount of employee paid taxes for the payroll.
	EmployeeTaxes *float64 `json:"employee_taxes,omitempty"`
	// The total amount of company contributed benefits for the payroll.
	EmployerBenefitContributions *float64 `json:"employer_benefit_contributions,omitempty"`
	// The total amount of employee deducted benefits for the payroll.
	EmployeeBenefitDeductions *float64 `json:"employee_benefit_deductions,omitempty"`
}

func (o *PayrollTotals) GetCompanyDebit() *float64 {
	if o == nil {
		return nil
	}
	return o.CompanyDebit
}

func (o *PayrollTotals) GetTaxDebit() *float64 {
	if o == nil {
		return nil
	}
	return o.TaxDebit
}

func (o *PayrollTotals) GetCheckAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.CheckAmount
}

func (o *PayrollTotals) GetNetPay() *float64 {
	if o == nil {
		return nil
	}
	return o.NetPay
}

func (o *PayrollTotals) GetGrossPay() *float64 {
	if o == nil {
		return nil
	}
	return o.GrossPay
}

func (o *PayrollTotals) GetEmployerTaxes() *float64 {
	if o == nil {
		return nil
	}
	return o.EmployerTaxes
}

func (o *PayrollTotals) GetEmployeeTaxes() *float64 {
	if o == nil {
		return nil
	}
	return o.EmployeeTaxes
}

func (o *PayrollTotals) GetEmployerBenefitContributions() *float64 {
	if o == nil {
		return nil
	}
	return o.EmployerBenefitContributions
}

func (o *PayrollTotals) GetEmployeeBenefitDeductions() *float64 {
	if o == nil {
		return nil
	}
	return o.EmployeeBenefitDeductions
}
