# BillsFilterStatus

Filter by bill status

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.BillsFilterStatusPaid

// Open enum: custom values can be created with a direct type cast
custom := components.BillsFilterStatus("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `BillsFilterStatusPaid`          | paid                             |
| `BillsFilterStatusUnpaid`        | unpaid                           |
| `BillsFilterStatusPartiallyPaid` | partially_paid                   |