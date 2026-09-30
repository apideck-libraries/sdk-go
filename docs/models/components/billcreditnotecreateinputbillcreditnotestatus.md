# BillCreditNoteCreateInputBillCreditNoteStatus

Status of bill credit notes

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.BillCreditNoteCreateInputBillCreditNoteStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.BillCreditNoteCreateInputBillCreditNoteStatus("custom_value")
```


## Values

| Name                                                         | Value                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| `BillCreditNoteCreateInputBillCreditNoteStatusDraft`         | draft                                                        |
| `BillCreditNoteCreateInputBillCreditNoteStatusAuthorised`    | authorised                                                   |
| `BillCreditNoteCreateInputBillCreditNoteStatusPosted`        | posted                                                       |
| `BillCreditNoteCreateInputBillCreditNoteStatusPartiallyPaid` | partially_paid                                               |
| `BillCreditNoteCreateInputBillCreditNoteStatusPaid`          | paid                                                         |
| `BillCreditNoteCreateInputBillCreditNoteStatusVoided`        | voided                                                       |
| `BillCreditNoteCreateInputBillCreditNoteStatusDeleted`       | deleted                                                      |