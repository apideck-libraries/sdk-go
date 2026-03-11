# PeriodType

The type of period to include in the resource: month, quarter, year.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.PeriodTypeMonth

// Open enum: custom values can be created with a direct type cast
custom := components.PeriodType("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `PeriodTypeMonth`   | month               |
| `PeriodTypeQuarter` | quarter             |
| `PeriodTypeYear`    | year                |