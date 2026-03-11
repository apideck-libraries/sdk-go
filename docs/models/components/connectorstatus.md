# ConnectorStatus

Status of the connector. Connectors with status live or beta are callable.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ConnectorStatusLive

// Open enum: custom values can be created with a direct type cast
custom := components.ConnectorStatus("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `ConnectorStatusLive`        | live                         |
| `ConnectorStatusBeta`        | beta                         |
| `ConnectorStatusEarlyAccess` | early-access                 |
| `ConnectorStatusDevelopment` | development                  |
| `ConnectorStatusConsidering` | considering                  |