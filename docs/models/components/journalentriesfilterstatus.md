# JournalEntriesFilterStatus

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.JournalEntriesFilterStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.JournalEntriesFilterStatus("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `JournalEntriesFilterStatusDraft`           | draft                                       |
| `JournalEntriesFilterStatusPendingApproval` | pending_approval                            |
| `JournalEntriesFilterStatusApproved`        | approved                                    |
| `JournalEntriesFilterStatusPosted`          | posted                                      |
| `JournalEntriesFilterStatusVoided`          | voided                                      |
| `JournalEntriesFilterStatusRejected`        | rejected                                    |
| `JournalEntriesFilterStatusDeleted`         | deleted                                     |
| `JournalEntriesFilterStatusOther`           | other                                       |