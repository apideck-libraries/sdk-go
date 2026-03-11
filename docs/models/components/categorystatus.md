# CategoryStatus

Based on the status some functionality is enabled or disabled.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.CategoryStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.CategoryStatus("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `CategoryStatusActive`   | active                   |
| `CategoryStatusInactive` | inactive                 |