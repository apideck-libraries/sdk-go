# EmploymentStatus

The employment status of the employee, indicating whether they are currently employed, inactive, terminated, or in another status.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.EmploymentStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.EmploymentStatus("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `EmploymentStatusActive`     | active                       |
| `EmploymentStatusInactive`   | inactive                     |
| `EmploymentStatusTerminated` | terminated                   |
| `EmploymentStatusOther`      | other                        |