# ManagedVia

How the subscription is managed in the downstream.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ManagedViaManual

// Open enum: custom values can be created with a direct type cast
custom := components.ManagedVia("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `ManagedViaManual` | manual             |
| `ManagedViaAPI`    | api                |