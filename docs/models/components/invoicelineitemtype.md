# InvoiceLineItemType

Item type

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.InvoiceLineItemTypeSalesItem

// Open enum: custom values can be created with a direct type cast
custom := components.InvoiceLineItemType("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `InvoiceLineItemTypeSalesItem` | sales_item                     |
| `InvoiceLineItemTypeDiscount`  | discount                       |
| `InvoiceLineItemTypeInfo`      | info                           |
| `InvoiceLineItemTypeSubTotal`  | sub_total                      |
| `InvoiceLineItemTypeService`   | service                        |
| `InvoiceLineItemTypeOther`     | other                          |