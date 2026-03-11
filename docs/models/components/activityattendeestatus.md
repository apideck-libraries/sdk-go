# ActivityAttendeeStatus

Status of the attendee

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ActivityAttendeeStatusAccepted

// Open enum: custom values can be created with a direct type cast
custom := components.ActivityAttendeeStatus("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `ActivityAttendeeStatusAccepted`  | accepted                          |
| `ActivityAttendeeStatusTentative` | tentative                         |
| `ActivityAttendeeStatusDeclined`  | declined                          |