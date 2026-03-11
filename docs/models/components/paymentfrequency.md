# PaymentFrequency

Frequency of employee compensation.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.PaymentFrequencyWeekly

// Open enum: custom values can be created with a direct type cast
custom := components.PaymentFrequency("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `PaymentFrequencyWeekly`   | weekly                     |
| `PaymentFrequencyBiweekly` | biweekly                   |
| `PaymentFrequencyMonthly`  | monthly                    |
| `PaymentFrequencyProRata`  | pro-rata                   |
| `PaymentFrequencyOther`    | other                      |