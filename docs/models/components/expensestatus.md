# ExpenseStatus

Expense status

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ExpenseStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.ExpenseStatus("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `ExpenseStatusDraft`  | draft                 |
| `ExpenseStatusPosted` | posted                |
| `ExpenseStatusVoided` | voided                |