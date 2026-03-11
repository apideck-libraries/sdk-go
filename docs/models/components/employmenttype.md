# EmploymentType

The type of employment relationship the employee has with the organization.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.EmploymentTypeContractor

// Open enum: custom values can be created with a direct type cast
custom := components.EmploymentType("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `EmploymentTypeContractor` | contractor                 |
| `EmploymentTypeEmployee`   | employee                   |
| `EmploymentTypeFreelance`  | freelance                  |
| `EmploymentTypeTemp`       | temp                       |
| `EmploymentTypeInternship` | internship                 |
| `EmploymentTypeOther`      | other                      |