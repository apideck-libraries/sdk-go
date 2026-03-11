# TimeOffRequestStatusStatus

The status of the time off request.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.TimeOffRequestStatusStatusRequested

// Open enum: custom values can be created with a direct type cast
custom := components.TimeOffRequestStatusStatus("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `TimeOffRequestStatusStatusRequested` | requested                             |
| `TimeOffRequestStatusStatusApproved`  | approved                              |
| `TimeOffRequestStatusStatusDeclined`  | declined                              |
| `TimeOffRequestStatusStatusCancelled` | cancelled                             |
| `TimeOffRequestStatusStatusDeleted`   | deleted                               |
| `TimeOffRequestStatusStatusOther`     | other                                 |