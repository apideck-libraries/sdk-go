# Status

The status of the webhook.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.StatusEnabled

// Open enum: custom values can be created with a direct type cast
custom := components.Status("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `StatusEnabled`  | enabled          |
| `StatusDisabled` | disabled         |