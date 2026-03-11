# EmployeeJobStatus

Indicates the status of the job.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.EmployeeJobStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.EmployeeJobStatus("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `EmployeeJobStatusActive`   | active                      |
| `EmployeeJobStatusInactive` | inactive                    |
| `EmployeeJobStatusOther`    | other                       |