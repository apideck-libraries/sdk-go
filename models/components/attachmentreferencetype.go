// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type AttachmentReferenceType string

const (
	AttachmentReferenceTypeInvoice AttachmentReferenceType = "invoice"
	AttachmentReferenceTypeBill    AttachmentReferenceType = "bill"
	AttachmentReferenceTypeExpense AttachmentReferenceType = "expense"
)

func (e AttachmentReferenceType) ToPointer() *AttachmentReferenceType {
	return &e
}
func (e *AttachmentReferenceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "invoice":
		fallthrough
	case "bill":
		fallthrough
	case "expense":
		*e = AttachmentReferenceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AttachmentReferenceType: %v", v)
	}
}
