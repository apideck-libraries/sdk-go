# PaymentUnit

Unit of measurement for employee compensation.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.PaymentUnitHour

// Open enum: custom values can be created with a direct type cast
custom := components.PaymentUnit("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `PaymentUnitHour`     | hour                  |
| `PaymentUnitWeek`     | week                  |
| `PaymentUnitMonth`    | month                 |
| `PaymentUnitYear`     | year                  |
| `PaymentUnitPaycheck` | paycheck              |
| `PaymentUnitOther`    | other                 |