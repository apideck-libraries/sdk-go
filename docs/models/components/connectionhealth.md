# ConnectionHealth

The operational health status of the connection

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ConnectionHealthOk

// Open enum: custom values can be created with a direct type cast
custom := components.ConnectionHealth("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `ConnectionHealthOk`                  | ok                                    |
| `ConnectionHealthPendingRefresh`      | pending_refresh                       |
| `ConnectionHealthNeedsAuth`           | needs_auth                            |
| `ConnectionHealthPendingConfirmation` | pending_confirmation                  |
| `ConnectionHealthNeedsConsent`        | needs_consent                         |
| `ConnectionHealthRevoked`             | revoked                               |
| `ConnectionHealthMissingSettings`     | missing_settings                      |