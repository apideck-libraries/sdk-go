# InvoiceItemFilterInvoiceItemType

The type of invoice item, indicating whether it is an inventory item, a service, or another type.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.InvoiceItemFilterInvoiceItemTypeInventory

// Open enum: custom values can be created with a direct type cast
custom := components.InvoiceItemFilterInvoiceItemType("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `InvoiceItemFilterInvoiceItemTypeInventory` | inventory                                   |
| `InvoiceItemFilterInvoiceItemTypeService`   | service                                     |
| `InvoiceItemFilterInvoiceItemTypeOther`     | other                                       |