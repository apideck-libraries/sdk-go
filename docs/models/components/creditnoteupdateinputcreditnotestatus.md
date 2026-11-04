# CreditNoteUpdateInputCreditNoteStatus

Status of credit notes

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.CreditNoteUpdateInputCreditNoteStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.CreditNoteUpdateInputCreditNoteStatus("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `CreditNoteUpdateInputCreditNoteStatusDraft`         | draft                                                |
| `CreditNoteUpdateInputCreditNoteStatusAuthorised`    | authorised                                           |
| `CreditNoteUpdateInputCreditNoteStatusPosted`        | posted                                               |
| `CreditNoteUpdateInputCreditNoteStatusPartiallyPaid` | partially_paid                                       |
| `CreditNoteUpdateInputCreditNoteStatusPaid`          | paid                                                 |
| `CreditNoteUpdateInputCreditNoteStatusVoided`        | voided                                               |
| `CreditNoteUpdateInputCreditNoteStatusDeleted`       | deleted                                              |