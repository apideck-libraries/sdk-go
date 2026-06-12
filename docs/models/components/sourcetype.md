# SourceType

Filter by the originating transaction type.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.SourceTypeOther

// Open enum: custom values can be created with a direct type cast
custom := components.SourceType("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `SourceTypeOther`        | other                    |
| `SourceTypeInvoice`      | invoice                  |
| `SourceTypeBill`         | bill                     |
| `SourceTypeCreditNote`   | credit_note              |
| `SourceTypePayment`      | payment                  |
| `SourceTypeRefund`       | refund                   |
| `SourceTypeExpense`      | expense                  |
| `SourceTypeJournalEntry` | journal_entry            |
| `SourceTypePayroll`      | payroll                  |