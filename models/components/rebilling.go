// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// RebillStatus - Status of the rebilling process for this line item.
type RebillStatus string

const (
	RebillStatusPending RebillStatus = "pending"
	RebillStatusBilled  RebillStatus = "billed"
	RebillStatusVoided  RebillStatus = "voided"
)

func (e RebillStatus) ToPointer() *RebillStatus {
	return &e
}
func (e *RebillStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pending":
		fallthrough
	case "billed":
		fallthrough
	case "voided":
		*e = RebillStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RebillStatus: %v", v)
	}
}

// Rebilling metadata for this line item.
type Rebilling struct {
	// Whether this line item is eligible for rebilling.
	Rebillable *bool `json:"rebillable,omitempty"`
	// Status of the rebilling process for this line item.
	RebillStatus *RebillStatus `json:"rebill_status,omitempty"`
	// The ID of the transaction this line item was rebilled to.
	LinkedTransactionID *string `json:"linked_transaction_id,omitempty"`
	// The ID of the line item in the rebilled transaction.
	LinkedTransactionLineID *string `json:"linked_transaction_line_id,omitempty"`
}

func (o *Rebilling) GetRebillable() *bool {
	if o == nil {
		return nil
	}
	return o.Rebillable
}

func (o *Rebilling) GetRebillStatus() *RebillStatus {
	if o == nil {
		return nil
	}
	return o.RebillStatus
}

func (o *Rebilling) GetLinkedTransactionID() *string {
	if o == nil {
		return nil
	}
	return o.LinkedTransactionID
}

func (o *Rebilling) GetLinkedTransactionLineID() *string {
	if o == nil {
		return nil
	}
	return o.LinkedTransactionLineID
}
