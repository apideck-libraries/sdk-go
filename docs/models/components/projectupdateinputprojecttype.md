# ProjectUpdateInputProjectType

Type or category of the project

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ProjectUpdateInputProjectTypeClientProject

// Open enum: custom values can be created with a direct type cast
custom := components.ProjectUpdateInputProjectType("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `ProjectUpdateInputProjectTypeClientProject`       | client_project                                     |
| `ProjectUpdateInputProjectTypeInternalProject`     | internal_project                                   |
| `ProjectUpdateInputProjectTypeMaintenance`         | maintenance                                        |
| `ProjectUpdateInputProjectTypeResearchDevelopment` | research_development                               |
| `ProjectUpdateInputProjectTypeTraining`            | training                                           |
| `ProjectUpdateInputProjectTypeOther`               | other                                              |