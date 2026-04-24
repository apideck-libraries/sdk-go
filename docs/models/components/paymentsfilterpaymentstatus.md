# PaymentsFilterPaymentStatus

Filter by payment status

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.PaymentsFilterPaymentStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.PaymentsFilterPaymentStatus("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `PaymentsFilterPaymentStatusDraft`      | draft                                   |
| `PaymentsFilterPaymentStatusAuthorised` | authorised                              |
| `PaymentsFilterPaymentStatusRejected`   | rejected                                |
| `PaymentsFilterPaymentStatusPaid`       | paid                                    |
| `PaymentsFilterPaymentStatusVoided`     | voided                                  |
| `PaymentsFilterPaymentStatusDeleted`    | deleted                                 |