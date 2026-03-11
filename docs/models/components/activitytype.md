# ActivityType

The type of the activity

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ActivityTypeCall

// Open enum: custom values can be created with a direct type cast
custom := components.ActivityType("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `ActivityTypeCall`       | call                     |
| `ActivityTypeMeeting`    | meeting                  |
| `ActivityTypeEmail`      | email                    |
| `ActivityTypeNote`       | note                     |
| `ActivityTypeTask`       | task                     |
| `ActivityTypeDeadline`   | deadline                 |
| `ActivityTypeSendLetter` | send-letter              |
| `ActivityTypeSendQuote`  | send-quote               |
| `ActivityTypeOther`      | other                    |