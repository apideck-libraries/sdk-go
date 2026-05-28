# BalanceSheetFilterAccountingMethod

The accounting method used for the report: cash or accrual.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.BalanceSheetFilterAccountingMethodCash

// Open enum: custom values can be created with a direct type cast
custom := components.BalanceSheetFilterAccountingMethod("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `BalanceSheetFilterAccountingMethodCash`    | cash                                        |
| `BalanceSheetFilterAccountingMethodAccrual` | accrual                                     |