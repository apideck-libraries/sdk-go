# AccountStatus

The status of the account.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.AccountStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.AccountStatus("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `AccountStatusActive`   | active                  |
| `AccountStatusInactive` | inactive                |
| `AccountStatusArchived` | archived                |