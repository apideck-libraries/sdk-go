// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// BankAccount1AccountType - The type of bank account.
type BankAccount1AccountType string

const (
	BankAccount1AccountTypeBankAccount BankAccount1AccountType = "bank_account"
	BankAccount1AccountTypeCreditCard  BankAccount1AccountType = "credit_card"
	BankAccount1AccountTypeOther       BankAccount1AccountType = "other"
)

func (e BankAccount1AccountType) ToPointer() *BankAccount1AccountType {
	return &e
}
func (e *BankAccount1AccountType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "bank_account":
		fallthrough
	case "credit_card":
		fallthrough
	case "other":
		*e = BankAccount1AccountType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BankAccount1AccountType: %v", v)
	}
}

type BankAccount1 struct {
	// The name of the bank
	BankName *string `json:"bank_name,omitempty"`
	// A bank account number is a number that is tied to your bank account. If you have several bank accounts, such as personal, joint, business (and so on), each account will have a different account number.
	AccountNumber *string `json:"account_number,omitempty"`
	// The name which you used in opening your bank account.
	AccountName *string `json:"account_name,omitempty"`
	// The type of bank account.
	AccountType *BankAccount1AccountType `json:"account_type,omitempty"`
	// The International Bank Account Number (IBAN).
	Iban *string `json:"iban,omitempty"`
	// The Bank Identifier Code (BIC).
	Bic *string `json:"bic,omitempty"`
	// A routing number is a nine-digit code used to identify a financial institution in the United States.
	RoutingNumber *string `json:"routing_number,omitempty"`
	// A BSB is a 6 digit numeric code used for identifying the branch of an Australian or New Zealand bank or financial institution.
	BsbNumber *string `json:"bsb_number,omitempty"`
	// A branch identifier is a unique identifier for a branch of a bank or financial institution.
	BranchIdentifier *string `json:"branch_identifier,omitempty"`
	// A bank code is a code assigned by a central bank, a bank supervisory body or a Bankers Association in a country to all its licensed member banks or financial institutions.
	BankCode *string `json:"bank_code,omitempty"`
	// Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Currency *Currency `json:"currency,omitempty"`
	// Country code according to ISO 3166-1 alpha-2.
	Country *string `json:"country,omitempty"`
}

func (o *BankAccount1) GetBankName() *string {
	if o == nil {
		return nil
	}
	return o.BankName
}

func (o *BankAccount1) GetAccountNumber() *string {
	if o == nil {
		return nil
	}
	return o.AccountNumber
}

func (o *BankAccount1) GetAccountName() *string {
	if o == nil {
		return nil
	}
	return o.AccountName
}

func (o *BankAccount1) GetAccountType() *BankAccount1AccountType {
	if o == nil {
		return nil
	}
	return o.AccountType
}

func (o *BankAccount1) GetIban() *string {
	if o == nil {
		return nil
	}
	return o.Iban
}

func (o *BankAccount1) GetBic() *string {
	if o == nil {
		return nil
	}
	return o.Bic
}

func (o *BankAccount1) GetRoutingNumber() *string {
	if o == nil {
		return nil
	}
	return o.RoutingNumber
}

func (o *BankAccount1) GetBsbNumber() *string {
	if o == nil {
		return nil
	}
	return o.BsbNumber
}

func (o *BankAccount1) GetBranchIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.BranchIdentifier
}

func (o *BankAccount1) GetBankCode() *string {
	if o == nil {
		return nil
	}
	return o.BankCode
}

func (o *BankAccount1) GetCurrency() *Currency {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *BankAccount1) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}
