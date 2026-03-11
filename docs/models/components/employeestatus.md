# EmployeeStatus

The status of the employee.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.EmployeeStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.EmployeeStatus("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `EmployeeStatusActive`     | active                     |
| `EmployeeStatusInactive`   | inactive                   |
| `EmployeeStatusTerminated` | terminated                 |