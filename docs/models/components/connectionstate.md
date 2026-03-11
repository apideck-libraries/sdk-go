# ConnectionState

[Connection state flow](#section/Connection-state)

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ConnectionStateAvailable

// Open enum: custom values can be created with a direct type cast
custom := components.ConnectionState("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `ConnectionStateAvailable`  | available                   |
| `ConnectionStateCallable`   | callable                    |
| `ConnectionStateAdded`      | added                       |
| `ConnectionStateAuthorized` | authorized                  |
| `ConnectionStateInvalid`    | invalid                     |