# StatementStatus

The current status of the bank feed statement.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.StatementStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.StatementStatus("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `StatementStatusPending`  | pending                   |
| `StatementStatusRejected` | rejected                  |
| `StatementStatusSuccess`  | success                   |