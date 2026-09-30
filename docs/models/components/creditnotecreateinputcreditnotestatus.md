# CreditNoteCreateInputCreditNoteStatus

Status of credit notes

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.CreditNoteCreateInputCreditNoteStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.CreditNoteCreateInputCreditNoteStatus("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `CreditNoteCreateInputCreditNoteStatusDraft`         | draft                                                |
| `CreditNoteCreateInputCreditNoteStatusAuthorised`    | authorised                                           |
| `CreditNoteCreateInputCreditNoteStatusPosted`        | posted                                               |
| `CreditNoteCreateInputCreditNoteStatusPartiallyPaid` | partially_paid                                       |
| `CreditNoteCreateInputCreditNoteStatusPaid`          | paid                                                 |
| `CreditNoteCreateInputCreditNoteStatusVoided`        | voided                                               |
| `CreditNoteCreateInputCreditNoteStatusDeleted`       | deleted                                              |