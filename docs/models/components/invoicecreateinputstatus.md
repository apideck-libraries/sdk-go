# InvoiceCreateInputStatus

Invoice status

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.InvoiceCreateInputStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.InvoiceCreateInputStatus("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `InvoiceCreateInputStatusDraft`         | draft                                   |
| `InvoiceCreateInputStatusSubmitted`     | submitted                               |
| `InvoiceCreateInputStatusAuthorised`    | authorised                              |
| `InvoiceCreateInputStatusPartiallyPaid` | partially_paid                          |
| `InvoiceCreateInputStatusPaid`          | paid                                    |
| `InvoiceCreateInputStatusUnpaid`        | unpaid                                  |
| `InvoiceCreateInputStatusVoid`          | void                                    |
| `InvoiceCreateInputStatusCredit`        | credit                                  |
| `InvoiceCreateInputStatusDeleted`       | deleted                                 |
| `InvoiceCreateInputStatusPosted`        | posted                                  |