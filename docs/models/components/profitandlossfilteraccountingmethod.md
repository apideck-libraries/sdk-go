# ProfitAndLossFilterAccountingMethod

The accounting method used for the report: cash or accrual.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ProfitAndLossFilterAccountingMethodCash

// Open enum: custom values can be created with a direct type cast
custom := components.ProfitAndLossFilterAccountingMethod("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `ProfitAndLossFilterAccountingMethodCash`    | cash                                         |
| `ProfitAndLossFilterAccountingMethodAccrual` | accrual                                      |