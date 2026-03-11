# CreditNoteStatus

Status of credit notes

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.CreditNoteStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.CreditNoteStatus("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `CreditNoteStatusDraft`         | draft                           |
| `CreditNoteStatusAuthorised`    | authorised                      |
| `CreditNoteStatusPosted`        | posted                          |
| `CreditNoteStatusPartiallyPaid` | partially_paid                  |
| `CreditNoteStatusPaid`          | paid                            |
| `CreditNoteStatusVoided`        | voided                          |
| `CreditNoteStatusDeleted`       | deleted                         |