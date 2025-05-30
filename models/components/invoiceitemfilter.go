// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// InvoiceItemFilterInvoiceItemType - The type of invoice item, indicating whether it is an inventory item, a service, or another type.
type InvoiceItemFilterInvoiceItemType string

const (
	InvoiceItemFilterInvoiceItemTypeInventory InvoiceItemFilterInvoiceItemType = "inventory"
	InvoiceItemFilterInvoiceItemTypeService   InvoiceItemFilterInvoiceItemType = "service"
	InvoiceItemFilterInvoiceItemTypeOther     InvoiceItemFilterInvoiceItemType = "other"
)

func (e InvoiceItemFilterInvoiceItemType) ToPointer() *InvoiceItemFilterInvoiceItemType {
	return &e
}
func (e *InvoiceItemFilterInvoiceItemType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "inventory":
		fallthrough
	case "service":
		fallthrough
	case "other":
		*e = InvoiceItemFilterInvoiceItemType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InvoiceItemFilterInvoiceItemType: %v", v)
	}
}

type InvoiceItemFilter struct {
	// The type of invoice item, indicating whether it is an inventory item, a service, or another type.
	Type *InvoiceItemFilterInvoiceItemType `queryParam:"name=type"`
}

func (o *InvoiceItemFilter) GetType() *InvoiceItemFilterInvoiceItemType {
	if o == nil {
		return nil
	}
	return o.Type
}
