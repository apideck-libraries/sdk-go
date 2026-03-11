# APIStatus

Status of the API. APIs with status live or beta are callable.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.APIStatusLive

// Open enum: custom values can be created with a direct type cast
custom := components.APIStatus("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `APIStatusLive`        | live                   |
| `APIStatusBeta`        | beta                   |
| `APIStatusDevelopment` | development            |
| `APIStatusConsidering` | considering            |