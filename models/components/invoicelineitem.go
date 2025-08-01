// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/apideck-libraries/sdk-go/internal/utils"
	"time"
)

// InvoiceLineItemType - Item type
type InvoiceLineItemType string

const (
	InvoiceLineItemTypeSalesItem InvoiceLineItemType = "sales_item"
	InvoiceLineItemTypeDiscount  InvoiceLineItemType = "discount"
	InvoiceLineItemTypeInfo      InvoiceLineItemType = "info"
	InvoiceLineItemTypeSubTotal  InvoiceLineItemType = "sub_total"
	InvoiceLineItemTypeService   InvoiceLineItemType = "service"
	InvoiceLineItemTypeOther     InvoiceLineItemType = "other"
)

func (e InvoiceLineItemType) ToPointer() *InvoiceLineItemType {
	return &e
}
func (e *InvoiceLineItemType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sales_item":
		fallthrough
	case "discount":
		fallthrough
	case "info":
		fallthrough
	case "sub_total":
		fallthrough
	case "service":
		fallthrough
	case "other":
		*e = InvoiceLineItemType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InvoiceLineItemType: %v", v)
	}
}

type InvoiceLineItem struct {
	// A unique identifier for an object.
	ID *string `json:"id,omitempty"`
	// Row ID
	RowID *string `json:"row_id,omitempty"`
	// User defined item code
	Code *string `json:"code,omitempty"`
	// Line number of the resource
	LineNumber *int64 `json:"line_number,omitempty"`
	// User defined description
	Description *string `json:"description,omitempty"`
	// Item type
	Type *InvoiceLineItemType `json:"type,omitempty"`
	// Tax amount
	TaxAmount *float64 `json:"tax_amount,omitempty"`
	// Total amount of the line item
	TotalAmount *float64 `json:"total_amount,omitempty"`
	Quantity    *float64 `json:"quantity,omitempty"`
	UnitPrice   *float64 `json:"unit_price,omitempty"`
	// Description of the unit type the item is sold as, ie: kg, hour.
	UnitOfMeasure *string `json:"unit_of_measure,omitempty"`
	// Discount percentage applied to the line item when supported downstream.
	DiscountPercentage *float64 `json:"discount_percentage,omitempty"`
	// Discount amount applied to the line item when supported downstream.
	DiscountAmount *float64 `json:"discount_amount,omitempty"`
	// ID of the category of the line item
	CategoryID *string `json:"category_id,omitempty"`
	// The ID of the location
	LocationID *string `json:"location_id,omitempty"`
	// The ID of the department
	DepartmentID *string `json:"department_id,omitempty"`
	// The ID of the subsidiary
	SubsidiaryID *string `json:"subsidiary_id,omitempty"`
	// Whether the line item is prepaid
	Prepaid *bool              `json:"prepaid,omitempty"`
	Item    *LinkedInvoiceItem `json:"item,omitempty"`
	TaxRate *LinkedTaxRate     `json:"tax_rate,omitempty"`
	// A list of linked tracking categories.
	TrackingCategories []*LinkedTrackingCategory `json:"tracking_categories,omitempty"`
	LedgerAccount      *LinkedLedgerAccount      `json:"ledger_account,omitempty"`
	CustomFields       []CustomField             `json:"custom_fields,omitempty"`
	// A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	RowVersion *string `json:"row_version,omitempty"`
	// The user who last updated the object.
	UpdatedBy *string `json:"updated_by,omitempty"`
	// The user who created the object.
	CreatedBy *string `json:"created_by,omitempty"`
	// The date and time when the object was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The date and time when the object was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (i InvoiceLineItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InvoiceLineItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *InvoiceLineItem) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *InvoiceLineItem) GetRowID() *string {
	if o == nil {
		return nil
	}
	return o.RowID
}

func (o *InvoiceLineItem) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *InvoiceLineItem) GetLineNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.LineNumber
}

