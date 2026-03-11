# AccountingMethod

The accounting method used for the report: cash or accrual.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.AccountingMethodCash

// Open enum: custom values can be created with a direct type cast
custom := components.AccountingMethod("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `AccountingMethodCash`    | cash                      |
| `AccountingMethodAccrual` | accrual                   |