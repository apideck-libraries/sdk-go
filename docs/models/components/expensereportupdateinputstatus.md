# ExpenseReportUpdateInputStatus

The status of the expense report.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ExpenseReportUpdateInputStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.ExpenseReportUpdateInputStatus("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `ExpenseReportUpdateInputStatusDraft`      | draft                                      |
| `ExpenseReportUpdateInputStatusSubmitted`  | submitted                                  |
| `ExpenseReportUpdateInputStatusApproved`   | approved                                   |
| `ExpenseReportUpdateInputStatusReimbursed` | reimbursed                                 |
| `ExpenseReportUpdateInputStatusRejected`   | rejected                                   |
| `ExpenseReportUpdateInputStatusReversed`   | reversed                                   |
| `ExpenseReportUpdateInputStatusVoided`     | voided                                     |