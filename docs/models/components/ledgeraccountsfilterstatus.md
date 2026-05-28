# LedgerAccountsFilterStatus

Filter by account status.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.LedgerAccountsFilterStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.LedgerAccountsFilterStatus("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `LedgerAccountsFilterStatusActive`   | active                               |
| `LedgerAccountsFilterStatusInactive` | inactive                             |