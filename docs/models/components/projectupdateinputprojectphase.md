# ProjectUpdateInputProjectPhase

Current phase of the project lifecycle

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ProjectUpdateInputProjectPhaseInitiation

// Open enum: custom values can be created with a direct type cast
custom := components.ProjectUpdateInputProjectPhase("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `ProjectUpdateInputProjectPhaseInitiation` | initiation                                 |
| `ProjectUpdateInputProjectPhasePlanning`   | planning                                   |
| `ProjectUpdateInputProjectPhaseExecution`  | execution                                  |
| `ProjectUpdateInputProjectPhaseMonitoring` | monitoring                                 |
| `ProjectUpdateInputProjectPhaseClosure`    | closure                                    |
| `ProjectUpdateInputProjectPhaseOther`      | other                                      |