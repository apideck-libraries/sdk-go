// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type PayrollsFilter struct {
	// Return payrolls whose pay period is after the start date
	StartDate *string `queryParam:"name=start_date"`
	// Return payrolls whose pay period is before the end date
	EndDate *string `queryParam:"name=end_date"`
}

func (o *PayrollsFilter) GetStartDate() *string {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *PayrollsFilter) GetEndDate() *string {
	if o == nil {
		return nil
	}
	return o.EndDate
}
