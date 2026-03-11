# InvoiceItemFilterTransactionType

The kind of transaction, indicating whether it is a sales transaction or a purchase transaction.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.InvoiceItemFilterTransactionTypeSale

// Open enum: custom values can be created with a direct type cast
custom := components.InvoiceItemFilterTransactionType("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `InvoiceItemFilterTransactionTypeSale`     | sale                                       |
| `InvoiceItemFilterTransactionTypePurchase` | purchase                                   |