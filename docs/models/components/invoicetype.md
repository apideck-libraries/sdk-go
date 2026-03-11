# InvoiceType

Invoice type

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.InvoiceTypeStandard

// Open enum: custom values can be created with a direct type cast
custom := components.InvoiceType("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `InvoiceTypeStandard` | standard              |
| `InvoiceTypeCredit`   | credit                |
| `InvoiceTypeService`  | service               |
| `InvoiceTypeProduct`  | product               |
| `InvoiceTypeSupplier` | supplier              |
| `InvoiceTypeOther`    | other                 |