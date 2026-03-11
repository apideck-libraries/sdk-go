# CreditOrDebit

Whether the amount is a credit or debit.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.CreditOrDebitCredit

// Open enum: custom values can be created with a direct type cast
custom := components.CreditOrDebit("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `CreditOrDebitCredit` | credit                |
| `CreditOrDebitDebit`  | debit                 |