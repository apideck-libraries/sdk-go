# BillCreditNoteStatus

Status of bill credit notes

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.BillCreditNoteStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.BillCreditNoteStatus("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `BillCreditNoteStatusDraft`         | draft                               |
| `BillCreditNoteStatusAuthorised`    | authorised                          |
| `BillCreditNoteStatusPosted`        | posted                              |
| `BillCreditNoteStatusPartiallyPaid` | partially_paid                      |
| `BillCreditNoteStatusPaid`          | paid                                |
| `BillCreditNoteStatusVoided`        | voided                              |
| `BillCreditNoteStatusDeleted`       | deleted                             |