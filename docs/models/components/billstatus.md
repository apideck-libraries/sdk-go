# BillStatus

Invoice status

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.BillStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.BillStatus("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `BillStatusDraft`         | draft                     |
| `BillStatusSubmitted`     | submitted                 |
| `BillStatusAuthorised`    | authorised                |
| `BillStatusPartiallyPaid` | partially_paid            |
| `BillStatusPaid`          | paid                      |
| `BillStatusVoid`          | void                      |
| `BillStatusCredit`        | credit                    |
| `BillStatusDeleted`       | deleted                   |
| `BillStatusPosted`        | posted                    |