# ExpenseType

The type of expense.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ExpenseTypeExpense

// Open enum: custom values can be created with a direct type cast
custom := components.ExpenseType("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `ExpenseTypeExpense` | expense              |
| `ExpenseTypeRefund`  | refund               |