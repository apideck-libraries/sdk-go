# CreditNoteCreateInputCreditNoteType

Whether this credit note reduces an amount owed by a customer (accounts receivable) or owed to a supplier (accounts payable). `accounts_payable_credit` support is connector-specific — most connectors only expose the accounts-receivable side. Check the connector's gotchas for known deviations.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.CreditNoteCreateInputCreditNoteTypeAccountsReceivableCredit

// Open enum: custom values can be created with a direct type cast
custom := components.CreditNoteCreateInputCreditNoteType("custom_value")
```


## Values

| Name                                                          | Value                                                         |
| ------------------------------------------------------------- | ------------------------------------------------------------- |
| `CreditNoteCreateInputCreditNoteTypeAccountsReceivableCredit` | accounts_receivable_credit                                    |
| `CreditNoteCreateInputCreditNoteTypeAccountsPayableCredit`    | accounts_payable_credit                                       |