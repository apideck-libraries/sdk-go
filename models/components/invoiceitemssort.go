// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/apideck-libraries/sdk-go/internal/utils"
)

// InvoiceItemsSortBy - The field on which to sort the Invoice Items
type InvoiceItemsSortBy string

const (
	InvoiceItemsSortByCreatedAt InvoiceItemsSortBy = "created_at"
	InvoiceItemsSortByUpdatedAt InvoiceItemsSortBy = "updated_at"
)

func (e InvoiceItemsSortBy) ToPointer() *InvoiceItemsSortBy {
	return &e
}
func (e *InvoiceItemsSortBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "created_at":
		fallthrough
	case "updated_at":
		*e = InvoiceItemsSortBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InvoiceItemsSortBy: %v", v)
	}
}

type InvoiceItemsSort struct {
	// The field on which to sort the Invoice Items
	By *InvoiceItemsSortBy `queryParam:"name=by"`
	// The direction in which to sort the results
	Direction *SortDirection `default:"asc" queryParam:"name=direction"`
}

func (i InvoiceItemsSort) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InvoiceItemsSort) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InvoiceItemsSort) GetBy() *InvoiceItemsSortBy {
	if o == nil {
		return nil
	}
	return o.By
}

func (o *InvoiceItemsSort) GetDirection() *SortDirection {
	if o == nil {
		return nil
	}
	return o.Direction
}
