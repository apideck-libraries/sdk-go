# Classification

Filter by account classification.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ClassificationAsset

// Open enum: custom values can be created with a direct type cast
custom := components.Classification("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `ClassificationAsset`        | asset                        |
| `ClassificationEquity`       | equity                       |
| `ClassificationExpense`      | expense                      |
| `ClassificationLiability`    | liability                    |
| `ClassificationRevenue`      | revenue                      |
| `ClassificationIncome`       | income                       |
| `ClassificationOtherIncome`  | other_income                 |
| `ClassificationOtherExpense` | other_expense                |
| `ClassificationCostsOfSales` | costs_of_sales               |
| `ClassificationOther`        | other                        |