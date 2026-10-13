# TrialBalanceFilterAccountingMethod

The accounting method used for the report: cash or accrual.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.TrialBalanceFilterAccountingMethodCash

// Open enum: custom values can be created with a direct type cast
custom := components.TrialBalanceFilterAccountingMethod("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `TrialBalanceFilterAccountingMethodCash`    | cash                                        |
| `TrialBalanceFilterAccountingMethodAccrual` | accrual                                     |