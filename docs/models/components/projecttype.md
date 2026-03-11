# ProjectType

Type or category of the project

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ProjectTypeClientProject

// Open enum: custom values can be created with a direct type cast
custom := components.ProjectType("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `ProjectTypeClientProject`       | client_project                   |
| `ProjectTypeInternalProject`     | internal_project                 |
| `ProjectTypeMaintenance`         | maintenance                      |
| `ProjectTypeResearchDevelopment` | research_development             |
| `ProjectTypeTraining`            | training                         |
| `ProjectTypeOther`               | other                            |