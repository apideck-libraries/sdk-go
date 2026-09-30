# BillCreateInputAmortizationType

Type of amortization

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.BillCreateInputAmortizationTypeManual

// Open enum: custom values can be created with a direct type cast
custom := components.BillCreateInputAmortizationType("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `BillCreateInputAmortizationTypeManual`   | manual                                    |
| `BillCreateInputAmortizationTypeReceipt`  | receipt                                   |
| `BillCreateInputAmortizationTypeSchedule` | schedule                                  |
| `BillCreateInputAmortizationTypeOther`    | other                                     |