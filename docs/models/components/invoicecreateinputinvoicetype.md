# InvoiceCreateInputInvoiceType

Invoice type

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.InvoiceCreateInputInvoiceTypeStandard

// Open enum: custom values can be created with a direct type cast
custom := components.InvoiceCreateInputInvoiceType("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `InvoiceCreateInputInvoiceTypeStandard` | standard                                |
| `InvoiceCreateInputInvoiceTypeCredit`   | credit                                  |
| `InvoiceCreateInputInvoiceTypeService`  | service                                 |
| `InvoiceCreateInputInvoiceTypeProduct`  | product                                 |
| `InvoiceCreateInputInvoiceTypeSupplier` | supplier                                |
| `InvoiceCreateInputInvoiceTypeOther`    | other                                   |