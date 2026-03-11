# LeavingReason

The reason because the employment ended.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.LeavingReasonDismissed

// Open enum: custom values can be created with a direct type cast
custom := components.LeavingReason("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `LeavingReasonDismissed`  | dismissed                 |
| `LeavingReasonResigned`   | resigned                  |
| `LeavingReasonRedundancy` | redundancy                |
| `LeavingReasonRetired`    | retired                   |
| `LeavingReasonOther`      | other                     |