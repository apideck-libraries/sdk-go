# AllocationType

Type of entity this payment should be attributed to.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.AllocationTypeInvoice

// Open enum: custom values can be created with a direct type cast
custom := components.AllocationType("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `AllocationTypeInvoice`      | invoice                      |
| `AllocationTypeOrder`        | order                        |
| `AllocationTypeExpense`      | expense                      |
| `AllocationTypeCreditMemo`   | credit_memo                  |
| `AllocationTypeOverPayment`  | over_payment                 |
| `AllocationTypePrePayment`   | pre_payment                  |
| `AllocationTypeJournalEntry` | journal_entry                |
| `AllocationTypeOther`        | other                        |
| `AllocationTypeBill`         | bill                         |