# BankAccountsFilterStatus

Filter by account status

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.BankAccountsFilterStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.BankAccountsFilterStatus("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `BankAccountsFilterStatusActive`   | active                             |
| `BankAccountsFilterStatusInactive` | inactive                           |
| `BankAccountsFilterStatusClosed`   | closed                             |