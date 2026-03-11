# LocationStatus

Based on the status some functionality is enabled or disabled.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.LocationStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.LocationStatus("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `LocationStatusActive`   | active                   |
| `LocationStatusInactive` | inactive                 |