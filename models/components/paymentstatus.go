// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// PaymentStatus - Status of payment
type PaymentStatus string

const (
	PaymentStatusAuthorised PaymentStatus = "authorised"
	PaymentStatusPaid       PaymentStatus = "paid"
	PaymentStatusVoided     PaymentStatus = "voided"
	PaymentStatusDeleted    PaymentStatus = "deleted"
)

func (e PaymentStatus) ToPointer() *PaymentStatus {
	return &e
}
func (e *PaymentStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "authorised":
		fallthrough
	case "paid":
		fallthrough
	case "voided":
		fallthrough
	case "deleted":
		*e = PaymentStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentStatus: %v", v)
	}
}
