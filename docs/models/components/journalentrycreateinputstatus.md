# JournalEntryCreateInputStatus

Journal entry status

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.JournalEntryCreateInputStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.JournalEntryCreateInputStatus("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `JournalEntryCreateInputStatusDraft`           | draft                                          |
| `JournalEntryCreateInputStatusPendingApproval` | pending_approval                               |
| `JournalEntryCreateInputStatusApproved`        | approved                                       |
| `JournalEntryCreateInputStatusPosted`          | posted                                         |
| `JournalEntryCreateInputStatusVoided`          | voided                                         |
| `JournalEntryCreateInputStatusRejected`        | rejected                                       |
| `JournalEntryCreateInputStatusDeleted`         | deleted                                        |
| `JournalEntryCreateInputStatusOther`           | other                                          |