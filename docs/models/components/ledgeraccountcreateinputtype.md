# LedgerAccountCreateInputType

The type of account.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.LedgerAccountCreateInputTypeAccountsPayable

// Open enum: custom values can be created with a direct type cast
custom := components.LedgerAccountCreateInputType("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `LedgerAccountCreateInputTypeAccountsPayable`     | accounts_payable                                  |
| `LedgerAccountCreateInputTypeAccountsReceivable`  | accounts_receivable                               |
| `LedgerAccountCreateInputTypeBalancesheet`        | balancesheet                                      |
| `LedgerAccountCreateInputTypeBank`                | bank                                              |
| `LedgerAccountCreateInputTypeCostsOfSales`        | costs_of_sales                                    |
| `LedgerAccountCreateInputTypeCreditCard`          | credit_card                                       |
| `LedgerAccountCreateInputTypeCurrentAsset`        | current_asset                                     |
| `LedgerAccountCreateInputTypeCurrentLiability`    | current_liability                                 |
| `LedgerAccountCreateInputTypeEquity`              | equity                                            |
| `LedgerAccountCreateInputTypeExpense`             | expense                                           |
| `LedgerAccountCreateInputTypeFixedAsset`          | fixed_asset                                       |
| `LedgerAccountCreateInputTypeNonCurrentAsset`     | non_current_asset                                 |
| `LedgerAccountCreateInputTypeNonCurrentLiability` | non_current_liability                             |
| `LedgerAccountCreateInputTypeOtherAsset`          | other_asset                                       |
| `LedgerAccountCreateInputTypeOtherExpense`        | other_expense                                     |
| `LedgerAccountCreateInputTypeOtherIncome`         | other_income                                      |
| `LedgerAccountCreateInputTypeOtherLiability`      | other_liability                                   |
| `LedgerAccountCreateInputTypeRevenue`             | revenue                                           |
| `LedgerAccountCreateInputTypeSales`               | sales                                             |
| `LedgerAccountCreateInputTypeOther`               | other                                             |