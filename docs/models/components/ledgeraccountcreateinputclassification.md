# LedgerAccountCreateInputClassification

The classification of account.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.LedgerAccountCreateInputClassificationAsset

// Open enum: custom values can be created with a direct type cast
custom := components.LedgerAccountCreateInputClassification("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `LedgerAccountCreateInputClassificationAsset`        | asset                                                |
| `LedgerAccountCreateInputClassificationEquity`       | equity                                               |
| `LedgerAccountCreateInputClassificationExpense`      | expense                                              |
| `LedgerAccountCreateInputClassificationLiability`    | liability                                            |
| `LedgerAccountCreateInputClassificationRevenue`      | revenue                                              |
| `LedgerAccountCreateInputClassificationIncome`       | income                                               |
| `LedgerAccountCreateInputClassificationOtherIncome`  | other_income                                         |
| `LedgerAccountCreateInputClassificationOtherExpense` | other_expense                                        |
| `LedgerAccountCreateInputClassificationCostsOfSales` | costs_of_sales                                       |
| `LedgerAccountCreateInputClassificationOther`        | other                                                |