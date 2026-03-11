# InvoiceStatus

Invoice status

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.InvoiceStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.InvoiceStatus("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `InvoiceStatusDraft`         | draft                        |
| `InvoiceStatusSubmitted`     | submitted                    |
| `InvoiceStatusAuthorised`    | authorised                   |
| `InvoiceStatusPartiallyPaid` | partially_paid               |
| `InvoiceStatusPaid`          | paid                         |
| `InvoiceStatusUnpaid`        | unpaid                       |
| `InvoiceStatusVoid`          | void                         |
| `InvoiceStatusCredit`        | credit                       |
| `InvoiceStatusDeleted`       | deleted                      |
| `InvoiceStatusPosted`        | posted                       |