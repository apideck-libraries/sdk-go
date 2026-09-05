# BillCreditNoteType

The type of credit note. A bill credit note is always an accounts payable (supplier-side) credit.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.BillCreditNoteTypeAccountsPayableCredit

// Open enum: custom values can be created with a direct type cast
custom := components.BillCreditNoteType("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `BillCreditNoteTypeAccountsPayableCredit` | accounts_payable_credit                   |