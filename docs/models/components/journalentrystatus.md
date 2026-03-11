# JournalEntryStatus

Journal entry status

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.JournalEntryStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.JournalEntryStatus("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `JournalEntryStatusDraft`           | draft                               |
| `JournalEntryStatusPendingApproval` | pending_approval                    |
| `JournalEntryStatusApproved`        | approved                            |
| `JournalEntryStatusPosted`          | posted                              |
| `JournalEntryStatusVoided`          | voided                              |
| `JournalEntryStatusRejected`        | rejected                            |
| `JournalEntryStatusDeleted`         | deleted                             |
| `JournalEntryStatusOther`           | other                               |