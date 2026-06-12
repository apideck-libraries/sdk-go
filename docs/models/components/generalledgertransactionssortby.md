# GeneralLedgerTransactionsSortBy

The field on which to sort the General Ledger Transactions.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.GeneralLedgerTransactionsSortByPostedAt

// Open enum: custom values can be created with a direct type cast
custom := components.GeneralLedgerTransactionsSortBy("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `GeneralLedgerTransactionsSortByPostedAt`  | posted_at                                  |
| `GeneralLedgerTransactionsSortByUpdatedAt` | updated_at                                 |
| `GeneralLedgerTransactionsSortByCreatedAt` | created_at                                 |