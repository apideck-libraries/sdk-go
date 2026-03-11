# QuoteStatus

Quote status

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.QuoteStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.QuoteStatus("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `QuoteStatusDraft`     | draft                  |
| `QuoteStatusSent`      | sent                   |
| `QuoteStatusAccepted`  | accepted               |
| `QuoteStatusRejected`  | rejected               |
| `QuoteStatusExpired`   | expired                |
| `QuoteStatusConverted` | converted              |
| `QuoteStatusVoid`      | void                   |
| `QuoteStatusDeleted`   | deleted                |