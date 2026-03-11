# MessageStatus

Status of the delivery of the message.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.MessageStatusAccepted

// Open enum: custom values can be created with a direct type cast
custom := components.MessageStatus("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `MessageStatusAccepted`    | accepted                   |
| `MessageStatusScheduled`   | scheduled                  |
| `MessageStatusCanceled`    | canceled                   |
| `MessageStatusQueued`      | queued                     |
| `MessageStatusSending`     | sending                    |
| `MessageStatusSent`        | sent                       |
| `MessageStatusFailed`      | failed                     |
| `MessageStatusDelivered`   | delivered                  |
| `MessageStatusUndelivered` | undelivered                |
| `MessageStatusReceiving`   | receiving                  |
| `MessageStatusReceived`    | received                   |
| `MessageStatusRead`        | read                       |