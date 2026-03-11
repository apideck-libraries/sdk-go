# PaymentStatus

Status of payment

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.PaymentStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.PaymentStatus("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `PaymentStatusDraft`      | draft                     |
| `PaymentStatusAuthorised` | authorised                |
| `PaymentStatusRejected`   | rejected                  |
| `PaymentStatusPaid`       | paid                      |
| `PaymentStatusVoided`     | voided                    |
| `PaymentStatusDeleted`    | deleted                   |