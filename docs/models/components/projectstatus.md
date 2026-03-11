# ProjectStatus

Status of projects to filter by

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ProjectStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.ProjectStatus("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `ProjectStatusActive`     | active                    |
| `ProjectStatusCompleted`  | completed                 |
| `ProjectStatusOnHold`     | on_hold                   |
| `ProjectStatusCancelled`  | cancelled                 |
| `ProjectStatusDraft`      | draft                     |
| `ProjectStatusInProgress` | in_progress               |
| `ProjectStatusApproved`   | approved                  |
| `ProjectStatusOther`      | other                     |