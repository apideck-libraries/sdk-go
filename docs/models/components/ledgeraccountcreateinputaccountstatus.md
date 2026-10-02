# LedgerAccountCreateInputAccountStatus

The status of the account.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.LedgerAccountCreateInputAccountStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.LedgerAccountCreateInputAccountStatus("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `LedgerAccountCreateInputAccountStatusActive`   | active                                          |
| `LedgerAccountCreateInputAccountStatusInactive` | inactive                                        |
| `LedgerAccountCreateInputAccountStatusArchived` | archived                                        |