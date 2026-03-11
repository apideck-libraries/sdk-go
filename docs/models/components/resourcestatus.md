# ResourceStatus

Status of the resource. Resources with status live or beta are callable.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ResourceStatusLive

// Open enum: custom values can be created with a direct type cast
custom := components.ResourceStatus("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `ResourceStatusLive`        | live                        |
| `ResourceStatusBeta`        | beta                        |
| `ResourceStatusDevelopment` | development                 |
| `ResourceStatusUpcoming`    | upcoming                    |
| `ResourceStatusConsidering` | considering                 |