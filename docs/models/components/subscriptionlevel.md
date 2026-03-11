# SubscriptionLevel

Received events are scoped to connection or across integration.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.SubscriptionLevelConnection

// Open enum: custom values can be created with a direct type cast
custom := components.SubscriptionLevel("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `SubscriptionLevelConnection`  | connection                     |
| `SubscriptionLevelIntegration` | integration                    |