func (o *InvoiceLineItem) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *InvoiceLineItem) GetType() *InvoiceLineItemType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *InvoiceLineItem) GetTaxAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *InvoiceLineItem) GetTotalAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *InvoiceLineItem) GetQuantity() *float64 {
	if o == nil {
		return nil
	}
	return o.Quantity
}

func (o *InvoiceLineItem) GetUnitPrice() *float64 {
	if o == nil {
		return nil
	}
	return o.UnitPrice
}

func (o *InvoiceLineItem) GetUnitOfMeasure() *string {
	if o == nil {
		return nil
	}
	return o.UnitOfMeasure
}

func (o *InvoiceLineItem) GetDiscountPercentage() *float64 {
	if o == nil {
		return nil
	}
	return o.DiscountPercentage
}

func (o *InvoiceLineItem) GetDiscountAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.DiscountAmount
}

func (o *InvoiceLineItem) GetCategoryID() *string {
	if o == nil {
		return nil
	}
	return o.CategoryID
}

func (o *InvoiceLineItem) GetLocationID() *string {
	if o == nil {
		return nil
	}
	return o.LocationID
}

func (o *InvoiceLineItem) GetDepartmentID() *string {
	if o == nil {
		return nil
	}
	return o.DepartmentID
}

func (o *InvoiceLineItem) GetSubsidiaryID() *string {
	if o == nil {
		return nil
	}
	return o.SubsidiaryID
}

func (o *InvoiceLineItem) GetPrepaid() *bool {
	if o == nil {
		return nil
	}
	return o.Prepaid
}

func (o *InvoiceLineItem) GetItem() *LinkedInvoiceItem {
	if o == nil {
		return nil
	}
	return o.Item
}

func (o *InvoiceLineItem) GetTaxRate() *LinkedTaxRate {
	if o == nil {
		return nil
	}
	return o.TaxRate
}

func (o *InvoiceLineItem) GetTrackingCategories() []*LinkedTrackingCategory {
	if o == nil {
		return nil
	}
	return o.TrackingCategories
}

func (o *InvoiceLineItem) GetLedgerAccount() *LinkedLedgerAccount {
	if o == nil {
		return nil
	}
	return o.LedgerAccount
}

func (o *InvoiceLineItem) GetCustomFields() []CustomField {
	if o == nil {
		return nil
	}
	return o.CustomFields
}

func (o *InvoiceLineItem) GetRowVersion() *string {
	if o == nil {
		return nil
	}
	return o.RowVersion
}

func (o *InvoiceLineItem) GetUpdatedBy() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedBy
}

func (o *InvoiceLineItem) GetCreatedBy() *string {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *InvoiceLineItem) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *InvoiceLineItem) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

type InvoiceLineItemInput struct {
	// A unique identifier for an object.
	ID *string `json:"id,omitempty"`
	// Row ID
	RowID *string `json:"row_id,omitempty"`
	// User defined item code
	Code *string `json:"code,omitempty"`
	// Line number of the resource
	LineNumber *int64 `json:"line_number,omitempty"`
	// User defined description
	Description *string `json:"description,omitempty"`
	// Item type
	Type *InvoiceLineItemType `json:"type,omitempty"`
	// Tax amount
	TaxAmount *float64 `json:"tax_amount,omitempty"`
	// Total amount of the line item
	TotalAmount *float64 `json:"total_amount,omitempty"`
	Quantity    *float64 `json:"quantity,omitempty"`
	UnitPrice   *float64 `json:"unit_price,omitempty"`
	// Description of the unit type the item is sold as, ie: kg, hour.
	UnitOfMeasure *string `json:"unit_of_measure,omitempty"`
	// Discount percentage applied to the line item when supported downstream.
	DiscountPercentage *float64 `json:"discount_percentage,omitempty"`
	// Discount amount applied to the line item when supported downstream.
	DiscountAmount *float64 `json:"discount_amount,omitempty"`
	// ID of the category of the line item
	CategoryID *string `json:"category_id,omitempty"`
	// The ID of the location
	LocationID *string `json:"location_id,omitempty"`
	// The ID of the department
	DepartmentID *string `json:"department_id,omitempty"`
	// The ID of the subsidiary
	SubsidiaryID *string `json:"subsidiary_id,omitempty"`
	// Whether the line item is prepaid
	Prepaid *bool               `json:"prepaid,omitempty"`
	Item    *LinkedInvoiceItem  `json:"item,omitempty"`
	TaxRate *LinkedTaxRateInput `json:"tax_rate,omitempty"`
	// A list of linked tracking categories.
	TrackingCategories []*LinkedTrackingCategory `json:"tracking_categories,omitempty"`
	LedgerAccount      *LinkedLedgerAccountInput `json:"ledger_account,omitempty"`
	CustomFields       []CustomField             `json:"custom_fields,omitempty"`
	// A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	RowVersion *string `json:"row_version,omitempty"`
}

