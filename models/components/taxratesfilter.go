// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type TaxRatesFilter struct {
	// Boolean to describe if tax rate can be used for asset accounts
	Assets *bool `queryParam:"name=assets"`
	// Boolean to describe if tax rate can be used for equity accounts
	Equity *bool `queryParam:"name=equity"`
	// Boolean to describe if tax rate can be used for expense accounts
	Expenses *bool `queryParam:"name=expenses"`
	// Boolean to describe if tax rate can be used for liability accounts
	Liabilities *bool `queryParam:"name=liabilities"`
	// Boolean to describe if tax rate can be used for revenue accounts
	Revenue *bool `queryParam:"name=revenue"`
}

func (o *TaxRatesFilter) GetAssets() *bool {
	if o == nil {
		return nil
	}
	return o.Assets
}

func (o *TaxRatesFilter) GetEquity() *bool {
	if o == nil {
		return nil
	}
	return o.Equity
}

func (o *TaxRatesFilter) GetExpenses() *bool {
	if o == nil {
		return nil
	}
	return o.Expenses
}

func (o *TaxRatesFilter) GetLiabilities() *bool {
	if o == nil {
		return nil
	}
	return o.Liabilities
}

func (o *TaxRatesFilter) GetRevenue() *bool {
	if o == nil {
		return nil
	}
	return o.Revenue
}
