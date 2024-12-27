// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type AgedReportFilter struct {
	// Filter by customer id
	CustomerID *string `queryParam:"name=customer_id"`
	// Filter by supplier id
	SupplierID *string `queryParam:"name=supplier_id"`
	// The cutoff date for considering transactions
	ReportAsOfDate *string `queryParam:"name=report_as_of_date"`
	// Number of periods to split the aged creditors report into
	PeriodCount *int64 `queryParam:"name=period_count"`
	// Length of each period in days
	PeriodLength *int64 `queryParam:"name=period_length"`
}

func (o *AgedReportFilter) GetCustomerID() *string {
	if o == nil {
		return nil
	}
	return o.CustomerID
}

func (o *AgedReportFilter) GetSupplierID() *string {
	if o == nil {
		return nil
	}
	return o.SupplierID
}

func (o *AgedReportFilter) GetReportAsOfDate() *string {
	if o == nil {
		return nil
	}
	return o.ReportAsOfDate
}

func (o *AgedReportFilter) GetPeriodCount() *int64 {
	if o == nil {
		return nil
	}
	return o.PeriodCount
}

func (o *AgedReportFilter) GetPeriodLength() *int64 {
	if o == nil {
		return nil
	}
	return o.PeriodLength
}
