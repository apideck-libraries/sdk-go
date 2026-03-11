# CategoryType

The type of the category.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.CategoryTypeSupplier

// Open enum: custom values can be created with a direct type cast
custom := components.CategoryType("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `CategoryTypeSupplier` | supplier               |
| `CategoryTypeExpense`  | expense                |
| `CategoryTypeRevenue`  | revenue                |
| `CategoryTypeCustomer` | customer               |