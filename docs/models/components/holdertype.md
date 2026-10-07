# HolderType

Whether the holder is an individual or a business.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.HolderTypeConsumer

// Open enum: custom values can be created with a direct type cast
custom := components.HolderType("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `HolderTypeConsumer` | consumer             |
| `HolderTypeBusiness` | business             |