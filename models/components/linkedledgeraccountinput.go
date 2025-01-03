// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type LinkedLedgerAccountInput struct {
	// The unique identifier for the account.
	ID *string `json:"id,omitempty"`
	// The nominal code of the account.
	NominalCode *string `json:"nominal_code,omitempty"`
	// The code assigned to the account.
	Code *string `json:"code,omitempty"`
}

func (o *LinkedLedgerAccountInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *LinkedLedgerAccountInput) GetNominalCode() *string {
	if o == nil {
		return nil
	}
	return o.NominalCode
}

func (o *LinkedLedgerAccountInput) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}
