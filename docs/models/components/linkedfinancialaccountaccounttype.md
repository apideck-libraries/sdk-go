# LinkedFinancialAccountAccountType

The type of account being referenced. Use `ledger_account` for GL accounts from the chart of accounts, or `bank_account` for bank account entities. When not specified, the connector will use its default behavior.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.LinkedFinancialAccountAccountTypeLedgerAccount

// Open enum: custom values can be created with a direct type cast
custom := components.LinkedFinancialAccountAccountType("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `LinkedFinancialAccountAccountTypeLedgerAccount` | ledger_account                                   |
| `LinkedFinancialAccountAccountTypeBankAccount`   | bank_account                                     |