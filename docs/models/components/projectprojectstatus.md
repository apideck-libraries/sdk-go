# ProjectProjectStatus

Current status of the project

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ProjectProjectStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.ProjectProjectStatus("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `ProjectProjectStatusActive`     | active                           |
| `ProjectProjectStatusCompleted`  | completed                        |
| `ProjectProjectStatusOnHold`     | on_hold                          |
| `ProjectProjectStatusCancelled`  | cancelled                        |
| `ProjectProjectStatusDraft`      | draft                            |
| `ProjectProjectStatusInProgress` | in_progress                      |
| `ProjectProjectStatusApproved`   | approved                         |
| `ProjectProjectStatusOther`      | other                            |