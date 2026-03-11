# LedgerAccountClassification

The classification of account.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.LedgerAccountClassificationAsset

// Open enum: custom values can be created with a direct type cast
custom := components.LedgerAccountClassification("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `LedgerAccountClassificationAsset`        | asset                                     |
| `LedgerAccountClassificationEquity`       | equity                                    |
| `LedgerAccountClassificationExpense`      | expense                                   |
| `LedgerAccountClassificationLiability`    | liability                                 |
| `LedgerAccountClassificationRevenue`      | revenue                                   |
| `LedgerAccountClassificationIncome`       | income                                    |
| `LedgerAccountClassificationOtherIncome`  | other_income                              |
| `LedgerAccountClassificationOtherExpense` | other_expense                             |
| `LedgerAccountClassificationCostsOfSales` | costs_of_sales                            |
| `LedgerAccountClassificationOther`        | other                                     |