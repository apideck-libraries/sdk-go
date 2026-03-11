# BankAccountFilterAccountType

Filter by account type

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.BankAccountFilterAccountTypeChecking

// Open enum: custom values can be created with a direct type cast
custom := components.BankAccountFilterAccountType("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `BankAccountFilterAccountTypeChecking`     | checking                                   |
| `BankAccountFilterAccountTypeSavings`      | savings                                    |
| `BankAccountFilterAccountTypeCreditCard`   | credit_card                                |
| `BankAccountFilterAccountTypeMoneyMarket`  | money_market                               |
| `BankAccountFilterAccountTypeLineOfCredit` | line_of_credit                             |
| `BankAccountFilterAccountTypeOther`        | other                                      |
| `BankAccountFilterAccountTypeCash`         | cash                                       |