# ConnectionStatus

Status of the connection.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ConnectionStatusLive

// Open enum: custom values can be created with a direct type cast
custom := components.ConnectionStatus("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `ConnectionStatusLive`      | live                        |
| `ConnectionStatusUpcoming`  | upcoming                    |
| `ConnectionStatusRequested` | requested                   |