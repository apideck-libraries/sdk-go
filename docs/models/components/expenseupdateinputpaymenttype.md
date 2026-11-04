# ExpenseUpdateInputPaymentType

The type of payment for the expense.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ExpenseUpdateInputPaymentTypeCash

// Open enum: custom values can be created with a direct type cast
custom := components.ExpenseUpdateInputPaymentType("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `ExpenseUpdateInputPaymentTypeCash`       | cash                                      |
| `ExpenseUpdateInputPaymentTypeCheck`      | check                                     |
| `ExpenseUpdateInputPaymentTypeCreditCard` | credit_card                               |
| `ExpenseUpdateInputPaymentTypeOther`      | other                                     |