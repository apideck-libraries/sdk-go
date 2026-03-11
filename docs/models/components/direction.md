# Direction

The direction of the message.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.DirectionInbound

// Open enum: custom values can be created with a direct type cast
custom := components.Direction("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `DirectionInbound`       | inbound                  |
| `DirectionOutboundAPI`   | outbound-api             |
| `DirectionOutboundCall`  | outbound-call            |
| `DirectionOutboundReply` | outbound-reply           |
| `DirectionUnknown`       | unknown                  |