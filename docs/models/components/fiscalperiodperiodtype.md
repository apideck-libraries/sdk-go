# FiscalPeriodPeriodType

Whether this record is a posting period, or a quarter or year that groups posting periods.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.FiscalPeriodPeriodTypePeriod

// Open enum: custom values can be created with a direct type cast
custom := components.FiscalPeriodPeriodType("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `FiscalPeriodPeriodTypePeriod`  | period                          |
| `FiscalPeriodPeriodTypeQuarter` | quarter                         |
| `FiscalPeriodPeriodTypeYear`    | year                            |