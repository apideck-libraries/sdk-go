# LineItemType

Line Item type

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.LineItemTypeExpenseItem

// Open enum: custom values can be created with a direct type cast
custom := components.LineItemType("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `LineItemTypeExpenseItem`    | expense_item                 |
| `LineItemTypeExpenseAccount` | expense_account              |
| `LineItemTypeOther`          | other                        |