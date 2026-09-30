# BillCreateInputStatus

Invoice status

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.BillCreateInputStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.BillCreateInputStatus("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `BillCreateInputStatusDraft`         | draft                                |
| `BillCreateInputStatusSubmitted`     | submitted                            |
| `BillCreateInputStatusAuthorised`    | authorised                           |
| `BillCreateInputStatusPartiallyPaid` | partially_paid                       |
| `BillCreateInputStatusPaid`          | paid                                 |
| `BillCreateInputStatusVoid`          | void                                 |
| `BillCreateInputStatusCredit`        | credit                               |
| `BillCreateInputStatusDeleted`       | deleted                              |
| `BillCreateInputStatusPosted`        | posted                               |