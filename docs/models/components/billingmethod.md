# BillingMethod

Method used for billing this project

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.BillingMethodFixedPrice

// Open enum: custom values can be created with a direct type cast
custom := components.BillingMethod("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `BillingMethodFixedPrice`       | fixed_price                     |
| `BillingMethodTimeAndMaterials` | time_and_materials              |
| `BillingMethodMilestoneBased`   | milestone_based                 |
| `BillingMethodRetainer`         | retainer                        |
| `BillingMethodNonBillable`      | non_billable                    |