# RefundStatus

Status of refund. Maps to: QBO (limited status), NetSuite CashRefund status, Sage Intacct state (draft/posted/voided), Zoho Books vis_state.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.RefundStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.RefundStatus("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `RefundStatusDraft`      | draft                    |
| `RefundStatusAuthorised` | authorised               |
| `RefundStatusPosted`     | posted                   |
| `RefundStatusPaid`       | paid                     |
| `RefundStatusVoided`     | voided                   |
| `RefundStatusDeleted`    | deleted                  |