# ProjectUpdateInputBillingMethod

Method used for billing this project

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ProjectUpdateInputBillingMethodFixedPrice

// Open enum: custom values can be created with a direct type cast
custom := components.ProjectUpdateInputBillingMethod("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `ProjectUpdateInputBillingMethodFixedPrice`       | fixed_price                                       |
| `ProjectUpdateInputBillingMethodTimeAndMaterials` | time_and_materials                                |
| `ProjectUpdateInputBillingMethodMilestoneBased`   | milestone_based                                   |
| `ProjectUpdateInputBillingMethodRetainer`         | retainer                                          |
| `ProjectUpdateInputBillingMethodNonBillable`      | non_billable                                      |