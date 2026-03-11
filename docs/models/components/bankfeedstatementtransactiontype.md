# BankFeedStatementTransactionType

Type of transaction.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.BankFeedStatementTransactionTypeCredit

// Open enum: custom values can be created with a direct type cast
custom := components.BankFeedStatementTransactionType("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `BankFeedStatementTransactionTypeCredit`   | credit                                     |
| `BankFeedStatementTransactionTypeDebit`    | debit                                      |
| `BankFeedStatementTransactionTypeDeposit`  | deposit                                    |
| `BankFeedStatementTransactionTypeTransfer` | transfer                                   |
| `BankFeedStatementTransactionTypePayment`  | payment                                    |
| `BankFeedStatementTransactionTypeOther`    | other                                      |