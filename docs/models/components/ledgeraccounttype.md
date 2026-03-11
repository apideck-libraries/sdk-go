# LedgerAccountType

The type of account.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.LedgerAccountTypeAccountsPayable

// Open enum: custom values can be created with a direct type cast
custom := components.LedgerAccountType("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `LedgerAccountTypeAccountsPayable`     | accounts_payable                       |
| `LedgerAccountTypeAccountsReceivable`  | accounts_receivable                    |
| `LedgerAccountTypeBalancesheet`        | balancesheet                           |
| `LedgerAccountTypeBank`                | bank                                   |
| `LedgerAccountTypeCostsOfSales`        | costs_of_sales                         |
| `LedgerAccountTypeCreditCard`          | credit_card                            |
| `LedgerAccountTypeCurrentAsset`        | current_asset                          |
| `LedgerAccountTypeCurrentLiability`    | current_liability                      |
| `LedgerAccountTypeEquity`              | equity                                 |
| `LedgerAccountTypeExpense`             | expense                                |
| `LedgerAccountTypeFixedAsset`          | fixed_asset                            |
| `LedgerAccountTypeNonCurrentAsset`     | non_current_asset                      |
| `LedgerAccountTypeNonCurrentLiability` | non_current_liability                  |
| `LedgerAccountTypeOtherAsset`          | other_asset                            |
| `LedgerAccountTypeOtherExpense`        | other_expense                          |
| `LedgerAccountTypeOtherIncome`         | other_income                           |
| `LedgerAccountTypeOtherLiability`      | other_liability                        |
| `LedgerAccountTypeRevenue`             | revenue                                |
| `LedgerAccountTypeSales`               | sales                                  |
| `LedgerAccountTypeOther`               | other                                  |