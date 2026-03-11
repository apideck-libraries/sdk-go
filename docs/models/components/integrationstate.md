# IntegrationState

The current state of the Integration.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.IntegrationStateDisabled

// Open enum: custom values can be created with a direct type cast
custom := components.IntegrationState("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `IntegrationStateDisabled`           | disabled                             |
| `IntegrationStateNeedsConfiguration` | needs_configuration                  |
| `IntegrationStateConfigured`         | configured                           |