// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"github.com/apideck-libraries/sdk-go/types"
	"time"
)

type AgedDebtors struct {
	// The exact date and time the report was generated.
	ReportGeneratedAt *time.Time `json:"report_generated_at,omitempty"`
	// The cutoff date for transactions included in the report.
	ReportAsOfDate *types.Date `json:"report_as_of_date,omitempty"`
	// Number of aging periods shown in the report.
	PeriodCount *int64 `default:"4" json:"period_count"`
	// Length of each aging period in days.
	PeriodLength        *int64               `default:"30" json:"period_length"`
	OutstandingBalances []OutstandingBalance `json:"outstanding_balances,omitempty"`
}

func (a AgedDebtors) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AgedDebtors) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AgedDebtors) GetReportGeneratedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ReportGeneratedAt
}

func (o *AgedDebtors) GetReportAsOfDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.ReportAsOfDate
}

func (o *AgedDebtors) GetPeriodCount() *int64 {
	if o == nil {
		return nil
	}
	return o.PeriodCount
}

func (o *AgedDebtors) GetPeriodLength() *int64 {
	if o == nil {
		return nil
	}
	return o.PeriodLength
}

func (o *AgedDebtors) GetOutstandingBalances() []OutstandingBalance {
	if o == nil {
		return nil
	}
	return o.OutstandingBalances
}