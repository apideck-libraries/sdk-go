# BillPaymentUpdateInputAllocationType

Type of entity this payment should be attributed to.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.BillPaymentUpdateInputAllocationTypeBill

// Open enum: custom values can be created with a direct type cast
custom := components.BillPaymentUpdateInputAllocationType("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `BillPaymentUpdateInputAllocationTypeBill`         | bill                                               |
| `BillPaymentUpdateInputAllocationTypeExpense`      | expense                                            |
| `BillPaymentUpdateInputAllocationTypeCreditMemo`   | credit_memo                                        |
| `BillPaymentUpdateInputAllocationTypeOverPayment`  | over_payment                                       |
| `BillPaymentUpdateInputAllocationTypePrePayment`   | pre_payment                                        |
| `BillPaymentUpdateInputAllocationTypeJournalEntry` | journal_entry                                      |
| `BillPaymentUpdateInputAllocationTypeOther`        | other                                              |