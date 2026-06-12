# GeneralLedgerTransactionLineItemType

Side of the entry. Redundant with the sign of net_amount but exposed as an explicit flag for filtering convenience.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.GeneralLedgerTransactionLineItemTypeDebit

// Open enum: custom values can be created with a direct type cast
custom := components.GeneralLedgerTransactionLineItemType("custom_value")
```


## Values

| Name                                         | Value                                        |
| -------------------------------------------- | -------------------------------------------- |
| `GeneralLedgerTransactionLineItemTypeDebit`  | debit                                        |
| `GeneralLedgerTransactionLineItemTypeCredit` | credit                                       |