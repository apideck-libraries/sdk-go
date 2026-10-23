# BankFeedAccountStatus

Whether the data consumer may see this account. `revoked` hides it from the consumer without deleting it; `active` restores it. Unrelated to `feed_status`. Only connectors that list it support it; others ignore the field, so setting it there changes nothing.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.BankFeedAccountStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.BankFeedAccountStatus("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `BankFeedAccountStatusActive`  | active                         |
| `BankFeedAccountStatusRevoked` | revoked                        |