# ExpensePaymentType

The type of payment for the expense.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ExpensePaymentTypeCash

// Open enum: custom values can be created with a direct type cast
custom := components.ExpensePaymentType("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `ExpensePaymentTypeCash`       | cash                           |
| `ExpensePaymentTypeCheck`      | check                          |
| `ExpensePaymentTypeCreditCard` | credit_card                    |
| `ExpensePaymentTypeOther`      | other                          |