# ExpenseReportStatus

The status of the expense report.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ExpenseReportStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.ExpenseReportStatus("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `ExpenseReportStatusDraft`      | draft                           |
| `ExpenseReportStatusSubmitted`  | submitted                       |
| `ExpenseReportStatusApproved`   | approved                        |
| `ExpenseReportStatusReimbursed` | reimbursed                      |
| `ExpenseReportStatusRejected`   | rejected                        |
| `ExpenseReportStatusReversed`   | reversed                        |
| `ExpenseReportStatusVoided`     | voided                          |