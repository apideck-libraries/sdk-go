# BillPaymentCreateInputAllocationType

Type of entity this payment should be attributed to.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.BillPaymentCreateInputAllocationTypeBill

// Open enum: custom values can be created with a direct type cast
custom := components.BillPaymentCreateInputAllocationType("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `BillPaymentCreateInputAllocationTypeBill`         | bill                                               |
| `BillPaymentCreateInputAllocationTypeExpense`      | expense                                            |
| `BillPaymentCreateInputAllocationTypeCreditMemo`   | credit_memo                                        |
| `BillPaymentCreateInputAllocationTypeOverPayment`  | over_payment                                       |
| `BillPaymentCreateInputAllocationTypePrePayment`   | pre_payment                                        |
| `BillPaymentCreateInputAllocationTypeJournalEntry` | journal_entry                                      |
| `BillPaymentCreateInputAllocationTypeOther`        | other                                              |