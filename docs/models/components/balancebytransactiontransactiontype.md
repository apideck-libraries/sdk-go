# BalanceByTransactionTransactionType

Type of the transaction.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.BalanceByTransactionTransactionTypeInvoice

// Open enum: custom values can be created with a direct type cast
custom := components.BalanceByTransactionTransactionType("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `BalanceByTransactionTransactionTypeInvoice`     | invoice                                          |
| `BalanceByTransactionTransactionTypeCreditNote`  | credit_note                                      |
| `BalanceByTransactionTransactionTypeBill`        | bill                                             |
| `BalanceByTransactionTransactionTypePayment`     | payment                                          |
| `BalanceByTransactionTransactionTypeBillPayment` | bill_payment                                     |