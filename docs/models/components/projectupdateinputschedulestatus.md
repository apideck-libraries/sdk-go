# ProjectUpdateInputScheduleStatus

Current status of project schedule compared to plan

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ProjectUpdateInputScheduleStatusAheadOfSchedule

// Open enum: custom values can be created with a direct type cast
custom := components.ProjectUpdateInputScheduleStatus("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `ProjectUpdateInputScheduleStatusAheadOfSchedule` | ahead_of_schedule                                 |
| `ProjectUpdateInputScheduleStatusOnSchedule`      | on_schedule                                       |
| `ProjectUpdateInputScheduleStatusBehindSchedule`  | behind_schedule                                   |
| `ProjectUpdateInputScheduleStatusCriticalDelay`   | critical_delay                                    |