// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"time"
)

// SupplierStatus - Supplier status
type SupplierStatus string

const (
	SupplierStatusActive             SupplierStatus = "active"
	SupplierStatusInactive           SupplierStatus = "inactive"
	SupplierStatusArchived           SupplierStatus = "archived"
	SupplierStatusGdprErasureRequest SupplierStatus = "gdpr-erasure-request"
	SupplierStatusUnknown            SupplierStatus = "unknown"
)

func (e SupplierStatus) ToPointer() *SupplierStatus {
	return &e
}
func (e *SupplierStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "active":
		fallthrough
	case "inactive":
		fallthrough
	case "archived":
		fallthrough
	case "gdpr-erasure-request":
		fallthrough
	case "unknown":
		*e = SupplierStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SupplierStatus: %v", v)
	}
}

type Supplier struct {
	// A unique identifier for an object.
	ID string `json:"id"`
	// The third-party API ID of original entity
	DownstreamID *string `json:"downstream_id,omitempty"`
	// Display ID
	DisplayID *string `json:"display_id,omitempty"`
	// Display name
	DisplayName *string `json:"display_name,omitempty"`
	// The name of the company.
	CompanyName *string `json:"company_name,omitempty"`
	// The company or subsidiary id the transaction belongs to
	CompanyID *string `json:"company_id,omitempty"`
	// The job title of the person.
	Title *string `json:"title,omitempty"`
	// The first name of the person.
	FirstName *string `json:"first_name,omitempty"`
	// Middle name of the person.
	MiddleName *string `json:"middle_name,omitempty"`
	// The last name of the person.
	LastName *string `json:"last_name,omitempty"`
	Suffix   *string `json:"suffix,omitempty"`
	// Is this an individual or business supplier
	Individual   *bool         `json:"individual,omitempty"`
	Addresses    []Address     `json:"addresses,omitempty"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers,omitempty"`
	Emails       []Email       `json:"emails,omitempty"`
	Websites     []Website     `json:"websites,omitempty"`
	BankAccounts []BankAccount `json:"bank_accounts,omitempty"`
	// Some notes about this supplier
	Notes     *string        `json:"notes,omitempty"`
	TaxRate   *LinkedTaxRate `json:"tax_rate,omitempty"`
	TaxNumber *string        `json:"tax_number,omitempty"`
	// Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Currency *Currency            `json:"currency,omitempty"`
	Account  *LinkedLedgerAccount `json:"account,omitempty"`
	// Supplier status
	Status *SupplierStatus `json:"status,omitempty"`
	// Payment method used for the transaction, such as cash, credit card, bank transfer, or check
	PaymentMethod *string `json:"payment_method,omitempty"`
	// The channel through which the transaction is processed.
	Channel *string `json:"channel,omitempty"`
	// When custom mappings are configured on the resource, the result is included here.
	CustomMappings *CustomMappings `json:"custom_mappings,omitempty"`
	// The user who last updated the object.
	UpdatedBy *string `json:"updated_by,omitempty"`
	// The user who created the object.
	CreatedBy *string `json:"created_by,omitempty"`
	// The date and time when the object was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The date and time when the object was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	RowVersion *string `json:"row_version,omitempty"`
	// The pass_through property allows passing service-specific, custom data or structured modifications in request body when creating or updating resources.
	PassThrough []PassThroughBody `json:"pass_through,omitempty"`
}

func (s Supplier) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Supplier) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Supplier) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Supplier) GetDownstreamID() *string {
	if o == nil {
		return nil
	}
	return o.DownstreamID
}

func (o *Supplier) GetDisplayID() *string {
	if o == nil {
		return nil
	}
	return o.DisplayID
}

func (o *Supplier) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *Supplier) GetCompanyName() *string {
	if o == nil {
		return nil
	}
	return o.CompanyName
}

func (o *Supplier) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *Supplier) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *Supplier) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *Supplier) GetMiddleName() *string {
	if o == nil {
		return nil
	}
	return o.MiddleName
}

func (o *Supplier) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *Supplier) GetSuffix() *string {
	if o == nil {
		return nil
	}
	return o.Suffix
}

func (o *Supplier) GetIndividual() *bool {
	if o == nil {
		return nil
	}
	return o.Individual
}

func (o *Supplier) GetAddresses() []Address {
	if o == nil {
		return nil
	}
	return o.Addresses
}

func (o *Supplier) GetPhoneNumbers() []PhoneNumber {
	if o == nil {
		return nil
	}
	return o.PhoneNumbers
}

func (o *Supplier) GetEmails() []Email {
	if o == nil {
		return nil
	}
	return o.Emails
}

func (o *Supplier) GetWebsites() []Website {
	if o == nil {
		return nil
	}
	return o.Websites
}

func (o *Supplier) GetBankAccounts() []BankAccount {
	if o == nil {
		return nil
	}
	return o.BankAccounts
}

func (o *Supplier) GetNotes() *string {
	if o == nil {
		return nil
	}
	return o.Notes
}

func (o *Supplier) GetTaxRate() *LinkedTaxRate {
	if o == nil {
		return nil
	}
	return o.TaxRate
}

func (o *Supplier) GetTaxNumber() *string {
	if o == nil {
		return nil
	}
	return o.TaxNumber
}

func (o *Supplier) GetCurrency() *Currency {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *Supplier) GetAccount() *LinkedLedgerAccount {
	if o == nil {
		return nil
	}
	return o.Account
}

