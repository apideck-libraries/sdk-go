# ProjectPhase

Current phase of the project lifecycle

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ProjectPhaseInitiation

// Open enum: custom values can be created with a direct type cast
custom := components.ProjectPhase("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `ProjectPhaseInitiation` | initiation               |
| `ProjectPhasePlanning`   | planning                 |
| `ProjectPhaseExecution`  | execution                |
| `ProjectPhaseMonitoring` | monitoring               |
| `ProjectPhaseClosure`    | closure                  |
| `ProjectPhaseOther`      | other                    |