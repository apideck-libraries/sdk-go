# AmortizationType

Type of amortization

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.AmortizationTypeManual

// Open enum: custom values can be created with a direct type cast
custom := components.AmortizationType("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `AmortizationTypeManual`   | manual                     |
| `AmortizationTypeReceipt`  | receipt                    |
| `AmortizationTypeSchedule` | schedule                   |
| `AmortizationTypeOther`    | other                      |