# JournalEntryLineItemType

Debit entries are considered positive, and credit entries are considered negative.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.JournalEntryLineItemTypeDebit

// Open enum: custom values can be created with a direct type cast
custom := components.JournalEntryLineItemType("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `JournalEntryLineItemTypeDebit`  | debit                            |
| `JournalEntryLineItemTypeCredit` | credit                           |