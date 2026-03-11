# MessageType

Set to sms for SMS messages and mms for MMS messages.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.MessageTypeSms

// Open enum: custom values can be created with a direct type cast
custom := components.MessageType("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `MessageTypeSms` | sms              |
| `MessageTypeMms` | mms              |