func (o *Supplier) GetStatus() *SupplierStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *Supplier) GetPaymentMethod() *string {
	if o == nil {
		return nil
	}
	return o.PaymentMethod
}

func (o *Supplier) GetChannel() *string {
	if o == nil {
		return nil
	}
	return o.Channel
}

func (o *Supplier) GetCustomMappings() *CustomMappings {
	if o == nil {
		return nil
	}
	return o.CustomMappings
}

func (o *Supplier) GetUpdatedBy() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedBy
}

func (o *Supplier) GetCreatedBy() *string {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *Supplier) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *Supplier) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Supplier) GetRowVersion() *string {
	if o == nil {
		return nil
	}
	return o.RowVersion
}

func (o *Supplier) GetPassThrough() []PassThroughBody {
	if o == nil {
		return nil
	}
	return o.PassThrough
}

type SupplierInput struct {
	// Display ID
	DisplayID *string `json:"display_id,omitempty"`
	// Display name
	DisplayName *string `json:"display_name,omitempty"`
	// The name of the company.
	CompanyName *string `json:"company_name,omitempty"`
	// The company or subsidiary id the transaction belongs to
	CompanyID *string `json:"company_id,omitempty"`
	// The job title of the person.
	Title *string `json:"title,omitempty"`
	// The first name of the person.
	FirstName *string `json:"first_name,omitempty"`
	// Middle name of the person.
	MiddleName *string `json:"middle_name,omitempty"`
	// The last name of the person.
	LastName *string `json:"last_name,omitempty"`
	Suffix   *string `json:"suffix,omitempty"`
	// Is this an individual or business supplier
	Individual   *bool         `json:"individual,omitempty"`
	Addresses    []Address     `json:"addresses,omitempty"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers,omitempty"`
	Emails       []Email       `json:"emails,omitempty"`
	Websites     []Website     `json:"websites,omitempty"`
	BankAccounts []BankAccount `json:"bank_accounts,omitempty"`
	// Some notes about this supplier
	Notes     *string             `json:"notes,omitempty"`
	TaxRate   *LinkedTaxRateInput `json:"tax_rate,omitempty"`
	TaxNumber *string             `json:"tax_number,omitempty"`
	// Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Currency *Currency                 `json:"currency,omitempty"`
	Account  *LinkedLedgerAccountInput `json:"account,omitempty"`
	// Supplier status
	Status *SupplierStatus `json:"status,omitempty"`
	// Payment method used for the transaction, such as cash, credit card, bank transfer, or check
	PaymentMethod *string `json:"payment_method,omitempty"`
	// The channel through which the transaction is processed.
	Channel *string `json:"channel,omitempty"`
	// A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	RowVersion *string `json:"row_version,omitempty"`
	// The pass_through property allows passing service-specific, custom data or structured modifications in request body when creating or updating resources.
	PassThrough []PassThroughBody `json:"pass_through,omitempty"`
}

func (o *SupplierInput) GetDisplayID() *string {
	if o == nil {
		return nil
	}
	return o.DisplayID
}

func (o *SupplierInput) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *SupplierInput) GetCompanyName() *string {
	if o == nil {
		return nil
	}
	return o.CompanyName
}

func (o *SupplierInput) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *SupplierInput) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *SupplierInput) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *SupplierInput) GetMiddleName() *string {
	if o == nil {
		return nil
	}
	return o.MiddleName
}

func (o *SupplierInput) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *SupplierInput) GetSuffix() *string {
	if o == nil {
		return nil
	}
	return o.Suffix
}

func (o *SupplierInput) GetIndividual() *bool {
	if o == nil {
		return nil
	}
	return o.Individual
}

func (o *SupplierInput) GetAddresses() []Address {
	if o == nil {
		return nil
	}
	return o.Addresses
}

func (o *SupplierInput) GetPhoneNumbers() []PhoneNumber {
	if o == nil {
		return nil
	}
	return o.PhoneNumbers
}

func (o *SupplierInput) GetEmails() []Email {
	if o == nil {
		return nil
	}
	return o.Emails
}

func (o *SupplierInput) GetWebsites() []Website {
	if o == nil {
		return nil
	}
	return o.Websites
}

func (o *SupplierInput) GetBankAccounts() []BankAccount {
	if o == nil {
		return nil
	}
	return o.BankAccounts
}

func (o *SupplierInput) GetNotes() *string {
	if o == nil {
		return nil
	}
	return o.Notes
}

func (o *SupplierInput) GetTaxRate() *LinkedTaxRateInput {
	if o == nil {
		return nil
	}
	return o.TaxRate
}

func (o *SupplierInput) GetTaxNumber() *string {
	if o == nil {
		return nil
	}
	return o.TaxNumber
}

func (o *SupplierInput) GetCurrency() *Currency {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *SupplierInput) GetAccount() *LinkedLedgerAccountInput {
	if o == nil {
		return nil
	}
	return o.Account
}

func (o *SupplierInput) GetStatus() *SupplierStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *SupplierInput) GetPaymentMethod() *string {
	if o == nil {
		return nil
	}
	return o.PaymentMethod
}

func (o *SupplierInput) GetChannel() *string {
	if o == nil {
		return nil
	}
	return o.Channel
}

func (o *SupplierInput) GetRowVersion() *string {
	if o == nil {
		return nil
	}
	return o.RowVersion
}

func (o *SupplierInput) GetPassThrough() []PassThroughBody {
	if o == nil {
		return nil
	}
	return o.PassThrough
}
