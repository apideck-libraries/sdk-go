# TimeOffRequestStatus

Time off request status to filter on

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.TimeOffRequestStatusRequested

// Open enum: custom values can be created with a direct type cast
custom := components.TimeOffRequestStatus("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `TimeOffRequestStatusRequested` | requested                       |
| `TimeOffRequestStatusApproved`  | approved                        |
| `TimeOffRequestStatusDeclined`  | declined                        |
| `TimeOffRequestStatusCancelled` | cancelled                       |
| `TimeOffRequestStatusDeleted`   | deleted                         |
| `TimeOffRequestStatusOther`     | other                           |