# Health

Operational health status of the connection

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.HealthRevoked

// Open enum: custom values can be created with a direct type cast
custom := components.Health("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `HealthRevoked`         | revoked                 |
| `HealthMissingSettings` | missing_settings        |
| `HealthNeedsConsent`    | needs_consent           |
| `HealthNeedsAuth`       | needs_auth              |
| `HealthPendingRefresh`  | pending_refresh         |
| `HealthOk`              | ok                      |