# ScheduleStatus

Current status of project schedule compared to plan

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ScheduleStatusAheadOfSchedule

// Open enum: custom values can be created with a direct type cast
custom := components.ScheduleStatus("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `ScheduleStatusAheadOfSchedule` | ahead_of_schedule               |
| `ScheduleStatusOnSchedule`      | on_schedule                     |
| `ScheduleStatusBehindSchedule`  | behind_schedule                 |
| `ScheduleStatusCriticalDelay`   | critical_delay                  |