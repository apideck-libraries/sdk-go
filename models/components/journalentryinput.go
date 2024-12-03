// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"time"
)

type JournalEntryInput struct {
	// Journal entry title
	Title *string `json:"title,omitempty"`
	// Currency Exchange Rate at the time entity was recorded/generated.
	CurrencyRate *float64 `json:"currency_rate,omitempty"`
	// Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Currency *Currency `json:"currency,omitempty"`
	// The company or subsidiary id the transaction belongs to
	CompanyID *string `json:"company_id,omitempty"`
	// Requires a minimum of 2 line items that sum to 0
	LineItems []JournalEntryLineItemInput `json:"line_items,omitempty"`
	// Reference for the journal entry.
	Memo *string `json:"memo,omitempty"`
	// This is the date on which the journal entry was added. This can be different from the creation date and can also be backdated.
	PostedAt *time.Time `json:"posted_at,omitempty"`
	// Journal symbol of the entry. For example IND for indirect costs
	JournalSymbol *string `json:"journal_symbol,omitempty"`
	// The specific category of tax associated with a transaction like sales or purchase
	TaxType *string `json:"tax_type,omitempty"`
	// Applicable tax id/code override if tax is not supplied on a line item basis.
	TaxCode *string `json:"tax_code,omitempty"`
	// Journal entry number.
	Number *string `json:"number,omitempty"`
	// A list of linked tracking categories.
	TrackingCategories []LinkedTrackingCategory `json:"tracking_categories,omitempty"`
	// Accounting period
	AccountingPeriod *string `json:"accounting_period,omitempty"`
	// A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	RowVersion *string `json:"row_version,omitempty"`
	// The pass_through property allows passing service-specific, custom data or structured modifications in request body when creating or updating resources.
	PassThrough []PassThroughBody `json:"pass_through,omitempty"`
}

func (j JournalEntryInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(j, "", false)
}

func (j *JournalEntryInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &j, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *JournalEntryInput) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *JournalEntryInput) GetCurrencyRate() *float64 {
	if o == nil {
		return nil
	}
	return o.CurrencyRate
}

func (o *JournalEntryInput) GetCurrency() *Currency {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *JournalEntryInput) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *JournalEntryInput) GetLineItems() []JournalEntryLineItemInput {
	if o == nil {
		return nil
	}
	return o.LineItems
}

func (o *JournalEntryInput) GetMemo() *string {
	if o == nil {
		return nil
	}
	return o.Memo
}

func (o *JournalEntryInput) GetPostedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.PostedAt
}

func (o *JournalEntryInput) GetJournalSymbol() *string {
	if o == nil {
		return nil
	}
	return o.JournalSymbol
}

func (o *JournalEntryInput) GetTaxType() *string {
	if o == nil {
		return nil
	}
	return o.TaxType
}

func (o *JournalEntryInput) GetTaxCode() *string {
	if o == nil {
		return nil
	}
	return o.TaxCode
}

func (o *JournalEntryInput) GetNumber() *string {
	if o == nil {
		return nil
	}
	return o.Number
}

func (o *JournalEntryInput) GetTrackingCategories() []LinkedTrackingCategory {
	if o == nil {
		return nil
	}
	return o.TrackingCategories
}

func (o *JournalEntryInput) GetAccountingPeriod() *string {
	if o == nil {
		return nil
	}
	return o.AccountingPeriod
}

func (o *JournalEntryInput) GetRowVersion() *string {
	if o == nil {
		return nil
	}
	return o.RowVersion
}

func (o *JournalEntryInput) GetPassThrough() []PassThroughBody {
	if o == nil {
		return nil
	}
	return o.PassThrough
}
