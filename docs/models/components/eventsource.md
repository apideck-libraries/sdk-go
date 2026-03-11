# EventSource

Unify event source

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.EventSourceNative

// Open enum: custom values can be created with a direct type cast
custom := components.EventSource("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `EventSourceNative`  | native               |
| `EventSourceVirtual` | virtual              |