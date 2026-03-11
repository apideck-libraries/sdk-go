# DepartmentStatus

Based on the status some functionality is enabled or disabled.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.DepartmentStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.DepartmentStatus("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `DepartmentStatusActive`   | active                     |
| `DepartmentStatusInactive` | inactive                   |