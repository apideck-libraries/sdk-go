# ProjectUpdateInputProjectStatus

Current status of the project

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ProjectUpdateInputProjectStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.ProjectUpdateInputProjectStatus("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `ProjectUpdateInputProjectStatusActive`     | active                                      |
| `ProjectUpdateInputProjectStatusCompleted`  | completed                                   |
| `ProjectUpdateInputProjectStatusOnHold`     | on_hold                                     |
| `ProjectUpdateInputProjectStatusCancelled`  | cancelled                                   |
| `ProjectUpdateInputProjectStatusDraft`      | draft                                       |
| `ProjectUpdateInputProjectStatusInProgress` | in_progress                                 |
| `ProjectUpdateInputProjectStatusApproved`   | approved                                    |
| `ProjectUpdateInputProjectStatusOther`      | other                                       |