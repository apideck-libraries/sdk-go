# BillPaymentAllocationType

Type of entity this payment should be attributed to.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.BillPaymentAllocationTypeBill

// Open enum: custom values can be created with a direct type cast
custom := components.BillPaymentAllocationType("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `BillPaymentAllocationTypeBill`         | bill                                    |
| `BillPaymentAllocationTypeExpense`      | expense                                 |
| `BillPaymentAllocationTypeCreditMemo`   | credit_memo                             |
| `BillPaymentAllocationTypeOverPayment`  | over_payment                            |
| `BillPaymentAllocationTypePrePayment`   | pre_payment                             |
| `BillPaymentAllocationTypeJournalEntry` | journal_entry                           |
| `BillPaymentAllocationTypeOther`        | other                                   |