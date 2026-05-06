# RebillStatus

Status of the rebilling process for this line item.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.RebillStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.RebillStatus("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `RebillStatusPending` | pending               |
| `RebillStatusBilled`  | billed                |
| `RebillStatusVoided`  | voided                |
| `RebillStatusOther`   | other                 |