func (o *InvoiceLineItemInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *InvoiceLineItemInput) GetRowID() *string {
	if o == nil {
		return nil
	}
	return o.RowID
}

func (o *InvoiceLineItemInput) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *InvoiceLineItemInput) GetLineNumber() *int64 {
	if o == nil {
		return nil
	}
	return o.LineNumber
}

func (o *InvoiceLineItemInput) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *InvoiceLineItemInput) GetType() *InvoiceLineItemType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *InvoiceLineItemInput) GetTaxAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *InvoiceLineItemInput) GetTotalAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *InvoiceLineItemInput) GetQuantity() *float64 {
	if o == nil {
		return nil
	}
	return o.Quantity
}

func (o *InvoiceLineItemInput) GetUnitPrice() *float64 {
	if o == nil {
		return nil
	}
	return o.UnitPrice
}

func (o *InvoiceLineItemInput) GetUnitOfMeasure() *string {
	if o == nil {
		return nil
	}
	return o.UnitOfMeasure
}

func (o *InvoiceLineItemInput) GetDiscountPercentage() *float64 {
	if o == nil {
		return nil
	}
	return o.DiscountPercentage
}

func (o *InvoiceLineItemInput) GetDiscountAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.DiscountAmount
}

func (o *InvoiceLineItemInput) GetCategoryID() *string {
	if o == nil {
		return nil
	}
	return o.CategoryID
}

func (o *InvoiceLineItemInput) GetLocationID() *string {
	if o == nil {
		return nil
	}
	return o.LocationID
}

func (o *InvoiceLineItemInput) GetDepartmentID() *string {
	if o == nil {
		return nil
	}
	return o.DepartmentID
}

func (o *InvoiceLineItemInput) GetSubsidiaryID() *string {
	if o == nil {
		return nil
	}
	return o.SubsidiaryID
}

func (o *InvoiceLineItemInput) GetPrepaid() *bool {
	if o == nil {
		return nil
	}
	return o.Prepaid
}

func (o *InvoiceLineItemInput) GetItem() *LinkedInvoiceItem {
	if o == nil {
		return nil
	}
	return o.Item
}

func (o *InvoiceLineItemInput) GetTaxRate() *LinkedTaxRateInput {
	if o == nil {
		return nil
	}
	return o.TaxRate
}

func (o *InvoiceLineItemInput) GetTrackingCategories() []*LinkedTrackingCategory {
	if o == nil {
		return nil
	}
	return o.TrackingCategories
}

func (o *InvoiceLineItemInput) GetLedgerAccount() *LinkedLedgerAccountInput {
	if o == nil {
		return nil
	}
	return o.LedgerAccount
}

func (o *InvoiceLineItemInput) GetCustomFields() []CustomField {
	if o == nil {
		return nil
	}
	return o.CustomFields
}

func (o *InvoiceLineItemInput) GetRowVersion() *string {
	if o == nil {
		return nil
	}
	return o.RowVersion
}
